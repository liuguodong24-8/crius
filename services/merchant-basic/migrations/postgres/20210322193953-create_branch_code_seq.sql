
-- +migrate Up
CREATE SEQUENCE merchant_basic.branch_code_seq;

-- +migrate Down
DROP SEQUENCE merchant_basic.branch_code_seq;