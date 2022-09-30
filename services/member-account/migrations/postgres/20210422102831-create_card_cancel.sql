
-- +migrate Up
CREATE TABLE IF NOT EXISTS "member_account"."card_cancel" (
  "id" uuid PRIMARY KEY NOT NULL,
  "card_id" uuid NOT NULL,
  "account_id" uuid[] NOT NULL,
  "bank_account" varchar(255) NULL,
  "bank_name" varchar(255) NULL,
  "money_receiver" varchar(255) NULL,
  "reason" text NOT NULL,
  "apply_staff_id" uuid NOT NULL,
  "apply_at" timestamptz(6) NOT NULL,
  "status" varchar(255) NOT NULL,
  "examine_staff_id" uuid NULL,
  "examine_at" timestamptz(6),
  "refund_value" int DEFAULT 0,
  "reject_reason" varchar(255) DEFAULT NULL,
  "merchant_id" uuid NOT NULL,
  "load_extra" jsonb,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL
)
;
COMMENT ON COLUMN "member_account"."card_cancel"."card_id" IS '卡ID';
COMMENT ON COLUMN "member_account"."card_cancel"."account_id" IS '账户ID';
COMMENT ON COLUMN "member_account"."card_cancel"."bank_account" IS '退款银行账户';
COMMENT ON COLUMN "member_account"."card_cancel"."bank_name" IS '开户行';
COMMENT ON COLUMN "member_account"."card_cancel"."money_receiver" IS '收款人';
COMMENT ON COLUMN "member_account"."card_cancel"."reason" IS '注销原因';
COMMENT ON COLUMN "member_account"."card_cancel"."apply_staff_id" IS '申请员工ID';
COMMENT ON COLUMN "member_account"."card_cancel"."apply_at" IS '申请时间';
COMMENT ON COLUMN "member_account"."card_cancel"."status" IS '申请状态';
COMMENT ON COLUMN "member_account"."card_cancel"."examine_staff_id" IS '审核人';
COMMENT ON COLUMN "member_account"."card_cancel"."examine_at" IS '审核时间';
COMMENT ON COLUMN "member_account"."card_cancel"."refund_value" IS '退款金额';
COMMENT ON COLUMN "member_account"."card_cancel"."reject_reason" IS '驳回原因';
COMMENT ON TABLE "member_account"."card_cancel" IS '补卡';

CREATE INDEX "idx_card_cancel_merchant" ON "member_account"."card_cancel" USING hash ("merchant_id");
CREATE INDEX "idx_card_cancel_account" ON "member_account"."card_cancel" USING hash ("account_id");
CREATE INDEX "idx_card_cancel_card" ON "member_account"."card_cancel" USING hash ("card_id");

-- +migrate Down
DROP TABLE IF EXISTS "member_account"."card_cancel";