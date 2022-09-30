
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."room_type_category" (
  "id" uuid primary key,
  "branch_id" uuid,
  "name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "category" int2 DEFAULT 1,
  "status" varchar COLLATE "pg_catalog"."default",
  "merchant_id" uuid,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;
COMMENT ON COLUMN "merchant_basic"."room_type_category"."name" IS '分类名';
COMMENT ON COLUMN "merchant_basic"."room_type_category"."category" IS '1普通预约 2主题预约';
COMMENT ON COLUMN "merchant_basic"."room_type_category"."status" IS '状态';
COMMENT ON COLUMN "merchant_basic"."room_type_category"."merchant_id" IS '商户id';
COMMENT ON TABLE "merchant_basic"."room_type_category" IS '房型分类';

create index idx_room_type_category_name on "merchant_basic"."room_type_category"(name);
create unique index idx_room_type_category_merchant_name on "merchant_basic"."room_type_category"(merchant_id,name);
create index idx_room_type_category_created_at on "merchant_basic"."room_type_category"(created_at);

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."room_type_category";