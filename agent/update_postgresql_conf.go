// Copyright (c) 2017-2021 VMware, Inc. or its affiliates
// SPDX-License-Identifier: Apache-2.0

package agent

import (
	"context"

	"github.com/greenplum-db/gp-common-go-libs/gplog"

	"github.com/greenplum-db/gpupgrade/hub"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/utils/errorlist"
)

func (s *Server) UpdatePostgresqlConf(ctx context.Context, in *idl.UpdatePostgresqlConfRequest) (*idl.UpdatePostgresqlConfReply, error) {
	gplog.Info("agent received request to update postgresql.conf")

	var errs error
	for _, opt := range in.GetOptions() {
		err := hub.UpdatePostgresqlConf(opt.GetPath(), int(opt.GetOldPort()), int(opt.GetNewPort()))
		if err != nil {
			errs = errorlist.Append(errs, err)
		}
	}

	return &idl.UpdatePostgresqlConfReply{}, errs
}