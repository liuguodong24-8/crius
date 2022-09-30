
-- +migrate Up
CREATE TABLE "centre_data"."branch_b_goods" (
     "branch_goods_id" uuid PRIMARY KEY,
     "branch_id" uuid NOT NULL,
     "goods_id" uuid NOT NULL,
     "price" int4 NOT NULL,
     "make_duration" int4 NOT NULL,
     "up_date" date NOT NULL,
     "down_date" date,
     "create_time" timestamp(6) NOT NULL,
     "update_time" timestamp(6) NOT NULL,
     "delete_time" timestamp(6),
     CONSTRAINT "uk_branch_goods" UNIQUE ("goods_id", "branch_id")
)
;

COMMENT ON COLUMN "centre_data"."branch_b_goods"."branch_goods_id" IS '门店商品主键id';

COMMENT ON COLUMN "centre_data"."branch_b_goods"."branch_id" IS '门店id';

COMMENT ON COLUMN "centre_data"."branch_b_goods"."goods_id" IS '商品id';

COMMENT ON COLUMN "centre_data"."branch_b_goods"."price" IS '门店价格';

COMMENT ON COLUMN "centre_data"."branch_b_goods"."make_duration" IS '制作时长（分钟）';

COMMENT ON COLUMN "centre_data"."branch_b_goods"."up_date" IS '上架时间';

COMMENT ON COLUMN "centre_data"."branch_b_goods"."down_date" IS '下架时间';

COMMENT ON COLUMN "centre_data"."branch_b_goods"."create_time" IS '创建时间';

COMMENT ON COLUMN "centre_data"."branch_b_goods"."update_time" IS '更新时间';

COMMENT ON COLUMN "centre_data"."branch_b_goods"."delete_time" IS '删除时间';

COMMENT ON TABLE "centre_data"."branch_b_goods" IS '中央基础数据-门店商品';

-- +migrate Down
DROP TABLE IF EXISTS "centre_data"."branch_b_goods";