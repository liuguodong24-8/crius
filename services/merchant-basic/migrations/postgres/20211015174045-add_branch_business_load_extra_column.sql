
-- +migrate Up
ALTER TABLE "merchant_basic"."branch_business" ADD "extra" jsonb;
ALTER TABLE "merchant_basic"."branch_business" ADD "load_extra" jsonb;
-- +migrate Down
ALTER TABLE "merchant_basic"."branch_business" DROP "extra";
ALTER TABLE "merchant_basic"."branch_business" DROP "load_extra";