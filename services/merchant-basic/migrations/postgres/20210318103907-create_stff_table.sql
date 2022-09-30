
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."staff" (
  "id" uuid primary key,
  "merchant_id" uuid,
  "name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "phone" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "phone_code" varchar(10) COLLATE "pg_catalog"."default",
  "gender" int2 NOT NULL DEFAULT 0,
  "status" varchar(10) COLLATE "pg_catalog"."default",
  "code" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "employee_code" varchar(20) COLLATE "pg_catalog"."default",
  "quit_at" timestamptz(6),
  "entry_at" timestamptz(6),
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "salt" varchar(20) COLLATE "pg_catalog"."default",
  "admin" bool DEFAULT false,
  "load_extra" jsonb,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "merchant_basic"."staff"."id" IS '员工ID';
COMMENT ON COLUMN "merchant_basic"."staff"."name" IS '名字';
COMMENT ON COLUMN "merchant_basic"."staff"."phone" IS '电话';
COMMENT ON COLUMN "merchant_basic"."staff"."phone_code" IS '区号';
COMMENT ON COLUMN "merchant_basic"."staff"."gender" IS '1男 2女';
COMMENT ON COLUMN "merchant_basic"."staff"."status" IS '启用状态opened closed';
COMMENT ON COLUMN "merchant_basic"."staff"."code" IS '编号(规则自动生成)';
COMMENT ON COLUMN "merchant_basic"."staff"."quit_at" IS '离职时间';
COMMENT ON COLUMN "merchant_basic"."staff"."entry_at" IS '入职时间';
COMMENT ON COLUMN "merchant_basic"."staff"."password" IS '密码';
COMMENT ON COLUMN "merchant_basic"."staff"."salt" IS '密码加密盐';
COMMENT ON COLUMN "merchant_basic"."staff"."employee_code" IS '工号(用户手动输入)';
COMMENT ON COLUMN "merchant_basic"."staff"."deleted_at" IS '删除时间';
COMMENT ON COLUMN "merchant_basic"."staff"."admin" IS '超级管理员';
COMMENT ON COLUMN "merchant_basic"."staff"."merchant_id" IS '商户ID';
COMMENT ON TABLE "merchant_basic"."staff" IS '员工表';

create index idx_staff_phone on "merchant_basic"."staff"(phone);
create index idx_staff_created_at on "merchant_basic"."staff"(created_at);
create unique index idx_staff_employee_code on "merchant_basic"."staff"(employee_code);
create unique index idx_staff_code on "merchant_basic"."staff"(code);

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."staff";
