<?php

declare(strict_types=1);

namespace Omy\MicroServers\Client;

use Omy\Crius\BasicServer\EmptyRequest;
use Omy\Crius\ExtensionServer\CouponStyle;
use Omy\Crius\ExtensionServer\CreateBlackListRequest;
use Omy\Crius\ExtensionServer\CreateCouponCategoryRequest;
use Omy\Crius\ExtensionServer\CreateCouponRequest;
use Omy\Crius\ExtensionServer\CreateDistributeCouponRequest;
use Omy\Crius\ExtensionServer\CreateDivisionResponse;
use Omy\Crius\ExtensionServer\CreateLevelResponse;
use Omy\Crius\ExtensionServer\CreateMemberCouponRequest;
use Omy\Crius\ExtensionServer\CreateOrderExpressRequest;
use Omy\Crius\ExtensionServer\CreateTaskRequest;
use Omy\Crius\ExtensionServer\DeleteBlackListRequest;
use Omy\Crius\ExtensionServer\DeleteDivisionRequest;
use Omy\Crius\ExtensionServer\DeleteLevelRequest;
use Omy\Crius\ExtensionServer\GetBlackListConfigsResponse;
use Omy\Crius\ExtensionServer\GetBlackListsRequest;
use Omy\Crius\ExtensionServer\GetBlackListsResponse;
use Omy\Crius\ExtensionServer\GetCollectableCouponsRequest;
use Omy\Crius\ExtensionServer\GetCollectableCouponsResponse;
use Omy\Crius\ExtensionServer\GetCouponCategoriesRequest;
use Omy\Crius\ExtensionServer\GetCouponCategoriesResponse;
use Omy\Crius\ExtensionServer\GetCouponsRequest;
use Omy\Crius\ExtensionServer\GetCouponsResponse;
use Omy\Crius\ExtensionServer\GetDistributeCouponsRequest;
use Omy\Crius\ExtensionServer\GetDistributeCouponsResponse;
use Omy\Crius\ExtensionServer\GetDivisionsAndLevelsResponse;
use Omy\Crius\ExtensionServer\GetLevelTasksRequest;
use Omy\Crius\ExtensionServer\GetLevelTasksResponse;
use Omy\Crius\ExtensionServer\GetMemberBenefitsRequest;
use Omy\Crius\ExtensionServer\GetMemberBenefitsResponse;
use Omy\Crius\ExtensionServer\GetMemberCouponsRequest;
use Omy\Crius\ExtensionServer\GetMemberCouponsResponse;
use Omy\Crius\ExtensionServer\GetMemberLevelBenefitsRequest;
use Omy\Crius\ExtensionServer\GetMemberLevelTasksRequest;
use Omy\Crius\ExtensionServer\GetMemberLevelTasksResponse;
use Omy\Crius\ExtensionServer\GetMemberTaskAvailableRequest;
use Omy\Crius\ExtensionServer\GetMemberTaskAvailableResponse;
use Omy\Crius\ExtensionServer\GetMemberTaskFinishedRequest;
use Omy\Crius\ExtensionServer\GetMemberTaskFinishedResponse;
use Omy\Crius\ExtensionServer\GetRewardsRequest;
use Omy\Crius\ExtensionServer\GetTasksRequest;
use Omy\Crius\ExtensionServer\GetTasksResponse;
use Omy\Crius\ExtensionServer\MemberUsePointResponse;
use Omy\Crius\ExtensionServer\PBEmpty;
use Omy\Crius\ExtensionServer\BatchGetMemberExtensionRequest;
use Omy\Crius\ExtensionServer\BatchGetMemberExtensionResponse;
use Omy\Crius\ExtensionServer\CreateBenefitRequest;
use Omy\Crius\ExtensionServer\CreateDivisionRequest;
use Omy\Crius\ExtensionServer\CreateLevelRequest;
use Omy\Crius\ExtensionServer\CreateProductCategoryRequest;
use Omy\Crius\ExtensionServer\CreateProductOrderRequest;
use Omy\Crius\ExtensionServer\CreateProductOrderResponse;
use Omy\Crius\ExtensionServer\CreateProductRequest;
use Omy\Crius\ExtensionServer\CreateProductResponse;
use Omy\Crius\ExtensionServer\GetBenefitsRequest;
use Omy\Crius\ExtensionServer\GetBenefitsResponse;
use Omy\Crius\ExtensionServer\GetDivisionsRequest;
use Omy\Crius\ExtensionServer\GetDivisionsResponse;
use Omy\Crius\ExtensionServer\GetGrowthBillsRequest;
use Omy\Crius\ExtensionServer\GetGrowthBillsResponse;
use Omy\Crius\ExtensionServer\GetLevelsRequest;
use Omy\Crius\ExtensionServer\GetLevelsResponse;
use Omy\Crius\ExtensionServer\GetLevelsWithDivisionRequest;
use Omy\Crius\ExtensionServer\GetLevelsWithDivisionResponse;
use Omy\Crius\ExtensionServer\GetMemberPointBillsRequest;
use Omy\Crius\ExtensionServer\GetMemberPointBillsResponse;
use Omy\Crius\ExtensionServer\GetMemberExtensionRequest;
use Omy\Crius\ExtensionServer\GetMemberExtensionResponse;
use Omy\Crius\ExtensionServer\GetProductOrdersRequest;
use Omy\Crius\ExtensionServer\GetProductOrdersResponse;
use Omy\Crius\ExtensionServer\GetProductQuantityBillsRequest;
use Omy\Crius\ExtensionServer\GetProductQuantityBillsResponse;
use Omy\Crius\ExtensionServer\GetProductsRequest;
use Omy\Crius\ExtensionServer\GetProductsResponse;
use Omy\Crius\ExtensionServer\ListProductCategoryRequest;
use Omy\Crius\ExtensionServer\ListProductCategoryResponse;
use Omy\Crius\ExtensionServer\MemberGainPointByRuleRequest;
use Omy\Crius\ExtensionServer\MemberGainPointRequest;
use Omy\Crius\ExtensionServer\MemberUsePointByRuleRequest;
use Omy\Crius\ExtensionServer\MemberUsePointRequest;
use Omy\Crius\ExtensionServer\RemoveMemberPointBillLockRequest;
use Omy\Crius\ExtensionServer\Response;
use Omy\Crius\ExtensionServer\SaveBlackListConfigRequest;
use Omy\Crius\ExtensionServer\SavePointConfigRequest;
use Omy\Crius\ExtensionServer\ShowBenefitRequest;
use Omy\Crius\ExtensionServer\ShowBenefitResponse;
use Omy\Crius\ExtensionServer\ShowBlackListByPhoneRequest;
use Omy\Crius\ExtensionServer\ShowBlackListByPhoneResponse;
use Omy\Crius\ExtensionServer\ShowCouponCategoryRequest;
use Omy\Crius\ExtensionServer\ShowCouponCategoryResponse;
use Omy\Crius\ExtensionServer\ShowCouponRequest;
use Omy\Crius\ExtensionServer\ShowCouponResponse;
use Omy\Crius\ExtensionServer\ShowCouponStyleResponse;
use Omy\Crius\ExtensionServer\ShowDistributeCouponRequest;
use Omy\Crius\ExtensionServer\ShowDistributeCouponResponse;
use Omy\Crius\ExtensionServer\ShowDivisionRequest;
use Omy\Crius\ExtensionServer\ShowDivisionResponse;
use Omy\Crius\ExtensionServer\ShowLevelRequest;
use Omy\Crius\ExtensionServer\ShowLevelResponse;
use Omy\Crius\ExtensionServer\SecondLevelSortProductCategoryResponse;
use Omy\Crius\ExtensionServer\ShowMemberCouponRequest;
use Omy\Crius\ExtensionServer\ShowMemberCouponResponse;
use Omy\Crius\ExtensionServer\ShowPointConfigRequest;
use Omy\Crius\ExtensionServer\ShowPointConfigResponse;
use Omy\Crius\ExtensionServer\ShowProductOrderRequest;
use Omy\Crius\ExtensionServer\ShowProductOrderResponse;
use Omy\Crius\ExtensionServer\ShowProductRequest;
use Omy\Crius\ExtensionServer\ShowProductResponse;
use Omy\Crius\ExtensionServer\ShowTaskRequest;
use Omy\Crius\ExtensionServer\ShowTaskResponse;
use Omy\Crius\ExtensionServer\SubmitDivisionLevelRequest;
use Omy\Crius\ExtensionServer\UpdateBenefitRequest;
use Omy\Crius\ExtensionServer\UpdateBenefitStatusRequest;
use Omy\Crius\ExtensionServer\UpdateCouponCategoryRequest;
use Omy\Crius\ExtensionServer\UpdateCouponRequest;
use Omy\Crius\ExtensionServer\UpdateCouponStatusRequest;
use Omy\Crius\ExtensionServer\UpdateDistributeCouponRequest;
use Omy\Crius\ExtensionServer\UpdateDistributeCouponStatusRequest;
use Omy\Crius\ExtensionServer\UpdateDivisionRequest;
use Omy\Crius\ExtensionServer\UpdateLevelRequest;
use Omy\Crius\ExtensionServer\UpdateProductCategoryRequest;
use Omy\Crius\ExtensionServer\UpdateProductOrderStatusRequest;
use Omy\Crius\ExtensionServer\UpdateProductQuantityRequest;
use Omy\Crius\ExtensionServer\UpdateProductRequest;
use Omy\Crius\ExtensionServer\UpdateProductStatusRequest;
use Omy\Crius\ExtensionServer\UpdateStatusRequest;
use Omy\Crius\ExtensionServer\UpdateTaskRequest;
use Omy\Crius\ExtensionServer\UpdateTaskStatusRequest;

