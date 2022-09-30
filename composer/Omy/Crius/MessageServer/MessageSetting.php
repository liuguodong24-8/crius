<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/message/setting.proto

namespace Omy\Crius\MessageServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>message.MessageSetting</code>
 */
class MessageSetting extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    protected $id = '';
    /**
     * Generated from protobuf field <code>string message_type = 2;</code>
     */
    protected $message_type = '';
    /**
     * Generated from protobuf field <code>string message_type_name = 3;</code>
     */
    protected $message_type_name = '';
    /**
     * Generated from protobuf field <code>string trigger_type = 4;</code>
     */
    protected $trigger_type = '';
    /**
     * Generated from protobuf field <code>double advance_hour = 5;</code>
     */
    protected $advance_hour = 0.0;
    /**
     * Generated from protobuf field <code>string sms_template_id = 6;</code>
     */
    protected $sms_template_id = '';
    /**
     * Generated from protobuf field <code>string sms_template_name = 7;</code>
     */
    protected $sms_template_name = '';
    /**
     * Generated from protobuf field <code>string wechat_template_id = 8;</code>
     */
    protected $wechat_template_id = '';
    /**
     * Generated from protobuf field <code>string wechat_template_name = 9;</code>
     */
    protected $wechat_template_name = '';
    /**
     * Generated from protobuf field <code>repeated .message.SpecialSetting special_setting = 10;</code>
     */
    private $special_setting;
    /**
     * Generated from protobuf field <code>repeated .message.Cc cc_list = 11;</code>
     */
    private $cc_list;
    /**
     * Generated from protobuf field <code>repeated string special_branches = 12;</code>
     */
    private $special_branches;
    /**
     * Generated from protobuf field <code>string status = 13;</code>
     */
    protected $status = '';
    /**
     * Generated from protobuf field <code>int64 created_at = 14;</code>
     */
    protected $created_at = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $message_type
     *     @type string $message_type_name
     *     @type string $trigger_type
     *     @type float $advance_hour
     *     @type string $sms_template_id
     *     @type string $sms_template_name
     *     @type string $wechat_template_id
     *     @type string $wechat_template_name
     *     @type \Omy\Crius\MessageServer\SpecialSetting[]|\Google\Protobuf\Internal\RepeatedField $special_setting
     *     @type \Omy\Crius\MessageServer\Cc[]|\Google\Protobuf\Internal\RepeatedField $cc_list
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $special_branches
     *     @type string $status
     *     @type int|string $created_at
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\Message\Setting::initOnce();
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
     * Generated from protobuf field <code>string message_type = 2;</code>
     * @return string
     */
    public function getMessageType()
    {
        return $this->message_type;
    }

    /**
     * Generated from protobuf field <code>string message_type = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setMessageType($var)
    {
        GPBUtil::checkString($var, True);
        $this->message_type = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string message_type_name = 3;</code>
     * @return string
     */
    public function getMessageTypeName()
    {
        return $this->message_type_name;
    }

    /**
     * Generated from protobuf field <code>string message_type_name = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setMessageTypeName($var)
    {
        GPBUtil::checkString($var, True);
        $this->message_type_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string trigger_type = 4;</code>
     * @return string
     */
    public function getTriggerType()
    {
        return $this->trigger_type;
    }

    /**
     * Generated from protobuf field <code>string trigger_type = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setTriggerType($var)
    {
        GPBUtil::checkString($var, True);
        $this->trigger_type = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>double advance_hour = 5;</code>
     * @return float
     */
    public function getAdvanceHour()
    {
        return $this->advance_hour;
    }

    /**
     * Generated from protobuf field <code>double advance_hour = 5;</code>
     * @param float $var
     * @return $this
     */
    public function setAdvanceHour($var)
    {
        GPBUtil::checkDouble($var);
        $this->advance_hour = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string sms_template_id = 6;</code>
     * @return string
     */
    public function getSmsTemplateId()
    {
        return $this->sms_template_id;
    }

    /**
     * Generated from protobuf field <code>string sms_template_id = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setSmsTemplateId($var)
    {
        GPBUtil::checkString($var, True);
        $this->sms_template_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string sms_template_name = 7;</code>
     * @return string
     */
    public function getSmsTemplateName()
    {
        return $this->sms_template_name;
    }

    /**
     * Generated from protobuf field <code>string sms_template_name = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setSmsTemplateName($var)
    {
        GPBUtil::checkString($var, True);
        $this->sms_template_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string wechat_template_id = 8;</code>
     * @return string
     */
    public function getWechatTemplateId()
    {
        return $this->wechat_template_id;
    }

    /**
     * Generated from protobuf field <code>string wechat_template_id = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setWechatTemplateId($var)
    {
        GPBUtil::checkString($var, True);
        $this->wechat_template_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string wechat_template_name = 9;</code>
     * @return string
     */
    public function getWechatTemplateName()
    {
        return $this->wechat_template_name;
    }

    /**
     * Generated from protobuf field <code>string wechat_template_name = 9;</code>
     * @param string $var
     * @return $this
     */
    public function setWechatTemplateName($var)
    {
        GPBUtil::checkString($var, True);
        $this->wechat_template_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .message.SpecialSetting special_setting = 10;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getSpecialSetting()
    {
        return $this->special_setting;
    }

    /**
     * Generated from protobuf field <code>repeated .message.SpecialSetting special_setting = 10;</code>
     * @param \Omy\Crius\MessageServer\SpecialSetting[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setSpecialSetting($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MessageServer\SpecialSetting::class);
        $this->special_setting = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .message.Cc cc_list = 11;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getCcList()
    {
        return $this->cc_list;
    }

    /**
     * Generated from protobuf field <code>repeated .message.Cc cc_list = 11;</code>
     * @param \Omy\Crius\MessageServer\Cc[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setCcList($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MessageServer\Cc::class);
        $this->cc_list = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string special_branches = 12;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getSpecialBranches()
    {
        return $this->special_branches;
    }

    /**
     * Generated from protobuf field <code>repeated string special_branches = 12;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setSpecialBranches($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->special_branches = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string status = 13;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * Generated from protobuf field <code>string status = 13;</code>
     * @param string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkString($var, True);
        $this->status = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 created_at = 14;</code>
     * @return int|string
     */
    public function getCreatedAt()
    {
        return $this->created_at;
    }

    /**
     * Generated from protobuf field <code>int64 created_at = 14;</code>
     * @param int|string $var
     * @return $this
     */
    public function setCreatedAt($var)
    {
        GPBUtil::checkInt64($var);
        $this->created_at = $var;

        return $this;
    }

}
