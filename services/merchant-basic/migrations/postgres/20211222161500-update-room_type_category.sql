
-- +migrate Up
ALTER TABLE "merchant_basic"."room_type_category" DROP COLUMN "branch_id";

-- +migrate Down
ALTER TABLE "merchant_basic"."room_type_category" ADD COLUMN "branch_id" uuid;