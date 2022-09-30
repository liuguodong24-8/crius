
-- +migrate Up
ALTER TABLE "appointment"."appointment_extend" ADD COLUMN "open_at" timestamptz;
COMMENT ON COLUMN "appointment"."appointment_extend"."open_at" IS '开房时间';
ALTER TABLE "appointment"."appointment_extend" ADD COLUMN "open_room_id" uuid;
COMMENT ON COLUMN "appointment"."appointment_extend"."open_at" IS '开房房间id';
ALTER TABLE "appointment"."appointment_extend" ADD COLUMN "sended" bool;
COMMENT ON COLUMN "appointment"."appointment_extend"."sended" IS '是否通知咨客';
ALTER TABLE "appointment"."appointment_extend" DROP COLUMN "created_at";
ALTER TABLE "appointment"."appointment_extend" DROP COLUMN "updated_at";

-- +migrate Down
ALTER TABLE "appointment"."appointment_extend" DROP COLUMN "open_at";
ALTER TABLE "appointment"."appointment_extend" DROP COLUMN "open_room_id";
ALTER TABLE "appointment"."appointment_extend" DROP COLUMN "sended";
ALTER TABLE "appointment"."appointment_extend" ADD COLUMN "created_at" timestamptz;
ALTER TABLE "appointment"."appointment_extend" ADD COLUMN "updated_at" timestamptz;