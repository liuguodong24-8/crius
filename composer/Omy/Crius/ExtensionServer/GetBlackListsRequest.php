<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/black_list.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.GetBlackListsRequest</code>
 */
class GetBlackListsRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string phone_suffix = 1;</code>
     */
    protected $phone_suffix = '';
    /**
     * Generated from protobuf field <code>string phone = 2;</code>
     */
    protected $phone = '';
    /**
     * Generated from protobuf field <code>string name = 3;</code>
     */
    protected $name = '';
    /**
     * Generated from protobuf field <code>int32 offset = 4;</code>
     */
    protected $offset = 0;
    /**
     * Generated from protobuf field <code>int32 limit = 5;</code>
     */
    protected $limit = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $phone_suffix
     *     @type string $phone
     *     @type string $name
     *     @type int $offset
     *     @type int $limit
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\BlackList::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string phone_suffix = 1;</code>
     * @return string
     */
    public function getPhoneSuffix()
    {
        return $this->phone_suffix;
    }

    /**
     * Generated from protobuf field <code>string phone_suffix = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setPhoneSuffix($var)
    {
        GPBUtil::checkString($var, True);
        $this->phone_suffix = $var;

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
     * Generated from protobuf field <code>string name = 3;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * Generated from protobuf field <code>string name = 3;</code>
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
     * Generated from protobuf field <code>int32 offset = 4;</code>
     * @return int
     */
    public function getOffset()
    {
        return $this->offset;
    }

    /**
     * Generated from protobuf field <code>int32 offset = 4;</code>
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
     * Generated from protobuf field <code>int32 limit = 5;</code>
     * @return int
     */
    public function getLimit()
    {
        return $this->limit;
    }

    /**
     * Generated from protobuf field <code>int32 limit = 5;</code>
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

