<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/message/setting.proto

namespace Omy\Crius\MessageServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>message.ListMessageSettingData</code>
 */
class ListMessageSettingData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .message.MessageSetting settings = 1;</code>
     */
    private $settings;
    /**
     * Generated from protobuf field <code>int64 total = 2;</code>
     */
    protected $total = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\MessageServer\MessageSetting[]|\Google\Protobuf\Internal\RepeatedField $settings
     *     @type int|string $total
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Message\Setting::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .message.MessageSetting settings = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getSettings()
    {
        return $this->settings;
    }

    /**
     * Generated from protobuf field <code>repeated .message.MessageSetting settings = 1;</code>
     * @param \Omy\Crius\MessageServer\MessageSetting[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setSettings($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MessageServer\MessageSetting::class);
        $this->settings = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 total = 2;</code>
     * @return int|string
     */
    public function getTotal()
    {
        return $this->total;
    }

    /**
     * Generated from protobuf field <code>int64 total = 2;</code>
     * @param int|string $var
     * @return $this
     */
    public function setTotal($var)
    {
        GPBUtil::checkInt64($var);
        $this->total = $var;

        return $this;
    }

}
