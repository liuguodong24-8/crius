<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/message/message.proto

namespace Omy\Crius\MessageServer\MessageVariableResponse;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>message.MessageVariableResponse.Variable</code>
 */
class Variable extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string system = 1;</code>
     */
    protected $system = '';
    /**
     * Generated from protobuf field <code>string system_key = 2;</code>
     */
    protected $system_key = '';
    /**
     * Generated from protobuf field <code>repeated .message.MessageVariableResponse.Variable.Message message = 3;</code>
     */
    private $message;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $system
     *     @type string $system_key
     *     @type \Omy\Crius\MessageServer\MessageVariableResponse\Variable\Message[]|\Google\Protobuf\Internal\RepeatedField $message
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Message\Message::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string system = 1;</code>
     * @return string
     */
    public function getSystem()
    {
        return $this->system;
    }

    /**
     * Generated from protobuf field <code>string system = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setSystem($var)
    {
        GPBUtil::checkString($var, True);
        $this->system = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string system_key = 2;</code>
     * @return string
     */
    public function getSystemKey()
    {
        return $this->system_key;
    }

    /**
     * Generated from protobuf field <code>string system_key = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setSystemKey($var)
    {
        GPBUtil::checkString($var, True);
        $this->system_key = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .message.MessageVariableResponse.Variable.Message message = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getMessage()
    {
        return $this->message;
    }

    /**
     * Generated from protobuf field <code>repeated .message.MessageVariableResponse.Variable.Message message = 3;</code>
     * @param \Omy\Crius\MessageServer\MessageVariableResponse\Variable\Message[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setMessage($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MessageServer\MessageVariableResponse\Variable\Message::class);
        $this->message = $arr;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Variable::class, \Omy\Crius\MessageServer\MessageVariableResponse_Variable::class);

