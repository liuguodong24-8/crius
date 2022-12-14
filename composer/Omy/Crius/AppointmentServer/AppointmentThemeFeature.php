<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment_theme_feature.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.AppointmentThemeFeature</code>
 */
class AppointmentThemeFeature extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * 特色名字
     *
     * Generated from protobuf field <code>string name = 2;</code>
     */
    protected $name = '';
    /**
     * 特色权值
     *
     * Generated from protobuf field <code>int32 weight = 3;</code>
     */
    protected $weight = 0;
    /**
     * 特色状态
     *
     * Generated from protobuf field <code>string status = 4;</code>
     */
    protected $status = '';
    /**
     * 特色图标
     *
     * Generated from protobuf field <code>string icon = 5;</code>
     */
    protected $icon = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $name
     *           特色名字
     *     @type int $weight
     *           特色权值
     *     @type string $status
     *           特色状态
     *     @type string $icon
     *           特色图标
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\AppointmentThemeFeature::initOnce();
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
     * 特色名字
     *
     * Generated from protobuf field <code>string name = 2;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * 特色名字
     *
     * Generated from protobuf field <code>string name = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     * 特色权值
     *
     * Generated from protobuf field <code>int32 weight = 3;</code>
     * @return int
     */
    public function getWeight()
    {
        return $this->weight;
    }

    /**
     * 特色权值
     *
     * Generated from protobuf field <code>int32 weight = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setWeight($var)
    {
        GPBUtil::checkInt32($var);
        $this->weight = $var;

        return $this;
    }

    /**
     * 特色状态
     *
     * Generated from protobuf field <code>string status = 4;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * 特色状态
     *
     * Generated from protobuf field <code>string status = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkString($var, True);
        $this->status = $var;

        return $this;
    }

    /**
     * 特色图标
     *
     * Generated from protobuf field <code>string icon = 5;</code>
     * @return string
     */
    public function getIcon()
    {
        return $this->icon;
    }

    /**
     * 特色图标
     *
     * Generated from protobuf field <code>string icon = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setIcon($var)
    {
        GPBUtil::checkString($var, True);
        $this->icon = $var;

        return $this;
    }

}

