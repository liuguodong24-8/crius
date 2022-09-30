
-- +migrate Up
CREATE TABLE IF NOT EXISTS "appointment"."appointment" (
  "id" uuid NOT NULL,
  "member_id" uuid,
  "merchant_id" uuid NOT NULL,
  "branch_id" uuid NOT NULL,
  "called_code" varchar(10) COLLATE "pg_catalog"."default",
  "called_phone" varchar(20) COLLATE "pg_catalog"."default",
  "appointment_code" varchar(10) COLLATE "pg_catalog"."default",
  "appointment_phone" varchar(20) COLLATE "pg_catalog"."default",
  "name" varchar(50) COLLATE "pg_catalog"."default",
  "gender" int2,
  "way" int2,
  "customer_num" int4,
  "appointment_at" timestamptz(6),
  "appointment_date" date,
  "room_type_id" uuid,
  "deposit_fee" int8,
  "flower_cake" bool,
  "flower_cake_remark" varchar(400) COLLATE "pg_catalog"."default",
  "remark" varchar(400) COLLATE "pg_catalog"."default",
  "operator" uuid,
  "status" varchar(20) COLLATE "pg_catalog"."default",
  "cancelled_reason" varchar(400) COLLATE "pg_catalog"."default",
  "related_id" uuid,
  "date_counter_id" uuid,
  "charging_way" int2,
  "trade_id" uuid,
  "expire_at" timestamptz(6),
  "keep_at" timestamptz(6),
  "cancel_at" timestamptz(6),
  "trade_type" varchar(50) COLLATE "pg_catalog"."default",
  "arrived_at" timestamptz(6),
  "breach" bool,
  "breach_reason" varchar(200),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "appointment"."appointment"."member_id" IS '会员id';
COMMENT ON COLUMN "appointment"."appointment"."branch_id" IS '预约门店';
COMMENT ON COLUMN "appointment"."appointment"."called_code" IS '来电区号';
COMMENT ON COLUMN "appointment"."appointment"."called_phone" IS '来电号码';
COMMENT ON COLUMN "appointment"."appointment"."appointment_code" IS '预约区号';
COMMENT ON COLUMN "appointment"."appointment"."appointment_phone" IS '预约号码';
COMMENT ON COLUMN "appointment"."appointment"."name" IS '名字';
COMMENT ON COLUMN "appointment"."appointment"."gender" IS '性别';
COMMENT ON COLUMN "appointment"."appointment"."way" IS '预约方式 1微信 2电话 4店长 8客服';
COMMENT ON COLUMN "appointment"."appointment"."customer_num" IS '客户人数';
COMMENT ON COLUMN "appointment"."appointment"."appointment_at" IS '预约时间';
COMMENT ON COLUMN "appointment"."appointment"."appointment_date" IS '预约营业日';
COMMENT ON COLUMN "appointment"."appointment"."expire_at" IS '未支付过期时间';
COMMENT ON COLUMN "appointment"."appointment"."room_type_id" IS '预约房型';
COMMENT ON COLUMN "appointment"."appointment"."deposit_fee" IS '订金';
COMMENT ON COLUMN "appointment"."appointment"."flower_cake" IS '鲜花蛋糕代收';
COMMENT ON COLUMN "appointment"."appointment"."flower_cake_remark" IS '鲜花蛋糕代收备注';
COMMENT ON COLUMN "appointment"."appointment"."remark" IS '备注';
COMMENT ON COLUMN "appointment"."appointment"."operator" IS '操作人';
COMMENT ON COLUMN "appointment"."appointment"."status" IS '状态 待支付 arrearage、已预约 appointed、已到店 arrived、已取消 cancelled、退款中 refunding 、退款成功 refunded、已过期 expired';
COMMENT ON COLUMN "appointment"."appointment"."cancelled_reason" IS '取消原因';
COMMENT ON COLUMN "appointment"."appointment"."related_id" IS '快照关联id';
COMMENT ON COLUMN "appointment"."appointment"."date_counter_id" IS '关联id';
COMMENT ON COLUMN "appointment"."appointment"."charging_way" IS '支付方式 1线上 2线下';
COMMENT ON COLUMN "appointment"."appointment"."trade_id" IS '交易id';
COMMENT ON COLUMN "appointment"."appointment"."keep_at" IS '锁定房型时间';
COMMENT ON COLUMN "appointment"."appointment"."cancel_at" IS '取消时间';
COMMENT ON COLUMN "appointment"."appointment"."trade_type" IS '线上支付方式 支付宝 alipay  微信 wechat';
COMMENT ON COLUMN "appointment"."appointment"."arrived_at" IS '到店时间';
COMMENT ON COLUMN "appointment"."appointment"."breach" IS '是否违约';
COMMENT ON COLUMN "appointment"."appointment"."breach_reason" IS '违约原因';
COMMENT ON TABLE "appointment"."appointment" IS '预约表';

CREATE INDEX "appointment_appointment_code_appointment_phone_status_idx" ON "appointment"."appointment" USING btree (
  "appointment_code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "appointment_phone" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "status" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "appointment_appointment_phone_idx" ON "appointment"."appointment" USING btree (
  "appointment_phone" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "appointment_branch_id_idx" ON "appointment"."appointment" USING btree (
  "branch_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "appointment_room_type_id_idx" ON "appointment"."appointment" USING btree (
  "room_type_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);

ALTER TABLE "appointment"."appointment" ADD CONSTRAINT "appointment_pkey" PRIMARY KEY ("id");

-- +migrate Down
DROP TABLE IF EXISTS "appointment"."appointment";