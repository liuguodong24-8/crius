<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-account/card.proto

namespace Omy\Crius\MemberAccountServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberAccount.CancelCardRequest</code>
 */
class CancelCardRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string card_id = 1;</code>
     */
    protected $card_id = '';
    /**
     * Generated from protobuf field <code>string bank_account = 2;</code>
     */
    protected $bank_account = '';
    /**
     * Generated from protobuf field <code>string bank_name = 3;</code>
     */
    protected $bank_name = '';
    /**
     * Generated from protobuf field <code>string money_receiver = 4;</code>
     */
    protected $money_receiver = '';
    /**
     * Generated from protobuf field <code>string reason = 5;</code>
     */
    protected $reason = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $card_id
     *     @type string $bank_account
     *     @type string $bank_name
     *     @type string $money_receiver
     *     @type string $reason
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberAccount\Card::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string card_id = 1;</code>
     * @return string
     */
    public function getCardId()
    {
        return $this->card_id;
    }

    /**
     * Generated from protobuf field <code>string card_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setCardId($var)
    {
        GPBUtil::checkString($var, True);
        $this->card_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string bank_account = 2;</code>
     * @return string
     */
    public function getBankAccount()
    {
        return $this->bank_account;
    }

    /**
     * Generated from protobuf field <code>string bank_account = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setBankAccount($var)
    {
        GPBUtil::checkString($var, True);
        $this->bank_account = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string bank_name = 3;</code>
     * @return string
     */
    public function getBankName()
    {
        return $this->bank_name;
    }

    /**
     * Generated from protobuf field <code>string bank_name = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setBankName($var)
    {
        GPBUtil::checkString($var, True);
        $this->bank_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string money_receiver = 4;</code>
     * @return string
     */
    public function getMoneyReceiver()
    {
        return $this->money_receiver;
    }

    /**
     * Generated from protobuf field <code>string money_receiver = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setMoneyReceiver($var)
    {
        GPBUtil::checkString($var, True);
        $this->money_receiver = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string reason = 5;</code>
     * @return string
     */
    public function getReason()
    {
        return $this->reason;
    }

    /**
     * Generated from protobuf field <code>string reason = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setReason($var)
    {
        GPBUtil::checkString($var, True);
        $this->reason = $var;

        return $this;
    }

}

