<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.AppointmentRemaining</code>
 */
class AppointmentRemaining extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string room_type_id = 1;</code>
     */
    protected $room_type_id = '';
    /**
     * Generated from protobuf field <code>repeated .appointment.RoomTypeNum room_type_num = 2;</code>
     */
    private $room_type_num;
    /**
     * Generated from protobuf field <code>int32 deposit_fee = 3;</code>
     */
    protected $deposit_fee = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $room_type_id
     *     @type \Omy\Crius\AppointmentServer\RoomTypeNum[]|\Google\Protobuf\Internal\RepeatedField $room_type_num
     *     @type int $deposit_fee
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\Appointment::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string room_type_id = 1;</code>
     * @return string
     */
    public function getRoomTypeId()
    {
        return $this->room_type_id;
    }

    /**
     * Generated from protobuf field <code>string room_type_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setRoomTypeId($var)
    {
        GPBUtil::checkString($var, True);
        $this->room_type_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .appointment.RoomTypeNum room_type_num = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getRoomTypeNum()
    {
        return $this->room_type_num;
    }

    /**
     * Generated from protobuf field <code>repeated .appointment.RoomTypeNum room_type_num = 2;</code>
     * @param \Omy\Crius\AppointmentServer\RoomTypeNum[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setRoomTypeNum($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\AppointmentServer\RoomTypeNum::class);
        $this->room_type_num = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 deposit_fee = 3;</code>
     * @return int
     */
    public function getDepositFee()
    {
        return $this->deposit_fee;
    }

    /**
     * Generated from protobuf field <code>int32 deposit_fee = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setDepositFee($var)
    {
        GPBUtil::checkInt32($var);
        $this->deposit_fee = $var;

        return $this;
    }

}

