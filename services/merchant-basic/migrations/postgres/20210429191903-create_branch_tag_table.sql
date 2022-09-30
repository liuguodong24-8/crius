
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."branch_tag" (
  "id" uuid primary key,
  "name" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "branches" uuid[],
  "create_staff_id" uuid,
  "status" varchar(32) COLLATE "pg_catalog"."default",
  "merchant_id" uuid,
  "extra" jsonb,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "merchant_basic"."branch_tag"."name" IS '标签名';
COMMENT ON COLUMN "merchant_basic"."branch_tag"."branches" IS '门店数组id';
COMMENT ON COLUMN "merchant_basic"."branch_tag"."create_staff_id" IS '创建人';
COMMENT ON COLUMN "merchant_basic"."branch_tag"."merchant_id" IS '商户id';
COMMENT ON TABLE "merchant_basic"."branch_tag" IS '门店标签';

create index idx_branch_tag_created_at on "merchant_basic"."branch_tag"(created_at);

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."branch_tag";