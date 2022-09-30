
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."brand" (
  "id" uuid NOT NULL primary key,
  "name" varchar(50) COLLATE "pg_catalog"."default",
  "status" varchar(20) COLLATE "pg_catalog"."default",
  "order" int2,
  "merchant_id" uuid,
  "load_extra" jsonb,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;
COMMENT ON COLUMN "merchant_basic"."brand"."name" IS '品牌名称';
COMMENT ON COLUMN "merchant_basic"."brand"."status" IS '状态 opened closed';
COMMENT ON COLUMN "merchant_basic"."brand"."order" IS '排序值';

create index idx_brand_created_at on "merchant_basic"."brand"(created_at);

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."brand";