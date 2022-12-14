<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/member.proto

namespace Omy\Crius\MerchantBasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.GetCouponMemberIDsRequest</code>
 */
class GetCouponMemberIDsRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated string city_codes = 1;</code>
     */
    private $city_codes;
    /**
     * Generated from protobuf field <code>repeated string branch_ids = 2;</code>
     */
    private $branch_ids;
    /**
     * Generated from protobuf field <code>int32 gender = 3;</code>
     */
    protected $gender = 0;
    /**
     * Generated from protobuf field <code>repeated string phones = 4;</code>
     */
    private $phones;
    /**
     * Generated from protobuf field <code>int64 create_at = 5;</code>
     */
    protected $create_at = 0;
    /**
     * Generated from protobuf field <code>int32 limit = 6;</code>
     */
    protected $limit = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $city_codes
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $branch_ids
     *     @type int $gender
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $phones
     *     @type int|string $create_at
     *     @type int $limit
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\Member::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated string city_codes = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getCityCodes()
    {
        return $this->city_codes;
    }

    /**
     * Generated from protobuf field <code>repeated string city_codes = 1;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setCityCodes($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->city_codes = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string branch_ids = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getBranchIds()
    {
        return $this->branch_ids;
    }

    /**
     * Generated from protobuf field <code>repeated string branch_ids = 2;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setBranchIds($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->branch_ids = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 gender = 3;</code>
     * @return int
     */
    public function getGender()
    {
        return $this->gender;
    }

    /**
     * Generated from protobuf field <code>int32 gender = 3;</code>
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
     * Generated from protobuf field <code>repeated string phones = 4;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getPhones()
    {
        return $this->phones;
    }

    /**
     * Generated from protobuf field <code>repeated string phones = 4;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setPhones($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->phones = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 create_at = 5;</code>
     * @return int|string
     */
    public function getCreateAt()
    {
        return $this->create_at;
    }

    /**
     * Generated from protobuf field <code>int64 create_at = 5;</code>
     * @param int|string $var
     * @return $this
     */
    public function setCreateAt($var)
    {
        GPBUtil::checkInt64($var);
        $this->create_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 limit = 6;</code>
     * @return int
     */
    public function getLimit()
    {
        return $this->limit;
    }

    /**
     * Generated from protobuf field <code>int32 limit = 6;</code>
     * @param int $var
     * @return $this
     */
    public function setLimit($var)
    {
        GPBUtil::checkInt32($var);
        $this->limit = $var;

        return $this;
    }

}

