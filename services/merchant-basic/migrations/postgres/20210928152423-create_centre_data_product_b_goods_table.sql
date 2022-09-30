
-- +migrate Up
CREATE TABLE "centre_data"."product_b_goods" (
  "goods_id" uuid PRIMARY KEY,
  "pos_code" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "code" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "cn_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "simplify" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "finance_type_id" uuid NOT NULL,
  "operate_type_id" uuid NOT NULL,
  "erp_code" varchar(32) COLLATE "pg_catalog"."default",
  "alias" varchar(64) COLLATE "pg_catalog"."default",
  "en_name" varchar(32) COLLATE "pg_catalog"."default",
  "crafts" int2 NOT NULL,
  "guide_price" int4,
  "bar_code" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "is_discountable" int2 NOT NULL DEFAULT 1,
  "sale_unit_id" uuid NOT NULL,
  "check_unit_id" uuid,
  "unit_relation" int4 NOT NULL,
  "address" varchar(255) COLLATE "pg_catalog"."default",
  "vendor" varchar(64) COLLATE "pg_catalog"."default",
  "content" varchar(127) COLLATE "pg_catalog"."default",
  "up_date" date NOT NULL,
  "down_date" date,
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "branch_ids" uuid[],
  "image_url" varchar(127) COLLATE "pg_catalog"."default",
  "make_duration" int4,
  "workshop_id" uuid,
  "image_urls" varchar[],
  "quick" int4[],
  "merchant_id" uuid,
  "create_time" timestamp(6) NOT NULL,
  "update_time" timestamp(6) NOT NULL,
  "delete_time" timestamp(6)
)
;

COMMENT ON TABLE "centre_data"."product_b_goods" IS '中央基础数据-商品';

-- +migrate Down
DROP TABLE IF EXISTS "centre_data"."product_b_goods";