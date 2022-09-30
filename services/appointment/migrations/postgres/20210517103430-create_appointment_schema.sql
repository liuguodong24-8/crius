
-- +migrate Up
CREATE SCHEMA IF NOT EXISTS "appointment";

-- +migrate Down
DROP SCHEMA IF EXISTS "appointment" CASCADE;
