<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/member.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * MemberUsePointByRuleRequest 用户通过规则消费积分
 *
 * Generated from protobuf message <code>memberExtension.MemberUsePointByRuleRequest</code>
 */
class MemberUsePointByRuleRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string branch_id = 1;</code>
     */
    protected $branch_id = '';
    /**
     * Generated from protobuf field <code>string member_id = 2;</code>
     */
    protected $member_id = '';
    /**
     * Generated from protobuf field <code>string point_category_code = 3;</code>
     */
    protected $point_category_code = '';
    /**
     * Generated from protobuf field <code>int64 fee = 4;</code>
     */
    protected $fee = 0;
    /**
     * Generated from protobuf field <code>string way = 5;</code>
     */
    protected $way = '';
    /**
     * Generated from protobuf field <code>string describe = 6;</code>
     */
    protected $describe = '';
    /**
     * Generated from protobuf field <code>bool is_lock = 7;</code>
     */
    protected $is_lock = false;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $branch_id
     *     @type string $member_id
     *     @type string $point_category_code
     *     @type int|string $fee
     *     @type string $way
     *     @type string $describe
     *     @type bool $is_lock
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\Member::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string branch_id = 1;</code>
     * @return string
     */
    public function getBranchId()
    {
        return $this->branch_id;
    }

    /**
     * Generated from protobuf field <code>string branch_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setBranchId($var)
    {
        GPBUtil::checkString($var, True);
        $this->branch_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string member_id = 2;</code>
     * @return string
     */
    public function getMemberId()
    {
        return $this->member_id;
    }

    /**
     * Generated from protobuf field <code>string member_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setMemberId($var)
    {
        GPBUtil::checkString($var, True);
        $this->member_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string point_category_code = 3;</code>
     * @return string
     */
    public function getPointCategoryCode()
    {
        return $this->point_category_code;
    }

    /**
     * Generated from protobuf field <code>string point_category_code = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setPointCategoryCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->point_category_code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 fee = 4;</code>
     * @return int|string
     */
    public function getFee()
    {
        return $this->fee;
    }

    /**
     * Generated from protobuf field <code>int64 fee = 4;</code>
     * @param int|string $var
     * @return $this
     */
    public function setFee($var)
    {
        GPBUtil::checkInt64($var);
        $this->fee = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string way = 5;</code>
     * @return string
     */
    public function getWay()
    {
        return $this->way;
    }

    /**
     * Generated from protobuf field <code>string way = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setWay($var)
    {
        GPBUtil::checkString($var, True);
        $this->way = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string describe = 6;</code>
     * @return string
     */
    public function getDescribe()
    {
        return $this->describe;
    }

    /**
     * Generated from protobuf field <code>string describe = 6;</code>
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
     * Generated from protobuf field <code>bool is_lock = 7;</code>
     * @return bool
     */
    public function getIsLock()
    {
        return $this->is_lock;
    }

    /**
     * Generated from protobuf field <code>bool is_lock = 7;</code>
     * @param bool $var
     * @return $this
     */
    public function setIsLock($var)
    {
        GPBUtil::checkBool($var);
        $this->is_lock = $var;

        return $this;
    }

}

