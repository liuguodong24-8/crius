
-- +migrate Up
CREATE TABLE if not exists "centre_data"."product_b_operate_type" (
    "operate_type_id" uuid PRIMARY KEY,
    "erp_code" varchar(32) COLLATE "pg_catalog"."default",
    "grade" int2 NOT NULL,
    "code" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
    "type_name" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
    "parent_id" uuid,
    "merchant_id" uuid,
    "create_time" timestamp(6) NOT NULL,
    "update_time" timestamp(6),
    "delete_time" timestamp(6),
    CONSTRAINT "uk_operate_code" UNIQUE ("code", "grade", "parent_id"),
    CONSTRAINT "uk_operate_type_name" UNIQUE ("type_name")
)
;

COMMENT ON COLUMN "centre_data"."product_b_operate_type"."operate_type_id" IS '运营类别id';

COMMENT ON COLUMN "centre_data"."product_b_operate_type"."erp_code" IS 'ERP编码';

COMMENT ON COLUMN "centre_data"."product_b_operate_type"."grade" IS '类别等级, 0:大类, 1:中类, 2:小类';

COMMENT ON COLUMN "centre_data"."product_b_operate_type"."code" IS '类别编号, sprintf("%03d", 1)';

COMMENT ON COLUMN "centre_data"."product_b_operate_type"."type_name" IS '类别名字';

COMMENT ON COLUMN "centre_data"."product_b_operate_type"."parent_id" IS '关联上级id';

COMMENT ON TABLE "centre_data"."product_b_operate_type" IS '中央基础数据-运营类别';

-- +migrate Down
DROP TABLE IF EXISTS "centre_data"."product_b_operate_type";