
-- +migrate Up
CREATE TABLE IF NOT EXISTS "appointment"."appointment_template" (
  "id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "branch_id" uuid NOT NULL,
  "name" varchar COLLATE "pg_catalog"."default",
  "color" varchar COLLATE "pg_catalog"."default",
  "status" varchar COLLATE "pg_catalog"."default",
  "room_type_ids" uuid[],
  "begin_time" time(6),
  "end_time" time(6),
  "is_next_day" int2,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;
COMMENT ON COLUMN "appointment"."appointment_template"."branch_id" IS '门店ID';
COMMENT ON COLUMN "appointment"."appointment_template"."name" IS '模板名称';
COMMENT ON COLUMN "appointment"."appointment_template"."color" IS '模板颜色';
COMMENT ON COLUMN "appointment"."appointment_template"."status" IS '状态 opened, closed';
COMMENT ON COLUMN "appointment"."appointment_template"."room_type_ids" IS '房型id';
COMMENT ON COLUMN "appointment"."appointment_template"."begin_time" IS '开始时间"12:00"';
COMMENT ON COLUMN "appointment"."appointment_template"."end_time" IS '结束时间"06:00"';
COMMENT ON COLUMN "appointment"."appointment_template"."is_next_day" IS '是否下一天';
COMMENT ON TABLE "appointment"."appointment_template" IS '预约模版';

CREATE INDEX "appointment_template_branch_id_idx" ON "appointment"."appointment_template" USING btree (
  "branch_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "appointment_template_name_idx" ON "appointment"."appointment_template" USING btree (
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

ALTER TABLE "appointment"."appointment_template" ADD CONSTRAINT "appointment_template_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "appointment"."appointment_template";