<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-account/card.proto

namespace Omy\Crius\MemberAccountServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberAccount.RechargeCardRequest</code>
 */
class RechargeCardRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 recharge_value = 1;</code>
     */
    protected $recharge_value = 0;
    /**
     * Generated from protobuf field <code>string branch_id = 2;</code>
     */
    protected $branch_id = '';
    /**
     * Generated from protobuf field <code>repeated .memberAccount.PromotionCount promotions = 3;</code>
     */
    private $promotions;
    /**
     * Generated from protobuf field <code>.memberAccount.Payments payments = 4;</code>
     */
    protected $payments = null;
    /**
     * Generated from protobuf field <code>string recommender = 5;</code>
     */
    protected $recommender = '';
    /**
     * Generated from protobuf field <code>string card_id = 6;</code>
     */
    protected $card_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $recharge_value
     *     @type string $branch_id
     *     @type \Omy\Crius\MemberAccountServer\PromotionCount[]|\Google\Protobuf\Internal\RepeatedField $promotions
     *     @type \Omy\Crius\MemberAccountServer\Payments $payments
     *     @type string $recommender
     *     @type string $card_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberAccount\Card::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 recharge_value = 1;</code>
     * @return int
     */
    public function getRechargeValue()
    {
        return $this->recharge_value;
    }

    /**
     * Generated from protobuf field <code>int32 recharge_value = 1;</code>
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
     * Generated from protobuf field <code>repeated .memberAccount.PromotionCount promotions = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getPromotions()
    {
        return $this->promotions;
    }

    /**
     * Generated from protobuf field <code>repeated .memberAccount.PromotionCount promotions = 3;</code>
     * @param \Omy\Crius\MemberAccountServer\PromotionCount[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setPromotions($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MemberAccountServer\PromotionCount::class);
        $this->promotions = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.memberAccount.Payments payments = 4;</code>
     * @return \Omy\Crius\MemberAccountServer\Payments|null
     */
    public function getPayments()
    {
        return $this->payments;
    }

    public function hasPayments()
    {
        return isset($this->payments);
    }

    public function clearPayments()
    {
        unset($this->payments);
    }

    /**
     * Generated from protobuf field <code>.memberAccount.Payments payments = 4;</code>
     * @param \Omy\Crius\MemberAccountServer\Payments $var
     * @return $this
     */
    public function setPayments($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\MemberAccountServer\Payments::class);
        $this->payments = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string recommender = 5;</code>
     * @return string
     */
    public function getRecommender()
    {
        return $this->recommender;
    }

    /**
     * Generated from protobuf field <code>string recommender = 5;</code>
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
     * Generated from protobuf field <code>string card_id = 6;</code>
     * @return string
     */
    public function getCardId()
    {
        return $this->card_id;
    }

    /**
     * Generated from protobuf field <code>string card_id = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setCardId($var)
    {
        GPBUtil::checkString($var, True);
        $this->card_id = $var;

        return $this;
    }

}

