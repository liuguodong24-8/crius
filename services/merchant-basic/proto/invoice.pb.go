// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/merchant-basic/invoice.proto

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

type CreateInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action      string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	InvoiceData string `protobuf:"bytes,2,opt,name=invoice_data,json=invoiceData,proto3" json:"invoice_data,omitempty"`
}

func (x *CreateInvoiceRequest) Reset() {
	*x = CreateInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_basic_invoice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvoiceRequest) ProtoMessage() {}

func (x *CreateInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_basic_invoice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInvoiceRequest.ProtoReflect.Descriptor instead.
func (*CreateInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_proto_merchant_basic_invoice_proto_rawDescGZIP(), []int{0}
}

func (x *CreateInvoiceRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *CreateInvoiceRequest) GetInvoiceData() string {
	if x != nil {
		return x.InvoiceData
	}
	return ""
}

type CreateInvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    int32  `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *CreateInvoiceResponse) Reset() {
	*x = CreateInvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_merchant_basic_invoice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvoiceResponse) ProtoMessage() {}

func (x *CreateInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_merchant_basic_invoice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInvoiceResponse.ProtoReflect.Descriptor instead.
func (*CreateInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_proto_merchant_basic_invoice_proto_rawDescGZIP(), []int{1}
}

func (x *CreateInvoiceResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *CreateInvoiceResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_proto_merchant_basic_invoice_proto protoreflect.FileDescriptor

var file_proto_merchant_basic_invoice_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x2d, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x22, 0x51, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x5b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x2f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xca, 0x02, 0x1d, 0x4f, 0x6d, 0x79, 0x5c, 0x43, 0x72, 0x69, 0x75, 0x73,
	0x5c, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_merchant_basic_invoice_proto_rawDescOnce sync.Once
	file_proto_merchant_basic_invoice_proto_rawDescData = file_proto_merchant_basic_invoice_proto_rawDesc
)

func file_proto_merchant_basic_invoice_proto_rawDescGZIP() []byte {
	file_proto_merchant_basic_invoice_proto_rawDescOnce.Do(func() {
		file_proto_merchant_basic_invoice_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_merchant_basic_invoice_proto_rawDescData)
	})
	return file_proto_merchant_basic_invoice_proto_rawDescData
}

var file_proto_merchant_basic_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_merchant_basic_invoice_proto_goTypes = []interface{}{
	(*CreateInvoiceRequest)(nil),  // 0: merchantBasic.CreateInvoiceRequest
	(*CreateInvoiceResponse)(nil), // 1: merchantBasic.CreateInvoiceResponse
}
var file_proto_merchant_basic_invoice_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_merchant_basic_invoice_proto_init() }
func file_proto_merchant_basic_invoice_proto_init() {
	if File_proto_merchant_basic_invoice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_merchant_basic_invoice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInvoiceRequest); i {
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
		file_proto_merchant_basic_invoice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInvoiceResponse); i {
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
			RawDescriptor: file_proto_merchant_basic_invoice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_merchant_basic_invoice_proto_goTypes,
		DependencyIndexes: file_proto_merchant_basic_invoice_proto_depIdxs,
		MessageInfos:      file_proto_merchant_basic_invoice_proto_msgTypes,
	}.Build()
	File_proto_merchant_basic_invoice_proto = out.File
	file_proto_merchant_basic_invoice_proto_rawDesc = nil
	file_proto_merchant_basic_invoice_proto_goTypes = nil
	file_proto_merchant_basic_invoice_proto_depIdxs = nil
}