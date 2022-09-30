<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-account/member_account.proto

namespace Omy\Crius\MemberAccountServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberAccount.Account</code>
 */
class Account extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>string member_id = 2;</code>
     */
    protected $member_id = '';
    /**
     * Generated from protobuf field <code>string branch_id = 3;</code>
     */
    protected $branch_id = '';
    /**
     * Generated from protobuf field <code>int32 base_value = 4;</code>
     */
    protected $base_value = 0;
    /**
     * Generated from protobuf field <code>int32 gift_value = 5;</code>
     */
    protected $gift_value = 0;
    /**
     * Generated from protobuf field <code>repeated .memberAccount.ProductPackage products = 6;</code>
     */
    private $products;
    /**
     * Generated from protobuf field <code>repeated .memberAccount.ProductPackage packages = 7;</code>
     */
    private $packages;
    /**
     * Generated from protobuf field <code>string status = 8;</code>
     */
    protected $status = '';
    /**
     * Generated from protobuf field <code>string tag_id = 9;</code>
     */
    protected $tag_id = '';
    /**
     * Generated from protobuf field <code>int32 create_at = 10;</code>
     */
    protected $create_at = 0;
    /**
     * Generated from protobuf field <code>string category = 11;</code>
     */
    protected $category = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $member_id
     *     @type string $branch_id
     *     @type int $base_value
     *     @type int $gift_value
     *     @type \Omy\Crius\MemberAccountServer\ProductPackage[]|\Google\Protobuf\Internal\RepeatedField $products
     *     @type \Omy\Crius\MemberAccountServer\ProductPackage[]|\Google\Protobuf\Internal\RepeatedField $packages
     *     @type string $status
     *     @type string $tag_id
     *     @type int $create_at
     *     @type string $category
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberAccount\MemberAccount::initOnce();
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
     * Generated from protobuf field <code>string member_id = 2;</code>
     * @return string
     */
    public function getMemberId()
    {
        return $this->member_id;
    }

    /**
     * Generated from protobuf field <code>string member_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setMemberId($var)
    {
        GPBUtil::checkString($var, True);
        $this->member_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string branch_id = 3;</code>
     * @return string
     */
    public function getBranchId()
    {
        return $this->branch_id;
    }

    /**
     * Generated from protobuf field <code>string branch_id = 3;</code>
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
     * Generated from protobuf field <code>int32 base_value = 4;</code>
     * @return int
     */
    public function getBaseValue()
    {
        return $this->base_value;
    }

    /**
     * Generated from protobuf field <code>int32 base_value = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setBaseValue($var)
    {
        GPBUtil::checkInt32($var);
        $this->base_value = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 gift_value = 5;</code>
     * @return int
     */
    public function getGiftValue()
    {
        return $this->gift_value;
    }

    /**
     * Generated from protobuf field <code>int32 gift_value = 5;</code>
     * @param int $var
     * @return $this
     */
    public function setGiftValue($var)
    {
        GPBUtil::checkInt32($var);
        $this->gift_value = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .memberAccount.ProductPackage products = 6;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getProducts()
    {
        return $this->products;
    }

    /**
     * Generated from protobuf field <code>repeated .memberAccount.ProductPackage products = 6;</code>
     * @param \Omy\Crius\MemberAccountServer\ProductPackage[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setProducts($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MemberAccountServer\ProductPackage::class);
        $this->products = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .memberAccount.ProductPackage packages = 7;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getPackages()
    {
        return $this->packages;
    }

    /**
     * Generated from protobuf field <code>repeated .memberAccount.ProductPackage packages = 7;</code>
     * @param \Omy\Crius\MemberAccountServer\ProductPackage[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setPackages($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MemberAccountServer\ProductPackage::class);
        $this->packages = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string status = 8;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>string status = 8;</code>
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
     * Generated from protobuf field <code>string tag_id = 9;</code>
     * @return string
     */
    public function getTagId()
    {
        return $this->tag_id;
    }

    /**
     * Generated from protobuf field <code>string tag_id = 9;</code>
     * @param string $var
     * @return $this
     */
    public function setTagId($var)
    {
        GPBUtil::checkString($var, True);
        $this->tag_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 create_at = 10;</code>
     * @return int
     */
    public function getCreateAt()
    {
        return $this->create_at;
    }

    /**
     * Generated from protobuf field <code>int32 create_at = 10;</code>
     * @param int $var
     * @return $this
     */
    public function setCreateAt($var)
    {
        GPBUtil::checkInt32($var);
        $this->create_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string category = 11;</code>
     * @return string
     */
    public function getCategory()
    {
        return $this->category;
    }

    /**
     * Generated from protobuf field <code>string category = 11;</code>
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

