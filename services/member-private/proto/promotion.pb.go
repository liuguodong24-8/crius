// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/member-private/promotion.proto

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

type PromotionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    int32  `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *PromotionResponse) Reset() {
	*x = PromotionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_member_private_promotion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromotionResponse) ProtoMessage() {}

func (x *PromotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_member_private_promotion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromotionResponse.ProtoReflect.Descriptor instead.
func (*PromotionResponse) Descriptor() ([]byte, []int) {
	return file_proto_member_private_promotion_proto_rawDescGZIP(), []int{0}
}

func (x *PromotionResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *PromotionResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type ListPromotionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BranchId  string   `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	Status    string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Limit     int32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset    int32    `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	OrderBy   string   `protobuf:"bytes,6,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	BranchIds []string `protobuf:"bytes,7,rep,name=branch_ids,json=branchIds,proto3" json:"branch_ids,omitempty"`
	WithPage  bool     `protobuf:"varint,8,opt,name=with_page,json=withPage,proto3" json:"with_page,omitempty"`
}

func (x *ListPromotionRequest) Reset() {
	*x = ListPromotionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_member_private_promotion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPromotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPromotionRequest) ProtoMessage() {}

func (x *ListPromotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_member_private_promotion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPromotionRequest.ProtoReflect.Descriptor instead.
func (*ListPromotionRequest) Descriptor() ([]byte, []int) {
	return file_proto_member_private_promotion_proto_rawDescGZIP(), []int{1}
}

func (x *ListPromotionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListPromotionRequest) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *ListPromotionRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ListPromotionRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListPromotionRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListPromotionRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListPromotionRequest) GetBranchIds() []string {
	if x != nil {
		return x.BranchIds
	}
	return nil
}

func (x *ListPromotionRequest) GetWithPage() bool {
	if x != nil {
		return x.WithPage
	}
	return false
}

type ListPromotionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    int32                                    `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string                                   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Data         *ListPromotionResponse_ListPromotionData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ListPromotionResponse) Reset() {
	*x = ListPromotionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_member_private_promotion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPromotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPromotionResponse) ProtoMessage() {}

func (x *ListPromotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_member_private_promotion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPromotionResponse.ProtoReflect.Descriptor instead.
func (*ListPromotionResponse) Descriptor() ([]byte, []int) {
	return file_proto_member_private_promotion_proto_rawDescGZIP(), []int{2}
}

func (x *ListPromotionResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *ListPromotionResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *ListPromotionResponse) GetData() *ListPromotionResponse_ListPromotionData {
	if x != nil {
		return x.Data
	}
	return nil
}

type Promotion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BeginAt   int64    `protobuf:"varint,3,opt,name=begin_at,json=beginAt,proto3" json:"begin_at,omitempty"`
	EndAt     int64    `protobuf:"varint,4,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	Status    string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt int64    `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	BranchIds []string `protobuf:"bytes,7,rep,name=branch_ids,json=branchIds,proto3" json:"branch_ids,omitempty"`
}

func (x *Promotion) Reset() {
	*x = Promotion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_member_private_promotion_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Promotion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Promotion) ProtoMessage() {}

func (x *Promotion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_member_private_promotion_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Promotion.ProtoReflect.Descriptor instead.
func (*Promotion) Descriptor() ([]byte, []int) {
	return file_proto_member_private_promotion_proto_rawDescGZIP(), []int{3}
}

func (x *Promotion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Promotion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Promotion) GetBeginAt() int64 {
	if x != nil {
		return x.BeginAt
	}
	return 0
}

func (x *Promotion) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *Promotion) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Promotion) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Promotion) GetBranchIds() []string {
	if x != nil {
		return x.BranchIds
	}
	return nil
}

type CreatePromotionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BeginAt   int64    `protobuf:"varint,2,opt,name=begin_at,json=beginAt,proto3" json:"begin_at,omitempty"`
	EndAt     int64    `protobuf:"varint,3,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	Status    string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	BranchIds []string `protobuf:"bytes,5,rep,name=branch_ids,json=branchIds,proto3" json:"branch_ids,omitempty"`
}

func (x *CreatePromotionRequest) Reset() {
	*x = CreatePromotionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_member_private_promotion_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePromotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePromotionRequest) ProtoMessage() {}

func (x *CreatePromotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_member_private_promotion_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePromotionRequest.ProtoReflect.Descriptor instead.
func (*CreatePromotionRequest) Descriptor() ([]byte, []int) {
	return file_proto_member_private_promotion_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePromotionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePromotionRequest) GetBeginAt() int64 {
	if x != nil {
		return x.BeginAt
	}
	return 0
}

func (x *CreatePromotionRequest) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *CreatePromotionRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreatePromotionRequest) GetBranchIds() []string {
	if x != nil {
		return x.BranchIds
	}
	return nil
}

type CreatePromotionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    int32      `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string     `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Data         *Promotion `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreatePromotionResponse) Reset() {
	*x = CreatePromotionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_member_private_promotion_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePromotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePromotionResponse) ProtoMessage() {}

