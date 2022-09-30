
-- +migrate Up
ALTER TABLE "merchant_basic"."room_type" ADD COLUMN "room_type_group_ids" uuid[];
ALTER TABLE "merchant_basic"."consume_category" ADD COLUMN "operator_types" uuid[];
ALTER TABLE "merchant_basic"."consume_category" ADD COLUMN "active_types" uuid[];
COMMENT ON COLUMN "merchant_basic"."room_type"."room_type_group_ids" IS '房型分组id';
COMMENT ON COLUMN "merchant_basic"."consume_category"."operator_types" IS '运营类型';
COMMENT ON COLUMN "merchant_basic"."consume_category"."active_types" IS '活动类型';

-- +migrate Down
ALTER TABLE "merchant_basic"."room_type" DROP COLUMN "room_type_group_ids";
ALTER TABLE "merchant_basic"."consume_category" DROP COLUMN "operator_types";
ALTER TABLE "merchant_basic"."consume_category" DROP COLUMN "active_types";