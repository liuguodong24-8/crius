
-- +migrate Up
CREATE SCHEMA IF NOT EXISTS "member_account";

-- +migrate Down
DROP SCHEMA IF EXISTS "member_account" CASCADE;
