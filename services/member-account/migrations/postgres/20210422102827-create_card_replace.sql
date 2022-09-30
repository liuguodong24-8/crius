
-- +migrate Up
CREATE TABLE IF NOT EXISTS "member_account"."card_replace" (
  "id" uuid PRIMARY KEY NOT NULL,
  "curr_card_id" uuid NOT NULL,
  "new_card_id" uuid NOT NULL,
  "staff_id" uuid NOT NULL,
  "payments" jsonb NOT NULL,
  "merchant_id" uuid NOT NULL,
  "extra" jsonb NULL,
  "load_extra" jsonb,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "member_account"."card_replace"."curr_card_id" IS '旧卡ID';
COMMENT ON COLUMN "member_account"."card_replace"."curr_card_id" IS '新卡ID';
COMMENT ON COLUMN "member_account"."card_replace"."staff_id" IS '操作人';
COMMENT ON COLUMN "member_account"."card_replace"."payments" IS '补卡手续费支付方式';
COMMENT ON TABLE "member_account"."card_replace" IS '补卡';

CREATE INDEX "idx_card_replace_merchant" ON "member_account"."card_replace" USING hash ("merchant_id");

-- +migrate Down
DROP TABLE IF EXISTS "member_account"."card_replace";