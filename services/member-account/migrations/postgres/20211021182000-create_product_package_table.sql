
-- +migrate Up
CREATE TABLE "member_account"."bill_product_package" (
  "product_package_id" uuid NOT NULL,
  "code" varchar(255) COLLATE "pg_catalog"."default",
  "number" int4,
  "price" int4,
  "title" varchar(50) COLLATE "pg_catalog"."default",
  "bill_id" uuid,
  "left" int4,
  "category" varchar(10) COLLATE "pg_catalog"."default",
  "load_extra" jsonb,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;
COMMENT ON COLUMN "member_account"."bill_product_package"."product_package_id" IS '商品/套餐id';
COMMENT ON COLUMN "member_account"."bill_product_package"."code" IS 'pos码';
COMMENT ON COLUMN "member_account"."bill_product_package"."number" IS '数量';
COMMENT ON COLUMN "member_account"."bill_product_package"."price" IS '价格';
COMMENT ON COLUMN "member_account"."bill_product_package"."title" IS '名称';
COMMENT ON COLUMN "member_account"."bill_product_package"."bill_id" IS '账单id';
COMMENT ON COLUMN "member_account"."bill_product_package"."left" IS '剩余数量';
COMMENT ON COLUMN "member_account"."bill_product_package"."category" IS '类型 商品 product, 套餐 package';

-- +migrate Down
DROP TABLE IF EXISTS "member_account"."bill_product_package";