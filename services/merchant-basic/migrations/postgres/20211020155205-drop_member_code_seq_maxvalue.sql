
-- +migrate Up
alter SEQUENCE "merchant_basic"."member_code_seq" NO MAXVALUE;
-- +migrate Down
alter SEQUENCE "merchant_basic"."member_code_seq" MAXVALUE  999999;