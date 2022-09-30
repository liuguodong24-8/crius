
-- +migrate Up
CREATE SCHEMA IF NOT EXISTS "message";
-- +migrate Down
DROP SCHEMA IF EXISTS "message" CASCADE;