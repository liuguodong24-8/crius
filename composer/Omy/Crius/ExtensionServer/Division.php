<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/division_level.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.Division</code>
 */
class Division extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>string name = 2;</code>
     */
    protected $name = '';
    /**
     * Generated from protobuf field <code>string background = 3;</code>
     */
    protected $background = '';
    /**
     * Generated from protobuf field <code>string color = 4;</code>
     */
    protected $color = '';
    /**
     * Generated from protobuf field <code>bool submit = 5;</code>
     */
    protected $submit = false;
    /**
     * Generated from protobuf field <code>int32 created_at = 6;</code>
     */
    protected $created_at = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $name
     *     @type string $background
     *     @type string $color
     *     @type bool $submit
     *     @type int $created_at
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\DivisionLevel::initOnce();
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
     * Generated from protobuf field <code>string name = 2;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
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
     * Generated from protobuf field <code>string background = 3;</code>
     * @return string
     */
    public function getBackground()
    {
        return $this->background;
    }

    /**
     * Generated from protobuf field <code>string background = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setBackground($var)
    {
        GPBUtil::checkString($var, True);
        $this->background = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string color = 4;</code>
     * @return string
     */
    public function getColor()
    {
        return $this->color;
    }

    /**
     * Generated from protobuf field <code>string color = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setColor($var)
    {
        GPBUtil::checkString($var, True);
        $this->color = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool submit = 5;</code>
     * @return bool
     */
    public function getSubmit()
    {
        return $this->submit;
    }

    /**
     * Generated from protobuf field <code>bool submit = 5;</code>
     * @param bool $var
     * @return $this
     */
    public function setSubmit($var)
    {
        GPBUtil::checkBool($var);
        $this->submit = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 created_at = 6;</code>
     * @return int
     */
    public function getCreatedAt()
    {
        return $this->created_at;
    }

    /**
     * Generated from protobuf field <code>int32 created_at = 6;</code>
     * @param int $var
     * @return $this
     */
    public function setCreatedAt($var)
    {
        GPBUtil::checkInt32($var);
        $this->created_at = $var;

        return $this;
    }

}

