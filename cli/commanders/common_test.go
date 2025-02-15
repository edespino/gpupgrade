// Copyright (c) 2017-2022 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package commanders

import (
	"io"
	"os"
	"sync"
	"testing"

	"golang.org/x/xerrors"

	"github.com/greenplum-db/gpupgrade/testutils/exectest"
)

func Success() {}

func FailedMain() {
	os.Stderr.WriteString("oops!")
	os.Exit(1)
}

const SuccessScriptOutput = "successfully executed data migration SQL script"

func SuccessScript() {
	os.Stdout.WriteString(SuccessScriptOutput)
	os.Exit(0)
}

func FailedSqlAlreadyExists() {
	os.Stdout.WriteString("ERROR:  language \"plpythonu\" already exists\n")
	os.Exit(1)
}

func init() {
	exectest.RegisterMains(
		Success,
		FailedMain,
		SuccessScript,
		FailedSqlAlreadyExists,
	)
}

// descriptors is a helper to redirect os.Stdout and os.Stderr and buffer the
// bytes that are written to them.
//
//    d := BufferStandardDescriptors(t)
//    defer d.Close()
//
//    // write to os.Stdout and os.Stderr
//
//    bytesOut, bytesErr := d.Collect()
//
// All errors are handled through a t.Fatalf().
type descriptors struct {
	t                  *testing.T
	wg                 sync.WaitGroup
	stdout, stderr     *os.File
	saveOut, saveErr   *os.File
	outBytes, errBytes []byte
}

func BufferStandardDescriptors(t *testing.T) *descriptors {
	d := &descriptors{t: t}

	var err error
	var rOut, rErr *os.File

	rOut, d.stdout, err = os.Pipe()
	if err != nil {
		d.t.Fatalf("opening stdout pipe: %+v", err)
	}

	rErr, d.stderr, err = os.Pipe()
	if err != nil {
		d.t.Fatalf("opening stderr pipe: %+v", err)
	}

	// Switch out the streams; they are replaced by d.Close().
	d.saveOut, d.saveErr = os.Stdout, os.Stderr
	os.Stdout, os.Stderr = d.stdout, d.stderr

	// Each stream must be read separately to avoid deadlock.
	errChan := make(chan error, 2)
	d.wg.Add(2)
	go func() {
		defer d.wg.Done()

		d.outBytes, err = io.ReadAll(rOut)
		if err != nil {
			errChan <- xerrors.Errorf("reading from stdout pipe: %w", err)
		}
	}()
	go func() {
		defer d.wg.Done()

		d.errBytes, err = io.ReadAll(rErr)
		if err != nil {
			errChan <- xerrors.Errorf("reading from stderr pipe: %w", err)
		}
	}()

	close(errChan)
	for err := range errChan {
		d.t.Fatal(err)
	}

	return d
}

// Collect drains the pipes and returns the contents of stdout and stderr. It's
// safe to call more than once.
func (d *descriptors) Collect() ([]byte, []byte) {
	// Close the write sides of the pipe so our goroutines will finish.
	if d.stdout != nil {
		err := d.stdout.Close()
		if err != nil {
			d.t.Fatalf("closing stdout pipe: %+v", err)
		}

		d.stdout = nil
	}

	if d.stderr != nil {
		err := d.stderr.Close()
		if err != nil {
			d.t.Fatalf("closing stderr pipe: %+v", err)
		}

		d.stderr = nil
	}

	d.wg.Wait()

	return d.outBytes, d.errBytes
}

// Close puts os.Stdout and os.Stderr back the way they were, after draining the
// redirected pipes if necessary.
func (d *descriptors) Close() {
	// Always make sure we've waited on the pipe contents before closing.
	// Collect() is safe to call more than once.
	d.Collect()

	os.Stdout = d.saveOut
	os.Stderr = d.saveErr
}
