<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/growth_config.proto

namespace GPBMetadata\Proto\MerchantBasic;

class GrowthConfig
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
(proto/merchant-basic/growth_config.proto
GrowthConfig
name (	
top (
rules (	"p
ShowGrowthConfigResponse

error_code (

data (2.merchantBasic.GrowthConfig"F
SaveGrowthConfigRequest+
config (2.merchantBasic.GrowthConfigB/Z
        , true);

        static::$is_initialized = true;
    }
}
