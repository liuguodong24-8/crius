
-- +migrate Up
INSERT INTO "merchant_basic"."branch_tag"("id", "name", "status", "create_staff_id", "merchant_id", "created_at", "updated_at") VALUES ('5260754d-e100-4fec-8f66-ff9d2a5e3545', '直营', 'opened', 'f24c596a-da55-43ae-94f8-ed1628a4f2e0', '1d6fac48-77df-4395-8a88-e1ec425baffe', '2021-10-20 16:59:55', '2021-10-20 16:59:57');
-- +migrate Down
delete from merchant_basic.branch_tag where id='5260754d-e100-4fec-8f66-ff9d2a5e3545';