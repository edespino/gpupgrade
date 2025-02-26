// Copyright (c) 2017-2022 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package hub

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/greenplum-db/gpupgrade/greenplum"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/upgrade"
)

var RenameDirectories = upgrade.RenameDirectories

type RenameMap = map[string][]*idl.RenameDirectories

func RenameDataDirectories(agentConns []*idl.Connection, source *greenplum.Cluster, intermediate *greenplum.Cluster) error {
	src := source.CoordinatorDataDir()
	dst := intermediate.CoordinatorDataDir()
	if err := RenameDirectories(src, dst); err != nil {
		return xerrors.Errorf("renaming master data directories: %w", err)
	}

	renameMap := getRenameMap(source, intermediate)
	if err := RenameSegmentDataDirs(agentConns, renameMap); err != nil {
		return xerrors.Errorf("renaming segment data directories: %w", err)
	}

	return nil
}

// getRenameMap() returns a map of host to cluster data directories to be renamed.
// This includes renaming source to archive, and target to source. In link mode
// the mirrors have been deleted to save disk space, so exclude them from the map.
// Since the upgraded mirrors will be added later to the correct directory there
// is no need to rename target to source, so only archive the source directory.
func getRenameMap(source *greenplum.Cluster, intermediate *greenplum.Cluster) RenameMap {
	m := make(RenameMap)

	for _, seg := range source.Primaries {
		if seg.IsCoordinator() {
			continue
		}

		m[seg.Hostname] = append(m[seg.Hostname], &idl.RenameDirectories{
			Source: seg.DataDir,
			Target: intermediate.Primaries[seg.ContentID].DataDir,
		})
	}

	for _, seg := range source.Mirrors {
		m[seg.Hostname] = append(m[seg.Hostname], &idl.RenameDirectories{
			Source: seg.DataDir,
			Target: intermediate.Mirrors[seg.ContentID].DataDir,
		})
	}

	return m
}

// e.g. for source /data/dbfast1/demoDataDir0 becomes /data/dbfast1/demoDataDir0_old
// e.g. for target /data/dbfast1/demoDataDir0_123ABC becomes /data/dbfast1/demoDataDir0
func RenameSegmentDataDirs(agentConns []*idl.Connection, renames RenameMap) error {
	request := func(conn *idl.Connection) error {
		if len(renames[conn.Hostname]) == 0 {
			return nil
		}

		req := &idl.RenameDirectoriesRequest{Dirs: renames[conn.Hostname]}
		_, err := conn.AgentClient.RenameDirectories(context.Background(), req)
		return err
	}

	return ExecuteRPC(agentConns, request)
}
