
-- +migrate Up
CREATE TABLE "appointment"."appointment_caller_records" (
  "id" uuid,
  "merchant_id" uuid NOT NULL,
  "caller_id" uuid NOT NULL,
  "phone" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "operator" uuid,
  "call_action" varchar(20) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "call_at" timestamptz(6),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
PARTITION BY RANGE (
  "call_at" "pg_catalog"."timestamptz_ops"
)
;
COMMENT ON TABLE "appointment"."appointment_caller_records" IS '来电用户记录';
COMMENT ON COLUMN "appointment"."appointment_caller_records"."phone" IS '手机号';
COMMENT ON COLUMN "appointment"."appointment_caller_records"."caller_id" IS '来电用户表ID';

CREATE TABLE "appointment"."appointment_caller_records_2021" PARTITION OF "appointment"."appointment_caller_records" FOR VALUES FROM ('2021-01-01') TO ('2022-01-01');
CREATE TABLE "appointment"."appointment_caller_records_2022" PARTITION OF "appointment"."appointment_caller_records" FOR VALUES FROM ('2022-01-01') TO ('2023-01-01');
CREATE TABLE "appointment"."appointment_caller_records_2023" PARTITION OF "appointment"."appointment_caller_records" FOR VALUES FROM ('2023-01-01') TO ('2024-01-01');
CREATE TABLE "appointment"."appointment_caller_records_2024" PARTITION OF "appointment"."appointment_caller_records" FOR VALUES FROM ('2024-01-01') TO ('2025-01-01');
CREATE TABLE "appointment"."appointment_caller_records_2025" PARTITION OF "appointment"."appointment_caller_records" FOR VALUES FROM ('2025-01-01') TO ('2026-01-01');
CREATE TABLE "appointment"."appointment_caller_records_2026" PARTITION OF "appointment"."appointment_caller_records" FOR VALUES FROM ('2026-01-01') TO ('2027-01-01');
CREATE TABLE "appointment"."appointment_caller_records_2027" PARTITION OF "appointment"."appointment_caller_records" FOR VALUES FROM ('2027-01-01') TO ('2028-01-01');
CREATE TABLE "appointment"."appointment_caller_records_2028" PARTITION OF "appointment"."appointment_caller_records" FOR VALUES FROM ('2028-01-01') TO ('2029-01-01');
CREATE TABLE "appointment"."appointment_caller_records_2029" PARTITION OF "appointment"."appointment_caller_records" FOR VALUES FROM ('2029-01-01') TO ('2030-01-01');
CREATE TABLE "appointment"."appointment_caller_records_other" PARTITION OF "appointment"."appointment_caller_records" FOR VALUES FROM ('2030-01-01') TO ('2100-01-01');

create index idx_appointment_call_record_2021_phone on "appointment"."appointment_caller_records_2021"(phone);
create index idx_appointment_call_record_2021_caller on "appointment"."appointment_caller_records_2021"(caller_id);
create index idx_appointment_call_record_2021_call_at on "appointment"."appointment_caller_records_2021"(call_at);

create index idx_appointment_call_record_2022_phone on "appointment"."appointment_caller_records_2022"(phone);
create index idx_appointment_call_record_2022_caller on "appointment"."appointment_caller_records_2022"(caller_id);
create index idx_appointment_call_record_2022_call_at on "appointment"."appointment_caller_records_2022"(call_at);

create index idx_appointment_call_record_2023_phone on "appointment"."appointment_caller_records_2023"(phone);
create index idx_appointment_call_record_2023_caller on "appointment"."appointment_caller_records_2023"(caller_id);
create index idx_appointment_call_record_2023_call_at on "appointment"."appointment_caller_records_2023"(call_at);

create index idx_appointment_call_record_2024_phone on "appointment"."appointment_caller_records_2024"(phone);
create index idx_appointment_call_record_2024_caller on "appointment"."appointment_caller_records_2024"(caller_id);
create index idx_appointment_call_record_2024_call_at on "appointment"."appointment_caller_records_2024"(call_at);

create index idx_appointment_call_record_2025_phone on "appointment"."appointment_caller_records_2025"(phone);
create index idx_appointment_call_record_2025_caller on "appointment"."appointment_caller_records_2025"(caller_id);
create index idx_appointment_call_record_2025_call_at on "appointment"."appointment_caller_records_2025"(call_at);

create index idx_appointment_call_record_2026_phone on "appointment"."appointment_caller_records_2026"(phone);
create index idx_appointment_call_record_2026_caller on "appointment"."appointment_caller_records_2026"(caller_id);
create index idx_appointment_call_record_2026_call_at on "appointment"."appointment_caller_records_2026"(call_at);

create index idx_appointment_call_record_2027_phone on "appointment"."appointment_caller_records_2027"(phone);
create index idx_appointment_call_record_2027_caller on "appointment"."appointment_caller_records_2027"(caller_id);
create index idx_appointment_call_record_2027_call_at on "appointment"."appointment_caller_records_2027"(call_at);

create index idx_appointment_call_record_2028_phone on "appointment"."appointment_caller_records_2028"(phone);
create index idx_appointment_call_record_2028_caller on "appointment"."appointment_caller_records_2028"(caller_id);
create index idx_appointment_call_record_2028_call_at on "appointment"."appointment_caller_records_2028"(call_at);

create index idx_appointment_call_record_2029_phone on "appointment"."appointment_caller_records_2029"(phone);
create index idx_appointment_call_record_2029_caller on "appointment"."appointment_caller_records_2029"(caller_id);
create index idx_appointment_call_record_2029_call_at on "appointment"."appointment_caller_records_2029"(call_at);

create index idx_appointment_call_record_other_phone on "appointment"."appointment_caller_records_other"(phone);
create index idx_appointment_call_record_other_caller on "appointment"."appointment_caller_records_other"(caller_id);
create index idx_appointment_call_record_other_call_at on "appointment"."appointment_caller_records_other"(call_at);

-- +migrate Down
DROP TABLE IF EXISTS "appointment"."appointment_caller_records_other";
DROP TABLE IF EXISTS "appointment"."appointment_caller_records_2029";
DROP TABLE IF EXISTS "appointment"."appointment_caller_records_2028";
DROP TABLE IF EXISTS "appointment"."appointment_caller_records_2027";
DROP TABLE IF EXISTS "appointment"."appointment_caller_records_2026";
DROP TABLE IF EXISTS "appointment"."appointment_caller_records_2025";
DROP TABLE IF EXISTS "appointment"."appointment_caller_records_2024";
DROP TABLE IF EXISTS "appointment"."appointment_caller_records_2023";
DROP TABLE IF EXISTS "appointment"."appointment_caller_records_2022";
DROP TABLE IF EXISTS "appointment"."appointment_caller_records_2021";
DROP TABLE IF EXISTS "appointment"."appointment_caller_records";

