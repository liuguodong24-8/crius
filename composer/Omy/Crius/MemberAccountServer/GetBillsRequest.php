<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-account/member_account.proto

namespace Omy\Crius\MemberAccountServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberAccount.GetBillsRequest</code>
 */
class GetBillsRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string account_id = 1;</code>
     */
    protected $account_id = '';
    /**
     * Generated from protobuf field <code>string category = 2;</code>
     */
    protected $category = '';
    /**
     * Generated from protobuf field <code>int32 offset = 3;</code>
     */
    protected $offset = 0;
    /**
     * Generated from protobuf field <code>int32 limit = 4;</code>
     */
    protected $limit = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $account_id
     *     @type string $category
     *     @type int $offset
     *     @type int $limit
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberAccount\MemberAccount::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string account_id = 1;</code>
     * @return string
     */
    public function getAccountId()
    {
        return $this->account_id;
    }

    /**
     * Generated from protobuf field <code>string account_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setAccountId($var)
    {
        GPBUtil::checkString($var, True);
        $this->account_id = $var;

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
     * Generated from protobuf field <code>int32 offset = 3;</code>
     * @return int
     */
    public function getOffset()
    {
        return $this->offset;
    }

    /**
     * Generated from protobuf field <code>int32 offset = 3;</code>
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
     * Generated from protobuf field <code>int32 limit = 4;</code>
     * @return int
     */
    public function getLimit()
    {
        return $this->limit;
    }

    /**
     * Generated from protobuf field <code>int32 limit = 4;</code>
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

