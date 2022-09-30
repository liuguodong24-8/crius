
-- +migrate Up
INSERT INTO "merchant_basic"."staff" ("id", "name", "phone", "phone_code", "gender", "status", "created_at", "updated_at", "password", "salt", "employee_code", "code", "admin") VALUES ('f24c596a-da55-43ae-94f8-ed1628a4f2e0', '管理员', '', '', '3', 'opened', now( ), now( ), '9e3709caa25a4017799bf865a90993d5', 'dhyq7s8twstlv', 'omytech', '', 'true');

-- +migrate Down
DELETE FROM "merchant_basic"."staff" WHERE "id"='f24c596a-da55-43ae-94f8-ed1628a4f2e0';