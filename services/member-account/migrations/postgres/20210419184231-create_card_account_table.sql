
-- +migrate Up
CREATE TABLE "member_account"."card_account" (
  "id" uuid NOT NULL,
  "member_id" uuid,
  "base_value" int4 DEFAULT 0,
  "gift_value" int4 DEFAULT 0,
  "products" jsonb,
  "packages" jsonb,
  "category" varchar(20) COLLATE "pg_catalog"."default",
  "status" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "branch_id" uuid,
  "merchant_id" uuid,
  "tag_id" uuid,
  "extra" jsonb,
  "load_extra" jsonb,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "member_account"."card_account"."member_id" IS '用户';
COMMENT ON COLUMN "member_account"."card_account"."base_value" IS '本金';
COMMENT ON COLUMN "member_account"."card_account"."gift_value" IS '赠金';
COMMENT ON COLUMN "member_account"."card_account"."products" IS '商品 [{"id":"uuid","title":"string","number":"int"}]';
COMMENT ON COLUMN "member_account"."card_account"."packages" IS '活动 [{"id":"uuid","title":"string","number":"int"}]';
COMMENT ON COLUMN "member_account"."card_account"."category" IS '账户类别: primary主卡账户, secondary副卡账户, blank不记名账户';
COMMENT ON COLUMN "member_account"."card_account"."status" IS '状态:activated,frozen,cancelled';
COMMENT ON COLUMN "member_account"."card_account"."branch_id" IS '账户开通门店';
COMMENT ON COLUMN "member_account"."card_account"."merchant_id" IS '商户id';
COMMENT ON COLUMN "member_account"."card_account"."tag_id" IS '标签';
COMMENT ON TABLE "member_account"."card_account" IS '补卡';

ALTER TABLE "member_account"."card_account" ADD CONSTRAINT "member_account_card_account_value_check" CHECK (base_value >= 0 AND gift_value >= 0);

ALTER TABLE "member_account"."card_account" ADD CONSTRAINT "card_account_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "member_account"."card_account";