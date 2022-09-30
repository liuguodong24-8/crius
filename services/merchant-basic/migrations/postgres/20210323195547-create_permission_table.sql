
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."permission" (
  "id" int4 NOT NULL,
  "permission" varchar(500) COLLATE "pg_catalog"."default",
  "service" varchar(50) COLLATE "pg_catalog"."default",
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;
COMMENT ON TABLE "merchant_basic"."permission" IS '权限表';

ALTER TABLE "merchant_basic"."permission" ADD CONSTRAINT "permission_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."permission";