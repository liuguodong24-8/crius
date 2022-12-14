<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment_theme_feature.proto

namespace Omy\Crius\AppointmentServer\GetAppointmentThemeFeaturesResponse;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.GetAppointmentThemeFeaturesResponse.Data</code>
 */
class Data extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 total = 1;</code>
     */
    protected $total = 0;
    /**
     * Generated from protobuf field <code>repeated .appointment.AppointmentThemeFeature features = 2;</code>
     */
    private $features;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $total
     *     @type \Omy\Crius\AppointmentServer\AppointmentThemeFeature[]|\Google\Protobuf\Internal\RepeatedField $features
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\AppointmentThemeFeature::initOnce();
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
     * Generated from protobuf field <code>repeated .appointment.AppointmentThemeFeature features = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getFeatures()
    {
        return $this->features;
    }

    /**
     * Generated from protobuf field <code>repeated .appointment.AppointmentThemeFeature features = 2;</code>
     * @param \Omy\Crius\AppointmentServer\AppointmentThemeFeature[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setFeatures($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\AppointmentServer\AppointmentThemeFeature::class);
        $this->features = $arr;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Data::class, \Omy\Crius\AppointmentServer\GetAppointmentThemeFeaturesResponse_Data::class);

