<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/member.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.RefundPointRequest</code>
 */
class RefundPointRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string pos_bill_id = 1;</code>
     */
    protected $pos_bill_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $pos_bill_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\Member::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string pos_bill_id = 1;</code>
     * @return string
     */
    public function getPosBillId()
    {
        return $this->pos_bill_id;
    }

    /**
     * Generated from protobuf field <code>string pos_bill_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setPosBillId($var)
    {
        GPBUtil::checkString($var, True);
        $this->pos_bill_id = $var;

        return $this;
    }

}

