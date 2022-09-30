
-- +migrate Up
CREATE TABLE "appointment"."appointment_caller" (
  "id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "phone_code" varchar(10) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "phone" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "phone_tail" char(1) COLLATE "pg_catalog"."default" NOT NULL,
  "phone_suffix" char(4) COLLATE "pg_catalog"."default" NOT NULL,
  "caller_name" varchar(20) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "tags" jsonb,
  "gender" int2 DEFAULT 0,
  "is_black" bool DEFAULT false,
  "black_reason" varchar(256) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "last_call_at" timestamptz(6),
  "last_operator" uuid,
  "last_call_action" varchar(20) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
PARTITION BY LIST (
  "phone_tail" COLLATE "pg_catalog"."default" "pg_catalog"."bpchar_ops"
)
;
COMMENT ON COLUMN "appointment"."appointment_caller"."phone_code" IS '手机区号';
COMMENT ON COLUMN "appointment"."appointment_caller"."phone" IS '手机号';
COMMENT ON COLUMN "appointment"."appointment_caller"."phone_suffix" IS '手机尾号';
COMMENT ON COLUMN "appointment"."appointment_caller"."caller_name" IS '姓名';
COMMENT ON COLUMN "appointment"."appointment_caller"."tags" IS '标签[{"tag":"阿巴巴巴","color":1},{"tag":"阿巴啊巴啊巴","color":2}]';
COMMENT ON COLUMN "appointment"."appointment_caller"."gender" IS '性别';
COMMENT ON COLUMN "appointment"."appointment_caller"."is_black" IS '是否黑名单';
COMMENT ON COLUMN "appointment"."appointment_caller"."black_reason" IS '黑名单原因';
COMMENT ON COLUMN "appointment"."appointment_caller"."last_call_at" IS '上次来电时间';
COMMENT ON COLUMN "appointment"."appointment_caller"."last_operator" IS '上次操作人';
COMMENT ON COLUMN "appointment"."appointment_caller"."last_call_action" IS '上次来电操作';
COMMENT ON TABLE "appointment"."appointment_caller" IS '来电用户表';

-- +migrate StatementBegin
create or replace function create_appointment_caller_sub() returns void as $BODY$
declare i integer;
declare table_name varchar;
begin
  i = 0;
  for i in 0..9 LOOP
    table_name = 'appointment_caller_' || i;
    EXECUTE FORMAT('CREATE TABLE IF NOT EXISTS %I PARTITION OF "appointment"."appointment_caller" FOR VALUES IN (%s)', table_name , i);

    EXECUTE format('CREATE INDEX idx_appointment_caller_%s_phone ON %I (phone);', i, table_name);
    EXECUTE format('CREATE INDEX idx_appointment_caller_%s_phone_suffix ON %I (phone_suffix);', i, table_name);
    EXECUTE format('CREATE INDEX idx_appointment_caller_%s_name ON %I (caller_name);', i, table_name);
    EXECUTE format('CREATE INDEX idx_appointment_caller_%s_call_at ON %I (last_call_at);', i, table_name);
    end LOOP;
    return;
end;
$BODY$ LANGUAGE plpgsql;

-- +migrate StatementEnd

select create_appointment_caller_sub();

-- +migrate Down
DROP FUNCTION IF EXISTS create_appointment_caller_sub();
DROP TABLE IF EXISTS "appointment"."appointment_caller";

