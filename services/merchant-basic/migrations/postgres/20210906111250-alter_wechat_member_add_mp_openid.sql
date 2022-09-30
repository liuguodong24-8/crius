
-- +migrate Up
ALTER TABLE "merchant_basic"."wechat_member" ADD COLUMN "mp_openid" varchar(50);
COMMENT ON COLUMN "merchant_basic"."wechat_member"."mp_openid" IS '小程序openid';
ALTER TABLE "merchant_basic"."wechat_member" ADD COLUMN "mp_appid" varchar(50);
COMMENT ON COLUMN "merchant_basic"."wechat_member"."mp_appid" IS '小程序appid';
create UNIQUE index idx_wechat_member_unionid on "merchant_basic"."wechat_member"(unionid);
-- +migrate Down
alter table "merchant_basic"."wechat_member" drop column if exists mp_openid;
alter table "merchant_basic"."wechat_member" drop column if exists mp_appid;
DROP index "merchant_basic"."idx_wechat_member_unionid";