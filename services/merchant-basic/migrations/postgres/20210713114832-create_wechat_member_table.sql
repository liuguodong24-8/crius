
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."wechat_member" (
  "id" uuid NOT NULL PRIMARY KEY,
  "openid" varchar(50) NOT NULL,
  "appid" varchar(50) NOT NULL,
  "member_id" uuid,
  "nickname" varchar(100),
  "sex" int2,
  "province" varchar(50),
  "city" varchar(50),
  "headimgurl" varchar(255),
  "unionid" varchar(50),
  "extra" json,
  "load_extra" jsonb,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;

COMMENT ON COLUMN "merchant_basic"."wechat_member"."member_id" IS 'member表id';

create index idx_wechat_member_created_at on "merchant_basic"."wechat_member"(created_at);
create index idx_wechat_member_member on "merchant_basic"."wechat_member"(member_id);
create UNIQUE index idx_wechat_member_openid on "merchant_basic"."wechat_member"(appid,openid);
-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."wechat_member";
