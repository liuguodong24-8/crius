
-- +migrate Up
CREATE TABLE "centre_data"."promotion_b_active_type" (
  "active_type_id" uuid PRIMARY KEY,
  "grade" int2 NOT NULL,
  "code" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
  "type_name" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "parent_id" uuid,
  "merchant_id" uuid,
  "create_time" timestamp(6) NOT NULL,
  "update_time" timestamp(6) NOT NULL,
  "delete_time" timestamp(6)
)
;

COMMENT ON TABLE "centre_data"."promotion_b_active_type" IS '中央基础数据-活动类型';

-- +migrate Down
DROP TABLE IF EXISTS "centre_data"."promotion_b_active_type";