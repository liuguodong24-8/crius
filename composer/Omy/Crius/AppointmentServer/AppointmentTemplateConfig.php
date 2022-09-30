<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment_template.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.AppointmentTemplateConfig</code>
 */
class AppointmentTemplateConfig extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>string room_type_id = 2;</code>
     */
    protected $room_type_id = '';
    /**
     * Generated from protobuf field <code>string template_id = 3;</code>
     */
    protected $template_id = '';
    /**
     * Generated from protobuf field <code>int32 advance_day = 4;</code>
     */
    protected $advance_day = 0;
    /**
     * Generated from protobuf field <code>int32 deposit_fee = 5;</code>
     */
    protected $deposit_fee = 0;
    /**
     * Generated from protobuf field <code>repeated .appointment.TemplateRoomConfigColumn configure = 6;</code>
     */
    private $configure;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $room_type_id
     *     @type string $template_id
     *     @type int $advance_day
     *     @type int $deposit_fee
     *     @type \Omy\Crius\AppointmentServer\TemplateRoomConfigColumn[]|\Google\Protobuf\Internal\RepeatedField $configure
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\AppointmentTemplate::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string id = 1;</code>
     * @return string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Generated from protobuf field <code>string id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkString($var, True);
        $this->id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string room_type_id = 2;</code>
     * @return string
     */
    public function getRoomTypeId()
    {
        return $this->room_type_id;
    }

    /**
     * Generated from protobuf field <code>string room_type_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setRoomTypeId($var)
    {
        GPBUtil::checkString($var, True);
        $this->room_type_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string template_id = 3;</code>
     * @return string
     */
    public function getTemplateId()
    {
        return $this->template_id;
    }

    /**
     * Generated from protobuf field <code>string template_id = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setTemplateId($var)
    {
        GPBUtil::checkString($var, True);
        $this->template_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 advance_day = 4;</code>
     * @return int
     */
    public function getAdvanceDay()
    {
        return $this->advance_day;
    }

    /**
     * Generated from protobuf field <code>int32 advance_day = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setAdvanceDay($var)
    {
        GPBUtil::checkInt32($var);
        $this->advance_day = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 deposit_fee = 5;</code>
     * @return int
     */
    public function getDepositFee()
    {
        return $this->deposit_fee;
    }

    /**
     * Generated from protobuf field <code>int32 deposit_fee = 5;</code>
     * @param int $var
     * @return $this
     */
    public function setDepositFee($var)
    {
        GPBUtil::checkInt32($var);
        $this->deposit_fee = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .appointment.TemplateRoomConfigColumn configure = 6;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getConfigure()
    {
        return $this->configure;
    }

    /**
     * Generated from protobuf field <code>repeated .appointment.TemplateRoomConfigColumn configure = 6;</code>
     * @param \Omy\Crius\AppointmentServer\TemplateRoomConfigColumn[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setConfigure($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\AppointmentServer\TemplateRoomConfigColumn::class);
        $this->configure = $arr;

        return $this;
    }

}

