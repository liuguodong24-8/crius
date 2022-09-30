<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/assist/assist.proto

namespace Omy\Crius\AssistServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * PayOrderRequest 订单预约付款
 *
 * Generated from protobuf message <code>assist.PayOrderRequest</code>
 */
class PayOrderRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string appointment_id = 1;</code>
     */
    protected $appointment_id = '';
    /**
     * Generated from protobuf field <code>string old_appointment_id = 2;</code>
     */
    protected $old_appointment_id = '';
    /**
     * Generated from protobuf field <code>string old_order_state = 3;</code>
     */
    protected $old_order_state = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $appointment_id
     *     @type string $old_appointment_id
     *     @type string $old_order_state
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Assist\Assist::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string appointment_id = 1;</code>
     * @return string
     */
    public function getAppointmentId()
    {
        return $this->appointment_id;
    }

    /**
     * Generated from protobuf field <code>string appointment_id = 1;</code>
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
     * Generated from protobuf field <code>string old_appointment_id = 2;</code>
     * @return string
     */
    public function getOldAppointmentId()
    {
        return $this->old_appointment_id;
    }

    /**
     * Generated from protobuf field <code>string old_appointment_id = 2;</code>
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
     * Generated from protobuf field <code>string old_order_state = 3;</code>
     * @return string
     */
    public function getOldOrderState()
    {
        return $this->old_order_state;
    }

    /**
     * Generated from protobuf field <code>string old_order_state = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setOldOrderState($var)
    {
        GPBUtil::checkString($var, True);
        $this->old_order_state = $var;

        return $this;
    }

}
