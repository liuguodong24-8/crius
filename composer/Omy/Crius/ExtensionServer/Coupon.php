<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/coupon.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.Coupon</code>
 */
class Coupon extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>string name = 2;</code>
     */
    protected $name = '';
    /**
     * Generated from protobuf field <code>string alias = 3;</code>
     */
    protected $alias = '';
    /**
     * Generated from protobuf field <code>string coupon_category_id = 4;</code>
     */
    protected $coupon_category_id = '';
    /**
     * Generated from protobuf field <code>int32 available_times = 5;</code>
     */
    protected $available_times = 0;
    /**
     * Generated from protobuf field <code>int32 collectable_times = 6;</code>
     */
    protected $collectable_times = 0;
    /**
     * Generated from protobuf field <code>string status = 7;</code>
     */
    protected $status = '';
    /**
     * Generated from protobuf field <code>repeated string distribution_way = 8;</code>
     */
    private $distribution_way;
    /**
     * Generated from protobuf field <code>repeated string available_items = 19;</code>
     */
    private $available_items;
    /**
     * Generated from protobuf field <code>string describe = 10;</code>
     */
    protected $describe = '';
    /**
     * Generated from protobuf field <code>.memberExtension.Coupon.Coupon coupon = 11;</code>
     */
    protected $coupon = null;
    /**
     * Generated from protobuf field <code>string coupon_category_first = 12;</code>
     */
    protected $coupon_category_first = '';
    /**
     * Generated from protobuf field <code>string coupon_category_name = 13;</code>
     */
    protected $coupon_category_name = '';
    /**
     * Generated from protobuf field <code>string style = 14;</code>
     */
    protected $style = '';
    /**
     * Generated from protobuf field <code>bool multi_available = 15;</code>
     */
    protected $multi_available = false;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $name
     *     @type string $alias
     *     @type string $coupon_category_id
     *     @type int $available_times
     *     @type int $collectable_times
     *     @type string $status
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $distribution_way
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $available_items
     *     @type string $describe
     *     @type \Omy\Crius\ExtensionServer\Coupon\Coupon $coupon
     *     @type string $coupon_category_first
     *     @type string $coupon_category_name
     *     @type string $style
     *     @type bool $multi_available
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\Coupon::initOnce();
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
     * Generated from protobuf field <code>string name = 2;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * Generated from protobuf field <code>string name = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string alias = 3;</code>
     * @return string
     */
    public function getAlias()
    {
        return $this->alias;
    }

    /**
     * Generated from protobuf field <code>string alias = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setAlias($var)
    {
        GPBUtil::checkString($var, True);
        $this->alias = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string coupon_category_id = 4;</code>
     * @return string
     */
    public function getCouponCategoryId()
    {
        return $this->coupon_category_id;
    }

    /**
     * Generated from protobuf field <code>string coupon_category_id = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setCouponCategoryId($var)
    {
        GPBUtil::checkString($var, True);
        $this->coupon_category_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 available_times = 5;</code>
     * @return int
     */
    public function getAvailableTimes()
    {
        return $this->available_times;
    }

    /**
     * Generated from protobuf field <code>int32 available_times = 5;</code>
     * @param int $var
     * @return $this
     */
    public function setAvailableTimes($var)
    {
        GPBUtil::checkInt32($var);
        $this->available_times = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 collectable_times = 6;</code>
     * @return int
     */
    public function getCollectableTimes()
    {
        return $this->collectable_times;
    }

    /**
     * Generated from protobuf field <code>int32 collectable_times = 6;</code>
     * @param int $var
     * @return $this
     */
    public function setCollectableTimes($var)
    {
        GPBUtil::checkInt32($var);
        $this->collectable_times = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string status = 7;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>string status = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkString($var, True);
        $this->status = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string distribution_way = 8;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getDistributionWay()
    {
        return $this->distribution_way;
    }

    /**
     * Generated from protobuf field <code>repeated string distribution_way = 8;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setDistributionWay($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->distribution_way = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string available_items = 19;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAvailableItems()
    {
        return $this->available_items;
    }

    /**
     * Generated from protobuf field <code>repeated string available_items = 19;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAvailableItems($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->available_items = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string describe = 10;</code>
     * @return string
     */
    public function getDescribe()
    {
        return $this->describe;
    }

    /**
     * Generated from protobuf field <code>string describe = 10;</code>
     * @param string $var
     * @return $this
     */
    public function setDescribe($var)
    {
        GPBUtil::checkString($var, True);
        $this->describe = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.memberExtension.Coupon.Coupon coupon = 11;</code>
     * @return \Omy\Crius\ExtensionServer\Coupon\Coupon|null
     */
    public function getCoupon()
    {
        return $this->coupon;
    }

    public function hasCoupon()
    {
        return isset($this->coupon);
    }

    public function clearCoupon()
    {
        unset($this->coupon);
    }

    /**
     * Generated from protobuf field <code>.memberExtension.Coupon.Coupon coupon = 11;</code>
     * @param \Omy\Crius\ExtensionServer\Coupon\Coupon $var
     * @return $this
     */
    public function setCoupon($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\ExtensionServer\Coupon\Coupon::class);
        $this->coupon = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string coupon_category_first = 12;</code>
     * @return string
     */
    public function getCouponCategoryFirst()
    {
        return $this->coupon_category_first;
    }

    /**
     * Generated from protobuf field <code>string coupon_category_first = 12;</code>
     * @param string $var
     * @return $this
     */
    public function setCouponCategoryFirst($var)
    {
        GPBUtil::checkString($var, True);
        $this->coupon_category_first = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string coupon_category_name = 13;</code>
     * @return string
     */
    public function getCouponCategoryName()
    {
        return $this->coupon_category_name;
    }

    /**
     * Generated from protobuf field <code>string coupon_category_name = 13;</code>
     * @param string $var
     * @return $this
     */
    public function setCouponCategoryName($var)
    {
        GPBUtil::checkString($var, True);
        $this->coupon_category_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string style = 14;</code>
     * @return string
     */
    public function getStyle()
    {
        return $this->style;
    }

    /**
     * Generated from protobuf field <code>string style = 14;</code>
     * @param string $var
     * @return $this
     */
    public function setStyle($var)
    {
        GPBUtil::checkString($var, True);
        $this->style = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool multi_available = 15;</code>
     * @return bool
     */
    public function getMultiAvailable()
    {
        return $this->multi_available;
    }

    /**
     * Generated from protobuf field <code>bool multi_available = 15;</code>
     * @param bool $var
     * @return $this
     */
    public function setMultiAvailable($var)
    {
        GPBUtil::checkBool($var);
        $this->multi_available = $var;

        return $this;
    }

}

