// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.0
// source: external/ingest/request/liveness.proto

package request

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

type Liveness struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	EventType        string                 `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	EntityUuid       string                 `protobuf:"bytes,2,opt,name=entity_uuid,json=entityUuid,proto3" json:"entity_uuid,omitempty"`
	ChefServerFqdn   string                 `protobuf:"bytes,3,opt,name=chef_server_fqdn,json=chefServerFqdn,proto3" json:"chef_server_fqdn,omitempty"`
	Source           string                 `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	MessageVersion   string                 `protobuf:"bytes,5,opt,name=message_version,json=messageVersion,proto3" json:"message_version,omitempty"`
	OrganizationName string                 `protobuf:"bytes,6,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty"`
	NodeName         string                 `protobuf:"bytes,7,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Liveness) Reset() {
	*x = Liveness{}
	mi := &file_external_ingest_request_liveness_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Liveness) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Liveness) ProtoMessage() {}

func (x *Liveness) ProtoReflect() protoreflect.Message {
	mi := &file_external_ingest_request_liveness_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Liveness.ProtoReflect.Descriptor instead.
func (*Liveness) Descriptor() ([]byte, []int) {
	return file_external_ingest_request_liveness_proto_rawDescGZIP(), []int{0}
}

func (x *Liveness) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *Liveness) GetEntityUuid() string {
	if x != nil {
		return x.EntityUuid
	}
	return ""
}

func (x *Liveness) GetChefServerFqdn() string {
	if x != nil {
		return x.ChefServerFqdn
	}
	return ""
}

func (x *Liveness) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Liveness) GetMessageVersion() string {
	if x != nil {
		return x.MessageVersion
	}
	return ""
}

func (x *Liveness) GetOrganizationName() string {
	if x != nil {
		return x.OrganizationName
	}
	return ""
}

func (x *Liveness) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

var File_external_ingest_request_liveness_proto protoreflect.FileDescriptor

var file_external_ingest_request_liveness_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6e, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xff, 0x01, 0x0a, 0x08, 0x4c,
	0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x55, 0x75, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x66, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x66, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x46, 0x71, 0x64,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_ingest_request_liveness_proto_rawDescOnce sync.Once
	file_external_ingest_request_liveness_proto_rawDescData = file_external_ingest_request_liveness_proto_rawDesc
)

func file_external_ingest_request_liveness_proto_rawDescGZIP() []byte {
	file_external_ingest_request_liveness_proto_rawDescOnce.Do(func() {
		file_external_ingest_request_liveness_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_ingest_request_liveness_proto_rawDescData)
	})
	return file_external_ingest_request_liveness_proto_rawDescData
}

var file_external_ingest_request_liveness_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_external_ingest_request_liveness_proto_goTypes = []any{
	(*Liveness)(nil), // 0: chef.automate.api.ingest.request.Liveness
}
var file_external_ingest_request_liveness_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_ingest_request_liveness_proto_init() }
func file_external_ingest_request_liveness_proto_init() {
	if File_external_ingest_request_liveness_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_ingest_request_liveness_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_ingest_request_liveness_proto_goTypes,
		DependencyIndexes: file_external_ingest_request_liveness_proto_depIdxs,
		MessageInfos:      file_external_ingest_request_liveness_proto_msgTypes,
	}.Build()
	File_external_ingest_request_liveness_proto = out.File
	file_external_ingest_request_liveness_proto_rawDesc = nil
	file_external_ingest_request_liveness_proto_goTypes = nil
	file_external_ingest_request_liveness_proto_depIdxs = nil
}
