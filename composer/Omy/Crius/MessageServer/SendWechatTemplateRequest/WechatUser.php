<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/message/wechat.proto

namespace Omy\Crius\MessageServer\SendWechatTemplateRequest;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>message.SendWechatTemplateRequest.WechatUser</code>
 */
class WechatUser extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string member_id = 1;</code>
     */
    protected $member_id = '';
    /**
     * Generated from protobuf field <code>string member_wechat_id = 2;</code>
     */
    protected $member_wechat_id = '';
    /**
     * Generated from protobuf field <code>string member_open_id = 3;</code>
     */
    protected $member_open_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $member_id
     *     @type string $member_wechat_id
     *     @type string $member_open_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Message\Wechat::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string member_id = 1;</code>
     * @return string
     */
    public function getMemberId()
    {
        return $this->member_id;
    }

    /**
     * Generated from protobuf field <code>string member_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setMemberId($var)
    {
        GPBUtil::checkString($var, True);
        $this->member_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string member_wechat_id = 2;</code>
     * @return string
     */
    public function getMemberWechatId()
    {
        return $this->member_wechat_id;
    }

    /**
     * Generated from protobuf field <code>string member_wechat_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setMemberWechatId($var)
    {
        GPBUtil::checkString($var, True);
        $this->member_wechat_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string member_open_id = 3;</code>
     * @return string
     */
    public function getMemberOpenId()
    {
        return $this->member_open_id;
    }

    /**
     * Generated from protobuf field <code>string member_open_id = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setMemberOpenId($var)
    {
        GPBUtil::checkString($var, True);
        $this->member_open_id = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(WechatUser::class, \Omy\Crius\MessageServer\SendWechatTemplateRequest_WechatUser::class);

