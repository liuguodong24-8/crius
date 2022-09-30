
-- +migrate Up
ALTER TABLE "member_account"."account_bills" ADD COLUMN "pos_bill_id" uuid;
ALTER TABLE "member_account"."account_bills" ADD COLUMN "status" varchar(20);
COMMENT ON COLUMN "member_account"."account_bills"."pos_bill_id" IS 'pos账单id';
COMMENT ON COLUMN "member_account"."account_bills"."status" IS '账单状态 成功success, 取消cancel';

-- +migrate Down
ALTER TABLE "member_account"."account_bills" DROP COLUMN "pos_bill_id";
ALTER TABLE "member_account"."account_bills" DROP COLUMN "status";