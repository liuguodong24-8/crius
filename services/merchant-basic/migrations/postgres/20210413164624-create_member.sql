
-- +migrate Up
CREATE SEQUENCE "merchant_basic"."member_code_seq" maxvalue 999999 start with 100001;

CREATE TABLE IF NOT EXISTS "merchant_basic"."member" (
  "id" uuid NOT NULL,
  "name" varchar(20) NOT NULL default '',
  "phone" varchar(20) NOT NULL,
  "phone_tail" char(1) NOT NULL,
  "phone_suffix" char(4) NOT NULL,
  "phone_code" varchar(10),
  "avatar" varchar(255) NULL,
  "gender" int2 NOT NULL DEFAULT 0,
  "birthday" date NULL,
  "code" int8 NOT NULL DEFAULT nextval('"merchant_basic".member_code_seq'::regclass),
  "first_branch_id" uuid,
  "first_brand" uuid,
  "staff_id" uuid,
  "first_channel" varchar(30) NOT NULL,
  "channels" VARCHAR[] NOT NULL,
  "city" varchar(30),
  "merchant_id" uuid NOT NULL,
  "load_extra" jsonb,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
) PARTITION BY LIST ("phone_tail")
;
COMMENT ON COLUMN "merchant_basic"."member"."id" IS '会员ID';
COMMENT ON COLUMN "merchant_basic"."member"."name" IS '名字';
COMMENT ON COLUMN "merchant_basic"."member"."phone" IS '电话';
COMMENT ON COLUMN "merchant_basic"."member"."phone_tail" IS '电话尾号';
COMMENT ON COLUMN "merchant_basic"."member"."phone_suffix" IS '电话后4位';
COMMENT ON COLUMN "merchant_basic"."member"."phone_code" IS '电话区号';
COMMENT ON COLUMN "merchant_basic"."member"."avatar" IS '头像';
COMMENT ON COLUMN "merchant_basic"."member"."gender" IS '1男 2女';
COMMENT ON COLUMN "merchant_basic"."member"."birthday" IS '生日';
COMMENT ON COLUMN "merchant_basic"."member"."code" IS '自增唯一编码';
COMMENT ON COLUMN "merchant_basic"."member"."first_branch_id" IS '首次交互门店';
COMMENT ON COLUMN "merchant_basic"."member"."first_brand" IS '首次交互门店所属品牌';
COMMENT ON COLUMN "merchant_basic"."member"."staff_id" IS '专属客服';
COMMENT ON COLUMN "merchant_basic"."member"."city" IS '首次开户门店所在城市code';
COMMENT ON COLUMN "merchant_basic"."member"."first_channel" IS '会员渠道来源 开卡open_card 微信wechat 预约 reservation';
COMMENT ON COLUMN "merchant_basic"."member"."channels" IS '会员渠道来源 开卡open_card 微信wechat 预约 reservation';
COMMENT ON COLUMN "merchant_basic"."member"."merchant_id" IS '商户';
COMMENT ON TABLE "merchant_basic"."member" IS '会员表';

ALTER TABLE "merchant_basic"."member" ADD CONSTRAINT "member_phone_key" UNIQUE ("phone", "phone_tail");

-- +migrate StatementBegin
create or replace function create_merchant_basic_member_sub() returns void as $BODY$
declare i integer;
declare table_name varchar;
begin
  i = 0;
  for i in 0..9 LOOP
    table_name = 'member_' || i;
    EXECUTE FORMAT('CREATE TABLE IF NOT EXISTS %I PARTITION OF "merchant_basic"."member" FOR VALUES IN (%s)', table_name , i);

    EXECUTE format('CREATE INDEX idx_member_%s_phone ON %I (phone, phone_code);', i, table_name);
    EXECUTE format('CREATE INDEX idx_member_%s_name ON %I (name);', i, table_name);
    EXECUTE format('CREATE INDEX idx_member_%s_created_at ON %I (created_at);', i, table_name);
    EXECUTE format('CREATE INDEX idx_member_%s_phone_suffix ON %I (phone_suffix);', i, table_name);
    end LOOP;
    return;
end;
$BODY$ LANGUAGE plpgsql;

-- +migrate StatementEnd

select create_merchant_basic_member_sub();
-- +migrate Down
DROP FUNCTION IF EXISTS create_merchant_basic_member_sub();
DROP TABLE IF EXISTS "merchant_basic"."member";
DROP SEQUENCE IF EXISTS "merchant_basic"."member_code_seq";