class MemberExtensionClient extends GrpcClient
{
    private const BASE_PATH = '/memberExtension.ExtensionServer/';

    // 用户直接获得积分
    public const MEMBER_GAIN_POINT = 'MemberGainPoint';

    // 用户用过规则获得积分
    public const MEMBER_GAIN_POINT_BY_RULE = 'MemberGainPointByRule';

    // 用户直接使用积分
    public const MEMBER_USE_POINT = 'MemberUsePoint';

    // 用户通过规则使用积分
    public const MEMBER_USE_POINT_BY_RULE = 'MemberUsePointByRule';

    // 移除用户积分流水锁
    public const REMOVE_MEMBER_POINT_BILL_LOCK = 'RemoveMemberPointBillLock';

    // 用户用户积分账户
    public const GET_MEMBER_EXTENSION= 'GetMemberExtension';

    // 批量获取用户积分账户
    public const BATCH_GET_MEMBER_EXTENSION = 'BatchGetMemberExtension';

    // 获取用户积分流水
    public const GET_MEMBER_POINT_BILLS = 'GetMemberPointBills';

    // 创建商品分类
    public const CREATE_PRODUCT_CATEGORY = 'CreateProductCategory';

    // 修改商品分类
    public const UPDATE_PRODUCT_CATEGORY = 'UpdateProductCategory';

