
-- +migrate Up
CREATE TABLE "centre_data"."product_b_unit" (
  "unit_id" uuid NOT NULL,
  "unit_name" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "create_time" timestamp(6) NOT NULL,
  "update_time" timestamp(6) NOT NULL,
  "delete_time" timestamp(6),
  CONSTRAINT "pk_unit" PRIMARY KEY ("unit_id"),
  CONSTRAINT "uk_unit_name" UNIQUE ("unit_name")
)
;

COMMENT ON COLUMN "centre_data"."product_b_unit"."unit_id" IS '主键';

COMMENT ON COLUMN "centre_data"."product_b_unit"."unit_name" IS '单位';

COMMENT ON COLUMN "centre_data"."product_b_unit"."create_time" IS '创建时间';

COMMENT ON COLUMN "centre_data"."product_b_unit"."update_time" IS '更新时间';

COMMENT ON COLUMN "centre_data"."product_b_unit"."delete_time" IS '删除时间';

COMMENT ON TABLE "centre_data"."product_b_unit" IS '单位';

-- +migrate Down
DROP TABLE IF EXISTS "centre_data"."product_b_unit";