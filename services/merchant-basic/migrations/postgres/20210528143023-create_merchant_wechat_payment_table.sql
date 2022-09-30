
-- +migrate Up
CREATE TABLE "merchant_basic"."wechat_payment" (
    "merchant_id" uuid primary key,
    "app_id" varchar(20),
    "mch_id" varchar(20),
    "headquarters_sub_mch_id" varchar(255),
    "private_key" varchar(255),
    "cert_filename" varchar(100),
    "cert_content" bytea,
    "created_at" timestamptz(6),
    "updated_at" timestamptz(6)
)
;
COMMENT ON COLUMN "merchant_basic"."wechat_payment"."merchant_id" IS '商户ID';
COMMENT ON COLUMN "merchant_basic"."wechat_payment"."app_id" IS '微信APPID';
COMMENT ON COLUMN "merchant_basic"."wechat_payment"."mch_id" IS '微信商户ID';
COMMENT ON COLUMN "merchant_basic"."wechat_payment"."headquarters_sub_mch_id" IS '微信总部子商户ID';
COMMENT ON COLUMN "merchant_basic"."wechat_payment"."private_key" IS '微信支付KEY';
COMMENT ON COLUMN "merchant_basic"."wechat_payment"."cert_filename" IS '微信支付证书名';
COMMENT ON COLUMN "merchant_basic"."wechat_payment"."cert_content" IS '微信支付证书';
COMMENT ON TABLE "merchant_basic"."wechat_payment" IS '商户微信支付';

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."wechat_payment";