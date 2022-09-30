<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/products.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.Carriage</code>
 */
class Carriage extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 cash = 1;</code>
     */
    protected $cash = 0;
    /**
     * Generated from protobuf field <code>int32 point = 2;</code>
     */
    protected $point = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $cash
     *     @type int $point
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\Products::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 cash = 1;</code>
     * @return int
     */
    public function getCash()
    {
        return $this->cash;
    }

    /**
     * Generated from protobuf field <code>int32 cash = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setCash($var)
    {
        GPBUtil::checkInt32($var);
        $this->cash = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 point = 2;</code>
     * @return int
     */
    public function getPoint()
    {
        return $this->point;
    }

    /**
     * Generated from protobuf field <code>int32 point = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setPoint($var)
    {
        GPBUtil::checkInt32($var);
        $this->point = $var;

        return $this;
    }

}
