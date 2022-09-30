<?php


namespace Omy\MicroServers\Client;


use Omy\Crius\AssistServer\CancelOrderRequest;
use Omy\Crius\AssistServer\CreateOrderRequest;
use Omy\Crius\AssistServer\DeleteOrderRequest;
use Omy\Crius\AssistServer\GetOrdersRequest;
use Omy\Crius\AssistServer\GetOrdersResponse;
use Omy\Crius\AssistServer\OrderEventResponse;
use Omy\Crius\AssistServer\PayOrderRequest;
use Omy\Crius\AssistServer\PosBillBandRequest;
use Omy\Crius\AssistServer\PosBillCancelBandRequest;
use Omy\Crius\AssistServer\PosBillCancelRequest;
use Omy\Crius\AssistServer\PosBillChangeRoomRequest;
use Omy\Crius\AssistServer\PosBillChangeValueRequest;
use Omy\Crius\AssistServer\PosBillDoneRequest;
use Omy\Crius\AssistServer\PosBillOrderRequest;
use Omy\Crius\AssistServer\PosBillPayRequest;
use Omy\Crius\AssistServer\PosBillPowerRequest;
use Omy\Crius\AssistServer\UpdateOrderRequest;

class AssistClient extends GrpcClient
{
    private const BASE_PATH = '/assist.AssistServer/';

    //pos开机
    public const POS_BILL_POWER = 'PosBillPower';
    //pos预约订单挂入
    public const POS_BILL_BAND = 'PosBillBand';
    //pos预约订单取消绑定
    public const POS_BILL_CANCEL_BAND = 'PosBillCancelBand';
    //pos下单
    public const POS_BILL_ORDER = 'PosBillOrder';
    //pos金额改变
    public const POS_BILL_CHANGE_VALUE = 'PosBillChangeValue';
    //pos转包
    public const POS_BILL_CHANGE_ROOM = 'PosBillChangeRoom';
    //pos支付
    public const POS_BILL_PAY = 'PosBillPay';
    //pos完成
    public const POS_BILL_DONE = 'PosBillDone';
    //pos取消
    public const POS_BILL_CANCEL = 'PosBillCancel';


    //创建订单
    public const CREATE_ORDER = 'CreateOrder';
    //更新订单
    public const UPDATE_ORDER = 'UpdateOrder';
    //删除订单
    public const DELETE_ORDER = 'DeleteOrder';
    //订单预约付款
    public const PAY_ORDER = 'PayOrder';
    //订单预约取消
    public const CANCEL_ORDER = 'CancelOrder';


    //订单列表
    public const GET_ORDERS = 'GetOrders';

    protected function getBasePath()
    {
        return self::BASE_PATH;
    }

    protected array $actionMap = [
        self::POS_BILL_POWER => [
            PosBillPowerRequest::class,
            OrderEventResponse::class
        ],
        self::POS_BILL_BAND => [
            PosBillBandRequest::class,
            OrderEventResponse::class
        ],
        self::POS_BILL_CANCEL_BAND => [
            PosBillCancelBandRequest::class,
            OrderEventResponse::class
        ],
        self::POS_BILL_ORDER => [
            PosBillOrderRequest::class,
            OrderEventResponse::class
        ],
        self::POS_BILL_CHANGE_VALUE => [
            PosBillChangeValueRequest::class,
            OrderEventResponse::class
        ],
        self::POS_BILL_CHANGE_ROOM => [
            PosBillChangeRoomRequest::class,
            OrderEventResponse::class
        ],
        self::POS_BILL_PAY => [
            PosBillPayRequest::class,
            OrderEventResponse::class
        ],
        self::POS_BILL_DONE => [
            PosBillDoneRequest::class,
            OrderEventResponse::class
        ],
        self::POS_BILL_CANCEL => [
            PosBillCancelRequest::class,
            OrderEventResponse::class
        ],
        self::CREATE_ORDER => [
            CreateOrderRequest::class,
            OrderEventResponse::class
        ],
        self::UPDATE_ORDER => [
            UpdateOrderRequest::class,
            OrderEventResponse::class
        ],
        self::DELETE_ORDER => [
            DeleteOrderRequest::class,
            OrderEventResponse::class
        ],
        self::PAY_ORDER => [
            PayOrderRequest::class,
            OrderEventResponse::class
        ],
        self::CANCEL_ORDER => [
            CancelOrderRequest::class,
            OrderEventResponse::class
        ],
        self::GET_ORDERS => [
            GetOrdersRequest::class,
            GetOrdersResponse::class
        ],
    ];
}