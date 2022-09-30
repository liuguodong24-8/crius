<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/member.proto

namespace Omy\Crius\ExtensionServer\GetMemberPointBillsResponse;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.GetMemberPointBillsResponse.Data</code>
 */
class Data extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 total = 1;</code>
     */
    protected $total = 0;
    /**
     * Generated from protobuf field <code>int32 total_point = 2;</code>
     */
    protected $total_point = 0;
    /**
     * Generated from protobuf field <code>int32 total_lock_point = 3;</code>
     */
    protected $total_lock_point = 0;
    /**
     * Generated from protobuf field <code>repeated .memberExtension.GetMemberPointBillsResponse.Bill bills = 4;</code>
     */
    private $bills;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $total
     *     @type int $total_point
     *     @type int $total_lock_point
     *     @type \Omy\Crius\ExtensionServer\GetMemberPointBillsResponse\Bill[]|\Google\Protobuf\Internal\RepeatedField $bills
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\Member::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 total = 1;</code>
     * @return int
     */
    public function getTotal()
    {
        return $this->total;
    }

    /**
     * Generated from protobuf field <code>int32 total = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setTotal($var)
    {
        GPBUtil::checkInt32($var);
        $this->total = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 total_point = 2;</code>
     * @return int
     */
    public function getTotalPoint()
    {
        return $this->total_point;
    }

    /**
     * Generated from protobuf field <code>int32 total_point = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setTotalPoint($var)
    {
        GPBUtil::checkInt32($var);
        $this->total_point = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 total_lock_point = 3;</code>
     * @return int
     */
    public function getTotalLockPoint()
    {
        return $this->total_lock_point;
    }

    /**
     * Generated from protobuf field <code>int32 total_lock_point = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setTotalLockPoint($var)
    {
        GPBUtil::checkInt32($var);
        $this->total_lock_point = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.GetMemberPointBillsResponse.Bill bills = 4;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getBills()
    {
        return $this->bills;
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.GetMemberPointBillsResponse.Bill bills = 4;</code>
     * @param \Omy\Crius\ExtensionServer\GetMemberPointBillsResponse\Bill[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setBills($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\ExtensionServer\GetMemberPointBillsResponse\Bill::class);
        $this->bills = $arr;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Data::class, \Omy\Crius\ExtensionServer\GetMemberPointBillsResponse_Data::class);

