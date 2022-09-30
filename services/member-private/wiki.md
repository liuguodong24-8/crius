# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/member-private/member_private.proto](#proto/member-private/member_private.proto)
    - [UpdateStatusRequest](#memberPrivate.UpdateStatusRequest)
    - [UpdateStatusResponse](#memberPrivate.UpdateStatusResponse)
  
    - [MemberPrivateServer](#memberPrivate.MemberPrivateServer)
  
- [proto/member-private/option.proto](#proto/member-private/option.proto)
    - [CreatePromotionOptionRequest](#memberPrivate.CreatePromotionOptionRequest)
    - [ListPromotionOptionRequest](#memberPrivate.ListPromotionOptionRequest)
    - [ListPromotionOptionRequest.FilterPromotion](#memberPrivate.ListPromotionOptionRequest.FilterPromotion)
    - [ListPromotionOptionResponse](#memberPrivate.ListPromotionOptionResponse)
    - [ListPromotionOptionResponse.ListPromotionOptionData](#memberPrivate.ListPromotionOptionResponse.ListPromotionOptionData)
    - [OptionResponse](#memberPrivate.OptionResponse)
    - [PromotionOption](#memberPrivate.PromotionOption)
    - [ShowPromotionOptionRequest](#memberPrivate.ShowPromotionOptionRequest)
    - [ShowPromotionOptionResponse](#memberPrivate.ShowPromotionOptionResponse)
    - [UpdatePromotionOptionRequest](#memberPrivate.UpdatePromotionOptionRequest)
  
- [proto/member-private/promotion.proto](#proto/member-private/promotion.proto)
    - [CreatePromotionRequest](#memberPrivate.CreatePromotionRequest)
    - [CreatePromotionResponse](#memberPrivate.CreatePromotionResponse)
    - [ListPromotionRequest](#memberPrivate.ListPromotionRequest)
    - [ListPromotionResponse](#memberPrivate.ListPromotionResponse)
    - [ListPromotionResponse.ListPromotionData](#memberPrivate.ListPromotionResponse.ListPromotionData)
    - [Promotion](#memberPrivate.Promotion)
    - [PromotionResponse](#memberPrivate.PromotionResponse)
    - [ShowPromotionRequest](#memberPrivate.ShowPromotionRequest)
    - [ShowPromotionResponse](#memberPrivate.ShowPromotionResponse)
    - [UpdatePromotionRequest](#memberPrivate.UpdatePromotionRequest)
  
- [proto/member-private/staff_shift.proto](#proto/member-private/staff_shift.proto)
    - [CreateStaffShiftRequest](#memberPrivate.CreateStaffShiftRequest)
    - [CreateStaffShiftResponse](#memberPrivate.CreateStaffShiftResponse)
    - [CreateStaffShiftResponse.Shift](#memberPrivate.CreateStaffShiftResponse.Shift)
    - [GetStaffShiftTimeRequest](#memberPrivate.GetStaffShiftTimeRequest)
    - [GetStaffShiftTimeResponse](#memberPrivate.GetStaffShiftTimeResponse)
    - [GetStaffShiftTimeResponse.Shift](#memberPrivate.GetStaffShiftTimeResponse.Shift)
    - [ListStaffShiftRequest](#memberPrivate.ListStaffShiftRequest)
    - [ListStaffShiftResponse](#memberPrivate.ListStaffShiftResponse)
    - [ListStaffShiftResponse.StaffShift](#memberPrivate.ListStaffShiftResponse.StaffShift)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto/member-private/member_private.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/member-private/member_private.proto



<a name="memberPrivate.UpdateStatusRequest"></a>

### UpdateStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="memberPrivate.UpdateStatusResponse"></a>

### UpdateStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |





 

 

 


<a name="memberPrivate.MemberPrivateServer"></a>

### MemberPrivateServer


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListPromotion | [ListPromotionRequest](#memberPrivate.ListPromotionRequest) | [ListPromotionResponse](#memberPrivate.ListPromotionResponse) |  |
| CreatePromotion | [CreatePromotionRequest](#memberPrivate.CreatePromotionRequest) | [CreatePromotionResponse](#memberPrivate.CreatePromotionResponse) |  |
| UpdatePromotion | [UpdatePromotionRequest](#memberPrivate.UpdatePromotionRequest) | [PromotionResponse](#memberPrivate.PromotionResponse) |  |
| ShowPromotion | [ShowPromotionRequest](#memberPrivate.ShowPromotionRequest) | [ShowPromotionResponse](#memberPrivate.ShowPromotionResponse) |  |
| UpdatePromotionStatus | [UpdateStatusRequest](#memberPrivate.UpdateStatusRequest) | [UpdateStatusResponse](#memberPrivate.UpdateStatusResponse) |  |
| ListPromotionOption | [ListPromotionOptionRequest](#memberPrivate.ListPromotionOptionRequest) | [ListPromotionOptionResponse](#memberPrivate.ListPromotionOptionResponse) |  |
| CreatePromotionOption | [CreatePromotionOptionRequest](#memberPrivate.CreatePromotionOptionRequest) | [OptionResponse](#memberPrivate.OptionResponse) |  |
| ShowPromotionOption | [ShowPromotionOptionRequest](#memberPrivate.ShowPromotionOptionRequest) | [ShowPromotionOptionResponse](#memberPrivate.ShowPromotionOptionResponse) |  |
| UpdatePromotionOption | [UpdatePromotionOptionRequest](#memberPrivate.UpdatePromotionOptionRequest) | [OptionResponse](#memberPrivate.OptionResponse) |  |
| UpdatePromotionOptionStatus | [UpdateStatusRequest](#memberPrivate.UpdateStatusRequest) | [UpdateStatusResponse](#memberPrivate.UpdateStatusResponse) |  |
| CreateStaffShift | [CreateStaffShiftRequest](#memberPrivate.CreateStaffShiftRequest) | [CreateStaffShiftResponse](#memberPrivate.CreateStaffShiftResponse) | CreateStaffShift 交班 |
| ListStaffShift | [ListStaffShiftRequest](#memberPrivate.ListStaffShiftRequest) | [ListStaffShiftResponse](#memberPrivate.ListStaffShiftResponse) | ListStaffShift 交班列表 |
| GetStaffShiftTime | [GetStaffShiftTimeRequest](#memberPrivate.GetStaffShiftTimeRequest) | [GetStaffShiftTimeResponse](#memberPrivate.GetStaffShiftTimeResponse) | GetStaffShiftTime 获取当前班次时间 |

 



<a name="proto/member-private/option.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/member-private/option.proto



<a name="memberPrivate.CreatePromotionOptionRequest"></a>

### CreatePromotionOptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| promotion_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| recharge_value | [int64](#int64) |  |  |
| base_value | [int64](#int64) |  |  |
| gift_value | [int64](#int64) |  |  |
| describe | [string](#string) |  |  |
| tag_id | [string](#string) |  |  |






<a name="memberPrivate.ListPromotionOptionRequest"></a>

### ListPromotionOptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| promotion_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |
| order_by | [string](#string) |  |  |
| with_page | [bool](#bool) |  |  |
| ids | [string](#string) | repeated |  |
| promotion_ids | [string](#string) | repeated |  |
| tag_id | [string](#string) |  |  |
| tag_ids | [string](#string) | repeated |  |
| filter_promotion | [ListPromotionOptionRequest.FilterPromotion](#memberPrivate.ListPromotionOptionRequest.FilterPromotion) |  |  |






<a name="memberPrivate.ListPromotionOptionRequest.FilterPromotion"></a>

### ListPromotionOptionRequest.FilterPromotion



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| begin | [int64](#int64) |  |  |
| end | [int64](#int64) |  |  |
| status | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |






<a name="memberPrivate.ListPromotionOptionResponse"></a>

### ListPromotionOptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ListPromotionOptionResponse.ListPromotionOptionData](#memberPrivate.ListPromotionOptionResponse.ListPromotionOptionData) |  |  |






<a name="memberPrivate.ListPromotionOptionResponse.ListPromotionOptionData"></a>

### ListPromotionOptionResponse.ListPromotionOptionData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [PromotionOption](#memberPrivate.PromotionOption) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="memberPrivate.OptionResponse"></a>

### OptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="memberPrivate.PromotionOption"></a>

### PromotionOption



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| promotion_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| recharge_value | [int64](#int64) |  |  |
| base_value | [int64](#int64) |  |  |
| gift_value | [int64](#int64) |  |  |
| describe | [string](#string) |  |  |
| tag_id | [string](#string) |  |  |
| created_at | [int64](#int64) |  |  |






<a name="memberPrivate.ShowPromotionOptionRequest"></a>

### ShowPromotionOptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| option_id | [string](#string) |  |  |






<a name="memberPrivate.ShowPromotionOptionResponse"></a>

### ShowPromotionOptionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [PromotionOption](#memberPrivate.PromotionOption) |  |  |






<a name="memberPrivate.UpdatePromotionOptionRequest"></a>

### UpdatePromotionOptionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| promotion_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| recharge_value | [int64](#int64) |  |  |
| base_value | [int64](#int64) |  |  |
| gift_value | [int64](#int64) |  |  |
| describe | [string](#string) |  |  |
| tag_id | [string](#string) |  |  |





 

 

 

 



<a name="proto/member-private/promotion.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/member-private/promotion.proto



<a name="memberPrivate.CreatePromotionRequest"></a>

### CreatePromotionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| begin_at | [int64](#int64) |  |  |
| end_at | [int64](#int64) |  |  |
| status | [string](#string) |  |  |
| branch_ids | [string](#string) | repeated |  |






<a name="memberPrivate.CreatePromotionResponse"></a>

### CreatePromotionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [Promotion](#memberPrivate.Promotion) |  |  |






<a name="memberPrivate.ListPromotionRequest"></a>

### ListPromotionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| status | [string](#string) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |
| order_by | [string](#string) |  |  |
| branch_ids | [string](#string) | repeated |  |
| with_page | [bool](#bool) |  |  |






<a name="memberPrivate.ListPromotionResponse"></a>

### ListPromotionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ListPromotionResponse.ListPromotionData](#memberPrivate.ListPromotionResponse.ListPromotionData) |  |  |






<a name="memberPrivate.ListPromotionResponse.ListPromotionData"></a>

### ListPromotionResponse.ListPromotionData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Promotion](#memberPrivate.Promotion) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="memberPrivate.Promotion"></a>

### Promotion



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| begin_at | [int64](#int64) |  |  |
| end_at | [int64](#int64) |  |  |
| status | [string](#string) |  |  |
| created_at | [int64](#int64) |  |  |
| branch_ids | [string](#string) | repeated |  |






<a name="memberPrivate.PromotionResponse"></a>

### PromotionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="memberPrivate.ShowPromotionRequest"></a>

### ShowPromotionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| promotion_id | [string](#string) |  |  |






<a name="memberPrivate.ShowPromotionResponse"></a>

### ShowPromotionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [Promotion](#memberPrivate.Promotion) |  |  |






<a name="memberPrivate.UpdatePromotionRequest"></a>

### UpdatePromotionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| begin_at | [int64](#int64) |  |  |
| end_at | [int64](#int64) |  |  |
| status | [string](#string) |  |  |
| branch_ids | [string](#string) | repeated |  |





 

 

 

 



<a name="proto/member-private/staff_shift.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/member-private/staff_shift.proto



<a name="memberPrivate.CreateStaffShiftRequest"></a>

### CreateStaffShiftRequest
CreateStaffShiftRequest 新增员工交班


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| shift_time | [int64](#int64) |  |  |






<a name="memberPrivate.CreateStaffShiftResponse"></a>

### CreateStaffShiftResponse
CreateStaffShiftResponse 新增员工交班


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [CreateStaffShiftResponse.Shift](#memberPrivate.CreateStaffShiftResponse.Shift) |  |  |






<a name="memberPrivate.CreateStaffShiftResponse.Shift"></a>

### CreateStaffShiftResponse.Shift



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| begin_time | [int64](#int64) |  |  |
| end_time | [int64](#int64) |  |  |






<a name="memberPrivate.GetStaffShiftTimeRequest"></a>

### GetStaffShiftTimeRequest
GetStaffShiftTimeRequest 获取员工当前班次时间


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |






<a name="memberPrivate.GetStaffShiftTimeResponse"></a>

### GetStaffShiftTimeResponse
GetStaffShiftTimeResponse 获取员工当前班次时间


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GetStaffShiftTimeResponse.Shift](#memberPrivate.GetStaffShiftTimeResponse.Shift) |  |  |






<a name="memberPrivate.GetStaffShiftTimeResponse.Shift"></a>

### GetStaffShiftTimeResponse.Shift



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| begin_time | [int64](#int64) |  |  |
| end_time | [int64](#int64) |  |  |






<a name="memberPrivate.ListStaffShiftRequest"></a>

### ListStaffShiftRequest
ListStaffShiftRequest 班次列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| begin_at | [int64](#int64) |  |  |
| end_at | [int64](#int64) |  |  |
| order_by | [string](#string) |  |  |






<a name="memberPrivate.ListStaffShiftResponse"></a>

### ListStaffShiftResponse
ListStaffShiftResponse 班次列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ListStaffShiftResponse.StaffShift](#memberPrivate.ListStaffShiftResponse.StaffShift) | repeated |  |






<a name="memberPrivate.ListStaffShiftResponse.StaffShift"></a>

### ListStaffShiftResponse.StaffShift



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| begin_time | [int64](#int64) |  |  |
| end_time | [int64](#int64) |  |  |
| created_at | [int64](#int64) |  |  |





 

 

 

 



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

