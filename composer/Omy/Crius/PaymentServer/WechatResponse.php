<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/payment/wechat.proto

namespace Omy\Crius\PaymentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * WechatResponse 微信返回信息
 *
 * Generated from protobuf message <code>payment.WechatResponse</code>
 */
class WechatResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string return_code = 1;</code>
     */
    protected $return_code = '';
    /**
     * Generated from protobuf field <code>string return_msg = 2;</code>
     */
    protected $return_msg = '';
    /**
     * Generated from protobuf field <code>string detail = 3;</code>
     */
    protected $detail = '';
    /**
     * Generated from protobuf field <code>map<string, string> data = 4;</code>
     */
    private $data;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $return_code
     *     @type string $return_msg
     *     @type string $detail
     *     @type array|\Google\Protobuf\Internal\MapField $data
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Payment\Wechat::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string return_code = 1;</code>
     * @return string
     */
    public function getReturnCode()
    {
        return $this->return_code;
    }

    /**
     * Generated from protobuf field <code>string return_code = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setReturnCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->return_code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string return_msg = 2;</code>
     * @return string
     */
    public function getReturnMsg()
    {
        return $this->return_msg;
    }

    /**
     * Generated from protobuf field <code>string return_msg = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setReturnMsg($var)
    {
        GPBUtil::checkString($var, True);
        $this->return_msg = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string detail = 3;</code>
     * @return string
     */
    public function getDetail()
    {
        return $this->detail;
    }

    /**
     * Generated from protobuf field <code>string detail = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setDetail($var)
    {
        GPBUtil::checkString($var, True);
        $this->detail = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>map<string, string> data = 4;</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getData()
    {
        return $this->data;
    }

    /**
     * Generated from protobuf field <code>map<string, string> data = 4;</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setData($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::STRING);
        $this->data = $arr;

        return $this;
    }

}

