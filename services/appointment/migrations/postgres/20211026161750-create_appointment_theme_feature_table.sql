
-- +migrate Up
CREATE TABLE "appointment"."appointment_theme_feature" (
  "id" uuid NOT NULL,
  "name" varchar(50) COLLATE "pg_catalog"."default",
  "icon" varchar(500) COLLATE "pg_catalog"."default",
  "weight" int4,
  "status" varchar(20) COLLATE "pg_catalog"."default",
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;
COMMENT ON COLUMN "appointment"."appointment_theme_feature"."name" IS '颜色';
COMMENT ON COLUMN "appointment"."appointment_theme_feature"."icon" IS '图标';
COMMENT ON COLUMN "appointment"."appointment_theme_feature"."weight" IS '权值';
COMMENT ON COLUMN "appointment"."appointment_theme_feature"."status" IS '状态 opened closed';

ALTER TABLE "appointment"."appointment_theme_feature" ADD CONSTRAINT "appointment_theme_feature_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "appointment"."appointment_theme_feature";