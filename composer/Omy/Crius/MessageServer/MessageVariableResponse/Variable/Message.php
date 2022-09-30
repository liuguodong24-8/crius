<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/message/message.proto

namespace Omy\Crius\MessageServer\MessageVariableResponse\Variable;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>message.MessageVariableResponse.Variable.Message</code>
 */
class Message extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string category = 1;</code>
     */
    protected $category = '';
    /**
     * Generated from protobuf field <code>string category_key = 2;</code>
     */
    protected $category_key = '';
    /**
     * Generated from protobuf field <code>repeated string variables = 3;</code>
     */
    private $variables;
    /**
     * Generated from protobuf field <code>string trigger = 4;</code>
     */
    protected $trigger = '';
    /**
     * Generated from protobuf field <code>bool setting_disable = 5;</code>
     */
    protected $setting_disable = false;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $category
     *     @type string $category_key
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $variables
     *     @type string $trigger
     *     @type bool $setting_disable
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Message\Message::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string category = 1;</code>
     * @return string
     */
    public function getCategory()
    {
        return $this->category;
    }

    /**
     * Generated from protobuf field <code>string category = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setCategory($var)
    {
        GPBUtil::checkString($var, True);
        $this->category = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string category_key = 2;</code>
     * @return string
     */
    public function getCategoryKey()
    {
        return $this->category_key;
    }

    /**
     * Generated from protobuf field <code>string category_key = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setCategoryKey($var)
    {
        GPBUtil::checkString($var, True);
        $this->category_key = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string variables = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getVariables()
    {
        return $this->variables;
    }

    /**
     * Generated from protobuf field <code>repeated string variables = 3;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setVariables($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->variables = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string trigger = 4;</code>
     * @return string
     */
    public function getTrigger()
    {
        return $this->trigger;
    }

    /**
     * Generated from protobuf field <code>string trigger = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setTrigger($var)
    {
        GPBUtil::checkString($var, True);
        $this->trigger = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool setting_disable = 5;</code>
     * @return bool
     */
    public function getSettingDisable()
    {
        return $this->setting_disable;
    }

    /**
     * Generated from protobuf field <code>bool setting_disable = 5;</code>
     * @param bool $var
     * @return $this
     */
    public function setSettingDisable($var)
    {
        GPBUtil::checkBool($var);
        $this->setting_disable = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Message::class, \Omy\Crius\MessageServer\MessageVariableResponse_Variable_Message::class);
