
-- +migrate Up
CREATE TABLE "centre_data"."product_b_finance_type" (
  "finance_type_id" uuid PRIMARY KEY,
  "erp_code" varchar(32) COLLATE "pg_catalog"."default",
  "grade" int2 NOT NULL,
  "code" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
  "type_name" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "parent_id" uuid,
  "merchant_id" uuid,
  "create_time" timestamp(6) NOT NULL,
  "update_time" timestamp(6),
  "delete_time" timestamp(6)
)
;

COMMENT ON TABLE "centre_data"."product_b_finance_type" IS '中央基础数据-财务类别';

-- +migrate Down
DROP TABLE IF EXISTS "centre_data"."product_b_finance_type";