<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment_calendar.proto

namespace GPBMetadata\Proto\Appointment;

class AppointmentCalendar
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
,proto/appointment/appointment_calendar.protoappointment"�
UpdateTemplateCalendarRequest
	branch_id (	

begin_date (
end_date (L
settings (2:.appointment.UpdateTemplateCalendarRequest.CalendarSettingZ
CalendarSetting
template_id (	
weeks (
category (	
	theme_ids (	"=
GetTemplateCalendarRequest
	branch_id (	
year ("�
GetTemplateCalendarResponse

error_code (
error_message (	?
data (21.appointment.GetTemplateCalendarResponse.Calendar�
Calendar
calendar_id (	
calendar_category (	
business_date (
template_id (	
template_name (	
template_color (	
	theme_ids (	B-Z./proto;proto�Omy\\Crius\\AppointmentServerbproto3'
        , true);

        static::$is_initialized = true;
    }
}

