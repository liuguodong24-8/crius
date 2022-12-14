<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/benefit.proto

namespace GPBMetadata\Proto\MemberExtension;

class Benefit
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
?
$proto/member-extension/benefit.protomemberExtension"A
CreateBenefitRequest)
benefit (2.memberExtension.Benefit"?
Benefit

id (	
name (	
icon (	
category (	
describe (	
status (	/
coupons (2.memberExtension.BenefitCoupon<
point_coefficient (2!.memberExtension.PointCoefficient
shop_benefit	 (	:
custom_wallpaper
 (2 .memberExtension.CustomWallpaper
custom_nameplate (	8
custom_discount (2.memberExtension.CustomDiscount
merchant_id (	"`
BenefitCoupon
code (	
	valid_day (
receive_day (
receive_month_day ("9
PointCoefficient
coefficient (
duration (",
CustomWallpaper
num (
gift (".
CustomDiscount

id (	
discount ("8
UpdateBenefitStatusRequest

id (	
status (	"^
GetBenefitsRequest
name (	
status (	
offset (
limit (
ids (	"p
GetBenefitsResponse

error_code (
error_message (	.
data (2 .memberExtension.GetBenefitsData"L
GetBenefitsData*
benefits (2.memberExtension.Benefit
total ("A
UpdateBenefitRequest)
benefit (2.memberExtension.Benefit" 
ShowBenefitRequest

id (	"h
ShowBenefitResponse

error_code (
error_message (	&
data (2.memberExtension.Benefit"-
GetMemberBenefitsRequest
	member_id (	"n
GetMemberBenefitsResponse

error_code (
error_message (	&
data (2.memberExtension.Benefit"Q
GetMemberLevelBenefitsRequest
	member_id (	
offset (
limit ("?
GetMemberLevelBenefitsResponse

error_code (
error_message (	9
data (2+.memberExtension.GetMemberLevelBenefitsData"f
GetMemberLevelBenefitsData*
benefits (2.memberExtension.Benefit
total (
level (B+Z./proto;proto?Omy\\Crius\\ExtensionServerbproto3'
        , true);

        static::$is_initialized = true;
    }
}

