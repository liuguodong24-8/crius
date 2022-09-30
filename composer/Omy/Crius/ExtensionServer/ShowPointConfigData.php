<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/point_config.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.ShowPointConfigData</code>
 */
class ShowPointConfigData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>string merchant_id = 2;</code>
     */
    protected $merchant_id = '';
    /**
     * Generated from protobuf field <code>repeated .memberExtension.PointBeginEnd begin_end = 3;</code>
     */
    private $begin_end;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $merchant_id
     *     @type \Omy\Crius\ExtensionServer\PointBeginEnd[]|\Google\Protobuf\Internal\RepeatedField $begin_end
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\PointConfig::initOnce();
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
     * Generated from protobuf field <code>string merchant_id = 2;</code>
     * @return string
     */
    public function getMerchantId()
    {
        return $this->merchant_id;
    }

    /**
     * Generated from protobuf field <code>string merchant_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setMerchantId($var)
    {
        GPBUtil::checkString($var, True);
        $this->merchant_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.PointBeginEnd begin_end = 3;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getBeginEnd()
    {
        return $this->begin_end;
    }

    /**
     * Generated from protobuf field <code>repeated .memberExtension.PointBeginEnd begin_end = 3;</code>
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

