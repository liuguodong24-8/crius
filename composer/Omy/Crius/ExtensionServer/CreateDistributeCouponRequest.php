<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/distribute_coupon.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.CreateDistributeCouponRequest</code>
 */
class CreateDistributeCouponRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.memberExtension.DistributeCoupon distribution_coupon = 1;</code>
     */
    protected $distribution_coupon = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\ExtensionServer\DistributeCoupon $distribution_coupon
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\DistributeCoupon::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.memberExtension.DistributeCoupon distribution_coupon = 1;</code>
     * @return \Omy\Crius\ExtensionServer\DistributeCoupon|null
     */
    public function getDistributionCoupon()
    {
        return $this->distribution_coupon;
    }

    public function hasDistributionCoupon()
    {
        return isset($this->distribution_coupon);
    }

    public function clearDistributionCoupon()
    {
        unset($this->distribution_coupon);
    }

    /**
     * Generated from protobuf field <code>.memberExtension.DistributeCoupon distribution_coupon = 1;</code>
     * @param \Omy\Crius\ExtensionServer\DistributeCoupon $var
     * @return $this
     */
    public function setDistributionCoupon($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\ExtensionServer\DistributeCoupon::class);
        $this->distribution_coupon = $var;

        return $this;
    }

}

