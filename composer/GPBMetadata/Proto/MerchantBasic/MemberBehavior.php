<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/member_behavior.proto

namespace GPBMetadata\Proto\MerchantBasic;

class MemberBehavior
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
?
*proto/merchant-basic/member_behavior.protomerchantBasic"z
MemberBehavior

id (	
	member_id (	
behavior (	
staff_id (	
	branch_id (	

created_at ("X
MemberBehaviors6
memberBehaviors (2.merchantBasic.MemberBehavior
total ("U
CreateMemberBehaviorRequest
behavior (	
	member_id (	
	branch_id (	"M
GetMemberBehaviorsRequest
	member_id (	
offset (
limit ("u
GetMemberBehaviorsResponse

error_code (
error_message (	,
data (2.merchantBasic.MemberBehaviorsB/Z./proto;proto?Omy\\Crius\\MerchantBasicServerbproto3'
        , true);

        static::$is_initialized = true;
    }
}

