-- +migrate Up
ALTER TABLE "merchant_basic"."room_type" ADD "load_extra" jsonb;
-- +migrate Down
ALTER TABLE "merchant_basic"."room_type" DROP "load_extra";