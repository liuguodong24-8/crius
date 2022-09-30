
-- +migrate Up
CREATE TABLE IF NOT EXISTS "member_account"."update_account_bill" (
  "id" uuid PRIMARY KEY NOT NULL,
  "account_bill_id" uuid NOT NULL,
  "old_account_bill_balance" jsonb NOT NULL,
  "new_account_bill_balance" jsonb NOT NULL,
  "branch_id" uuid NOT NULL,
  "staff_id" uuid NOT NULL,
  "reason" varchar(255) NOT NULL,
  "merchant_id" uuid NOT NULL,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL
)
;
COMMENT ON COLUMN "member_account"."update_account_bill"."account_bill_id" IS '账户ID';
COMMENT ON COLUMN "member_account"."update_account_bill"."old_account_bill_balance" IS '操作前账户余额';
COMMENT ON COLUMN "member_account"."update_account_bill"."new_account_bill_balance" IS '操作后账户余额';
COMMENT ON COLUMN "member_account"."update_account_bill"."branch_id" IS '操作门店';
COMMENT ON COLUMN "member_account"."update_account_bill"."staff_id" IS '操作人';
COMMENT ON COLUMN "member_account"."update_account_bill"."merchant_id" IS '商户';
COMMENT ON COLUMN "member_account"."update_account_bill"."reason" IS '原因';
COMMENT ON TABLE "member_account"."update_account_bill" IS '账户修改记录';

CREATE INDEX "idx_update_account_bill_account_id" ON "member_account"."update_account_bill" USING hash ("account_bill_id");

-- +migrate Down
DROP TABLE IF EXISTS "member_account"."update_account_bill";