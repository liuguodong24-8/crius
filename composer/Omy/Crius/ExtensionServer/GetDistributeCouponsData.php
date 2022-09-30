<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/distribute_coupon.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.GetDistributeCouponsData</code>
 */
class GetDistributeCouponsData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 total = 1;</code>
     */
    protected $total = 0;
    /**
     * Generated from protobuf field <code>repeated .memberExtension.DistributeCoupon coupons = 2;</code>
     */
    private $coupons;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $total
     *     @type \Omy\Crius\ExtensionServer\DistributeCoupon[]|\Google\Protobuf\Internal\RepeatedField $coupons
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\DistributeCoupon::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 total = 1;</code>
     * @return int
     */
    public function getTotal()
    {
        return $this->total;
    }

    /**
     * Generated from protobuf field <code>int32 total = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setTotal($var)
    {
        GPBUtil::checkInt32($var);
        $this->total = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.DistributeCoupon coupons = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getCoupons()
    {
        return $this->coupons;
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.DistributeCoupon coupons = 2;</code>
     * @param \Omy\Crius\ExtensionServer\DistributeCoupon[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setCoupons($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\ExtensionServer\DistributeCoupon::class);
        $this->coupons = $arr;

        return $this;
    }

}

