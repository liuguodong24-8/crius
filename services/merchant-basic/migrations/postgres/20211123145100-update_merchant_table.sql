
-- +migrate Up
ALTER TABLE "merchant_basic"."merchant" ADD COLUMN "logo" varchar(500);
ALTER TABLE "merchant_basic"."merchant" ADD COLUMN "user_agreement" jsonb;
COMMENT ON COLUMN "merchant_basic"."merchant"."logo" IS '商户logo';
COMMENT ON COLUMN "merchant_basic"."merchant"."user_agreement" IS '商户用户协议';

-- +migrate Down
ALTER TABLE "merchant_basic"."merchant" DROP COLUMN "logo";
ALTER TABLE "merchant_basic"."merchant" DROP COLUMN "user_agreement";