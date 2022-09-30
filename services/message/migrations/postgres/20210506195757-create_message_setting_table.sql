-- +migrate Up
CREATE TABLE IF NOT EXISTS "message".message_setting (
    id uuid primary key,
    status varchar(10) not null,
    merchant_id uuid not null,
    message_type varchar(50) not null,
    trigger_type varchar(10) not null,
    advance_hour float default 0,
    sms_template_id uuid,
    wechat_template_id uuid,
    special_setting jsonb null,
    cc_list jsonb null,
    special_branches uuid [] null,
    extra jsonb,
    created_at timestamptz(6) not null,
    updated_at timestamptz(6) not null
);
COMMENT ON TABLE "message".message_setting IS '发送设置';
COMMENT ON COLUMN "message".message_setting.merchant_id IS '商户';
COMMENT ON COLUMN "message".message_setting.message_type IS '消息类型';
COMMENT ON COLUMN "message".message_setting.trigger_type IS '触发类型';
COMMENT ON COLUMN "message".message_setting.advance_hour IS '提前小时';
COMMENT ON COLUMN "message".message_setting.special_setting IS '特殊内容 [{"begin":"2021-05-01","end":"2021-06-01","sms_template_id":"xxx", "wechat_template_id":"xxxx"}]';
COMMENT ON COLUMN "message".message_setting.cc_list IS '抄送人 [{"code":"86","phone":"13800138000"}]';
COMMENT ON COLUMN "message".message_setting.special_branches IS '特殊门店 {"xxxx"}';
-- +migrate Down
DROP TABLE IF EXISTS "message"."message_setting";