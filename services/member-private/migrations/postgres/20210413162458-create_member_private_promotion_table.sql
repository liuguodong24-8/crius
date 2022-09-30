
-- +migrate Up
CREATE TABLE "member_private"."promotions" (
  "id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "name" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "begin_at" date,
  "end_at" date,
  "status" varchar(10) COLLATE "pg_catalog"."default" NOT NULL DEFAULT 'opened'::character varying,
  "extra" jsonb,
  "branch_ids" UUID[] null,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL,
  "deleted_at" timestamptz(6),
  CONSTRAINT "promotions_pkey" PRIMARY KEY ("id")
)
;

create index idx_member_private_promotion_created_at on "member_private"."promotions"(created_at);
create index idx_member_private_promotion_merchant on "member_private"."promotions"(merchant_id);
create index idx_member_private_promotion_name on "member_private"."promotions"(name);

COMMENT ON COLUMN "member_private"."promotions"."id" IS '优惠方案组';
COMMENT ON COLUMN "member_private"."promotions"."merchant_id" IS '商户ID';
COMMENT ON COLUMN "member_private"."promotions"."name" IS '方案组名';
COMMENT ON COLUMN "member_private"."promotions"."begin_at" IS '生效时间 Y-m-d';
COMMENT ON COLUMN "member_private"."promotions"."end_at" IS '失效时间 Y-m-d';
COMMENT ON COLUMN "member_private"."promotions"."status" IS '状态 opened启用 closed禁用';
-- +migrate Down
DROP TABLE IF EXISTS "member_private"."promotions";