<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/role.proto

namespace Omy\Crius\MerchantBasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.RoleInfo</code>
 */
class RoleInfo extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>string name = 2;</code>
     */
    protected $name = '';
    /**
     * Generated from protobuf field <code>string status = 3;</code>
     */
    protected $status = '';
    /**
     * Generated from protobuf field <code>int32 property = 4;</code>
     */
    protected $property = 0;
    /**
     * Generated from protobuf field <code>string staff_id = 5;</code>
     */
    protected $staff_id = '';
    /**
     * Generated from protobuf field <code>string staff_name = 6;</code>
     */
    protected $staff_name = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $name
     *     @type string $status
     *     @type int $property
     *     @type string $staff_id
     *     @type string $staff_name
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\Role::initOnce();
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
     * Generated from protobuf field <code>string name = 2;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * Generated from protobuf field <code>string name = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string status = 3;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>string status = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkString($var, True);
        $this->status = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 property = 4;</code>
     * @return int
     */
    public function getProperty()
    {
        return $this->property;
    }

    /**
     * Generated from protobuf field <code>int32 property = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setProperty($var)
    {
        GPBUtil::checkInt32($var);
        $this->property = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string staff_id = 5;</code>
     * @return string
     */
    public function getStaffId()
    {
        return $this->staff_id;
    }

    /**
     * Generated from protobuf field <code>string staff_id = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setStaffId($var)
    {
        GPBUtil::checkString($var, True);
        $this->staff_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string staff_name = 6;</code>
     * @return string
     */
    public function getStaffName()
    {
        return $this->staff_name;
    }

    /**
     * Generated from protobuf field <code>string staff_name = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setStaffName($var)
    {
        GPBUtil::checkString($var, True);
        $this->staff_name = $var;

        return $this;
    }

}

