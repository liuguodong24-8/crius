<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/basic/basic.proto

namespace Omy\Crius\BasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>basic.ProvinceArea</code>
 */
class ProvinceArea extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.basic.Area city = 1;</code>
     */
    protected $city = null;
    /**
     * Generated from protobuf field <code>repeated .basic.Area district = 2;</code>
     */
    private $district;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\BasicServer\Area $city
     *     @type \Omy\Crius\BasicServer\Area[]|\Google\Protobuf\Internal\RepeatedField $district
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Basic\Basic::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.basic.Area city = 1;</code>
     * @return \Omy\Crius\BasicServer\Area|null
     */
    public function getCity()
    {
        return $this->city;
    }

    public function hasCity()
    {
        return isset($this->city);
    }

    public function clearCity()
    {
        unset($this->city);
    }

    /**
     * Generated from protobuf field <code>.basic.Area city = 1;</code>
     * @param \Omy\Crius\BasicServer\Area $var
     * @return $this
     */
    public function setCity($var)
    {
        GPBUtil::checkMessage($var, \Omy\Crius\BasicServer\Area::class);
        $this->city = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .basic.Area district = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getDistrict()
    {
        return $this->district;
    }

    /**
     * Generated from protobuf field <code>repeated .basic.Area district = 2;</code>
     * @param \Omy\Crius\BasicServer\Area[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setDistrict($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\BasicServer\Area::class);
        $this->district = $arr;

        return $this;
    }

}

