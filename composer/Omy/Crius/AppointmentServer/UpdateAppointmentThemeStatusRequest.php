<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment_theme.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.UpdateAppointmentThemeStatusRequest</code>
 */
class UpdateAppointmentThemeStatusRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * 主题状态
     *
     * Generated from protobuf field <code>string status = 2;</code>
     */
    protected $status = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $status
     *           主题状态
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\AppointmentTheme::initOnce();
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
     * 主题状态
     *
     * Generated from protobuf field <code>string status = 2;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * 主题状态
     *
     * Generated from protobuf field <code>string status = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkString($var, True);
        $this->status = $var;

        return $this;
    }

}

