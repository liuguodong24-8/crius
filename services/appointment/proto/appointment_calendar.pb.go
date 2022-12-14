// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/appointment/appointment_calendar.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateTemplateCalendarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId  string                                           `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	BeginDate int64                                            `protobuf:"varint,2,opt,name=begin_date,json=beginDate,proto3" json:"begin_date,omitempty"`
	EndDate   int64                                            `protobuf:"varint,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Settings  []*UpdateTemplateCalendarRequest_CalendarSetting `protobuf:"bytes,4,rep,name=settings,proto3" json:"settings,omitempty"`
}

func (x *UpdateTemplateCalendarRequest) Reset() {
	*x = UpdateTemplateCalendarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_appointment_appointment_calendar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTemplateCalendarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTemplateCalendarRequest) ProtoMessage() {}

func (x *UpdateTemplateCalendarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_appointment_appointment_calendar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTemplateCalendarRequest.ProtoReflect.Descriptor instead.
func (*UpdateTemplateCalendarRequest) Descriptor() ([]byte, []int) {
	return file_proto_appointment_appointment_calendar_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateTemplateCalendarRequest) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *UpdateTemplateCalendarRequest) GetBeginDate() int64 {
	if x != nil {
		return x.BeginDate
	}
	return 0
}

func (x *UpdateTemplateCalendarRequest) GetEndDate() int64 {
	if x != nil {
		return x.EndDate
	}
	return 0
}

func (x *UpdateTemplateCalendarRequest) GetSettings() []*UpdateTemplateCalendarRequest_CalendarSetting {
	if x != nil {
		return x.Settings
	}
	return nil
}

type GetTemplateCalendarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId string `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	Year     int32  `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *GetTemplateCalendarRequest) Reset() {
	*x = GetTemplateCalendarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_appointment_appointment_calendar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateCalendarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateCalendarRequest) ProtoMessage() {}

func (x *GetTemplateCalendarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_appointment_appointment_calendar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateCalendarRequest.ProtoReflect.Descriptor instead.
func (*GetTemplateCalendarRequest) Descriptor() ([]byte, []int) {
	return file_proto_appointment_appointment_calendar_proto_rawDescGZIP(), []int{1}
}

func (x *GetTemplateCalendarRequest) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *GetTemplateCalendarRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type GetTemplateCalendarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    int32                                   `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string                                  `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Data         []*GetTemplateCalendarResponse_Calendar `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetTemplateCalendarResponse) Reset() {
	*x = GetTemplateCalendarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_appointment_appointment_calendar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateCalendarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateCalendarResponse) ProtoMessage() {}

func (x *GetTemplateCalendarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_appointment_appointment_calendar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateCalendarResponse.ProtoReflect.Descriptor instead.
func (*GetTemplateCalendarResponse) Descriptor() ([]byte, []int) {
	return file_proto_appointment_appointment_calendar_proto_rawDescGZIP(), []int{2}
}

func (x *GetTemplateCalendarResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *GetTemplateCalendarResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *GetTemplateCalendarResponse) GetData() []*GetTemplateCalendarResponse_Calendar {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateTemplateCalendarRequest_CalendarSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId string   `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	Weeks      []int32  `protobuf:"varint,2,rep,packed,name=weeks,proto3" json:"weeks,omitempty"`
	Category   string   `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	ThemeIds   []string `protobuf:"bytes,4,rep,name=theme_ids,json=themeIds,proto3" json:"theme_ids,omitempty"`
}

func (x *UpdateTemplateCalendarRequest_CalendarSetting) Reset() {
	*x = UpdateTemplateCalendarRequest_CalendarSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_appointment_appointment_calendar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTemplateCalendarRequest_CalendarSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTemplateCalendarRequest_CalendarSetting) ProtoMessage() {}

func (x *UpdateTemplateCalendarRequest_CalendarSetting) ProtoReflect() protoreflect.Message {
	mi := &file_proto_appointment_appointment_calendar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTemplateCalendarRequest_CalendarSetting.ProtoReflect.Descriptor instead.
func (*UpdateTemplateCalendarRequest_CalendarSetting) Descriptor() ([]byte, []int) {
	return file_proto_appointment_appointment_calendar_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UpdateTemplateCalendarRequest_CalendarSetting) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *UpdateTemplateCalendarRequest_CalendarSetting) GetWeeks() []int32 {
	if x != nil {
		return x.Weeks
	}
	return nil
}

func (x *UpdateTemplateCalendarRequest_CalendarSetting) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *UpdateTemplateCalendarRequest_CalendarSetting) GetThemeIds() []string {
	if x != nil {
		return x.ThemeIds
	}
	return nil
}

type GetTemplateCalendarResponse_Calendar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CalendarId       string   `protobuf:"bytes,1,opt,name=calendar_id,json=calendarId,proto3" json:"calendar_id,omitempty"`
	CalendarCategory string   `protobuf:"bytes,4,opt,name=calendar_category,json=calendarCategory,proto3" json:"calendar_category,omitempty"`
	BusinessDate     int64    `protobuf:"varint,2,opt,name=business_date,json=businessDate,proto3" json:"business_date,omitempty"`
	TemplateId       string   `protobuf:"bytes,3,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	TemplateName     string   `protobuf:"bytes,5,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
	TemplateColor    string   `protobuf:"bytes,6,opt,name=template_color,json=templateColor,proto3" json:"template_color,omitempty"`
	ThemeIds         []string `protobuf:"bytes,7,rep,name=theme_ids,json=themeIds,proto3" json:"theme_ids,omitempty"`
}

