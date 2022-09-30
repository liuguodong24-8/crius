<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/branch.proto

namespace Omy\Crius\MerchantBasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.UpdateBranchAccountRequest</code>
 */
class UpdateBranchAccountRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>string wechat_app_id = 2;</code>
     */
    protected $wechat_app_id = '';
    /**
     * Generated from protobuf field <code>string wechat_mch_id = 3;</code>
     */
    protected $wechat_mch_id = '';
    /**
     * Generated from protobuf field <code>string wechat_key = 4;</code>
     */
    protected $wechat_key = '';
    /**
     * Generated from protobuf field <code>string wechat_cert_path = 5;</code>
     */
    protected $wechat_cert_path = '';
    /**
     * Generated from protobuf field <code>string wechat_key_path = 6;</code>
     */
    protected $wechat_key_path = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $wechat_app_id
     *     @type string $wechat_mch_id
     *     @type string $wechat_key
     *     @type string $wechat_cert_path
     *     @type string $wechat_key_path
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\Branch::initOnce();
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
     * Generated from protobuf field <code>string wechat_app_id = 2;</code>
     * @return string
     */
    public function getWechatAppId()
    {
        return $this->wechat_app_id;
    }

    /**
     * Generated from protobuf field <code>string wechat_app_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setWechatAppId($var)
    {
        GPBUtil::checkString($var, True);
        $this->wechat_app_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string wechat_mch_id = 3;</code>
     * @return string
     */
    public function getWechatMchId()
    {
        return $this->wechat_mch_id;
    }

    /**
     * Generated from protobuf field <code>string wechat_mch_id = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setWechatMchId($var)
    {
        GPBUtil::checkString($var, True);
        $this->wechat_mch_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string wechat_key = 4;</code>
     * @return string
     */
    public function getWechatKey()
    {
        return $this->wechat_key;
    }

    /**
     * Generated from protobuf field <code>string wechat_key = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setWechatKey($var)
    {
        GPBUtil::checkString($var, True);
        $this->wechat_key = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string wechat_cert_path = 5;</code>
     * @return string
     */
    public function getWechatCertPath()
    {
        return $this->wechat_cert_path;
    }

    /**
     * Generated from protobuf field <code>string wechat_cert_path = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setWechatCertPath($var)
    {
        GPBUtil::checkString($var, True);
        $this->wechat_cert_path = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string wechat_key_path = 6;</code>
     * @return string
     */
    public function getWechatKeyPath()
    {
        return $this->wechat_key_path;
    }

    /**
     * Generated from protobuf field <code>string wechat_key_path = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setWechatKeyPath($var)
    {
        GPBUtil::checkString($var, True);
        $this->wechat_key_path = $var;

        return $this;
    }

}
