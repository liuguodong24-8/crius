<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/assist/assist.proto

namespace Omy\Crius\AssistServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * DeleteOrderRequest 删除订单
 *
 * Generated from protobuf message <code>assist.DeleteOrderRequest</code>
 */
class DeleteOrderRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string appointment_id = 1;</code>
     */
    protected $appointment_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $appointment_id
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

}

