<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/product_order.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.ProductOrderWithProductInfo</code>
 */
class ProductOrderWithProductInfo extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.memberExtension.ProductOrder order = 1;</code>
     */
    protected $order = null;
    /**
     * Generated from protobuf field <code>string product_name = 2;</code>
     */
    protected $product_name = '';
    /**
     * Generated from protobuf field <code>string product_code = 3;</code>
     */
    protected $product_code = '';
    /**
     * Generated from protobuf field <code>int32 point = 4;</code>
     */
    protected $point = 0;
    /**
     * Generated from protobuf field <code>int32 price = 5;</code>
     */
    protected $price = 0;
    /**
     * Generated from protobuf field <code>int32 line_price = 6;</code>
     */
    protected $line_price = 0;
    /**
     * Generated from protobuf field <code>repeated string images = 7;</code>
     */
    private $images;
    /**
     * Generated from protobuf field <code>repeated string videos = 8;</code>
     */
    private $videos;
    /**
     * Generated from protobuf field <code>repeated string graphic_detail = 9;</code>
     */
    private $graphic_detail;
    /**
     * Generated from protobuf field <code>string describe = 10;</code>
     */
    protected $describe = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\ExtensionServer\ProductOrder $order
     *     @type string $product_name
     *     @type string $product_code
     *     @type int $point
     *     @type int $price
     *     @type int $line_price
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $images
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $videos
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $graphic_detail
     *     @type string $describe
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\ProductOrder::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.memberExtension.ProductOrder order = 1;</code>
     * @return \Omy\Crius\ExtensionServer\ProductOrder|null
     */
    public function getOrder()
    {
        return $this->order;
    }

    public function hasOrder()
    {
        return isset($this->order);
    }

    public function clearOrder()
    {
        unset($this->order);
    }

    /**
     * Generated from protobuf field <code>.memberExtension.ProductOrder order = 1;</code>
     * @param \Omy\Crius\ExtensionServer\ProductOrder $var
     * @return $this
     */
    public function setOrder($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\ExtensionServer\ProductOrder::class);
        $this->order = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string product_name = 2;</code>
     * @return string
     */
    public function getProductName()
    {
        return $this->product_name;
    }

    /**
     * Generated from protobuf field <code>string product_name = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setProductName($var)
    {
        GPBUtil::checkString($var, True);
        $this->product_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string product_code = 3;</code>
     * @return string
     */
    public function getProductCode()
    {
        return $this->product_code;
    }

    /**
     * Generated from protobuf field <code>string product_code = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setProductCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->product_code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 point = 4;</code>
     * @return int
     */
    public function getPoint()
    {
        return $this->point;
    }

    /**
     * Generated from protobuf field <code>int32 point = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setPoint($var)
    {
        GPBUtil::checkInt32($var);
        $this->point = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 price = 5;</code>
     * @return int
     */
    public function getPrice()
    {
        return $this->price;
    }

    /**
     * Generated from protobuf field <code>int32 price = 5;</code>
     * @param int $var
     * @return $this
     */
    public function setPrice($var)
    {
        GPBUtil::checkInt32($var);
        $this->price = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 line_price = 6;</code>
     * @return int
     */
    public function getLinePrice()
    {
        return $this->line_price;
    }

    /**
     * Generated from protobuf field <code>int32 line_price = 6;</code>
     * @param int $var
     * @return $this
     */
    public function setLinePrice($var)
    {
        GPBUtil::checkInt32($var);
        $this->line_price = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string images = 7;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getImages()
    {
        return $this->images;
    }

    /**
     * Generated from protobuf field <code>repeated string images = 7;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setImages($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->images = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string videos = 8;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getVideos()
    {
        return $this->videos;
    }

    /**
     * Generated from protobuf field <code>repeated string videos = 8;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setVideos($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->videos = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string graphic_detail = 9;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getGraphicDetail()
    {
        return $this->graphic_detail;
    }

    /**
     * Generated from protobuf field <code>repeated string graphic_detail = 9;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setGraphicDetail($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->graphic_detail = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string describe = 10;</code>
     * @return string
     */
    public function getDescribe()
    {
        return $this->describe;
    }

    /**
     * Generated from protobuf field <code>string describe = 10;</code>
     * @param string $var
     * @return $this
     */
    public function setDescribe($var)
    {
        GPBUtil::checkString($var, True);
        $this->describe = $var;

        return $this;
    }

}
