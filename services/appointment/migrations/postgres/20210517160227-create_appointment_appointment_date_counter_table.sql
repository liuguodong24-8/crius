
-- +migrate Up
CREATE TABLE IF NOT EXISTS "appointment"."appointment_date_counter" (
  "id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "branch_id" uuid NOT NULL,
  "appointment_date" date,
  "appointment_time" timestamptz(6),
  "room_group_id" uuid,
  "number" int4,
  "way" int2,
  "appoint_num" int4
)
;
COMMENT ON COLUMN "appointment"."appointment_date_counter"."branch_id" IS '门店';
COMMENT ON COLUMN "appointment"."appointment_date_counter"."appointment_date" IS '营业时间Y-m-d';
COMMENT ON COLUMN "appointment"."appointment_date_counter"."appointment_time" IS '预约时间 Y-m-d H:i:s';
COMMENT ON COLUMN "appointment"."appointment_date_counter"."room_group_id" IS '房型';
COMMENT ON COLUMN "appointment"."appointment_date_counter"."number" IS '临时总数量';
COMMENT ON COLUMN "appointment"."appointment_date_counter"."way" IS '预约方式 1微信 2电话 4店长 8客服';
COMMENT ON COLUMN "appointment"."appointment_date_counter"."appoint_num" IS '已锁定房型数量';

CREATE INDEX "appointment_date_counter_merchant_id_branch_id_appointment__idx" ON "appointment"."appointment_date_counter" USING btree (
  "merchant_id" "pg_catalog"."uuid_ops" ASC NULLS LAST,
  "branch_id" "pg_catalog"."uuid_ops" ASC NULLS LAST,
  "appointment_time" "pg_catalog"."timestamptz_ops" ASC NULLS LAST,
  "room_group_id" "pg_catalog"."uuid_ops" ASC NULLS LAST,
  "way" "pg_catalog"."int2_ops" ASC NULLS LAST
);
CREATE INDEX "appointment_date_counter_merchant_id_branch_id_appointment_idx1" ON "appointment"."appointment_date_counter" USING btree (
  "merchant_id" "pg_catalog"."uuid_ops" ASC NULLS LAST,
  "branch_id" "pg_catalog"."uuid_ops" ASC NULLS LAST,
  "appointment_date" "pg_catalog"."date_ops" ASC NULLS LAST
);

ALTER TABLE "appointment"."appointment_date_counter" ADD CONSTRAINT "appointment_date_counter_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "appointment"."appointment_date_counter";