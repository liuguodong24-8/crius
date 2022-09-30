<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/room_type_group.proto

namespace Omy\Crius\MerchantBasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.GetRoomTypeGroupsByRoomTypeIDRequest</code>
 */
class GetRoomTypeGroupsByRoomTypeIDRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string room_type_id = 1;</code>
     */
    protected $room_type_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $room_type_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\RoomTypeGroup::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string room_type_id = 1;</code>
     * @return string
     */
    public function getRoomTypeId()
    {
        return $this->room_type_id;
    }

    /**
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

}

