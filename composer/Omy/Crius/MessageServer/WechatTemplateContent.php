<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/message/wechat.proto

namespace Omy\Crius\MessageServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>message.WechatTemplateContent</code>
 */
class WechatTemplateContent extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.message.WechatTemplateContentBase first = 1;</code>
     */
    protected $first = null;
    /**
     * Generated from protobuf field <code>repeated .message.WechatTemplateContentDetail detail = 2;</code>
     */
    private $detail;
    /**
     * Generated from protobuf field <code>.message.WechatTemplateContentBase remark = 3;</code>
     */
    protected $remark = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\MessageServer\WechatTemplateContentBase $first
     *     @type \Omy\Crius\MessageServer\WechatTemplateContentDetail[]|\Google\Protobuf\Internal\RepeatedField $detail
     *     @type \Omy\Crius\MessageServer\WechatTemplateContentBase $remark
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Message\Wechat::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.message.WechatTemplateContentBase first = 1;</code>
     * @return \Omy\Crius\MessageServer\WechatTemplateContentBase|null
     */
    public function getFirst()
    {
        return $this->first;
    }

    public function hasFirst()
    {
        return isset($this->first);
    }

    public function clearFirst()
    {
        unset($this->first);
    }

    /**
     * Generated from protobuf field <code>.message.WechatTemplateContentBase first = 1;</code>
     * @param \Omy\Crius\MessageServer\WechatTemplateContentBase $var
     * @return $this
     */
    public function setFirst($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\MessageServer\WechatTemplateContentBase::class);
        $this->first = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .message.WechatTemplateContentDetail detail = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getDetail()
    {
        return $this->detail;
    }

    /**
     * Generated from protobuf field <code>repeated .message.WechatTemplateContentDetail detail = 2;</code>
     * @param \Omy\Crius\MessageServer\WechatTemplateContentDetail[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setDetail($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MessageServer\WechatTemplateContentDetail::class);
        $this->detail = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.message.WechatTemplateContentBase remark = 3;</code>
     * @return \Omy\Crius\MessageServer\WechatTemplateContentBase|null
     */
    public function getRemark()
    {
        return $this->remark;
    }

    public function hasRemark()
    {
        return isset($this->remark);
    }

    public function clearRemark()
    {
        unset($this->remark);
    }

    /**
     * Generated from protobuf field <code>.message.WechatTemplateContentBase remark = 3;</code>
     * @param \Omy\Crius\MessageServer\WechatTemplateContentBase $var
     * @return $this
     */
    public function setRemark($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\MessageServer\WechatTemplateContentBase::class);
        $this->remark = $var;

        return $this;
    }

}

