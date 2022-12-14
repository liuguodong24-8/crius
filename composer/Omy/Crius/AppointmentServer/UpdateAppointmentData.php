<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.UpdateAppointmentData</code>
 */
class UpdateAppointmentData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string appointment_id = 1;</code>
     */
    protected $appointment_id = '';
    /**
     * Generated from protobuf field <code>string new_appointment_id = 2;</code>
     */
    protected $new_appointment_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $appointment_id
     *     @type string $new_appointment_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\Appointment::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string appointment_id = 1;</code>
     * @return string
     */
    public function getAppointmentId()
    {
        return $this->appointment_id;
    }

    /**
     * Generated from protobuf field <code>string appointment_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setAppointmentId($var)
    {
        GPBUtil::checkString($var, True);
        $this->appointment_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string new_appointment_id = 2;</code>
     * @return string
     */
    public function getNewAppointmentId()
    {
        return $this->new_appointment_id;
    }

    /**
     * Generated from protobuf field <code>string new_appointment_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setNewAppointmentId($var)
    {
        GPBUtil::checkString($var, True);
        $this->new_appointment_id = $var;

        return $this;
    }

}

