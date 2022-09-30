-- +migrate Up
CREATE TABLE IF NOT EXISTS "message".sms_templates (
    id uuid primary key,
    merchant_id uuid not null,
    name varchar(20) not null,
    category varchar(50) not null,
    category_key varchar(50) not null,
    sign varchar(20) not null,
    content varchar(1024) not null,
    status varchar(6) default 'opened',
    extra jsonb,
    created_at timestamptz(6) not null,
    updated_at timestamptz(6) not null
);
COMMENT ON TABLE "message".sms_templates IS '短信模版';
COMMENT ON COLUMN "message".sms_templates.merchant_id IS '商户';
COMMENT ON COLUMN "message".sms_templates.name IS '模版名';
COMMENT ON COLUMN "message".sms_templates.category IS '类型';
COMMENT ON COLUMN "message".sms_templates.sign IS '签章';
COMMENT ON COLUMN "message".sms_templates.content IS '内容';
-- +migrate Down
DROP TABLE IF EXISTS "message"."sms_templates";