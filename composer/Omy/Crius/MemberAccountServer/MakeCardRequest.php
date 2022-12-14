<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-account/card.proto

namespace Omy\Crius\MemberAccountServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberAccount.MakeCardRequest</code>
 */
class MakeCardRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string category = 1;</code>
     */
    protected $category = '';
    /**
     * Generated from protobuf field <code>string branch_id = 2;</code>
     */
    protected $branch_id = '';
    /**
     * Generated from protobuf field <code>string code = 3;</code>
     */
    protected $code = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $category
     *     @type string $branch_id
     *     @type string $code
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberAccount\Card::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string category = 1;</code>
     * @return string
     */
    public function getCategory()
    {
        return $this->category;
    }

    /**
     * Generated from protobuf field <code>string category = 1;</code>
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
     * Generated from protobuf field <code>string code = 3;</code>
     * @return string
     */
    public function getCode()
    {
        return $this->code;
    }

    /**
     * Generated from protobuf field <code>string code = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->code = $var;

        return $this;
    }

}

