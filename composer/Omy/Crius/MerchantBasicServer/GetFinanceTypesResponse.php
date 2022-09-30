<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/finance_type.proto

namespace Omy\Crius\MerchantBasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.GetFinanceTypesResponse</code>
 */
class GetFinanceTypesResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 error_code = 1;</code>
     */
    protected $error_code = 0;
    /**
     * Generated from protobuf field <code>string error_message = 2;</code>
     */
    protected $error_message = '';
    /**
     * Generated from protobuf field <code>repeated .merchantBasic.FinanceType data = 3;</code>
     */
    private $data;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $error_code
     *     @type string $error_message
     *     @type \Omy\Crius\MerchantBasicServer\FinanceType[]|\Google\Protobuf\Internal\RepeatedField $data
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\FinanceType::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 error_code = 1;</code>
     * @return int
     */
    public function getErrorCode()
    {
        return $this->error_code;
    }

    /**
     * Generated from protobuf field <code>int32 error_code = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setErrorCode($var)
    {
        GPBUtil::checkInt32($var);
        $this->error_code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string error_message = 2;</code>
     * @return string
     */
    public function getErrorMessage()
    {
        return $this->error_message;
    }

    /**
     * Generated from protobuf field <code>string error_message = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setErrorMessage($var)
    {
        GPBUtil::checkString($var, True);
        $this->error_message = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .merchantBasic.FinanceType data = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getData()
    {
        return $this->data;
    }

    /**
     * Generated from protobuf field <code>repeated .merchantBasic.FinanceType data = 3;</code>
     * @param \Omy\Crius\MerchantBasicServer\FinanceType[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setData($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MerchantBasicServer\FinanceType::class);
        $this->data = $arr;

        return $this;
    }

}

