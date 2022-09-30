<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.GetAppointmentBreachLimitRequest</code>
 */
class GetAppointmentBreachLimitRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string phone = 1;</code>
     */
    protected $phone = '';
    /**
     * Generated from protobuf field <code>string phone_code = 2;</code>
     */
    protected $phone_code = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $phone
     *     @type string $phone_code
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\Appointment::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string phone = 1;</code>
     * @return string
     */
    public function getPhone()
    {
        return $this->phone;
    }

    /**
     * Generated from protobuf field <code>string phone = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setPhone($var)
    {
        GPBUtil::checkString($var, True);
        $this->phone = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string phone_code = 2;</code>
     * @return string
     */
    public function getPhoneCode()
    {
        return $this->phone_code;
    }

    /**
     * Generated from protobuf field <code>string phone_code = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setPhoneCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->phone_code = $var;

        return $this;
    }

}
