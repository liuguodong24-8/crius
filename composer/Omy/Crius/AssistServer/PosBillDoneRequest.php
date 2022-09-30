<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/assist/assist.proto

namespace Omy\Crius\AssistServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * PosBillDoneRequest 账单完成
 *
 * Generated from protobuf message <code>assist.PosBillDoneRequest</code>
 */
class PosBillDoneRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * 咨客账单信息
     *
     * Generated from protobuf field <code>string bill_id = 1;</code>
     */
    protected $bill_id = '';
    /**
     * 账单金额
     *
     * Generated from protobuf field <code>int64 bill_fee = 2;</code>
     */
    protected $bill_fee = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $bill_id
     *           咨客账单信息
     *     @type int|string $bill_fee
     *           账单金额
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Assist\Assist::initOnce();
        parent::__construct($data);
    }

    /**
     * 咨客账单信息
     *
     * Generated from protobuf field <code>string bill_id = 1;</code>
     * @return string
     */
    public function getBillId()
    {
        return $this->bill_id;
    }

    /**
     * 咨客账单信息
     *
     * Generated from protobuf field <code>string bill_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setBillId($var)
    {
        GPBUtil::checkString($var, True);
        $this->bill_id = $var;

        return $this;
    }

    /**
     * 账单金额
     *
     * Generated from protobuf field <code>int64 bill_fee = 2;</code>
     * @return int|string
     */
    public function getBillFee()
    {
        return $this->bill_fee;
    }

    /**
     * 账单金额
     *
     * Generated from protobuf field <code>int64 bill_fee = 2;</code>
     * @param int|string $var
     * @return $this
     */
    public function setBillFee($var)
    {
        GPBUtil::checkInt64($var);
        $this->bill_fee = $var;

        return $this;
    }

}

