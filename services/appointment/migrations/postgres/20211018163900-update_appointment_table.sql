
-- +migrate Up
ALTER TABLE "appointment"."appointment_config" ADD COLUMN "decoration_fee" varchar(500);
ALTER TABLE "appointment"."appointment_template_calendar" ADD COLUMN "theme_ids" uuid[];
ALTER TABLE "appointment"."appointment_config" ADD COLUMN "theme_keep_time" int4;
ALTER TABLE "appointment"."appointment_config" ADD COLUMN "theme_cancel_time" float4;
ALTER TABLE "appointment"."appointment_config" ADD COLUMN "theme_refund_percent_before" float4;
ALTER TABLE "appointment"."appointment_config" ADD COLUMN "theme_refund_percent_after" float4;
ALTER TABLE "appointment"."appointment" ADD COLUMN "wechat_ids" jsonb;
COMMENT ON COLUMN "appointment"."appointment_config"."decoration_fee" IS '布置费说明';
COMMENT ON COLUMN "appointment"."appointment_template_calendar"."theme_ids" IS '主题id';
COMMENT ON COLUMN "appointment"."appointment_config"."theme_keep_time" IS '主题预约到期保留时间(分钟)';
COMMENT ON COLUMN "appointment"."appointment_config"."theme_cancel_time" IS '主题预约可提前取消时间(小时)';
COMMENT ON COLUMN "appointment"."appointment_config"."theme_refund_percent_before" IS '主题预约规定取消时间前退款百分比';
COMMENT ON COLUMN "appointment"."appointment_config"."theme_refund_percent_after" IS '主题预约规定取消时间后退款百分比';
COMMENT ON COLUMN "appointment"."appointment"."wechat_ids" IS '微信id';

DROP TABLE IF EXISTS "appointment"."appointment_refund";

-- +migrate Down