<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/branch_tag.proto

namespace Omy\Crius\MerchantBasicServer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>merchantBasic.GetBranchTagsData</code>
 */
class GetBranchTagsData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .merchantBasic.BranchTagData branch_tags = 1;</code>
     */
    private $branch_tags;
    /**
     * Generated from protobuf field <code>int32 total = 2;</code>
     */
    protected $total = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Omy\Crius\MerchantBasicServer\BranchTagData[]|\Google\Protobuf\Internal\RepeatedField $branch_tags
     *     @type int $total
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Proto\MerchantBasic\BranchTag::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .merchantBasic.BranchTagData branch_tags = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getBranchTags()
    {
        return $this->branch_tags;
    }

    /**
     * Generated from protobuf field <code>repeated .merchantBasic.BranchTagData branch_tags = 1;</code>
     * @param \Omy\Crius\MerchantBasicServer\BranchTagData[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setBranchTags($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Omy\Crius\MerchantBasicServer\BranchTagData::class);
        $this->branch_tags = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 total = 2;</code>
     * @return int
     */
    public function getTotal()
    {
        return $this->total;
    }

    /**
     * Generated from protobuf field <code>int32 total = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setTotal($var)
    {
        GPBUtil::checkInt32($var);
        $this->total = $var;

        return $this;
    }

}
