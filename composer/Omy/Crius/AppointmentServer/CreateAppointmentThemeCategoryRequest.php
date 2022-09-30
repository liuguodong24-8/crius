<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment_theme_category.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.CreateAppointmentThemeCategoryRequest</code>
 */
class CreateAppointmentThemeCategoryRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.appointment.AppointmentThemeCategory category = 1;</code>
     */
    protected $category = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\AppointmentServer\AppointmentThemeCategory $category
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\AppointmentThemeCategory::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.appointment.AppointmentThemeCategory category = 1;</code>
     * @return \Omy\Crius\AppointmentServer\AppointmentThemeCategory|null
     */
    public function getCategory()
    {
        return $this->category;
    }

    public function hasCategory()
    {
        return isset($this->category);
    }

    public function clearCategory()
    {
        unset($this->category);
    }

    /**
     * Generated from protobuf field <code>.appointment.AppointmentThemeCategory category = 1;</code>
     * @param \Omy\Crius\AppointmentServer\AppointmentThemeCategory $var
     * @return $this
     */
    public function setCategory($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\AppointmentServer\AppointmentThemeCategory::class);
        $this->category = $var;

        return $this;
    }

}

