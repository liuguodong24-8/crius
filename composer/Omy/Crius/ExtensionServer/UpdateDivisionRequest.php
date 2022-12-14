<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/division_level.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.UpdateDivisionRequest</code>
 */
class UpdateDivisionRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.memberExtension.Division division = 1;</code>
     */
    protected $division = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\ExtensionServer\Division $division
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\DivisionLevel::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.memberExtension.Division division = 1;</code>
     * @return \Omy\Crius\ExtensionServer\Division|null
     */
    public function getDivision()
    {
        return $this->division;
    }

    public function hasDivision()
    {
        return isset($this->division);
    }

    public function clearDivision()
    {
        unset($this->division);
    }

    /**
     * Generated from protobuf field <code>.memberExtension.Division division = 1;</code>
     * @param \Omy\Crius\ExtensionServer\Division $var
     * @return $this
     */
    public function setDivision($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\ExtensionServer\Division::class);
        $this->division = $var;

        return $this;
    }

}

