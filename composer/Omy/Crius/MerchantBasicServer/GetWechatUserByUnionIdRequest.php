<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/wechat_member.proto

namespace Omy\Crius\MerchantBasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.GetWechatUserByUnionIdRequest</code>
 */
class GetWechatUserByUnionIdRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string unionid = 1;</code>
     */
    protected $unionid = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $unionid
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\WechatMember::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string unionid = 1;</code>
     * @return string
     */
    public function getUnionid()
    {
        return $this->unionid;
    }

    /**
     * Generated from protobuf field <code>string unionid = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setUnionid($var)
    {
        GPBUtil::checkString($var, True);
        $this->unionid = $var;

        return $this;
    }

}

