<?php

declare(strict_types=1);

namespace Omy\MicroServers\Client;

use Omy\Crius\MemberPrivateServer\CreatePromotionOptionRequest;
use Omy\Crius\MemberPrivateServer\CreatePromotionRequest;
use Omy\Crius\MemberPrivateServer\CreatePromotionResponse;
use Omy\Crius\MemberPrivateServer\CreateStaffShiftRequest;
use Omy\Crius\MemberPrivateServer\CreateStaffShiftResponse;
use Omy\Crius\MemberPrivateServer\GetStaffShiftTimeRequest;
use Omy\Crius\MemberPrivateServer\GetStaffShiftTimeResponse;
use Omy\Crius\MemberPrivateServer\ListPromotionOptionRequest;
use Omy\Crius\MemberPrivateServer\ListPromotionOptionResponse;
use Omy\Crius\MemberPrivateServer\ListPromotionRequest;
use Omy\Crius\MemberPrivateServer\ListPromotionResponse;
use Omy\Crius\MemberPrivateServer\ListStaffShiftRequest;
use Omy\Crius\MemberPrivateServer\ListStaffShiftResponse;
use Omy\Crius\MemberPrivateServer\OptionResponse;
use Omy\Crius\MemberPrivateServer\PromotionResponse;
use Omy\Crius\MemberPrivateServer\ShowPromotionOptionRequest;
use Omy\Crius\MemberPrivateServer\ShowPromotionOptionResponse;
use Omy\Crius\MemberPrivateServer\ShowPromotionRequest;
use Omy\Crius\MemberPrivateServer\ShowPromotionResponse;
use Omy\Crius\MemberPrivateServer\UpdatePromotionOptionRequest;
use Omy\Crius\MemberPrivateServer\UpdatePromotionRequest;
use Omy\Crius\MemberPrivateServer\UpdateStatusRequest;
use Omy\Crius\MemberPrivateServer\UpdateStatusResponse;

class MemberPrivateClient extends GrpcClient
{
    private const BASE_PATH = '/memberPrivate.MemberPrivateServer/';

    // 优惠方案组列表
    public const LIST_PROMOTION = 'ListPromotion';

    // 创建优惠方案组
    public const CREATE_PROMOTION = 'CreatePromotion';

    // 修改优惠方案组
    public const UPDATE_PROMOTION = 'UpdatePromotion';

    // 修改优惠方案组状态
    public const UPDATE_PROMOTION_STATUS = 'UpdatePromotionStatus';

    // 优惠方案组详情
    public const SHOW_PROMOTION = 'ShowPromotion';

    // 创建优惠方案
    public const CREATE_OPTION = 'CreatePromotionOption';

    // 优惠方案列表
    public const LIST_OPTION = 'ListPromotionOption';

    // 修改优惠方案
    public const UPDATE_OPTION = 'UpdatePromotionOption';

    // 优惠方案详情
    public const SHOW_OPTION = 'ShowPromotionOption';

    // 修改优惠方案状态
    public const UPDATE_OPTION_STATUS = 'UpdatePromotionOptionStatus';

    public const CREATE_STAFF_SHIFT = 'CreateStaffShift';

    public const LIST_STAFF_SHIFT = 'ListStaffShift';

    public const GET_STAFF_SHIFT_TIME = 'GetStaffShiftTime';

    protected array $actionMap = [
        self::GET_STAFF_SHIFT_TIME => [
            GetStaffShiftTimeRequest::class,
            GetStaffShiftTimeResponse::class
        ],
        self::LIST_PROMOTION => [
            ListPromotionRequest::class,
            ListPromotionResponse::class
        ],
        self::CREATE_PROMOTION => [
            CreatePromotionRequest::class,
            CreatePromotionResponse::class
        ],
        self::UPDATE_PROMOTION => [
            UpdatePromotionRequest::class,
            PromotionResponse::class
        ],
        self::UPDATE_PROMOTION_STATUS => [
            UpdateStatusRequest::class,
            UpdateStatusResponse::class
        ],
        self::SHOW_PROMOTION => [
            ShowPromotionRequest::class,
            ShowPromotionResponse::class
        ],
        self::CREATE_OPTION => [
            CreatePromotionOptionRequest::class,
            OptionResponse::class
        ],
        self::LIST_OPTION => [
            ListPromotionOptionRequest::class,
            ListPromotionOptionResponse::class
        ],
        self::UPDATE_OPTION => [
            UpdatePromotionOptionRequest::class,
            OptionResponse::class
        ],
        self::SHOW_OPTION => [
            ShowPromotionOptionRequest::class,
            ShowPromotionOptionResponse::class
        ],
        self::UPDATE_OPTION_STATUS => [
            UpdateStatusRequest::class,
            UpdateStatusResponse::class
        ],
        self::CREATE_STAFF_SHIFT => [
            CreateStaffShiftRequest::class,
            CreateStaffShiftResponse::class
        ],
        self::LIST_STAFF_SHIFT => [
            ListStaffShiftRequest::class,
            ListStaffShiftResponse::class
        ],
    ];

    protected function getBasePath()
    {
        return self::BASE_PATH;
    }
}
