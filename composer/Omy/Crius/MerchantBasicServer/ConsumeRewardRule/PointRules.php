<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/branch.proto

namespace Omy\Crius\MerchantBasicServer\ConsumeRewardRule;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.ConsumeRewardRule.PointRules</code>
 */
class PointRules extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .merchantBasic.ConsumeRewardRule.PointRule rules = 1;</code>
     */
    private $rules;
    /**
     * Generated from protobuf field <code>int32 validity_day = 2;</code>
     */
    protected $validity_day = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\MerchantBasicServer\ConsumeRewardRule\PointRule[]|\Google\Protobuf\Internal\RepeatedField $rules
     *     @type int $validity_day
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\Branch::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .merchantBasic.ConsumeRewardRule.PointRule rules = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getRules()
    {
        return $this->rules;
    }

    /**
     * Generated from protobuf field <code>repeated .merchantBasic.ConsumeRewardRule.PointRule rules = 1;</code>
     * @param \Omy\Crius\MerchantBasicServer\ConsumeRewardRule\PointRule[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setRules($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MerchantBasicServer\ConsumeRewardRule\PointRule::class);
        $this->rules = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 validity_day = 2;</code>
     * @return int
     */
    public function getValidityDay()
    {
        return $this->validity_day;
    }

    /**
     * Generated from protobuf field <code>int32 validity_day = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setValidityDay($var)
    {
        GPBUtil::checkInt32($var);
        $this->validity_day = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(PointRules::class, \Omy\Crius\MerchantBasicServer\ConsumeRewardRule_PointRules::class);

