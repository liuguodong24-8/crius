<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment_theme.proto

namespace Omy\Crius\AppointmentServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>appointment.GetAppointmentThemesByRoomTypeRequest</code>
 */
class GetAppointmentThemesByRoomTypeRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * 房型id
     *
     * Generated from protobuf field <code>string room_type_id = 1;</code>
     */
    protected $room_type_id = '';
    /**
     * 门店id
     *
     * Generated from protobuf field <code>string branch_id = 2;</code>
     */
    protected $branch_id = '';
    /**
     * 查询日期
     *
     * Generated from protobuf field <code>int32 date = 3;</code>
     */
    protected $date = 0;
    /**
     * 主题id
     *
     * Generated from protobuf field <code>string theme_id = 4;</code>
     */
    protected $theme_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $room_type_id
     *           房型id
     *     @type string $branch_id
     *           门店id
     *     @type int $date
     *           查询日期
     *     @type string $theme_id
     *           主题id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Appointment\AppointmentTheme::initOnce();
        parent::__construct($data);
    }

    /**
     * 房型id
     *
     * Generated from protobuf field <code>string room_type_id = 1;</code>
     * @return string
     */
    public function getRoomTypeId()
    {
        return $this->room_type_id;
    }

    /**
     * 房型id
     *
     * Generated from protobuf field <code>string room_type_id = 1;</code>
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
     * 门店id
     *
     * Generated from protobuf field <code>string branch_id = 2;</code>
     * @return string
     */
    public function getBranchId()
    {
        return $this->branch_id;
    }

    /**
     * 门店id
     *
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
     * 查询日期
     *
     * Generated from protobuf field <code>int32 date = 3;</code>
     * @return int
     */
    public function getDate()
    {
        return $this->date;
    }

    /**
     * 查询日期
     *
     * Generated from protobuf field <code>int32 date = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setDate($var)
    {
        GPBUtil::checkInt32($var);
        $this->date = $var;

        return $this;
    }

    /**
     * 主题id
     *
     * Generated from protobuf field <code>string theme_id = 4;</code>
     * @return string
     */
    public function getThemeId()
    {
        return $this->theme_id;
    }

    /**
     * 主题id
     *
     * Generated from protobuf field <code>string theme_id = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setThemeId($var)
    {
        GPBUtil::checkString($var, True);
        $this->theme_id = $var;

        return $this;
    }

}

