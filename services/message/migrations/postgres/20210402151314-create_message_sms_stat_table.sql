-- +migrate Up
CREATE TABLE IF NOT EXISTS "message".sms_stats (
    id uuid not null,
    merchant_id uuid null,
    branch_id uuid null,
    message_type varchar(50) not null,
    area_code varchar(16) default '86',
    phone varchar(32) not null,
    sign varchar(32) default '',
    content varchar(1024),
    system varchar(100) not null,
    status varchar(10) not null,
    extra jsonb,
    created_at timestamptz(6) not null,
    updated_at timestamptz(6) not null
) PARTITION BY RANGE (created_at);
CREATE INDEX idx_message_sms_stats_phone ON "message".sms_stats (phone);
CREATE INDEX idx_message_sms_stats_created_at ON "message".sms_stats (created_at);
CREATE INDEX idx_message_sms_stats_merchant ON "message".sms_stats (merchant_id);
CREATE INDEX idx_message_sms_stats_branch ON "message".sms_stats (branch_id);
CREATE INDEX idx_message_sms_stats_message_type ON "message".sms_stats (message_type);
CREATE TABLE "message"."sms_stats_2021" PARTITION OF "message"."sms_stats" FOR
VALUES
FROM ('2021-01-01') TO ('2022-01-01');
CREATE TABLE "message"."sms_stats_2022" PARTITION OF "message"."sms_stats" FOR
VALUES
FROM ('2022-01-01') TO ('2023-01-01');
CREATE TABLE "message"."sms_stats_2023" PARTITION OF "message"."sms_stats" FOR
VALUES
FROM ('2023-01-01') TO ('2024-01-01');
CREATE TABLE "message"."sms_stats_2024" PARTITION OF "message"."sms_stats" FOR
VALUES
FROM ('2024-01-01') TO ('2025-01-01');
CREATE TABLE "message"."sms_stats_2025" PARTITION OF "message"."sms_stats" FOR
VALUES
FROM ('2025-01-01') TO ('2026-01-01');
CREATE TABLE "message"."sms_stats_2026" PARTITION OF "message"."sms_stats" FOR
VALUES
FROM ('2026-01-01') TO ('2027-01-01');
CREATE TABLE "message"."sms_stats_2027" PARTITION OF "message"."sms_stats" FOR
VALUES
FROM ('2027-01-01') TO ('2028-01-01');
CREATE TABLE "message"."sms_stats_2028" PARTITION OF "message"."sms_stats" FOR
VALUES
FROM ('2028-01-01') TO ('2029-01-01');
CREATE TABLE "message"."sms_stats_2029" PARTITION OF "message"."sms_stats" FOR
VALUES
FROM ('2029-01-01') TO ('2030-01-01');
CREATE TABLE "message"."sms_stats_other" PARTITION OF "message"."sms_stats" FOR
VALUES
FROM ('2030-01-01') TO ('2100-01-01');
CREATE INDEX idx_message_sms_stat_2021_phone ON "message".sms_stats_2021 (phone);
CREATE INDEX idx_message_sms_stat_2021_merchant ON "message".sms_stats_2021 (merchant_id);
CREATE INDEX idx_message_sms_stat_2021_branch ON "message".sms_stats_2021 (branch_id);
CREATE INDEX idx_message_sms_stat_2021_created_at ON "message".sms_stats_2021 (created_at);
CREATE INDEX idx_message_sms_stat_2022_phone ON "message".sms_stats_2022 (phone);
CREATE INDEX idx_message_sms_stat_2022_merchant ON "message".sms_stats_2022 (merchant_id);
CREATE INDEX idx_message_sms_stat_2022_branch ON "message".sms_stats_2022 (branch_id);
CREATE INDEX idx_message_sms_stat_2022_created_at ON "message".sms_stats_2022 (created_at);
CREATE INDEX idx_message_sms_stat_2023_phone ON "message".sms_stats_2023 (phone);
CREATE INDEX idx_message_sms_stat_2023_merchant ON "message".sms_stats_2023 (merchant_id);
CREATE INDEX idx_message_sms_stat_2023_branch ON "message".sms_stats_2023 (branch_id);
CREATE INDEX idx_message_sms_stat_2023_created_at ON "message".sms_stats_2023 (created_at);
CREATE INDEX idx_message_sms_stat_2024_phone ON "message".sms_stats_2024 (phone);
CREATE INDEX idx_message_sms_stat_2024_merchant ON "message".sms_stats_2024 (merchant_id);
CREATE INDEX idx_message_sms_stat_2024_branch ON "message".sms_stats_2024 (branch_id);
CREATE INDEX idx_message_sms_stat_2024_created_at ON "message".sms_stats_2024 (created_at);
CREATE INDEX idx_message_sms_stat_2025_phone ON "message".sms_stats_2025 (phone);
CREATE INDEX idx_message_sms_stat_2025_merchant ON "message".sms_stats_2025 (merchant_id);
CREATE INDEX idx_message_sms_stat_2025_branch ON "message".sms_stats_2025 (branch_id);
CREATE INDEX idx_message_sms_stat_2025_created_at ON "message".sms_stats_2025 (created_at);
CREATE INDEX idx_message_sms_stat_2026_phone ON "message".sms_stats_2026 (phone);
CREATE INDEX idx_message_sms_stat_2026_merchant ON "message".sms_stats_2026 (merchant_id);
CREATE INDEX idx_message_sms_stat_2026_branch ON "message".sms_stats_2026 (branch_id);
CREATE INDEX idx_message_sms_stat_2026_created_at ON "message".sms_stats_2026 (created_at);
CREATE INDEX idx_message_sms_stat_2027_phone ON "message".sms_stats_2027 (phone);
CREATE INDEX idx_message_sms_stat_2027_merchant ON "message".sms_stats_2027 (merchant_id);
CREATE INDEX idx_message_sms_stat_2027_branch ON "message".sms_stats_2027 (branch_id);
CREATE INDEX idx_message_sms_stat_2027_created_at ON "message".sms_stats_2027 (created_at);
CREATE INDEX idx_message_sms_stat_2028_phone ON "message".sms_stats_2028 (phone);
CREATE INDEX idx_message_sms_stat_2028_merchant ON "message".sms_stats_2028 (merchant_id);
CREATE INDEX idx_message_sms_stat_2028_branch ON "message".sms_stats_2028 (branch_id);
CREATE INDEX idx_message_sms_stat_2028_created_at ON "message".sms_stats_2028 (created_at);
CREATE INDEX idx_message_sms_stat_2029_phone ON "message".sms_stats_2029 (phone);
CREATE INDEX idx_message_sms_stat_2029_merchant ON "message".sms_stats_2029 (merchant_id);
CREATE INDEX idx_message_sms_stat_2029_branch ON "message".sms_stats_2029 (branch_id);
CREATE INDEX idx_message_sms_stat_2029_created_at ON "message".sms_stats_2029 (created_at);
CREATE INDEX idx_message_sms_stat_other_phone ON "message".sms_stats_other (phone);
CREATE INDEX idx_message_sms_stat_other_merchant ON "message".sms_stats_other (merchant_id);
CREATE INDEX idx_message_sms_stat_other_branch ON "message".sms_stats_other (branch_id);
CREATE INDEX idx_message_sms_stat_other_created_at ON "message".sms_stats_other (created_at);
COMMENT ON TABLE "message".sms_stats IS '短信发送记录';
COMMENT ON COLUMN "message".sms_stats.phone IS '电话';
COMMENT ON COLUMN "message".sms_stats.sign IS '签章';
COMMENT ON COLUMN "message".sms_stats.content IS '内容';
COMMENT ON COLUMN "message".sms_stats.system IS '调用系统';
COMMENT ON COLUMN "message".sms_stats.status IS '结果 success fail';
-- +migrate Down
DROP TABLE IF EXISTS "message"."sms_stats_other";
DROP TABLE IF EXISTS "message"."sms_stats_2029";
DROP TABLE IF EXISTS "message"."sms_stats_2028";
DROP TABLE IF EXISTS "message"."sms_stats_2027";
DROP TABLE IF EXISTS "message"."sms_stats_2026";
DROP TABLE IF EXISTS "message"."sms_stats_2025";
DROP TABLE IF EXISTS "message"."sms_stats_2024";
DROP TABLE IF EXISTS "message"."sms_stats_2023";
DROP TABLE IF EXISTS "message"."sms_stats_2022";
DROP TABLE IF EXISTS "message"."sms_stats_2021";
DROP TABLE IF EXISTS "message"."sms_stats";