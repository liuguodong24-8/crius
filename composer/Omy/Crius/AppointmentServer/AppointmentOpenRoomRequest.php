<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.AppointmentOpenRoomRequest</code>
 */
class AppointmentOpenRoomRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>string room_id = 2;</code>
     */
    protected $room_id = '';
    /**
     * Generated from protobuf field <code>int32 open_at = 3;</code>
     */
    protected $open_at = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $room_id
     *     @type int $open_at
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\Appointment::initOnce();
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
     * Generated from protobuf field <code>string room_id = 2;</code>
     * @return string
     */
    public function getRoomId()
    {
        return $this->room_id;
    }

    /**
     * Generated from protobuf field <code>string room_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setRoomId($var)
    {
        GPBUtil::checkString($var, True);
        $this->room_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 open_at = 3;</code>
     * @return int
     */
    public function getOpenAt()
    {
        return $this->open_at;
    }

    /**
     * Generated from protobuf field <code>int32 open_at = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setOpenAt($var)
    {
        GPBUtil::checkInt32($var);
        $this->open_at = $var;

        return $this;
    }

}

