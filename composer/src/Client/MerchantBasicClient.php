<?php

declare(strict_types=1);

namespace Omy\MicroServers\Client;

use Omy\Crius\BasicServer\EmptyRequest;
use Omy\Crius\MemberAccountServer\NoDataResponse;
use Omy\Crius\MerchantBasicServer\CreateBranchBusinessSpecialRequest;
use Omy\Crius\MerchantBasicServer\CreateBranchRequest;
use Omy\Crius\MerchantBasicServer\CreateBranchResponse;
use Omy\Crius\MerchantBasicServer\CreateBranchTagRequest;
use Omy\Crius\MerchantBasicServer\CreateBranchTagResponse;
use Omy\Crius\MerchantBasicServer\CreateBrandRequest;
use Omy\Crius\MerchantBasicServer\CreateConsumeCategoryRequest;
use Omy\Crius\MerchantBasicServer\CreateDistrictRequest;
use Omy\Crius\MerchantBasicServer\CreateGrowthRuleRequest;
use Omy\Crius\MerchantBasicServer\CreateInvoiceRequest;
use Omy\Crius\MerchantBasicServer\CreateInvoiceResponse;
use Omy\Crius\MerchantBasicServer\CreateMemberBehaviorRequest;
use Omy\Crius\MerchantBasicServer\CreateMemberRequest;
use Omy\Crius\MerchantBasicServer\CreateMemberResponse;
use Omy\Crius\MerchantBasicServer\CreateOrUpdateCallingMemberRequest;
use Omy\Crius\MerchantBasicServer\CreateOrUpdateCallingMemberResponse;
use Omy\Crius\MerchantBasicServer\CreateOrUpdateWechatUserRequest;
use Omy\Crius\MerchantBasicServer\CreateOrUpdateWechatUserResponse;
use Omy\Crius\MerchantBasicServer\CreatePermissionsRequest;
use Omy\Crius\MerchantBasicServer\CreatePermissionsResponse;
use Omy\Crius\MerchantBasicServer\CreateRoleRequest;
use Omy\Crius\MerchantBasicServer\CreateRoleResponse;
use Omy\Crius\MerchantBasicServer\CreateRoomTypeCategoryRequest;
use Omy\Crius\MerchantBasicServer\CreateRoomTypeCategoryResponse;
use Omy\Crius\MerchantBasicServer\CreateRoomTypeRequest;
use Omy\Crius\MerchantBasicServer\CreateStaffRequest;
use Omy\Crius\MerchantBasicServer\CreateStaffResponse;
use Omy\Crius\MerchantBasicServer\CreateWechatMemberRequest;
use Omy\Crius\MerchantBasicServer\CreateWechatMemberResponse;
use Omy\Crius\MerchantBasicServer\DeleteBranchRequest;
use Omy\Crius\MerchantBasicServer\DeleteBranchResponse;
use Omy\Crius\MerchantBasicServer\DeleteMemberAddressRequest;
use Omy\Crius\MerchantBasicServer\DeleteRoleRequest;
use Omy\Crius\MerchantBasicServer\DeleteRoleResponse;
use Omy\Crius\MerchantBasicServer\DeleteStaffRequest;
use Omy\Crius\MerchantBasicServer\DeleteStaffResponse;
use Omy\Crius\MerchantBasicServer\GetActiveTypesResponse;
use Omy\Crius\MerchantBasicServer\GetBranchBusinessesRequest;
use Omy\Crius\MerchantBasicServer\GetBranchBusinessesResponse;
use Omy\Crius\MerchantBasicServer\GetBranchesByTagIDsRequest;
use Omy\Crius\MerchantBasicServer\GetBranchesByTagIDsResponse;
use Omy\Crius\MerchantBasicServer\GetBranchesHasGrowthRuleResponse;
use Omy\Crius\MerchantBasicServer\GetBranchesRequest;
use Omy\Crius\MerchantBasicServer\GetBranchesResponse;
use Omy\Crius\MerchantBasicServer\GetBranchGrowthRuleRequest;
use Omy\Crius\MerchantBasicServer\GetBranchGrowthRuleResponse;
use Omy\Crius\MerchantBasicServer\GetBranchPointRuleRequest;
use Omy\Crius\MerchantBasicServer\GetBranchPointRuleResponse;
use Omy\Crius\MerchantBasicServer\GetBranchPointRulesRequest;
use Omy\Crius\MerchantBasicServer\GetBranchPointRulesResponse;
use Omy\Crius\MerchantBasicServer\GetBranchTagsByIDsRequest;
use Omy\Crius\MerchantBasicServer\GetBranchTagsByIDsResponse;
use Omy\Crius\MerchantBasicServer\GetBranchTagsRequest;
use Omy\Crius\MerchantBasicServer\GetBranchTagsResponse;
use Omy\Crius\MerchantBasicServer\GetBranchWithSubMchIDRequest;
use Omy\Crius\MerchantBasicServer\GetBranchWithSubMchIDResponse;
use Omy\Crius\MerchantBasicServer\GetBrandsRequest;
use Omy\Crius\MerchantBasicServer\GetBrandsResponse;
use Omy\Crius\MerchantBasicServer\GetConsumeCategoryTypesResponse;
use Omy\Crius\MerchantBasicServer\GetDistrictsRequest;
use Omy\Crius\MerchantBasicServer\GetDistrictsResponse;
use Omy\Crius\MerchantBasicServer\GetGrowthRulesRequest;
use Omy\Crius\MerchantBasicServer\GetGrowthRulesResponse;
use Omy\Crius\MerchantBasicServer\GetMemberBehaviorsRequest;
use Omy\Crius\MerchantBasicServer\GetMemberBehaviorsResponse;
use Omy\Crius\MerchantBasicServer\GetMemberDefaultAddressRequest;
use Omy\Crius\MerchantBasicServer\GetMemberDefaultAddressResponse;
use Omy\Crius\MerchantBasicServer\GetMembersByIDsRequest;
use Omy\Crius\MerchantBasicServer\GetMembersByIDsResponse;
use Omy\Crius\MerchantBasicServer\GetMembersRequest;
use Omy\Crius\MerchantBasicServer\GetMembersResponse;
use Omy\Crius\MerchantBasicServer\GetOperateTypesResponse;
use Omy\Crius\MerchantBasicServer\GetPermissionsRequest;
use Omy\Crius\MerchantBasicServer\GetPermissionsResponse;
use Omy\Crius\MerchantBasicServer\GetRoleHistoriesRequest;
use Omy\Crius\MerchantBasicServer\GetRoleHistoriesResponse;
use Omy\Crius\MerchantBasicServer\GetRolesRequest;
use Omy\Crius\MerchantBasicServer\GetRolesResponse;
use Omy\Crius\MerchantBasicServer\GetRoomTypeCategoriesRequest;
use Omy\Crius\MerchantBasicServer\GetRoomTypeCategoriesResponse;
use Omy\Crius\MerchantBasicServer\GetRoomTypeGroupsResponse;
use Omy\Crius\MerchantBasicServer\GetRoomTypesRequest;
use Omy\Crius\MerchantBasicServer\GetRoomTypesResponse;
use Omy\Crius\MerchantBasicServer\GetStaffsByRoleIDRequest;
use Omy\Crius\MerchantBasicServer\GetStaffsByRoleIDResponse;
use Omy\Crius\MerchantBasicServer\GetStaffsRequest;
use Omy\Crius\MerchantBasicServer\GetStaffsResponse;
use Omy\Crius\MerchantBasicServer\GetWechatPaySettingRequest;
use Omy\Crius\MerchantBasicServer\GetWechatPaySettingResponse;
use Omy\Crius\MerchantBasicServer\GetWechatUserByUnionIdRequest;
use Omy\Crius\MerchantBasicServer\GetWechatUserByUnionIdResponse;
use Omy\Crius\MerchantBasicServer\GetWechatUsersRequest;
use Omy\Crius\MerchantBasicServer\GetWechatUsersResponse;
use Omy\Crius\MerchantBasicServer\ListConsumeCategoryRequest;
use Omy\Crius\MerchantBasicServer\ListConsumeCategoryResponse;
use Omy\Crius\MerchantBasicServer\MemberAddress;
use Omy\Crius\MerchantBasicServer\MultiGetGoodsAndPackagesRequest;
use Omy\Crius\MerchantBasicServer\MultiGetGoodsAndPackagesResponse;
use Omy\Crius\MerchantBasicServer\SaveBranchAppointmentRequest;
use Omy\Crius\MerchantBasicServer\SaveBranchAppointmentResponse;
use Omy\Crius\MerchantBasicServer\SaveGrowthConfigRequest;
use Omy\Crius\MerchantBasicServer\SearchGoodsAndPackageRequest;
use Omy\Crius\MerchantBasicServer\SearchGoodsAndPackageResponse;
use Omy\Crius\MerchantBasicServer\SearchGoodsOrPackageRequest;
use Omy\Crius\MerchantBasicServer\SearchGoodsOrPackageResponse;
use Omy\Crius\MerchantBasicServer\SearchMemberRequest;
use Omy\Crius\MerchantBasicServer\SearchMemberResponse;
use Omy\Crius\MerchantBasicServer\SetBranchSubMchIDRequest;
use Omy\Crius\MerchantBasicServer\SetMemberAddressDefaultRequest;
use Omy\Crius\MerchantBasicServer\SetWechatPaySettingRequest;
use Omy\Crius\MerchantBasicServer\ShowBranchAppointmentBusinessTimeRequest;
use Omy\Crius\MerchantBasicServer\ShowBranchAppointmentBusinessTimeResponse;
use Omy\Crius\MerchantBasicServer\GetWechatUserRequest;
use Omy\Crius\MerchantBasicServer\GetWechatUserResponse;
use Omy\Crius\MerchantBasicServer\ShowBranchAppointmentRequest;
use Omy\Crius\MerchantBasicServer\ShowBranchAppointmentResponse;
use Omy\Crius\MerchantBasicServer\ShowBranchRequest;
use Omy\Crius\MerchantBasicServer\ShowBranchResponse;
use Omy\Crius\MerchantBasicServer\ShowBranchTagRequest;
use Omy\Crius\MerchantBasicServer\ShowBranchTagResponse;
use Omy\Crius\MerchantBasicServer\ShowConsumeCategoryRequest;
use Omy\Crius\MerchantBasicServer\ShowConsumeCategoryResponse;
use Omy\Crius\MerchantBasicServer\ShowGoodsRequest;
use Omy\Crius\MerchantBasicServer\ShowGoodsResponse;
use Omy\Crius\MerchantBasicServer\ShowGrowthConfigResponse;
use Omy\Crius\MerchantBasicServer\ShowGrowthRuleRequest;
use Omy\Crius\MerchantBasicServer\ShowGrowthRuleResponse;
use Omy\Crius\MerchantBasicServer\ShowMemberByAccuratePhoneRequest;
use Omy\Crius\MerchantBasicServer\ShowMemberByAccuratePhoneResponse;
use Omy\Crius\MerchantBasicServer\ShowMemberRequest;
use Omy\Crius\MerchantBasicServer\ShowMemberResponse;
use Omy\Crius\MerchantBasicServer\ShowMerchantResponse;
use Omy\Crius\MerchantBasicServer\ShowPackageRequest;
use Omy\Crius\MerchantBasicServer\ShowPackageResponse;
use Omy\Crius\MerchantBasicServer\ShowRoleRequest;
use Omy\Crius\MerchantBasicServer\ShowRoleResponse;
use Omy\Crius\MerchantBasicServer\ShowStaffByPhoneRequest;
use Omy\Crius\MerchantBasicServer\ShowStaffByPhoneResponse;
use Omy\Crius\MerchantBasicServer\ShowStaffRequest;
use Omy\Crius\MerchantBasicServer\ShowStaffResponse;
use Omy\Crius\MerchantBasicServer\ShowWechatUserByMemberRequest;
use Omy\Crius\MerchantBasicServer\ShowWechatUserByMemberResponse;
use Omy\Crius\MerchantBasicServer\ShowWechatUserRequest;
use Omy\Crius\MerchantBasicServer\ShowWechatUserResponse;
use Omy\Crius\MerchantBasicServer\SignInRequest;
use Omy\Crius\MerchantBasicServer\SignInResponse;
use Omy\Crius\MerchantBasicServer\UpdateBranchAppointmentRoomTypeRequest;
use Omy\Crius\MerchantBasicServer\UpdateBranchAppointmentRoomTypeResponse;
use Omy\Crius\MerchantBasicServer\UpdateBranchBusinessNormalRequest;
use Omy\Crius\MerchantBasicServer\UpdateBranchBusinessSpecialRequest;
use Omy\Crius\MerchantBasicServer\UpdateBranchBusinessStatusRequest;
use Omy\Crius\MerchantBasicServer\UpdateBranchRequest;
use Omy\Crius\MerchantBasicServer\UpdateBranchResponse;
use Omy\Crius\MerchantBasicServer\UpdateBranchStatusRequest;
use Omy\Crius\MerchantBasicServer\UpdateBranchStatusResponse;
use Omy\Crius\MerchantBasicServer\UpdateBranchTagRequest;
use Omy\Crius\MerchantBasicServer\UpdateBranchTagStatusRequest;
use Omy\Crius\MerchantBasicServer\UpdateBranchTagStatusResponse;
use Omy\Crius\MerchantBasicServer\UpdateBrandRequest;
use Omy\Crius\MerchantBasicServer\UpdateBrandStatusRequest;
use Omy\Crius\MerchantBasicServer\UpdateConsumeCategoryRequest;
use Omy\Crius\MerchantBasicServer\UpdateDistrictRequest;
use Omy\Crius\MerchantBasicServer\UpdateGrowthRuleRequest;
use Omy\Crius\MerchantBasicServer\UpdateMemberBranchInfoRequest;
use Omy\Crius\MerchantBasicServer\UpdateMemberRequest;
use Omy\Crius\MerchantBasicServer\UpdateMemberResponse;
use Omy\Crius\MerchantBasicServer\CreatePointRuleRequest;
use Omy\Crius\MerchantBasicServer\GetMemberAddressRequest;
use Omy\Crius\MerchantBasicServer\GetMemberAddressResponse;
use Omy\Crius\MerchantBasicServer\GetPointRuleAllBranchResponse;
use Omy\Crius\MerchantBasicServer\GetPointRuleDescribeResponse;
use Omy\Crius\MerchantBasicServer\ListPointRuleRequest;
use Omy\Crius\MerchantBasicServer\ListPointRuleResponse;
use Omy\Crius\MerchantBasicServer\SetPointRuleDescribeRequest;
use Omy\Crius\MerchantBasicServer\ShowPointRuleRequest;
use Omy\Crius\MerchantBasicServer\ShowPointRuleResponse;
use Omy\Crius\MerchantBasicServer\UpdateMerchantLogoRequest;
use Omy\Crius\MerchantBasicServer\UpdateMerchantRequest;
use Omy\Crius\MerchantBasicServer\UpdateMerchantUserAgreementRequest;
use Omy\Crius\MerchantBasicServer\UpdatePasswordRequest;
use Omy\Crius\MerchantBasicServer\UpdatePasswordResponse;
use Omy\Crius\MerchantBasicServer\UpdatePointRuleRequest;
use Omy\Crius\MerchantBasicServer\Response;
use Omy\Crius\MerchantBasicServer\PBEmpty;
use Omy\Crius\MerchantBasicServer\UpdateRoleRequest;
use Omy\Crius\MerchantBasicServer\UpdateRoleResponse;
use Omy\Crius\MerchantBasicServer\UpdateRoleStatusRequest;
use Omy\Crius\MerchantBasicServer\UpdateRoleStatusResponse;
use Omy\Crius\MerchantBasicServer\UpdateRoomTypeCategoryRequest;
use Omy\Crius\MerchantBasicServer\UpdateRoomTypeCategoryResponse;
use Omy\Crius\MerchantBasicServer\UpdateRoomTypeRequest;
use Omy\Crius\MerchantBasicServer\UpdateRoomTypeStatusRequest;
use Omy\Crius\MerchantBasicServer\UpdateStaffRequest;
use Omy\Crius\MerchantBasicServer\UpdateStaffResponse;
use Omy\Crius\MerchantBasicServer\UpdateStaffStatusRequest;
use Omy\Crius\MerchantBasicServer\UpdateStaffStatusResponse;
use Omy\Crius\MerchantBasicServer\UpdateStatusRequest;


