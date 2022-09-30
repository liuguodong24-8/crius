
-- +migrate Up
CREATE TABLE "member_account"."account_bills" (
  "id" uuid NOT NULL,
  "bill_code" varchar(20) COLLATE "pg_catalog"."default",
  "account_id" uuid NOT NULL,
  "card_id" uuid,
  "card_code" varchar(20) COLLATE "pg_catalog"."default",
  "branch_id" uuid,
  "change_value" int4 DEFAULT 0,
  "change_type" varchar(20) COLLATE "pg_catalog"."default",
  "base_value" int4 DEFAULT 0,
  "gift_value" int4 DEFAULT 0,
  "payments" jsonb,
  "staff_id" uuid,
  "operator_comment" varchar(20) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "merchant_id" uuid,
  "base_value_left" int4,
  "gift_value_left" int4,
  "promotion_options" jsonb,
  "extra" jsonb,
  "load_extra" jsonb,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  "change_category" varchar(20) COLLATE "pg_catalog"."default",
  "after_account" jsonb
)
;
COMMENT ON COLUMN "member_account"."account_bills"."bill_code" IS '编号';
COMMENT ON COLUMN "member_account"."account_bills"."account_id" IS '账户';
COMMENT ON COLUMN "member_account"."account_bills"."card_id" IS '卡ID';
COMMENT ON COLUMN "member_account"."account_bills"."card_code" IS '卡号';
COMMENT ON COLUMN "member_account"."account_bills"."branch_id" IS '门店ID';
COMMENT ON COLUMN "member_account"."account_bills"."change_value" IS '改变金额';
COMMENT ON COLUMN "member_account"."account_bills"."change_type" IS '操作类型详情 open开卡 sub开副卡 nobody不记名卡 recharge充值 sub-consume划副卡消费 change修改余额 deduction添加扣款 replace补卡 consume消费';
COMMENT ON COLUMN "member_account"."account_bills"."base_value" IS '本金';
COMMENT ON COLUMN "member_account"."account_bills"."gift_value" IS '赠金';
COMMENT ON COLUMN "member_account"."account_bills"."payments" IS '支付方式 {"wechat":0,"cash":0,"alipay":0,"card":0}';
COMMENT ON COLUMN "member_account"."account_bills"."staff_id" IS '操作人';
COMMENT ON COLUMN "member_account"."account_bills"."operator_comment" IS '操作备注 开卡经办人 充值推荐人';
COMMENT ON COLUMN "member_account"."account_bills"."merchant_id" IS '商户id';
COMMENT ON COLUMN "member_account"."account_bills"."base_value_left" IS '本金剩余';
COMMENT ON COLUMN "member_account"."account_bills"."gift_value_left" IS '赠金剩余';
COMMENT ON COLUMN "member_account"."account_bills"."promotion_options" IS '优惠方案详情 [{"id":"uuid","count":1}]';
COMMENT ON COLUMN "member_account"."account_bills"."change_category" IS '操作类型 recharge充值 consume消费';
COMMENT ON TABLE "member_account"."account_bills" IS '账户流水';

ALTER TABLE "member_account"."account_bills" ADD CONSTRAINT "account_bills_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "member_account"."account_bills";