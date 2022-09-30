// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/merchant-basic/room_type_category.proto

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

type CreateRoomTypeCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Category int32  `protobuf:"varint,2,opt,name=category,proto3" json:"category,omitempty"`
	Status   string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateRoomTypeCategoryRequest) Reset() {
	*x = CreateRoomTypeCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomTypeCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomTypeCategoryRequest) ProtoMessage() {}

func (x *CreateRoomTypeCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomTypeCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomTypeCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_merchant_basic_room_type_category_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoomTypeCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRoomTypeCategoryRequest) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *CreateRoomTypeCategoryRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CreateRoomTypeCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    int32  `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *CreateRoomTypeCategoryResponse) Reset() {
	*x = CreateRoomTypeCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomTypeCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomTypeCategoryResponse) ProtoMessage() {}

func (x *CreateRoomTypeCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomTypeCategoryResponse.ProtoReflect.Descriptor instead.
func (*CreateRoomTypeCategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_merchant_basic_room_type_category_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoomTypeCategoryResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *CreateRoomTypeCategoryResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type GetRoomTypeCategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Offset   int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit    int32  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Category int32  `protobuf:"varint,5,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *GetRoomTypeCategoriesRequest) Reset() {
	*x = GetRoomTypeCategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomTypeCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomTypeCategoriesRequest) ProtoMessage() {}

func (x *GetRoomTypeCategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomTypeCategoriesRequest.ProtoReflect.Descriptor instead.
func (*GetRoomTypeCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_proto_merchant_basic_room_type_category_proto_rawDescGZIP(), []int{2}
}

func (x *GetRoomTypeCategoriesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetRoomTypeCategoriesRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetRoomTypeCategoriesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetRoomTypeCategoriesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetRoomTypeCategoriesRequest) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

type GetRoomTypeCategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    int32                   `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string                  `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Data         *RoomTypeCategoriesData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetRoomTypeCategoriesResponse) Reset() {
	*x = GetRoomTypeCategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomTypeCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomTypeCategoriesResponse) ProtoMessage() {}

func (x *GetRoomTypeCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomTypeCategoriesResponse.ProtoReflect.Descriptor instead.
func (*GetRoomTypeCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_proto_merchant_basic_room_type_category_proto_rawDescGZIP(), []int{3}
}

func (x *GetRoomTypeCategoriesResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *GetRoomTypeCategoriesResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *GetRoomTypeCategoriesResponse) GetData() *RoomTypeCategoriesData {
	if x != nil {
		return x.Data
	}
	return nil
}

type RoomTypeCategoriesData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomTypeCategories []*RoomTypeCategory `protobuf:"bytes,1,rep,name=room_type_categories,json=roomTypeCategories,proto3" json:"room_type_categories,omitempty"`
	Total              int32               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *RoomTypeCategoriesData) Reset() {
	*x = RoomTypeCategoriesData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomTypeCategoriesData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomTypeCategoriesData) ProtoMessage() {}

func (x *RoomTypeCategoriesData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomTypeCategoriesData.ProtoReflect.Descriptor instead.
func (*RoomTypeCategoriesData) Descriptor() ([]byte, []int) {
	return file_proto_merchant_basic_room_type_category_proto_rawDescGZIP(), []int{4}
}

func (x *RoomTypeCategoriesData) GetRoomTypeCategories() []*RoomTypeCategory {
	if x != nil {
		return x.RoomTypeCategories
	}
	return nil
}

func (x *RoomTypeCategoriesData) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type RoomTypeCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category int32  `protobuf:"varint,3,opt,name=category,proto3" json:"category,omitempty"`
	Status   string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RoomTypeCategory) Reset() {
	*x = RoomTypeCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomTypeCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomTypeCategory) ProtoMessage() {}

func (x *RoomTypeCategory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomTypeCategory.ProtoReflect.Descriptor instead.
func (*RoomTypeCategory) Descriptor() ([]byte, []int) {
	return file_proto_merchant_basic_room_type_category_proto_rawDescGZIP(), []int{5}
}

func (x *RoomTypeCategory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoomTypeCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoomTypeCategory) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *RoomTypeCategory) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateRoomTypeCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category int32  `protobuf:"varint,3,opt,name=category,proto3" json:"category,omitempty"`
	Status   string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateRoomTypeCategoryRequest) Reset() {
	*x = UpdateRoomTypeCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoomTypeCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoomTypeCategoryRequest) ProtoMessage() {}

