<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/basic/basic.proto

namespace Omy\Crius\BasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>basic.GetAllAreasByLevelRequest</code>
 */
class GetAllAreasByLevelRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 level = 1;</code>
     */
    protected $level = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $level
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Basic\Basic::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 level = 1;</code>
     * @return int
     */
    public function getLevel()
    {
        return $this->level;
    }

    /**
     * Generated from protobuf field <code>int32 level = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setLevel($var)
    {
        GPBUtil::checkInt32($var);
        $this->level = $var;

        return $this;
    }

}
