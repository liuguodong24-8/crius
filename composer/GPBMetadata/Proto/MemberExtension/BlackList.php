<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/black_list.proto

namespace GPBMetadata\Proto\MemberExtension;

class BlackList
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
?
\'proto/member-extension/black_list.protomemberExtension"B
CreateBlackListRequest(
list (2.memberExtension.BlackList"$
DeleteBlackListRequest

id (	"h
GetBlackListsRequest
phone_suffix (	
phone (	
name (	
offset (
limit ("?
GetBlackListsResponse

error_code (
error_message (	9
data (2+.memberExtension.GetBlackListsResponse.Data?
Data
total ((
list (2.memberExtension.BlackList"?
	BlackList

id (	

phone_code (	
phone (	
name (	
black_list_config_id (	
black_list_level (
join_at (
joint_reason (	
tag	 (	
remark
 (	
staff_id (	
black_list_config_name (	

block_call (",
ShowBlackListByPhoneRequest
phone (	"s
ShowBlackListByPhoneResponse

error_code (
error_message (	(
data (2.memberExtension.BlackListB+Z./proto;proto?Omy\\Crius\\ExtensionServerbproto3'
        , true);

        static::$is_initialized = true;
    }
}

