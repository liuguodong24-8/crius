<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/member_address.proto

namespace GPBMetadata\Proto\MerchantBasic;

class MemberAddress
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
)proto/merchant-basic/member_address.protomerchantBasic"?
MemberAddress

id (	
	member_id (	
name (	
phone (	

phone_code (	
province_id (	
city_id (	
district_id (	
address	 (	

is_default
 (",
GetMemberAddressRequest
	member_id (	"q
GetMemberAddressResponse

error_code (
error_message (	*
data (2.merchantBasic.MemberAddress"?
SetMemberAddressDefaultRequest
	member_id (	

id (	"3
GetMemberDefaultAddressRequest
	member_id (	"x
GetMemberDefaultAddressResponse

error_code (
error_message (	*
data (2.merchantBasic.MemberAddress"(
DeleteMemberAddressRequest

id (	B/Z./proto;proto?Omy\\Crius\\MerchantBasicServerbproto3'
        , true);

        static::$is_initialized = true;
    }
}