class MerchantBasicClient extends GrpcClient
{
    private const BASE_PATH = '/merchantBasic.MerchantBasicService/';

    //创建会员
    public const CREATE_MEMBER = 'CreateMember';

    //登录
    public const SIGN_IN = 'SignIn';

    //获取员工权限
    public const GET_PERMISSIONS = 'GetPermissions';

    //创建微信会员
    public const CREATE_WECHAT_MEMBER = 'CreateWechatMember';

    //更新会员
    public const UPDATE_MEMBER = 'UpdateMember';

    //会员详情
    public const SHOW_MEMBER = 'ShowMember';

    //通过ids批量获取会员信息
    public const GET_MEMBERS_BY_IDS = 'GetMembersByIDs';

    //通过手机号获取会员信息
    public const SHOW_MEMBER_ACCURATE_PHONE = 'ShowMemberByAccuratePhone';

    //获取营业时间
    public const SHOW_BRANCH_APPOINTMENT_BUSINESS_TIME = 'ShowBranchAppointmentBusinessTime';

    //门店预约详情
    public const SHOW_BRANCH_APPOINTMENT = 'ShowBranchAppointment';

    // 获取房型
    public const GET_ROOM_TYPE_CATEGORIES = 'GetRoomTypeCategories';

    //获取房型
    public const GET_ROOM_TYPES = 'GetRoomTypes';

