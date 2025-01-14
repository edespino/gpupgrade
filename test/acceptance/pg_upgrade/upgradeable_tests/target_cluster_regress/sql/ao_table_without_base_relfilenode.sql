-- Copyright (c) 2017-2022 VMware, Inc. or its affiliates
-- SPDX-License-Identifier: Apache-2.0

--------------------------------------------------------------------------------
-- Validate that the upgradeable objects are functional post-upgrade
--------------------------------------------------------------------------------

-- Check that the base relfilenode of the upgraded tables exist and are empty on
-- each segment.
SELECT (pg_stat_file(pg_relation_filepath('ao_without_base_relfilenode'::regclass))).size FROM gp_dist_random('pg_class') WHERE relname='ao_without_base_relfilenode';
SELECT (pg_stat_file(pg_relation_filepath('aoco_without_base_relfilenode'::regclass))).size FROM gp_dist_random('pg_class') WHERE relname='aoco_without_base_relfilenode';

-- In GPDB 6+ changing tablespace expects the base relfilenode to be present.
-- However GPDB 5 did not make this assumption. Here we check that after
-- upgrading a GPDB 5 that did not have base relfilenode that we can still
-- change the tablespace as pg_upgrade would have created an empty base
-- relfilenode on the upgraded cluster
! mkdir /tmp/ao_table_without_base_relfilenode_tablespace;
CREATE TABLESPACE ao_table_without_base_relfilenode_tablespace LOCATION '/tmp/ao_table_without_base_relfilenode_tablespace';

ALTER TABLE ao_without_base_relfilenode SET TABLESPACE ao_table_without_base_relfilenode_tablespace;
SELECT * FROM ao_without_base_relfilenode;

ALTER TABLE aoco_without_base_relfilenode SET TABLESPACE ao_table_without_base_relfilenode_tablespace;
SELECT * FROM aoco_without_base_relfilenode;

--------------------------------------------------------------------------------
-- Cleanup
--------------------------------------------------------------------------------
DROP TABLE ao_without_base_relfilenode;
DROP TABLE aoco_without_base_relfilenode;
DROP TABLESPACE ao_table_without_base_relfilenode_tablespace;
! rm -rf /tmp/ao_table_without_base_relfilenode_tablespace;
