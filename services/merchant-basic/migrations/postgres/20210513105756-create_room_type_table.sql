
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."room_type" (
  "id" uuid primary key,
  "branch_id" uuid,
  "category_id" uuid NOT NULL,
  "name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "category" int2 DEFAULT 1,
  "status" varchar COLLATE "pg_catalog"."default",
  "merchant_id" uuid,
  "customer_min" int2,
  "customer_max" int2,
  "order" int8,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;
COMMENT ON COLUMN "merchant_basic"."room_type"."category_id" IS '分类ID';
COMMENT ON COLUMN "merchant_basic"."room_type"."name" IS '房型名';
COMMENT ON COLUMN "merchant_basic"."room_type"."category" IS '1普通预约 2主题预约';
COMMENT ON COLUMN "merchant_basic"."room_type"."status" IS '状态';
COMMENT ON COLUMN "merchant_basic"."room_type"."merchant_id" IS '商户id';
COMMENT ON COLUMN "merchant_basic"."room_type"."customer_min" IS '房型最少人数';
COMMENT ON COLUMN "merchant_basic"."room_type"."customer_max" IS '房型最多人数';
COMMENT ON COLUMN "merchant_basic"."room_type"."order" IS '房型排序号';
COMMENT ON TABLE "merchant_basic"."room_type" IS '房型分类';

create index idx_room_type_name on "merchant_basic"."room_type"(name);
create unique index idx_room_type_merchant_name on "merchant_basic"."room_type"(merchant_id,name);
create index idx_room_type_created_at on "merchant_basic"."room_type"(created_at);

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."room_type";