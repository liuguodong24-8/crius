
-- +migrate Up
ALTER TABLE "merchant_basic"."branch" DROP "deleted_at";
-- +migrate Down
ALTER TABLE "merchant_basic"."branch" ADD "deleted_at" timestamptz(6);