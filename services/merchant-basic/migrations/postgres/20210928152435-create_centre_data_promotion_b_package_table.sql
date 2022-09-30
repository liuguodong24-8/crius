
-- +migrate Up
CREATE TABLE "centre_data"."promotion_b_package" (
  "package_id" uuid PRIMARY KEY,
  "package_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "code" varchar(64) COLLATE "pg_catalog"."default",
  "image_url" varchar(128) COLLATE "pg_catalog"."default",
  "pos_code" varchar(64) COLLATE "pg_catalog"."default",
  "simplify" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "active_type_id" uuid NOT NULL,
  "begin_date" date NOT NULL,
  "end_date" date,
  "goods_set" jsonb,
  "branch_ids" uuid[],
  "merchant_id" uuid,
  "create_time" timestamptz(6) NOT NULL,
  "update_time" timestamptz(6),
  "delete_time" timestamptz(6)
)
;

COMMENT ON TABLE "centre_data"."promotion_b_package" IS '中央基础数据-套餐';

-- +migrate Down
DROP TABLE IF EXISTS "centre_data"."promotion_b_package";