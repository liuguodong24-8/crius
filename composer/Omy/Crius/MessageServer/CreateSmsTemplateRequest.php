<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/message/sms.proto

namespace Omy\Crius\MessageServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>message.CreateSmsTemplateRequest</code>
 */
class CreateSmsTemplateRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string name = 1;</code>
     */
    protected $name = '';
    /**
     * Generated from protobuf field <code>string sign = 2;</code>
     */
    protected $sign = '';
    /**
     * Generated from protobuf field <code>string category = 3;</code>
     */
    protected $category = '';
    /**
     * Generated from protobuf field <code>string category_key = 4;</code>
     */
    protected $category_key = '';
    /**
     * Generated from protobuf field <code>string content = 5;</code>
     */
    protected $content = '';
    /**
     * Generated from protobuf field <code>string status = 6;</code>
     */
    protected $status = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *     @type string $sign
     *     @type string $category
     *     @type string $category_key
     *     @type string $content
     *     @type string $status
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Message\Sms::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string name = 1;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * Generated from protobuf field <code>string name = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string sign = 2;</code>
     * @return string
     */
    public function getSign()
    {
        return $this->sign;
    }

    /**
     * Generated from protobuf field <code>string sign = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setSign($var)
    {
        GPBUtil::checkString($var, True);
        $this->sign = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string category = 3;</code>
     * @return string
     */
    public function getCategory()
    {
        return $this->category;
    }

    /**
     * Generated from protobuf field <code>string category = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setCategory($var)
    {
        GPBUtil::checkString($var, True);
        $this->category = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string category_key = 4;</code>
     * @return string
     */
    public function getCategoryKey()
    {
        return $this->category_key;
    }

    /**
     * Generated from protobuf field <code>string category_key = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setCategoryKey($var)
    {
        GPBUtil::checkString($var, True);
        $this->category_key = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string content = 5;</code>
     * @return string
     */
    public function getContent()
    {
        return $this->content;
    }

    /**
     * Generated from protobuf field <code>string content = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setContent($var)
    {
        GPBUtil::checkString($var, True);
        $this->content = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string status = 6;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>string status = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkString($var, True);
        $this->status = $var;

        return $this;
    }

}

