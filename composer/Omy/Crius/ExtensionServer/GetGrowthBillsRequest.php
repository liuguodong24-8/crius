<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/growth_bill.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.GetGrowthBillsRequest</code>
 */
class GetGrowthBillsRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string member_id = 1;</code>
     */
    protected $member_id = '';
    /**
     * Generated from protobuf field <code>string category = 2;</code>
     */
    protected $category = '';
    /**
     * Generated from protobuf field <code>int32 dateStart = 3;</code>
     */
    protected $dateStart = 0;
    /**
     * Generated from protobuf field <code>int32 dateEnd = 4;</code>
     */
    protected $dateEnd = 0;
    /**
     * Generated from protobuf field <code>int32 offset = 5;</code>
     */
    protected $offset = 0;
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
     *     @type string $member_id
     *     @type string $category
     *     @type int $dateStart
     *     @type int $dateEnd
     *     @type int $offset
     *     @type int $limit
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\GrowthBill::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string member_id = 1;</code>
     * @return string
     */
    public function getMemberId()
    {
        return $this->member_id;
    }

    /**
     * Generated from protobuf field <code>string member_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setMemberId($var)
    {
        GPBUtil::checkString($var, True);
        $this->member_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string category = 2;</code>
     * @return string
     */
    public function getCategory()
    {
        return $this->category;
    }

    /**
     * Generated from protobuf field <code>string category = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setCategory($var)
    {
        GPBUtil::checkString($var, True);
        $this->category = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 dateStart = 3;</code>
     * @return int
     */
    public function getDateStart()
    {
        return $this->dateStart;
    }

    /**
     * Generated from protobuf field <code>int32 dateStart = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setDateStart($var)
    {
        GPBUtil::checkInt32($var);
        $this->dateStart = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 dateEnd = 4;</code>
     * @return int
     */
    public function getDateEnd()
    {
        return $this->dateEnd;
    }

    /**
     * Generated from protobuf field <code>int32 dateEnd = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setDateEnd($var)
    {
        GPBUtil::checkInt32($var);
        $this->dateEnd = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 offset = 5;</code>
     * @return int
     */
    public function getOffset()
    {
        return $this->offset;
    }

    /**
     * Generated from protobuf field <code>int32 offset = 5;</code>
     * @param int $var
     * @return $this
     */
    public function setOffset($var)
    {
        GPBUtil::checkInt32($var);
        $this->offset = $var;

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
