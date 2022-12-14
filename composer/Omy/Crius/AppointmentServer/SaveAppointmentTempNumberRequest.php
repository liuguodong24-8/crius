<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.SaveAppointmentTempNumberRequest</code>
 */
class SaveAppointmentTempNumberRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string branch_id = 1;</code>
     */
    protected $branch_id = '';
    /**
     * Generated from protobuf field <code>int32 num = 2;</code>
     */
    protected $num = 0;
    /**
     * Generated from protobuf field <code>int32 appointment_at = 3;</code>
     */
    protected $appointment_at = 0;
    /**
     * Generated from protobuf field <code>int32 way = 4;</code>
     */
    protected $way = 0;
    /**
     * Generated from protobuf field <code>string room_type_id = 5;</code>
     */
    protected $room_type_id = '';
    /**
     * Generated from protobuf field <code>int32 appointment_date = 6;</code>
     */
    protected $appointment_date = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $branch_id
     *     @type int $num
     *     @type int $appointment_at
     *     @type int $way
     *     @type string $room_type_id
     *     @type int $appointment_date
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\Appointment::initOnce();
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
     * Generated from protobuf field <code>int32 num = 2;</code>
     * @return int
     */
    public function getNum()
    {
        return $this->num;
    }

    /**
     * Generated from protobuf field <code>int32 num = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setNum($var)
    {
        GPBUtil::checkInt32($var);
        $this->num = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 appointment_at = 3;</code>
     * @return int
     */
    public function getAppointmentAt()
    {
        return $this->appointment_at;
    }

    /**
     * Generated from protobuf field <code>int32 appointment_at = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setAppointmentAt($var)
    {
        GPBUtil::checkInt32($var);
        $this->appointment_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 way = 4;</code>
     * @return int
     */
    public function getWay()
    {
        return $this->way;
    }

    /**
     * Generated from protobuf field <code>int32 way = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setWay($var)
    {
        GPBUtil::checkInt32($var);
        $this->way = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string room_type_id = 5;</code>
     * @return string
     */
    public function getRoomTypeId()
    {
        return $this->room_type_id;
    }

    /**
     * Generated from protobuf field <code>string room_type_id = 5;</code>
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
     * Generated from protobuf field <code>int32 appointment_date = 6;</code>
     * @return int
     */
    public function getAppointmentDate()
    {
        return $this->appointment_date;
    }

    /**
     * Generated from protobuf field <code>int32 appointment_date = 6;</code>
     * @param int $var
     * @return $this
     */
    public function setAppointmentDate($var)
    {
        GPBUtil::checkInt32($var);
        $this->appointment_date = $var;

        return $this;
    }

}

