-- +migrate Up
CREATE TABLE IF NOT EXISTS "message".wechat_templates (
    id uuid primary key,
    merchant_id uuid not null,
    template_name varchar(20) not null,
    template_code varchar(64) not null,
    template_content jsonb not null,
    official_link varchar(255),
    miniprogram_link varchar(255),
    category varchar(50) not null,
    category_key varchar(50) not null,
    status varchar(6) default 'opened',
    extra jsonb,
    created_at timestamptz(6) not null,
    updated_at timestamptz(6) not null
);
COMMENT ON TABLE "message".wechat_templates IS '微信模版';
COMMENT ON COLUMN "message".wechat_templates.merchant_id IS '商户';
COMMENT ON COLUMN "message".wechat_templates.template_name IS '模版名';
COMMENT ON COLUMN "message".wechat_templates.template_code IS '微信模版ID';
COMMENT ON COLUMN "message".wechat_templates.template_content IS '模版内容{
                  "first": {
                    "value" : "xxx",
                    "color" : "#abcdef"
                  },
                  "detail" : [
                   {
                      "name" : "thing2",
                       "value" : "xxx",
                       "color" : "#abcdef"
                   }
                  ],
                  "remark" : {
                    "value" : "备注",
                    "color" : "#abcdef"
                  }
              }';
COMMENT ON COLUMN "message".wechat_templates.official_link IS '公众号链接';
COMMENT ON COLUMN "message".wechat_templates.miniprogram_link IS '小程序链接';
COMMENT ON COLUMN "message".wechat_templates.category IS '类型';

create index idx_wechat_template_name on message.wechat_templates (template_name);
create index idx_wechat_merchant on message.wechat_templates (merchant_id);
-- +migrate Down
DROP TABLE IF EXISTS "message"."wechat_templates";