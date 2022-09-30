# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/message/message.proto](#proto/message/message.proto)
    - [Empty](#message.Empty)
    - [GetBranchTemplateRequest](#message.GetBranchTemplateRequest)
    - [GetBranchTemplateResponse](#message.GetBranchTemplateResponse)
    - [GetBranchTemplateResponse.GetBranchTemplateResponseDate](#message.GetBranchTemplateResponse.GetBranchTemplateResponseDate)
    - [GetBranchTemplateResponse.GetBranchTemplateResponseDate.Cc](#message.GetBranchTemplateResponse.GetBranchTemplateResponseDate.Cc)
    - [GetShortURLRequest](#message.GetShortURLRequest)
    - [GetShortURLResponse](#message.GetShortURLResponse)
    - [GetShortURLResponse.Data](#message.GetShortURLResponse.Data)
    - [MessageVariableResponse](#message.MessageVariableResponse)
    - [MessageVariableResponse.Variable](#message.MessageVariableResponse.Variable)
    - [MessageVariableResponse.Variable.Message](#message.MessageVariableResponse.Variable.Message)
  
    - [MessageService](#message.MessageService)
  
- [proto/message/setting.proto](#proto/message/setting.proto)
    - [Cc](#message.Cc)
    - [ChangeMessageSettingStatusRequest](#message.ChangeMessageSettingStatusRequest)
    - [ChangeMessageSettingStatusResponse](#message.ChangeMessageSettingStatusResponse)
    - [CreateMessageSettingRequest](#message.CreateMessageSettingRequest)
    - [CreateMessageSettingResponse](#message.CreateMessageSettingResponse)
    - [ListMessageSettingData](#message.ListMessageSettingData)
    - [ListMessageSettingRequest](#message.ListMessageSettingRequest)
    - [ListMessageSettingResponse](#message.ListMessageSettingResponse)
    - [MessageSetting](#message.MessageSetting)
    - [ShowMessageSettingRequest](#message.ShowMessageSettingRequest)
    - [ShowMessageSettingResponse](#message.ShowMessageSettingResponse)
    - [SpecialSetting](#message.SpecialSetting)
    - [UpdateMessageSettingRequest](#message.UpdateMessageSettingRequest)
    - [UpdateMessageSettingResponse](#message.UpdateMessageSettingResponse)
  
- [proto/message/sms.proto](#proto/message/sms.proto)
    - [ChangeSmsTemplateStatusRequest](#message.ChangeSmsTemplateStatusRequest)
    - [ChangeSmsTemplateStatusResponse](#message.ChangeSmsTemplateStatusResponse)
    - [CreateSmsTemplateRequest](#message.CreateSmsTemplateRequest)
    - [CreateSmsTemplateResponse](#message.CreateSmsTemplateResponse)
    - [ListSmsTemplateData](#message.ListSmsTemplateData)
    - [ListSmsTemplateRequest](#message.ListSmsTemplateRequest)
    - [ListSmsTemplateResponse](#message.ListSmsTemplateResponse)
    - [SendSmsRequest](#message.SendSmsRequest)
    - [SendSmsResponse](#message.SendSmsResponse)
    - [ShowSmsTemplateRequest](#message.ShowSmsTemplateRequest)
    - [ShowSmsTemplateResponse](#message.ShowSmsTemplateResponse)
    - [SmsTemplate](#message.SmsTemplate)
    - [UpdateSmsTemplateRequest](#message.UpdateSmsTemplateRequest)
    - [UpdateSmsTemplateResponse](#message.UpdateSmsTemplateResponse)
  
- [proto/message/sms_stat.proto](#proto/message/sms_stat.proto)
    - [DetailSmsHistory](#message.DetailSmsHistory)
    - [ListSmsHistoryData](#message.ListSmsHistoryData)
    - [ListSmsHistoryRequest](#message.ListSmsHistoryRequest)
    - [ListSmsHistoryResponse](#message.ListSmsHistoryResponse)
    - [StatSmsHistory](#message.StatSmsHistory)
    - [StatSmsHistoryData](#message.StatSmsHistoryData)
    - [StatSmsHistoryRequest](#message.StatSmsHistoryRequest)
    - [StatSmsHistoryResponse](#message.StatSmsHistoryResponse)
  
- [proto/message/wechat.proto](#proto/message/wechat.proto)
    - [CreateWechatTemplateRequest](#message.CreateWechatTemplateRequest)
    - [CreateWechatTemplateResponse](#message.CreateWechatTemplateResponse)
    - [ListMiniprogramLinkResponse](#message.ListMiniprogramLinkResponse)
    - [ListMiniprogramLinkResponse.MiniprogramLink](#message.ListMiniprogramLinkResponse.MiniprogramLink)
    - [ListOfficialLinkResponse](#message.ListOfficialLinkResponse)
    - [ListOfficialLinkResponse.OfficialLink](#message.ListOfficialLinkResponse.OfficialLink)
    - [ListWechatTemplateData](#message.ListWechatTemplateData)
    - [ListWechatTemplateRequest](#message.ListWechatTemplateRequest)
    - [ListWechatTemplateResponse](#message.ListWechatTemplateResponse)
    - [SendWechatTemplateRequest](#message.SendWechatTemplateRequest)
    - [SendWechatTemplateRequest.Miniprogram](#message.SendWechatTemplateRequest.Miniprogram)
    - [SendWechatTemplateRequest.WechatUser](#message.SendWechatTemplateRequest.WechatUser)
    - [SendWechatTemplateResponse](#message.SendWechatTemplateResponse)
    - [ShowWechatTemplateRequest](#message.ShowWechatTemplateRequest)
    - [ShowWechatTemplateResponse](#message.ShowWechatTemplateResponse)
    - [UpdateWechatTemplateRequest](#message.UpdateWechatTemplateRequest)
    - [UpdateWechatTemplateResponse](#message.UpdateWechatTemplateResponse)
    - [UpdateWechatTemplateStatusRequest](#message.UpdateWechatTemplateStatusRequest)
    - [UpdateWechatTemplateStatusResponse](#message.UpdateWechatTemplateStatusResponse)
    - [WechatTemplate](#message.WechatTemplate)
    - [WechatTemplateContent](#message.WechatTemplateContent)
    - [WechatTemplateContentBase](#message.WechatTemplateContentBase)
    - [WechatTemplateContentDetail](#message.WechatTemplateContentDetail)
  
- [proto/message/wechat_stat.proto](#proto/message/wechat_stat.proto)
    - [WechatStat](#message.WechatStat)
    - [WechatStatData](#message.WechatStatData)
    - [WechatStatRequest](#message.WechatStatRequest)
    - [WechatStatResponse](#message.WechatStatResponse)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto/message/message.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/message/message.proto



<a name="message.Empty"></a>

### Empty







<a name="message.GetBranchTemplateRequest"></a>

### GetBranchTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| message_type | [string](#string) |  |  |
| time | [int64](#int64) |  |  |






<a name="message.GetBranchTemplateResponse"></a>

### GetBranchTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GetBranchTemplateResponse.GetBranchTemplateResponseDate](#message.GetBranchTemplateResponse.GetBranchTemplateResponseDate) |  |  |






<a name="message.GetBranchTemplateResponse.GetBranchTemplateResponseDate"></a>

### GetBranchTemplateResponse.GetBranchTemplateResponseDate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| trigger_type | [string](#string) |  |  |
| advance_hour | [int32](#int32) |  |  |
| sms_template | [SmsTemplate](#message.SmsTemplate) |  |  |
| wechat_template | [WechatTemplate](#message.WechatTemplate) |  |  |
| cc | [GetBranchTemplateResponse.GetBranchTemplateResponseDate.Cc](#message.GetBranchTemplateResponse.GetBranchTemplateResponseDate.Cc) | repeated |  |






<a name="message.GetBranchTemplateResponse.GetBranchTemplateResponseDate.Cc"></a>

### GetBranchTemplateResponse.GetBranchTemplateResponseDate.Cc



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  |  |
| phone | [string](#string) |  |  |






<a name="message.GetShortURLRequest"></a>

### GetShortURLRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |






<a name="message.GetShortURLResponse"></a>

### GetShortURLResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GetShortURLResponse.Data](#message.GetShortURLResponse.Data) |  |  |






<a name="message.GetShortURLResponse.Data"></a>

### GetShortURLResponse.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |






<a name="message.MessageVariableResponse"></a>

### MessageVariableResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [MessageVariableResponse.Variable](#message.MessageVariableResponse.Variable) | repeated |  |






<a name="message.MessageVariableResponse.Variable"></a>

### MessageVariableResponse.Variable



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| system | [string](#string) |  |  |
| system_key | [string](#string) |  |  |
| message | [MessageVariableResponse.Variable.Message](#message.MessageVariableResponse.Variable.Message) | repeated |  |






<a name="message.MessageVariableResponse.Variable.Message"></a>

### MessageVariableResponse.Variable.Message



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [string](#string) |  |  |
| category_key | [string](#string) |  |  |
| variables | [string](#string) | repeated |  |
| trigger | [string](#string) |  |  |
| setting_disable | [bool](#bool) |  |  |





 

 

 


<a name="message.MessageService"></a>

### MessageService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SendSms | [SendSmsRequest](#message.SendSmsRequest) | [SendSmsResponse](#message.SendSmsResponse) |  |
| GetMessageVariable | [Empty](#message.Empty) | [MessageVariableResponse](#message.MessageVariableResponse) |  |
| CreateSmsTemplate | [CreateSmsTemplateRequest](#message.CreateSmsTemplateRequest) | [CreateSmsTemplateResponse](#message.CreateSmsTemplateResponse) | 短信模版 |
| ListSmsTemplate | [ListSmsTemplateRequest](#message.ListSmsTemplateRequest) | [ListSmsTemplateResponse](#message.ListSmsTemplateResponse) |  |
| UpdateSmsTemplate | [UpdateSmsTemplateRequest](#message.UpdateSmsTemplateRequest) | [UpdateSmsTemplateResponse](#message.UpdateSmsTemplateResponse) |  |
| ShowSmsTemplate | [ShowSmsTemplateRequest](#message.ShowSmsTemplateRequest) | [ShowSmsTemplateResponse](#message.ShowSmsTemplateResponse) |  |
| ChangeSmsTemplateStatus | [ChangeSmsTemplateStatusRequest](#message.ChangeSmsTemplateStatusRequest) | [ChangeSmsTemplateStatusResponse](#message.ChangeSmsTemplateStatusResponse) |  |
| CreateMessageSetting | [CreateMessageSettingRequest](#message.CreateMessageSettingRequest) | [CreateMessageSettingResponse](#message.CreateMessageSettingResponse) | 消息设置 |
| ListMessageSetting | [ListMessageSettingRequest](#message.ListMessageSettingRequest) | [ListMessageSettingResponse](#message.ListMessageSettingResponse) |  |
| UpdateMessageSetting | [UpdateMessageSettingRequest](#message.UpdateMessageSettingRequest) | [UpdateMessageSettingResponse](#message.UpdateMessageSettingResponse) |  |
| ShowMessageSetting | [ShowMessageSettingRequest](#message.ShowMessageSettingRequest) | [ShowMessageSettingResponse](#message.ShowMessageSettingResponse) |  |
| ChangeMessageSettingStatus | [ChangeMessageSettingStatusRequest](#message.ChangeMessageSettingStatusRequest) | [ChangeMessageSettingStatusResponse](#message.ChangeMessageSettingStatusResponse) |  |
| StatSmsHistory | [StatSmsHistoryRequest](#message.StatSmsHistoryRequest) | [StatSmsHistoryResponse](#message.StatSmsHistoryResponse) | 发送记录 |
| ListSmsHistory | [ListSmsHistoryRequest](#message.ListSmsHistoryRequest) | [ListSmsHistoryResponse](#message.ListSmsHistoryResponse) |  |
| GetBranchTemplate | [GetBranchTemplateRequest](#message.GetBranchTemplateRequest) | [GetBranchTemplateResponse](#message.GetBranchTemplateResponse) | 获取门店短信模版 |
| GetShortURL | [GetShortURLRequest](#message.GetShortURLRequest) | [GetShortURLResponse](#message.GetShortURLResponse) | 短链 |
| ListOfficialLink | [Empty](#message.Empty) | [ListOfficialLinkResponse](#message.ListOfficialLinkResponse) | 微信模版 |
| ListMiniprogramLink | [Empty](#message.Empty) | [ListMiniprogramLinkResponse](#message.ListMiniprogramLinkResponse) |  |
| CreateWechatTemplate | [CreateWechatTemplateRequest](#message.CreateWechatTemplateRequest) | [CreateWechatTemplateResponse](#message.CreateWechatTemplateResponse) |  |
| ListWechatTemplate | [ListWechatTemplateRequest](#message.ListWechatTemplateRequest) | [ListWechatTemplateResponse](#message.ListWechatTemplateResponse) |  |
| UpdateWechatTemplateStatus | [UpdateWechatTemplateStatusRequest](#message.UpdateWechatTemplateStatusRequest) | [UpdateWechatTemplateStatusResponse](#message.UpdateWechatTemplateStatusResponse) |  |
| UpdateWechatTemplate | [UpdateWechatTemplateRequest](#message.UpdateWechatTemplateRequest) | [UpdateWechatTemplateResponse](#message.UpdateWechatTemplateResponse) |  |
| ShowWechatTemplate | [ShowWechatTemplateRequest](#message.ShowWechatTemplateRequest) | [ShowWechatTemplateResponse](#message.ShowWechatTemplateResponse) |  |
| SendWechatTemplate | [SendWechatTemplateRequest](#message.SendWechatTemplateRequest) | [SendWechatTemplateResponse](#message.SendWechatTemplateResponse) | 发送微信模版消息 |
| WechatStat | [WechatStatRequest](#message.WechatStatRequest) | [WechatStatResponse](#message.WechatStatResponse) |  |

 



<a name="proto/message/setting.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/message/setting.proto



<a name="message.Cc"></a>

### Cc



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  |  |
| phone | [string](#string) |  |  |






<a name="message.ChangeMessageSettingStatusRequest"></a>

### ChangeMessageSettingStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="message.ChangeMessageSettingStatusResponse"></a>

### ChangeMessageSettingStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [MessageSetting](#message.MessageSetting) |  |  |






<a name="message.CreateMessageSettingRequest"></a>

### CreateMessageSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message_type | [string](#string) |  |  |
| message_type_name | [string](#string) |  |  |
| trigger_type | [string](#string) |  |  |
| advance_hour | [double](#double) |  |  |
| sms_template_id | [string](#string) |  |  |
| sms_template_name | [string](#string) |  |  |
| wechat_template_id | [string](#string) |  |  |
| wechat_template_name | [string](#string) |  |  |
| special_setting | [SpecialSetting](#message.SpecialSetting) | repeated |  |
| cc_list | [Cc](#message.Cc) | repeated |  |
| special_branches | [string](#string) | repeated |  |
| status | [string](#string) |  |  |






<a name="message.CreateMessageSettingResponse"></a>

### CreateMessageSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="message.ListMessageSettingData"></a>

### ListMessageSettingData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| settings | [MessageSetting](#message.MessageSetting) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="message.ListMessageSettingRequest"></a>

### ListMessageSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message_type | [string](#string) |  |  |
| trigger_type | [string](#string) |  |  |
| status | [string](#string) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |
| order_by | [string](#string) |  |  |
| with_page | [bool](#bool) |  |  |






<a name="message.ListMessageSettingResponse"></a>

### ListMessageSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ListMessageSettingData](#message.ListMessageSettingData) |  |  |






<a name="message.MessageSetting"></a>

### MessageSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| message_type | [string](#string) |  |  |
| message_type_name | [string](#string) |  |  |
| trigger_type | [string](#string) |  |  |
| advance_hour | [double](#double) |  |  |
| sms_template_id | [string](#string) |  |  |
| sms_template_name | [string](#string) |  |  |
| wechat_template_id | [string](#string) |  |  |
| wechat_template_name | [string](#string) |  |  |
| special_setting | [SpecialSetting](#message.SpecialSetting) | repeated |  |
| cc_list | [Cc](#message.Cc) | repeated |  |
| special_branches | [string](#string) | repeated |  |
| status | [string](#string) |  |  |
| created_at | [int64](#int64) |  |  |






<a name="message.ShowMessageSettingRequest"></a>

### ShowMessageSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="message.ShowMessageSettingResponse"></a>

### ShowMessageSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [MessageSetting](#message.MessageSetting) |  |  |






<a name="message.SpecialSetting"></a>

### SpecialSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| begin | [string](#string) |  |  |
| end | [string](#string) |  |  |
| sms_template_id | [string](#string) |  |  |
| sms_template_name | [string](#string) |  |  |
| wechat_template_id | [string](#string) |  |  |
| wechat_template_name | [string](#string) |  |  |






<a name="message.UpdateMessageSettingRequest"></a>

### UpdateMessageSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| message_type | [string](#string) |  |  |
| message_type_name | [string](#string) |  |  |
| trigger_type | [string](#string) |  |  |
| advance_hour | [double](#double) |  |  |
| sms_template_id | [string](#string) |  |  |
| sms_template_name | [string](#string) |  |  |
| wechat_template_id | [string](#string) |  |  |
| wechat_template_name | [string](#string) |  |  |
| special_setting | [SpecialSetting](#message.SpecialSetting) | repeated |  |
| cc_list | [Cc](#message.Cc) | repeated |  |
| special_branches | [string](#string) | repeated |  |
| status | [string](#string) |  |  |






<a name="message.UpdateMessageSettingResponse"></a>

### UpdateMessageSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |





 

 

 

 



<a name="proto/message/sms.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/message/sms.proto



<a name="message.ChangeSmsTemplateStatusRequest"></a>

### ChangeSmsTemplateStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="message.ChangeSmsTemplateStatusResponse"></a>

### ChangeSmsTemplateStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [SmsTemplate](#message.SmsTemplate) |  |  |






<a name="message.CreateSmsTemplateRequest"></a>

### CreateSmsTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| sign | [string](#string) |  |  |
| category | [string](#string) |  |  |
| category_key | [string](#string) |  |  |
| content | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="message.CreateSmsTemplateResponse"></a>

### CreateSmsTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="message.ListSmsTemplateData"></a>

### ListSmsTemplateData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| templates | [SmsTemplate](#message.SmsTemplate) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="message.ListSmsTemplateRequest"></a>

### ListSmsTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| category | [string](#string) |  |  |
| category_key | [string](#string) |  |  |
| status | [string](#string) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |
| order_by | [string](#string) |  |  |
| with_page | [bool](#bool) |  |  |






<a name="message.ListSmsTemplateResponse"></a>

### ListSmsTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ListSmsTemplateData](#message.ListSmsTemplateData) |  |  |






<a name="message.SendSmsRequest"></a>

### SendSmsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sign | [string](#string) |  |  |
| area_code | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| message | [string](#string) |  |  |
| system | [string](#string) |  |  |
| message_type | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |






<a name="message.SendSmsResponse"></a>

### SendSmsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="message.ShowSmsTemplateRequest"></a>

### ShowSmsTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="message.ShowSmsTemplateResponse"></a>

### ShowSmsTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [SmsTemplate](#message.SmsTemplate) |  |  |






<a name="message.SmsTemplate"></a>

### SmsTemplate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| sign | [string](#string) |  |  |
| category | [string](#string) |  |  |
| category_key | [string](#string) |  |  |
| content | [string](#string) |  |  |
| status | [string](#string) |  |  |
| created_at | [int64](#int64) |  |  |






<a name="message.UpdateSmsTemplateRequest"></a>

### UpdateSmsTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| sign | [string](#string) |  |  |
| content | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="message.UpdateSmsTemplateResponse"></a>

### UpdateSmsTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |





 

 

 

 



<a name="proto/message/sms_stat.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/message/sms_stat.proto



<a name="message.DetailSmsHistory"></a>

### DetailSmsHistory



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| area_code | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| sign | [string](#string) |  |  |
| content | [string](#string) |  |  |
| status | [string](#string) |  |  |
| created_at | [int64](#int64) |  |  |






<a name="message.ListSmsHistoryData"></a>

### ListSmsHistoryData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| histories | [DetailSmsHistory](#message.DetailSmsHistory) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="message.ListSmsHistoryRequest"></a>

### ListSmsHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [string](#string) | repeated |  |
| branch_id | [string](#string) |  |  |
| message_type | [string](#string) |  |  |
| begin_date | [string](#string) |  |  |
| end_date | [string](#string) |  |  |
| sms_status | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |
| order_by | [string](#string) |  |  |
| with_page | [bool](#bool) |  |  |






<a name="message.ListSmsHistoryResponse"></a>

### ListSmsHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ListSmsHistoryData](#message.ListSmsHistoryData) |  |  |






<a name="message.StatSmsHistory"></a>

### StatSmsHistory



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| total | [int64](#int64) |  |  |
| succeed | [int64](#int64) |  |  |
| failured | [int64](#int64) |  |  |






<a name="message.StatSmsHistoryData"></a>

### StatSmsHistoryData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stats | [StatSmsHistory](#message.StatSmsHistory) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="message.StatSmsHistoryRequest"></a>

### StatSmsHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| message_type | [string](#string) |  |  |
| begin_date | [string](#string) |  |  |
| end_date | [string](#string) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |
| order_by | [string](#string) |  |  |
| with_page | [bool](#bool) |  |  |
| branch_ids | [string](#string) | repeated |  |






<a name="message.StatSmsHistoryResponse"></a>

### StatSmsHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [StatSmsHistoryData](#message.StatSmsHistoryData) |  |  |





 

 

 

 



<a name="proto/message/wechat.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/message/wechat.proto



<a name="message.CreateWechatTemplateRequest"></a>

### CreateWechatTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| template_name | [string](#string) |  |  |
| template_code | [string](#string) |  |  |
| category | [string](#string) |  |  |
| category_key | [string](#string) |  |  |
| content | [WechatTemplateContent](#message.WechatTemplateContent) |  |  |
| official_link | [string](#string) |  |  |
| miniprogram_link | [string](#string) |  |  |






<a name="message.CreateWechatTemplateResponse"></a>

### CreateWechatTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="message.ListMiniprogramLinkResponse"></a>

### ListMiniprogramLinkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ListMiniprogramLinkResponse.MiniprogramLink](#message.ListMiniprogramLinkResponse.MiniprogramLink) | repeated |  |






<a name="message.ListMiniprogramLinkResponse.MiniprogramLink"></a>

### ListMiniprogramLinkResponse.MiniprogramLink



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| name | [string](#string) |  |  |






<a name="message.ListOfficialLinkResponse"></a>

### ListOfficialLinkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ListOfficialLinkResponse.OfficialLink](#message.ListOfficialLinkResponse.OfficialLink) | repeated |  |






<a name="message.ListOfficialLinkResponse.OfficialLink"></a>

### ListOfficialLinkResponse.OfficialLink



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| name | [string](#string) |  |  |






<a name="message.ListWechatTemplateData"></a>

### ListWechatTemplateData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| templates | [WechatTemplate](#message.WechatTemplate) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="message.ListWechatTemplateRequest"></a>

### ListWechatTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| template_name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |
| order_by | [string](#string) |  |  |
| with_page | [bool](#bool) |  |  |






<a name="message.ListWechatTemplateResponse"></a>

### ListWechatTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ListWechatTemplateData](#message.ListWechatTemplateData) |  |  |






<a name="message.SendWechatTemplateRequest"></a>

### SendWechatTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| channel | [string](#string) |  |  |
| system | [string](#string) |  |  |
| template_id | [string](#string) |  |  |
| content | [WechatTemplateContent](#message.WechatTemplateContent) |  |  |
| branch_id | [string](#string) |  |  |
| message_type | [string](#string) |  |  |
| wechat_user | [SendWechatTemplateRequest.WechatUser](#message.SendWechatTemplateRequest.WechatUser) |  |  |
| official_link | [string](#string) |  |  |
| miniprogram | [SendWechatTemplateRequest.Miniprogram](#message.SendWechatTemplateRequest.Miniprogram) |  |  |






<a name="message.SendWechatTemplateRequest.Miniprogram"></a>

### SendWechatTemplateRequest.Miniprogram



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| appid | [string](#string) |  |  |
| pagepath | [string](#string) |  |  |






<a name="message.SendWechatTemplateRequest.WechatUser"></a>

### SendWechatTemplateRequest.WechatUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member_id | [string](#string) |  |  |
| member_wechat_id | [string](#string) |  |  |
| member_open_id | [string](#string) |  |  |






<a name="message.SendWechatTemplateResponse"></a>

### SendWechatTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="message.ShowWechatTemplateRequest"></a>

### ShowWechatTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="message.ShowWechatTemplateResponse"></a>

### ShowWechatTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [WechatTemplate](#message.WechatTemplate) |  |  |






<a name="message.UpdateWechatTemplateRequest"></a>

### UpdateWechatTemplateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| template_name | [string](#string) |  |  |
| template_code | [string](#string) |  |  |
| content | [WechatTemplateContent](#message.WechatTemplateContent) |  |  |
| official_link | [string](#string) |  |  |
| miniprogram_link | [string](#string) |  |  |






<a name="message.UpdateWechatTemplateResponse"></a>

### UpdateWechatTemplateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="message.UpdateWechatTemplateStatusRequest"></a>

### UpdateWechatTemplateStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="message.UpdateWechatTemplateStatusResponse"></a>

### UpdateWechatTemplateStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="message.WechatTemplate"></a>

### WechatTemplate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| template_name | [string](#string) |  |  |
| template_code | [string](#string) |  |  |
| content | [WechatTemplateContent](#message.WechatTemplateContent) |  |  |
| official_link | [string](#string) |  |  |
| official_link_name | [string](#string) |  |  |
| miniprogram_link | [string](#string) |  |  |
| miniprogram_link_name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| category | [string](#string) |  |  |
| category_key | [string](#string) |  |  |
| created_at | [int64](#int64) |  |  |






<a name="message.WechatTemplateContent"></a>

### WechatTemplateContent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| first | [WechatTemplateContentBase](#message.WechatTemplateContentBase) |  |  |
| detail | [WechatTemplateContentDetail](#message.WechatTemplateContentDetail) | repeated |  |
| remark | [WechatTemplateContentBase](#message.WechatTemplateContentBase) |  |  |






<a name="message.WechatTemplateContentBase"></a>

### WechatTemplateContentBase



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |
| color | [string](#string) |  |  |






<a name="message.WechatTemplateContentDetail"></a>

### WechatTemplateContentDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| value | [string](#string) |  |  |
| color | [string](#string) |  |  |





 

 

 

 



<a name="proto/message/wechat_stat.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/message/wechat_stat.proto



<a name="message.WechatStat"></a>

### WechatStat



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| total | [int64](#int64) |  |  |
| success_total | [int64](#int64) |  |  |
| fail_total | [int64](#int64) |  |  |






<a name="message.WechatStatData"></a>

### WechatStatData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stats | [WechatStat](#message.WechatStat) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="message.WechatStatRequest"></a>

### WechatStatRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| message_type | [string](#string) |  |  |
| begin_date | [int64](#int64) |  |  |
| end_date | [int64](#int64) |  |  |
| with_page | [bool](#bool) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






<a name="message.WechatStatResponse"></a>

### WechatStatResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [WechatStatData](#message.WechatStatData) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

