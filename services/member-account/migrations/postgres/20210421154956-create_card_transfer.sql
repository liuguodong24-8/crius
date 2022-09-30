
-- +migrate Up

CREATE TABLE IF NOT EXISTS "member_account"."card_transfer" (
  "id" uuid PRIMARY KEY NOT NULL,
  "source_account_id" uuid NOT NULL,
  "dest_account_id" uuid NOT NULL,
  "transfer_value" int4 DEFAULT 0,
  "staff_id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "extra" jsonb NULL,
  "load_extra" jsonb,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL
)
;
COMMENT ON COLUMN "member_account"."card_transfer"."source_account_id" IS '源卡ID';
COMMENT ON COLUMN "member_account"."card_transfer"."dest_account_id" IS '目标卡ID';
COMMENT ON COLUMN "member_account"."card_transfer"."transfer_value" IS '划帐金额';
COMMENT ON COLUMN "member_account"."card_transfer"."staff_id" IS '操作人';
COMMENT ON TABLE "member_account"."card_transfer" IS '划账记录';

CREATE INDEX "idx_card_transfer_merchant" ON "member_account"."card_transfer" USING hash ("merchant_id");

-- +migrate Down
DROP TABLE IF EXISTS "member_account"."card_transfer";
