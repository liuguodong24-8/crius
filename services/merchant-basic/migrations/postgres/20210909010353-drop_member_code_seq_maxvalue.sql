
-- +migrate Up
ALTER TABLE "merchant_basic"."member" ALTER COLUMN code SET default null;
 DROP SEQUENCE  IF EXISTS  "merchant_basic"."member_code_seq";
 CREATE SEQUENCE "merchant_basic"."member_code_seq" NO MAXVALUE start with 100001;
 ALTER TABLE "merchant_basic"."member" ALTER COLUMN code SET DEFAULT nextval('"merchant_basic".member_code_seq'::regclass);
-- +migrate Down
