#!/bin/bash
# Copyright (c) 2017-2022 VMware, Inc. or its affiliates
# SPDX-License-Identifier: Apache-2.0

set -eux -o pipefail

function run_migration_scripts_and_tests() {
    time ssh mdw '
        set -eux -o pipefail

        export GPHOME_SOURCE=/usr/local/greenplum-db-source
        export GPHOME_TARGET=/usr/local/greenplum-db-target
        source "${GPHOME_SOURCE}"/greenplum_path.sh
        export MASTER_DATA_DIRECTORY=/data/gpdata/master/gpseg-1
        export PGPORT=5432

        echo "Running data migration scripts to ensure a clean cluster..."
        gpupgrade generate --non-interactive --gphome "$GPHOME_SOURCE" --port "$PGPORT"
        gpupgrade apply    --non-interactive --gphome "$GPHOME_SOURCE" --port "$PGPORT" --phase initialize

        ./gpupgrade_src/test/acceptance/gpupgrade/revert.bats
  '
}

main() {
    echo "Enabling ssh to cluster..."
    ./ccp_src/scripts/setup_ssh_to_cluster.sh

    echo "Installing gpupgrade_src on mdw..."
    scp -rpq gpupgrade_src gpadmin@mdw:/home/gpadmin

    echo "Installing BATS..."
    rsync --archive bats centos@mdw:
    ssh centos@mdw sudo ./bats/install.sh /usr/local

    echo "Running data migration scripts and tests..."
    run_migration_scripts_and_tests
}

main
