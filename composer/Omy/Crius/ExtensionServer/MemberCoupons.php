<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/member_coupon.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.MemberCoupons</code>
 */
class MemberCoupons extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .memberExtension.MemberCoupon coupons = 1;</code>
     */
    private $coupons;
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
     *     @type \Omy\Crius\ExtensionServer\MemberCoupon[]|\Google\Protobuf\Internal\RepeatedField $coupons
     *     @type int $total
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\MemberCoupon::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.MemberCoupon coupons = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getCoupons()
    {
        return $this->coupons;
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.MemberCoupon coupons = 1;</code>
     * @param \Omy\Crius\ExtensionServer\MemberCoupon[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setCoupons($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\ExtensionServer\MemberCoupon::class);
        $this->coupons = $arr;

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

