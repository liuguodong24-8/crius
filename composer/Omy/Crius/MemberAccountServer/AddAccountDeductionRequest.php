<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-account/member_account.proto

namespace Omy\Crius\MemberAccountServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberAccount.AddAccountDeductionRequest</code>
 */
class AddAccountDeductionRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string bill_number = 1;</code>
     */
    protected $bill_number = '';
    /**
     * Generated from protobuf field <code>string card_id = 2;</code>
     */
    protected $card_id = '';
    /**
     * Generated from protobuf field <code>string branch_id = 3;</code>
     */
    protected $branch_id = '';
    /**
     * Generated from protobuf field <code>int32 cost_value = 4;</code>
     */
    protected $cost_value = 0;
    /**
     * Generated from protobuf field <code>string reason = 5;</code>
     */
    protected $reason = '';
    /**
     * Generated from protobuf field <code>repeated .memberAccount.CostProductPackageTicketItem products = 6;</code>
     */
    private $products;
    /**
     * Generated from protobuf field <code>repeated .memberAccount.CostProductPackageTicketItem packages = 7;</code>
     */
    private $packages;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $bill_number
     *     @type string $card_id
     *     @type string $branch_id
     *     @type int $cost_value
     *     @type string $reason
     *     @type \Omy\Crius\MemberAccountServer\CostProductPackageTicketItem[]|\Google\Protobuf\Internal\RepeatedField $products
     *     @type \Omy\Crius\MemberAccountServer\CostProductPackageTicketItem[]|\Google\Protobuf\Internal\RepeatedField $packages
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberAccount\MemberAccount::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string bill_number = 1;</code>
     * @return string
     */
    public function getBillNumber()
    {
        return $this->bill_number;
    }

    /**
     * Generated from protobuf field <code>string bill_number = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setBillNumber($var)
    {
        GPBUtil::checkString($var, True);
        $this->bill_number = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string card_id = 2;</code>
     * @return string
     */
    public function getCardId()
    {
        return $this->card_id;
    }

    /**
     * Generated from protobuf field <code>string card_id = 2;</code>
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
     * Generated from protobuf field <code>string branch_id = 3;</code>
     * @return string
     */
    public function getBranchId()
    {
        return $this->branch_id;
    }

    /**
     * Generated from protobuf field <code>string branch_id = 3;</code>
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
     * Generated from protobuf field <code>int32 cost_value = 4;</code>
     * @return int
     */
    public function getCostValue()
    {
        return $this->cost_value;
    }

    /**
     * Generated from protobuf field <code>int32 cost_value = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setCostValue($var)
    {
        GPBUtil::checkInt32($var);
        $this->cost_value = $var;

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

    /**
     * Generated from protobuf field <code>repeated .memberAccount.CostProductPackageTicketItem products = 6;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getProducts()
    {
        return $this->products;
    }

    /**
     * Generated from protobuf field <code>repeated .memberAccount.CostProductPackageTicketItem products = 6;</code>
     * @param \Omy\Crius\MemberAccountServer\CostProductPackageTicketItem[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setProducts($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MemberAccountServer\CostProductPackageTicketItem::class);
        $this->products = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .memberAccount.CostProductPackageTicketItem packages = 7;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getPackages()
    {
        return $this->packages;
    }

    /**
     * Generated from protobuf field <code>repeated .memberAccount.CostProductPackageTicketItem packages = 7;</code>
     * @param \Omy\Crius\MemberAccountServer\CostProductPackageTicketItem[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setPackages($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MemberAccountServer\CostProductPackageTicketItem::class);
        $this->packages = $arr;

        return $this;
    }

}
