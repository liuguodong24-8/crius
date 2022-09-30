
-- +migrate Up
CREATE TABLE "merchant_basic"."consume_category" (
  "id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "category" varchar(50) NOT NULL,
  "code" varchar(20) NOT NULL,
  "status" varchar(10) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'opened'::character varying,
  "extra" jsonb,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  "deleted_at" timestamptz(6),
  CONSTRAINT "merchant_basic_point_category_pkey" PRIMARY KEY ("id")
)
;

create index idx_merchant_basic_point_category_created_at on "merchant_basic"."consume_category"(created_at);
create index idx_merchant_basic_point_category_merchant on "merchant_basic"."consume_category"(merchant_id);
create index idx_merchant_basic_point_category_category on "merchant_basic"."consume_category"(category);
create index idx_merchant_basic_point_category_code on "merchant_basic"."consume_category"(code);

COMMENT ON TABLE "merchant_basic"."consume_category" IS '积分消费类型';
COMMENT ON COLUMN "merchant_basic"."consume_category"."id" IS '消费类型';
COMMENT ON COLUMN "merchant_basic"."consume_category"."merchant_id" IS '商户ID';
COMMENT ON COLUMN "merchant_basic"."consume_category"."category" IS '类型名称';
COMMENT ON COLUMN "merchant_basic"."consume_category"."code" IS '类型编码';
COMMENT ON COLUMN "merchant_basic"."consume_category"."status" IS '状态 opened启用 closed禁用';
-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."consume_category";