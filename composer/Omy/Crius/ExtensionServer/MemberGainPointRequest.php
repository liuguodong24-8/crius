<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/member.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * MemberGainPointRequest 直接获取积分
 *
 * Generated from protobuf message <code>memberExtension.MemberGainPointRequest</code>
 */
class MemberGainPointRequest extends \Google\Protobuf\Internal\Message
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
     * Generated from protobuf field <code>int64 point = 3;</code>
     */
    protected $point = 0;
    /**
     * Generated from protobuf field <code>string way = 4;</code>
     */
    protected $way = '';
    /**
     * Generated from protobuf field <code>string describe = 5;</code>
     */
    protected $describe = '';
    /**
     * Generated from protobuf field <code>bool is_lock = 6;</code>
     */
    protected $is_lock = false;
    /**
     * Generated from protobuf field <code>int64 expire_at = 7;</code>
     */
    protected $expire_at = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $branch_id
     *     @type string $member_id
     *     @type int|string $point
     *     @type string $way
     *     @type string $describe
     *     @type bool $is_lock
     *     @type int|string $expire_at
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
     * Generated from protobuf field <code>int64 point = 3;</code>
     * @return int|string
     */
    public function getPoint()
    {
        return $this->point;
    }

    /**
     * Generated from protobuf field <code>int64 point = 3;</code>
     * @param int|string $var
     * @return $this
     */
    public function setPoint($var)
    {
        GPBUtil::checkInt64($var);
        $this->point = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string way = 4;</code>
     * @return string
     */
    public function getWay()
    {
        return $this->way;
    }

    /**
     * Generated from protobuf field <code>string way = 4;</code>
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
     * Generated from protobuf field <code>bool is_lock = 6;</code>
     * @return bool
     */
    public function getIsLock()
    {
        return $this->is_lock;
    }

    /**
     * Generated from protobuf field <code>bool is_lock = 6;</code>
     * @param bool $var
     * @return $this
     */
    public function setIsLock($var)
    {
        GPBUtil::checkBool($var);
        $this->is_lock = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 expire_at = 7;</code>
     * @return int|string
     */
    public function getExpireAt()
    {
        return $this->expire_at;
    }

    /**
     * Generated from protobuf field <code>int64 expire_at = 7;</code>
     * @param int|string $var
     * @return $this
     */
    public function setExpireAt($var)
    {
        GPBUtil::checkInt64($var);
        $this->expire_at = $var;

        return $this;
    }

}
