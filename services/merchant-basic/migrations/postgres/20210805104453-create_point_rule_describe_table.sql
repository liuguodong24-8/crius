
-- +migrate Up
CREATE TABLE "merchant_basic"."point_rule_describe" (
  "id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "graphic_detail" varchar[],
  "extra" jsonb,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  CONSTRAINT "merchant_basic_point_rule_describe_pkey" PRIMARY KEY ("id")
)
;

create index idx_point_rule_describe_created_at on "merchant_basic"."point_rule_describe"(created_at);

COMMENT ON TABLE "merchant_basic"."point_rule_describe" IS '积分规则图文详情';
COMMENT ON COLUMN "merchant_basic"."point_rule_describe"."merchant_id" IS '商户ID';
-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."point_rule_describe";