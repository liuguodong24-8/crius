<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment_template.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.SaveAppointmentTemplateRequest</code>
 */
class SaveAppointmentTemplateRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.appointment.AppointmentTemplate template = 1;</code>
     */
    protected $template = null;
    /**
     * Generated from protobuf field <code>repeated .appointment.AppointmentTemplateConfig config = 2;</code>
     */
    private $config;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\AppointmentServer\AppointmentTemplate $template
     *     @type \Omy\Crius\AppointmentServer\AppointmentTemplateConfig[]|\Google\Protobuf\Internal\RepeatedField $config
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\AppointmentTemplate::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.appointment.AppointmentTemplate template = 1;</code>
     * @return \Omy\Crius\AppointmentServer\AppointmentTemplate|null
     */
    public function getTemplate()
    {
        return $this->template;
    }

    public function hasTemplate()
    {
        return isset($this->template);
    }

    public function clearTemplate()
    {
        unset($this->template);
    }

    /**
     * Generated from protobuf field <code>.appointment.AppointmentTemplate template = 1;</code>
     * @param \Omy\Crius\AppointmentServer\AppointmentTemplate $var
     * @return $this
     */
    public function setTemplate($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\AppointmentServer\AppointmentTemplate::class);
        $this->template = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .appointment.AppointmentTemplateConfig config = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getConfig()
    {
        return $this->config;
    }

    /**
     * Generated from protobuf field <code>repeated .appointment.AppointmentTemplateConfig config = 2;</code>
     * @param \Omy\Crius\AppointmentServer\AppointmentTemplateConfig[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setConfig($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\AppointmentServer\AppointmentTemplateConfig::class);
        $this->config = $arr;

        return $this;
    }

}

