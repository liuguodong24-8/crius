
-- +migrate Up
CREATE TABLE IF NOT EXISTS "appointment"."appointment_template_calendar" (
  "id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "branch_id" uuid NOT NULL,
  "business_date" date NOT NULL,
  "template_id" uuid NOT NULL,
  "category" varchar(10) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;
COMMENT ON COLUMN "appointment"."appointment_template_calendar"."business_date" IS '时间 Y-m-d 营业日';
COMMENT ON COLUMN "appointment"."appointment_template_calendar"."template_id" IS '模版ID';
COMMENT ON COLUMN "appointment"."appointment_template_calendar"."category" IS '应用类型 holiday节假日 normal普通';
COMMENT ON TABLE "appointment"."appointment_template_calendar" IS '预约模版日历';

CREATE INDEX "idx_appointment_template_calendar_branch" ON "appointment"."appointment_template_calendar" USING btree (
  "branch_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_template_calendar_business" ON "appointment"."appointment_template_calendar" USING btree (
  "business_date" "pg_catalog"."date_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_template_calendar_merchant" ON "appointment"."appointment_template_calendar" USING btree (
  "merchant_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);

ALTER TABLE "appointment"."appointment_template_calendar" ADD CONSTRAINT "appointment_template_calendar_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "appointment"."appointment_template_calendar";