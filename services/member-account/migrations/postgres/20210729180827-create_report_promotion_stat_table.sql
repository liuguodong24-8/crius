
-- +migrate Up
CREATE TABLE "member_account"."report_promotion_stat" (
  "id" uuid primary key,
  "bill_id" uuid NOT NULL,
  "branch_id" uuid NOT NULL,
  "card_id" uuid NOT NULL,
  "card_code" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "promotion_option_id" uuid,
  "promotion_option_name" varchar(50) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "recharge_value" int8 default 0,
  "total" int default 1,
  "change_type" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "change_value" int4 DEFAULT 0,
  "created_at" timestamptz(6)
)
;

create index idx_member_account_report_promotion_stat_create on "member_account"."report_promotion_stat" (created_at);
create index idx_member_account_report_promotion_stat_branch on "member_account"."report_promotion_stat" (branch_id);
create index idx_member_account_report_promotion_stat_type on "member_account"."report_promotion_stat" (change_type);

COMMENT ON COLUMN "member_account"."report_promotion_stat"."bill_id" IS '账单ID account_bill.id';

COMMENT ON COLUMN "member_account"."report_promotion_stat"."promotion_option_id" IS '优惠方案ID';

COMMENT ON COLUMN "member_account"."report_promotion_stat"."promotion_option_name" IS '优惠方案名字';

COMMENT ON COLUMN "member_account"."report_promotion_stat"."change_type" IS '类型 参照account_bill.change_type';

COMMENT ON COLUMN "member_account"."report_promotion_stat"."change_value" IS '金额';

COMMENT ON TABLE "member_account"."report_promotion_stat" IS '账户流水，按优惠方案汇总';

-- +migrate Down
DROP TABLE IF EXISTS "member_account"."report_promotion_stat";