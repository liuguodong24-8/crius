
-- +migrate Up
ALTER TABLE "appointment"."appointment_extend" ADD COLUMN "load_extra" jsonb;
COMMENT ON COLUMN "appointment"."appointment_extend"."load_extra" IS '导入信息';

-- +migrate Down
ALTER TABLE "appointment"."appointment_extend" DROP COLUMN "load_extra";