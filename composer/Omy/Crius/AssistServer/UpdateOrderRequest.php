<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/assist/assist.proto

namespace Omy\Crius\AssistServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * UpdateOrderRequest 更新订单
 *
 * Generated from protobuf message <code>assist.UpdateOrderRequest</code>
 */
class UpdateOrderRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string old_appointment_id = 1;</code>
     */
    protected $old_appointment_id = '';
    /**
     * Generated from protobuf field <code>string appointment_id = 2;</code>
     */
    protected $appointment_id = '';
    /**
     * Generated from protobuf field <code>string branch_id = 3;</code>
     */
    protected $branch_id = '';
    /**
     * Generated from protobuf field <code>string member_id = 4;</code>
     */
    protected $member_id = '';
    /**
     * Generated from protobuf field <code>string order_state = 5;</code>
     */
    protected $order_state = '';
    /**
     * Generated from protobuf field <code>int32 fee = 6;</code>
     */
    protected $fee = 0;
    /**
     * Generated from protobuf field <code>string old_order_state = 7;</code>
     */
    protected $old_order_state = '';
    /**
     * Generated from protobuf field <code>string called_phone = 8;</code>
     */
    protected $called_phone = '';
    /**
     * Generated from protobuf field <code>string called_code = 9;</code>
     */
    protected $called_code = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $old_appointment_id
     *     @type string $appointment_id
     *     @type string $branch_id
     *     @type string $member_id
     *     @type string $order_state
     *     @type int $fee
     *     @type string $old_order_state
     *     @type string $called_phone
     *     @type string $called_code
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Assist\Assist::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string old_appointment_id = 1;</code>
     * @return string
     */
    public function getOldAppointmentId()
    {
        return $this->old_appointment_id;
    }

    /**
     * Generated from protobuf field <code>string old_appointment_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setOldAppointmentId($var)
    {
        GPBUtil::checkString($var, True);
        $this->old_appointment_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string appointment_id = 2;</code>
     * @return string
     */
    public function getAppointmentId()
    {
        return $this->appointment_id;
    }

    /**
     * Generated from protobuf field <code>string appointment_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setAppointmentId($var)
    {
        GPBUtil::checkString($var, True);
        $this->appointment_id = $var;

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
     * Generated from protobuf field <code>string member_id = 4;</code>
     * @return string
     */
    public function getMemberId()
    {
        return $this->member_id;
    }

    /**
     * Generated from protobuf field <code>string member_id = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setMemberId($var)
    {
        GPBUtil::checkString($var, True);
        $this->member_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string order_state = 5;</code>
     * @return string
     */
    public function getOrderState()
    {
        return $this->order_state;
    }

    /**
     * Generated from protobuf field <code>string order_state = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setOrderState($var)
    {
        GPBUtil::checkString($var, True);
        $this->order_state = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 fee = 6;</code>
     * @return int
     */
    public function getFee()
    {
        return $this->fee;
    }

    /**
     * Generated from protobuf field <code>int32 fee = 6;</code>
     * @param int $var
     * @return $this
     */
    public function setFee($var)
    {
        GPBUtil::checkInt32($var);
        $this->fee = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string old_order_state = 7;</code>
     * @return string
     */
    public function getOldOrderState()
    {
        return $this->old_order_state;
    }

    /**
     * Generated from protobuf field <code>string old_order_state = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setOldOrderState($var)
    {
        GPBUtil::checkString($var, True);
        $this->old_order_state = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string called_phone = 8;</code>
     * @return string
     */
    public function getCalledPhone()
    {
        return $this->called_phone;
    }

    /**
     * Generated from protobuf field <code>string called_phone = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setCalledPhone($var)
    {
        GPBUtil::checkString($var, True);
        $this->called_phone = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string called_code = 9;</code>
     * @return string
     */
    public function getCalledCode()
    {
        return $this->called_code;
    }

    /**
     * Generated from protobuf field <code>string called_code = 9;</code>
     * @param string $var
     * @return $this
     */
    public function setCalledCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->called_code = $var;

        return $this;
    }

}
