# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/merchant-basic/branch_appointment.proto](#proto/merchant-basic/branch_appointment.proto)
    - [BranchAppointmentData](#merchantBasic.BranchAppointmentData)
    - [BranchAppointmentVR](#merchantBasic.BranchAppointmentVR)
    - [RoomTypeNum](#merchantBasic.RoomTypeNum)
    - [SaveBranchAppointmentRequest](#merchantBasic.SaveBranchAppointmentRequest)
    - [SaveBranchAppointmentResponse](#merchantBasic.SaveBranchAppointmentResponse)
    - [ShowBranchAppointmentBusinessTimeData](#merchantBasic.ShowBranchAppointmentBusinessTimeData)
    - [ShowBranchAppointmentBusinessTimeRequest](#merchantBasic.ShowBranchAppointmentBusinessTimeRequest)
    - [ShowBranchAppointmentBusinessTimeResponse](#merchantBasic.ShowBranchAppointmentBusinessTimeResponse)
    - [ShowBranchAppointmentRequest](#merchantBasic.ShowBranchAppointmentRequest)
    - [ShowBranchAppointmentResponse](#merchantBasic.ShowBranchAppointmentResponse)
    - [UpdateBranchAppointmentRoomTypeRequest](#merchantBasic.UpdateBranchAppointmentRoomTypeRequest)
    - [UpdateBranchAppointmentRoomTypeResponse](#merchantBasic.UpdateBranchAppointmentRoomTypeResponse)
  
- [proto/merchant-basic/branch_business.proto](#proto/merchant-basic/branch_business.proto)
    - [BranchBusiness](#merchantBasic.BranchBusiness)
    - [CreateBranchBusinessSpecialRequest](#merchantBasic.CreateBranchBusinessSpecialRequest)
    - [GetBranchBusinessesData](#merchantBasic.GetBranchBusinessesData)
    - [GetBranchBusinessesRequest](#merchantBasic.GetBranchBusinessesRequest)
    - [GetBranchBusinessesResponse](#merchantBasic.GetBranchBusinessesResponse)
    - [GetBranchLatelyBusinessRequest](#merchantBasic.GetBranchLatelyBusinessRequest)
    - [GetBranchLatelyBusinessResponse](#merchantBasic.GetBranchLatelyBusinessResponse)
    - [GetBranchLatelyBusinessResponse.Business](#merchantBasic.GetBranchLatelyBusinessResponse.Business)
    - [UpdateBranchBusinessNormalRequest](#merchantBasic.UpdateBranchBusinessNormalRequest)
    - [UpdateBranchBusinessSpecialRequest](#merchantBasic.UpdateBranchBusinessSpecialRequest)
    - [UpdateBranchBusinessStatusRequest](#merchantBasic.UpdateBranchBusinessStatusRequest)
  
- [proto/merchant-basic/branch.proto](#proto/merchant-basic/branch.proto)
    - [BranchInfo](#merchantBasic.BranchInfo)
    - [BranchesData](#merchantBasic.BranchesData)
    - [CreateBranchRequest](#merchantBasic.CreateBranchRequest)
    - [CreateBranchResponse](#merchantBasic.CreateBranchResponse)
    - [DeleteBranchRequest](#merchantBasic.DeleteBranchRequest)
    - [DeleteBranchResponse](#merchantBasic.DeleteBranchResponse)
    - [GetBranchesByTagIDsData](#merchantBasic.GetBranchesByTagIDsData)
    - [GetBranchesByTagIDsRequest](#merchantBasic.GetBranchesByTagIDsRequest)
    - [GetBranchesByTagIDsResponse](#merchantBasic.GetBranchesByTagIDsResponse)
    - [GetBranchesRequest](#merchantBasic.GetBranchesRequest)
    - [GetBranchesResponse](#merchantBasic.GetBranchesResponse)
    - [ShowBranchRequest](#merchantBasic.ShowBranchRequest)
    - [ShowBranchResponse](#merchantBasic.ShowBranchResponse)
    - [UpdateBranchAccountRequest](#merchantBasic.UpdateBranchAccountRequest)
    - [UpdateBranchAccountResponse](#merchantBasic.UpdateBranchAccountResponse)
    - [UpdateBranchRequest](#merchantBasic.UpdateBranchRequest)
    - [UpdateBranchResponse](#merchantBasic.UpdateBranchResponse)
    - [UpdateBranchStatusRequest](#merchantBasic.UpdateBranchStatusRequest)
    - [UpdateBranchStatusResponse](#merchantBasic.UpdateBranchStatusResponse)
  
- [proto/merchant-basic/branch_tag.proto](#proto/merchant-basic/branch_tag.proto)
    - [BranchTagData](#merchantBasic.BranchTagData)
    - [CreateBranchTagRequest](#merchantBasic.CreateBranchTagRequest)
    - [CreateBranchTagResponse](#merchantBasic.CreateBranchTagResponse)
    - [GetBranchTagsByIDsRequest](#merchantBasic.GetBranchTagsByIDsRequest)
    - [GetBranchTagsByIDsResponse](#merchantBasic.GetBranchTagsByIDsResponse)
    - [GetBranchTagsData](#merchantBasic.GetBranchTagsData)
    - [GetBranchTagsRequest](#merchantBasic.GetBranchTagsRequest)
    - [GetBranchTagsResponse](#merchantBasic.GetBranchTagsResponse)
    - [ShowBranchTagRequest](#merchantBasic.ShowBranchTagRequest)
    - [ShowBranchTagResponse](#merchantBasic.ShowBranchTagResponse)
    - [UpdateBranchTagRequest](#merchantBasic.UpdateBranchTagRequest)
    - [UpdateBranchTagResponse](#merchantBasic.UpdateBranchTagResponse)
    - [UpdateBranchTagStatusRequest](#merchantBasic.UpdateBranchTagStatusRequest)
    - [UpdateBranchTagStatusResponse](#merchantBasic.UpdateBranchTagStatusResponse)
  
- [proto/merchant-basic/brand.proto](#proto/merchant-basic/brand.proto)
    - [Brand](#merchantBasic.Brand)
    - [CreateBrandRequest](#merchantBasic.CreateBrandRequest)
    - [GetBrandsData](#merchantBasic.GetBrandsData)
    - [GetBrandsRequest](#merchantBasic.GetBrandsRequest)
    - [GetBrandsResponse](#merchantBasic.GetBrandsResponse)
    - [UpdateBrandRequest](#merchantBasic.UpdateBrandRequest)
    - [UpdateBrandStatusRequest](#merchantBasic.UpdateBrandStatusRequest)
  
- [proto/merchant-basic/district.proto](#proto/merchant-basic/district.proto)
    - [CreateDistrictRequest](#merchantBasic.CreateDistrictRequest)
    - [CreateDistrictResponse](#merchantBasic.CreateDistrictResponse)
    - [District](#merchantBasic.District)
    - [GetDistrictsData](#merchantBasic.GetDistrictsData)
    - [GetDistrictsRequest](#merchantBasic.GetDistrictsRequest)
    - [GetDistrictsResponse](#merchantBasic.GetDistrictsResponse)
    - [UpdateDistrictRequest](#merchantBasic.UpdateDistrictRequest)
    - [UpdateDistrictResponse](#merchantBasic.UpdateDistrictResponse)
  
- [proto/merchant-basic/growth_config.proto](#proto/merchant-basic/growth_config.proto)
    - [GrowthConfig](#merchantBasic.GrowthConfig)
    - [SaveGrowthConfigRequest](#merchantBasic.SaveGrowthConfigRequest)
    - [ShowGrowthConfigResponse](#merchantBasic.ShowGrowthConfigResponse)
  
- [proto/merchant-basic/growth_rule.proto](#proto/merchant-basic/growth_rule.proto)
    - [CreateGrowthRuleRequest](#merchantBasic.CreateGrowthRuleRequest)
    - [GetBranchGrowthRuleRequest](#merchantBasic.GetBranchGrowthRuleRequest)
    - [GetBranchGrowthRuleResponse](#merchantBasic.GetBranchGrowthRuleResponse)
    - [GetBranchesHasGrowthRuleResponse](#merchantBasic.GetBranchesHasGrowthRuleResponse)
    - [GetGrowthRulesRequest](#merchantBasic.GetGrowthRulesRequest)
    - [GetGrowthRulesResponse](#merchantBasic.GetGrowthRulesResponse)
    - [GrowthGain](#merchantBasic.GrowthGain)
    - [GrowthRule](#merchantBasic.GrowthRule)
    - [GrowthRulesData](#merchantBasic.GrowthRulesData)
    - [ShowGrowthRuleRequest](#merchantBasic.ShowGrowthRuleRequest)
    - [ShowGrowthRuleResponse](#merchantBasic.ShowGrowthRuleResponse)
    - [UpdateGrowthRuleRequest](#merchantBasic.UpdateGrowthRuleRequest)
  
- [proto/merchant-basic/invoice.proto](#proto/merchant-basic/invoice.proto)
    - [CreateInvoiceRequest](#merchantBasic.CreateInvoiceRequest)
    - [CreateInvoiceResponse](#merchantBasic.CreateInvoiceResponse)
  
- [proto/merchant-basic/member_address.proto](#proto/merchant-basic/member_address.proto)
    - [DeleteMemberAddressRequest](#merchantBasic.DeleteMemberAddressRequest)
    - [GetMemberAddressRequest](#merchantBasic.GetMemberAddressRequest)
    - [GetMemberAddressResponse](#merchantBasic.GetMemberAddressResponse)
    - [GetMemberDefaultAddressRequest](#merchantBasic.GetMemberDefaultAddressRequest)
    - [GetMemberDefaultAddressResponse](#merchantBasic.GetMemberDefaultAddressResponse)
    - [MemberAddress](#merchantBasic.MemberAddress)
    - [SetMemberAddressDefaultRequest](#merchantBasic.SetMemberAddressDefaultRequest)
  
- [proto/merchant-basic/member_behavior.proto](#proto/merchant-basic/member_behavior.proto)
    - [CreateMemberBehaviorRequest](#merchantBasic.CreateMemberBehaviorRequest)
    - [GetMemberBehaviorsRequest](#merchantBasic.GetMemberBehaviorsRequest)
    - [GetMemberBehaviorsResponse](#merchantBasic.GetMemberBehaviorsResponse)
    - [MemberBehavior](#merchantBasic.MemberBehavior)
    - [MemberBehaviors](#merchantBasic.MemberBehaviors)
  
- [proto/merchant-basic/member.proto](#proto/merchant-basic/member.proto)
    - [CreateMemberRequest](#merchantBasic.CreateMemberRequest)
    - [CreateMemberResponse](#merchantBasic.CreateMemberResponse)
    - [CreateOrUpdateCallingMemberRequest](#merchantBasic.CreateOrUpdateCallingMemberRequest)
    - [CreateWechatMemberRequest](#merchantBasic.CreateWechatMemberRequest)
    - [CreateWechatMemberResponse](#merchantBasic.CreateWechatMemberResponse)
    - [GetBirthdayMembersRequest](#merchantBasic.GetBirthdayMembersRequest)
    - [GetBirthdayMembersResponse](#merchantBasic.GetBirthdayMembersResponse)
    - [GetMembersByIDsRequest](#merchantBasic.GetMembersByIDsRequest)
    - [GetMembersByIDsResponse](#merchantBasic.GetMembersByIDsResponse)
    - [GetMembersByPhoneSuffixRequest](#merchantBasic.GetMembersByPhoneSuffixRequest)
    - [GetMembersByPhoneSuffixResponse](#merchantBasic.GetMembersByPhoneSuffixResponse)
    - [GetMembersRequest](#merchantBasic.GetMembersRequest)
    - [GetMembersResponse](#merchantBasic.GetMembersResponse)
    - [MemberInfo](#merchantBasic.MemberInfo)
    - [MembersData](#merchantBasic.MembersData)
    - [ShowMemberByAccuratePhoneRequest](#merchantBasic.ShowMemberByAccuratePhoneRequest)
    - [ShowMemberByAccuratePhoneResponse](#merchantBasic.ShowMemberByAccuratePhoneResponse)
    - [ShowMemberRequest](#merchantBasic.ShowMemberRequest)
    - [ShowMemberResponse](#merchantBasic.ShowMemberResponse)
    - [UpdateMemberBranchInfoRequest](#merchantBasic.UpdateMemberBranchInfoRequest)
    - [UpdateMemberRequest](#merchantBasic.UpdateMemberRequest)
    - [UpdateMemberResponse](#merchantBasic.UpdateMemberResponse)
  
- [proto/merchant-basic/merchant_basic.proto](#proto/merchant-basic/merchant_basic.proto)
    - [Empty](#merchantBasic.Empty)
    - [Response](#merchantBasic.Response)
    - [UpdateStatusRequest](#merchantBasic.UpdateStatusRequest)
  
    - [MerchantBasicService](#merchantBasic.MerchantBasicService)
  
- [proto/merchant-basic/payment.proto](#proto/merchant-basic/payment.proto)
    - [BranchSubMchID](#merchantBasic.BranchSubMchID)
    - [BranchSubMchIDData](#merchantBasic.BranchSubMchIDData)
    - [BranchWechatPaymentSetting](#merchantBasic.BranchWechatPaymentSetting)
    - [GetBranchWechatPaymentSettingRequest](#merchantBasic.GetBranchWechatPaymentSettingRequest)
    - [GetBranchWechatPaymentSettingResponse](#merchantBasic.GetBranchWechatPaymentSettingResponse)
    - [GetBranchWithSubMchIDRequest](#merchantBasic.GetBranchWithSubMchIDRequest)
    - [GetBranchWithSubMchIDResponse](#merchantBasic.GetBranchWithSubMchIDResponse)
    - [GetWechatPaySettingByAppIDRequest](#merchantBasic.GetWechatPaySettingByAppIDRequest)
    - [GetWechatPaySettingRequest](#merchantBasic.GetWechatPaySettingRequest)
    - [GetWechatPaySettingResponse](#merchantBasic.GetWechatPaySettingResponse)
    - [SetBranchSubMchIDRequest](#merchantBasic.SetBranchSubMchIDRequest)
    - [SetWechatPaySettingRequest](#merchantBasic.SetWechatPaySettingRequest)
    - [WechatPaySetting](#merchantBasic.WechatPaySetting)
  
- [proto/merchant-basic/permission.proto](#proto/merchant-basic/permission.proto)
    - [CreatePermissionsRequest](#merchantBasic.CreatePermissionsRequest)
    - [CreatePermissionsResponse](#merchantBasic.CreatePermissionsResponse)
    - [GetPermissionsRequest](#merchantBasic.GetPermissionsRequest)
    - [GetPermissionsResponse](#merchantBasic.GetPermissionsResponse)
    - [PermissionInfo](#merchantBasic.PermissionInfo)
  
- [proto/merchant-basic/point.proto](#proto/merchant-basic/point.proto)
    - [CreatePointCategoryRequest](#merchantBasic.CreatePointCategoryRequest)
    - [CreatePointRuleRequest](#merchantBasic.CreatePointRuleRequest)
    - [GetBranchPointRuleRequest](#merchantBasic.GetBranchPointRuleRequest)
    - [GetBranchPointRuleResponse](#merchantBasic.GetBranchPointRuleResponse)
    - [GetBranchPointRuleResponse.Data](#merchantBasic.GetBranchPointRuleResponse.Data)
    - [GetBranchPointRuleResponse.Rule](#merchantBasic.GetBranchPointRuleResponse.Rule)
    - [GetPointRuleAllBranchResponse](#merchantBasic.GetPointRuleAllBranchResponse)
    - [GetPointRuleAllBranchResponse.Data](#merchantBasic.GetPointRuleAllBranchResponse.Data)
    - [GetPointRuleDescribeResponse](#merchantBasic.GetPointRuleDescribeResponse)
    - [GetPointRuleDescribeResponse.Data](#merchantBasic.GetPointRuleDescribeResponse.Data)
    - [ListPointCategoryRequest](#merchantBasic.ListPointCategoryRequest)
    - [ListPointCategoryResponse](#merchantBasic.ListPointCategoryResponse)
    - [ListPointCategoryResponse.Data](#merchantBasic.ListPointCategoryResponse.Data)
    - [ListPointRuleRequest](#merchantBasic.ListPointRuleRequest)
    - [ListPointRuleResponse](#merchantBasic.ListPointRuleResponse)
    - [ListPointRuleResponse.Data](#merchantBasic.ListPointRuleResponse.Data)
    - [PointCategory](#merchantBasic.PointCategory)
    - [PointRule](#merchantBasic.PointRule)
    - [PointRuleDetail](#merchantBasic.PointRuleDetail)
    - [SetPointRuleDescribeRequest](#merchantBasic.SetPointRuleDescribeRequest)
    - [ShowPointCategoryRequest](#merchantBasic.ShowPointCategoryRequest)
    - [ShowPointCategoryResponse](#merchantBasic.ShowPointCategoryResponse)
    - [ShowPointRuleRequest](#merchantBasic.ShowPointRuleRequest)
    - [ShowPointRuleResponse](#merchantBasic.ShowPointRuleResponse)
    - [UpdatePointCategoryRequest](#merchantBasic.UpdatePointCategoryRequest)
    - [UpdatePointRuleRequest](#merchantBasic.UpdatePointRuleRequest)
  
- [proto/merchant-basic/role.proto](#proto/merchant-basic/role.proto)
    - [CreateRoleRequest](#merchantBasic.CreateRoleRequest)
    - [CreateRoleResponse](#merchantBasic.CreateRoleResponse)
    - [DeleteRoleRequest](#merchantBasic.DeleteRoleRequest)
    - [DeleteRoleResponse](#merchantBasic.DeleteRoleResponse)
    - [GetRoleHistoriesRequest](#merchantBasic.GetRoleHistoriesRequest)
    - [GetRoleHistoriesResponse](#merchantBasic.GetRoleHistoriesResponse)
    - [GetRolesRequest](#merchantBasic.GetRolesRequest)
    - [GetRolesResponse](#merchantBasic.GetRolesResponse)
    - [RoleHistoriesData](#merchantBasic.RoleHistoriesData)
    - [RoleInfo](#merchantBasic.RoleInfo)
    - [RolePermissionInfo](#merchantBasic.RolePermissionInfo)
    - [RolesData](#merchantBasic.RolesData)
    - [ShowRoleRequest](#merchantBasic.ShowRoleRequest)
    - [ShowRoleResponse](#merchantBasic.ShowRoleResponse)
    - [Snapshot](#merchantBasic.Snapshot)
    - [UpdateRoleRequest](#merchantBasic.UpdateRoleRequest)
    - [UpdateRoleResponse](#merchantBasic.UpdateRoleResponse)
    - [UpdateRoleStatusRequest](#merchantBasic.UpdateRoleStatusRequest)
    - [UpdateRoleStatusResponse](#merchantBasic.UpdateRoleStatusResponse)
  
- [proto/merchant-basic/room_type_category.proto](#proto/merchant-basic/room_type_category.proto)
    - [CreateRoomTypeCategoryRequest](#merchantBasic.CreateRoomTypeCategoryRequest)
    - [CreateRoomTypeCategoryResponse](#merchantBasic.CreateRoomTypeCategoryResponse)
    - [GetRoomTypeCategoriesRequest](#merchantBasic.GetRoomTypeCategoriesRequest)
    - [GetRoomTypeCategoriesResponse](#merchantBasic.GetRoomTypeCategoriesResponse)
    - [RoomTypeCategoriesData](#merchantBasic.RoomTypeCategoriesData)
    - [RoomTypeCategory](#merchantBasic.RoomTypeCategory)
    - [UpdateRoomTypeCategoryRequest](#merchantBasic.UpdateRoomTypeCategoryRequest)
    - [UpdateRoomTypeCategoryResponse](#merchantBasic.UpdateRoomTypeCategoryResponse)
  
- [proto/merchant-basic/room_type.proto](#proto/merchant-basic/room_type.proto)
    - [CreateRoomTypeRequest](#merchantBasic.CreateRoomTypeRequest)
    - [CreateRoomTypeResponse](#merchantBasic.CreateRoomTypeResponse)
    - [GetRoomTypesByIDsRequest](#merchantBasic.GetRoomTypesByIDsRequest)
    - [GetRoomTypesByIDsResponse](#merchantBasic.GetRoomTypesByIDsResponse)
    - [GetRoomTypesRequest](#merchantBasic.GetRoomTypesRequest)
    - [GetRoomTypesResponse](#merchantBasic.GetRoomTypesResponse)
    - [RoomType](#merchantBasic.RoomType)
    - [RoomTypesData](#merchantBasic.RoomTypesData)
    - [ShowRoomTypeRequest](#merchantBasic.ShowRoomTypeRequest)
    - [ShowRoomTypeResponse](#merchantBasic.ShowRoomTypeResponse)
    - [UpdateRoomTypeRequest](#merchantBasic.UpdateRoomTypeRequest)
    - [UpdateRoomTypeResponse](#merchantBasic.UpdateRoomTypeResponse)
    - [UpdateRoomTypeStatusRequest](#merchantBasic.UpdateRoomTypeStatusRequest)
  
- [proto/merchant-basic/staff.proto](#proto/merchant-basic/staff.proto)
    - [CreateStaffRequest](#merchantBasic.CreateStaffRequest)
    - [CreateStaffResponse](#merchantBasic.CreateStaffResponse)
    - [DeleteStaffRequest](#merchantBasic.DeleteStaffRequest)
    - [DeleteStaffResponse](#merchantBasic.DeleteStaffResponse)
    - [GetStaffsByRoleIDRequest](#merchantBasic.GetStaffsByRoleIDRequest)
    - [GetStaffsByRoleIDResponse](#merchantBasic.GetStaffsByRoleIDResponse)
    - [GetStaffsRequest](#merchantBasic.GetStaffsRequest)
    - [GetStaffsResponse](#merchantBasic.GetStaffsResponse)
    - [ResetPasswordRequest](#merchantBasic.ResetPasswordRequest)
    - [ResetPasswordResponse](#merchantBasic.ResetPasswordResponse)
    - [ShowStaffByPhoneRequest](#merchantBasic.ShowStaffByPhoneRequest)
    - [ShowStaffByPhoneResponse](#merchantBasic.ShowStaffByPhoneResponse)
    - [ShowStaffRequest](#merchantBasic.ShowStaffRequest)
    - [ShowStaffResponse](#merchantBasic.ShowStaffResponse)
    - [SignInData](#merchantBasic.SignInData)
    - [SignInRequest](#merchantBasic.SignInRequest)
    - [SignInResponse](#merchantBasic.SignInResponse)
    - [StaffFullInfo](#merchantBasic.StaffFullInfo)
    - [StaffInfo](#merchantBasic.StaffInfo)
    - [StaffsData](#merchantBasic.StaffsData)
    - [UpdatePasswordRequest](#merchantBasic.UpdatePasswordRequest)
    - [UpdatePasswordResponse](#merchantBasic.UpdatePasswordResponse)
    - [UpdateStaffRequest](#merchantBasic.UpdateStaffRequest)
    - [UpdateStaffResponse](#merchantBasic.UpdateStaffResponse)
    - [UpdateStaffStatusRequest](#merchantBasic.UpdateStaffStatusRequest)
    - [UpdateStaffStatusResponse](#merchantBasic.UpdateStaffStatusResponse)
  
- [proto/merchant-basic/wechat_member.proto](#proto/merchant-basic/wechat_member.proto)
    - [CreateOrUpdateWechatUserRequest](#merchantBasic.CreateOrUpdateWechatUserRequest)
    - [CreateOrUpdateWechatUserResponse](#merchantBasic.CreateOrUpdateWechatUserResponse)
    - [GetWechatUserRequest](#merchantBasic.GetWechatUserRequest)
    - [GetWechatUserResponse](#merchantBasic.GetWechatUserResponse)
    - [ShowWechatUserByMemberRequest](#merchantBasic.ShowWechatUserByMemberRequest)
    - [ShowWechatUserByMemberResponse](#merchantBasic.ShowWechatUserByMemberResponse)
    - [ShowWechatUserRequest](#merchantBasic.ShowWechatUserRequest)
    - [ShowWechatUserResponse](#merchantBasic.ShowWechatUserResponse)
    - [WechatUser](#merchantBasic.WechatUser)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto/merchant-basic/branch_appointment.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/branch_appointment.proto



<a name="merchantBasic.BranchAppointmentData"></a>

### BranchAppointmentData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| open_appointment | [bool](#bool) |  |  |
| appointment_granularity | [int32](#int32) |  |  |
| vr | [BranchAppointmentVR](#merchantBasic.BranchAppointmentVR) | repeated |  |
| video | [string](#string) | repeated |  |
| environment | [string](#string) | repeated |  |
| meal | [string](#string) | repeated |  |
| price | [string](#string) | repeated |  |
| parking | [string](#string) |  |  |
| hot | [bool](#bool) |  |  |
| branch_id | [string](#string) |  |  |
| room_types | [string](#string) |  |  |
| business_hours | [BranchBusiness](#merchantBasic.BranchBusiness) | repeated |  |






<a name="merchantBasic.BranchAppointmentVR"></a>

### BranchAppointmentVR



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |






<a name="merchantBasic.RoomTypeNum"></a>

### RoomTypeNum



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_type_id | [string](#string) |  |  |
| room_type_num | [int32](#int32) |  |  |






<a name="merchantBasic.SaveBranchAppointmentRequest"></a>

### SaveBranchAppointmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| open_appointment | [bool](#bool) |  |  |
| appointment_granularity | [int32](#int32) |  |  |
| vr | [BranchAppointmentVR](#merchantBasic.BranchAppointmentVR) | repeated |  |
| video | [string](#string) | repeated |  |
| environment | [string](#string) | repeated |  |
| meal | [string](#string) | repeated |  |
| price | [string](#string) | repeated |  |
| parking | [string](#string) |  |  |
| hot | [bool](#bool) |  |  |






<a name="merchantBasic.SaveBranchAppointmentResponse"></a>

### SaveBranchAppointmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.ShowBranchAppointmentBusinessTimeData"></a>

### ShowBranchAppointmentBusinessTimeData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| begin_time | [string](#string) |  |  |
| end_time | [string](#string) |  |  |
| is_next_day | [bool](#bool) |  |  |
| granularity | [int32](#int32) |  |  |
| open_appointment | [bool](#bool) |  |  |
| hot | [bool](#bool) |  |  |






<a name="merchantBasic.ShowBranchAppointmentBusinessTimeRequest"></a>

### ShowBranchAppointmentBusinessTimeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| date | [int32](#int32) |  |  |






<a name="merchantBasic.ShowBranchAppointmentBusinessTimeResponse"></a>

### ShowBranchAppointmentBusinessTimeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ShowBranchAppointmentBusinessTimeData](#merchantBasic.ShowBranchAppointmentBusinessTimeData) |  |  |






<a name="merchantBasic.ShowBranchAppointmentRequest"></a>

### ShowBranchAppointmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |






<a name="merchantBasic.ShowBranchAppointmentResponse"></a>

### ShowBranchAppointmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [BranchAppointmentData](#merchantBasic.BranchAppointmentData) |  |  |






<a name="merchantBasic.UpdateBranchAppointmentRoomTypeRequest"></a>

### UpdateBranchAppointmentRoomTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_type | [RoomTypeNum](#merchantBasic.RoomTypeNum) | repeated |  |
| branch_id | [string](#string) |  |  |






<a name="merchantBasic.UpdateBranchAppointmentRoomTypeResponse"></a>

### UpdateBranchAppointmentRoomTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/branch_business.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/branch_business.proto



<a name="merchantBasic.BranchBusiness"></a>

### BranchBusiness



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| begin_date | [string](#string) |  |  |
| end_date | [string](#string) |  |  |
| weeks | [int32](#int32) | repeated |  |
| begin_time | [string](#string) |  |  |
| end_time | [string](#string) |  |  |
| is_next_day | [bool](#bool) |  |  |
| status | [string](#string) |  |  |
| category | [string](#string) |  |  |






<a name="merchantBasic.CreateBranchBusinessSpecialRequest"></a>

### CreateBranchBusinessSpecialRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| business | [BranchBusiness](#merchantBasic.BranchBusiness) |  |  |






<a name="merchantBasic.GetBranchBusinessesData"></a>

### GetBranchBusinessesData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int32](#int32) |  |  |
| businesses | [BranchBusiness](#merchantBasic.BranchBusiness) | repeated |  |






<a name="merchantBasic.GetBranchBusinessesRequest"></a>

### GetBranchBusinessesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| status | [string](#string) |  |  |
| category | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |






<a name="merchantBasic.GetBranchBusinessesResponse"></a>

### GetBranchBusinessesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GetBranchBusinessesData](#merchantBasic.GetBranchBusinessesData) |  |  |






<a name="merchantBasic.GetBranchLatelyBusinessRequest"></a>

### GetBranchLatelyBusinessRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| date_time | [int64](#int64) |  |  |






<a name="merchantBasic.GetBranchLatelyBusinessResponse"></a>

### GetBranchLatelyBusinessResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GetBranchLatelyBusinessResponse.Business](#merchantBasic.GetBranchLatelyBusinessResponse.Business) |  |  |






<a name="merchantBasic.GetBranchLatelyBusinessResponse.Business"></a>

### GetBranchLatelyBusinessResponse.Business



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| business_date | [string](#string) |  |  |
| begin_time | [string](#string) |  |  |
| end_time | [string](#string) |  |  |
| is_next_day | [bool](#bool) |  |  |






<a name="merchantBasic.UpdateBranchBusinessNormalRequest"></a>

### UpdateBranchBusinessNormalRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| businesses | [BranchBusiness](#merchantBasic.BranchBusiness) | repeated |  |






<a name="merchantBasic.UpdateBranchBusinessSpecialRequest"></a>

### UpdateBranchBusinessSpecialRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| business | [BranchBusiness](#merchantBasic.BranchBusiness) |  |  |






<a name="merchantBasic.UpdateBranchBusinessStatusRequest"></a>

### UpdateBranchBusinessStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/branch.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/branch.proto



<a name="merchantBasic.BranchInfo"></a>

### BranchInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| province_id | [string](#string) |  |  |
| city_id | [string](#string) |  |  |
| district_id | [string](#string) |  |  |
| address | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| created_at | [int32](#int32) |  |  |
| extra | [string](#string) |  |  |
| code | [string](#string) |  |  |
| id | [string](#string) |  |  |
| latitude | [float](#float) |  |  |
| longitude | [float](#float) |  |  |
| status | [string](#string) |  |  |
| opened_at | [int32](#int32) |  |  |
| photo | [string](#string) | repeated |  |
| parking | [string](#string) |  |  |
| area_id | [string](#string) |  |  |
| weight | [int32](#int32) |  |  |
| domain | [string](#string) |  |  |
| biz_type | [int32](#int32) |  |  |
| business_status | [string](#string) |  |  |
| alias | [string](#string) |  |  |
| simplify | [string](#string) |  |  |
| location | [string](#string) |  |  |
| brand_id | [string](#string) |  |  |
| brand_name | [string](#string) |  |  |
| business_hours | [BranchBusiness](#merchantBasic.BranchBusiness) | repeated |  |






<a name="merchantBasic.BranchesData"></a>

### BranchesData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branches | [BranchInfo](#merchantBasic.BranchInfo) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="merchantBasic.CreateBranchRequest"></a>

### CreateBranchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch | [BranchInfo](#merchantBasic.BranchInfo) |  |  |
| business_hours | [BranchBusiness](#merchantBasic.BranchBusiness) | repeated |  |






<a name="merchantBasic.CreateBranchResponse"></a>

### CreateBranchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [string](#string) |  |  |






<a name="merchantBasic.DeleteBranchRequest"></a>

### DeleteBranchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.DeleteBranchResponse"></a>

### DeleteBranchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.GetBranchesByTagIDsData"></a>

### GetBranchesByTagIDsData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_tag | [BranchTagData](#merchantBasic.BranchTagData) |  |  |
| branches | [BranchInfo](#merchantBasic.BranchInfo) | repeated |  |






<a name="merchantBasic.GetBranchesByTagIDsRequest"></a>

### GetBranchesByTagIDsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag_ids | [string](#string) | repeated |  |
| status | [string](#string) |  |  |
| business_status | [string](#string) | repeated |  |






<a name="merchantBasic.GetBranchesByTagIDsResponse"></a>

### GetBranchesByTagIDsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GetBranchesByTagIDsData](#merchantBasic.GetBranchesByTagIDsData) | repeated |  |






<a name="merchantBasic.GetBranchesRequest"></a>

### GetBranchesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| province_id | [string](#string) |  |  |
| city_id | [string](#string) |  |  |
| district_id | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |
| status | [string](#string) |  |  |
| staff_id | [string](#string) |  |  |
| area_id | [string](#string) |  |  |
| business_status | [string](#string) | repeated |  |
| brand_id | [string](#string) |  |  |






<a name="merchantBasic.GetBranchesResponse"></a>

### GetBranchesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [BranchesData](#merchantBasic.BranchesData) |  |  |






<a name="merchantBasic.ShowBranchRequest"></a>

### ShowBranchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.ShowBranchResponse"></a>

### ShowBranchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [BranchInfo](#merchantBasic.BranchInfo) |  |  |






<a name="merchantBasic.UpdateBranchAccountRequest"></a>

### UpdateBranchAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| wechat_app_id | [string](#string) |  |  |
| wechat_mch_id | [string](#string) |  |  |
| wechat_key | [string](#string) |  |  |
| wechat_cert_path | [string](#string) |  |  |
| wechat_key_path | [string](#string) |  |  |






<a name="merchantBasic.UpdateBranchAccountResponse"></a>

### UpdateBranchAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.UpdateBranchRequest"></a>

### UpdateBranchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch | [BranchInfo](#merchantBasic.BranchInfo) |  |  |
| business_hours | [BranchBusiness](#merchantBasic.BranchBusiness) | repeated |  |






<a name="merchantBasic.UpdateBranchResponse"></a>

### UpdateBranchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.UpdateBranchStatusRequest"></a>

### UpdateBranchStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.UpdateBranchStatusResponse"></a>

### UpdateBranchStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/branch_tag.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/branch_tag.proto



<a name="merchantBasic.BranchTagData"></a>

### BranchTagData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| branches | [string](#string) | repeated |  |
| create_staff_id | [string](#string) |  |  |
| staff_name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| created_at | [int32](#int32) |  |  |
| updated_at | [int32](#int32) |  |  |






<a name="merchantBasic.CreateBranchTagRequest"></a>

### CreateBranchTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| branch_ids | [string](#string) | repeated |  |






<a name="merchantBasic.CreateBranchTagResponse"></a>

### CreateBranchTagResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.GetBranchTagsByIDsRequest"></a>

### GetBranchTagsByIDsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [string](#string) | repeated |  |






<a name="merchantBasic.GetBranchTagsByIDsResponse"></a>

### GetBranchTagsByIDsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [BranchTagData](#merchantBasic.BranchTagData) | repeated |  |






<a name="merchantBasic.GetBranchTagsData"></a>

### GetBranchTagsData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_tags | [BranchTagData](#merchantBasic.BranchTagData) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="merchantBasic.GetBranchTagsRequest"></a>

### GetBranchTagsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| date_start | [int32](#int32) |  |  |
| date_end | [int32](#int32) |  |  |
| status | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |
| branch_ids | [string](#string) | repeated |  |






<a name="merchantBasic.GetBranchTagsResponse"></a>

### GetBranchTagsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GetBranchTagsData](#merchantBasic.GetBranchTagsData) |  |  |






<a name="merchantBasic.ShowBranchTagRequest"></a>

### ShowBranchTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.ShowBranchTagResponse"></a>

### ShowBranchTagResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [BranchTagData](#merchantBasic.BranchTagData) |  |  |






<a name="merchantBasic.UpdateBranchTagRequest"></a>

### UpdateBranchTagRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| branch_ids | [string](#string) | repeated |  |






<a name="merchantBasic.UpdateBranchTagResponse"></a>

### UpdateBranchTagResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.UpdateBranchTagStatusRequest"></a>

### UpdateBranchTagStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.UpdateBranchTagStatusResponse"></a>

### UpdateBranchTagStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/brand.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/brand.proto



<a name="merchantBasic.Brand"></a>

### Brand



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| order | [int32](#int32) |  |  |
| status | [string](#string) |  |  |
| created_at | [int32](#int32) |  |  |






<a name="merchantBasic.CreateBrandRequest"></a>

### CreateBrandRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| order | [int32](#int32) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.GetBrandsData"></a>

### GetBrandsData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int32](#int32) |  |  |
| brands | [Brand](#merchantBasic.Brand) | repeated |  |






<a name="merchantBasic.GetBrandsRequest"></a>

### GetBrandsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.GetBrandsResponse"></a>

### GetBrandsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GetBrandsData](#merchantBasic.GetBrandsData) |  |  |






<a name="merchantBasic.UpdateBrandRequest"></a>

### UpdateBrandRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| order | [int32](#int32) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.UpdateBrandStatusRequest"></a>

### UpdateBrandStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/district.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/district.proto



<a name="merchantBasic.CreateDistrictRequest"></a>

### CreateDistrictRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.CreateDistrictResponse"></a>

### CreateDistrictResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.District"></a>

### District



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| code | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.GetDistrictsData"></a>

### GetDistrictsData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| districts | [District](#merchantBasic.District) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="merchantBasic.GetDistrictsRequest"></a>

### GetDistrictsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |






<a name="merchantBasic.GetDistrictsResponse"></a>

### GetDistrictsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GetDistrictsData](#merchantBasic.GetDistrictsData) |  |  |






<a name="merchantBasic.UpdateDistrictRequest"></a>

### UpdateDistrictRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.UpdateDistrictResponse"></a>

### UpdateDistrictResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/growth_config.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/growth_config.proto



<a name="merchantBasic.GrowthConfig"></a>

### GrowthConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| top | [uint32](#uint32) |  |  |
| rules | [string](#string) | repeated |  |






<a name="merchantBasic.SaveGrowthConfigRequest"></a>

### SaveGrowthConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [GrowthConfig](#merchantBasic.GrowthConfig) |  |  |






<a name="merchantBasic.ShowGrowthConfigResponse"></a>

### ShowGrowthConfigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GrowthConfig](#merchantBasic.GrowthConfig) |  |  |





 

 

 

 



<a name="proto/merchant-basic/growth_rule.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/growth_rule.proto



<a name="merchantBasic.CreateGrowthRuleRequest"></a>

### CreateGrowthRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rule | [GrowthRule](#merchantBasic.GrowthRule) |  |  |






<a name="merchantBasic.GetBranchGrowthRuleRequest"></a>

### GetBranchGrowthRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |






<a name="merchantBasic.GetBranchGrowthRuleResponse"></a>

### GetBranchGrowthRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GrowthRule](#merchantBasic.GrowthRule) |  |  |






<a name="merchantBasic.GetBranchesHasGrowthRuleResponse"></a>

### GetBranchesHasGrowthRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [string](#string) | repeated |  |






<a name="merchantBasic.GetGrowthRulesRequest"></a>

### GetGrowthRulesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |






<a name="merchantBasic.GetGrowthRulesResponse"></a>

### GetGrowthRulesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GrowthRulesData](#merchantBasic.GrowthRulesData) |  |  |






<a name="merchantBasic.GrowthGain"></a>

### GrowthGain



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| consume_type | [string](#string) |  |  |
| cost | [uint32](#uint32) |  |  |






<a name="merchantBasic.GrowthRule"></a>

### GrowthRule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| growth_gain | [GrowthGain](#merchantBasic.GrowthGain) | repeated |  |
| expire_day | [int32](#int32) |  |  |
| branches | [string](#string) | repeated |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.GrowthRulesData"></a>

### GrowthRulesData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rules | [GrowthRule](#merchantBasic.GrowthRule) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="merchantBasic.ShowGrowthRuleRequest"></a>

### ShowGrowthRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.ShowGrowthRuleResponse"></a>

### ShowGrowthRuleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GrowthRule](#merchantBasic.GrowthRule) |  |  |






<a name="merchantBasic.UpdateGrowthRuleRequest"></a>

### UpdateGrowthRuleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rule | [GrowthRule](#merchantBasic.GrowthRule) |  |  |





 

 

 

 



<a name="proto/merchant-basic/invoice.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/invoice.proto



<a name="merchantBasic.CreateInvoiceRequest"></a>

### CreateInvoiceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [string](#string) |  |  |
| invoice_data | [string](#string) |  |  |






<a name="merchantBasic.CreateInvoiceResponse"></a>

### CreateInvoiceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/member_address.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/member_address.proto



<a name="merchantBasic.DeleteMemberAddressRequest"></a>

### DeleteMemberAddressRequest
DeleteMemberAddressRequest 删除地址


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.GetMemberAddressRequest"></a>

### GetMemberAddressRequest
GetMemberAddressRequest 获取用户地址


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member_id | [string](#string) |  |  |






<a name="merchantBasic.GetMemberAddressResponse"></a>

### GetMemberAddressResponse
GetMemberAddressResponse 获取用户地址


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [MemberAddress](#merchantBasic.MemberAddress) | repeated |  |






<a name="merchantBasic.GetMemberDefaultAddressRequest"></a>

### GetMemberDefaultAddressRequest
GetMemberDefaultAddressRequest 获取用户默认地址


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member_id | [string](#string) |  |  |






<a name="merchantBasic.GetMemberDefaultAddressResponse"></a>

### GetMemberDefaultAddressResponse
GetMemberDefaultAddressResponse 获取用户默认地址


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [MemberAddress](#merchantBasic.MemberAddress) |  |  |






<a name="merchantBasic.MemberAddress"></a>

### MemberAddress
用户地址


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| member_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| phone_code | [string](#string) |  |  |
| province_id | [string](#string) |  |  |
| city_id | [string](#string) |  |  |
| district_id | [string](#string) |  |  |
| address | [string](#string) |  |  |
| is_default | [bool](#bool) |  |  |






<a name="merchantBasic.SetMemberAddressDefaultRequest"></a>

### SetMemberAddressDefaultRequest
SetMemberAddressDefaultRequest 设置为默认地址


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member_id | [string](#string) |  |  |
| id | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/member_behavior.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/member_behavior.proto



<a name="merchantBasic.CreateMemberBehaviorRequest"></a>

### CreateMemberBehaviorRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| behavior | [string](#string) |  |  |
| member_id | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |






<a name="merchantBasic.GetMemberBehaviorsRequest"></a>

### GetMemberBehaviorsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member_id | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |






<a name="merchantBasic.GetMemberBehaviorsResponse"></a>

### GetMemberBehaviorsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [MemberBehaviors](#merchantBasic.MemberBehaviors) |  |  |






<a name="merchantBasic.MemberBehavior"></a>

### MemberBehavior



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| member_id | [string](#string) |  |  |
| behavior | [string](#string) |  |  |
| staff_id | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| created_at | [int32](#int32) |  |  |






<a name="merchantBasic.MemberBehaviors"></a>

### MemberBehaviors



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| memberBehaviors | [MemberBehavior](#merchantBasic.MemberBehavior) | repeated |  |
| total | [int32](#int32) |  |  |





 

 

 

 



<a name="proto/merchant-basic/member.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/member.proto



<a name="merchantBasic.CreateMemberRequest"></a>

### CreateMemberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member | [MemberInfo](#merchantBasic.MemberInfo) |  |  |






<a name="merchantBasic.CreateMemberResponse"></a>

### CreateMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [string](#string) |  |  |






<a name="merchantBasic.CreateOrUpdateCallingMemberRequest"></a>

### CreateOrUpdateCallingMemberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| phone_code | [string](#string) |  |  |
| gender | [int32](#int32) |  |  |
| channel | [string](#string) |  |  |
| can_overwrite | [bool](#bool) |  |  |
| branch_id | [string](#string) |  |  |






<a name="merchantBasic.CreateWechatMemberRequest"></a>

### CreateWechatMemberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  |  |
| phone_code | [string](#string) |  |  |
| name | [string](#string) |  |  |
| gender | [int32](#int32) |  |  |
| birthday | [string](#string) |  |  |






<a name="merchantBasic.CreateWechatMemberResponse"></a>

### CreateWechatMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [string](#string) |  |  |






<a name="merchantBasic.GetBirthdayMembersRequest"></a>

### GetBirthdayMembersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| birthday | [string](#string) |  |  |






<a name="merchantBasic.GetBirthdayMembersResponse"></a>

### GetBirthdayMembersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [MemberInfo](#merchantBasic.MemberInfo) | repeated |  |






<a name="merchantBasic.GetMembersByIDsRequest"></a>

### GetMembersByIDsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [string](#string) | repeated |  |






<a name="merchantBasic.GetMembersByIDsResponse"></a>

### GetMembersByIDsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [MemberInfo](#merchantBasic.MemberInfo) | repeated |  |






<a name="merchantBasic.GetMembersByPhoneSuffixRequest"></a>

### GetMembersByPhoneSuffixRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone_suffix | [string](#string) |  |  |






<a name="merchantBasic.GetMembersByPhoneSuffixResponse"></a>

### GetMembersByPhoneSuffixResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [MemberInfo](#merchantBasic.MemberInfo) | repeated |  |






<a name="merchantBasic.GetMembersRequest"></a>

### GetMembersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| phone_code | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| channel | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |
| gender | [int32](#int32) |  |  |
| first_brand | [string](#string) |  |  |






<a name="merchantBasic.GetMembersResponse"></a>

### GetMembersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [MembersData](#merchantBasic.MembersData) |  |  |






<a name="merchantBasic.MemberInfo"></a>

### MemberInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| phone_code | [string](#string) |  |  |
| gender | [int32](#int32) |  |  |
| avatar | [string](#string) |  |  |
| city_code | [string](#string) |  |  |
| code | [string](#string) |  |  |
| first_branch_id | [string](#string) |  |  |
| staff_id | [string](#string) |  |  |
| created_at | [int32](#int32) |  |  |
| channel | [string](#string) |  |  |
| birthday | [string](#string) |  |  |
| first_brand | [string](#string) |  |  |
| behavior_count | [uint32](#uint32) |  |  |






<a name="merchantBasic.MembersData"></a>

### MembersData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| members | [MemberInfo](#merchantBasic.MemberInfo) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="merchantBasic.ShowMemberByAccuratePhoneRequest"></a>

### ShowMemberByAccuratePhoneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone_code | [string](#string) |  |  |
| phone | [string](#string) |  |  |






<a name="merchantBasic.ShowMemberByAccuratePhoneResponse"></a>

### ShowMemberByAccuratePhoneResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [MemberInfo](#merchantBasic.MemberInfo) |  |  |






<a name="merchantBasic.ShowMemberRequest"></a>

### ShowMemberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.ShowMemberResponse"></a>

### ShowMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [MemberInfo](#merchantBasic.MemberInfo) |  |  |






<a name="merchantBasic.UpdateMemberBranchInfoRequest"></a>

### UpdateMemberBranchInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member_id | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |






<a name="merchantBasic.UpdateMemberRequest"></a>

### UpdateMemberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| phone_code | [string](#string) |  |  |
| gender | [int32](#int32) |  |  |
| avatar | [string](#string) |  |  |
| birthday | [string](#string) |  |  |






<a name="merchantBasic.UpdateMemberResponse"></a>

### UpdateMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/merchant_basic.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/merchant_basic.proto



<a name="merchantBasic.Empty"></a>

### Empty







<a name="merchantBasic.Response"></a>

### Response



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.UpdateStatusRequest"></a>

### UpdateStatusRequest
UpdateStatusRequest 统一定义修改状态request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |





 

 

 


<a name="merchantBasic.MerchantBasicService"></a>

### MerchantBasicService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBranch | [CreateBranchRequest](#merchantBasic.CreateBranchRequest) | [CreateBranchResponse](#merchantBasic.CreateBranchResponse) |  |
| UpdateBranch | [UpdateBranchRequest](#merchantBasic.UpdateBranchRequest) | [UpdateBranchResponse](#merchantBasic.UpdateBranchResponse) |  |
| GetBranches | [GetBranchesRequest](#merchantBasic.GetBranchesRequest) | [GetBranchesResponse](#merchantBasic.GetBranchesResponse) |  |
| DeleteBranch | [DeleteBranchRequest](#merchantBasic.DeleteBranchRequest) | [DeleteBranchResponse](#merchantBasic.DeleteBranchResponse) |  |
| ShowBranch | [ShowBranchRequest](#merchantBasic.ShowBranchRequest) | [ShowBranchResponse](#merchantBasic.ShowBranchResponse) |  |
| UpdateBranchStatus | [UpdateBranchStatusRequest](#merchantBasic.UpdateBranchStatusRequest) | [UpdateBranchStatusResponse](#merchantBasic.UpdateBranchStatusResponse) |  |
| UpdateBranchAccount | [UpdateBranchAccountRequest](#merchantBasic.UpdateBranchAccountRequest) | [UpdateBranchAccountResponse](#merchantBasic.UpdateBranchAccountResponse) |  |
| GetBranchesByTagIDs | [GetBranchesByTagIDsRequest](#merchantBasic.GetBranchesByTagIDsRequest) | [GetBranchesByTagIDsResponse](#merchantBasic.GetBranchesByTagIDsResponse) |  |
| CreateStaff | [CreateStaffRequest](#merchantBasic.CreateStaffRequest) | [CreateStaffResponse](#merchantBasic.CreateStaffResponse) |  |
| UpdateStaff | [UpdateStaffRequest](#merchantBasic.UpdateStaffRequest) | [UpdateStaffResponse](#merchantBasic.UpdateStaffResponse) |  |
| GetStaffs | [GetStaffsRequest](#merchantBasic.GetStaffsRequest) | [GetStaffsResponse](#merchantBasic.GetStaffsResponse) |  |
| DeleteStaff | [DeleteStaffRequest](#merchantBasic.DeleteStaffRequest) | [DeleteStaffResponse](#merchantBasic.DeleteStaffResponse) |  |
| ShowStaff | [ShowStaffRequest](#merchantBasic.ShowStaffRequest) | [ShowStaffResponse](#merchantBasic.ShowStaffResponse) |  |
| ShowStaffByPhone | [ShowStaffByPhoneRequest](#merchantBasic.ShowStaffByPhoneRequest) | [ShowStaffByPhoneResponse](#merchantBasic.ShowStaffByPhoneResponse) |  |
| UpdateStaffStatus | [UpdateStaffStatusRequest](#merchantBasic.UpdateStaffStatusRequest) | [UpdateStaffStatusResponse](#merchantBasic.UpdateStaffStatusResponse) |  |
| GetStaffsByRoleID | [GetStaffsByRoleIDRequest](#merchantBasic.GetStaffsByRoleIDRequest) | [GetStaffsByRoleIDResponse](#merchantBasic.GetStaffsByRoleIDResponse) |  |
| SignIn | [SignInRequest](#merchantBasic.SignInRequest) | [SignInResponse](#merchantBasic.SignInResponse) |  |
| UpdatePassword | [UpdatePasswordRequest](#merchantBasic.UpdatePasswordRequest) | [UpdatePasswordResponse](#merchantBasic.UpdatePasswordResponse) |  |
| ResetPassword | [ResetPasswordRequest](#merchantBasic.ResetPasswordRequest) | [ResetPasswordResponse](#merchantBasic.ResetPasswordResponse) |  |
| CreateRole | [CreateRoleRequest](#merchantBasic.CreateRoleRequest) | [CreateRoleResponse](#merchantBasic.CreateRoleResponse) |  |
| UpdateRole | [UpdateRoleRequest](#merchantBasic.UpdateRoleRequest) | [UpdateRoleResponse](#merchantBasic.UpdateRoleResponse) |  |
| GetRoles | [GetRolesRequest](#merchantBasic.GetRolesRequest) | [GetRolesResponse](#merchantBasic.GetRolesResponse) |  |
| DeleteRole | [DeleteRoleRequest](#merchantBasic.DeleteRoleRequest) | [DeleteRoleResponse](#merchantBasic.DeleteRoleResponse) |  |
| UpdateRoleStatus | [UpdateRoleStatusRequest](#merchantBasic.UpdateRoleStatusRequest) | [UpdateRoleStatusResponse](#merchantBasic.UpdateRoleStatusResponse) |  |
| GetRoleHistories | [GetRoleHistoriesRequest](#merchantBasic.GetRoleHistoriesRequest) | [GetRoleHistoriesResponse](#merchantBasic.GetRoleHistoriesResponse) |  |
| ShowRole | [ShowRoleRequest](#merchantBasic.ShowRoleRequest) | [ShowRoleResponse](#merchantBasic.ShowRoleResponse) |  |
| GetPermissions | [GetPermissionsRequest](#merchantBasic.GetPermissionsRequest) | [GetPermissionsResponse](#merchantBasic.GetPermissionsResponse) |  |
| CreatePermissions | [CreatePermissionsRequest](#merchantBasic.CreatePermissionsRequest) | [CreatePermissionsResponse](#merchantBasic.CreatePermissionsResponse) |  |
| CreateMember | [CreateMemberRequest](#merchantBasic.CreateMemberRequest) | [CreateMemberResponse](#merchantBasic.CreateMemberResponse) |  |
| GetMembers | [GetMembersRequest](#merchantBasic.GetMembersRequest) | [GetMembersResponse](#merchantBasic.GetMembersResponse) |  |
| GetMembersByIDs | [GetMembersByIDsRequest](#merchantBasic.GetMembersByIDsRequest) | [GetMembersByIDsResponse](#merchantBasic.GetMembersByIDsResponse) |  |
| ShowMember | [ShowMemberRequest](#merchantBasic.ShowMemberRequest) | [ShowMemberResponse](#merchantBasic.ShowMemberResponse) |  |
| ShowMemberByAccuratePhone | [ShowMemberByAccuratePhoneRequest](#merchantBasic.ShowMemberByAccuratePhoneRequest) | [ShowMemberByAccuratePhoneResponse](#merchantBasic.ShowMemberByAccuratePhoneResponse) |  |
| UpdateMember | [UpdateMemberRequest](#merchantBasic.UpdateMemberRequest) | [UpdateMemberResponse](#merchantBasic.UpdateMemberResponse) |  |
| GetBirthdayMembers | [GetBirthdayMembersRequest](#merchantBasic.GetBirthdayMembersRequest) | [GetBirthdayMembersResponse](#merchantBasic.GetBirthdayMembersResponse) |  |
| CreateOrUpdateCallingMember | [CreateOrUpdateCallingMemberRequest](#merchantBasic.CreateOrUpdateCallingMemberRequest) | [Response](#merchantBasic.Response) |  |
| GetMembersByPhoneSuffix | [GetMembersByPhoneSuffixRequest](#merchantBasic.GetMembersByPhoneSuffixRequest) | [GetMembersByPhoneSuffixResponse](#merchantBasic.GetMembersByPhoneSuffixResponse) |  |
| CreateWechatMember | [CreateWechatMemberRequest](#merchantBasic.CreateWechatMemberRequest) | [CreateWechatMemberResponse](#merchantBasic.CreateWechatMemberResponse) |  |
| UpdateMemberBranchInfo | [UpdateMemberBranchInfoRequest](#merchantBasic.UpdateMemberBranchInfoRequest) | [Response](#merchantBasic.Response) |  |
| CreateBranchTag | [CreateBranchTagRequest](#merchantBasic.CreateBranchTagRequest) | [CreateBranchTagResponse](#merchantBasic.CreateBranchTagResponse) |  |
| GetBranchTags | [GetBranchTagsRequest](#merchantBasic.GetBranchTagsRequest) | [GetBranchTagsResponse](#merchantBasic.GetBranchTagsResponse) |  |
| UpdateBranchTag | [UpdateBranchTagRequest](#merchantBasic.UpdateBranchTagRequest) | [UpdateBranchTagResponse](#merchantBasic.UpdateBranchTagResponse) |  |
| UpdateBranchTagStatus | [UpdateBranchTagStatusRequest](#merchantBasic.UpdateBranchTagStatusRequest) | [UpdateBranchTagStatusResponse](#merchantBasic.UpdateBranchTagStatusResponse) |  |
| GetBranchTagsByIDs | [GetBranchTagsByIDsRequest](#merchantBasic.GetBranchTagsByIDsRequest) | [GetBranchTagsByIDsResponse](#merchantBasic.GetBranchTagsByIDsResponse) |  |
| ShowBranchTag | [ShowBranchTagRequest](#merchantBasic.ShowBranchTagRequest) | [ShowBranchTagResponse](#merchantBasic.ShowBranchTagResponse) |  |
| SaveBranchAppointment | [SaveBranchAppointmentRequest](#merchantBasic.SaveBranchAppointmentRequest) | [SaveBranchAppointmentResponse](#merchantBasic.SaveBranchAppointmentResponse) |  |
| ShowBranchAppointment | [ShowBranchAppointmentRequest](#merchantBasic.ShowBranchAppointmentRequest) | [ShowBranchAppointmentResponse](#merchantBasic.ShowBranchAppointmentResponse) |  |
| UpdateBranchAppointmentRoomType | [UpdateBranchAppointmentRoomTypeRequest](#merchantBasic.UpdateBranchAppointmentRoomTypeRequest) | [UpdateBranchAppointmentRoomTypeResponse](#merchantBasic.UpdateBranchAppointmentRoomTypeResponse) |  |
| CreateRoomTypeCategory | [CreateRoomTypeCategoryRequest](#merchantBasic.CreateRoomTypeCategoryRequest) | [CreateRoomTypeCategoryResponse](#merchantBasic.CreateRoomTypeCategoryResponse) |  |
| CreateRoomType | [CreateRoomTypeRequest](#merchantBasic.CreateRoomTypeRequest) | [CreateRoomTypeResponse](#merchantBasic.CreateRoomTypeResponse) |  |
| GetRoomTypeCategories | [GetRoomTypeCategoriesRequest](#merchantBasic.GetRoomTypeCategoriesRequest) | [GetRoomTypeCategoriesResponse](#merchantBasic.GetRoomTypeCategoriesResponse) |  |
| GetRoomTypes | [GetRoomTypesRequest](#merchantBasic.GetRoomTypesRequest) | [GetRoomTypesResponse](#merchantBasic.GetRoomTypesResponse) |  |
| GetRoomTypesByIDs | [GetRoomTypesByIDsRequest](#merchantBasic.GetRoomTypesByIDsRequest) | [GetRoomTypesByIDsResponse](#merchantBasic.GetRoomTypesByIDsResponse) |  |
| UpdateRoomTypeCategory | [UpdateRoomTypeCategoryRequest](#merchantBasic.UpdateRoomTypeCategoryRequest) | [UpdateRoomTypeCategoryResponse](#merchantBasic.UpdateRoomTypeCategoryResponse) |  |
| UpdateRoomType | [UpdateRoomTypeRequest](#merchantBasic.UpdateRoomTypeRequest) | [UpdateRoomTypeResponse](#merchantBasic.UpdateRoomTypeResponse) |  |
| ShowBranchAppointmentBusinessTime | [ShowBranchAppointmentBusinessTimeRequest](#merchantBasic.ShowBranchAppointmentBusinessTimeRequest) | [ShowBranchAppointmentBusinessTimeResponse](#merchantBasic.ShowBranchAppointmentBusinessTimeResponse) |  |
| ShowRoomType | [ShowRoomTypeRequest](#merchantBasic.ShowRoomTypeRequest) | [ShowRoomTypeResponse](#merchantBasic.ShowRoomTypeResponse) |  |
| UpdateRoomTypeStatus | [UpdateRoomTypeStatusRequest](#merchantBasic.UpdateRoomTypeStatusRequest) | [Response](#merchantBasic.Response) |  |
| CreateInvoice | [CreateInvoiceRequest](#merchantBasic.CreateInvoiceRequest) | [CreateInvoiceResponse](#merchantBasic.CreateInvoiceResponse) |  |
| CreateDistrict | [CreateDistrictRequest](#merchantBasic.CreateDistrictRequest) | [CreateDistrictResponse](#merchantBasic.CreateDistrictResponse) |  |
| UpdateDistrict | [UpdateDistrictRequest](#merchantBasic.UpdateDistrictRequest) | [UpdateDistrictResponse](#merchantBasic.UpdateDistrictResponse) |  |
| GetDistricts | [GetDistrictsRequest](#merchantBasic.GetDistrictsRequest) | [GetDistrictsResponse](#merchantBasic.GetDistrictsResponse) |  |
| GetBranchWithSubMchID | [GetBranchWithSubMchIDRequest](#merchantBasic.GetBranchWithSubMchIDRequest) | [GetBranchWithSubMchIDResponse](#merchantBasic.GetBranchWithSubMchIDResponse) |  |
| SetBranchSubMchID | [SetBranchSubMchIDRequest](#merchantBasic.SetBranchSubMchIDRequest) | [Response](#merchantBasic.Response) |  |
| GetWechatPaySetting | [GetWechatPaySettingRequest](#merchantBasic.GetWechatPaySettingRequest) | [GetWechatPaySettingResponse](#merchantBasic.GetWechatPaySettingResponse) |  |
| GetWechatPaySettingByAppID | [GetWechatPaySettingByAppIDRequest](#merchantBasic.GetWechatPaySettingByAppIDRequest) | [GetWechatPaySettingResponse](#merchantBasic.GetWechatPaySettingResponse) |  |
| SetWechatPaySetting | [SetWechatPaySettingRequest](#merchantBasic.SetWechatPaySettingRequest) | [Response](#merchantBasic.Response) |  |
| GetBranchWechatPaymentSetting | [GetBranchWechatPaymentSettingRequest](#merchantBasic.GetBranchWechatPaymentSettingRequest) | [GetBranchWechatPaymentSettingResponse](#merchantBasic.GetBranchWechatPaymentSettingResponse) |  |
| CreateMemberBehavior | [CreateMemberBehaviorRequest](#merchantBasic.CreateMemberBehaviorRequest) | [Response](#merchantBasic.Response) |  |
| GetMemberBehaviors | [GetMemberBehaviorsRequest](#merchantBasic.GetMemberBehaviorsRequest) | [GetMemberBehaviorsResponse](#merchantBasic.GetMemberBehaviorsResponse) |  |
| GetBrands | [GetBrandsRequest](#merchantBasic.GetBrandsRequest) | [GetBrandsResponse](#merchantBasic.GetBrandsResponse) |  |
| CreateBrand | [CreateBrandRequest](#merchantBasic.CreateBrandRequest) | [Response](#merchantBasic.Response) |  |
| UpdateBrand | [UpdateBrandRequest](#merchantBasic.UpdateBrandRequest) | [Response](#merchantBasic.Response) |  |
| UpdateBrandStatus | [UpdateBrandStatusRequest](#merchantBasic.UpdateBrandStatusRequest) | [Response](#merchantBasic.Response) |  |
| CreateBranchBusinessSpecial | [CreateBranchBusinessSpecialRequest](#merchantBasic.CreateBranchBusinessSpecialRequest) | [Response](#merchantBasic.Response) |  |
| UpdateBranchBusinessSpecial | [UpdateBranchBusinessSpecialRequest](#merchantBasic.UpdateBranchBusinessSpecialRequest) | [Response](#merchantBasic.Response) |  |
| UpdateBranchBusinessNormal | [UpdateBranchBusinessNormalRequest](#merchantBasic.UpdateBranchBusinessNormalRequest) | [Response](#merchantBasic.Response) |  |
| UpdateBranchBusinessStatus | [UpdateBranchBusinessStatusRequest](#merchantBasic.UpdateBranchBusinessStatusRequest) | [Response](#merchantBasic.Response) |  |
| GetBranchBusinesses | [GetBranchBusinessesRequest](#merchantBasic.GetBranchBusinessesRequest) | [GetBranchBusinessesResponse](#merchantBasic.GetBranchBusinessesResponse) |  |
| GetWechatUser | [GetWechatUserRequest](#merchantBasic.GetWechatUserRequest) | [GetWechatUserResponse](#merchantBasic.GetWechatUserResponse) |  |
| CreateOrUpdateWechatUser | [CreateOrUpdateWechatUserRequest](#merchantBasic.CreateOrUpdateWechatUserRequest) | [CreateOrUpdateWechatUserResponse](#merchantBasic.CreateOrUpdateWechatUserResponse) |  |
| ShowWechatUser | [ShowWechatUserRequest](#merchantBasic.ShowWechatUserRequest) | [ShowWechatUserResponse](#merchantBasic.ShowWechatUserResponse) |  |
| GetBranchLatelyBusiness | [GetBranchLatelyBusinessRequest](#merchantBasic.GetBranchLatelyBusinessRequest) | [GetBranchLatelyBusinessResponse](#merchantBasic.GetBranchLatelyBusinessResponse) | 获取门店指定时间最近一次营业日 |
| ShowWechatUserByMember | [ShowWechatUserByMemberRequest](#merchantBasic.ShowWechatUserByMemberRequest) | [ShowWechatUserByMemberResponse](#merchantBasic.ShowWechatUserByMemberResponse) |  |
| CreatePointCategory | [CreatePointCategoryRequest](#merchantBasic.CreatePointCategoryRequest) | [Response](#merchantBasic.Response) | CreatePointCategory 创建积分类型 |
| ListPointCategory | [ListPointCategoryRequest](#merchantBasic.ListPointCategoryRequest) | [ListPointCategoryResponse](#merchantBasic.ListPointCategoryResponse) | ListPointCategory 积分类型列表 |
| UpdatePointCategory | [UpdatePointCategoryRequest](#merchantBasic.UpdatePointCategoryRequest) | [Response](#merchantBasic.Response) | UpdatePointCategory 修改积分类型 |
| UpdatePointCategoryStatus | [UpdateStatusRequest](#merchantBasic.UpdateStatusRequest) | [Response](#merchantBasic.Response) | UpdatePointCategoryStatus 修改积分类型状态 |
| ShowPointCategory | [ShowPointCategoryRequest](#merchantBasic.ShowPointCategoryRequest) | [ShowPointCategoryResponse](#merchantBasic.ShowPointCategoryResponse) | ShowPointCategory 积分类型详情 |
| SetPointRuleDescribe | [SetPointRuleDescribeRequest](#merchantBasic.SetPointRuleDescribeRequest) | [Response](#merchantBasic.Response) | SetPointRuleDescribe 设置积分规则说明 |
| GetPointRuleDescribe | [Empty](#merchantBasic.Empty) | [GetPointRuleDescribeResponse](#merchantBasic.GetPointRuleDescribeResponse) | GetPointRuleDescribe 获取积分规则说明 |
| CreatePointRule | [CreatePointRuleRequest](#merchantBasic.CreatePointRuleRequest) | [Response](#merchantBasic.Response) | CreatePointRule 创建积分规则 |
| UpdatePointRule | [UpdatePointRuleRequest](#merchantBasic.UpdatePointRuleRequest) | [Response](#merchantBasic.Response) | UpdatePointRule 修改积分规则 |
| UpdatePointRuleStatus | [UpdateStatusRequest](#merchantBasic.UpdateStatusRequest) | [Response](#merchantBasic.Response) | UpdatePointRuleStatus 修改积分规则状态 |
| ListPointRule | [ListPointRuleRequest](#merchantBasic.ListPointRuleRequest) | [ListPointRuleResponse](#merchantBasic.ListPointRuleResponse) | ListPointRuleRequest 积分规则列表 |
| ShowPointRule | [ShowPointRuleRequest](#merchantBasic.ShowPointRuleRequest) | [ShowPointRuleResponse](#merchantBasic.ShowPointRuleResponse) | ShowPointRule 积分规则详情 |
| GetPointRuleAllBranch | [Empty](#merchantBasic.Empty) | [GetPointRuleAllBranchResponse](#merchantBasic.GetPointRuleAllBranchResponse) | GetPointRuleAllBranch 获取所有已设置积分规则的门店 |
| GetBranchPointRule | [GetBranchPointRuleRequest](#merchantBasic.GetBranchPointRuleRequest) | [GetBranchPointRuleResponse](#merchantBasic.GetBranchPointRuleResponse) | GetBranchPointRule 获取门店积分规则 |
| CreateMemberAddress | [MemberAddress](#merchantBasic.MemberAddress) | [Response](#merchantBasic.Response) | CreateMemberAddress 新增用户地址 |
| UpdateMemberAddress | [MemberAddress](#merchantBasic.MemberAddress) | [Response](#merchantBasic.Response) | UpdateMemberAddress 修改用户地址 |
| GetMemberAddress | [GetMemberAddressRequest](#merchantBasic.GetMemberAddressRequest) | [GetMemberAddressResponse](#merchantBasic.GetMemberAddressResponse) | GetMemberAddress 获取用户地址 |
| SetMemberAddressDefault | [SetMemberAddressDefaultRequest](#merchantBasic.SetMemberAddressDefaultRequest) | [Response](#merchantBasic.Response) | SetMemberAddressDefault 标记为默认地址 |
| GetMemberDefaultAddress | [GetMemberDefaultAddressRequest](#merchantBasic.GetMemberDefaultAddressRequest) | [GetMemberDefaultAddressResponse](#merchantBasic.GetMemberDefaultAddressResponse) | GetMemberDefaultAddress 获取默认地址 |
| DeleteMemberAddress | [DeleteMemberAddressRequest](#merchantBasic.DeleteMemberAddressRequest) | [Response](#merchantBasic.Response) | DeleteMemberAddress 删除地址 |
| ShowGrowthConfig | [Empty](#merchantBasic.Empty) | [ShowGrowthConfigResponse](#merchantBasic.ShowGrowthConfigResponse) | 成长值 ShowGrowthConfig 查询 |
| SaveGrowthConfig | [SaveGrowthConfigRequest](#merchantBasic.SaveGrowthConfigRequest) | [Response](#merchantBasic.Response) | SaveGrowthConfig 保存 |
| CreateGrowthRule | [CreateGrowthRuleRequest](#merchantBasic.CreateGrowthRuleRequest) | [Response](#merchantBasic.Response) | CreateGrowthRule 新建 |
| GetGrowthRules | [GetGrowthRulesRequest](#merchantBasic.GetGrowthRulesRequest) | [GetGrowthRulesResponse](#merchantBasic.GetGrowthRulesResponse) | GetGrowthRules 列表 |
| ShowGrowthRule | [ShowGrowthRuleRequest](#merchantBasic.ShowGrowthRuleRequest) | [ShowGrowthRuleResponse](#merchantBasic.ShowGrowthRuleResponse) | ShowGrowthRule 详情 |
| UpdateGrowthRule | [UpdateGrowthRuleRequest](#merchantBasic.UpdateGrowthRuleRequest) | [Response](#merchantBasic.Response) | UpdateGrowthRule 更新 |
| GetBranchesHasGrowthRule | [Empty](#merchantBasic.Empty) | [GetBranchesHasGrowthRuleResponse](#merchantBasic.GetBranchesHasGrowthRuleResponse) | GetBranchesHasGrowthRule 已设置了规则的门店列表 |
| GetBranchGrowthRule | [GetBranchGrowthRuleRequest](#merchantBasic.GetBranchGrowthRuleRequest) | [GetBranchGrowthRuleResponse](#merchantBasic.GetBranchGrowthRuleResponse) | GetBranchGrowthRule 获取门店的成长值规则 |

 



<a name="proto/merchant-basic/payment.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/payment.proto



<a name="merchantBasic.BranchSubMchID"></a>

### BranchSubMchID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| branch_name | [string](#string) |  |  |
| sub_mch_id | [string](#string) |  |  |






<a name="merchantBasic.BranchSubMchIDData"></a>

### BranchSubMchIDData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branches | [BranchSubMchID](#merchantBasic.BranchSubMchID) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="merchantBasic.BranchWechatPaymentSetting"></a>

### BranchWechatPaymentSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| merchant_id | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| mch_id | [string](#string) |  |  |
| sub_mch_id | [string](#string) |  |  |
| private_key | [string](#string) |  |  |
| cert_filename | [string](#string) |  |  |
| cert_content | [bytes](#bytes) |  |  |
| headquarters_sub_mch_id | [string](#string) |  |  |






<a name="merchantBasic.GetBranchWechatPaymentSettingRequest"></a>

### GetBranchWechatPaymentSettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |






<a name="merchantBasic.GetBranchWechatPaymentSettingResponse"></a>

### GetBranchWechatPaymentSettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [BranchWechatPaymentSetting](#merchantBasic.BranchWechatPaymentSetting) |  |  |






<a name="merchantBasic.GetBranchWithSubMchIDRequest"></a>

### GetBranchWithSubMchIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| city_id | [string](#string) |  |  |
| branch_name | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |
| province_id | [string](#string) |  |  |
| district_id | [string](#string) |  |  |






<a name="merchantBasic.GetBranchWithSubMchIDResponse"></a>

### GetBranchWithSubMchIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [BranchSubMchIDData](#merchantBasic.BranchSubMchIDData) |  |  |






<a name="merchantBasic.GetWechatPaySettingByAppIDRequest"></a>

### GetWechatPaySettingByAppIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| app_id | [string](#string) |  |  |






<a name="merchantBasic.GetWechatPaySettingRequest"></a>

### GetWechatPaySettingRequest







<a name="merchantBasic.GetWechatPaySettingResponse"></a>

### GetWechatPaySettingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [WechatPaySetting](#merchantBasic.WechatPaySetting) |  |  |






<a name="merchantBasic.SetBranchSubMchIDRequest"></a>

### SetBranchSubMchIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| sub_mch_id | [string](#string) |  |  |






<a name="merchantBasic.SetWechatPaySettingRequest"></a>

### SetWechatPaySettingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| app_id | [string](#string) |  |  |
| mch_id | [string](#string) |  |  |
| private_key | [string](#string) |  |  |
| cert_filename | [string](#string) |  |  |
| cert_content | [bytes](#bytes) |  |  |
| headquarters_sub_mch_id | [string](#string) |  |  |






<a name="merchantBasic.WechatPaySetting"></a>

### WechatPaySetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| mch_id | [string](#string) |  |  |
| private_key | [string](#string) |  |  |
| cert_filename | [string](#string) |  |  |
| cert_content | [bytes](#bytes) |  |  |
| headquarters_sub_mch_id | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/permission.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/permission.proto



<a name="merchantBasic.CreatePermissionsRequest"></a>

### CreatePermissionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permissions | [PermissionInfo](#merchantBasic.PermissionInfo) | repeated |  |
| service | [string](#string) |  |  |






<a name="merchantBasic.CreatePermissionsResponse"></a>

### CreatePermissionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.GetPermissionsRequest"></a>

### GetPermissionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| service | [string](#string) |  |  |






<a name="merchantBasic.GetPermissionsResponse"></a>

### GetPermissionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [PermissionInfo](#merchantBasic.PermissionInfo) | repeated |  |






<a name="merchantBasic.PermissionInfo"></a>

### PermissionInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| permission | [string](#string) |  |  |
| service | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/point.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/point.proto



<a name="merchantBasic.CreatePointCategoryRequest"></a>

### CreatePointCategoryRequest
CreatePointCategoryRequest 创建积分类型


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [string](#string) |  |  |
| code | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.CreatePointRuleRequest"></a>

### CreatePointRuleRequest
CreatePointRule 保存积分设置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rule_name | [string](#string) |  |  |
| gain_rules | [PointRuleDetail](#merchantBasic.PointRuleDetail) | repeated |  |
| use_rules | [PointRuleDetail](#merchantBasic.PointRuleDetail) | repeated |  |
| validity_day | [int32](#int32) |  |  |
| branch_ids | [string](#string) | repeated |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.GetBranchPointRuleRequest"></a>

### GetBranchPointRuleRequest
GetBranchPointRuleRequest 获取门店积分规则


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| category_code | [string](#string) |  |  |






<a name="merchantBasic.GetBranchPointRuleResponse"></a>

### GetBranchPointRuleResponse
GetBranchPointRuleResponse 获取门店积分规则


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GetBranchPointRuleResponse.Data](#merchantBasic.GetBranchPointRuleResponse.Data) |  |  |






<a name="merchantBasic.GetBranchPointRuleResponse.Data"></a>

### GetBranchPointRuleResponse.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gain_rule | [GetBranchPointRuleResponse.Rule](#merchantBasic.GetBranchPointRuleResponse.Rule) |  |  |
| use_rule | [GetBranchPointRuleResponse.Rule](#merchantBasic.GetBranchPointRuleResponse.Rule) |  |  |






<a name="merchantBasic.GetBranchPointRuleResponse.Rule"></a>

### GetBranchPointRuleResponse.Rule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| point | [int64](#int64) |  |  |
| fee | [int64](#int64) |  |  |
| validity_day | [int32](#int32) |  |  |






<a name="merchantBasic.GetPointRuleAllBranchResponse"></a>

### GetPointRuleAllBranchResponse
GetPointRuleAllBranchResponse 获取所有已设置积分规则的门店


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GetPointRuleAllBranchResponse.Data](#merchantBasic.GetPointRuleAllBranchResponse.Data) |  |  |






<a name="merchantBasic.GetPointRuleAllBranchResponse.Data"></a>

### GetPointRuleAllBranchResponse.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_ids | [string](#string) | repeated |  |






<a name="merchantBasic.GetPointRuleDescribeResponse"></a>

### GetPointRuleDescribeResponse
GetPointRuleDescribeResponse 获取积分规则说明


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [GetPointRuleDescribeResponse.Data](#merchantBasic.GetPointRuleDescribeResponse.Data) |  |  |






<a name="merchantBasic.GetPointRuleDescribeResponse.Data"></a>

### GetPointRuleDescribeResponse.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| images | [string](#string) | repeated |  |






<a name="merchantBasic.ListPointCategoryRequest"></a>

### ListPointCategoryRequest
ListPointCategoryRequest 积分类型列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [string](#string) |  |  |
| status | [string](#string) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |
| order_by | [string](#string) |  |  |
| with_page | [bool](#bool) |  |  |






<a name="merchantBasic.ListPointCategoryResponse"></a>

### ListPointCategoryResponse
ListPointCategoryResponse 积分类型列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ListPointCategoryResponse.Data](#merchantBasic.ListPointCategoryResponse.Data) |  |  |






<a name="merchantBasic.ListPointCategoryResponse.Data"></a>

### ListPointCategoryResponse.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [PointCategory](#merchantBasic.PointCategory) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="merchantBasic.ListPointRuleRequest"></a>

### ListPointRuleRequest
ListPointRuleRequest 查看积分规则列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rule_name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |
| order_by | [string](#string) |  |  |
| with_page | [bool](#bool) |  |  |






<a name="merchantBasic.ListPointRuleResponse"></a>

### ListPointRuleResponse
ListPointRuleRequest 查看积分规则列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ListPointRuleResponse.Data](#merchantBasic.ListPointRuleResponse.Data) |  |  |






<a name="merchantBasic.ListPointRuleResponse.Data"></a>

### ListPointRuleResponse.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [PointRule](#merchantBasic.PointRule) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="merchantBasic.PointCategory"></a>

### PointCategory
PointCategory 积分类型


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| category | [string](#string) |  |  |
| code | [string](#string) |  |  |
| status | [string](#string) |  |  |
| created_at | [int64](#int64) |  |  |






<a name="merchantBasic.PointRule"></a>

### PointRule
PointRule 积分规则


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| rule_name | [string](#string) |  |  |
| gain_rules | [PointRuleDetail](#merchantBasic.PointRuleDetail) | repeated |  |
| use_rules | [PointRuleDetail](#merchantBasic.PointRuleDetail) | repeated |  |
| validity_day | [int32](#int32) |  |  |
| branch_ids | [string](#string) | repeated |  |
| status | [string](#string) |  |  |
| created_at | [int64](#int64) |  |  |






<a name="merchantBasic.PointRuleDetail"></a>

### PointRuleDetail
PointRuleDetail 积分规则详情


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category_id | [string](#string) |  |  |
| point | [int64](#int64) |  |  |
| fee | [int64](#int64) |  |  |






<a name="merchantBasic.SetPointRuleDescribeRequest"></a>

### SetPointRuleDescribeRequest
SetPointRuleDescribeRequest 设置积分规则说明


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| images | [string](#string) | repeated |  |






<a name="merchantBasic.ShowPointCategoryRequest"></a>

### ShowPointCategoryRequest
ShowPointCategoryRequest 查看积分类型


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.ShowPointCategoryResponse"></a>

### ShowPointCategoryResponse
ShowPointCategoryResponse 查看积分类型


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [PointCategory](#merchantBasic.PointCategory) |  |  |






<a name="merchantBasic.ShowPointRuleRequest"></a>

### ShowPointRuleRequest
ShowPointRuleRequest 查看积分规则详情


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |






<a name="merchantBasic.ShowPointRuleResponse"></a>

### ShowPointRuleResponse
ShowPointRuleResponse 查看积分规则详情


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [PointRule](#merchantBasic.PointRule) |  |  |






<a name="merchantBasic.UpdatePointCategoryRequest"></a>

### UpdatePointCategoryRequest
UpdatePointCategoryRequest 修改积分类型


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| category | [string](#string) |  |  |
| code | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.UpdatePointRuleRequest"></a>

### UpdatePointRuleRequest
UpdatePointRuleRequest 修改积分规则


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| rule_name | [string](#string) |  |  |
| gain_rules | [PointRuleDetail](#merchantBasic.PointRuleDetail) | repeated |  |
| use_rules | [PointRuleDetail](#merchantBasic.PointRuleDetail) | repeated |  |
| validity_day | [int32](#int32) |  |  |
| branch_ids | [string](#string) | repeated |  |
| status | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/role.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/role.proto



<a name="merchantBasic.CreateRoleRequest"></a>

### CreateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| property | [int32](#int32) |  |  |
| permissions | [int32](#int32) | repeated |  |






<a name="merchantBasic.CreateRoleResponse"></a>

### CreateRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.DeleteRoleRequest"></a>

### DeleteRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.DeleteRoleResponse"></a>

### DeleteRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.GetRoleHistoriesRequest"></a>

### GetRoleHistoriesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.GetRoleHistoriesResponse"></a>

### GetRoleHistoriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [RoleHistoriesData](#merchantBasic.RoleHistoriesData) |  |  |






<a name="merchantBasic.GetRolesRequest"></a>

### GetRolesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |
| staff_id | [string](#string) |  |  |






<a name="merchantBasic.GetRolesResponse"></a>

### GetRolesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [RolesData](#merchantBasic.RolesData) |  |  |






<a name="merchantBasic.RoleHistoriesData"></a>

### RoleHistoriesData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [RoleInfo](#merchantBasic.RoleInfo) |  |  |
| snapshots | [Snapshot](#merchantBasic.Snapshot) | repeated |  |






<a name="merchantBasic.RoleInfo"></a>

### RoleInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| property | [int32](#int32) |  |  |
| staff_id | [string](#string) |  |  |
| staff_name | [string](#string) |  |  |






<a name="merchantBasic.RolePermissionInfo"></a>

### RolePermissionInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [RoleInfo](#merchantBasic.RoleInfo) |  |  |
| permissions | [int32](#int32) | repeated |  |






<a name="merchantBasic.RolesData"></a>

### RolesData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| roles | [RoleInfo](#merchantBasic.RoleInfo) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="merchantBasic.ShowRoleRequest"></a>

### ShowRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.ShowRoleResponse"></a>

### ShowRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [RolePermissionInfo](#merchantBasic.RolePermissionInfo) |  |  |






<a name="merchantBasic.Snapshot"></a>

### Snapshot



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| staff_id | [string](#string) |  |  |
| staff_name | [string](#string) |  |  |
| before | [string](#string) |  |  |
| after | [string](#string) |  |  |
| created_at | [int32](#int32) |  |  |
| method | [string](#string) |  |  |






<a name="merchantBasic.UpdateRoleRequest"></a>

### UpdateRoleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| property | [int32](#int32) |  |  |
| permissions | [int32](#int32) | repeated |  |






<a name="merchantBasic.UpdateRoleResponse"></a>

### UpdateRoleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.UpdateRoleStatusRequest"></a>

### UpdateRoleStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.UpdateRoleStatusResponse"></a>

### UpdateRoleStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/room_type_category.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/room_type_category.proto



<a name="merchantBasic.CreateRoomTypeCategoryRequest"></a>

### CreateRoomTypeCategoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| category | [int32](#int32) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.CreateRoomTypeCategoryResponse"></a>

### CreateRoomTypeCategoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.GetRoomTypeCategoriesRequest"></a>

### GetRoomTypeCategoriesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |
| category | [int32](#int32) |  |  |






<a name="merchantBasic.GetRoomTypeCategoriesResponse"></a>

### GetRoomTypeCategoriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [RoomTypeCategoriesData](#merchantBasic.RoomTypeCategoriesData) |  |  |






<a name="merchantBasic.RoomTypeCategoriesData"></a>

### RoomTypeCategoriesData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_type_categories | [RoomTypeCategory](#merchantBasic.RoomTypeCategory) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="merchantBasic.RoomTypeCategory"></a>

### RoomTypeCategory



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| category | [int32](#int32) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.UpdateRoomTypeCategoryRequest"></a>

### UpdateRoomTypeCategoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| category | [int32](#int32) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.UpdateRoomTypeCategoryResponse"></a>

### UpdateRoomTypeCategoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/room_type.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/room_type.proto



<a name="merchantBasic.CreateRoomTypeRequest"></a>

### CreateRoomTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| category_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| customer_min | [int32](#int32) |  |  |
| customer_max | [int32](#int32) |  |  |
| order | [int32](#int32) |  |  |






<a name="merchantBasic.CreateRoomTypeResponse"></a>

### CreateRoomTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.GetRoomTypesByIDsRequest"></a>

### GetRoomTypesByIDsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ids | [string](#string) | repeated |  |






<a name="merchantBasic.GetRoomTypesByIDsResponse"></a>

### GetRoomTypesByIDsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [RoomType](#merchantBasic.RoomType) | repeated |  |






<a name="merchantBasic.GetRoomTypesRequest"></a>

### GetRoomTypesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |
| category_id | [string](#string) |  |  |






<a name="merchantBasic.GetRoomTypesResponse"></a>

### GetRoomTypesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [RoomTypesData](#merchantBasic.RoomTypesData) |  |  |






<a name="merchantBasic.RoomType"></a>

### RoomType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| category_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| category_name | [string](#string) |  |  |
| customer_min | [int32](#int32) |  |  |
| customer_max | [int32](#int32) |  |  |
| order | [int32](#int32) |  |  |






<a name="merchantBasic.RoomTypesData"></a>

### RoomTypesData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_types | [RoomType](#merchantBasic.RoomType) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="merchantBasic.ShowRoomTypeRequest"></a>

### ShowRoomTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.ShowRoomTypeResponse"></a>

### ShowRoomTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [RoomType](#merchantBasic.RoomType) |  |  |






<a name="merchantBasic.UpdateRoomTypeRequest"></a>

### UpdateRoomTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| category_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| status | [string](#string) |  |  |
| customer_min | [int32](#int32) |  |  |
| customer_max | [int32](#int32) |  |  |
| order | [int32](#int32) |  |  |






<a name="merchantBasic.UpdateRoomTypeResponse"></a>

### UpdateRoomTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.UpdateRoomTypeStatusRequest"></a>

### UpdateRoomTypeStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/staff.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/staff.proto



<a name="merchantBasic.CreateStaffRequest"></a>

### CreateStaffRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| staff | [StaffInfo](#merchantBasic.StaffInfo) |  |  |
| roles | [string](#string) | repeated |  |
| branches | [string](#string) | repeated |  |






<a name="merchantBasic.CreateStaffResponse"></a>

### CreateStaffResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.DeleteStaffRequest"></a>

### DeleteStaffRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.DeleteStaffResponse"></a>

### DeleteStaffResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.GetStaffsByRoleIDRequest"></a>

### GetStaffsByRoleIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role_id | [string](#string) |  |  |






<a name="merchantBasic.GetStaffsByRoleIDResponse"></a>

### GetStaffsByRoleIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [StaffInfo](#merchantBasic.StaffInfo) | repeated |  |






<a name="merchantBasic.GetStaffsRequest"></a>

### GetStaffsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| gender | [int32](#int32) |  |  |
| date_start | [int32](#int32) |  |  |
| date_end | [int32](#int32) |  |  |
| branches | [string](#string) | repeated |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.GetStaffsResponse"></a>

### GetStaffsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [StaffsData](#merchantBasic.StaffsData) |  |  |






<a name="merchantBasic.ResetPasswordRequest"></a>

### ResetPasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.ResetPasswordResponse"></a>

### ResetPasswordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.ShowStaffByPhoneRequest"></a>

### ShowStaffByPhoneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  |  |
| phone_code | [string](#string) |  |  |






<a name="merchantBasic.ShowStaffByPhoneResponse"></a>

### ShowStaffByPhoneResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [StaffFullInfo](#merchantBasic.StaffFullInfo) |  |  |






<a name="merchantBasic.ShowStaffRequest"></a>

### ShowStaffRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.ShowStaffResponse"></a>

### ShowStaffResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [StaffFullInfo](#merchantBasic.StaffFullInfo) |  |  |






<a name="merchantBasic.SignInData"></a>

### SignInData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| staff_full | [StaffFullInfo](#merchantBasic.StaffFullInfo) |  |  |
| default_password | [bool](#bool) |  |  |






<a name="merchantBasic.SignInRequest"></a>

### SignInRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="merchantBasic.SignInResponse"></a>

### SignInResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [SignInData](#merchantBasic.SignInData) |  |  |






<a name="merchantBasic.StaffFullInfo"></a>

### StaffFullInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| staff | [StaffInfo](#merchantBasic.StaffInfo) |  |  |
| roles | [RoleInfo](#merchantBasic.RoleInfo) | repeated |  |
| branches | [BranchInfo](#merchantBasic.BranchInfo) | repeated |  |






<a name="merchantBasic.StaffInfo"></a>

### StaffInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| phone_code | [string](#string) |  |  |
| gender | [int32](#int32) |  |  |
| status | [string](#string) |  |  |
| code | [string](#string) |  |  |
| entry_at | [int32](#int32) |  |  |
| quit_at | [int32](#int32) |  |  |
| created_at | [int32](#int32) |  |  |
| id | [string](#string) |  |  |
| employee_code | [string](#string) |  |  |
| admin | [bool](#bool) |  |  |






<a name="merchantBasic.StaffsData"></a>

### StaffsData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| staffs | [StaffFullInfo](#merchantBasic.StaffFullInfo) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="merchantBasic.UpdatePasswordRequest"></a>

### UpdatePasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| old_password | [string](#string) |  |  |
| new_password | [string](#string) |  |  |






<a name="merchantBasic.UpdatePasswordResponse"></a>

### UpdatePasswordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.UpdateStaffRequest"></a>

### UpdateStaffRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| staff | [StaffInfo](#merchantBasic.StaffInfo) |  |  |
| roles | [string](#string) | repeated |  |
| branches | [string](#string) | repeated |  |






<a name="merchantBasic.UpdateStaffResponse"></a>

### UpdateStaffResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="merchantBasic.UpdateStaffStatusRequest"></a>

### UpdateStaffStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| status | [string](#string) |  |  |






<a name="merchantBasic.UpdateStaffStatusResponse"></a>

### UpdateStaffStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |





 

 

 

 



<a name="proto/merchant-basic/wechat_member.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/merchant-basic/wechat_member.proto



<a name="merchantBasic.CreateOrUpdateWechatUserRequest"></a>

### CreateOrUpdateWechatUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [WechatUser](#merchantBasic.WechatUser) |  |  |






<a name="merchantBasic.CreateOrUpdateWechatUserResponse"></a>

### CreateOrUpdateWechatUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [WechatUser](#merchantBasic.WechatUser) |  |  |






<a name="merchantBasic.GetWechatUserRequest"></a>

### GetWechatUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| openid | [string](#string) |  |  |
| appid | [string](#string) |  |  |






<a name="merchantBasic.GetWechatUserResponse"></a>

### GetWechatUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [WechatUser](#merchantBasic.WechatUser) |  |  |






<a name="merchantBasic.ShowWechatUserByMemberRequest"></a>

### ShowWechatUserByMemberRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member_id | [string](#string) |  |  |






<a name="merchantBasic.ShowWechatUserByMemberResponse"></a>

### ShowWechatUserByMemberResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [WechatUser](#merchantBasic.WechatUser) |  |  |






<a name="merchantBasic.ShowWechatUserRequest"></a>

### ShowWechatUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="merchantBasic.ShowWechatUserResponse"></a>

### ShowWechatUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [WechatUser](#merchantBasic.WechatUser) |  |  |






<a name="merchantBasic.WechatUser"></a>

### WechatUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| openid | [string](#string) |  |  |
| appid | [string](#string) |  |  |
| member_id | [string](#string) |  |  |
| nickname | [string](#string) |  |  |
| sex | [int32](#int32) |  |  |
| province | [string](#string) |  |  |
| city | [string](#string) |  |  |
| headimgurl | [string](#string) |  |  |
| unionid | [string](#string) |  |  |
| created_at | [string](#string) |  |  |
| member_phone | [string](#string) |  |  |
| member_phone_code | [string](#string) |  |  |
| member_name | [string](#string) |  |  |





 

 

 

 



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