    //门店列表
    public const GET_BRANCHES = 'GetBranches';

    //门店详情
    public const SHOW_BRANCH = 'ShowBranch';

    //获取微信会员列表
    public const GET_WECHAT_USER = 'GetWechatUser';

    public const GET_WECHAT_USER_BY_UNION_ID = 'GetWechatUserByUnionId';

    public const GET_WECHAT_USERS = 'GetWechatUsers';

    //创新更新微信会员信息
    public const CREATE_OR_UPDATE_WECHAT_USER = 'CreateOrUpdateWechatUser';

    public const UPDATE_MEMBER_BRANCH_INFO = 'UpdateMemberBranchInfo';

    //获取微信会员信息
    public const SHOW_WECHAT_USER = 'ShowWechatUser';

    //获取门店标签
    public const GET_BRANCH_TAGS = 'GetBranchTags';

    //通过branchids获取门店信息
    public const GET_BRANCH_TAGS_BY_IDS = 'GetBranchTagsByIDs';

    //创建小票账单
    public const CREATE_INVOICE = 'CreateInvoice';

    //获取会员列表
    public const GET_MEMBERS = 'GetMembers';

    //会员查找
    public const SEARCH_MEMBER = 'SearchMember';

    //创建用户行为
    public const CREATE_MEMBER_BEHAVIOR = 'CreateMemberBehavior';

