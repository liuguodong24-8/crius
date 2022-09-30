
-- +migrate Up
CREATE TABLE "appointment"."appointment_theme_package" (
  "name" varchar(50) COLLATE "pg_catalog"."default",
  "packages" jsonb,
  "decoration" varchar(500) COLLATE "pg_catalog"."default",
  "staffing" varchar(500) COLLATE "pg_catalog"."default",
  "custom_configs" jsonb,
  "room_types" jsonb,
  "theme_id" uuid,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "id" uuid NOT NULL
)
;
COMMENT ON COLUMN "appointment"."appointment_theme_package"."name" IS '套餐名称';
COMMENT ON COLUMN "appointment"."appointment_theme_package"."packages" IS '[{"id":"xxxxxxx", "category":"xxxx"}] category为product(商品)或者package(套餐) 商品或套餐';
COMMENT ON COLUMN "appointment"."appointment_theme_package"."decoration" IS '装饰布置';
COMMENT ON COLUMN "appointment"."appointment_theme_package"."staffing" IS '人员配置';
COMMENT ON COLUMN "appointment"."appointment_theme_package"."custom_configs" IS '[{"name":"xxxx", "config":"xxxx"}] 自定义配置';
COMMENT ON COLUMN "appointment"."appointment_theme_package"."room_types" IS '[{"id":"xxxx", "price":5000}] 可约房型';
COMMENT ON COLUMN "appointment"."appointment_theme_package"."theme_id" IS '主题id';

ALTER TABLE "appointment"."appointment_theme_package" ADD CONSTRAINT "appointment_theme_package_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "appointment"."appointment_theme_package";