<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/branch_tag.proto

namespace GPBMetadata\Proto\MerchantBasic;

class BranchTag
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

%proto/merchant-basic/branch_tag.proto
CreateBranchTagRequest
name (	

branch_ids (	"D
CreateBranchTagResponse

error_code (

GetBranchTagsRequest
name (	

date_start (
date_end (
status (	
offset (
limit (

branch_ids (	"r
GetBranchTagsResponse

error_code (

data (2 .merchantBasic.GetBranchTagsData"(
GetBranchTagsByIDsRequest
ids (	"s
GetBranchTagsByIDsResponse

error_code (

data (2.merchantBasic.BranchTagData"U
GetBranchTagsData1
branch_tags (2.merchantBasic.BranchTagData
total ("�


id (	
name (	
branches (	
create_staff_id (	

staff_name (	
status (	

created_at (

updated_at ("F
UpdateBranchTagRequest

id (	
name (	

branch_ids (	"D
UpdateBranchTagResponse

error_code (

UpdateBranchTagStatusRequest

id (	
status (	"J
UpdateBranchTagStatusResponse

error_code (

ShowBranchTagRequest

id (	"n
ShowBranchTagResponse

error_code (

data (2.merchantBasic.BranchTagDataB/Z
        , true);

        static::$is_initialized = true;
    }
}
