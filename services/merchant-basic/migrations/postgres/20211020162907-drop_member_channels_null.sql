
-- +migrate Up
ALTER TABLE merchant_basic."member" ALTER channels DROP NOT NULL, ALTER first_channel DROP NOT NULL;
-- +migrate Down
ALTER TABLE "merchant_basic"."member" ALTER COLUMN "first_channel" SET NOT NULL, ALTER COLUMN "channels" SET NOT NULL;