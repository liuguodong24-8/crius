<?php

declare(strict_types=1);

namespace Omy\MicroServers\Client;

use Omy\Crius\PaymentServer\WechatOrderQueryRequest;
use Omy\Crius\PaymentServer\WechatOrderQueryResponse;
use Omy\Crius\PaymentServer\WechatUnifiedOrderRequest;
use Omy\Crius\PaymentServer\WechatUnifiedOrderResponse;

class PaymentClient extends GrpcClient
{
    private const BASE_PATH = '/payment.PaymentServer/';

    // 微信下单
    public const WECHAT_UNIFIED_ORDER = 'WechatUnifiedOrder';

    public const WECHAT_ORDER_QUERY = 'WechatOrderQuery';

    protected array $actionMap = [
        self::WECHAT_UNIFIED_ORDER => [
            WechatUnifiedOrderRequest::class,
            WechatUnifiedOrderResponse::class,
        ],
        self::WECHAT_ORDER_QUERY => [
            WechatOrderQueryRequest::class,
            WechatOrderQueryResponse::class
        ]
    ];

    protected function getBasePath()
    {
        return self::BASE_PATH;
    }
}
