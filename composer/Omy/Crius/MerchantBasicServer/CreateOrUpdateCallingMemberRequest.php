<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/member.proto

namespace Omy\Crius\MerchantBasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.CreateOrUpdateCallingMemberRequest</code>
 */
class CreateOrUpdateCallingMemberRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string name = 1;</code>
     */
    protected $name = '';
    /**
     * Generated from protobuf field <code>string phone = 2;</code>
     */
    protected $phone = '';
    /**
     * Generated from protobuf field <code>string phone_code = 3;</code>
     */
    protected $phone_code = '';
    /**
     * Generated from protobuf field <code>int32 gender = 4;</code>
     */
    protected $gender = 0;
    /**
     * Generated from protobuf field <code>string channel = 5;</code>
     */
    protected $channel = '';
    /**
     * Generated from protobuf field <code>bool can_overwrite = 6;</code>
     */
    protected $can_overwrite = false;
    /**
     * Generated from protobuf field <code>string branch_id = 7;</code>
     */
    protected $branch_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *     @type string $phone
     *     @type string $phone_code
     *     @type int $gender
     *     @type string $channel
     *     @type bool $can_overwrite
     *     @type string $branch_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\Member::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string name = 1;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * Generated from protobuf field <code>string name = 1;</code>
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
     * Generated from protobuf field <code>string phone = 2;</code>
     * @return string
     */
    public function getPhone()
    {
        return $this->phone;
    }

    /**
     * Generated from protobuf field <code>string phone = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setPhone($var)
    {
        GPBUtil::checkString($var, True);
        $this->phone = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string phone_code = 3;</code>
     * @return string
     */
    public function getPhoneCode()
    {
        return $this->phone_code;
    }

    /**
     * Generated from protobuf field <code>string phone_code = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setPhoneCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->phone_code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 gender = 4;</code>
     * @return int
     */
    public function getGender()
    {
        return $this->gender;
    }

    /**
     * Generated from protobuf field <code>int32 gender = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setGender($var)
    {
        GPBUtil::checkInt32($var);
        $this->gender = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string channel = 5;</code>
     * @return string
     */
    public function getChannel()
    {
        return $this->channel;
    }

    /**
     * Generated from protobuf field <code>string channel = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setChannel($var)
    {
        GPBUtil::checkString($var, True);
        $this->channel = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool can_overwrite = 6;</code>
     * @return bool
     */
    public function getCanOverwrite()
    {
        return $this->can_overwrite;
    }

    /**
     * Generated from protobuf field <code>bool can_overwrite = 6;</code>
     * @param bool $var
     * @return $this
     */
    public function setCanOverwrite($var)
    {
        GPBUtil::checkBool($var);
        $this->can_overwrite = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string branch_id = 7;</code>
     * @return string
     */
    public function getBranchId()
    {
        return $this->branch_id;
    }

    /**
     * Generated from protobuf field <code>string branch_id = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setBranchId($var)
    {
        GPBUtil::checkString($var, True);
        $this->branch_id = $var;

        return $this;
    }

}

