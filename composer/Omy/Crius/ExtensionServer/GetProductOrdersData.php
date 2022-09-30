<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/product_order.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.GetProductOrdersData</code>
 */
class GetProductOrdersData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .memberExtension.ProductOrderWithProductInfo orders = 1;</code>
     */
    private $orders;
    /**
     * Generated from protobuf field <code>int32 total = 2;</code>
     */
    protected $total = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\ExtensionServer\ProductOrderWithProductInfo[]|\Google\Protobuf\Internal\RepeatedField $orders
     *     @type int $total
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\ProductOrder::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.ProductOrderWithProductInfo orders = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getOrders()
    {
        return $this->orders;
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.ProductOrderWithProductInfo orders = 1;</code>
     * @param \Omy\Crius\ExtensionServer\ProductOrderWithProductInfo[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setOrders($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\ExtensionServer\ProductOrderWithProductInfo::class);
        $this->orders = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 total = 2;</code>
     * @return int
     */
    public function getTotal()
    {
        return $this->total;
    }

    /**
     * Generated from protobuf field <code>int32 total = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setTotal($var)
    {
        GPBUtil::checkInt32($var);
        $this->total = $var;

        return $this;
    }

}
