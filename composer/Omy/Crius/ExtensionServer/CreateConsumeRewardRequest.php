<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/consume_reward.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.CreateConsumeRewardRequest</code>
 */
class CreateConsumeRewardRequest extends \Google\Protobuf\Internal\Message
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
     * Generated from protobuf field <code>repeated .memberExtension.CreateConsumeRewardRequest.Consume consumes = 3;</code>
     */
    private $consumes;
    /**
     *订单号
     *
     * Generated from protobuf field <code>string code = 4;</code>
     */
    protected $code = '';
    /**
     * Generated from protobuf field <code>string describe = 5;</code>
     */
    protected $describe = '';
    /**
     * Generated from protobuf field <code>bool point_is_lock = 6;</code>
     */
    protected $point_is_lock = false;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $branch_id
     *     @type string $member_id
     *     @type \Omy\Crius\ExtensionServer\CreateConsumeRewardRequest\Consume[]|\Google\Protobuf\Internal\RepeatedField $consumes
     *     @type string $code
     *          订单号
     *     @type string $describe
     *     @type bool $point_is_lock
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\ConsumeReward::initOnce();
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
     * Generated from protobuf field <code>repeated .memberExtension.CreateConsumeRewardRequest.Consume consumes = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getConsumes()
    {
        return $this->consumes;
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.CreateConsumeRewardRequest.Consume consumes = 3;</code>
     * @param \Omy\Crius\ExtensionServer\CreateConsumeRewardRequest\Consume[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setConsumes($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\ExtensionServer\CreateConsumeRewardRequest\Consume::class);
        $this->consumes = $arr;

        return $this;
    }

    /**
     *订单号
     *
     * Generated from protobuf field <code>string code = 4;</code>
     * @return string
     */
    public function getCode()
    {
        return $this->code;
    }

    /**
     *订单号
     *
     * Generated from protobuf field <code>string code = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string describe = 5;</code>
     * @return string
     */
    public function getDescribe()
    {
        return $this->describe;
    }

    /**
     * Generated from protobuf field <code>string describe = 5;</code>
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
     * Generated from protobuf field <code>bool point_is_lock = 6;</code>
     * @return bool
     */
    public function getPointIsLock()
    {
        return $this->point_is_lock;
    }

    /**
     * Generated from protobuf field <code>bool point_is_lock = 6;</code>
     * @param bool $var
     * @return $this
     */
    public function setPointIsLock($var)
    {
        GPBUtil::checkBool($var);
        $this->point_is_lock = $var;

        return $this;
    }

}

