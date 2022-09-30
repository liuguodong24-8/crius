<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/branch_business.proto

namespace Omy\Crius\MerchantBasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.BranchBusiness</code>
 */
class BranchBusiness extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>string branch_id = 2;</code>
     */
    protected $branch_id = '';
    /**
     * Generated from protobuf field <code>string begin_date = 3;</code>
     */
    protected $begin_date = '';
    /**
     * Generated from protobuf field <code>string end_date = 4;</code>
     */
    protected $end_date = '';
    /**
     * Generated from protobuf field <code>repeated int32 weeks = 5;</code>
     */
    private $weeks;
    /**
     * Generated from protobuf field <code>string begin_time = 6;</code>
     */
    protected $begin_time = '';
    /**
     * Generated from protobuf field <code>string end_time = 7;</code>
     */
    protected $end_time = '';
    /**
     * Generated from protobuf field <code>bool is_next_day = 8;</code>
     */
    protected $is_next_day = false;
    /**
     * Generated from protobuf field <code>string status = 9;</code>
     */
    protected $status = '';
    /**
     * Generated from protobuf field <code>string category = 10;</code>
     */
    protected $category = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $branch_id
     *     @type string $begin_date
     *     @type string $end_date
     *     @type int[]|\Google\Protobuf\Internal\RepeatedField $weeks
     *     @type string $begin_time
     *     @type string $end_time
     *     @type bool $is_next_day
     *     @type string $status
     *     @type string $category
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\BranchBusiness::initOnce();
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
     * Generated from protobuf field <code>string begin_date = 3;</code>
     * @return string
     */
    public function getBeginDate()
    {
        return $this->begin_date;
    }

    /**
     * Generated from protobuf field <code>string begin_date = 3;</code>
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
     * Generated from protobuf field <code>string end_date = 4;</code>
     * @return string
     */
    public function getEndDate()
    {
        return $this->end_date;
    }

    /**
     * Generated from protobuf field <code>string end_date = 4;</code>
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
     * Generated from protobuf field <code>repeated int32 weeks = 5;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getWeeks()
    {
        return $this->weeks;
    }

    /**
     * Generated from protobuf field <code>repeated int32 weeks = 5;</code>
     * @param int[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setWeeks($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::INT32);
        $this->weeks = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string begin_time = 6;</code>
     * @return string
     */
    public function getBeginTime()
    {
        return $this->begin_time;
    }

    /**
     * Generated from protobuf field <code>string begin_time = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setBeginTime($var)
    {
        GPBUtil::checkString($var, True);
        $this->begin_time = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string end_time = 7;</code>
     * @return string
     */
    public function getEndTime()
    {
        return $this->end_time;
    }

    /**
     * Generated from protobuf field <code>string end_time = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setEndTime($var)
    {
        GPBUtil::checkString($var, True);
        $this->end_time = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool is_next_day = 8;</code>
     * @return bool
     */
    public function getIsNextDay()
    {
        return $this->is_next_day;
    }

    /**
     * Generated from protobuf field <code>bool is_next_day = 8;</code>
     * @param bool $var
     * @return $this
     */
    public function setIsNextDay($var)
    {
        GPBUtil::checkBool($var);
        $this->is_next_day = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string status = 9;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>string status = 9;</code>
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
     * Generated from protobuf field <code>string category = 10;</code>
     * @return string
     */
    public function getCategory()
    {
        return $this->category;
    }

    /**
     * Generated from protobuf field <code>string category = 10;</code>
     * @param string $var
     * @return $this
     */
    public function setCategory($var)
    {
        GPBUtil::checkString($var, True);
        $this->category = $var;

        return $this;
    }

}
