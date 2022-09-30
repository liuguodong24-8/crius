
-- +migrate Up
CREATE SCHEMA IF NOT EXISTS "member_private";

-- +migrate Down
DROP SCHEMA IF EXISTS "member_private" CASCADE;