func (x *CreatePromotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_member_private_promotion_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePromotionResponse.ProtoReflect.Descriptor instead.
func (*CreatePromotionResponse) Descriptor() ([]byte, []int) {
	return file_proto_member_private_promotion_proto_rawDescGZIP(), []int{5}
}

func (x *CreatePromotionResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *CreatePromotionResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *CreatePromotionResponse) GetData() *Promotion {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdatePromotionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BeginAt   int64    `protobuf:"varint,3,opt,name=begin_at,json=beginAt,proto3" json:"begin_at,omitempty"`
	EndAt     int64    `protobuf:"varint,4,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	Status    string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	BranchIds []string `protobuf:"bytes,6,rep,name=branch_ids,json=branchIds,proto3" json:"branch_ids,omitempty"`
}

func (x *UpdatePromotionRequest) Reset() {
	*x = UpdatePromotionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_member_private_promotion_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePromotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePromotionRequest) ProtoMessage() {}

func (x *UpdatePromotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_member_private_promotion_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePromotionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePromotionRequest) Descriptor() ([]byte, []int) {
	return file_proto_member_private_promotion_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePromotionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePromotionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePromotionRequest) GetBeginAt() int64 {
	if x != nil {
		return x.BeginAt
	}
	return 0
}

func (x *UpdatePromotionRequest) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *UpdatePromotionRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdatePromotionRequest) GetBranchIds() []string {
	if x != nil {
		return x.BranchIds
	}
	return nil
}

type ShowPromotionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PromotionId string `protobuf:"bytes,1,opt,name=promotion_id,json=promotionId,proto3" json:"promotion_id,omitempty"`
}

func (x *ShowPromotionRequest) Reset() {
	*x = ShowPromotionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_member_private_promotion_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowPromotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowPromotionRequest) ProtoMessage() {}

func (x *ShowPromotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_member_private_promotion_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowPromotionRequest.ProtoReflect.Descriptor instead.
func (*ShowPromotionRequest) Descriptor() ([]byte, []int) {
	return file_proto_member_private_promotion_proto_rawDescGZIP(), []int{7}
}

func (x *ShowPromotionRequest) GetPromotionId() string {
	if x != nil {
		return x.PromotionId
	}
	return ""
}

type ShowPromotionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    int32      `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string     `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Data         *Promotion `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ShowPromotionResponse) Reset() {
	*x = ShowPromotionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_member_private_promotion_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowPromotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowPromotionResponse) ProtoMessage() {}

func (x *ShowPromotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_member_private_promotion_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowPromotionResponse.ProtoReflect.Descriptor instead.
func (*ShowPromotionResponse) Descriptor() ([]byte, []int) {
	return file_proto_member_private_promotion_proto_rawDescGZIP(), []int{8}
}

func (x *ShowPromotionResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *ShowPromotionResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *ShowPromotionResponse) GetData() *Promotion {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListPromotionResponse_ListPromotionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*Promotion `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total int32        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListPromotionResponse_ListPromotionData) Reset() {
	*x = ListPromotionResponse_ListPromotionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_member_private_promotion_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPromotionResponse_ListPromotionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPromotionResponse_ListPromotionData) ProtoMessage() {}

func (x *ListPromotionResponse_ListPromotionData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_member_private_promotion_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPromotionResponse_ListPromotionData.ProtoReflect.Descriptor instead.
func (*ListPromotionResponse_ListPromotionData) Descriptor() ([]byte, []int) {
	return file_proto_member_private_promotion_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListPromotionResponse_ListPromotionData) GetData() []*Promotion {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListPromotionResponse_ListPromotionData) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_member_private_promotion_proto protoreflect.FileDescriptor

var file_proto_member_private_promotion_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x57, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xe4,
	0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69, 0x74, 0x68,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x77, 0x69, 0x74,
	0x68, 0x50, 0x61, 0x67, 0x65, 0x22, 0x80, 0x02, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x4a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x57, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xb7, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49,
	0x64, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x6e,
	0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa5, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x65, 0x67, 0x69, 0x6e,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x65, 0x67, 0x69, 0x6e,
	0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x73,
	0x22, 0x39, 0x0a, 0x14, 0x53, 0x68, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x15,
	0x53, 0x68, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x1d, 0x4f, 0x6d, 0x79, 0x5c, 0x43,
	0x72, 0x69, 0x75, 0x73, 0x5c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_member_private_promotion_proto_rawDescOnce sync.Once
	file_proto_member_private_promotion_proto_rawDescData = file_proto_member_private_promotion_proto_rawDesc
)

func file_proto_member_private_promotion_proto_rawDescGZIP() []byte {
	file_proto_member_private_promotion_proto_rawDescOnce.Do(func() {
		file_proto_member_private_promotion_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_member_private_promotion_proto_rawDescData)
	})
	return file_proto_member_private_promotion_proto_rawDescData
}

var file_proto_member_private_promotion_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_member_private_promotion_proto_goTypes = []interface{}{
	(*PromotionResponse)(nil),                       // 0: memberPrivate.PromotionResponse
	(*ListPromotionRequest)(nil),                    // 1: memberPrivate.ListPromotionRequest
	(*ListPromotionResponse)(nil),                   // 2: memberPrivate.ListPromotionResponse
	(*Promotion)(nil),                               // 3: memberPrivate.Promotion
	(*CreatePromotionRequest)(nil),                  // 4: memberPrivate.CreatePromotionRequest
	(*CreatePromotionResponse)(nil),                 // 5: memberPrivate.CreatePromotionResponse
	(*UpdatePromotionRequest)(nil),                  // 6: memberPrivate.UpdatePromotionRequest
	(*ShowPromotionRequest)(nil),                    // 7: memberPrivate.ShowPromotionRequest
	(*ShowPromotionResponse)(nil),                   // 8: memberPrivate.ShowPromotionResponse
	(*ListPromotionResponse_ListPromotionData)(nil), // 9: memberPrivate.ListPromotionResponse.ListPromotionData
}
var file_proto_member_private_promotion_proto_depIdxs = []int32{
	9, // 0: memberPrivate.ListPromotionResponse.data:type_name -> memberPrivate.ListPromotionResponse.ListPromotionData
	3, // 1: memberPrivate.CreatePromotionResponse.data:type_name -> memberPrivate.Promotion
	3, // 2: memberPrivate.ShowPromotionResponse.data:type_name -> memberPrivate.Promotion
	3, // 3: memberPrivate.ListPromotionResponse.ListPromotionData.data:type_name -> memberPrivate.Promotion
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_member_private_promotion_proto_init() }
func file_proto_member_private_promotion_proto_init() {
	if File_proto_member_private_promotion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_member_private_promotion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromotionResponse); i {
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
		file_proto_member_private_promotion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPromotionRequest); i {
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
		file_proto_member_private_promotion_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPromotionResponse); i {
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
		file_proto_member_private_promotion_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Promotion); i {
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
		file_proto_member_private_promotion_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePromotionRequest); i {
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
		file_proto_member_private_promotion_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePromotionResponse); i {
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
		file_proto_member_private_promotion_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePromotionRequest); i {
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
		file_proto_member_private_promotion_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowPromotionRequest); i {
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
		file_proto_member_private_promotion_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowPromotionResponse); i {
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
		file_proto_member_private_promotion_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPromotionResponse_ListPromotionData); i {
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
			RawDescriptor: file_proto_member_private_promotion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_member_private_promotion_proto_goTypes,
		DependencyIndexes: file_proto_member_private_promotion_proto_depIdxs,
		MessageInfos:      file_proto_member_private_promotion_proto_msgTypes,
	}.Build()
	File_proto_member_private_promotion_proto = out.File
	file_proto_member_private_promotion_proto_rawDesc = nil
	file_proto_member_private_promotion_proto_goTypes = nil
	file_proto_member_private_promotion_proto_depIdxs = nil
}
