
-- +migrate Up
CREATE TABLE IF NOT EXISTS "member_account"."add_account_deduction" (
  "id" uuid PRIMARY KEY NOT NULL,
  "bill_number" varchar(255) NOT NULL,
  "consume_ids" uuid[] NOT NULL,
  "branch_id" uuid NOT NULL,
  "staff_id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "reason" varchar(255) NOT NULL,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL
)
;
COMMENT ON COLUMN "member_account"."add_account_deduction"."bill_number" IS '关联账单号';
COMMENT ON COLUMN "member_account"."add_account_deduction"."consume_ids" IS '消费流水ID';
COMMENT ON COLUMN "member_account"."add_account_deduction"."branch_id" IS '操作门店';
COMMENT ON COLUMN "member_account"."add_account_deduction"."staff_id" IS '操作人';
COMMENT ON COLUMN "member_account"."add_account_deduction"."merchant_id" IS '商户';
COMMENT ON COLUMN "member_account"."add_account_deduction"."reason" IS '原因';
COMMENT ON TABLE "member_account"."add_account_deduction" IS '账户手动扣款记录';

-- +migrate Down
DROP TABLE IF EXISTS "member_account"."add_account_deduction";