
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."district" (
  "id" uuid primary key,
  "code" serial,
  "name" varchar(40) COLLATE "pg_catalog"."default" NOT NULL,
  "merchant_id" uuid,
  "status" varchar(255) COLLATE "pg_catalog"."default",
  "load_extra" jsonb,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
)
;
COMMENT ON COLUMN "merchant_basic"."district"."id" IS '主键';
COMMENT ON COLUMN "merchant_basic"."district"."code" IS '地区编号';
COMMENT ON COLUMN "merchant_basic"."district"."name" IS '地区名';
COMMENT ON TABLE "merchant_basic"."district" IS '地区管理';

ALTER TABLE "merchant_basic"."district" ADD CONSTRAINT "uk_district_name" UNIQUE ("name");
ALTER TABLE "merchant_basic"."district" ADD CONSTRAINT "uk_district_code" UNIQUE ("code");
create index idx_district_created_at on "merchant_basic"."district"(created_at);

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."district";