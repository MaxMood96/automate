// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.0
// source: external/secrets/secrets.proto

package secrets

import (
	context "context"
	_ "github.com/chef/automate/api/external/annotations/iam"
	query "github.com/chef/automate/api/external/common/query"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Return the results in ascending or descending order.
type Query_OrderType int32

const (
	Query_ASC  Query_OrderType = 0
	Query_DESC Query_OrderType = 1
)

// Enum value maps for Query_OrderType.
var (
	Query_OrderType_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	Query_OrderType_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x Query_OrderType) Enum() *Query_OrderType {
	p := new(Query_OrderType)
	*p = x
	return p
}

func (x Query_OrderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Query_OrderType) Descriptor() protoreflect.EnumDescriptor {
	return file_external_secrets_secrets_proto_enumTypes[0].Descriptor()
}

func (Query_OrderType) Type() protoreflect.EnumType {
	return &file_external_secrets_secrets_proto_enumTypes[0]
}

func (x Query_OrderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Query_OrderType.Descriptor instead.
func (Query_OrderType) EnumDescriptor() ([]byte, []int) {
	return file_external_secrets_secrets_proto_rawDescGZIP(), []int{3, 0}
}

type UpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	mi := &file_external_secrets_secrets_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_secrets_secrets_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_external_secrets_secrets_proto_rawDescGZIP(), []int{0}
}

type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_external_secrets_secrets_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_secrets_secrets_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_external_secrets_secrets_proto_rawDescGZIP(), []int{1}
}

type Id struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique node ID (UUID).
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Id) Reset() {
	*x = Id{}
	mi := &file_external_secrets_secrets_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_external_secrets_secrets_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_external_secrets_secrets_proto_rawDescGZIP(), []int{2}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Query struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Use filters to limit the set of secrets.
	Filters []*query.Filter `protobuf:"bytes,20,rep,name=filters,proto3" json:"filters,omitempty"`
	Order   Query_OrderType `protobuf:"varint,21,opt,name=order,proto3,enum=chef.automate.api.secrets.Query_OrderType" json:"order,omitempty"`
	// Sort the results on a specific field.
	Sort string `protobuf:"bytes,22,opt,name=sort,proto3" json:"sort,omitempty"`
	// Starting page for the results.
	Page int32 `protobuf:"varint,23,opt,name=page,proto3" json:"page,omitempty"`
	// The number of results on each page.
	PerPage       int32 `protobuf:"varint,24,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Query) Reset() {
	*x = Query{}
	mi := &file_external_secrets_secrets_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_external_secrets_secrets_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_external_secrets_secrets_proto_rawDescGZIP(), []int{3}
}

func (x *Query) GetFilters() []*query.Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *Query) GetOrder() Query_OrderType {
	if x != nil {
		return x.Order
	}
	return Query_ASC
}

func (x *Query) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *Query) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Query) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type Secret struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique node ID (UUID).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// User-specified name for the secret.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Type of credential: ssh, winrm, sudo, aws, azure, gcp, service_now
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Timestamp denoting when the secret was last modified.
	LastModified *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	// Tags to associate with the secret.
	Tags []*query.Kv `protobuf:"bytes,21,rep,name=tags,proto3" json:"tags,omitempty"`
	// Secret data, where the kv structs for the credential data live.
	Data          []*query.Kv `protobuf:"bytes,22,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Secret) Reset() {
	*x = Secret{}
	mi := &file_external_secrets_secrets_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_external_secrets_secrets_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_external_secrets_secrets_proto_rawDescGZIP(), []int{4}
}

func (x *Secret) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Secret) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Secret) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Secret) GetLastModified() *timestamppb.Timestamp {
	if x != nil {
		return x.LastModified
	}
	return nil
}

func (x *Secret) GetTags() []*query.Kv {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Secret) GetData() []*query.Kv {
	if x != nil {
		return x.Data
	}
	return nil
}

