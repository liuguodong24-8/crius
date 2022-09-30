<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/payment.proto

namespace GPBMetadata\Proto\MerchantBasic;

class Payment
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
"proto/merchant-basic/payment.protomerchantBasic"L
BranchSubMchID
	branch_id (	
branch_name (	

sub_mch_id (	"�
GetBranchWithSubMchIDRequest
city_id (	
branch_name (	
offset (
limit (
province_id (	
district_id (	"T
BranchSubMchIDData/
branches (2.merchantBasic.BranchSubMchID
total ("{
GetBranchWithSubMchIDResponse

error_code (
error_message (	/
data (2!.merchantBasic.BranchSubMchIDData"A
SetBranchSubMchIDRequest
	branch_id (	

sub_mch_id (	"
GetWechatPaySettingRequest"3
!GetWechatPaySettingByAppIDRequest
app_id (	"�
WechatPaySetting
merchant_id (	
app_id (	
mch_id (	
private_key (	
cert_filename (	
cert_content (
headquarters_sub_mch_id (	"w
GetWechatPaySettingResponse

error_code (
error_message (	-
data (2.merchantBasic.WechatPaySetting"�
SetWechatPaySettingRequest
app_id (	
mch_id (	
private_key (	
cert_filename (	
cert_content (
headquarters_sub_mch_id (	"�
BranchWechatPaymentSetting
	branch_id (	
merchant_id (	
app_id (	
mch_id (	

sub_mch_id (	
private_key (	
cert_filename (	
cert_content (
headquarters_sub_mch_id	 (	"9
$GetBranchWechatPaymentSettingRequest
	branch_id (	"�
%GetBranchWechatPaymentSettingResponse

error_code (
error_message (	7
data (2).merchantBasic.BranchWechatPaymentSettingB/Z./proto;proto�Omy\\Crius\\MerchantBasicServerbproto3'
        , true);

        static::$is_initialized = true;
    }
}

