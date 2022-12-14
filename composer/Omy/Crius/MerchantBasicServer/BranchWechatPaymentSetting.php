<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/payment.proto

namespace Omy\Crius\MerchantBasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.BranchWechatPaymentSetting</code>
 */
class BranchWechatPaymentSetting extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string branch_id = 1;</code>
     */
    protected $branch_id = '';
    /**
     * Generated from protobuf field <code>string merchant_id = 2;</code>
     */
    protected $merchant_id = '';
    /**
     * Generated from protobuf field <code>string app_id = 3;</code>
     */
    protected $app_id = '';
    /**
     * Generated from protobuf field <code>string mch_id = 4;</code>
     */
    protected $mch_id = '';
    /**
     * Generated from protobuf field <code>string sub_mch_id = 5;</code>
     */
    protected $sub_mch_id = '';
    /**
     * Generated from protobuf field <code>string private_key = 6;</code>
     */
    protected $private_key = '';
    /**
     * Generated from protobuf field <code>string cert_filename = 7;</code>
     */
    protected $cert_filename = '';
    /**
     * Generated from protobuf field <code>bytes cert_content = 8;</code>
     */
    protected $cert_content = '';
    /**
     * Generated from protobuf field <code>string headquarters_sub_mch_id = 9;</code>
     */
    protected $headquarters_sub_mch_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $branch_id
     *     @type string $merchant_id
     *     @type string $app_id
     *     @type string $mch_id
     *     @type string $sub_mch_id
     *     @type string $private_key
     *     @type string $cert_filename
     *     @type string $cert_content
     *     @type string $headquarters_sub_mch_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\Payment::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string branch_id = 1;</code>
     * @return string
     */
    public function getBranchId()
    {
        return $this->branch_id;
    }

    /**
     * Generated from protobuf field <code>string branch_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setBranchId($var)
    {
        GPBUtil::checkString($var, True);
        $this->branch_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string merchant_id = 2;</code>
     * @return string
     */
    public function getMerchantId()
    {
        return $this->merchant_id;
    }

    /**
     * Generated from protobuf field <code>string merchant_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setMerchantId($var)
    {
        GPBUtil::checkString($var, True);
        $this->merchant_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string app_id = 3;</code>
     * @return string
     */
    public function getAppId()
    {
        return $this->app_id;
    }

    /**
     * Generated from protobuf field <code>string app_id = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setAppId($var)
    {
        GPBUtil::checkString($var, True);
        $this->app_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string mch_id = 4;</code>
     * @return string
     */
    public function getMchId()
    {
        return $this->mch_id;
    }

    /**
     * Generated from protobuf field <code>string mch_id = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setMchId($var)
    {
        GPBUtil::checkString($var, True);
        $this->mch_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string sub_mch_id = 5;</code>
     * @return string
     */
    public function getSubMchId()
    {
        return $this->sub_mch_id;
    }

    /**
     * Generated from protobuf field <code>string sub_mch_id = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setSubMchId($var)
    {
        GPBUtil::checkString($var, True);
        $this->sub_mch_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string private_key = 6;</code>
     * @return string
     */
    public function getPrivateKey()
    {
        return $this->private_key;
    }

    /**
     * Generated from protobuf field <code>string private_key = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setPrivateKey($var)
    {
        GPBUtil::checkString($var, True);
        $this->private_key = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string cert_filename = 7;</code>
     * @return string
     */
    public function getCertFilename()
    {
        return $this->cert_filename;
    }

    /**
     * Generated from protobuf field <code>string cert_filename = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setCertFilename($var)
    {
        GPBUtil::checkString($var, True);
        $this->cert_filename = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bytes cert_content = 8;</code>
     * @return string
     */
    public function getCertContent()
    {
        return $this->cert_content;
    }

    /**
     * Generated from protobuf field <code>bytes cert_content = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setCertContent($var)
    {
        GPBUtil::checkString($var, False);
        $this->cert_content = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string headquarters_sub_mch_id = 9;</code>
     * @return string
     */
    public function getHeadquartersSubMchId()
    {
        return $this->headquarters_sub_mch_id;
    }

    /**
     * Generated from protobuf field <code>string headquarters_sub_mch_id = 9;</code>
     * @param string $var
     * @return $this
     */
    public function setHeadquartersSubMchId($var)
    {
        GPBUtil::checkString($var, True);
        $this->headquarters_sub_mch_id = $var;

        return $this;
    }

}

