-- +migrate Up
ALTER TABLE "merchant_basic"."room_type_category" ADD "load_extra" jsonb;
-- +migrate Down
ALTER TABLE "merchant_basic"."room_type_category" DROP "load_extra";