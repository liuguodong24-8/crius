<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/assist/assist.proto

namespace Omy\Crius\AssistServer\Order;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>assist.Order.JoinPeople</code>
 */
class JoinPeople extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string wechat_id = 1;</code>
     */
    protected $wechat_id = '';
    /**
     * Generated from protobuf field <code>int32 time = 2;</code>
     */
    protected $time = 0;
    /**
     * Generated from protobuf field <code>string source = 3;</code>
     */
    protected $source = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $wechat_id
     *     @type int $time
     *     @type string $source
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Assist\Assist::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string wechat_id = 1;</code>
     * @return string
     */
    public function getWechatId()
    {
        return $this->wechat_id;
    }

    /**
     * Generated from protobuf field <code>string wechat_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setWechatId($var)
    {
        GPBUtil::checkString($var, True);
        $this->wechat_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 time = 2;</code>
     * @return int
     */
    public function getTime()
    {
        return $this->time;
    }

    /**
     * Generated from protobuf field <code>int32 time = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setTime($var)
    {
        GPBUtil::checkInt32($var);
        $this->time = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string source = 3;</code>
     * @return string
     */
    public function getSource()
    {
        return $this->source;
    }

    /**
     * Generated from protobuf field <code>string source = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setSource($var)
    {
        GPBUtil::checkString($var, True);
        $this->source = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(JoinPeople::class, \Omy\Crius\AssistServer\Order_JoinPeople::class);
