<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/message/setting.proto

namespace GPBMetadata\Proto\Message;

class Setting
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
?
proto/message/setting.protomessage"?
SpecialSetting
begin (	
end (	
sms_template_id (	
sms_template_name (	
wechat_template_id (	
wechat_template_name (	"!
Cc
code (	
phone (	"?
ListMessageSettingRequest
message_type (	
trigger_type (	
status (	
limit (
offset (
order_by (	
	with_page ("?
MessageSetting

id (	
message_type (	
message_type_name (	
trigger_type (	
advance_hour (
sms_template_id (	
sms_template_name (	
wechat_template_id (	
wechat_template_name	 (	0
special_setting
 (2.message.SpecialSetting
cc_list (2.message.Cc
special_branches (	
status (	

created_at ("v
ListMessageSettingResponse

error_code (
error_message (	-
data (2.message.ListMessageSettingData"R
ListMessageSettingData)
settings (2.message.MessageSetting
total ("?
CreateMessageSettingRequest
message_type (	
trigger_type (	
advance_hour (
sms_template_id (	
wechat_template_id (	0
special_setting (2.message.SpecialSetting
cc_list (2.message.Cc
special_branches (	
status	 (	"I
CreateMessageSettingResponse

error_code (
error_message (	"?
UpdateMessageSettingRequest

id (	
message_type (	
trigger_type (	
advance_hour (
sms_template_id (	
wechat_template_id (	0
special_setting (2.message.SpecialSetting
cc_list (2.message.Cc
special_branches	 (	
status
 (	"I
UpdateMessageSettingResponse

error_code (
error_message (	"\'
ShowMessageSettingRequest

id (	"n
ShowMessageSettingResponse

error_code (
error_message (	%
data (2.message.MessageSetting"?
!ChangeMessageSettingStatusRequest

id (	
status (	"v
"ChangeMessageSettingStatusResponse

error_code (
error_message (	%
data (2.message.MessageSettingB)Z./proto;proto?Omy\\Crius\\MessageServerbproto3'
        , true);

        static::$is_initialized = true;
    }
}

