<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-private/staff_shift.proto

namespace GPBMetadata\Proto\MemberPrivate;

class StaffShift
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
&proto/member-private/staff_shift.protomemberPrivate"@
CreateStaffShiftRequest
	branch_id (	

shift_time ("�
CreateStaffShiftResponse

error_code (
error_message (	;
data (2-.memberPrivate.CreateStaffShiftResponse.Shift-
Shift

begin_time (
end_time ("-
GetStaffShiftTimeRequest
	branch_id (	"�
GetStaffShiftTimeResponse

error_code (
error_message (	<
data (2..memberPrivate.GetStaffShiftTimeResponse.Shift-
Shift

begin_time (
end_time ("^
ListStaffShiftRequest
	branch_id (	
begin_at (
end_at (
order_by (	"�
ListStaffShiftResponse

error_code (
error_message (	>
data (20.memberPrivate.ListStaffShiftResponse.StaffShiftR

StaffShift

id (	

begin_time (
end_time (

created_at (B/Z./proto;proto�Omy\\Crius\\MemberPrivateServerbproto3'
        , true);

        static::$is_initialized = true;
    }
}

