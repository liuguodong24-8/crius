
-- +migrate Up
ALTER TABLE "merchant_basic"."consume_category" ADD COLUMN "is_room_fee" bool;
COMMENT ON COLUMN "merchant_basic"."consume_category"."is_room_fee" IS '是否是房费';

-- +migrate Down
ALTER TABLE "merchant_basic"."consume_category" DROP COLUMN "is_room_fee";