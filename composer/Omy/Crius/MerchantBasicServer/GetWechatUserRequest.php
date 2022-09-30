<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/wechat_member.proto

namespace Omy\Crius\MerchantBasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.GetWechatUserRequest</code>
 */
class GetWechatUserRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string openid = 1;</code>
     */
    protected $openid = '';
    /**
     * Generated from protobuf field <code>string appid = 2;</code>
     */
    protected $appid = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $openid
     *     @type string $appid
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\WechatMember::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string openid = 1;</code>
     * @return string
     */
    public function getOpenid()
    {
        return $this->openid;
    }

    /**
     * Generated from protobuf field <code>string openid = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setOpenid($var)
    {
        GPBUtil::checkString($var, True);
        $this->openid = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string appid = 2;</code>
     * @return string
     */
    public function getAppid()
    {
        return $this->appid;
    }

    /**
     * Generated from protobuf field <code>string appid = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setAppid($var)
    {
        GPBUtil::checkString($var, True);
        $this->appid = $var;

        return $this;
    }

}

