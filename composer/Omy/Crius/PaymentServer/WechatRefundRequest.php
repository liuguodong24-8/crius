<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/payment/wechat.proto

namespace Omy\Crius\PaymentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * WechatRefundRequest 退款
 *
 * Generated from protobuf message <code>payment.WechatRefundRequest</code>
 */
class WechatRefundRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string branch_id = 1;</code>
     */
    protected $branch_id = '';
    /**
     * Generated from protobuf field <code>string wechat_trade_id = 2;</code>
     */
    protected $wechat_trade_id = '';
    /**
     * Generated from protobuf field <code>map<string, string> string_map = 3;</code>
     */
    private $string_map;
    /**
     * Generated from protobuf field <code>map<string, int64> int64_map = 4;</code>
     */
    private $int64_map;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $branch_id
     *     @type string $wechat_trade_id
     *     @type array|\Google\Protobuf\Internal\MapField $string_map
     *     @type array|\Google\Protobuf\Internal\MapField $int64_map
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Payment\Wechat::initOnce();
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
     * Generated from protobuf field <code>string wechat_trade_id = 2;</code>
     * @return string
     */
    public function getWechatTradeId()
    {
        return $this->wechat_trade_id;
    }

    /**
     * Generated from protobuf field <code>string wechat_trade_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setWechatTradeId($var)
    {
        GPBUtil::checkString($var, True);
        $this->wechat_trade_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>map<string, string> string_map = 3;</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getStringMap()
    {
        return $this->string_map;
    }

    /**
     * Generated from protobuf field <code>map<string, string> string_map = 3;</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setStringMap($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::STRING);
        $this->string_map = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>map<string, int64> int64_map = 4;</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getInt64Map()
    {
        return $this->int64_map;
    }

    /**
     * Generated from protobuf field <code>map<string, int64> int64_map = 4;</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setInt64Map($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::INT64);
        $this->int64_map = $arr;

        return $this;
    }

}

