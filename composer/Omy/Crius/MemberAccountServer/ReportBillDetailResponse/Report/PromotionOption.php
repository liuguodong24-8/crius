<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-account/report.proto

namespace Omy\Crius\MemberAccountServer\ReportBillDetailResponse\Report;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberAccount.ReportBillDetailResponse.Report.PromotionOption</code>
 */
class PromotionOption extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string promotion_option_id = 1;</code>
     */
    protected $promotion_option_id = '';
    /**
     * Generated from protobuf field <code>string promotion_option_name = 2;</code>
     */
    protected $promotion_option_name = '';
    /**
     * Generated from protobuf field <code>int32 count = 3;</code>
     */
    protected $count = 0;
    /**
     * Generated from protobuf field <code>int32 recharge_value = 4;</code>
     */
    protected $recharge_value = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $promotion_option_id
     *     @type string $promotion_option_name
     *     @type int $count
     *     @type int $recharge_value
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberAccount\Report::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string promotion_option_id = 1;</code>
     * @return string
     */
    public function getPromotionOptionId()
    {
        return $this->promotion_option_id;
    }

    /**
     * Generated from protobuf field <code>string promotion_option_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setPromotionOptionId($var)
    {
        GPBUtil::checkString($var, True);
        $this->promotion_option_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string promotion_option_name = 2;</code>
     * @return string
     */
    public function getPromotionOptionName()
    {
        return $this->promotion_option_name;
    }

    /**
     * Generated from protobuf field <code>string promotion_option_name = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setPromotionOptionName($var)
    {
        GPBUtil::checkString($var, True);
        $this->promotion_option_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 count = 3;</code>
     * @return int
     */
    public function getCount()
    {
        return $this->count;
    }

    /**
     * Generated from protobuf field <code>int32 count = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setCount($var)
    {
        GPBUtil::checkInt32($var);
        $this->count = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 recharge_value = 4;</code>
     * @return int
     */
    public function getRechargeValue()
    {
        return $this->recharge_value;
    }

    /**
     * Generated from protobuf field <code>int32 recharge_value = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setRechargeValue($var)
    {
        GPBUtil::checkInt32($var);
        $this->recharge_value = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(PromotionOption::class, \Omy\Crius\MemberAccountServer\ReportBillDetailResponse_Report_PromotionOption::class);

