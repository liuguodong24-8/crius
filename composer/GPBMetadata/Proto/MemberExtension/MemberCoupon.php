<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/member_coupon.proto

namespace GPBMetadata\Proto\MemberExtension;

class MemberCoupon
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Proto\MemberExtension\Coupon::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
*proto/member-extension/member_coupon.protomemberExtension"�
MemberCoupon

id (	
	coupon_id (	
coupon_code (	
	member_id (	
status (	
	branch_id (	
used_at (

life_begin (
life_end	 (=
distribute_coupon
 (2".memberExtension.CollectableCoupon
style (	

order_code (	

created_at
GetMemberCouponsRequest
member_phone (	
status (	
offset (
limit (
member_name (	
	member_id (	
	coupon_id (	
coupon_code (	

time_begin	 (
time_end
 ("N

coupons (2.memberExtension.MemberCoupon
total ("s
GetMemberCouponsResponse

error_code (

data (2.memberExtension.MemberCoupons"%
ShowMemberCouponRequest

id (	"r
ShowMemberCouponResponse

error_code (

data (2.memberExtension.MemberCoupon"Q
CreateMemberCouponRequest4

CreateMemberCouponsRequest5
member_coupons (2.memberExtension.MemberCoupon"l
GetCollectableCouponsRequest

distribute_way (	
offset (
limit ("}
GetCollectableCouponsResponse

error_code (

data (2#.memberExtension.CollectableCoupons"�
CollectableCoupon

id (	
distribution_way (	
distribution_num (@

valid_time (2,.memberExtension.CollectableCoupon.ValidTimeQ
valid_time_interval (24.memberExtension.CollectableCoupon.ValidTimeInterval
branches (	
collection_num (U
invalid_time_interval (26.memberExtension.CollectableCoupon.InvalidTimeInterval\'
coupon	 (2.memberExtension.CouponM
	ValidTime
begin (
end (
	after_day (
	valid_day (A
ValidTimeInterval
week_day (
begin (	
end (	1
InvalidTimeInterval
begin (
end ("X
CollectableCoupons
total (3
coupons (2".memberExtension.CollectableCouponB+Z
        , true);

        static::$is_initialized = true;
    }
}
