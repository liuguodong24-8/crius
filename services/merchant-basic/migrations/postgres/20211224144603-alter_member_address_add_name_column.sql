
-- +migrate Up
ALTER TABLE "merchant_basic"."member_address" ADD "name" varchar(30) DEFAULT '';
-- +migrate Down
ALTER TABLE "merchant_basic"."member_address" DROP "name";