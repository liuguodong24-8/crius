
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."mapping_staff_branch" (
  "id" uuid NOT NULL,
  "branch_id" uuid,
  "staff_id" uuid
)
;
COMMENT ON TABLE "merchant_basic"."mapping_staff_branch" IS '员工门店中间表';

ALTER TABLE "merchant_basic"."mapping_staff_branch" ADD CONSTRAINT "mapping_staff_role_pkey" PRIMARY KEY ("id");
ALTER TABLE "merchant_basic"."mapping_staff_branch" ADD CONSTRAINT "fk_branch_id" FOREIGN KEY ("branch_id") REFERENCES "merchant_basic"."branch" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "merchant_basic"."mapping_staff_branch" ADD CONSTRAINT "fk_staff_id" FOREIGN KEY ("staff_id") REFERENCES "merchant_basic"."staff" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."mapping_staff_branch";