
-- +migrate Up
CREATE TABLE "merchant_basic"."point_rule" (
  "id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "rule_name" varchar(20) NOT NULL,
  "gain_rules" jsonb,
  "use_rules" jsonb,
  "validity_day" int default 1,
  "branch_ids" UUID[],
  "status" varchar(10) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'opened'::character varying,
  "extra" jsonb,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  "deleted_at" timestamptz(6),
  CONSTRAINT "merchant_basic_point_rule_pkey" PRIMARY KEY ("id")
)
;

create index idx_merchant_basic_point_rule_created_at on "merchant_basic"."point_rule"(created_at);
create index idx_merchant_basic_point_rule_merchant on "merchant_basic"."point_rule"(merchant_id);
create index idx_merchant_basic_point_rule_name on "merchant_basic"."point_rule"(rule_name);


COMMENT ON TABLE "merchant_basic"."point_rule" IS '积分规则';
COMMENT ON COLUMN "merchant_basic"."point_rule"."merchant_id" IS '商户ID';
COMMENT ON COLUMN "merchant_basic"."point_rule"."rule_name" IS '规则名';
COMMENT ON COLUMN "merchant_basic"."point_rule"."gain_rules" IS '获取规则[{"category_id":"uuid","point":10,"fee":1}]';
COMMENT ON COLUMN "merchant_basic"."point_rule"."use_rules" IS '抵扣规则[{"category_id":"uuid","point":10,"fee":1}]';
COMMENT ON COLUMN "merchant_basic"."point_rule"."status" IS '状态 opened启用 closed禁用';
COMMENT ON COLUMN "merchant_basic"."point_rule"."validity_day" IS '有效天数 -1永久有效';
-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."point_rule";