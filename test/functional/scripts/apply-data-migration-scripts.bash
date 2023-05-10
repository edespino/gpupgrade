#! /bin/bash
# Copyright (c) 2017-2023 VMware, Inc. or its affiliates
# SPDX-License-Identifier: Apache-2.0

set -eux -o pipefail

export GPHOME_SOURCE=/usr/local/greenplum-db-source
export GPHOME_TARGET=/usr/local/greenplum-db-target
export MASTER_DATA_DIRECTORY=/data/gpdata/coordinator/gpseg-1
export PGPORT=5432

./ccp_src/scripts/setup_ssh_to_cluster.sh

echo "Running the data migration scripts on the source cluster..."
time ssh -n cdw "
    set -eux -o pipefail

    source /usr/local/greenplum-db-source/greenplum_path.sh

    gpupgrade generate --non-interactive --gphome "$GPHOME_SOURCE" --port "$PGPORT" --output-dir /home/gpadmin/gpupgrade
    gpupgrade apply    --non-interactive --gphome "$GPHOME_SOURCE" --port "$PGPORT" --input-dir /home/gpadmin/gpupgrade --phase stats
    gpupgrade apply    --non-interactive --gphome "$GPHOME_SOURCE" --port "$PGPORT" --input-dir /home/gpadmin/gpupgrade --phase initialize
"
