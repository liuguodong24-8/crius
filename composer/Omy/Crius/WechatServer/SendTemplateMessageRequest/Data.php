<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/wechat/message.proto

namespace Omy\Crius\WechatServer\SendTemplateMessageRequest;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>wechat.SendTemplateMessageRequest.Data</code>
 */
class Data extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string name = 1;</code>
     */
    protected $name = '';
    /**
     * Generated from protobuf field <code>string value = 2;</code>
     */
    protected $value = '';
    /**
     * Generated from protobuf field <code>string color = 3;</code>
     */
    protected $color = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *     @type string $value
     *     @type string $color
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Wechat\Message::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string name = 1;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * Generated from protobuf field <code>string name = 1;</code>
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
     * Generated from protobuf field <code>string value = 2;</code>
     * @return string
     */
    public function getValue()
    {
        return $this->value;
    }

    /**
     * Generated from protobuf field <code>string value = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setValue($var)
    {
        GPBUtil::checkString($var, True);
        $this->value = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string color = 3;</code>
     * @return string
     */
    public function getColor()
    {
        return $this->color;
    }

    /**
     * Generated from protobuf field <code>string color = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setColor($var)
    {
        GPBUtil::checkString($var, True);
        $this->color = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Data::class, \Omy\Crius\WechatServer\SendTemplateMessageRequest_Data::class);

