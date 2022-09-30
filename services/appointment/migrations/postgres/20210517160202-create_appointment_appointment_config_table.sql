
-- +migrate Up
CREATE TABLE IF NOT EXISTS "appointment"."appointment_config" (
  "id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "keep_time" int4,
  "remind_time" float4,
  "order_limit" int4,
  "room_num_warn" int4,
  "payment_time" int4,
  "cancel_time" float4,
  "refund_percent_before" float4,
  "refund_percent_after" float4,
  "breach_months" int4,
  "breach_total" int4,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;
COMMENT ON COLUMN "appointment"."appointment_config"."keep_time" IS '预约到期保留时间(分钟)';
COMMENT ON COLUMN "appointment"."appointment_config"."remind_time" IS '预约到期提示时间(小时)';
COMMENT ON COLUMN "appointment"."appointment_config"."order_limit" IS '同一时段订单上限';
COMMENT ON COLUMN "appointment"."appointment_config"."room_num_warn" IS '房型数量预警';
COMMENT ON COLUMN "appointment"."appointment_config"."payment_time" IS '支付倒计时(分钟)';
COMMENT ON COLUMN "appointment"."appointment_config"."cancel_time" IS '可提前取消时间(小时)';
COMMENT ON COLUMN "appointment"."appointment_config"."refund_percent_before" IS '规定取消时间前退款百分比';
COMMENT ON COLUMN "appointment"."appointment_config"."refund_percent_after" IS '规定取消时间后退款百分比';
COMMENT ON COLUMN "appointment"."appointment_config"."breach_months" IS '违约周期月';
COMMENT ON COLUMN "appointment"."appointment_config"."breach_total" IS '违约数量';
COMMENT ON TABLE "appointment"."appointment_config" IS '预约配置';

CREATE INDEX "idx_appointment_config_merchant" ON "appointment"."appointment_config" USING btree (
  "merchant_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);

ALTER TABLE "appointment"."appointment_config" ADD CONSTRAINT "appointment_config_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "appointment"."appointment_config";