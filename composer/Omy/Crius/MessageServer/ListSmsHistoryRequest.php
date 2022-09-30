<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/message/sms_stat.proto

namespace Omy\Crius\MessageServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>message.ListSmsHistoryRequest</code>
 */
class ListSmsHistoryRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated string ids = 1;</code>
     */
    private $ids;
    /**
     * Generated from protobuf field <code>string branch_id = 2;</code>
     */
    protected $branch_id = '';
    /**
     * Generated from protobuf field <code>string message_type = 3;</code>
     */
    protected $message_type = '';
    /**
     * Generated from protobuf field <code>string begin_date = 4;</code>
     */
    protected $begin_date = '';
    /**
     * Generated from protobuf field <code>string end_date = 5;</code>
     */
    protected $end_date = '';
    /**
     * Generated from protobuf field <code>string sms_status = 6;</code>
     */
    protected $sms_status = '';
    /**
     * Generated from protobuf field <code>string phone = 7;</code>
     */
    protected $phone = '';
    /**
     * Generated from protobuf field <code>int32 limit = 8;</code>
     */
    protected $limit = 0;
    /**
     * Generated from protobuf field <code>int32 offset = 9;</code>
     */
    protected $offset = 0;
    /**
     * Generated from protobuf field <code>string order_by = 10;</code>
     */
    protected $order_by = '';
    /**
     * Generated from protobuf field <code>bool with_page = 11;</code>
     */
    protected $with_page = false;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $ids
     *     @type string $branch_id
     *     @type string $message_type
     *     @type string $begin_date
     *     @type string $end_date
     *     @type string $sms_status
     *     @type string $phone
     *     @type int $limit
     *     @type int $offset
     *     @type string $order_by
     *     @type bool $with_page
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Message\SmsStat::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated string ids = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getIds()
    {
        return $this->ids;
    }

    /**
     * Generated from protobuf field <code>repeated string ids = 1;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setIds($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->ids = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string branch_id = 2;</code>
     * @return string
     */
    public function getBranchId()
    {
        return $this->branch_id;
    }

    /**
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
     * Generated from protobuf field <code>string message_type = 3;</code>
     * @return string
     */
    public function getMessageType()
    {
        return $this->message_type;
    }

    /**
     * Generated from protobuf field <code>string message_type = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setMessageType($var)
    {
        GPBUtil::checkString($var, True);
        $this->message_type = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string begin_date = 4;</code>
     * @return string
     */
    public function getBeginDate()
    {
        return $this->begin_date;
    }

    /**
     * Generated from protobuf field <code>string begin_date = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setBeginDate($var)
    {
        GPBUtil::checkString($var, True);
        $this->begin_date = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string end_date = 5;</code>
     * @return string
     */
    public function getEndDate()
    {
        return $this->end_date;
    }

    /**
     * Generated from protobuf field <code>string end_date = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setEndDate($var)
    {
        GPBUtil::checkString($var, True);
        $this->end_date = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string sms_status = 6;</code>
     * @return string
     */
    public function getSmsStatus()
    {
        return $this->sms_status;
    }

    /**
     * Generated from protobuf field <code>string sms_status = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setSmsStatus($var)
    {
        GPBUtil::checkString($var, True);
        $this->sms_status = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string phone = 7;</code>
     * @return string
     */
    public function getPhone()
    {
        return $this->phone;
    }

    /**
     * Generated from protobuf field <code>string phone = 7;</code>
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
     * Generated from protobuf field <code>int32 limit = 8;</code>
     * @return int
     */
    public function getLimit()
    {
        return $this->limit;
    }

    /**
     * Generated from protobuf field <code>int32 limit = 8;</code>
     * @param int $var
     * @return $this
     */
    public function setLimit($var)
    {
        GPBUtil::checkInt32($var);
        $this->limit = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 offset = 9;</code>
     * @return int
     */
    public function getOffset()
    {
        return $this->offset;
    }

    /**
     * Generated from protobuf field <code>int32 offset = 9;</code>
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
     * Generated from protobuf field <code>string order_by = 10;</code>
     * @return string
     */
    public function getOrderBy()
    {
        return $this->order_by;
    }

    /**
     * Generated from protobuf field <code>string order_by = 10;</code>
     * @param string $var
     * @return $this
     */
    public function setOrderBy($var)
    {
        GPBUtil::checkString($var, True);
        $this->order_by = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool with_page = 11;</code>
     * @return bool
     */
    public function getWithPage()
    {
        return $this->with_page;
    }

    /**
     * Generated from protobuf field <code>bool with_page = 11;</code>
     * @param bool $var
     * @return $this
     */
    public function setWithPage($var)
    {
        GPBUtil::checkBool($var);
        $this->with_page = $var;

        return $this;
    }

}
