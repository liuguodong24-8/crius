
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."merchant" (
  "id" uuid NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "merchant_basic"."merchant"."name" IS '商户名字';
COMMENT ON TABLE "merchant_basic"."merchant" IS '员工表';

ALTER TABLE "merchant_basic"."merchant" ADD CONSTRAINT "merchant_pkey" PRIMARY KEY ("id");

INSERT INTO "merchant_basic"."merchant" VALUES	( '1d6fac48-77df-4395-8a88-e1ec425baffe', '纯K' );

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."merchant";