    // 创建消费类型
    public const CREATE_CONSUME_CATEGORY = 'CreateConsumeCategory';

    // 消费类型列表
    public const LIST_CONSUME_CATEGORY = 'ListConsumeCategory';

    // 修改消费类型
    public const UPDATE_CONSUME_CATEGORY = 'UpdateConsumeCategory';

    // 修改消费类型状态
    public const UPDATE_CONSUME_CATEGORY_STATUS = 'UpdateConsumeCategoryStatus';

    // 查看积分类型详情
    public const SHOW_CONSUME_CATEGORY = 'ShowConsumeCategory';

    // 设置积分规则说明
    public const SET_POINT_RULE_DESCRIBE = 'SetPointRuleDescribe';

    // 获取积分规则说明
    public const GET_POINT_RULE_DESCRIBE = 'GetPointRuleDescribe';

    // 创建积分规则
    public const CREATE_POINT_RULE = 'CreatePointRule';

    // 修改积分规则
    public const UPDATE_POINT_RULE = 'UpdatePointRule';

    // 修改积分规则状态
    public const UPDATE_POINT_RULE_STATUS = 'UpdatePointRuleStatus';

    // 积分规则列表
    public const LIST_POINT_RULE = 'ListPointRule';

    // 积分规则详情
    public const SHOW_POINT_RULE = 'ShowPointRule';

    // 获取所有积分规则已设置门店
    public const GET_POINT_RULE_ALL_BRANCH = 'GetPointRuleAllBranch';

    //通过手机号获取员工信息
    public const SHOW_STAFF_BY_PHONE = 'ShowStaffByPhone';

    //更新密码
    public const UPDATE_PASSWORD = 'UpdatePassword';

    // CREATE_MEMBER_ADDRESS 新增用户地址
    public const CREATE_MEMBER_ADDRESS = 'CreateMemberAddress';

    // UPDATE_MEMBER_ADDRESS 修改用户地址
    public const UPDATE_MEMBER_ADDRESS = 'UpdateMemberAddress';

    // GET_MEMBER_ADDRESS 获取用户地址
    public const GET_MEMBER_ADDRESS = 'GetMemberAddress';

    public const GET_MEMBER_DEFAULT_ADDRESS = 'GetMemberDefaultAddress';

    public const SET_MEMBER_DEFAULT_ADDRESS = 'SetMemberAddressDefault';

    public const DELETE_MEMBER_ADDRESS = 'DeleteMemberAddress';

    //获取会员信息
    public const GET_STAFFS = 'GetStaffs';

    //创建权限
    public const CREATE_PERMISSION = 'CreatePermissions';

    //通过会员id  查询微信信息
    public const SHOW_WECHAT_USER_BY_MEMBER = 'ShowWechatUserByMember';

    public const SET_WECHAT_PAY_SETTING = 'SetWechatPaySetting';

    public const GET_WECHAT_PAY_SETTING = 'GetWechatPaySetting';

    //获取门店标签
    public const SHOW_BRANCH_TAG = 'ShowBranchTag';

    //更新房型状态
    public const UPDATE_ROOM_TYPE_STATUS = 'UpdateRoomTypeStatus';

