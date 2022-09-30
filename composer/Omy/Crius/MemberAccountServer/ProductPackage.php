<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-account/card.proto

namespace Omy\Crius\MemberAccountServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberAccount.ProductPackage</code>
 */
class ProductPackage extends \Google\Protobuf\Internal\Message
{
    /**
     * 商品/套餐id
     *
     * Generated from protobuf field <code>string product_package_id = 1;</code>
     */
    protected $product_package_id = '';
    /**
     * 商品/套餐code
     *
     * Generated from protobuf field <code>string code = 2;</code>
     */
    protected $code = '';
    /**
     * 消费/获得数量
     *
     * Generated from protobuf field <code>int32 number = 3;</code>
     */
    protected $number = 0;
    /**
     * 商品/套餐价格
     *
     * Generated from protobuf field <code>int32 price = 4;</code>
     */
    protected $price = 0;
    /**
     * 商品/套餐名字
     *
     * Generated from protobuf field <code>string title = 5;</code>
     */
    protected $title = '';
    /**
     * 获得商品/套餐剩余数量
     *
     * Generated from protobuf field <code>int32 left = 6;</code>
     */
    protected $left = 0;
    /**
     * 类别 商品 product, 套餐 package
     *
     * Generated from protobuf field <code>string category = 7;</code>
     */
    protected $category = '';
    /**
     * Generated from protobuf field <code>string id = 8;</code>
     */
    protected $id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $product_package_id
     *           商品/套餐id
     *     @type string $code
     *           商品/套餐code
     *     @type int $number
     *           消费/获得数量
     *     @type int $price
     *           商品/套餐价格
     *     @type string $title
     *           商品/套餐名字
     *     @type int $left
     *           获得商品/套餐剩余数量
     *     @type string $category
     *           类别 商品 product, 套餐 package
     *     @type string $id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberAccount\Card::initOnce();
        parent::__construct($data);
    }

    /**
     * 商品/套餐id
     *
     * Generated from protobuf field <code>string product_package_id = 1;</code>
     * @return string
     */
    public function getProductPackageId()
    {
        return $this->product_package_id;
    }

    /**
     * 商品/套餐id
     *
     * Generated from protobuf field <code>string product_package_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setProductPackageId($var)
    {
        GPBUtil::checkString($var, True);
        $this->product_package_id = $var;

        return $this;
    }

    /**
     * 商品/套餐code
     *
     * Generated from protobuf field <code>string code = 2;</code>
     * @return string
     */
    public function getCode()
    {
        return $this->code;
    }

    /**
     * 商品/套餐code
     *
     * Generated from protobuf field <code>string code = 2;</code>
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
     * 消费/获得数量
     *
     * Generated from protobuf field <code>int32 number = 3;</code>
     * @return int
     */
    public function getNumber()
    {
        return $this->number;
    }

    /**
     * 消费/获得数量
     *
     * Generated from protobuf field <code>int32 number = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setNumber($var)
    {
        GPBUtil::checkInt32($var);
        $this->number = $var;

        return $this;
    }

    /**
     * 商品/套餐价格
     *
     * Generated from protobuf field <code>int32 price = 4;</code>
     * @return int
     */
    public function getPrice()
    {
        return $this->price;
    }

    /**
     * 商品/套餐价格
     *
     * Generated from protobuf field <code>int32 price = 4;</code>
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
     * 商品/套餐名字
     *
     * Generated from protobuf field <code>string title = 5;</code>
     * @return string
     */
    public function getTitle()
    {
        return $this->title;
    }

    /**
     * 商品/套餐名字
     *
     * Generated from protobuf field <code>string title = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setTitle($var)
    {
        GPBUtil::checkString($var, True);
        $this->title = $var;

        return $this;
    }

    /**
     * 获得商品/套餐剩余数量
     *
     * Generated from protobuf field <code>int32 left = 6;</code>
     * @return int
     */
    public function getLeft()
    {
        return $this->left;
    }

    /**
     * 获得商品/套餐剩余数量
     *
     * Generated from protobuf field <code>int32 left = 6;</code>
     * @param int $var
     * @return $this
     */
    public function setLeft($var)
    {
        GPBUtil::checkInt32($var);
        $this->left = $var;

        return $this;
    }

    /**
     * 类别 商品 product, 套餐 package
     *
     * Generated from protobuf field <code>string category = 7;</code>
     * @return string
     */
    public function getCategory()
    {
        return $this->category;
    }

    /**
     * 类别 商品 product, 套餐 package
     *
     * Generated from protobuf field <code>string category = 7;</code>
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
     * Generated from protobuf field <code>string id = 8;</code>
     * @return string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Generated from protobuf field <code>string id = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkString($var, True);
        $this->id = $var;

        return $this;
    }

}
