
-- +migrate Up
CREATE TABLE IF NOT EXISTS "member_account"."consume_bills" (
  "id" uuid NOT NULL,
  "base_value" int4 DEFAULT 0,
  "gift_value" int4 DEFAULT 0,
  "products" jsonb,
  "packages" jsonb,
  "tickets" jsonb,
  "bill_id" uuid,
  "consume_bill_id" uuid,
  "load_extra" jsonb,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL
)
;
COMMENT ON COLUMN "member_account"."consume_bills"."base_value" IS '本金';
COMMENT ON COLUMN "member_account"."consume_bills"."gift_value" IS '赠金';
COMMENT ON COLUMN "member_account"."consume_bills"."products" IS '商品 [{"id":"uuid","title":"string","count":"int"}]';
COMMENT ON COLUMN "member_account"."consume_bills"."packages" IS '活动 [{"id":"uuid","title":"string","count":"int"}]';
COMMENT ON COLUMN "member_account"."consume_bills"."tickets" IS '优惠券 [{"id":"uuid","title":"string","count":"int", "type" : "string"}]';
COMMENT ON COLUMN "member_account"."consume_bills"."bill_id" IS '消费账单id';
COMMENT ON COLUMN "member_account"."consume_bills"."consume_bill_id" IS '消费所扣充值账单id';
COMMENT ON TABLE "member_account"."consume_bills" IS '账户流水';

ALTER TABLE "member_account"."consume_bills" ADD CONSTRAINT "consume_bills_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "member_account"."consume_bills";