    // 修改商品分类状态
    public const UPDATE_PRODUCT_CATEGORY_STATUS = 'UpdateProductCategoryStatus';

    // 商品分类列表
    public const LIST_PRODUCT_CATEGORY = 'ListProductCategory';
    // 商品二级分类返回
    public const SECOND_LEVEL_SORT_PRODUCT_CATEGORY = 'SecondLevelSortProductCategory';


    //新建商品
    public const CREATE_PRODUCT = 'CreateProduct';

    //商品详情
    public const SHOW_PRODUCT = 'ShowProduct';

    //获取商品列表
    public const GET_PRODUCTS = 'GetProducts';

    //更新库存
    public const UPDATE_PRODUCT_QUANTITY = 'UpdateProductQuantity';

    //修改商品
    public const UPDATE_PRODUCT = 'UpdateProduct';

    //获取商品库存流水
    public const GET_PRODUCT_QUANTITY_BILLS = 'GetProductQuantityBills';


    //订单
    public const SHOW_PRODUCT_ORDER = 'ShowProductOrder';

    //创建订单
    public const CREATE_PRODUCT_ORDER = 'CreateProductOrder';

    //订单列表
    public const GET_PRODUCT_ORDERS = 'GetProductOrders';

    //更新订单状态
    public const UPDATE_PRODUCT_ORDER_STATUS = 'UpdateProductOrderStatus';

