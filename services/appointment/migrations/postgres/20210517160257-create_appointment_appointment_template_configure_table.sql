
-- +migrate Up
CREATE TABLE IF NOT EXISTS "appointment"."appointment_template_configure" (
  "id" uuid NOT NULL,
  "room_type_id" uuid,
  "template_id" uuid,
  "advance_day" int4,
  "deposit_fee" int8,
  "configure" jsonb
)
;
COMMENT ON COLUMN "appointment"."appointment_template_configure"."room_type_id" IS '房型id';
COMMENT ON COLUMN "appointment"."appointment_template_configure"."template_id" IS '模板id';
COMMENT ON COLUMN "appointment"."appointment_template_configure"."advance_day" IS '提前预定时间(天)';
COMMENT ON COLUMN "appointment"."appointment_template_configure"."deposit_fee" IS '订金';
COMMENT ON COLUMN "appointment"."appointment_template_configure"."configure" IS '[{{"way":"或存值，与查询","time":"12:00:00","num":1,"is_next_day":false}]';
COMMENT ON TABLE "appointment"."appointment_template_configure" IS '预约模版配置';

CREATE INDEX "appointment_template_configure_room_type_id_template_id_idx" ON "appointment"."appointment_template_configure" USING btree (
  "room_type_id" "pg_catalog"."uuid_ops" ASC NULLS LAST,
  "template_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "appointment_template_configure_template_id_idx" ON "appointment"."appointment_template_configure" USING btree (
  "template_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);

ALTER TABLE "appointment"."appointment_template_configure" ADD CONSTRAINT "appointment_template_configure_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "appointment"."appointment_template_configure";