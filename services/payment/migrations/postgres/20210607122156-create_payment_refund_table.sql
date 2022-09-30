
-- +migrate Up
CREATE TABLE "payment"."wechat_refunds" (
  "id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "branch_id" uuid NOT NULL,
  "refund_fee" int8 DEFAULT 0,
  "wechat_trade_id" uuid NOT NULL,
  "refund_id" varchar(64) default '',
  "out_refund_no" varchar(32) COLLATE "pg_catalog"."default",
  "notify_url" varchar(64) COLLATE "pg_catalog"."default",
  "request_params" jsonb,
  "wechat_request" jsonb,
  "wechat_response" jsonb,
  "notify_content" jsonb,
  "query_content" jsonb,
  "refund_status" varchar COLLATE "pg_catalog"."default" DEFAULT 'init'::character varying,
  "notify_state" int2 DEFAULT 0,
  "extra" jsonb,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  CONSTRAINT "wechat_refunds_pkey" PRIMARY KEY ("id")
)
;
CREATE INDEX "idx_payment_wechat_redund_trade" ON "payment"."wechat_refunds" USING btree (
  "wechat_trade_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);

CREATE INDEX "idx_payment_wechat_refund_create" ON "payment"."wechat_refunds" USING btree (
  "created_at" "pg_catalog"."timestamp_ops" ASC NULLS LAST
);

CREATE INDEX "idx_payment_wechat_refund_out_refund_no" ON "payment"."wechat_refunds" USING btree (
  "out_refund_no" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

COMMENT ON COLUMN "payment"."wechat_refunds"."wechat_trade_id" IS '交易ID';
COMMENT ON COLUMN "payment"."wechat_refunds"."refund_id" IS '微信退款ID';
COMMENT ON COLUMN "payment"."wechat_refunds"."out_refund_no" IS '商户退款号';

COMMENT ON COLUMN "payment"."wechat_refunds"."wechat_request" IS '微信请求参数';

COMMENT ON COLUMN "payment"."wechat_refunds"."wechat_response" IS '微信请求返回';

COMMENT ON COLUMN "payment"."wechat_refunds"."notify_content" IS '微信回调参数';

COMMENT ON COLUMN "payment"."wechat_refunds"."refund_status" IS '退款状态 init success fail';

COMMENT ON COLUMN "payment"."wechat_refunds"."notify_state" IS '通知状态';

COMMENT ON TABLE "payment"."wechat_refunds" IS '微信退款';
-- +migrate Down
DROP TABLE IF EXISTS "payment"."wechat_refunds";