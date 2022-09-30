<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/product_order.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.UpdateProductOrderStatusRequest</code>
 */
class UpdateProductOrderStatusRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>string status = 2;</code>
     */
    protected $status = '';
    /**
     * Generated from protobuf field <code>string trade_id = 3;</code>
     */
    protected $trade_id = '';
    /**
     * Generated from protobuf field <code>string express_company = 4;</code>
     */
    protected $express_company = '';
    /**
     * Generated from protobuf field <code>string express_code = 5;</code>
     */
    protected $express_code = '';
    /**
     * Generated from protobuf field <code>string express_company_cn = 6;</code>
     */
    protected $express_company_cn = '';
    /**
     * Generated from protobuf field <code>string transaction_id = 7;</code>
     */
    protected $transaction_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $status
     *     @type string $trade_id
     *     @type string $express_company
     *     @type string $express_code
     *     @type string $express_company_cn
     *     @type string $transaction_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\ProductOrder::initOnce();
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
     * Generated from protobuf field <code>string status = 2;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>string status = 2;</code>
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
     * Generated from protobuf field <code>string trade_id = 3;</code>
     * @return string
     */
    public function getTradeId()
    {
        return $this->trade_id;
    }

    /**
     * Generated from protobuf field <code>string trade_id = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setTradeId($var)
    {
        GPBUtil::checkString($var, True);
        $this->trade_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string express_company = 4;</code>
     * @return string
     */
    public function getExpressCompany()
    {
        return $this->express_company;
    }

    /**
     * Generated from protobuf field <code>string express_company = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setExpressCompany($var)
    {
        GPBUtil::checkString($var, True);
        $this->express_company = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string express_code = 5;</code>
     * @return string
     */
    public function getExpressCode()
    {
        return $this->express_code;
    }

    /**
     * Generated from protobuf field <code>string express_code = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setExpressCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->express_code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string express_company_cn = 6;</code>
     * @return string
     */
    public function getExpressCompanyCn()
    {
        return $this->express_company_cn;
    }

    /**
     * Generated from protobuf field <code>string express_company_cn = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setExpressCompanyCn($var)
    {
        GPBUtil::checkString($var, True);
        $this->express_company_cn = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string transaction_id = 7;</code>
     * @return string
     */
    public function getTransactionId()
    {
        return $this->transaction_id;
    }

    /**
     * Generated from protobuf field <code>string transaction_id = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setTransactionId($var)
    {
        GPBUtil::checkString($var, True);
        $this->transaction_id = $var;

        return $this;
    }

}
