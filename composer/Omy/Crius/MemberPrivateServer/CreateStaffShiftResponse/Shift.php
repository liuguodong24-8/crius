<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-private/staff_shift.proto

namespace Omy\Crius\MemberPrivateServer\CreateStaffShiftResponse;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberPrivate.CreateStaffShiftResponse.Shift</code>
 */
class Shift extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int64 begin_time = 1;</code>
     */
    protected $begin_time = 0;
    /**
     * Generated from protobuf field <code>int64 end_time = 2;</code>
     */
    protected $end_time = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int|string $begin_time
     *     @type int|string $end_time
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberPrivate\StaffShift::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int64 begin_time = 1;</code>
     * @return int|string
     */
    public function getBeginTime()
    {
        return $this->begin_time;
    }

    /**
     * Generated from protobuf field <code>int64 begin_time = 1;</code>
     * @param int|string $var
     * @return $this
     */
    public function setBeginTime($var)
    {
        GPBUtil::checkInt64($var);
        $this->begin_time = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 end_time = 2;</code>
     * @return int|string
     */
    public function getEndTime()
    {
        return $this->end_time;
    }

    /**
     * Generated from protobuf field <code>int64 end_time = 2;</code>
     * @param int|string $var
     * @return $this
     */
    public function setEndTime($var)
    {
        GPBUtil::checkInt64($var);
        $this->end_time = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Shift::class, \Omy\Crius\MemberPrivateServer\CreateStaffShiftResponse_Shift::class);

