<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/point_config.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.SavePointConfigRequest</code>
 */
class SavePointConfigRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .memberExtension.PointBeginEnd begin_end = 1;</code>
     */
    private $begin_end;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\ExtensionServer\PointBeginEnd[]|\Google\Protobuf\Internal\RepeatedField $begin_end
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\PointConfig::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.PointBeginEnd begin_end = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getBeginEnd()
    {
        return $this->begin_end;
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.PointBeginEnd begin_end = 1;</code>
     * @param \Omy\Crius\ExtensionServer\PointBeginEnd[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setBeginEnd($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\ExtensionServer\PointBeginEnd::class);
        $this->begin_end = $arr;

        return $this;
    }

}

