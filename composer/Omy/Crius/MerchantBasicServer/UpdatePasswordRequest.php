<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/staff.proto

namespace Omy\Crius\MerchantBasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.UpdatePasswordRequest</code>
 */
class UpdatePasswordRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>string old_password = 2;</code>
     */
    protected $old_password = '';
    /**
     * Generated from protobuf field <code>string new_password = 3;</code>
     */
    protected $new_password = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $old_password
     *     @type string $new_password
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\Staff::initOnce();
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
     * Generated from protobuf field <code>string old_password = 2;</code>
     * @return string
     */
    public function getOldPassword()
    {
        return $this->old_password;
    }

    /**
     * Generated from protobuf field <code>string old_password = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setOldPassword($var)
    {
        GPBUtil::checkString($var, True);
        $this->old_password = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string new_password = 3;</code>
     * @return string
     */
    public function getNewPassword()
    {
        return $this->new_password;
    }

    /**
     * Generated from protobuf field <code>string new_password = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setNewPassword($var)
    {
        GPBUtil::checkString($var, True);
        $this->new_password = $var;

        return $this;
    }

}
