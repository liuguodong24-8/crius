<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.SaveAppointmentTemplateConfigRequest</code>
 */
class SaveAppointmentTemplateConfigRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string template_id = 1;</code>
     */
    protected $template_id = '';
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
     *     @type string $template_id
     *     @type \Omy\Crius\AppointmentServer\AppointmentTemplateConfig[]|\Google\Protobuf\Internal\RepeatedField $config
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\Appointment::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string template_id = 1;</code>
     * @return string
     */
    public function getTemplateId()
    {
        return $this->template_id;
    }

    /**
     * Generated from protobuf field <code>string template_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setTemplateId($var)
    {
        GPBUtil::checkString($var, True);
        $this->template_id = $var;

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

