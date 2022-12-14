<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-account/member_account.proto

namespace Omy\Crius\MemberAccountServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberAccount.CardWithAccount</code>
 */
class CardWithAccount extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.memberAccount.Card card = 1;</code>
     */
    protected $card = null;
    /**
     * Generated from protobuf field <code>repeated .memberAccount.Account accounts = 2;</code>
     */
    private $accounts;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\MemberAccountServer\Card $card
     *     @type \Omy\Crius\MemberAccountServer\Account[]|\Google\Protobuf\Internal\RepeatedField $accounts
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberAccount\MemberAccount::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.memberAccount.Card card = 1;</code>
     * @return \Omy\Crius\MemberAccountServer\Card|null
     */
    public function getCard()
    {
        return $this->card;
    }

    public function hasCard()
    {
        return isset($this->card);
    }

    public function clearCard()
    {
        unset($this->card);
    }

    /**
     * Generated from protobuf field <code>.memberAccount.Card card = 1;</code>
     * @param \Omy\Crius\MemberAccountServer\Card $var
     * @return $this
     */
    public function setCard($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\MemberAccountServer\Card::class);
        $this->card = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .memberAccount.Account accounts = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAccounts()
    {
        return $this->accounts;
    }

    /**
     * Generated from protobuf field <code>repeated .memberAccount.Account accounts = 2;</code>
     * @param \Omy\Crius\MemberAccountServer\Account[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAccounts($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MemberAccountServer\Account::class);
        $this->accounts = $arr;

        return $this;
    }

}

