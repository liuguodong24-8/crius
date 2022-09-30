<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/division_level.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.GetDivisionsAndLevelsData</code>
 */
class GetDivisionsAndLevelsData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .memberExtension.GetDivisionsAndLevelsData.DivisionWithLevel division_level = 1;</code>
     */
    private $division_level;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\ExtensionServer\GetDivisionsAndLevelsData\DivisionWithLevel[]|\Google\Protobuf\Internal\RepeatedField $division_level
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\DivisionLevel::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.GetDivisionsAndLevelsData.DivisionWithLevel division_level = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getDivisionLevel()
    {
        return $this->division_level;
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.GetDivisionsAndLevelsData.DivisionWithLevel division_level = 1;</code>
     * @param \Omy\Crius\ExtensionServer\GetDivisionsAndLevelsData\DivisionWithLevel[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setDivisionLevel($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\ExtensionServer\GetDivisionsAndLevelsData\DivisionWithLevel::class);
        $this->division_level = $arr;

        return $this;
    }

}

