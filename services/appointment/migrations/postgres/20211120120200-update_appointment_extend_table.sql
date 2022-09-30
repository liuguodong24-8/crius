
-- +migrate Up
ALTER TABLE "appointment"."appointment_extend" ADD COLUMN "package_price" int4;
COMMENT ON COLUMN "appointment"."appointment_extend"."package_price" IS '套餐价格';

-- +migrate Down
ALTER TABLE "appointment"."appointment_extend" DROP COLUMN "package_price";