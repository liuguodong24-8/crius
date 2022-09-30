
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."member_behavior" (
  "id" uuid NOT NULL primary key,
  "member_id" uuid NOT NULL,
  "behavior" varchar(255) NOT NULL,
  "staff_id" uuid,
  "branch_id" uuid,
  "merchant_id" uuid NOT NULL,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
);

COMMENT ON COLUMN "merchant_basic"."member_behavior"."behavior" IS '用户交互行为 open_card...';

create index idx_member_behavior_member on "merchant_basic"."member_behavior"(member_id);
create index idx_member_behavior_created_at on "merchant_basic"."member_behavior"(created_at);

-- +migrate Down

DROP TABLE IF EXISTS "merchant_basic"."member_behavior";
