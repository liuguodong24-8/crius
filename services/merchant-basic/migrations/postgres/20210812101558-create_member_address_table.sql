
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."member_address" (
  "id" uuid NOT NULL primary key,
  "member_id" uuid NOT NULL,
  "phone" varchar(20) NOT NULL,
  "phone_code" varchar(10),
  "province_id" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "city_id" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "district_id" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "address" varchar(1024) COLLATE "pg_catalog"."default",
  "is_default" bool default false,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
);

COMMENT ON TABLE "merchant_basic"."member_address" IS '用户地址信息';

create index idx_member_address_member on "merchant_basic"."member_address"(member_id);
create index idx_member_address_created_at on "merchant_basic"."member_address"(created_at);

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."member_address";