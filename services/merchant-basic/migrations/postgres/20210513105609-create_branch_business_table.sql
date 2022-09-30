
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."branch_business" (
  "id" uuid NOT NULL,
  "branch_id" uuid,
  "begin_date" date,
  "end_date" date,
  "weeks" int4[],
  "begin_time" time(6),
  "end_time" time(6),
  "is_next_day" bool,
  "merchant_id" uuid,
  "status" varchar(20),
  "category" varchar(20),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;
COMMENT ON COLUMN "merchant_basic"."branch_business"."begin_date" IS '开始时间';
COMMENT ON COLUMN "merchant_basic"."branch_business"."end_date" IS '结束时间';
COMMENT ON COLUMN "merchant_basic"."branch_business"."weeks" IS '星期';
COMMENT ON COLUMN "merchant_basic"."branch_business"."begin_time" IS '开始 12:00';
COMMENT ON COLUMN "merchant_basic"."branch_business"."end_time" IS '结束 12:00';
COMMENT ON COLUMN "merchant_basic"."branch_business"."is_next_day" IS '结束时间明日';
COMMENT ON COLUMN "merchant_basic"."branch_business"."merchant_id" IS '商户id';
COMMENT ON COLUMN "merchant_basic"."branch_business"."status" IS '状态 opened closed';
COMMENT ON COLUMN "merchant_basic"."branch_business"."category" IS '类型 normal special';
COMMENT ON TABLE "merchant_basic"."branch_business" IS '门店营业日';

CREATE INDEX "branch_business_branch_id_merchant_id_idx" ON "merchant_basic"."branch_business" USING btree (
  "branch_id" "pg_catalog"."uuid_ops" ASC NULLS LAST,
  "merchant_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);

ALTER TABLE "merchant_basic"."branch_business" ADD CONSTRAINT "branch_business_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."branch_business";