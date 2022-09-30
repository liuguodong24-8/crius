
-- +migrate Up
alter table "merchant_basic"."wechat_member" drop column if exists mp_openid;
alter table "merchant_basic"."wechat_member" drop column if exists mp_appid;

DROP index if exists "merchant_basic"."idx_wechat_member_unionid";
DROP index if exists "merchant_basic"."idx_wechat_member_openid";
DROP index if exists "merchant_basic"."idx_wechat_member_member";

create UNIQUE index idx_wechat_member_unionid on "merchant_basic"."wechat_member"(appid,unionid);
create index idx_wechat_member_member on "merchant_basic"."wechat_member"(member_id);

-- +migrate Down
ALTER TABLE "merchant_basic"."wechat_member" ADD COLUMN "mp_openid" varchar(50);
COMMENT ON COLUMN "merchant_basic"."wechat_member"."mp_openid" IS '小程序openid';
ALTER TABLE "merchant_basic"."wechat_member" ADD COLUMN "mp_appid" varchar(50);
COMMENT ON COLUMN "merchant_basic"."wechat_member"."mp_appid" IS '小程序appid';

DROP index if exists "merchant_basic"."idx_wechat_member_unionid";
DROP index if exists "merchant_basic"."idx_wechat_member_member";

create UNIQUE index idx_wechat_member_unionid on "merchant_basic"."wechat_member"(unionid);
create UNIQUE index idx_wechat_member_member on "merchant_basic"."wechat_member"(member_id);
