<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/member.proto

namespace Omy\Crius\ExtensionServer\GetMemberPointBillsResponse;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.GetMemberPointBillsResponse.Bill</code>
 */
class Bill extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>int64 code = 2;</code>
     */
    protected $code = 0;
    /**
     * Generated from protobuf field <code>string category = 3;</code>
     */
    protected $category = '';
    /**
     * Generated from protobuf field <code>string way = 4;</code>
     */
    protected $way = '';
    /**
     * Generated from protobuf field <code>string way_desc = 5;</code>
     */
    protected $way_desc = '';
    /**
     * Generated from protobuf field <code>string describe = 6;</code>
     */
    protected $describe = '';
    /**
     * Generated from protobuf field <code>int64 point = 7;</code>
     */
    protected $point = 0;
    /**
     * Generated from protobuf field <code>bool is_lock = 8;</code>
     */
    protected $is_lock = false;
    /**
     * Generated from protobuf field <code>int64 created_at = 9;</code>
     */
    protected $created_at = 0;
    /**
     * Generated from protobuf field <code>int64 expire_at = 10;</code>
     */
    protected $expire_at = 0;
    /**
     * Generated from protobuf field <code>int64 Left = 11;</code>
     */
    protected $Left = 0;
    /**
     * Generated from protobuf field <code>string member_id = 12;</code>
     */
    protected $member_id = '';
    /**
     * Generated from protobuf field <code>string branch_id = 13;</code>
     */
    protected $branch_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type int|string $code
     *     @type string $category
     *     @type string $way
     *     @type string $way_desc
     *     @type string $describe
     *     @type int|string $point
     *     @type bool $is_lock
     *     @type int|string $created_at
     *     @type int|string $expire_at
     *     @type int|string $Left
     *     @type string $member_id
     *     @type string $branch_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\Member::initOnce();
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
     * Generated from protobuf field <code>int64 code = 2;</code>
     * @return int|string
     */
    public function getCode()
    {
        return $this->code;
    }

    /**
     * Generated from protobuf field <code>int64 code = 2;</code>
     * @param int|string $var
     * @return $this
     */
    public function setCode($var)
    {
        GPBUtil::checkInt64($var);
        $this->code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string category = 3;</code>
     * @return string
     */
    public function getCategory()
    {
        return $this->category;
    }

    /**
     * Generated from protobuf field <code>string category = 3;</code>
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
     * Generated from protobuf field <code>string way_desc = 5;</code>
     * @return string
     */
    public function getWayDesc()
    {
        return $this->way_desc;
    }

    /**
     * Generated from protobuf field <code>string way_desc = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setWayDesc($var)
    {
        GPBUtil::checkString($var, True);
        $this->way_desc = $var;

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
     * Generated from protobuf field <code>int64 point = 7;</code>
     * @return int|string
     */
    public function getPoint()
    {
        return $this->point;
    }

    /**
     * Generated from protobuf field <code>int64 point = 7;</code>
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
     * Generated from protobuf field <code>bool is_lock = 8;</code>
     * @return bool
     */
    public function getIsLock()
    {
        return $this->is_lock;
    }

    /**
     * Generated from protobuf field <code>bool is_lock = 8;</code>
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
     * Generated from protobuf field <code>int64 created_at = 9;</code>
     * @return int|string
     */
    public function getCreatedAt()
    {
        return $this->created_at;
    }

    /**
     * Generated from protobuf field <code>int64 created_at = 9;</code>
     * @param int|string $var
     * @return $this
     */
    public function setCreatedAt($var)
    {
        GPBUtil::checkInt64($var);
        $this->created_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 expire_at = 10;</code>
     * @return int|string
     */
    public function getExpireAt()
    {
        return $this->expire_at;
    }

    /**
     * Generated from protobuf field <code>int64 expire_at = 10;</code>
     * @param int|string $var
     * @return $this
     */
    public function setExpireAt($var)
    {
        GPBUtil::checkInt64($var);
        $this->expire_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 Left = 11;</code>
     * @return int|string
     */
    public function getLeft()
    {
        return $this->Left;
    }

    /**
     * Generated from protobuf field <code>int64 Left = 11;</code>
     * @param int|string $var
     * @return $this
     */
    public function setLeft($var)
    {
        GPBUtil::checkInt64($var);
        $this->Left = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string member_id = 12;</code>
     * @return string
     */
    public function getMemberId()
    {
        return $this->member_id;
    }

    /**
     * Generated from protobuf field <code>string member_id = 12;</code>
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
     * Generated from protobuf field <code>string branch_id = 13;</code>
     * @return string
     */
    public function getBranchId()
    {
        return $this->branch_id;
    }

    /**
     * Generated from protobuf field <code>string branch_id = 13;</code>
     * @param string $var
     * @return $this
     */
    public function setBranchId($var)
    {
        GPBUtil::checkString($var, True);
        $this->branch_id = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Bill::class, \Omy\Crius\ExtensionServer\GetMemberPointBillsResponse_Bill::class);
