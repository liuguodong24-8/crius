
-- +migrate Up
CREATE TABLE "merchant_basic"."growth_rule" (
  "id" uuid NOT NULL,
  "name" varchar(50) COLLATE "pg_catalog"."default",
  "growth_gain" jsonb,
  "expire_day" int4,
  "branches" uuid[],
  "status" varchar(20) COLLATE "pg_catalog"."default",
  "merchant_id" uuid,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  CONSTRAINT "growth_rule_pkey" PRIMARY KEY ("id")
)
;

ALTER TABLE "merchant_basic"."growth_rule"
  OWNER TO "micro";

COMMENT ON COLUMN "merchant_basic"."growth_rule"."name" IS '规则名字';

COMMENT ON COLUMN "merchant_basic"."growth_rule"."growth_gain" IS '[{“id”,"uuid", "consume":10 //消费10元=1成长值}]';

COMMENT ON COLUMN "merchant_basic"."growth_rule"."expire_day" IS '有效期(天) 0永久有效';

COMMENT ON COLUMN "merchant_basic"."growth_rule"."branches" IS '可用门店列表';

COMMENT ON COLUMN "merchant_basic"."growth_rule"."status" IS '状态 opened, closed';

COMMENT ON TABLE "merchant_basic"."growth_rule" IS '成长值规则表';
-- +migrate Down

DROP TABLE IF EXISTS "merchant_basic"."growth_rule";