type Secrets struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of secrets.
	Secrets []*Secret `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty"`
	// Total count of secrets
	Total         int32 `protobuf:"varint,20,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Secrets) Reset() {
	*x = Secrets{}
	mi := &file_external_secrets_secrets_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Secrets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secrets) ProtoMessage() {}

func (x *Secrets) ProtoReflect() protoreflect.Message {
	mi := &file_external_secrets_secrets_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secrets.ProtoReflect.Descriptor instead.
func (*Secrets) Descriptor() ([]byte, []int) {
	return file_external_secrets_secrets_proto_rawDescGZIP(), []int{5}
}

func (x *Secrets) GetSecrets() []*Secret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *Secrets) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_external_secrets_secrets_proto protoreflect.FileDescriptor

var file_external_secrets_secrets_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xee, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x40, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x14, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x40, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x1e, 0x0a, 0x09, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x22, 0xf1, 0x01, 0x0a, 0x06, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x0d,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x36, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4b, 0x76, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x16, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x4b, 0x76, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5c, 0x0a,
	0x07, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x07, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xb5, 0x06, 0x0a, 0x0e,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x93,
	0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x1a, 0x1d, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x49, 0x64, 0x22, 0x47, 0x8a, 0xb5, 0x18,
	0x29, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x12, 0x16, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x3a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x12, 0x98, 0x01, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1d, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x49, 0x64, 0x1a, 0x21, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22,
	0x4e, 0x8a, 0xb5, 0x18, 0x2b, 0x0a, 0x14, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x13, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x3a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x67, 0x65, 0x74,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0xac, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x1a, 0x29, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x8a, 0xb5, 0x18, 0x2e, 0x0a, 0x14,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x3a,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x16, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x32, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xa5,
	0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x49, 0x64, 0x1a, 0x29, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x51, 0x8a, 0xb5, 0x18, 0x2e, 0x0a, 0x14, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x3a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x16, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x3a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x69,
	0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x9a, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x20, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x1a, 0x22, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x22, 0x4c, 0x8a, 0xb5, 0x18, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x3a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x14, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x3a, 0x6c, 0x69,
	0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_secrets_secrets_proto_rawDescOnce sync.Once
	file_external_secrets_secrets_proto_rawDescData = file_external_secrets_secrets_proto_rawDesc
)

func file_external_secrets_secrets_proto_rawDescGZIP() []byte {
	file_external_secrets_secrets_proto_rawDescOnce.Do(func() {
		file_external_secrets_secrets_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_secrets_secrets_proto_rawDescData)
	})
	return file_external_secrets_secrets_proto_rawDescData
}

var file_external_secrets_secrets_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_external_secrets_secrets_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_external_secrets_secrets_proto_goTypes = []any{
	(Query_OrderType)(0),          // 0: chef.automate.api.secrets.Query.OrderType
	(*UpdateResponse)(nil),        // 1: chef.automate.api.secrets.UpdateResponse
	(*DeleteResponse)(nil),        // 2: chef.automate.api.secrets.DeleteResponse
	(*Id)(nil),                    // 3: chef.automate.api.secrets.Id
	(*Query)(nil),                 // 4: chef.automate.api.secrets.Query
	(*Secret)(nil),                // 5: chef.automate.api.secrets.Secret
	(*Secrets)(nil),               // 6: chef.automate.api.secrets.Secrets
	(*query.Filter)(nil),          // 7: chef.automate.api.common.query.Filter
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*query.Kv)(nil),              // 9: chef.automate.api.common.query.Kv
}
var file_external_secrets_secrets_proto_depIdxs = []int32{
	7,  // 0: chef.automate.api.secrets.Query.filters:type_name -> chef.automate.api.common.query.Filter
	0,  // 1: chef.automate.api.secrets.Query.order:type_name -> chef.automate.api.secrets.Query.OrderType
	8,  // 2: chef.automate.api.secrets.Secret.last_modified:type_name -> google.protobuf.Timestamp
	9,  // 3: chef.automate.api.secrets.Secret.tags:type_name -> chef.automate.api.common.query.Kv
	9,  // 4: chef.automate.api.secrets.Secret.data:type_name -> chef.automate.api.common.query.Kv
	5,  // 5: chef.automate.api.secrets.Secrets.secrets:type_name -> chef.automate.api.secrets.Secret
	5,  // 6: chef.automate.api.secrets.SecretsService.Create:input_type -> chef.automate.api.secrets.Secret
	3,  // 7: chef.automate.api.secrets.SecretsService.Read:input_type -> chef.automate.api.secrets.Id
	5,  // 8: chef.automate.api.secrets.SecretsService.Update:input_type -> chef.automate.api.secrets.Secret
	3,  // 9: chef.automate.api.secrets.SecretsService.Delete:input_type -> chef.automate.api.secrets.Id
	4,  // 10: chef.automate.api.secrets.SecretsService.List:input_type -> chef.automate.api.secrets.Query
	3,  // 11: chef.automate.api.secrets.SecretsService.Create:output_type -> chef.automate.api.secrets.Id
	5,  // 12: chef.automate.api.secrets.SecretsService.Read:output_type -> chef.automate.api.secrets.Secret
	1,  // 13: chef.automate.api.secrets.SecretsService.Update:output_type -> chef.automate.api.secrets.UpdateResponse
	2,  // 14: chef.automate.api.secrets.SecretsService.Delete:output_type -> chef.automate.api.secrets.DeleteResponse
	6,  // 15: chef.automate.api.secrets.SecretsService.List:output_type -> chef.automate.api.secrets.Secrets
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_external_secrets_secrets_proto_init() }
func file_external_secrets_secrets_proto_init() {
	if File_external_secrets_secrets_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_secrets_secrets_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_secrets_secrets_proto_goTypes,
		DependencyIndexes: file_external_secrets_secrets_proto_depIdxs,
		EnumInfos:         file_external_secrets_secrets_proto_enumTypes,
		MessageInfos:      file_external_secrets_secrets_proto_msgTypes,
	}.Build()
	File_external_secrets_secrets_proto = out.File
	file_external_secrets_secrets_proto_rawDesc = nil
	file_external_secrets_secrets_proto_goTypes = nil
	file_external_secrets_secrets_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SecretsServiceClient is the client API for SecretsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretsServiceClient interface {
	// Create a secret
	//
	// Creates a secret. Requires values for name, type, and data.
	//
	// Supported types: ssh, winrm, sudo, aws, azure, gcp, service_now
	// Supported keys by type:
	// ssh: username, password, key
	// winrm: username, password
	// sudo: username, password
	// service_now: username, password
	// aws: AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_SESSION_TOKEN
	// azure: AZURE_CLIENT_ID, AZURE_CLIENT_SECRET, AZURE_TENANT_ID
	// azure: AZURE_SUBSCRIPTION_ID is optional
	// gcp: GOOGLE_CREDENTIALS_JSON
	//
	// Example:
	// ```
	// {
	// "name": "my ssh secret",
	// "type": "ssh",
	// "data": [
	// { "key": "username", "value": "vagrant" },
	// { "key": "password", "value": "vagrant"}
	// ]
	// }
	// ```
	//
	// Authorization Action:
	// ```
	// ```
	//
	//secrets:secrets:create
	Create(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Id, error)
	// Read a secret
	//
	// Reads a secret given the ID of the secret.
	// Note that the secret information (password and key values) will not be returned by the API, as a safety measure.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//secrets:secrets:get
	Read(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Secret, error)
	// Update a secret
	//
	// Updates a secret.
	// This is a PATCH operation, meaning the details sent in will override/replace those stored in the DB.
	// Secret information that is not in the body of the request will persist.
	//
	// Example:
	// ```
	// given a credential with a username and password, a user could update the password by passing in the following body,
	// and the name of the secret as well as the username for the secret be unchanged:
	//
	// {
	// "id": "525c013a-2ab3-4e6f-9005-51bc620e9157",
	// "data": [
	// { "key": "password", "value": "new-value"}
	// ]
	// }
	// ```
	//
	// Authorization Action:
	// ```
	// ```
	//
	//secrets:secrets:update
	Update(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete a secret
	//
	// Deletes a secret given the ID of the secret.
	// Note that any nodes that were using the secret will no longer be associated with the deleted secret.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//secrets:secrets:delete
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*DeleteResponse, error)
	// List and filter secrets
	//
	// Makes a list of secrets.
	// Supports filtering, pagination, and sorting.
	// Adding a filter narrows the list of secrets to only those that match the filter or filters.
	// Supported filters: type, name
	// Supported sort types: name, type, last modified
	//
	// Example:
	// ```
	// {
	// "sort": "type",
	// "order": "ASC",
	// "filters": [
	// { "key": "type", "values": ["ssh","winrm","sudo"] }
	// ],
	// "page":1,
	// "per_page":100
	// }
	// ```
	//
	// Authorization Action:
	// ```
	// ```
	//
	//secrets:secrets:list
	List(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Secrets, error)
}

type secretsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretsServiceClient(cc grpc.ClientConnInterface) SecretsServiceClient {
	return &secretsServiceClient{cc}
}

func (c *secretsServiceClient) Create(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/chef.automate.api.secrets.SecretsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsServiceClient) Read(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, "/chef.automate.api.secrets.SecretsService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsServiceClient) Update(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.secrets.SecretsService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.secrets.SecretsService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsServiceClient) List(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Secrets, error) {
	out := new(Secrets)
	err := c.cc.Invoke(ctx, "/chef.automate.api.secrets.SecretsService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretsServiceServer is the server API for SecretsService service.
type SecretsServiceServer interface {
	// Create a secret
	//
	// Creates a secret. Requires values for name, type, and data.
	//
	// Supported types: ssh, winrm, sudo, aws, azure, gcp, service_now
	// Supported keys by type:
	// ssh: username, password, key
	// winrm: username, password
	// sudo: username, password
	// service_now: username, password
	// aws: AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_SESSION_TOKEN
	// azure: AZURE_CLIENT_ID, AZURE_CLIENT_SECRET, AZURE_TENANT_ID
	// azure: AZURE_SUBSCRIPTION_ID is optional
	// gcp: GOOGLE_CREDENTIALS_JSON
	//
	// Example:
	// ```
	// {
	// "name": "my ssh secret",
	// "type": "ssh",
	// "data": [
	// { "key": "username", "value": "vagrant" },
	// { "key": "password", "value": "vagrant"}
	// ]
	// }
	// ```
	//
	// Authorization Action:
	// ```
	// ```
	//
	//secrets:secrets:create
	Create(context.Context, *Secret) (*Id, error)
	// Read a secret
	//
	// Reads a secret given the ID of the secret.
	// Note that the secret information (password and key values) will not be returned by the API, as a safety measure.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//secrets:secrets:get
	Read(context.Context, *Id) (*Secret, error)
	// Update a secret
	//
	// Updates a secret.
	// This is a PATCH operation, meaning the details sent in will override/replace those stored in the DB.
	// Secret information that is not in the body of the request will persist.
	//
	// Example:
	// ```
	// given a credential with a username and password, a user could update the password by passing in the following body,
	// and the name of the secret as well as the username for the secret be unchanged:
	//
	// {
	// "id": "525c013a-2ab3-4e6f-9005-51bc620e9157",
	// "data": [
	// { "key": "password", "value": "new-value"}
	// ]
	// }
	// ```
	//
	// Authorization Action:
	// ```
	// ```
	//
	//secrets:secrets:update
	Update(context.Context, *Secret) (*UpdateResponse, error)
	// Delete a secret
	//
	// Deletes a secret given the ID of the secret.
	// Note that any nodes that were using the secret will no longer be associated with the deleted secret.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//secrets:secrets:delete
	Delete(context.Context, *Id) (*DeleteResponse, error)
	// List and filter secrets
	//
	// Makes a list of secrets.
	// Supports filtering, pagination, and sorting.
	// Adding a filter narrows the list of secrets to only those that match the filter or filters.
	// Supported filters: type, name
	// Supported sort types: name, type, last modified
	//
	// Example:
	// ```
	// {
	// "sort": "type",
	// "order": "ASC",
	// "filters": [
	// { "key": "type", "values": ["ssh","winrm","sudo"] }
	// ],
	// "page":1,
	// "per_page":100
	// }
	// ```
	//
	// Authorization Action:
	// ```
	// ```
	//
	//secrets:secrets:list
	List(context.Context, *Query) (*Secrets, error)
}

// UnimplementedSecretsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSecretsServiceServer struct {
}

func (*UnimplementedSecretsServiceServer) Create(context.Context, *Secret) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedSecretsServiceServer) Read(context.Context, *Id) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedSecretsServiceServer) Update(context.Context, *Secret) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedSecretsServiceServer) Delete(context.Context, *Id) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedSecretsServiceServer) List(context.Context, *Query) (*Secrets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterSecretsServiceServer(s *grpc.Server, srv SecretsServiceServer) {
	s.RegisterService(&_SecretsService_serviceDesc, srv)
}

func _SecretsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.secrets.SecretsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsServiceServer).Create(ctx, req.(*Secret))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.secrets.SecretsService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsServiceServer).Read(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.secrets.SecretsService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsServiceServer).Update(ctx, req.(*Secret))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.secrets.SecretsService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.secrets.SecretsService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsServiceServer).List(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.secrets.SecretsService",
	HandlerType: (*SecretsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SecretsService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _SecretsService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SecretsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SecretsService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SecretsService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external/secrets/secrets.proto",
}
