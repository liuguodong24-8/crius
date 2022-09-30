<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/task.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.Phone</code>
 */
class Phone extends \Google\Protobuf\Internal\Message
{
    /**
     * 手机
     *
     * Generated from protobuf field <code>string phone = 1;</code>
     */
    protected $phone = '';
    /**
     * 手机区号
     *
     * Generated from protobuf field <code>string phone_code = 2;</code>
     */
    protected $phone_code = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $phone
     *           手机
     *     @type string $phone_code
     *           手机区号
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\Task::initOnce();
        parent::__construct($data);
    }

    /**
     * 手机
     *
     * Generated from protobuf field <code>string phone = 1;</code>
     * @return string
     */
    public function getPhone()
    {
        return $this->phone;
    }

    /**
     * 手机
     *
     * Generated from protobuf field <code>string phone = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setPhone($var)
    {
        GPBUtil::checkString($var, True);
        $this->phone = $var;

        return $this;
    }

    /**
     * 手机区号
     *
     * Generated from protobuf field <code>string phone_code = 2;</code>
     * @return string
     */
    public function getPhoneCode()
    {
        return $this->phone_code;
    }

    /**
     * 手机区号
     *
     * Generated from protobuf field <code>string phone_code = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setPhoneCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->phone_code = $var;

        return $this;
    }

}

