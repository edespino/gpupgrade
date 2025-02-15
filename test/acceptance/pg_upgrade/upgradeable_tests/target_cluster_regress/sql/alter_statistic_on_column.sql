-- Copyright (c) 2017-2022 VMware, Inc. or its affiliates
-- SPDX-License-Identifier: Apache-2.0

--------------------------------------------------------------------------------
-- Validate that the upgradeable objects are functional post-upgrade
--------------------------------------------------------------------------------

-- Check that the table has been migrated over with the custom statistics
-- target.
SELECT count(*) FROM explicitly_set_statistic_table;
SELECT attname, attstattarget from pg_attribute, pg_class where attrelid=oid and relname='explicitly_set_statistic_table' and attname='col1';
