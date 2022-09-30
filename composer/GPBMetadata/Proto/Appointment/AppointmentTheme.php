<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/appointment/appointment_theme.proto

namespace GPBMetadata\Proto\Appointment;

class AppointmentTheme
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
)proto/appointment/appointment_theme.protoappointment"
Theme8
appointment_theme (2.appointment.AppointmentTheme<
theme_packages (2$.appointment.AppointmentThemePackage"�

appointment_theme (2.appointment.AppointmentTheme;


AppointmentTheme

id (	
color (	
feature_ids (	7
contents (2%.appointment.AppointmentTheme.Content
style (	
images (	
video (	
details (	
weight	 (
status
 (	
name (	
category_id (	(
Content
name (	
content (	"�
AppointmentThemePackage
name (	>
packages (2,.appointment.AppointmentThemePackage.Package

decoration (	
staffing (	I
custom_configs (21.appointment.AppointmentThemePackage.CustomConfigA

room_types (2-.appointment.AppointmentThemePackage.RoomType

id (	\'
Package

id (	
category (	,
CustomConfig
name (	
config (	%
RoomType

id (	
price ("B
CreateAppointmentThemeRequest!
theme (2.appointment.Theme"B
UpdateAppointmentThemeRequest!
theme (2.appointment.Theme"A
#UpdateAppointmentThemeStatusRequest

id (	
status (	"o
GetAppointmentThemesRequest
name (	
status (	
category_id (	
offset (
limit ("�
GetAppointmentThemesResponse

error_code (

data (2..appointment.GetAppointmentThemesResponse.DataD
Data
total (-
themes (2.appointment.AppointmentTheme")
ShowAppointmentThemeRequest

id (	"k
ShowAppointmentThemeResponse

error_code (

data (2.appointment.Theme"p
%GetAppointmentThemesByRoomTypeRequest
room_type_id (	
	branch_id (	
date (
theme_id (	"}
&GetAppointmentThemesByRoomTypeResponse

error_code (

data (2.appointment.ThemeRoomTypeB-Z
        , true);

        static::$is_initialized = true;
    }
}
