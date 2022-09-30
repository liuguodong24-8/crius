
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."mapping_role_permission" (
  "id" uuid NOT NULL,
  "role_id" uuid,
  "permission_id" int4
)
;
COMMENT ON COLUMN "merchant_basic"."mapping_role_permission"."role_id" IS '角色ID';
COMMENT ON COLUMN "merchant_basic"."mapping_role_permission"."permission_id" IS '权限ID';
COMMENT ON TABLE "merchant_basic"."mapping_role_permission" IS '角色权限中间表';

ALTER TABLE "merchant_basic"."mapping_role_permission" ADD CONSTRAINT "mapping_role_permission_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."mapping_role_permission";
