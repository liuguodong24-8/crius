<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/brand.proto

namespace GPBMetadata\Proto\MerchantBasic;

class Brand
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
?
 proto/merchant-basic/brand.protomerchantBasic"A
CreateBrandRequest
name (	
order (
status (	"M
UpdateBrandRequest

id (	
name (	
order (
status (	"6
UpdateBrandStatusRequest

id (	
status (	"O
GetBrandsRequest
name (	
offset (
limit (
status (	"j
GetBrandsResponse

error_code (
error_message (	*
data (2.merchantBasic.GetBrandsData"D
GetBrandsData
total ($
brands (2.merchantBasic.Brand"T
Brand

id (	
name (	
order (
status (	

created_at (B/Z./proto;proto?Omy\\Crius\\MerchantBasicServerbproto3'
        , true);

        static::$is_initialized = true;
    }
}

