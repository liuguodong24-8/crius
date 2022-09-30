<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/member_task.proto

namespace Omy\Crius\ExtensionServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>memberExtension.GetRewardsRequest</code>
 */
class GetRewardsRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string member_id = 1;</code>
     */
    protected $member_id = '';
    /**
     * Generated from protobuf field <code>string member_task_id = 2;</code>
     */
    protected $member_task_id = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $member_id
     *     @type string $member_task_id
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MemberExtension\MemberTask::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string member_id = 1;</code>
     * @return string
     */
    public function getMemberId()
    {
        return $this->member_id;
    }

    /**
     * Generated from protobuf field <code>string member_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setMemberId($var)
    {
        GPBUtil::checkString($var, True);
        $this->member_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string member_task_id = 2;</code>
     * @return string
     */
    public function getMemberTaskId()
    {
        return $this->member_task_id;
    }

    /**
     * Generated from protobuf field <code>string member_task_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setMemberTaskId($var)
    {
        GPBUtil::checkString($var, True);
        $this->member_task_id = $var;

        return $this;
    }

}