    public const SAVE_POINT_CONFIG = 'SavePointConfig';

    public const SHOW_POINT_CONFIG = 'ShowPointConfig';

    //更新商品状态
    public const UPDATE_PRODUCT_STATUS = 'UpdateProductStatus';

    //权益
    public const CREATE_BENEFITS = 'CreateBenefit';
    public const UPDATE_BENEFITS = 'UpdateBenefit';
    public const GET_BENEFITS = 'GetBenefits';
    public const UPDATE_BENEFITS_STATUS = 'UpdateBenefitStatus';
    public const SHOW_BENEFIT = 'ShowBenefit';

    //段位
    public const CREATE_DIVISION = 'CreateDivision';
    public const UPDATE_DIVISION = 'UpdateDivision';
    public const GET_DIVISIONS = 'GetDivisions';
    public const SHOW_DIVISIONS = 'ShowDivision';

    //等级
    public const CREATE_LEVEL = 'CreateLevel';
    public const UPDATE_LEVEL = 'UpdateLevel';
    public const GET_LEVELS = 'GetLevels';
    public const SHOW_LEVEL = 'ShowLevel';
    public const SUBMIT_LEVEL = 'SubmitDivisionLevel';
    public const GET_LEVELS_WITH_DIVISION = 'GetLevelsWithDivision';

    //成长值
    public const GET_GROWTH_BILLS = 'GetGrowthBills';

    //创建订单express
    public const CREATE_ORDER_EXPRESS = 'CreateOrderExpress';

    //删除等级
    public const DELETE_LEVEL = 'DeleteLevel';
    public const DELETE_DIVISION = 'DeleteDivision';

    public const GET_DIVISIONS_AND_LEVELS = 'GetDivisionsAndLevels';

    //创建优惠券
    public const CREATE_COUPON = 'CreateCoupon';
    public const UPDATE_COUPON = 'UpdateCoupon';
    public const GET_COUPONS = 'GetCoupons';
    public const SHOW_COUPON = 'ShowCoupon';
    public const UPDATE_COUPON_STATUS = 'UpdateCouponStatus';

    //创建优惠券类型
    public const CREATE_COUPON_CATEGORY = 'CreateCouponCategory';
    public const UPDATE_COUPON_CATEGORY = 'UpdateCouponCategory';
    public const GET_COUPON_CATEGORY = 'GetCouponCategories';
    public const SHOW_COUPON_CATEGORY = 'ShowCouponCategory';

    //创建优惠券
    public const CREATE_COUPON_STYLE = 'SaveCouponStyle';
    public const SHOW_COUPON_STYLE = 'ShowCouponStyle';

    //编辑推送优惠券
    public const CREATE_DISTRIBUTE_COUPON = 'CreateDistributeCoupon';
    public const UPDATE_DISTRIBUTE_COUPON = 'UpdateDistributeCoupon';
    public const GET_DISTRIBUTE_COUPONS = 'GetDistributeCoupons';
    public const SHOW_DISTRIBUTE_COUPON = 'ShowDistributeCoupon';
    public const UPDATE_DISTRIBUTE_COUPON_STATUS = 'UpdateDistributeCouponStatus';

    //用户券
    public const GET_MEMBER_COUPONS = 'GetMemberCoupons';
    public const SHOW_MEMBER_COUPON = 'ShowMemberCoupon';
    public const GET_COLLECTABLE_COUPONS = 'GetCollectableCoupons';
    public const CREATE_MEMBER_COUPON = 'CreateMemberCoupon';

    //任务
    public const CREATE_TASK = 'CreateTask';
    public const UPDATE_TASK = 'UpdateTask';
    public const GET_TASKS = 'GetTasks';
    public const SHOW_TASK = 'ShowTask';
    public const UPDATE_TASK_STATUS = 'UpdateTaskStatus';