    public const CREATE_OR_UPDATE_CALLING_MEMBER = 'CreateOrUpdateCallingMember';

    public const CREATE_ROOM_TYPE = 'CreateRoomType';

    public const UPDATE_ROOM_TYPE = 'UpdateRoomType';

    //成长值
    public const SHOW_GROWTH_CONFIG = 'ShowGrowthConfig';

    public const SAVE_GROWTH_CONFIG = 'SaveGrowthConfig';

    public const CREATE_GROWTH_RULE = 'CreateGrowthRule';

    public const GET_GROWTH_RULES = 'GetGrowthRules';

    public const SHOW_GROWTH_RULE = 'ShowGrowthRule';

    public const UPDATE_GROWTH_RULE = 'UpdateGrowthRule';

    public const GET_BRANCHES_HAS_GROWTH_RULE = 'GetBranchesHasGrowthRule';

    public const GET_BRANCH_GROWTH_RULE = 'GetBranchGrowthRule';

    //门店
    public const CREATE_BRANCH = 'CreateBranch';

    public const UPDATE_BRANCH = 'UpdateBranch';

    public const DELETE_BRANCH = 'DeleteBranch';

    public const UPDATE_BRANCH_STATUS = 'UpdateBranchStatus';

    // 获取标签
    public const CREATE_BRANCH_TAG = 'CreateBranchTag';

    //更新标签
    public const UPDATE_BRANCH_TAG = 'UpdateBranchTag';

    // 更新标签状态
    public const UPDATE_BRANCH_TAG_STATUS = 'UpdateBranchTagStatus';


    //品牌管理
    public const CREATE_BRAND = 'CreateBrand';

    public const UPDATE_BRAND = 'UpdateBrand';

    public const GET_BRANDS = 'GetBrands';

    public const UPDATE_BRAND_STATUS = 'UpdateBrandStatus';

//地区
    public const CREATE_DISTRICT = 'CreateDistrict';

    public const UPDATE_DISTRICT = 'UpdateDistrict';

    public const GET_DISTRICTS = 'GetDistricts';

    //角色
    public const CREATE_ROLE = 'CreateRole';
    public const UPDATE_ROLE = 'UpdateRole';
    public const GET_ROLES = 'GetRoles';
    public const SHOW_ROLE = 'ShowRole';
    public const UPDATE_ROLE_STATUS = 'UpdateRoleStatus';
    public const DELETE_ROLE = 'DeleteRole';
    public const GET_ROLE_HISTORIES = 'GetRoleHistories';
    public const GET_STAFFS_BY_ROLE_ID = 'GetStaffsByRoleID';

    //员工
    public const CREATE_STAFF = 'CreateStaff';
    public const SHOW_STAFF = 'ShowStaff';
    public const UPDATE_STAFF = 'UpdateStaff';
    public const DELETE_STAFF = 'DeleteStaff';
    public const UPDATE_STAFF_STATUS = 'UpdateStaffStatus';
    //房型分类
    public const CREATE_ROOM_TYPE_CATEGORY = 'CreateRoomTypeCategory';
    public const UPDATE_ROOM_TYPE_CATEGORY = 'UpdateRoomTypeCategory';
    public const SAVE_BRANCH_APPOINTMENT = "SaveBranchAppointment";
    public const UPDATE_BRANCH_APPOINTMENT_ROOM_TYPE = "UpdateBranchAppointmentRoomType";

    //营业时间管理
    public const CREATE_BRANCH_BUSINESS = 'CreateBranchBusinessSpecial';

    public const UPDATE_BRANCH_BUSINESS_NORMAL = 'UpdateBranchBusinessNormal';

    public const UPDATE_BRANCH_BUSINESS_SPECIAL = 'UpdateBranchBusinessSpecial';

    public const GET_BRANCH_BUSINESSES = 'GetBranchBusinesses';

    public const UPDATE_BRANCH_BUSINESS_STATUS = 'UpdateBranchBusinessStatus';

    public const GET_MEMBER_BEHAVIORS = 'GetMemberBehaviors';

    public const GET_BRANCH_WITH_SUB_MCH_ID = 'GetBranchWithSubMchID';

    public const SET_BRANCH_SUB_MCH_ID = 'SetBranchSubMchID';

    //商品套餐
    public const SHOW_GOODS = 'ShowGoods';
    public const SHOW_PACKAGE = 'ShowPackage';
    public const MULTI_GET_GOODS_AND_PACKAGES = 'MultiGetGoodsAndPackages';
    public const SEARCH_GOODS_AND_PACKAGE = 'SearchGoodsAndPackage';
    public const SEARCH_GOODS_OR_PACKAGE = 'SearchGoodsOrPackage';

    //运营商品类型
    public const GET_OPERATE_TYPES = 'GetOperateTypes';
    public const GET_ACTIVE_TYPES = 'GetActiveTypes';
    public const GET_CONSUME_CATEGORY_TYPES = 'GetConsumeCategoryTypes';

    //房型组
    public const GET_ROOM_TYPE_GROUPS = 'GetRoomTypeGroups';

    public const GET_BRANCH_POINT_RULE = 'GetBranchPointRule';

    public const GET_BRANCH_POINT_RULES = 'GetBranchPointRules';

    public const SHOW_MERCHANT = 'ShowMerchant';

    public const UPDATE_MERCHANT_LOGO = 'UpdateMerchantLogo';

    public const UPDATE_MERCHANT_USER_AGREEMENT = 'UpdateMerchantUserAgreement';

    public const GET_BRANCHES_BY_TAG_IDS = 'GetBranchesByTagIds';

