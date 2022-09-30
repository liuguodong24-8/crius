<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: crius.proto

namespace Omy\Crius;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>proto.DiscoverServersResponse</code>
 */
class DiscoverServersResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .proto.Service services = 1;</code>
     */
    private $services;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\Service[]|\Google\Protobuf\Internal\RepeatedField $services
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Crius::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .proto.Service services = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getServices()
    {
        return $this->services;
    }

    /**
     * Generated from protobuf field <code>repeated .proto.Service services = 1;</code>
     * @param \Omy\Crius\Service[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setServices($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\Service::class);
        $this->services = $arr;

        return $this;
    }

}

