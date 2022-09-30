
-- +migrate Up
CREATE TABLE "appointment"."appointment_theme_category" (
  "id" uuid NOT NULL,
  "name" varchar(50) COLLATE "pg_catalog"."default",
  "weight" int4,
  "status" varchar(20) COLLATE "pg_catalog"."default",
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;
COMMENT ON COLUMN "appointment"."appointment_theme_category"."name" IS '主题名称';
COMMENT ON COLUMN "appointment"."appointment_theme_category"."weight" IS '权值';
COMMENT ON COLUMN "appointment"."appointment_theme_category"."status" IS '状态 opened closed';

ALTER TABLE "appointment"."appointment_theme_category" ADD CONSTRAINT "appointment_theme_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "appointment"."appointment_theme_category";