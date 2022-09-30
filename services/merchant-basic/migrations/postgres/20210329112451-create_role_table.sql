
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."role" (
  "id" uuid primary key,
  "name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "status" varchar(10) COLLATE "pg_catalog"."default" NOT NULL,
  "property" int2 NOT NULL,
  "staff_id" uuid,
  "merchant_id" uuid,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "merchant_basic"."role"."name" IS '角色名称';
COMMENT ON COLUMN "merchant_basic"."role"."status" IS '状态opened closed';
COMMENT ON COLUMN "merchant_basic"."role"."property" IS '1公有 2私有';
COMMENT ON COLUMN "merchant_basic"."role"."staff_id" IS '创造者id';
COMMENT ON COLUMN "merchant_basic"."role"."merchant_id" IS '商户ID';
COMMENT ON TABLE "merchant_basic"."role" IS '角色表';

create index idx_role_created_at on "merchant_basic"."role"(created_at);
create index idx_role_staff on "merchant_basic"."role"("staff_id");
create index idx_role_name on "merchant_basic"."role"("name");

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."role";