
-- +migrate Up
CREATE SCHEMA IF NOT EXISTS "payment";

-- +migrate Down
DROP SCHEMA IF EXISTS "payment" CASCADE;
