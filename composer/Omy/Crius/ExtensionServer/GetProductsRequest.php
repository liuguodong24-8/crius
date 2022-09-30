<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/products.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.GetProductsRequest</code>
 */
class GetProductsRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string name = 1;</code>
     */
    protected $name = '';
    /**
     * Generated from protobuf field <code>string category = 2;</code>
     */
    protected $category = '';
    /**
     * Generated from protobuf field <code>string status = 3;</code>
     */
    protected $status = '';
    /**
     * Generated from protobuf field <code>int32 point_begin = 4;</code>
     */
    protected $point_begin = 0;
    /**
     * Generated from protobuf field <code>int32 point_end = 5;</code>
     */
    protected $point_end = 0;
    /**
     * Generated from protobuf field <code>string category_id = 6;</code>
     */
    protected $category_id = '';
    /**
     * Generated from protobuf field <code>int32 offset = 7;</code>
     */
    protected $offset = 0;
    /**
     * Generated from protobuf field <code>int32 limit = 8;</code>
     */
    protected $limit = 0;
    /**
     * Generated from protobuf field <code>string order_field = 9;</code>
     */
    protected $order_field = '';
    /**
     * Generated from protobuf field <code>bool order_asc = 10;</code>
     */
    protected $order_asc = false;
    /**
     * Generated from protobuf field <code>bool on_sale = 11;</code>
     */
    protected $on_sale = false;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *     @type string $category
     *     @type string $status
     *     @type int $point_begin
     *     @type int $point_end
     *     @type string $category_id
     *     @type int $offset
     *     @type int $limit
     *     @type string $order_field
     *     @type bool $order_asc
     *     @type bool $on_sale
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\Products::initOnce();
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
     * Generated from protobuf field <code>int32 point_begin = 4;</code>
     * @return int
     */
    public function getPointBegin()
    {
        return $this->point_begin;
    }

    /**
     * Generated from protobuf field <code>int32 point_begin = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setPointBegin($var)
    {
        GPBUtil::checkInt32($var);
        $this->point_begin = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 point_end = 5;</code>
     * @return int
     */
    public function getPointEnd()
    {
        return $this->point_end;
    }

    /**
     * Generated from protobuf field <code>int32 point_end = 5;</code>
     * @param int $var
     * @return $this
     */
    public function setPointEnd($var)
    {
        GPBUtil::checkInt32($var);
        $this->point_end = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string category_id = 6;</code>
     * @return string
     */
    public function getCategoryId()
    {
        return $this->category_id;
    }

    /**
     * Generated from protobuf field <code>string category_id = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setCategoryId($var)
    {
        GPBUtil::checkString($var, True);
        $this->category_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 offset = 7;</code>
     * @return int
     */
    public function getOffset()
    {
        return $this->offset;
    }

    /**
     * Generated from protobuf field <code>int32 offset = 7;</code>
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
     * Generated from protobuf field <code>string order_field = 9;</code>
     * @return string
     */
    public function getOrderField()
    {
        return $this->order_field;
    }

    /**
     * Generated from protobuf field <code>string order_field = 9;</code>
     * @param string $var
     * @return $this
     */
    public function setOrderField($var)
    {
        GPBUtil::checkString($var, True);
        $this->order_field = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool order_asc = 10;</code>
     * @return bool
     */
    public function getOrderAsc()
    {
        return $this->order_asc;
    }

    /**
     * Generated from protobuf field <code>bool order_asc = 10;</code>
     * @param bool $var
     * @return $this
     */
    public function setOrderAsc($var)
    {
        GPBUtil::checkBool($var);
        $this->order_asc = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool on_sale = 11;</code>
     * @return bool
     */
    public function getOnSale()
    {
        return $this->on_sale;
    }

    /**
     * Generated from protobuf field <code>bool on_sale = 11;</code>
     * @param bool $var
     * @return $this
     */
    public function setOnSale($var)
    {
        GPBUtil::checkBool($var);
        $this->on_sale = $var;

        return $this;
    }

}