    //会员任务
    public const GET_MEMBER_TASK_AVAILABLE = 'GetMemberTaskAvailable';
    public const GET_MEMBER_TASK_FINISHED = 'GetMemberTaskFinished';
    public const GET_REWARDS = 'GetRewards';
    public const GET_LEVEL_TASKS = 'GetLevelTasks';
    public const GET_MEMBER_LEVEL_TASKS = 'GetMemberLevelTasks';

    //会员权益,仅等级权益
    public const GET_MEMBER_LEVEL_BENEFITS = 'GetMemberLevelBenefits';

    //会员所有权益
    public const GET_MEMBER_BENEFITS = 'GetMemberBenefits';

    //黑名单配置
    public const GET_BLACK_LIST_CONFIGS = 'GetBlackListConfigs';
    public const SAVE_BLACK_LIST_CONFIG = 'SaveBlackListConfig';
    public const CREATE_BLACK_LIST = 'CreateBlackList';
    public const GET_BLACK_LISTS = 'GetBlackLists';
    public const DELETE_BLACK_LIST = 'DeleteBlackList';

    public const SHOW_BLACK_LIST_BY_PHONE = 'ShowBlackListByPhone';

    protected array $actionMap = [
        self::SHOW_BLACK_LIST_BY_PHONE => [
            ShowBlackListByPhoneRequest::class,
            ShowBlackListByPhoneResponse::class
        ],
        self::DELETE_BLACK_LIST => [
            DeleteBlackListRequest::class,
            Response::class
        ],
        self::GET_BLACK_LISTS => [
            GetBlackListsRequest::class,
            GetBlackListsResponse::class
        ],
        self::CREATE_BLACK_LIST => [
            CreateBlackListRequest::class,
            Response::class
        ],
        self::SAVE_BLACK_LIST_CONFIG => [
            SaveBlackListConfigRequest::class,
            Response::class
        ],
        self::GET_BLACK_LIST_CONFIGS => [
            PBEmpty::class,
            GetBlackListConfigsResponse::class
        ],
        self::GET_MEMBER_LEVEL_TASKS => [
            GetMemberLevelTasksRequest::class,
            GetMemberLevelTasksResponse::class
        ],
        self::GET_LEVEL_TASKS => [
            GetLevelTasksRequest::class,
            GetLevelTasksResponse::class
        ],
        self::GET_MEMBER_BENEFITS => [
            GetMemberBenefitsRequest::class,
            GetMemberBenefitsResponse::class
        ],
        self::UPDATE_TASK_STATUS => [
            UpdateTaskStatusRequest::class,
            Response::class,
        ],
        self::SHOW_TASK => [
            ShowTaskRequest::class,
            ShowTaskResponse::class,
        ],
        self::GET_TASKS => [
            GetTasksRequest::class,
            GetTasksResponse::class,
        ],
        self::UPDATE_TASK => [
            UpdateTaskRequest::class,
            Response::class,
        ],
        self::CREATE_TASK => [
            CreateTaskRequest::class,
            Response::class,
        ],
        self::UPDATE_DISTRIBUTE_COUPON_STATUS => [
            UpdateDistributeCouponStatusRequest::class,
            Response::class,
        ],
        self::SHOW_DISTRIBUTE_COUPON => [
            ShowDistributeCouponRequest::class,
            ShowDistributeCouponResponse::class,
        ],
        self::GET_DISTRIBUTE_COUPONS => [
            GetDistributeCouponsRequest::class,
            GetDistributeCouponsResponse::class,
        ],
        self::CREATE_DISTRIBUTE_COUPON => [
            CreateDistributeCouponRequest::class,
            Response::class,
        ],
        self::UPDATE_DISTRIBUTE_COUPON => [
            UpdateDistributeCouponRequest::class,
            Response::class,
        ],
        self::CREATE_COUPON_STYLE => [
            CouponStyle::class,
            Response::class,
        ],
        self::SHOW_COUPON_STYLE => [
            PBEmpty::class,
            ShowCouponStyleResponse::class,
        ],
        self::SHOW_COUPON_CATEGORY => [
            ShowCouponCategoryRequest::class,
            ShowCouponCategoryResponse::class,
        ],
        self::GET_COUPON_CATEGORY => [
            GetCouponCategoriesRequest::class,
            GetCouponCategoriesResponse::class,
        ],
        self::UPDATE_COUPON_CATEGORY => [
            UpdateCouponCategoryRequest::class,
            Response::class,
        ],
        self::CREATE_COUPON_CATEGORY => [
            CreateCouponCategoryRequest::class,
            Response::class,
        ],
        self::UPDATE_COUPON_STATUS => [
            UpdateCouponStatusRequest::class,
            Response::class,
        ],
        self::SHOW_COUPON => [
            ShowCouponRequest::class,
            ShowCouponResponse::class,
        ],
        self::GET_COUPONS => [
            GetCouponsRequest::class,
            GetCouponsResponse::class,
        ],
        self::UPDATE_COUPON => [
            UpdateCouponRequest::class,
            Response::class,
        ],
        self::CREATE_COUPON => [
            CreateCouponRequest::class,
            Response::class,
        ],
        self::GET_DIVISIONS_AND_LEVELS => [
            PBEmpty::class,
            GetDivisionsAndLevelsResponse::class,
        ],
        self::DELETE_DIVISION => [
            DeleteDivisionRequest::class,
            Response::class,
        ],
        self::DELETE_LEVEL => [
            DeleteLevelRequest::class,
            Response::class,
        ],
        self::CREATE_ORDER_EXPRESS => [
            CreateOrderExpressRequest::class,
            Response::class,
        ],
        self::GET_GROWTH_BILLS => [
            GetGrowthBillsRequest::class,
            GetGrowthBillsResponse::class,
        ],
        self::GET_LEVELS_WITH_DIVISION => [
            GetLevelsWithDivisionRequest::class,
            GetLevelsWithDivisionResponse::class,
        ],
        self::SUBMIT_LEVEL => [
            SubmitDivisionLevelRequest::class,
            Response::class,
        ],
        self::CREATE_LEVEL => [
            CreateLevelRequest::class,
            CreateLevelResponse::class,
        ],
        self::UPDATE_LEVEL => [
            UpdateLevelRequest::class,
            Response::class,
        ],
        self::GET_LEVELS => [
            GetLevelsRequest::class,
            GetLevelsResponse::class,
        ],
        self::SHOW_LEVEL => [
            ShowLevelRequest::class,
            ShowLevelResponse::class,
        ],
        self::SHOW_DIVISIONS => [
            ShowDivisionRequest::class,
            ShowDivisionResponse::class,
        ],
        self::GET_DIVISIONS => [
            GetDivisionsRequest::class,
            GetDivisionsResponse::class,
        ],
        self::UPDATE_DIVISION => [
            UpdateDivisionRequest::class,
            Response::class,
        ],
        self::CREATE_DIVISION => [
            CreateDivisionRequest::class,
            CreateDivisionResponse::class,
        ],
        self::SHOW_BENEFIT => [
            ShowBenefitRequest::class,
            ShowBenefitResponse::class,
        ],
        self::UPDATE_BENEFITS_STATUS => [
            UpdateBenefitStatusRequest::class,
            Response::class,
        ],
        self::GET_BENEFITS => [
            GetBenefitsRequest::class,
            GetBenefitsResponse::class,
        ],
        self::UPDATE_BENEFITS => [
            UpdateBenefitRequest::class,
            Response::class,
        ],
        self::CREATE_BENEFITS => [
            CreateBenefitRequest::class,
            Response::class,
        ],
        self::UPDATE_PRODUCT_STATUS => [
            UpdateProductStatusRequest::class,
            Response::class,
        ],
        self::SHOW_POINT_CONFIG => [
            ShowPointConfigRequest::class,
            ShowPointConfigResponse::class,
        ],
        self::SAVE_POINT_CONFIG => [
            SavePointConfigRequest::class,
            Response::class,
        ],
        self::MEMBER_GAIN_POINT => [
            MemberGainPointRequest::class,
            Response::class,
        ],
        self::MEMBER_GAIN_POINT_BY_RULE => [
            MemberGainPointByRuleRequest::class,
            Response::class
        ],
        self::MEMBER_USE_POINT => [
            MemberUsePointRequest::class,
            MemberUsePointResponse::class
        ],
        self::MEMBER_USE_POINT_BY_RULE => [
            MemberUsePointByRuleRequest::class,
            Response::class
        ],
        self::REMOVE_MEMBER_POINT_BILL_LOCK => [
            RemoveMemberPointBillLockRequest::class,
            Response::class
        ],
        self::GET_MEMBER_EXTENSION => [
            GetMemberExtensionRequest::class,
            GetMemberExtensionResponse::class
        ],
        self::BATCH_GET_MEMBER_EXTENSION => [
            BatchGetMemberExtensionRequest::class,
            BatchGetMemberExtensionResponse::class
        ],
        self::GET_MEMBER_POINT_BILLS => [
            GetMemberPointBillsRequest::class,
            GetMemberPointBillsResponse::class
        ],
        self::CREATE_PRODUCT_CATEGORY => [
            CreateProductCategoryRequest::class,
            Response::class
        ],
        self::UPDATE_PRODUCT_CATEGORY => [
            UpdateProductCategoryRequest::class,
            Response::class
        ],
        self::UPDATE_PRODUCT_CATEGORY_STATUS => [
            UpdateStatusRequest::class,
            Response::class
        ],
        self::LIST_PRODUCT_CATEGORY => [
            ListProductCategoryRequest::class,
            ListProductCategoryResponse::class
        ],
        self::SECOND_LEVEL_SORT_PRODUCT_CATEGORY => [
            PBEmpty::class,
            SecondLevelSortProductCategoryResponse::class
        ],
        self::CREATE_PRODUCT => [
            CreateProductRequest::class,
            CreateProductResponse::class,
        ],
        self::GET_PRODUCTS => [
            GetProductsRequest::class,
            GetProductsResponse::class,
        ],
        self::UPDATE_PRODUCT_QUANTITY => [
            UpdateProductQuantityRequest::class,
            Response::class,
        ],
        self::UPDATE_PRODUCT => [
            UpdateProductRequest::class,
            Response::class,
        ],
        self::GET_PRODUCT_QUANTITY_BILLS => [
            GetProductQuantityBillsRequest::class,
            GetProductQuantityBillsResponse::class
        ],
        self::SHOW_PRODUCT_ORDER => [
            ShowProductOrderRequest::class,
            ShowProductOrderResponse::class,
        ],
        self::CREATE_PRODUCT_ORDER => [
            CreateProductOrderRequest::class,
            CreateProductOrderResponse::class,
        ],
        self::GET_PRODUCT_ORDERS => [
            GetProductOrdersRequest::class,
            GetProductOrdersResponse::class,
        ],
        self::UPDATE_PRODUCT_ORDER_STATUS => [
            UpdateProductOrderStatusRequest::class,
            Response::class
        ],
        self::SHOW_PRODUCT => [
            ShowProductRequest::class,
            ShowProductResponse::class,
        ],
        self::GET_MEMBER_COUPONS => [
            GetMemberCouponsRequest::class,
            GetMemberCouponsResponse::class,
        ],
        self::SHOW_MEMBER_COUPON => [
            ShowMemberCouponRequest::class,
            ShowMemberCouponResponse::class,
        ],
        self::GET_COLLECTABLE_COUPONS => [
            GetCollectableCouponsRequest::class,
            GetCollectableCouponsResponse::class
        ],
        self::CREATE_MEMBER_COUPON => [
            CreateMemberCouponRequest::class,
            Response::class
        ],
        self::GET_MEMBER_TASK_AVAILABLE => [
            GetMemberTaskAvailableRequest::class,
            GetMemberTaskAvailableResponse::class
        ],
        self::GET_MEMBER_TASK_FINISHED => [
            GetMemberTaskFinishedRequest::class,
            GetMemberTaskFinishedResponse::class
        ],
        self::GET_REWARDS => [
            GetRewardsRequest::class,
            Response::class
        ],
        self::GET_MEMBER_LEVEL_BENEFITS => [
            GetMemberLevelBenefitsRequest::class,
            GetBenefitsResponse::class
        ]
    ];

    protected function getBasePath()
    {
        return self::BASE_PATH;
    }
}
