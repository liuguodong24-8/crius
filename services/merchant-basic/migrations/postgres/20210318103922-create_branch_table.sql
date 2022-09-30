
-- +migrate Up
CREATE TABLE IF NOT EXISTS "merchant_basic"."branch" (
  "id" uuid primary key,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "alias" varchar(50) COLLATE "pg_catalog"."default",
  "simplify" varchar(50) COLLATE "pg_catalog"."default",
  "contact_phone" varchar(20) COLLATE "pg_catalog"."default",
  "code" varchar(20) COLLATE "pg_catalog"."default",
  "province_id" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "city_id" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "district_id" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "area_id" uuid,
  "address" varchar(1024) COLLATE "pg_catalog"."default",
  "merchant_id" uuid,
  "brand_id" uuid,
  "sub_mch_id" varchar(20),
  "domain" varchar(300) COLLATE "pg_catalog"."default",
  "latitude" float4,
  "longitude" float4,
  "status" varchar(10) COLLATE "pg_catalog"."default" DEFAULT 'opened'::character varying,
  "opened_at" date,
  "photo" varchar[] COLLATE "pg_catalog"."default",
  "parking" varchar(300) COLLATE "pg_catalog"."default",
  "weight" int8,
  "biz_type" int2,
  "business_status" varchar(30) COLLATE "pg_catalog"."default",
  "location" varchar(450) COLLATE "pg_catalog"."default",
  "authorization_salt" varchar(255) COLLATE "pg_catalog"."default",
  "signature_salt" varchar(255) COLLATE "pg_catalog"."default",
  "extra" json,
  "load_extra" jsonb,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6)
)
;
COMMENT ON COLUMN "merchant_basic"."branch"."id" IS '门店ID';
COMMENT ON COLUMN "merchant_basic"."branch"."name" IS '门店名';
COMMENT ON COLUMN "merchant_basic"."branch"."province_id" IS '省份';
COMMENT ON COLUMN "merchant_basic"."branch"."city_id" IS '城市';
COMMENT ON COLUMN "merchant_basic"."branch"."district_id" IS '区域';
COMMENT ON COLUMN "merchant_basic"."branch"."address" IS '详细地址(手动输入地址)';
COMMENT ON COLUMN "merchant_basic"."branch"."contact_phone" IS '联系电话';
COMMENT ON COLUMN "merchant_basic"."branch"."extra" IS '冗余信息';
COMMENT ON COLUMN "merchant_basic"."branch"."code" IS '编号';
COMMENT ON COLUMN "merchant_basic"."branch"."deleted_at" IS '删除时间';
COMMENT ON COLUMN "merchant_basic"."branch"."latitude" IS '纬度';
COMMENT ON COLUMN "merchant_basic"."branch"."longitude" IS '经度';
COMMENT ON COLUMN "merchant_basic"."branch"."status" IS '状态 opened closed';
COMMENT ON COLUMN "merchant_basic"."branch"."merchant_id" IS '商户ID';
COMMENT ON COLUMN "merchant_basic"."branch"."opened_at" IS '开通时间';
COMMENT ON COLUMN "merchant_basic"."branch"."photo" IS '照片';
COMMENT ON COLUMN "merchant_basic"."branch"."parking" IS '停车信息';
COMMENT ON COLUMN "merchant_basic"."branch"."area_id" IS '地区id(西南区 华东区)';
COMMENT ON COLUMN "merchant_basic"."branch"."weight" IS '权值';
COMMENT ON COLUMN "merchant_basic"."branch"."domain" IS '域名';
COMMENT ON COLUMN "merchant_basic"."branch"."biz_type" IS '类型 0直营 1加盟';
COMMENT ON COLUMN "merchant_basic"."branch"."business_status" IS '营业状态 已关闭closed 不开放not_opened 试营业soft_opening 营业opening 暂停营业suspend_business';
COMMENT ON COLUMN "merchant_basic"."branch"."alias" IS '简称';
COMMENT ON COLUMN "merchant_basic"."branch"."simplify" IS '拼音';
COMMENT ON COLUMN "merchant_basic"."branch"."location" IS '导航地址';
COMMENT ON COLUMN "merchant_basic"."branch"."sub_mch_id" IS '微信支付子商户号';
COMMENT ON COLUMN "merchant_basic"."branch"."brand_id" IS '品牌ID';
COMMENT ON TABLE "merchant_basic"."branch" IS '门店表';

create index idx_branch_created_at on "merchant_basic"."branch"(created_at);
create index idx_branch_province on "merchant_basic"."branch"(province_id);
create index idx_branch_city on "merchant_basic"."branch"(city_id);
create index idx_branch_district on "merchant_basic"."branch"(district_id);

-- +migrate Down
DROP TABLE IF EXISTS "merchant_basic"."branch";
