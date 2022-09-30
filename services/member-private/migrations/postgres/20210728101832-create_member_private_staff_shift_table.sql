
-- +migrate Up
CREATE TABLE "member_private"."staff_shifts" (
  "id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "branch_id" uuid NOT NULL,
  "staff_id" uuid NOT NULL,
  "begin_time" timestamptz(6) NOT NULL,
  "end_time" timestamptz(6) NOT NULL,
  "extra" jsonb,
  "created_at" timestamptz(6) NOT NULL,
  CONSTRAINT "staff_shifts_pkey" PRIMARY KEY ("id")
);

create index idx_member_private_staff_shift_created_at on "member_private"."staff_shifts"(created_at);
create index idx_member_private_staff_shift_merchant on "member_private"."staff_shifts"(merchant_id);
create index idx_member_private_staff_shift_branch on "member_private"."staff_shifts"(branch_id);
create index idx_member_private_staff_shift_staff on "member_private"."staff_shifts"(staff_id);

COMMENT ON COLUMN "member_private"."staff_shifts"."id" IS '员工交班信息';
COMMENT ON COLUMN "member_private"."staff_shifts"."merchant_id" IS '商户ID';
COMMENT ON COLUMN "member_private"."staff_shifts"."branch_id" IS '门店ID';
COMMENT ON COLUMN "member_private"."staff_shifts"."staff_id" IS '员工';
COMMENT ON COLUMN "member_private"."staff_shifts"."begin_time" IS '班次开始时间';
COMMENT ON COLUMN "member_private"."staff_shifts"."end_time" IS '班次结束时间';
-- +migrate Down
DROP TABLE IF EXISTS "member_private"."staff_shifts";