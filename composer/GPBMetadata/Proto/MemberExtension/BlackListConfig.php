<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/black_list_config.proto

namespace GPBMetadata\Proto\MemberExtension;

class BlackListConfig
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
.proto/member-extension/black_list_config.protomemberExtension"O
SaveBlackListConfigRequest1
configs (2 .memberExtension.BlackListConfig"x
GetBlackListConfigsResponse

error_code (

data (2 .memberExtension.BlackListConfig"�
BlackListConfig

id (	
name (	
in_days (
in_times (
out_days (
	out_times (

block_call (
level (B+Z
        , true);

        static::$is_initialized = true;
    }
}
