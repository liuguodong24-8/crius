<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment_theme_feature.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.UpdateAppointmentThemeFeatureRequest</code>
 */
class UpdateAppointmentThemeFeatureRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.appointment.AppointmentThemeFeature feature = 1;</code>
     */
    protected $feature = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\AppointmentServer\AppointmentThemeFeature $feature
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\AppointmentThemeFeature::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.appointment.AppointmentThemeFeature feature = 1;</code>
     * @return \Omy\Crius\AppointmentServer\AppointmentThemeFeature|null
     */
    public function getFeature()
    {
        return $this->feature;
    }

    public function hasFeature()
    {
        return isset($this->feature);
    }

    public function clearFeature()
    {
        unset($this->feature);
    }

    /**
     * Generated from protobuf field <code>.appointment.AppointmentThemeFeature feature = 1;</code>
     * @param \Omy\Crius\AppointmentServer\AppointmentThemeFeature $var
     * @return $this
     */
    public function setFeature($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\AppointmentServer\AppointmentThemeFeature::class);
        $this->feature = $var;

        return $this;
    }

}

