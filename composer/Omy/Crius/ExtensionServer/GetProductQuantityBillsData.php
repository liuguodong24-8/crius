<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/products.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.GetProductQuantityBillsData</code>
 */
class GetProductQuantityBillsData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string code = 1;</code>
     */
    protected $code = '';
    /**
     * Generated from protobuf field <code>int32 created_at = 2;</code>
     */
    protected $created_at = 0;
    /**
     * Generated from protobuf field <code>int32 change = 3;</code>
     */
    protected $change = 0;
    /**
     * Generated from protobuf field <code>int32 quantity = 4;</code>
     */
    protected $quantity = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $code
     *     @type int $created_at
     *     @type int $change
     *     @type int $quantity
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\Products::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string code = 1;</code>
     * @return string
     */
    public function getCode()
    {
        return $this->code;
    }

    /**
     * Generated from protobuf field <code>string code = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 created_at = 2;</code>
     * @return int
     */
    public function getCreatedAt()
    {
        return $this->created_at;
    }

    /**
     * Generated from protobuf field <code>int32 created_at = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setCreatedAt($var)
    {
        GPBUtil::checkInt32($var);
        $this->created_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 change = 3;</code>
     * @return int
     */
    public function getChange()
    {
        return $this->change;
    }

    /**
     * Generated from protobuf field <code>int32 change = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setChange($var)
    {
        GPBUtil::checkInt32($var);
        $this->change = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 quantity = 4;</code>
     * @return int
     */
    public function getQuantity()
    {
        return $this->quantity;
    }

    /**
     * Generated from protobuf field <code>int32 quantity = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setQuantity($var)
    {
        GPBUtil::checkInt32($var);
        $this->quantity = $var;

        return $this;
    }

}
