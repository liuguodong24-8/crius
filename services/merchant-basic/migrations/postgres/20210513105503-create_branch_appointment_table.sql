
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."branch_appointment" (
  "branch_id" uuid,
  "open_appointment" int2 DEFAULT 1,
  "appointment_granularity" int2 DEFAULT 15,
  "vr" jsonb,
  "video" varchar[] COLLATE "pg_catalog"."default",
  "environment" varchar[] COLLATE "pg_catalog"."default",
  "meal" varchar[] COLLATE "pg_catalog"."default",
  "price" varchar[] COLLATE "pg_catalog"."default",
  "hot" bool DEFAULT false,
  "room_types" jsonb
)
;
COMMENT ON COLUMN "merchant_basic"."branch_appointment"."open_appointment" IS '是否开发预约';
COMMENT ON COLUMN "merchant_basic"."branch_appointment"."appointment_granularity" IS '预约时间粒度，分钟';
COMMENT ON COLUMN "merchant_basic"."branch_appointment"."vr" IS 'VR';
COMMENT ON COLUMN "merchant_basic"."branch_appointment"."video" IS '门店视频';
COMMENT ON COLUMN "merchant_basic"."branch_appointment"."environment" IS '环境设施';
COMMENT ON COLUMN "merchant_basic"."branch_appointment"."meal" IS '餐点饮料';
COMMENT ON COLUMN "merchant_basic"."branch_appointment"."price" IS '价目';
COMMENT ON COLUMN "merchant_basic"."branch_appointment"."hot" IS '热门门店';
COMMENT ON COLUMN "merchant_basic"."branch_appointment"."branch_id" IS '门店id';
COMMENT ON COLUMN "merchant_basic"."branch_appointment"."room_types" IS '预约数据 [{"room_type_id":"xxx","num" : 12}]';
COMMENT ON TABLE "merchant_basic"."branch_appointment" IS '门店预约配置信息';

ALTER TABLE "merchant_basic"."branch_appointment" ADD CONSTRAINT "branch_appointment_pkey" PRIMARY KEY ("branch_id");

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."branch_appointment";