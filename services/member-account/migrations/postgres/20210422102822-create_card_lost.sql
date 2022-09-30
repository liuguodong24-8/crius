
-- +migrate Up
CREATE TABLE IF NOT EXISTS "member_account"."card_lost" (
  "id" uuid PRIMARY KEY NOT NULL,
  "card_id" uuid NOT NULL,
  "staff_id" uuid NOT NULL,
  "action" varchar(255) NULL,
  "merchant_id" uuid NOT NULL,
  "extra" jsonb NULL,
  "load_extra" jsonb,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "member_account"."card_lost"."card_id" IS '卡ID';
COMMENT ON COLUMN "member_account"."card_lost"."staff_id" IS '操作人';
COMMENT ON COLUMN "member_account"."card_lost"."action" IS 'lost挂失 find找回';
COMMENT ON TABLE "member_account"."card_lost" IS '卡挂失找回';

CREATE INDEX "idx_card_lost_merchant" ON "member_account"."card_lost" USING hash ("merchant_id");


-- +migrate Down
DROP TABLE IF EXISTS "member_account"."card_lost";

