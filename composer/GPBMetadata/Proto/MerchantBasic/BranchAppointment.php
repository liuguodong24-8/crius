<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/merchant-basic/branch_appointment.proto

namespace GPBMetadata\Proto\MerchantBasic;

class BranchAppointment
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Proto\MerchantBasic\BranchBusiness::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
-proto/merchant-basic/branch_appointment.proto
SaveBranchAppointmentRequest
	branch_id (	
open_appointment (
appointment_granularity (.
vr (2".merchantBasic.BranchAppointmentVR
video (	
environment (	
meal (	
price (	
hot	 ("J
SaveBranchAppointmentResponse

error_code (

ShowBranchAppointmentRequest
	branch_id (	"~
ShowBranchAppointmentResponse

error_code (

data (2$.merchantBasic.BranchAppointmentData"j
&UpdateBranchAppointmentRoomTypeRequest-
	room_type (2.merchantBasic.RoomTypeNum
	branch_id (	":
RoomTypeNum
room_type_id (	

\'UpdateBranchAppointmentRoomTypeResponse

error_code (

BranchAppointmentData5
business_hours (2.merchantBasic.BranchBusiness
open_appointment (
appointment_granularity (.
vr (2".merchantBasic.BranchAppointmentVR
video (	
environment (	
meal (	
price (	
hot	 (
	branch_id
 (	

room_types (	"K
(ShowBranchAppointmentBusinessTimeRequest
	branch_id (	
date ("�
)ShowBranchAppointmentBusinessTimeResponse

error_code (

data (24.merchantBasic.ShowBranchAppointmentBusinessTimeData"�
%ShowBranchAppointmentBusinessTimeData

begin_time (	
end_time (	
is_next_day (
granularity (
open_appointment (
hot ("E
BranchAppointmentVR
url (	
name (	
description (	B/Z
        , true);

        static::$is_initialized = true;
    }
}