    protected array $actionMap = [
        self::GET_BRANCHES_BY_TAG_IDS => [
            GetBranchesByTagIDsRequest::class,
            GetBranchesByTagIDsResponse::class,
        ],
        self::SHOW_MERCHANT => [
            EmptyRequest::class,
            ShowMerchantResponse::class,
        ],
        self::UPDATE_MERCHANT_LOGO => [
            UpdateMerchantLogoRequest::class,
            Response::class,
        ],
        self::UPDATE_MERCHANT_USER_AGREEMENT => [
            UpdateMerchantUserAgreementRequest::class,
            Response::class,
        ],
        self::GET_BRANCH_POINT_RULES => [
            GetBranchPointRulesRequest::class,
            GetBranchPointRulesResponse::class,
        ],
        self::GET_BRANCH_POINT_RULE => [
            GetBranchPointRuleRequest::class,
            GetBranchPointRuleResponse::class
        ],
        self::SET_BRANCH_SUB_MCH_ID => [
            SetBranchSubMchIDRequest::class,
            PBEmpty::class
        ],
        self::GET_BRANCH_WITH_SUB_MCH_ID => [
            GetBranchWithSubMchIDRequest::class,
            GetBranchWithSubMchIDResponse::class
        ],
        self::GET_MEMBER_BEHAVIORS => [
            GetMemberBehaviorsRequest::class,
            GetMemberBehaviorsResponse::class
        ],
        self::UPDATE_BRANCH_BUSINESS_STATUS => [
            UpdateBranchBusinessStatusRequest::class,
            PBEmpty::class
        ],
        self::GET_BRANCH_BUSINESSES => [
            GetBranchBusinessesRequest::class,
            GetBranchBusinessesResponse::class
        ],
        self::UPDATE_BRANCH_BUSINESS_SPECIAL => [
            UpdateBranchBusinessSpecialRequest::class,
            PBEmpty::class
        ],
        self::UPDATE_BRANCH_BUSINESS_NORMAL => [
            UpdateBranchBusinessNormalRequest::class,
            PBEmpty::class
        ],
        self::CREATE_BRANCH_BUSINESS => [
            CreateBranchBusinessSpecialRequest::class,
            PBEmpty::class
        ],
        self::UPDATE_BRANCH_APPOINTMENT_ROOM_TYPE => [
            UpdateBranchAppointmentRoomTypeRequest::class,
            UpdateBranchAppointmentRoomTypeResponse::class
        ],
        self::SAVE_BRANCH_APPOINTMENT => [
            SaveBranchAppointmentRequest::class,
            SaveBranchAppointmentResponse::class
        ],
        self::CREATE_ROOM_TYPE_CATEGORY => [
            CreateRoomTypeCategoryRequest::class,
            CreateRoomTypeCategoryResponse::class
        ],
        self::UPDATE_ROOM_TYPE_CATEGORY => [
            UpdateRoomTypeCategoryRequest::class,
            UpdateRoomTypeCategoryResponse::class
        ],
        self::SHOW_STAFF => [
            ShowStaffRequest::class,
            ShowStaffResponse::class
        ],
        self::CREATE_STAFF => [
            CreateStaffRequest::class,
            CreateStaffResponse::class
        ],
        self::UPDATE_STAFF => [
            UpdateStaffRequest::class,
            UpdateStaffResponse::class
        ],
        self::DELETE_STAFF => [
            DeleteStaffRequest::class,
            DeleteStaffResponse::class
        ],
        self::UPDATE_STAFF_STATUS => [
            UpdateStaffStatusRequest::class,
            UpdateStaffStatusResponse::class
        ],
        self::CREATE_ROLE => [
            CreateRoleRequest::class,
            CreateRoleResponse::class
        ],
        self::UPDATE_ROLE => [
            UpdateRoleRequest::class,
            UpdateRoleResponse::class
        ],
        self::GET_ROLES => [
            GetRolesRequest::class,
            GetRolesResponse::class
        ],
        self::SHOW_ROLE => [
            ShowRoleRequest::class,
            ShowRoleResponse::class
        ],
        self::UPDATE_ROLE_STATUS => [
            UpdateRoleStatusRequest::class,
            UpdateRoleStatusResponse::class
        ],
        self::DELETE_ROLE => [
            DeleteRoleRequest::class,
            DeleteRoleResponse::class
        ],
        self::GET_ROLE_HISTORIES => [
            GetRoleHistoriesRequest::class,
            GetRoleHistoriesResponse::class
        ],
        self::GET_STAFFS_BY_ROLE_ID => [
            GetStaffsByRoleIDRequest::class,
            GetStaffsByRoleIDResponse::class
        ],

        self::CREATE_DISTRICT => [
            CreateDistrictRequest::class,
            Response::class
        ],
        self::UPDATE_DISTRICT => [
            UpdateDistrictRequest::class,
            Response::class
        ],

        self::GET_DISTRICTS => [
            GetDistrictsRequest::class,
            GetDistrictsResponse::class
        ],

        self::CREATE_BRAND => [
            CreateBrandRequest::class,
            Response::class
        ],
        self::UPDATE_BRAND => [
            UpdateBrandRequest::class,
            Response::class
        ],
        self::GET_BRANDS => [
            GetBrandsRequest::class,
            GetBrandsResponse::class
        ],
        self::UPDATE_BRAND_STATUS => [
            UpdateBrandStatusRequest::class,
            Response::class
        ],

        self::CREATE_BRANCH_TAG => [
            CreateBranchTagRequest::class,
            CreateBranchTagResponse::class,
        ],
        self::UPDATE_BRANCH_TAG => [
            UpdateBranchTagRequest::class,
            NoDataResponse::class
        ],
        self::UPDATE_BRANCH_TAG_STATUS => [
            UpdateBranchTagStatusRequest::class,
            UpdateBranchTagStatusResponse::class
        ],

        self::UPDATE_BRANCH_STATUS => [
            UpdateBranchStatusRequest::class,
            UpdateBranchStatusResponse::class
        ],
        self::CREATE_BRANCH => [
            CreateBranchRequest::class,
            CreateBranchResponse::class
        ],
        self::UPDATE_BRANCH => [
            UpdateBranchRequest::class,
            UpdateBranchResponse::class
        ],
        self::DELETE_BRANCH => [
            DeleteBranchRequest::class,
            DeleteBranchResponse::class
        ],
        self::SHOW_GROWTH_CONFIG => [
            PBEmpty::class,
            ShowGrowthConfigResponse::class
        ],
        self::SAVE_GROWTH_CONFIG => [
            SaveGrowthConfigRequest::class,
            Response::class
        ],
        self::CREATE_GROWTH_RULE => [
            CreateGrowthRuleRequest::class,
            Response::class
        ],
        self::GET_GROWTH_RULES => [
            GetGrowthRulesRequest::class,
            GetGrowthRulesResponse::class
        ],
        self::SHOW_GROWTH_RULE => [
            ShowGrowthRuleRequest::class,
            ShowGrowthRuleResponse::class,
        ],
        self::UPDATE_GROWTH_RULE => [
            UpdateGrowthRuleRequest::class,
            Response::class
        ],
        self::GET_BRANCHES_HAS_GROWTH_RULE => [
            PBEmpty::class,
            GetBranchesHasGrowthRuleResponse::class
        ],
        self::GET_BRANCH_GROWTH_RULE => [
            GetBranchGrowthRuleRequest::class,
            GetBranchGrowthRuleResponse::class
        ],
        self::UPDATE_ROOM_TYPE => [
            UpdateRoomTypeRequest::class,
            Response::class
        ],
        self::CREATE_ROOM_TYPE => [
            CreateRoomTypeRequest::class,
            Response::class
        ],
        self::CREATE_OR_UPDATE_CALLING_MEMBER => [
            CreateOrUpdateCallingMemberRequest::class,
            CreateOrUpdateCallingMemberResponse::class
        ],
        self::UPDATE_ROOM_TYPE_STATUS => [
            UpdateRoomTypeStatusRequest::class,
            Response::class
        ],
        self::SHOW_BRANCH_TAG => [
            ShowBranchTagRequest::class,
            ShowBranchTagResponse::class
        ],
        self::SHOW_WECHAT_USER_BY_MEMBER => [
            ShowWechatUserByMemberRequest::class,
            ShowWechatUserByMemberResponse::class
        ],
        self::CREATE_PERMISSION => [
            CreatePermissionsRequest::class,
            CreatePermissionsResponse::class
        ],
        self::GET_STAFFS => [
            GetStaffsRequest::class,
            GetStaffsResponse::class
        ],
        self::GET_PERMISSIONS => [
            GetPermissionsRequest::class,
            GetPermissionsResponse::class
        ],
        self::UPDATE_PASSWORD => [
            UpdatePasswordRequest::class,
            UpdatePasswordResponse::class
        ],
        self::SHOW_STAFF_BY_PHONE => [
            ShowStaffByPhoneRequest::class,
            ShowStaffByPhoneResponse::class
        ],
        self::CREATE_MEMBER_BEHAVIOR => [
            CreateMemberBehaviorRequest::class,
            NoDataResponse::class
        ],
        self::GET_MEMBERS => [
            GetMembersRequest::class,
            GetMembersResponse::class
        ],
        self::SIGN_IN => [
            SignInRequest::class,
            SignInResponse::class
        ],
        self::CREATE_MEMBER => [
            CreateMemberRequest::class,
            CreateMemberResponse::class
        ],
        self::CREATE_WECHAT_MEMBER => [
            CreateWechatMemberRequest::class,
            CreateWechatMemberResponse::class
        ],
        self::UPDATE_MEMBER => [
            UpdateMemberRequest::class,
            UpdateMemberResponse::class
        ],
        self::SHOW_MEMBER => [
            ShowMemberRequest::class,
            ShowMemberResponse::class
        ],
        self::GET_MEMBERS_BY_IDS => [
            GetMembersByIDsRequest::class,
            GetMembersByIDsResponse::class
        ],
        self::SHOW_MEMBER_ACCURATE_PHONE => [
            ShowMemberByAccuratePhoneRequest::class,
            ShowMemberByAccuratePhoneResponse::class
        ],
        self::GET_WECHAT_USER => [
            GetWechatUserRequest::class,
            GetWechatUserResponse::class
        ],
        self::GET_WECHAT_USER_BY_UNION_ID => [
            GetWechatUserByUnionIdRequest::class,
            GetWechatUserByUnionIdResponse::class
        ],
        self::CREATE_OR_UPDATE_WECHAT_USER => [
            CreateOrUpdateWechatUserRequest::class,
            CreateOrUpdateWechatUserResponse::class
        ],
        self::SHOW_WECHAT_USER => [
            ShowWechatUserRequest::class,
            ShowWechatUserResponse::class
        ],
        self::GET_BRANCHES => [
            GetBranchesRequest::class,
            GetBranchesResponse::class
        ],
        self::SHOW_BRANCH => [
            ShowBranchRequest::class,
            ShowBranchResponse::class
        ],
        self::SHOW_BRANCH_APPOINTMENT_BUSINESS_TIME => [
            ShowBranchAppointmentBusinessTimeRequest::class,
            ShowBranchAppointmentBusinessTimeResponse::class
        ],
        self::GET_ROOM_TYPE_CATEGORIES => [
            GetRoomTypeCategoriesRequest::class,
            GetRoomTypeCategoriesResponse::class
        ],
        self::GET_ROOM_TYPES => [
            GetRoomTypesRequest::class,
            GetRoomTypesResponse::class
        ],
        self::GET_BRANCH_TAGS => [
            GetBranchTagsRequest::class,
            GetBranchTagsResponse::class
        ],
        self::GET_BRANCH_TAGS_BY_IDS => [
            GetBranchTagsByIDsRequest::class,
            GetBranchTagsByIDsResponse::class
        ],
        self::SHOW_BRANCH_APPOINTMENT => [
            ShowBranchAppointmentRequest::class,
            ShowBranchAppointmentResponse::class
        ],
        self::CREATE_INVOICE => [
            CreateInvoiceRequest::class,
            CreateInvoiceResponse::class
        ],
        self::CREATE_CONSUME_CATEGORY => [
            CreateConsumeCategoryRequest::class,
            Response::class,
        ],
        self::LIST_CONSUME_CATEGORY => [
            ListConsumeCategoryRequest::class,
            ListConsumeCategoryResponse::class
        ],
        self::UPDATE_CONSUME_CATEGORY => [
            UpdateConsumeCategoryRequest::class,
            Response::class
        ],
        self::UPDATE_CONSUME_CATEGORY_STATUS => [
            UpdateStatusRequest::class,
            Response::class
        ],
        self::SHOW_CONSUME_CATEGORY => [
            ShowConsumeCategoryRequest::class,
            ShowConsumeCategoryResponse::class
        ],
        self::SET_POINT_RULE_DESCRIBE => [
            SetPointRuleDescribeRequest::class,
            Response::class
        ],
        self::GET_POINT_RULE_DESCRIBE => [
            PBEmpty::class,
            GetPointRuleDescribeResponse::class,
        ],
        self::CREATE_POINT_RULE => [
            CreatePointRuleRequest::class,
            Response::class
        ],
        self::UPDATE_POINT_RULE => [
            UpdatePointRuleRequest::class,
            Response::class,
        ],
        self::UPDATE_POINT_RULE_STATUS => [
            UpdateStatusRequest::class,
            Response::class
        ],
        self::LIST_POINT_RULE => [
            ListPointRuleRequest::class,
            ListPointRuleResponse::class,
        ],
        self::SHOW_POINT_RULE => [
            ShowPointRuleRequest::class,
            ShowPointRuleResponse::class
        ],
        self::GET_POINT_RULE_ALL_BRANCH => [
            PBEmpty::class,
            GetPointRuleAllBranchResponse::class
        ],
        self::CREATE_MEMBER_ADDRESS => [
            MemberAddress::class,
            Response::class,
        ],
        self::UPDATE_MEMBER_ADDRESS => [
            MemberAddress::class,
            Response::class,
        ],
        self::GET_MEMBER_ADDRESS => [
            GetMemberAddressRequest::class,
            GetMemberAddressResponse::class
        ],
        self::SET_WECHAT_PAY_SETTING => [
            SetWechatPaySettingRequest::class,
            Response::class
        ],
        self::GET_WECHAT_PAY_SETTING => [
            GetWechatPaySettingRequest::class,
            GetWechatPaySettingResponse::class
        ],
        self::GET_MEMBER_DEFAULT_ADDRESS => [
            GetMemberDefaultAddressRequest::class,
            GetMemberDefaultAddressResponse::class,
        ],
        self::SET_MEMBER_DEFAULT_ADDRESS => [
            SetMemberAddressDefaultRequest::class,
            Response::class
        ],
        self::DELETE_MEMBER_ADDRESS => [
            DeleteMemberAddressRequest::class,
            Response::class
        ],
        self::SHOW_GOODS => [
            ShowGoodsRequest::class,
            ShowGoodsResponse::class
        ],
        self::SHOW_PACKAGE => [
            ShowPackageRequest::class,
            ShowPackageResponse::class
        ],
        self::MULTI_GET_GOODS_AND_PACKAGES => [
            MultiGetGoodsAndPackagesRequest::class,
            MultiGetGoodsAndPackagesResponse::class
        ],
        self::SEARCH_GOODS_AND_PACKAGE => [
            SearchGoodsAndPackageRequest::class,
            SearchGoodsAndPackageResponse::class
        ],
        self::SEARCH_GOODS_OR_PACKAGE => [
            SearchGoodsOrPackageRequest::class,
            SearchGoodsOrPackageResponse::class
        ],
        self::GET_WECHAT_USERS => [
            GetWechatUsersRequest::class,
            GetWechatUsersResponse::class,
        ],
        self::GET_OPERATE_TYPES => [
            PBEmpty::class,
            GetOperateTypesResponse::class
        ],
        self::GET_ACTIVE_TYPES => [
            PBEmpty::class,
            GetActiveTypesResponse::class
        ],
        self::GET_ROOM_TYPE_GROUPS => [
            PBEmpty::class,
            GetRoomTypeGroupsResponse::class
        ],
        self::GET_CONSUME_CATEGORY_TYPES => [
            PBEmpty::class,
            GetConsumeCategoryTypesResponse::class
        ],
        self::SEARCH_MEMBER => [
            SearchMemberRequest::class,
            SearchMemberResponse::class
        ],
        self::UPDATE_MEMBER_BRANCH_INFO => [
            UpdateMemberBranchInfoRequest::class,
            Response::class
        ]
    ];

    protected function getBasePath()
    {
        return self::BASE_PATH;
    }
}
