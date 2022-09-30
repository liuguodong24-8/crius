<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/assist/assist.proto

namespace Omy\Crius\AssistServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * PosBillBandRequest 咨客订单绑定预约订单
 *
 * Generated from protobuf message <code>assist.PosBillBandRequest</code>
 */
class PosBillBandRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * 咨客账单信息
     *
     * Generated from protobuf field <code>string bill_id = 1;</code>
     */
    protected $bill_id = '';
    /**
     * 预约订单信息
     *
     * Generated from protobuf field <code>string appointment_id = 2;</code>
     */
    protected $appointment_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $bill_id
     *           咨客账单信息
     *     @type string $appointment_id
     *           预约订单信息
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
     * 预约订单信息
     *
     * Generated from protobuf field <code>string appointment_id = 2;</code>
     * @return string
     */
    public function getAppointmentId()
    {
        return $this->appointment_id;
    }

    /**
     * 预约订单信息
     *
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

}
