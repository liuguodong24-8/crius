<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/task.proto

namespace Omy\Crius\ExtensionServer\Task;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.Task.Reward</code>
 */
class Reward extends \Google\Protobuf\Internal\Message
{
    /**
     * 积分
     *
     * Generated from protobuf field <code>int32 point = 1;</code>
     */
    protected $point = 0;
    /**
     * 积分过期时间
     *
     * Generated from protobuf field <code>int32 point_expire = 2;</code>
     */
    protected $point_expire = 0;
    /**
     * 成长值
     *
     * Generated from protobuf field <code>int32 growth = 3;</code>
     */
    protected $growth = 0;
    /**
     * 成长值过期时间
     *
     * Generated from protobuf field <code>int32 growth_expire = 4;</code>
     */
    protected $growth_expire = 0;
    /**
     * 优惠券
     *
     * Generated from protobuf field <code>repeated string coupon_ids = 5;</code>
     */
    private $coupon_ids;
    /**
     * 奖励类型
     *
     * Generated from protobuf field <code>string category = 6;</code>
     */
    protected $category = '';
    /**
     * 权益id
     *
     * Generated from protobuf field <code>repeated string benefit_ids = 7;</code>
     */
    private $benefit_ids;
    /**
     * 权益过期时间
     *
     * Generated from protobuf field <code>int32 benefit_expire = 8;</code>
     */
    protected $benefit_expire = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $point
     *           积分
     *     @type int $point_expire
     *           积分过期时间
     *     @type int $growth
     *           成长值
     *     @type int $growth_expire
     *           成长值过期时间
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $coupon_ids
     *           优惠券
     *     @type string $category
     *           奖励类型
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $benefit_ids
     *           权益id
     *     @type int $benefit_expire
     *           权益过期时间
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\Task::initOnce();
        parent::__construct($data);
    }

    /**
     * 积分
     *
     * Generated from protobuf field <code>int32 point = 1;</code>
     * @return int
     */
    public function getPoint()
    {
        return $this->point;
    }

    /**
     * 积分
     *
     * Generated from protobuf field <code>int32 point = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setPoint($var)
    {
        GPBUtil::checkInt32($var);
        $this->point = $var;

        return $this;
    }

    /**
     * 积分过期时间
     *
     * Generated from protobuf field <code>int32 point_expire = 2;</code>
     * @return int
     */
    public function getPointExpire()
    {
        return $this->point_expire;
    }

    /**
     * 积分过期时间
     *
     * Generated from protobuf field <code>int32 point_expire = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setPointExpire($var)
    {
        GPBUtil::checkInt32($var);
        $this->point_expire = $var;

        return $this;
    }

    /**
     * 成长值
     *
     * Generated from protobuf field <code>int32 growth = 3;</code>
     * @return int
     */
    public function getGrowth()
    {
        return $this->growth;
    }

    /**
     * 成长值
     *
     * Generated from protobuf field <code>int32 growth = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setGrowth($var)
    {
        GPBUtil::checkInt32($var);
        $this->growth = $var;

        return $this;
    }

    /**
     * 成长值过期时间
     *
     * Generated from protobuf field <code>int32 growth_expire = 4;</code>
     * @return int
     */
    public function getGrowthExpire()
    {
        return $this->growth_expire;
    }

    /**
     * 成长值过期时间
     *
     * Generated from protobuf field <code>int32 growth_expire = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setGrowthExpire($var)
    {
        GPBUtil::checkInt32($var);
        $this->growth_expire = $var;

        return $this;
    }

    /**
     * 优惠券
     *
     * Generated from protobuf field <code>repeated string coupon_ids = 5;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getCouponIds()
    {
        return $this->coupon_ids;
    }

    /**
     * 优惠券
     *
     * Generated from protobuf field <code>repeated string coupon_ids = 5;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setCouponIds($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->coupon_ids = $arr;

        return $this;
    }

    /**
     * 奖励类型
     *
     * Generated from protobuf field <code>string category = 6;</code>
     * @return string
     */
    public function getCategory()
    {
        return $this->category;
    }

    /**
     * 奖励类型
     *
     * Generated from protobuf field <code>string category = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setCategory($var)
    {
        GPBUtil::checkString($var, True);
        $this->category = $var;

        return $this;
    }

    /**
     * 权益id
     *
     * Generated from protobuf field <code>repeated string benefit_ids = 7;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getBenefitIds()
    {
        return $this->benefit_ids;
    }

    /**
     * 权益id
     *
     * Generated from protobuf field <code>repeated string benefit_ids = 7;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setBenefitIds($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->benefit_ids = $arr;

        return $this;
    }

    /**
     * 权益过期时间
     *
     * Generated from protobuf field <code>int32 benefit_expire = 8;</code>
     * @return int
     */
    public function getBenefitExpire()
    {
        return $this->benefit_expire;
    }

    /**
     * 权益过期时间
     *
     * Generated from protobuf field <code>int32 benefit_expire = 8;</code>
     * @param int $var
     * @return $this
     */
    public function setBenefitExpire($var)
    {
        GPBUtil::checkInt32($var);
        $this->benefit_expire = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Reward::class, \Omy\Crius\ExtensionServer\Task_Reward::class);

