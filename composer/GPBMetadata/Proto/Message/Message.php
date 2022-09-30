<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/message/message.proto

namespace GPBMetadata\Proto\Message;

class Message
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Proto\Message\Sms::initOnce();
        \GPBMetadata\Proto\Message\Setting::initOnce();
        \GPBMetadata\Proto\Message\SmsStat::initOnce();
        \GPBMetadata\Proto\Message\Wechat::initOnce();
        \GPBMetadata\Proto\Message\WechatStat::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
proto/message/message.protomessageproto/message/setting.protoproto/message/sms_stat.protoproto/message/wechat.protoproto/message/wechat_stat.proto"
Empty"�
MessageVariableResponse

error_code (
error_message (	7
data (2).message.MessageVariableResponse.Variable�
Variable
system (	

system_key (	B
message (21.message.MessageVariableResponse.Variable.Messagen
Message
category (	
category_key (	
	variables (	
trigger (	
setting_disable ("Q
GetBranchTemplateRequest
	branch_id (	
message_type (	
time ("�
GetBranchTemplateResponse

error_code (
error_message (	N
data (2@.message.GetBranchTemplateResponse.GetBranchTemplateResponseDate�
GetBranchTemplateResponseDate

id (	
trigger_type (	
advance_hour (*
sms_template (2.message.SmsTemplate0
wechat_template (2.message.WechatTemplateO
cc (2C.message.GetBranchTemplateResponse.GetBranchTemplateResponseDate.Cc!
Cc
code (	
phone (	"!
GetShortURLRequest
url (	"�
GetShortURLResponse

error_code (
error_message (	/
data (2!.message.GetShortURLResponse.Data
Data
url (	2�
MessageService>
SendSms.message.SendSmsRequest.message.SendSmsResponse" H
GetMessageVariable.message.Empty .message.MessageVariableResponse" \\
CreateSmsTemplate!.message.CreateSmsTemplateRequest".message.CreateSmsTemplateResponse" V
ListSmsTemplate.message.ListSmsTemplateRequest .message.ListSmsTemplateResponse" \\
UpdateSmsTemplate!.message.UpdateSmsTemplateRequest".message.UpdateSmsTemplateResponse" V
ShowSmsTemplate.message.ShowSmsTemplateRequest .message.ShowSmsTemplateResponse" n
ChangeSmsTemplateStatus\'.message.ChangeSmsTemplateStatusRequest(.message.ChangeSmsTemplateStatusResponse" e
CreateMessageSetting$.message.CreateMessageSettingRequest%.message.CreateMessageSettingResponse" _
ListMessageSetting".message.ListMessageSettingRequest#.message.ListMessageSettingResponse" e
UpdateMessageSetting$.message.UpdateMessageSettingRequest%.message.UpdateMessageSettingResponse" _
ShowMessageSetting".message.ShowMessageSettingRequest#.message.ShowMessageSettingResponse" w
ChangeMessageSettingStatus*.message.ChangeMessageSettingStatusRequest+.message.ChangeMessageSettingStatusResponse" S
StatSmsHistory.message.StatSmsHistoryRequest.message.StatSmsHistoryResponse" S
ListSmsHistory.message.ListSmsHistoryRequest.message.ListSmsHistoryResponse" \\
GetBranchTemplate!.message.GetBranchTemplateRequest".message.GetBranchTemplateResponse" J
GetShortURL.message.GetShortURLRequest.message.GetShortURLResponse" G
ListOfficialLink.message.Empty!.message.ListOfficialLinkResponse" M
ListMiniprogramLink.message.Empty$.message.ListMiniprogramLinkResponse" e
CreateWechatTemplate$.message.CreateWechatTemplateRequest%.message.CreateWechatTemplateResponse" _
ListWechatTemplate".message.ListWechatTemplateRequest#.message.ListWechatTemplateResponse" w
UpdateWechatTemplateStatus*.message.UpdateWechatTemplateStatusRequest+.message.UpdateWechatTemplateStatusResponse" e
UpdateWechatTemplate$.message.UpdateWechatTemplateRequest%.message.UpdateWechatTemplateResponse" _
ShowWechatTemplate".message.ShowWechatTemplateRequest#.message.ShowWechatTemplateResponse" _
SendWechatTemplate".message.SendWechatTemplateRequest#.message.SendWechatTemplateResponse" G

WechatStat.message.WechatStatRequest.message.WechatStatResponse" B)Z./proto;proto�Omy\\Crius\\MessageServerbproto3'
        , true);

        static::$is_initialized = true;
    }
}