func (x *UpdateRoomTypeCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoomTypeCategoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoomTypeCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_merchant_basic_room_type_category_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateRoomTypeCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRoomTypeCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRoomTypeCategoryRequest) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *UpdateRoomTypeCategoryRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateRoomTypeCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    int32  `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *UpdateRoomTypeCategoryResponse) Reset() {
	*x = UpdateRoomTypeCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoomTypeCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoomTypeCategoryResponse) ProtoMessage() {}

func (x *UpdateRoomTypeCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_basic_room_type_category_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoomTypeCategoryResponse.ProtoReflect.Descriptor instead.
func (*UpdateRoomTypeCategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_merchant_basic_room_type_category_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRoomTypeCategoryResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *UpdateRoomTypeCategoryResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_proto_merchant_basic_room_type_category_proto protoreflect.FileDescriptor

var file_proto_merchant_basic_room_type_category_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x2d, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x22, 0x67,
	0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x64, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x94, 0x01,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x22, 0x9e, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x81, 0x01, 0x0a, 0x16, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x51, 0x0a, 0x14, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x12, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x6a, 0x0a, 0x10, 0x52, 0x6f, 0x6f,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x77, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x64,
	0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x2f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x1d, 0x4f, 0x6d, 0x79, 0x5c, 0x43, 0x72, 0x69, 0x75,
	0x73, 0x5c, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_merchant_basic_room_type_category_proto_rawDescOnce sync.Once
	file_proto_merchant_basic_room_type_category_proto_rawDescData = file_proto_merchant_basic_room_type_category_proto_rawDesc
)

func file_proto_merchant_basic_room_type_category_proto_rawDescGZIP() []byte {
	file_proto_merchant_basic_room_type_category_proto_rawDescOnce.Do(func() {
		file_proto_merchant_basic_room_type_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_merchant_basic_room_type_category_proto_rawDescData)
	})
	return file_proto_merchant_basic_room_type_category_proto_rawDescData
}

var file_proto_merchant_basic_room_type_category_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_merchant_basic_room_type_category_proto_goTypes = []interface{}{
	(*CreateRoomTypeCategoryRequest)(nil),  // 0: merchantBasic.CreateRoomTypeCategoryRequest
	(*CreateRoomTypeCategoryResponse)(nil), // 1: merchantBasic.CreateRoomTypeCategoryResponse
	(*GetRoomTypeCategoriesRequest)(nil),   // 2: merchantBasic.GetRoomTypeCategoriesRequest
	(*GetRoomTypeCategoriesResponse)(nil),  // 3: merchantBasic.GetRoomTypeCategoriesResponse
	(*RoomTypeCategoriesData)(nil),         // 4: merchantBasic.RoomTypeCategoriesData
	(*RoomTypeCategory)(nil),               // 5: merchantBasic.RoomTypeCategory
	(*UpdateRoomTypeCategoryRequest)(nil),  // 6: merchantBasic.UpdateRoomTypeCategoryRequest
	(*UpdateRoomTypeCategoryResponse)(nil), // 7: merchantBasic.UpdateRoomTypeCategoryResponse
}
var file_proto_merchant_basic_room_type_category_proto_depIdxs = []int32{
	4, // 0: merchantBasic.GetRoomTypeCategoriesResponse.data:type_name -> merchantBasic.RoomTypeCategoriesData
	5, // 1: merchantBasic.RoomTypeCategoriesData.room_type_categories:type_name -> merchantBasic.RoomTypeCategory
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_merchant_basic_room_type_category_proto_init() }
func file_proto_merchant_basic_room_type_category_proto_init() {
	if File_proto_merchant_basic_room_type_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_merchant_basic_room_type_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomTypeCategoryRequest); i {
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
		file_proto_merchant_basic_room_type_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomTypeCategoryResponse); i {
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
		file_proto_merchant_basic_room_type_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomTypeCategoriesRequest); i {
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
		file_proto_merchant_basic_room_type_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomTypeCategoriesResponse); i {
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
		file_proto_merchant_basic_room_type_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomTypeCategoriesData); i {
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
		file_proto_merchant_basic_room_type_category_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomTypeCategory); i {
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
		file_proto_merchant_basic_room_type_category_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoomTypeCategoryRequest); i {
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
		file_proto_merchant_basic_room_type_category_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoomTypeCategoryResponse); i {
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
			RawDescriptor: file_proto_merchant_basic_room_type_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_merchant_basic_room_type_category_proto_goTypes,
		DependencyIndexes: file_proto_merchant_basic_room_type_category_proto_depIdxs,
		MessageInfos:      file_proto_merchant_basic_room_type_category_proto_msgTypes,
	}.Build()
	File_proto_merchant_basic_room_type_category_proto = out.File
	file_proto_merchant_basic_room_type_category_proto_rawDesc = nil
	file_proto_merchant_basic_room_type_category_proto_goTypes = nil
	file_proto_merchant_basic_room_type_category_proto_depIdxs = nil
}
