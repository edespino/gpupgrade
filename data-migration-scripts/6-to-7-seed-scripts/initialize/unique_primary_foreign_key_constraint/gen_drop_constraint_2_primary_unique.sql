-- Copyright (c) 2017-2022 VMware, Inc. or its affiliates
-- SPDX-License-Identifier: Apache-2.0

-- Generate a script to drop unique/primary key constraints.

SELECT
   'ALTER TABLE ' || pg_catalog.quote_ident(n.nspname) || '.' || pg_catalog.quote_ident(cc.relname) || ' DROP CONSTRAINT ' || pg_catalog.quote_ident(conname) || ' CASCADE;'
FROM
   pg_constraint con
   JOIN
      pg_depend dep
      ON (refclassid, classid, objsubid) =
      (
         'pg_constraint'::regclass,
         'pg_class'::regclass,
         0
      )
      AND refobjid = con.oid
      AND deptype = 'i'
      AND contype IN
      (
         'u',
         'p'
      )
   JOIN
      pg_class c
      ON objid = c.oid
      AND relkind = 'i'
   JOIN
      pg_class cc
      ON cc.oid = con.conrelid
   JOIN
      pg_namespace n
      ON (n.oid = cc.relnamespace);
