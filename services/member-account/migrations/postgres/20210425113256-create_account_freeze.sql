
-- +migrate Up
CREATE TABLE IF NOT EXISTS "member_account"."account_freeze" (
  "id" uuid PRIMARY KEY NOT NULL,
  "account_id" uuid NOT NULL,
  "action" varchar(20) NOT NULL,
  "reason" varchar(255) NOT NULL,
  "staff_id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL
)
;

COMMENT ON COLUMN "member_account"."account_freeze"."account_id" IS '账户ID';
COMMENT ON COLUMN "member_account"."account_freeze"."action" IS '冻结:freeze;解冻unfreeze';
COMMENT ON COLUMN "member_account"."account_freeze"."reason" IS '原因';
COMMENT ON COLUMN "member_account"."account_freeze"."staff_id" IS '操作人';
COMMENT ON COLUMN "member_account"."account_freeze"."merchant_id" IS '商户';
COMMENT ON TABLE "member_account"."account_freeze" IS '账户冻结/解冻记录';

CREATE INDEX "idx_account_freeze_account_id" ON "member_account"."account_freeze" USING hash ("account_id");

-- +migrate Down
DROP TABLE IF EXISTS "member_account"."account_freeze";