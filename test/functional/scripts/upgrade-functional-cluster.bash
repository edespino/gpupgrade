#!/bin/bash
# Copyright (c) 2017-2023 VMware, Inc. or its affiliates
# SPDX-License-Identifier: Apache-2.0

set -eux -o pipefail

source gpupgrade_src/ci/scripts/ci-helpers.bash

MODE=${MODE:-"copy"}

export GPHOME_SOURCE=/usr/local/greenplum-db-source
export GPHOME_TARGET=/usr/local/greenplum-db-target
export MASTER_DATA_DIRECTORY=/data/gpdata/coordinator/gpseg-1
export PGPORT=5432

echo "Enabling ssh to the ccp cluster..."
tar -xzvf saved_cluster_env_files/cluster_env_files.tar.gz
cp -R cluster_env_files/.ssh /root/.ssh

echo "Performing gpupgrade initialize..."
time ssh -n cdw "
    set -eux -o pipefail

    gpupgrade initialize \
              --non-interactive \
              --target-gphome $GPHOME_TARGET \
              --source-gphome $GPHOME_SOURCE \
              --source-master-port $PGPORT \
              --mode $MODE \
              --temp-port-range 6020-6040 \
              --disk-free-ratio 0
"

echo "Performing gpupgrade execute and finalize..."
time ssh -n cdw "
    set -eux -o pipefail

    gpupgrade execute --non-interactive
    gpupgrade finalize --non-interactive

    gpupgrade apply --non-interactive --gphome "$GPHOME_SOURCE" --port "$PGPORT" --input-dir /home/gpadmin/gpupgrade --phase finalize

    gpcheckcat -A
"

echo "Upgrade successful..."