func (x *GetTemplateCalendarResponse_Calendar) Reset() {
	*x = GetTemplateCalendarResponse_Calendar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_appointment_appointment_calendar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemplateCalendarResponse_Calendar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemplateCalendarResponse_Calendar) ProtoMessage() {}

func (x *GetTemplateCalendarResponse_Calendar) ProtoReflect() protoreflect.Message {
	mi := &file_proto_appointment_appointment_calendar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemplateCalendarResponse_Calendar.ProtoReflect.Descriptor instead.
func (*GetTemplateCalendarResponse_Calendar) Descriptor() ([]byte, []int) {
	return file_proto_appointment_appointment_calendar_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetTemplateCalendarResponse_Calendar) GetCalendarId() string {
	if x != nil {
		return x.CalendarId
	}
	return ""
}

func (x *GetTemplateCalendarResponse_Calendar) GetCalendarCategory() string {
	if x != nil {
		return x.CalendarCategory
	}
	return ""
}

func (x *GetTemplateCalendarResponse_Calendar) GetBusinessDate() int64 {
	if x != nil {
		return x.BusinessDate
	}
	return 0
}

func (x *GetTemplateCalendarResponse_Calendar) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *GetTemplateCalendarResponse_Calendar) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *GetTemplateCalendarResponse_Calendar) GetTemplateColor() string {
	if x != nil {
		return x.TemplateColor
	}
	return ""
}

func (x *GetTemplateCalendarResponse_Calendar) GetThemeIds() []string {
	if x != nil {
		return x.ThemeIds
	}
	return nil
}

var File_proto_appointment_appointment_calendar_proto protoreflect.FileDescriptor

var file_proto_appointment_appointment_calendar_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xd2, 0x02, 0x0a, 0x1d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x62, 0x65, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x56, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x81, 0x01, 0x0a,
	0x0f, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x05, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x49, 0x64, 0x73,
	0x22, 0x4d, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22,
	0xb2, 0x03, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x87, 0x02, 0x0a, 0x08, 0x43,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x65, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x65, 0x6d,
	0x65, 0x49, 0x64, 0x73, 0x42, 0x2d, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x1b, 0x4f, 0x6d, 0x79, 0x5c, 0x43, 0x72, 0x69, 0x75,
	0x73, 0x5c, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_appointment_appointment_calendar_proto_rawDescOnce sync.Once
	file_proto_appointment_appointment_calendar_proto_rawDescData = file_proto_appointment_appointment_calendar_proto_rawDesc
)

func file_proto_appointment_appointment_calendar_proto_rawDescGZIP() []byte {
	file_proto_appointment_appointment_calendar_proto_rawDescOnce.Do(func() {
		file_proto_appointment_appointment_calendar_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_appointment_appointment_calendar_proto_rawDescData)
	})
	return file_proto_appointment_appointment_calendar_proto_rawDescData
}

var file_proto_appointment_appointment_calendar_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_appointment_appointment_calendar_proto_goTypes = []interface{}{
	(*UpdateTemplateCalendarRequest)(nil),                 // 0: appointment.UpdateTemplateCalendarRequest
	(*GetTemplateCalendarRequest)(nil),                    // 1: appointment.GetTemplateCalendarRequest
	(*GetTemplateCalendarResponse)(nil),                   // 2: appointment.GetTemplateCalendarResponse
	(*UpdateTemplateCalendarRequest_CalendarSetting)(nil), // 3: appointment.UpdateTemplateCalendarRequest.CalendarSetting
	(*GetTemplateCalendarResponse_Calendar)(nil),          // 4: appointment.GetTemplateCalendarResponse.Calendar
}
var file_proto_appointment_appointment_calendar_proto_depIdxs = []int32{
	3, // 0: appointment.UpdateTemplateCalendarRequest.settings:type_name -> appointment.UpdateTemplateCalendarRequest.CalendarSetting
	4, // 1: appointment.GetTemplateCalendarResponse.data:type_name -> appointment.GetTemplateCalendarResponse.Calendar
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_appointment_appointment_calendar_proto_init() }
func file_proto_appointment_appointment_calendar_proto_init() {
	if File_proto_appointment_appointment_calendar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_appointment_appointment_calendar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTemplateCalendarRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_appointment_appointment_calendar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateCalendarRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_appointment_appointment_calendar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateCalendarResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_appointment_appointment_calendar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTemplateCalendarRequest_CalendarSetting); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_appointment_appointment_calendar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTemplateCalendarResponse_Calendar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_appointment_appointment_calendar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_appointment_appointment_calendar_proto_goTypes,
		DependencyIndexes: file_proto_appointment_appointment_calendar_proto_depIdxs,
		MessageInfos:      file_proto_appointment_appointment_calendar_proto_msgTypes,
	}.Build()
	File_proto_appointment_appointment_calendar_proto = out.File
	file_proto_appointment_appointment_calendar_proto_rawDesc = nil
	file_proto_appointment_appointment_calendar_proto_goTypes = nil
	file_proto_appointment_appointment_calendar_proto_depIdxs = nil
}
