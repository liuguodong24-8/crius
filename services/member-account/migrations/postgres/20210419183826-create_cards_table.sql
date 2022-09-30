
-- +migrate Up
CREATE TABLE IF NOT EXISTS "member_account"."cards" (
  "id" uuid NOT NULL,
  "category" varchar(10) COLLATE "pg_catalog"."default" NOT NULL,
  "code" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "create_branch_id" uuid NOT NULL,
  "create_staff_id" uuid,
  "status" varchar(10) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'init'::character varying,
  "password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "open_staff_id" uuid,
  "account_ids" uuid[],
  "open_branch_id" uuid,
  "sub_category" varchar(20) COLLATE "pg_catalog"."default",
  "primary_id" uuid,
  "merchant_id" uuid,
  "member_id" uuid,
  "extra" jsonb,
  "load_extra" jsonb,
  "opened_at" timestamptz(6),
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "member_account"."cards"."category" IS '类型 member会员卡 gift礼品卡';
COMMENT ON COLUMN "member_account"."cards"."code" IS '卡号';
COMMENT ON COLUMN "member_account"."cards"."create_branch_id" IS '创建门店id';
COMMENT ON COLUMN "member_account"."cards"."create_staff_id" IS '创建人';
COMMENT ON COLUMN "member_account"."cards"."status" IS '状态：init默认，active已激活, lost挂失, cancelling注销审核中, cancelled注销';
COMMENT ON COLUMN "member_account"."cards"."opened_at" IS '开通时间';
COMMENT ON COLUMN "member_account"."cards"."password" IS '密码';
COMMENT ON COLUMN "member_account"."cards"."open_staff_id" IS '开卡经办人';
COMMENT ON COLUMN "member_account"."cards"."account_ids" IS '关联账户ids';
COMMENT ON COLUMN "member_account"."cards"."open_branch_id" IS '开卡门店id';
COMMENT ON COLUMN "member_account"."cards"."sub_category" IS '账户类别: primary主卡账户, secondary副卡账户, blank不记名账户';
COMMENT ON COLUMN "member_account"."cards"."primary_id" IS '此卡为副卡时，主卡id';
COMMENT ON COLUMN "member_account"."cards"."merchant_id" IS '商户id';
COMMENT ON COLUMN "member_account"."cards"."member_id" IS '用户id';
COMMENT ON TABLE "member_account"."cards" IS '卡';

CREATE INDEX "cards_code_idx" ON "member_account"."cards" USING btree (
  "code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

ALTER TABLE "member_account"."cards" ADD CONSTRAINT "cards_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "member_account"."cards";