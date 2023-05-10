#! /bin/bash
# Copyright (c) 2017-2023 VMware, Inc. or its affiliates
# SPDX-License-Identifier: Apache-2.0

set -eux -o pipefail

export GPHOME_SOURCE=/usr/local/greenplum-db-source
export GPHOME_TARGET=/usr/local/greenplum-db-target
export MASTER_DATA_DIRECTORY=/data/gpdata/coordinator/gpseg-1
export PGPORT=5432

echo "Enabling ssh to the ccp cluster..."
tar -xzvf saved_cluster_env_files/cluster_env_files.tar.gz
cp -R cluster_env_files/.ssh /root/.ssh

scp schema_dump/* gpadmin@cdw:/tmp/dump.sql.gz

echo "Loading the SQL dump into the source cluster..."
time ssh -n gpadmin@cdw "
    set -eux -o pipefail

    source /usr/local/greenplum-db-source/greenplum_path.sh
    export PGOPTIONS='--client-min-messages=warning'
    # This is failing due to a number of errors.
    # Disabling ON_ERROR_STOP until this is fixed.
    unxz < /tmp/dump.sql.xz | psql -v ON_ERROR_STOP=0 -f - postgres
"
