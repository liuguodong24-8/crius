
-- +migrate Up
CREATE TABLE "member_account"."snapshot" (
  "id" uuid NOT NULL,
  "staff_id" uuid,
  "sleuth_code" varchar(20) COLLATE "pg_catalog"."default",
  "table_name" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "table_id" uuid NOT NULL,
  "method" varchar(50) COLLATE "pg_catalog"."default",
  "before" jsonb,
  "after" jsonb,
  "created_at" timestamptz(6)
)PARTITION BY RANGE (created_at)
;
COMMENT ON COLUMN "member_account"."snapshot"."staff_id" IS '操作员工';
COMMENT ON COLUMN "member_account"."snapshot"."sleuth_code" IS '调用链唯一码';
COMMENT ON COLUMN "member_account"."snapshot"."table_name" IS '表名';
COMMENT ON COLUMN "member_account"."snapshot"."table_id" IS '操作数据id';
COMMENT ON COLUMN "member_account"."snapshot"."method" IS '增、删、改';
COMMENT ON COLUMN "member_account"."snapshot"."before" IS '修改前数据';
COMMENT ON COLUMN "member_account"."snapshot"."after" IS '修改后数据';
COMMENT ON COLUMN "member_account"."snapshot"."created_at" IS '数据创建时间';
COMMENT ON TABLE "member_account"."snapshot" IS '快照表';

CREATE INDEX "index_table_id" ON "member_account"."snapshot" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "index_snapshot_created_at" ON "member_account"."snapshot" (
  "created_at"
);

CREATE TABLE "member_account"."snapshot_2021" PARTITION OF "member_account"."snapshot" FOR VALUES FROM ('2021-01-01') TO ('2022-01-01');
CREATE TABLE "member_account"."snapshot_2022" PARTITION OF "member_account"."snapshot" FOR VALUES FROM ('2022-01-01') TO ('2023-01-01');
CREATE TABLE "member_account"."snapshot_2023" PARTITION OF "member_account"."snapshot" FOR VALUES FROM ('2023-01-01') TO ('2024-01-01');
CREATE TABLE "member_account"."snapshot_2024" PARTITION OF "member_account"."snapshot" FOR VALUES FROM ('2024-01-01') TO ('2025-01-01');
CREATE TABLE "member_account"."snapshot_2025" PARTITION OF "member_account"."snapshot" FOR VALUES FROM ('2025-01-01') TO ('2026-01-01');
CREATE TABLE "member_account"."snapshot_2026" PARTITION OF "member_account"."snapshot" FOR VALUES FROM ('2026-01-01') TO ('2027-01-01');
CREATE TABLE "member_account"."snapshot_2027" PARTITION OF "member_account"."snapshot" FOR VALUES FROM ('2027-01-01') TO ('2028-01-01');
CREATE TABLE "member_account"."snapshot_2028" PARTITION OF "member_account"."snapshot" FOR VALUES FROM ('2028-01-01') TO ('2029-01-01');
CREATE TABLE "member_account"."snapshot_2029" PARTITION OF "member_account"."snapshot" FOR VALUES FROM ('2029-01-01') TO ('2030-01-01');
CREATE TABLE "member_account"."snapshot_other" PARTITION OF "member_account"."snapshot" FOR VALUES FROM ('2030-01-01') TO ('2100-01-01');

CREATE INDEX "index_table_id_2021" ON "member_account"."snapshot_2021" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "index_snapshot_2021_created_at" ON "member_account"."snapshot_2021" (
  "created_at"
);
CREATE INDEX "index_table_id_2022" ON "member_account"."snapshot_2022" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "index_snapshot_2022_created_at" ON "member_account"."snapshot_2022" (
  "created_at"
);
CREATE INDEX "index_table_id_2023" ON "member_account"."snapshot_2023" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "index_snapshot_2023_created_at" ON "member_account"."snapshot_2023" (
  "created_at"
);
CREATE INDEX "index_table_id_2024" ON "member_account"."snapshot_2024" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "index_snapshot_2024_created_at" ON "member_account"."snapshot_2024" (
  "created_at"
);
CREATE INDEX "index_table_id_2025" ON "member_account"."snapshot_2025" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "index_snapshot_2025_created_at" ON "member_account"."snapshot_2025" (
  "created_at"
);
CREATE INDEX "index_table_id_2026" ON "member_account"."snapshot_2026" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "index_snapshot_2026_created_at" ON "member_account"."snapshot_2026" (
  "created_at"
);
CREATE INDEX "index_table_id_2027" ON "member_account"."snapshot_2027" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "index_snapshot_2027_created_at" ON "member_account"."snapshot_2027" (
  "created_at"
);
CREATE INDEX "index_table_id_2028" ON "member_account"."snapshot_2028" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "index_snapshot_2028_created_at" ON "member_account"."snapshot_2028" (
  "created_at"
);
CREATE INDEX "index_table_id_2029" ON "member_account"."snapshot_2029" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "index_snapshot_2029_created_at" ON "member_account"."snapshot_2029" (
  "created_at"
);
CREATE INDEX "index_table_id_other" ON "member_account"."snapshot_other" USING btree (
  "table_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "index_snapshot_other_created_at" ON "member_account"."snapshot_other" (
  "created_at"
);

-- +migrate Down
DROP TABLE IF EXISTS "member_account"."snapshot_other";
DROP TABLE IF EXISTS "member_account"."snapshot_2029";
DROP TABLE IF EXISTS "member_account"."snapshot_2028";
DROP TABLE IF EXISTS "member_account"."snapshot_2027";
DROP TABLE IF EXISTS "member_account"."snapshot_2026";
DROP TABLE IF EXISTS "member_account"."snapshot_2025";
DROP TABLE IF EXISTS "member_account"."snapshot_2024";
DROP TABLE IF EXISTS "member_account"."snapshot_2023";
DROP TABLE IF EXISTS "member_account"."snapshot_2022";
DROP TABLE IF EXISTS "member_account"."snapshot_2021";
DROP TABLE IF EXISTS "member_account"."snapshot";