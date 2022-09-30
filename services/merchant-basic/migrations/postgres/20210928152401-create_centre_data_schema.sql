
-- +migrate Up
CREATE SCHEMA IF NOT EXISTS "centre_data";
-- +migrate Down
DROP SCHEMA IF EXISTS "centre_data" CASCADE;