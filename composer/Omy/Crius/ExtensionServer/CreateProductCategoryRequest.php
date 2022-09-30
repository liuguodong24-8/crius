<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/product_category.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * CreateProductCategoryRequest 新增商品分类
 *
 * Generated from protobuf message <code>memberExtension.CreateProductCategoryRequest</code>
 */
class CreateProductCategoryRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string name = 1;</code>
     */
    protected $name = '';
    /**
     * Generated from protobuf field <code>int32 level = 2;</code>
     */
    protected $level = 0;
    /**
     * Generated from protobuf field <code>string status = 3;</code>
     */
    protected $status = '';
    /**
     * Generated from protobuf field <code>string code = 4;</code>
     */
    protected $code = '';
    /**
     * Generated from protobuf field <code>int32 weight = 5;</code>
     */
    protected $weight = 0;
    /**
     * Generated from protobuf field <code>string parent_id = 6;</code>
     */
    protected $parent_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *     @type int $level
     *     @type string $status
     *     @type string $code
     *     @type int $weight
     *     @type string $parent_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\ProductCategory::initOnce();
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
     * Generated from protobuf field <code>int32 level = 2;</code>
     * @return int
     */
    public function getLevel()
    {
        return $this->level;
    }

    /**
     * Generated from protobuf field <code>int32 level = 2;</code>
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
     * Generated from protobuf field <code>string status = 3;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>string status = 3;</code>
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
     * Generated from protobuf field <code>string code = 4;</code>
     * @return string
     */
    public function getCode()
    {
        return $this->code;
    }

    /**
     * Generated from protobuf field <code>string code = 4;</code>
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
     * Generated from protobuf field <code>int32 weight = 5;</code>
     * @return int
     */
    public function getWeight()
    {
        return $this->weight;
    }

    /**
     * Generated from protobuf field <code>int32 weight = 5;</code>
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
     * Generated from protobuf field <code>string parent_id = 6;</code>
     * @return string
     */
    public function getParentId()
    {
        return $this->parent_id;
    }

    /**
     * Generated from protobuf field <code>string parent_id = 6;</code>
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

