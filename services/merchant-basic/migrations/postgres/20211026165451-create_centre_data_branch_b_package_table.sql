
-- +migrate Up
CREATE TABLE "centre_data"."branch_b_package" (
      "branch_package_id" uuid PRIMARY KEY,
      "branch_id" uuid NOT NULL,
      "package_id" uuid NOT NULL,
      "price" int4 NOT NULL,
      "begin_date" date NOT NULL,
      "end_date" date,
      "create_time" timestamp(6) NOT NULL,
      "update_time" timestamp(6) NOT NULL,
      "delete_time" timestamp(6),
      CONSTRAINT "uk_branch_package" UNIQUE ("package_id", "branch_id")
)
;

COMMENT ON COLUMN "centre_data"."branch_b_package"."branch_package_id" IS '门店套餐活动id';

COMMENT ON COLUMN "centre_data"."branch_b_package"."branch_id" IS '门店id';

COMMENT ON COLUMN "centre_data"."branch_b_package"."package_id" IS '套餐活动id';

COMMENT ON COLUMN "centre_data"."branch_b_package"."price" IS '套餐活动价格';

COMMENT ON COLUMN "centre_data"."branch_b_package"."begin_date" IS '有效开始日期';

COMMENT ON COLUMN "centre_data"."branch_b_package"."end_date" IS '有效结束日期';

COMMENT ON COLUMN "centre_data"."branch_b_package"."create_time" IS '创建时间';

COMMENT ON COLUMN "centre_data"."branch_b_package"."update_time" IS '更新时间';

COMMENT ON COLUMN "centre_data"."branch_b_package"."delete_time" IS '删除时间';

COMMENT ON TABLE "centre_data"."branch_b_package" IS '中央基础数据-门店套餐活动';

-- +migrate Down
DROP TABLE IF EXISTS "centre_data"."branch_b_package";