<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/member_coupon.proto

namespace Omy\Crius\ExtensionServer\CollectableCoupon;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.CollectableCoupon.ValidTimeInterval</code>
 */
class ValidTimeInterval extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 week_day = 1;</code>
     */
    protected $week_day = 0;
    /**
     * Generated from protobuf field <code>string begin = 2;</code>
     */
    protected $begin = '';
    /**
     * Generated from protobuf field <code>string end = 3;</code>
     */
    protected $end = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $week_day
     *     @type string $begin
     *     @type string $end
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\MemberCoupon::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 week_day = 1;</code>
     * @return int
     */
    public function getWeekDay()
    {
        return $this->week_day;
    }

    /**
     * Generated from protobuf field <code>int32 week_day = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setWeekDay($var)
    {
        GPBUtil::checkInt32($var);
        $this->week_day = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string begin = 2;</code>
     * @return string
     */
    public function getBegin()
    {
        return $this->begin;
    }

    /**
     * Generated from protobuf field <code>string begin = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setBegin($var)
    {
        GPBUtil::checkString($var, True);
        $this->begin = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string end = 3;</code>
     * @return string
     */
    public function getEnd()
    {
        return $this->end;
    }

    /**
     * Generated from protobuf field <code>string end = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setEnd($var)
    {
        GPBUtil::checkString($var, True);
        $this->end = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(ValidTimeInterval::class, \Omy\Crius\ExtensionServer\CollectableCoupon_ValidTimeInterval::class);

