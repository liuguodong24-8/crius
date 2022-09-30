
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."mapping_staff_role" (
  "id" uuid NOT NULL,
  "staff_id" uuid,
  "role_id" uuid
)
;
COMMENT ON COLUMN "merchant_basic"."mapping_staff_role"."staff_id" IS '员工ID';
COMMENT ON COLUMN "merchant_basic"."mapping_staff_role"."role_id" IS '角色ID';
COMMENT ON TABLE "merchant_basic"."mapping_staff_role" IS '员工角色中间表';

ALTER TABLE "merchant_basic"."mapping_staff_role" ADD CONSTRAINT "mapping_staff_role_pkey1" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."mapping_staff_role";