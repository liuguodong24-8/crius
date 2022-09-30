
-- +migrate Up
CREATE TABLE "payment"."wechat_trades" (
  "id" uuid NOT NULL,
  "merchant_id" uuid NOT NULL,
  "branch_id" uuid NOT NULL,
  "total_fee" int8 DEFAULT 0,
  "trade_type" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "transaction_id" varchar(32) COLLATE "pg_catalog"."default",
  "out_trade_no" varchar(32) COLLATE "pg_catalog"."default",
  "notify_url" varchar(256) COLLATE "pg_catalog"."default",
  "request_params" jsonb,
  "wechat_request" jsonb,
  "wechat_response" jsonb,
  "notify_content" jsonb,
  "query_content" jsonb,
  "trade_status" varchar DEFAULT 'init',
  "notify_state" int2 DEFAULT 0,
  "extra" jsonb,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  CONSTRAINT "wechat_trades_pkey" PRIMARY KEY ("id")
)
;

CREATE INDEX "idx_payment_wechat_trade_create" ON "payment"."wechat_trades" USING btree (
  "created_at" "pg_catalog"."timestamp_ops" ASC NULLS LAST
);

CREATE INDEX "idx_payment_wechat_trade_out_trade_no" ON "payment"."wechat_trades" USING btree (
  "out_trade_no" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

CREATE INDEX "idx_payment_wechat_trade_transaction" ON "payment"."wechat_trades" USING btree (
  "transaction_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);


COMMENT ON COLUMN "payment"."wechat_trades"."trade_type" IS '交易类型 MWEB JSAPI NATIVE APP';
COMMENT ON COLUMN "payment"."wechat_trades"."transaction_id" IS '微信的订单号';
COMMENT ON COLUMN "payment"."wechat_trades"."out_trade_no" IS '商户订单号';
COMMENT ON COLUMN "payment"."wechat_trades"."request_params" IS '请求参数';
COMMENT ON COLUMN "payment"."wechat_trades"."wechat_request" IS '微信请求参数';
COMMENT ON COLUMN "payment"."wechat_trades"."wechat_response" IS '微信请求返回';
COMMENT ON COLUMN "payment"."wechat_trades"."notify_content" IS '微信回调参数';
COMMENT ON COLUMN "payment"."wechat_trades"."query_content" IS '微信查询返回';
COMMENT ON COLUMN "payment"."wechat_trades"."trade_status" IS '支付状态 init success fail';
COMMENT ON COLUMN "payment"."wechat_trades"."notify_state" IS '通知状态';
COMMENT ON TABLE "payment"."wechat_trades" IS '微信交易';
-- +migrate Down
DROP TABLE IF EXISTS "payment"."wechat_trades";