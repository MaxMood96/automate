// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.0
// source: external/ingest/response/chef.proto

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

type ProcessChefRunResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessChefRunResponse) Reset() {
	*x = ProcessChefRunResponse{}
	mi := &file_external_ingest_response_chef_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessChefRunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessChefRunResponse) ProtoMessage() {}

func (x *ProcessChefRunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_ingest_response_chef_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessChefRunResponse.ProtoReflect.Descriptor instead.
func (*ProcessChefRunResponse) Descriptor() ([]byte, []int) {
	return file_external_ingest_response_chef_proto_rawDescGZIP(), []int{0}
}

var File_external_ingest_response_chef_proto protoreflect.FileDescriptor

var file_external_ingest_response_chef_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x66, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_external_ingest_response_chef_proto_rawDescOnce sync.Once
	file_external_ingest_response_chef_proto_rawDescData = file_external_ingest_response_chef_proto_rawDesc
)

func file_external_ingest_response_chef_proto_rawDescGZIP() []byte {
	file_external_ingest_response_chef_proto_rawDescOnce.Do(func() {
		file_external_ingest_response_chef_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_ingest_response_chef_proto_rawDescData)
	})
	return file_external_ingest_response_chef_proto_rawDescData
}

var file_external_ingest_response_chef_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_external_ingest_response_chef_proto_goTypes = []any{
	(*ProcessChefRunResponse)(nil), // 0: chef.automate.api.ingest.response.ProcessChefRunResponse
}
var file_external_ingest_response_chef_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_ingest_response_chef_proto_init() }
func file_external_ingest_response_chef_proto_init() {
	if File_external_ingest_response_chef_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_ingest_response_chef_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_ingest_response_chef_proto_goTypes,
		DependencyIndexes: file_external_ingest_response_chef_proto_depIdxs,
		MessageInfos:      file_external_ingest_response_chef_proto_msgTypes,
	}.Build()
	File_external_ingest_response_chef_proto = out.File
	file_external_ingest_response_chef_proto_rawDesc = nil
	file_external_ingest_response_chef_proto_goTypes = nil
	file_external_ingest_response_chef_proto_depIdxs = nil
}
