<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.RoomTypeNum</code>
 */
class RoomTypeNum extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 way = 1;</code>
     */
    protected $way = 0;
    /**
     * Generated from protobuf field <code>int32 num = 2;</code>
     */
    protected $num = 0;
    /**
     * Generated from protobuf field <code>string time = 3;</code>
     */
    protected $time = '';
    /**
     * Generated from protobuf field <code>int32 total = 4;</code>
     */
    protected $total = 0;
    /**
     * Generated from protobuf field <code>bool is_next_day = 5;</code>
     */
    protected $is_next_day = false;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $way
     *     @type int $num
     *     @type string $time
     *     @type int $total
     *     @type bool $is_next_day
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\Appointment::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 way = 1;</code>
     * @return int
     */
    public function getWay()
    {
        return $this->way;
    }

    /**
     * Generated from protobuf field <code>int32 way = 1;</code>
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
     * Generated from protobuf field <code>string time = 3;</code>
     * @return string
     */
    public function getTime()
    {
        return $this->time;
    }

    /**
     * Generated from protobuf field <code>string time = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setTime($var)
    {
        GPBUtil::checkString($var, True);
        $this->time = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 total = 4;</code>
     * @return int
     */
    public function getTotal()
    {
        return $this->total;
    }

    /**
     * Generated from protobuf field <code>int32 total = 4;</code>
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
     * Generated from protobuf field <code>bool is_next_day = 5;</code>
     * @return bool
     */
    public function getIsNextDay()
    {
        return $this->is_next_day;
    }

    /**
     * Generated from protobuf field <code>bool is_next_day = 5;</code>
     * @param bool $var
     * @return $this
     */
    public function setIsNextDay($var)
    {
        GPBUtil::checkBool($var);
        $this->is_next_day = $var;

        return $this;
    }

}
