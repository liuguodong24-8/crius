<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/product_category.proto

namespace Omy\Crius\ExtensionServer\ListProductCategoryResponse;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.ListProductCategoryResponse.Data</code>
 */
class Data extends \Google\Protobuf\Internal\Message
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
     * Generated from protobuf field <code>int32 level = 3;</code>
     */
    protected $level = 0;
    /**
     * Generated from protobuf field <code>string status = 4;</code>
     */
    protected $status = '';
    /**
     * Generated from protobuf field <code>string code = 5;</code>
     */
    protected $code = '';
    /**
     * Generated from protobuf field <code>int32 weight = 6;</code>
     */
    protected $weight = 0;
    /**
     * Generated from protobuf field <code>int64 created_at = 7;</code>
     */
    protected $created_at = 0;
    /**
     * Generated from protobuf field <code>string parent_id = 8;</code>
     */
    protected $parent_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $name
     *     @type int $level
     *     @type string $status
     *     @type string $code
     *     @type int $weight
     *     @type int|string $created_at
     *     @type string $parent_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\ProductCategory::initOnce();
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
     * Generated from protobuf field <code>int32 level = 3;</code>
     * @return int
     */
    public function getLevel()
    {
        return $this->level;
    }

    /**
     * Generated from protobuf field <code>int32 level = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setLevel($var)
    {
        GPBUtil::checkInt32($var);
        $this->level = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string status = 4;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>string status = 4;</code>
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
     * Generated from protobuf field <code>string code = 5;</code>
     * @return string
     */
    public function getCode()
    {
        return $this->code;
    }

    /**
     * Generated from protobuf field <code>string code = 5;</code>
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
     * Generated from protobuf field <code>int32 weight = 6;</code>
     * @return int
     */
    public function getWeight()
    {
        return $this->weight;
    }

    /**
     * Generated from protobuf field <code>int32 weight = 6;</code>
     * @param int $var
     * @return $this
     */
    public function setWeight($var)
    {
        GPBUtil::checkInt32($var);
        $this->weight = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 created_at = 7;</code>
     * @return int|string
     */
    public function getCreatedAt()
    {
        return $this->created_at;
    }

    /**
     * Generated from protobuf field <code>int64 created_at = 7;</code>
     * @param int|string $var
     * @return $this
     */
    public function setCreatedAt($var)
    {
        GPBUtil::checkInt64($var);
        $this->created_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string parent_id = 8;</code>
     * @return string
     */
    public function getParentId()
    {
        return $this->parent_id;
    }

    /**
     * Generated from protobuf field <code>string parent_id = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setParentId($var)
    {
        GPBUtil::checkString($var, True);
        $this->parent_id = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Data::class, \Omy\Crius\ExtensionServer\ListProductCategoryResponse_Data::class);

