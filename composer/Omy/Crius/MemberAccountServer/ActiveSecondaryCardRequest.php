<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-account/card.proto

namespace Omy\Crius\MemberAccountServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberAccount.ActiveSecondaryCardRequest</code>
 */
class ActiveSecondaryCardRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string card_id = 1;</code>
     */
    protected $card_id = '';
    /**
     * Generated from protobuf field <code>string branch_id = 2;</code>
     */
    protected $branch_id = '';
    /**
     * Generated from protobuf field <code>string password = 3;</code>
     */
    protected $password = '';
    /**
     * Generated from protobuf field <code>string recommender = 4;</code>
     */
    protected $recommender = '';
    /**
     * Generated from protobuf field <code>int32 recharge_value = 5;</code>
     */
    protected $recharge_value = 0;
    /**
     * Generated from protobuf field <code>string primary_id = 6;</code>
     */
    protected $primary_id = '';
    /**
     * Generated from protobuf field <code>string primary_password = 7;</code>
     */
    protected $primary_password = '';
    /**
     * Generated from protobuf field <code>bool primary_verified = 8;</code>
     */
    protected $primary_verified = false;
    /**
     * Generated from protobuf field <code>string primary_account_id = 9;</code>
     */
    protected $primary_account_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $card_id
     *     @type string $branch_id
     *     @type string $password
     *     @type string $recommender
     *     @type int $recharge_value
     *     @type string $primary_id
     *     @type string $primary_password
     *     @type bool $primary_verified
     *     @type string $primary_account_id
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
     * Generated from protobuf field <code>string branch_id = 2;</code>
     * @return string
     */
    public function getBranchId()
    {
        return $this->branch_id;
    }

    /**
     * Generated from protobuf field <code>string branch_id = 2;</code>
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
     * Generated from protobuf field <code>string password = 3;</code>
     * @return string
     */
    public function getPassword()
    {
        return $this->password;
    }

    /**
     * Generated from protobuf field <code>string password = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setPassword($var)
    {
        GPBUtil::checkString($var, True);
        $this->password = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string recommender = 4;</code>
     * @return string
     */
    public function getRecommender()
    {
        return $this->recommender;
    }

    /**
     * Generated from protobuf field <code>string recommender = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setRecommender($var)
    {
        GPBUtil::checkString($var, True);
        $this->recommender = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 recharge_value = 5;</code>
     * @return int
     */
    public function getRechargeValue()
    {
        return $this->recharge_value;
    }

    /**
     * Generated from protobuf field <code>int32 recharge_value = 5;</code>
     * @param int $var
     * @return $this
     */
    public function setRechargeValue($var)
    {
        GPBUtil::checkInt32($var);
        $this->recharge_value = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string primary_id = 6;</code>
     * @return string
     */
    public function getPrimaryId()
    {
        return $this->primary_id;
    }

    /**
     * Generated from protobuf field <code>string primary_id = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setPrimaryId($var)
    {
        GPBUtil::checkString($var, True);
        $this->primary_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string primary_password = 7;</code>
     * @return string
     */
    public function getPrimaryPassword()
    {
        return $this->primary_password;
    }

    /**
     * Generated from protobuf field <code>string primary_password = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setPrimaryPassword($var)
    {
        GPBUtil::checkString($var, True);
        $this->primary_password = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool primary_verified = 8;</code>
     * @return bool
     */
    public function getPrimaryVerified()
    {
        return $this->primary_verified;
    }

    /**
     * Generated from protobuf field <code>bool primary_verified = 8;</code>
     * @param bool $var
     * @return $this
     */
    public function setPrimaryVerified($var)
    {
        GPBUtil::checkBool($var);
        $this->primary_verified = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string primary_account_id = 9;</code>
     * @return string
     */
    public function getPrimaryAccountId()
    {
        return $this->primary_account_id;
    }

    /**
     * Generated from protobuf field <code>string primary_account_id = 9;</code>
     * @param string $var
     * @return $this
     */
    public function setPrimaryAccountId($var)
    {
        GPBUtil::checkString($var, True);
        $this->primary_account_id = $var;

        return $this;
    }

}

