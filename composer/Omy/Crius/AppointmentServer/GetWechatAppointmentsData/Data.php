<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment.proto

namespace Omy\Crius\AppointmentServer\GetWechatAppointmentsData;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.GetWechatAppointmentsData.Data</code>
 */
class Data extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.appointment.Appointment appointment = 1;</code>
     */
    protected $appointment = null;
    /**
     * Generated from protobuf field <code>.appointment.ThemeRoomType theme = 2;</code>
     */
    protected $theme = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\AppointmentServer\Appointment $appointment
     *     @type \Omy\Crius\AppointmentServer\ThemeRoomType $theme
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\Appointment::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.appointment.Appointment appointment = 1;</code>
     * @return \Omy\Crius\AppointmentServer\Appointment|null
     */
    public function getAppointment()
    {
        return $this->appointment;
    }

    public function hasAppointment()
    {
        return isset($this->appointment);
    }

    public function clearAppointment()
    {
        unset($this->appointment);
    }

    /**
     * Generated from protobuf field <code>.appointment.Appointment appointment = 1;</code>
     * @param \Omy\Crius\AppointmentServer\Appointment $var
     * @return $this
     */
    public function setAppointment($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\AppointmentServer\Appointment::class);
        $this->appointment = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.appointment.ThemeRoomType theme = 2;</code>
     * @return \Omy\Crius\AppointmentServer\ThemeRoomType|null
     */
    public function getTheme()
    {
        return $this->theme;
    }

    public function hasTheme()
    {
        return isset($this->theme);
    }

    public function clearTheme()
    {
        unset($this->theme);
    }

    /**
     * Generated from protobuf field <code>.appointment.ThemeRoomType theme = 2;</code>
     * @param \Omy\Crius\AppointmentServer\ThemeRoomType $var
     * @return $this
     */
    public function setTheme($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\AppointmentServer\ThemeRoomType::class);
        $this->theme = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Data::class, \Omy\Crius\AppointmentServer\GetWechatAppointmentsData_Data::class);

