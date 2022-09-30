<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/black_list_config.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.BlackListConfig</code>
 */
class BlackListConfig extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * 显示名称
     *
     * Generated from protobuf field <code>string name = 2;</code>
     */
    protected $name = '';
    /**
     * 加入周期
     *
     * Generated from protobuf field <code>int32 in_days = 3;</code>
     */
    protected $in_days = 0;
    /**
     * 加入连续次数
     *
     * Generated from protobuf field <code>int32 in_times = 4;</code>
     */
    protected $in_times = 0;
    /**
     * 退出周期
     *
     * Generated from protobuf field <code>int32 out_days = 5;</code>
     */
    protected $out_days = 0;
    /**
     * 退出连续次数
     *
     * Generated from protobuf field <code>int32 out_times = 6;</code>
     */
    protected $out_times = 0;
    /**
     * 拦截来电
     *
     * Generated from protobuf field <code>bool block_call = 7;</code>
     */
    protected $block_call = false;
    /**
     * 黑名单等级
     *
     * Generated from protobuf field <code>int32 level = 8;</code>
     */
    protected $level = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $name
     *           显示名称
     *     @type int $in_days
     *           加入周期
     *     @type int $in_times
     *           加入连续次数
     *     @type int $out_days
     *           退出周期
     *     @type int $out_times
     *           退出连续次数
     *     @type bool $block_call
     *           拦截来电
     *     @type int $level
     *           黑名单等级
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\BlackListConfig::initOnce();
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
     * 显示名称
     *
     * Generated from protobuf field <code>string name = 2;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * 显示名称
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
     * 加入周期
     *
     * Generated from protobuf field <code>int32 in_days = 3;</code>
     * @return int
     */
    public function getInDays()
    {
        return $this->in_days;
    }

    /**
     * 加入周期
     *
     * Generated from protobuf field <code>int32 in_days = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setInDays($var)
    {
        GPBUtil::checkInt32($var);
        $this->in_days = $var;

        return $this;
    }

    /**
     * 加入连续次数
     *
     * Generated from protobuf field <code>int32 in_times = 4;</code>
     * @return int
     */
    public function getInTimes()
    {
        return $this->in_times;
    }

    /**
     * 加入连续次数
     *
     * Generated from protobuf field <code>int32 in_times = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setInTimes($var)
    {
        GPBUtil::checkInt32($var);
        $this->in_times = $var;

        return $this;
    }

    /**
     * 退出周期
     *
     * Generated from protobuf field <code>int32 out_days = 5;</code>
     * @return int
     */
    public function getOutDays()
    {
        return $this->out_days;
    }

    /**
     * 退出周期
     *
     * Generated from protobuf field <code>int32 out_days = 5;</code>
     * @param int $var
     * @return $this
     */
    public function setOutDays($var)
    {
        GPBUtil::checkInt32($var);
        $this->out_days = $var;

        return $this;
    }

    /**
     * 退出连续次数
     *
     * Generated from protobuf field <code>int32 out_times = 6;</code>
     * @return int
     */
    public function getOutTimes()
    {
        return $this->out_times;
    }

    /**
     * 退出连续次数
     *
     * Generated from protobuf field <code>int32 out_times = 6;</code>
     * @param int $var
     * @return $this
     */
    public function setOutTimes($var)
    {
        GPBUtil::checkInt32($var);
        $this->out_times = $var;

        return $this;
    }

    /**
     * 拦截来电
     *
     * Generated from protobuf field <code>bool block_call = 7;</code>
     * @return bool
     */
    public function getBlockCall()
    {
        return $this->block_call;
    }

    /**
     * 拦截来电
     *
     * Generated from protobuf field <code>bool block_call = 7;</code>
     * @param bool $var
     * @return $this
     */
    public function setBlockCall($var)
    {
        GPBUtil::checkBool($var);
        $this->block_call = $var;

        return $this;
    }

    /**
     * 黑名单等级
     *
     * Generated from protobuf field <code>int32 level = 8;</code>
     * @return int
     */
    public function getLevel()
    {
        return $this->level;
    }

    /**
     * 黑名单等级
     *
     * Generated from protobuf field <code>int32 level = 8;</code>
     * @param int $var
     * @return $this
     */
    public function setLevel($var)
    {
        GPBUtil::checkInt32($var);
        $this->level = $var;

        return $this;
    }

}
