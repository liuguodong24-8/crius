
-- +migrate Up
ALTER TABLE "appointment"."appointment_extend" ADD COLUMN "code" varchar(50);
COMMENT ON COLUMN "appointment"."appointment_extend"."code" IS '预约号';

-- +migrate Down
ALTER TABLE "appointment"."appointment_extend" DROP COLUMN "code";