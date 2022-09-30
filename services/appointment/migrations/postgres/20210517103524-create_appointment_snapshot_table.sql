
-- +migrate Up
CREATE TABLE "appointment"."snapshot" (
  "id" uuid NOT NULL,
  "staff_id" uuid,
  "sleuth_code" varchar(20) COLLATE "pg_catalog"."default",
  "table_name" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "table_id" uuid,
  "method" varchar(50) COLLATE "pg_catalog"."default",
  "before" jsonb,
  "after" jsonb,
  "related_id" uuid,
  "created_at" timestamptz(6)
)PARTITION BY RANGE (created_at)
;
COMMENT ON COLUMN "appointment"."snapshot"."staff_id" IS '操作员工';
COMMENT ON COLUMN "appointment"."snapshot"."sleuth_code" IS '调用链唯一码';
COMMENT ON COLUMN "appointment"."snapshot"."table_name" IS '表名';
COMMENT ON COLUMN "appointment"."snapshot"."table_id" IS '操作数据id';
COMMENT ON COLUMN "appointment"."snapshot"."created_at" IS '数据创建时间';
COMMENT ON COLUMN "appointment"."snapshot"."method" IS '增、删、改';
COMMENT ON COLUMN "appointment"."snapshot"."before" IS '修改前数据';
COMMENT ON COLUMN "appointment"."snapshot"."after" IS '修改后数据';
COMMENT ON TABLE "appointment"."snapshot" IS '快照表';

CREATE TABLE "appointment"."snapshot_2021" PARTITION OF "appointment"."snapshot" FOR VALUES FROM ('2021-01-01') TO ('2022-01-01');
CREATE TABLE "appointment"."snapshot_2022" PARTITION OF "appointment"."snapshot" FOR VALUES FROM ('2022-01-01') TO ('2023-01-01');
CREATE TABLE "appointment"."snapshot_2023" PARTITION OF "appointment"."snapshot" FOR VALUES FROM ('2023-01-01') TO ('2024-01-01');
CREATE TABLE "appointment"."snapshot_2024" PARTITION OF "appointment"."snapshot" FOR VALUES FROM ('2024-01-01') TO ('2025-01-01');
CREATE TABLE "appointment"."snapshot_2025" PARTITION OF "appointment"."snapshot" FOR VALUES FROM ('2025-01-01') TO ('2026-01-01');
CREATE TABLE "appointment"."snapshot_2026" PARTITION OF "appointment"."snapshot" FOR VALUES FROM ('2026-01-01') TO ('2027-01-01');
CREATE TABLE "appointment"."snapshot_2027" PARTITION OF "appointment"."snapshot" FOR VALUES FROM ('2027-01-01') TO ('2028-01-01');
CREATE TABLE "appointment"."snapshot_2028" PARTITION OF "appointment"."snapshot" FOR VALUES FROM ('2028-01-01') TO ('2029-01-01');
CREATE TABLE "appointment"."snapshot_2029" PARTITION OF "appointment"."snapshot" FOR VALUES FROM ('2029-01-01') TO ('2030-01-01');
CREATE TABLE "appointment"."snapshot_other" PARTITION OF "appointment"."snapshot" FOR VALUES FROM ('2030-01-01') TO ('2100-01-01');
CREATE INDEX "idx_appointment_2021_table_id" ON "appointment"."snapshot_2021" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2021_table_name" ON "appointment"."snapshot_2021" USING btree (
  "table_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2021_created_at" ON "appointment"."snapshot_2021" (
  "created_at"
);

CREATE INDEX "idx_appointment_2022_table_id" ON "appointment"."snapshot_2022" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2022_table_name" ON "appointment"."snapshot_2022" USING btree (
  "table_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2022_created_at" ON "appointment"."snapshot_2022" (
  "created_at"
);

CREATE INDEX "idx_appointment_2023_table_id" ON "appointment"."snapshot_2023" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2023_table_name" ON "appointment"."snapshot_2023" USING btree (
  "table_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2023_created_at" ON "appointment"."snapshot_2023" (
  "created_at"
);

CREATE INDEX "idx_appointment_2024_table_id" ON "appointment"."snapshot_2024" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2024_table_name" ON "appointment"."snapshot_2024" USING btree (
  "table_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2024_created_at" ON "appointment"."snapshot_2024" (
  "created_at"
);

CREATE INDEX "idx_appointment_2025_table_id" ON "appointment"."snapshot_2025" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2025_table_name" ON "appointment"."snapshot_2025" USING btree (
  "table_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2025_created_at" ON "appointment"."snapshot_2025" (
  "created_at"
);

CREATE INDEX "idx_appointment_2026_table_id" ON "appointment"."snapshot_2026" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2026_table_name" ON "appointment"."snapshot_2026" USING btree (
  "table_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2026_created_at" ON "appointment"."snapshot_2026" (
  "created_at"
);

CREATE INDEX "idx_appointment_2027_table_id" ON "appointment"."snapshot_2027" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2027_table_name" ON "appointment"."snapshot_2027" USING btree (
  "table_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2027_created_at" ON "appointment"."snapshot_2027" (
  "created_at"
);

CREATE INDEX "idx_appointment_2028_table_id" ON "appointment"."snapshot_2028" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2028_table_name" ON "appointment"."snapshot_2028" USING btree (
  "table_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2028_created_at" ON "appointment"."snapshot_2028" (
  "created_at"
);

CREATE INDEX "idx_appointment_2029_table_id" ON "appointment"."snapshot_2029" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2029_table_name" ON "appointment"."snapshot_2029" USING btree (
  "table_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_2029_created_at" ON "appointment"."snapshot_2029" (
  "created_at"
);

CREATE INDEX "idx_appointment_other_table_id" ON "appointment"."snapshot_other" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_other_table_name" ON "appointment"."snapshot_other" USING btree (
  "table_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_appointment_other_created_at" ON "appointment"."snapshot_other" (
  "created_at"
);

-- +migrate Down
DROP TABLE IF EXISTS "appointment"."snapshot_other";
DROP TABLE IF EXISTS "appointment"."snapshot_2029";
DROP TABLE IF EXISTS "appointment"."snapshot_2028";
DROP TABLE IF EXISTS "appointment"."snapshot_2027";
DROP TABLE IF EXISTS "appointment"."snapshot_2026";
DROP TABLE IF EXISTS "appointment"."snapshot_2025";
DROP TABLE IF EXISTS "appointment"."snapshot_2024";
DROP TABLE IF EXISTS "appointment"."snapshot_2023";
DROP TABLE IF EXISTS "appointment"."snapshot_2022";
DROP TABLE IF EXISTS "appointment"."snapshot_2021";
DROP TABLE IF EXISTS "appointment"."snapshot";