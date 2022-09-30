
-- +migrate Up
CREATE TABLE "appointment"."appointment_extend" (
  "appointment_id" uuid NOT NULL,
  "package_id" uuid,
  "packages" jsonb,
  "decoration" varchar(500) COLLATE "pg_catalog"."default",
  "staffing" varchar(500) COLLATE "pg_catalog"."default",
  "custom_configs" jsonb,
  "theme_id" uuid,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "package_name" varchar(50) COLLATE "pg_catalog"."default",
  "refunding_at" timestamptz(6),
  "refunded_at" timestamptz(6),
  "refund_amount" int8,
  "refund_id" uuid,
  "refund_response" json,
  "share_message" varchar(500) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "appointment"."appointment_extend"."package_id" IS '套餐id';
COMMENT ON COLUMN "appointment"."appointment_extend"."packages" IS '[{"id":"xxxxxxx", "category":"xxxx"}] category为product(商品)或者package(套餐) 商品或套餐';
COMMENT ON COLUMN "appointment"."appointment_extend"."decoration" IS '装饰布置';
COMMENT ON COLUMN "appointment"."appointment_extend"."staffing" IS '人员配置';
COMMENT ON COLUMN "appointment"."appointment_extend"."custom_configs" IS '[{"name":"xxxx", "config":"xxxx"}] 自定义配置';
COMMENT ON COLUMN "appointment"."appointment_extend"."theme_id" IS '主题id';
COMMENT ON COLUMN "appointment"."appointment_extend"."package_name" IS '套餐名称';
COMMENT ON COLUMN "appointment"."appointment_extend"."refunding_at" IS '退款时间';
COMMENT ON COLUMN "appointment"."appointment_extend"."refunded_at" IS '退款成功时间';
COMMENT ON COLUMN "appointment"."appointment_extend"."refund_amount" IS '退款金额';
COMMENT ON COLUMN "appointment"."appointment_extend"."refund_id" IS '交易退款id';
COMMENT ON COLUMN "appointment"."appointment_extend"."refund_response" IS '退款返回json';
COMMENT ON COLUMN "appointment"."appointment_extend"."share_message" IS '留言消息';

-- +migrate Down
DROP TABLE IF EXISTS "appointment"."appointment_extend";