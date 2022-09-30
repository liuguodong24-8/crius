
-- +migrate Up
CREATE TABLE "member_private"."promotion_options" (
  "id" uuid NOT NULL,
  "name" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "merchant_id" uuid NOT NULL,
  "promotion_id" uuid NOT NULL,
  "recharge_value" int8 DEFAULT 0,
  "base_value" int8 DEFAULT 0,
  "gift_value" int8 DEFAULT 0,
  "describe" text COLLATE "pg_catalog"."default",
  "products" jsonb,
  "packages" jsonb,
  "tickets" jsonb,
  "status" varchar(10) COLLATE "pg_catalog"."default" DEFAULT 'opened'::character varying,
  tag_id uuid,
  "extra" jsonb,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  "deleted_at" timestamptz(6),
  CONSTRAINT "promotion_options_pkey" PRIMARY KEY ("id")
);

create index idx_member_private_promotion_option_created_at on "member_private"."promotion_options"(created_at);
create index idx_member_private_promotion_option_name on "member_private"."promotion_options"(name);
create index idx_member_private_promotion_option_pmerchant on "member_private"."promotion_options"(merchant_id);
create index idx_member_private_promotion_option_promotion on "member_private"."promotion_options"(promotion_id);

COMMENT ON COLUMN "member_private"."promotion_options"."id" IS '优惠方案组';
COMMENT ON COLUMN "member_private"."promotion_options"."name" IS '方案组名';
COMMENT ON COLUMN "member_private"."promotion_options"."merchant_id" IS '商户ID';
COMMENT ON COLUMN "member_private"."promotion_options"."promotion_id" IS '优惠方案组ID';
COMMENT ON COLUMN "member_private"."promotion_options"."recharge_value" IS '充值金额';
COMMENT ON COLUMN "member_private"."promotion_options"."base_value" IS '本金';
COMMENT ON COLUMN "member_private"."promotion_options"."gift_value" IS '赠金';
COMMENT ON COLUMN "member_private"."promotion_options"."describe" IS '说明';
COMMENT ON COLUMN "member_private"."promotion_options"."products" IS '赠品';
COMMENT ON COLUMN "member_private"."promotion_options"."packages" IS '套餐活动';
COMMENT ON COLUMN "member_private"."promotion_options"."tickets" IS '券';
COMMENT ON COLUMN "member_private"."promotion_options"."tag_id" IS '对应标签';
COMMENT ON COLUMN "member_private"."promotion_options"."status" IS '状态 opened启用 closed禁用';
-- +migrate Down
DROP TABLE IF EXISTS "member_private"."promotion_options";