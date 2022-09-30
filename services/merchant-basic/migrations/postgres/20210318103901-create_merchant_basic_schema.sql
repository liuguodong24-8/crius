
-- +migrate Up
CREATE SCHEMA IF NOT EXISTS "merchant_basic";

-- +migrate Down
DROP SCHEMA IF EXISTS "merchant_basic" CASCADE;
