
-- +migrate Up
CREATE TABLE "merchant_basic"."growth_config" (
  "id" uuid NOT NULL,
  "merchant_id" uuid,
  "name" varchar(50) COLLATE "pg_catalog"."default",
  "top" int4,
  "rules" varchar[] COLLATE "pg_catalog"."default",
  CONSTRAINT "growth_config_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "unique_growth_config_merid" UNIQUE ("merchant_id")
)
;

ALTER TABLE "merchant_basic"."growth_config"
  OWNER TO "micro";

COMMENT ON COLUMN "merchant_basic"."growth_config"."name" IS '成长值名称';

COMMENT ON COLUMN "merchant_basic"."growth_config"."top" IS '成长值上限';

COMMENT ON COLUMN "merchant_basic"."growth_config"."rules" IS '成长值规则';

COMMENT ON TABLE "merchant_basic"."growth_config" IS '成长值上限表';
-- +migrate Down

DROP TABLE IF EXISTS "merchant_basic"."growth_config";