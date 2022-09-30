
-- +migrate Up
CREATE TABLE "appointment"."appointment_theme" (
  "id" uuid NOT NULL,
  "color" varchar(20) COLLATE "pg_catalog"."default",
  "feature_ids" uuid[],
  "contents" jsonb,
  "style" varchar(500) COLLATE "pg_catalog"."default",
  "images" varchar[] COLLATE "pg_catalog"."default",
  "video" varchar(500) COLLATE "pg_catalog"."default",
  "weight" int4,
  "status" varchar(20) COLLATE "pg_catalog"."default",
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "name" varchar(50) COLLATE "pg_catalog"."default",
  "category_id" uuid,
  "details" varchar[] COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "appointment"."appointment_theme"."color" IS '颜色';
COMMENT ON COLUMN "appointment"."appointment_theme"."feature_ids" IS '主题特色id';
COMMENT ON COLUMN "appointment"."appointment_theme"."contents" IS '内容[{"name":"xxx", "content":"xxx"}]';
COMMENT ON COLUMN "appointment"."appointment_theme"."style" IS '风格';
COMMENT ON COLUMN "appointment"."appointment_theme"."images" IS '图片';
COMMENT ON COLUMN "appointment"."appointment_theme"."video" IS '视频';
COMMENT ON COLUMN "appointment"."appointment_theme"."weight" IS '权值';
COMMENT ON COLUMN "appointment"."appointment_theme"."status" IS '状态 opened closed';
COMMENT ON COLUMN "appointment"."appointment_theme"."name" IS '名称';
COMMENT ON COLUMN "appointment"."appointment_theme"."category_id" IS '分类id';
COMMENT ON COLUMN "appointment"."appointment_theme"."details" IS '详情';

ALTER TABLE "appointment"."appointment_theme" ADD CONSTRAINT "appointment_theme_pkey1" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "appointment"."appointment_theme";