// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.0
// source: external/cfgmgmt/response/errors.proto

package response

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

// Errors contains a list of the most common Chef Infra error type/message
// combinations among nodes in the active project as filtered according to the
// request.
type Errors struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Errors        []*ErrorCount          `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Errors) Reset() {
	*x = Errors{}
	mi := &file_external_cfgmgmt_response_errors_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Errors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Errors) ProtoMessage() {}

func (x *Errors) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_response_errors_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Errors.ProtoReflect.Descriptor instead.
func (*Errors) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_response_errors_proto_rawDescGZIP(), []int{0}
}

func (x *Errors) GetErrors() []*ErrorCount {
	if x != nil {
		return x.Errors
	}
	return nil
}

// ErrorCount gives the number of occurrences (count) of the error specified by
// the type and message among the nodes included by the request parameters
type ErrorCount struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Count         int32                  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ErrorMessage  string                 `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ErrorCount) Reset() {
	*x = ErrorCount{}
	mi := &file_external_cfgmgmt_response_errors_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrorCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorCount) ProtoMessage() {}

func (x *ErrorCount) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_response_errors_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorCount.ProtoReflect.Descriptor instead.
func (*ErrorCount) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_response_errors_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ErrorCount) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ErrorCount) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_external_cfgmgmt_response_errors_proto protoreflect.FileDescriptor

var file_external_cfgmgmt_response_errors_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67,
	0x6d, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x66, 0x67, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0x0a, 0x06,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x46, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67,
	0x6d, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x5b,
	0x0a, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_cfgmgmt_response_errors_proto_rawDescOnce sync.Once
	file_external_cfgmgmt_response_errors_proto_rawDescData = file_external_cfgmgmt_response_errors_proto_rawDesc
)

func file_external_cfgmgmt_response_errors_proto_rawDescGZIP() []byte {
	file_external_cfgmgmt_response_errors_proto_rawDescOnce.Do(func() {
		file_external_cfgmgmt_response_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_cfgmgmt_response_errors_proto_rawDescData)
	})
	return file_external_cfgmgmt_response_errors_proto_rawDescData
}

var file_external_cfgmgmt_response_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_external_cfgmgmt_response_errors_proto_goTypes = []any{
	(*Errors)(nil),     // 0: chef.automate.api.cfgmgmt.response.Errors
	(*ErrorCount)(nil), // 1: chef.automate.api.cfgmgmt.response.ErrorCount
}
var file_external_cfgmgmt_response_errors_proto_depIdxs = []int32{
	1, // 0: chef.automate.api.cfgmgmt.response.Errors.errors:type_name -> chef.automate.api.cfgmgmt.response.ErrorCount
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_external_cfgmgmt_response_errors_proto_init() }
func file_external_cfgmgmt_response_errors_proto_init() {
	if File_external_cfgmgmt_response_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_cfgmgmt_response_errors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_cfgmgmt_response_errors_proto_goTypes,
		DependencyIndexes: file_external_cfgmgmt_response_errors_proto_depIdxs,
		MessageInfos:      file_external_cfgmgmt_response_errors_proto_msgTypes,
	}.Build()
	File_external_cfgmgmt_response_errors_proto = out.File
	file_external_cfgmgmt_response_errors_proto_rawDesc = nil
	file_external_cfgmgmt_response_errors_proto_goTypes = nil
	file_external_cfgmgmt_response_errors_proto_depIdxs = nil
}
