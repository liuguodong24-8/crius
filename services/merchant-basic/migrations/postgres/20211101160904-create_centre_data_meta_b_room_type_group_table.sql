
-- +migrate Up
CREATE TABLE if not exists "centre_data"."meta_b_room_type_group" (
    "room_type_group_id" uuid PRIMARY KEY,
    "room_type_group_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
    "room_type_ids" uuid[] NOT NULL,
    "customer_count_min" int4 NOT NULL,
    "customer_count_max" int4 NOT NULL,
    "weight" int4 NOT NULL,
    "type" varchar(16) COLLATE "pg_catalog"."default" DEFAULT 'room'::character varying,
    "short_name" varchar(255) COLLATE "pg_catalog"."default",
    "create_time" timestamp(6) NOT NULL,
    "update_time" timestamp(6),
    "delete_time" timestamp(6),
    CONSTRAINT "uk_room_type_group" UNIQUE ("room_type_group_name")
)
;

COMMENT ON COLUMN "centre_data"."meta_b_room_type_group"."room_type_group_id" IS '房型组主键';

COMMENT ON COLUMN "centre_data"."meta_b_room_type_group"."room_type_group_name" IS '房型组名';

COMMENT ON COLUMN "centre_data"."meta_b_room_type_group"."room_type_ids" IS '房型列表';

COMMENT ON COLUMN "centre_data"."meta_b_room_type_group"."customer_count_min" IS '标准最小人数';

COMMENT ON COLUMN "centre_data"."meta_b_room_type_group"."customer_count_max" IS '标准最大人数';

COMMENT ON COLUMN "centre_data"."meta_b_room_type_group"."weight" IS '排序权重';

COMMENT ON COLUMN "centre_data"."meta_b_room_type_group"."type" IS '房台类型 房间|桌台';

COMMENT ON COLUMN "centre_data"."meta_b_room_type_group"."short_name" IS '简称';

COMMENT ON TABLE "centre_data"."meta_b_room_type_group" IS '中央基础数据-房型组';

-- +migrate Down
DROP TABLE IF EXISTS "centre_data"."meta_b_room_type_group";