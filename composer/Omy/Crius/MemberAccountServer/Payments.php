<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-account/card.proto

namespace Omy\Crius\MemberAccountServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberAccount.Payments</code>
 */
class Payments extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 wechat = 1;</code>
     */
    protected $wechat = 0;
    /**
     * Generated from protobuf field <code>int32 cash = 2;</code>
     */
    protected $cash = 0;
    /**
     * Generated from protobuf field <code>int32 alipay = 3;</code>
     */
    protected $alipay = 0;
    /**
     * Generated from protobuf field <code>int32 card = 4;</code>
     */
    protected $card = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $wechat
     *     @type int $cash
     *     @type int $alipay
     *     @type int $card
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberAccount\Card::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 wechat = 1;</code>
     * @return int
     */
    public function getWechat()
    {
        return $this->wechat;
    }

    /**
     * Generated from protobuf field <code>int32 wechat = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setWechat($var)
    {
        GPBUtil::checkInt32($var);
        $this->wechat = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 cash = 2;</code>
     * @return int
     */
    public function getCash()
    {
        return $this->cash;
    }

    /**
     * Generated from protobuf field <code>int32 cash = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setCash($var)
    {
        GPBUtil::checkInt32($var);
        $this->cash = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 alipay = 3;</code>
     * @return int
     */
    public function getAlipay()
    {
        return $this->alipay;
    }

    /**
     * Generated from protobuf field <code>int32 alipay = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setAlipay($var)
    {
        GPBUtil::checkInt32($var);
        $this->alipay = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 card = 4;</code>
     * @return int
     */
    public function getCard()
    {
        return $this->card;
    }

    /**
     * Generated from protobuf field <code>int32 card = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setCard($var)
    {
        GPBUtil::checkInt32($var);
        $this->card = $var;

        return $this;
    }

}
