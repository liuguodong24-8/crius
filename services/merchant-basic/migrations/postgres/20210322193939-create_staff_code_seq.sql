
-- +migrate Up
CREATE SEQUENCE merchant_basic.staff_code_seq;

-- +migrate Down
DROP SEQUENCE merchant_basic.staff_code_seq;