<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/point.proto

namespace Omy\Crius\MerchantBasicServer\ListPointRuleResponse;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.ListPointRuleResponse.Data</code>
 */
class Data extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .merchantBasic.PointRule data = 1;</code>
     */
    private $data;
    /**
     * Generated from protobuf field <code>int64 total = 2;</code>
     */
    protected $total = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\MerchantBasicServer\PointRule[]|\Google\Protobuf\Internal\RepeatedField $data
     *     @type int|string $total
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\Point::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .merchantBasic.PointRule data = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getData()
    {
        return $this->data;
    }

    /**
     * Generated from protobuf field <code>repeated .merchantBasic.PointRule data = 1;</code>
     * @param \Omy\Crius\MerchantBasicServer\PointRule[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setData($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MerchantBasicServer\PointRule::class);
        $this->data = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 total = 2;</code>
     * @return int|string
     */
    public function getTotal()
    {
        return $this->total;
    }

    /**
     * Generated from protobuf field <code>int64 total = 2;</code>
     * @param int|string $var
     * @return $this
     */
    public function setTotal($var)
    {
        GPBUtil::checkInt64($var);
        $this->total = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(Data::class, \Omy\Crius\MerchantBasicServer\ListPointRuleResponse_Data::class);

