
-- +migrate Up
ALTER TABLE "merchant_basic"."branch_tag" ADD "load_extra" jsonb;
-- +migrate Down
ALTER TABLE "merchant_basic"."branch_tag" DROP "load_extra";