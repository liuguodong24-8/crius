<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/assist/assist.proto

namespace Omy\Crius\AssistServer\PosBillOrderRequest;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>assist.PosBillOrderRequest.OrderInfo</code>
 */
class OrderInfo extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * 下单类型
     *
     * Generated from protobuf field <code>.assist.PosBillOrderRequest.OrderCategory order_category = 2;</code>
     */
    protected $order_category = 0;
    /**
     * 下单数量
     *
     * Generated from protobuf field <code>int32 order_number = 3;</code>
     */
    protected $order_number = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type int $order_category
     *           下单类型
     *     @type int $order_number
     *           下单数量
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Assist\Assist::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string id = 1;</code>
     * @return string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Generated from protobuf field <code>string id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkString($var, True);
        $this->id = $var;

        return $this;
    }

    /**
     * 下单类型
     *
     * Generated from protobuf field <code>.assist.PosBillOrderRequest.OrderCategory order_category = 2;</code>
     * @return int
     */
    public function getOrderCategory()
    {
        return $this->order_category;
    }

    /**
     * 下单类型
     *
     * Generated from protobuf field <code>.assist.PosBillOrderRequest.OrderCategory order_category = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setOrderCategory($var)
    {
        GPBUtil::checkEnum($var, \Omy\Crius\AssistServer\PosBillOrderRequest\OrderCategory::class);
        $this->order_category = $var;

        return $this;
    }

    /**
     * 下单数量
     *
     * Generated from protobuf field <code>int32 order_number = 3;</code>
     * @return int
     */
    public function getOrderNumber()
    {
        return $this->order_number;
    }

    /**
     * 下单数量
     *
     * Generated from protobuf field <code>int32 order_number = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setOrderNumber($var)
    {
        GPBUtil::checkInt32($var);
        $this->order_number = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(OrderInfo::class, \Omy\Crius\AssistServer\PosBillOrderRequest_OrderInfo::class);

