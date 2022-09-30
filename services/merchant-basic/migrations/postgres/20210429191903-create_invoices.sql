
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."invoices" (
  "id" uuid NOT NULL primary key,
  "action" varchar(255) NOT NULL,
  "data" jsonb NOT NULL,
  "merchant_id" uuid NOT NULL,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
);

COMMENT ON COLUMN "merchant_basic"."invoices"."action" IS '开票场景';
COMMENT ON COLUMN "merchant_basic"."invoices"."merchant_id" IS '商户';
COMMENT ON COLUMN "merchant_basic"."invoices"."data" IS '票据数据';

create index idx_invoices_created_at on "merchant_basic"."invoices"(created_at);

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."invoices";