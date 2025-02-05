// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.0
// source: external/iam/v2/request/rules.proto

package request

import (
	common "github.com/chef/automate/api/external/iam/v2/common"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type CreateRuleReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique ID. Cannot be changed.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique ID of the project this rule belongs to. Cannot be changed.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Name for the project rule.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Whether the rule affects nodes (`NODE`) or events (`EVENT`).
	Type common.RuleType `protobuf:"varint,4,opt,name=type,proto3,enum=chef.automate.api.iam.v2.RuleType" json:"type,omitempty"`
	// Conditions that ingested resources must match to belong to the project.
	// Will contain one or more.
	Conditions    []*common.Condition `protobuf:"bytes,5,rep,name=conditions,proto3" json:"conditions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRuleReq) Reset() {
	*x = CreateRuleReq{}
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRuleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRuleReq) ProtoMessage() {}

func (x *CreateRuleReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRuleReq.ProtoReflect.Descriptor instead.
func (*CreateRuleReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_rules_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRuleReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateRuleReq) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateRuleReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRuleReq) GetType() common.RuleType {
	if x != nil {
		return x.Type
	}
	return common.RuleType(0)
}

func (x *CreateRuleReq) GetConditions() []*common.Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

type UpdateRuleReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique ID. Cannot be changed.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique ID of the project this rule belongs to. Cannot be changed.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Name for the project rule.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Whether the rule applies to ingested `NODE` or `EVENT` resources.
	// Cannot be changed.
	Type common.RuleType `protobuf:"varint,4,opt,name=type,proto3,enum=chef.automate.api.iam.v2.RuleType" json:"type,omitempty"`
	// Conditions that ingested resources must match to belong to the project.
	// Will contain one or more.
	Conditions    []*common.Condition `protobuf:"bytes,5,rep,name=conditions,proto3" json:"conditions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRuleReq) Reset() {
	*x = UpdateRuleReq{}
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRuleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRuleReq) ProtoMessage() {}

func (x *UpdateRuleReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRuleReq.ProtoReflect.Descriptor instead.
func (*UpdateRuleReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_rules_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRuleReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRuleReq) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *UpdateRuleReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRuleReq) GetType() common.RuleType {
	if x != nil {
		return x.Type
	}
	return common.RuleType(0)
}

func (x *UpdateRuleReq) GetConditions() []*common.Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

type GetRuleReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the project rule.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the project the rule belongs to.
	ProjectId     string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRuleReq) Reset() {
	*x = GetRuleReq{}
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRuleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuleReq) ProtoMessage() {}

func (x *GetRuleReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuleReq.ProtoReflect.Descriptor instead.
func (*GetRuleReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_rules_proto_rawDescGZIP(), []int{2}
}

func (x *GetRuleReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetRuleReq) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type ListRulesReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRulesReq) Reset() {
	*x = ListRulesReq{}
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRulesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRulesReq) ProtoMessage() {}

func (x *ListRulesReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRulesReq.ProtoReflect.Descriptor instead.
func (*ListRulesReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_rules_proto_rawDescGZIP(), []int{3}
}

type ListRulesForProjectReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the project.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRulesForProjectReq) Reset() {
	*x = ListRulesForProjectReq{}
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRulesForProjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRulesForProjectReq) ProtoMessage() {}

func (x *ListRulesForProjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRulesForProjectReq.ProtoReflect.Descriptor instead.
func (*ListRulesForProjectReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_rules_proto_rawDescGZIP(), []int{4}
}

func (x *ListRulesForProjectReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteRuleReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the project rule.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the project the rule belongs to.
	ProjectId     string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRuleReq) Reset() {
	*x = DeleteRuleReq{}
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRuleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRuleReq) ProtoMessage() {}

func (x *DeleteRuleReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRuleReq.ProtoReflect.Descriptor instead.
func (*DeleteRuleReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_rules_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRuleReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteRuleReq) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type ApplyRulesStartReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApplyRulesStartReq) Reset() {
	*x = ApplyRulesStartReq{}
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApplyRulesStartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyRulesStartReq) ProtoMessage() {}

func (x *ApplyRulesStartReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyRulesStartReq.ProtoReflect.Descriptor instead.
func (*ApplyRulesStartReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_rules_proto_rawDescGZIP(), []int{6}
}

type ApplyRulesCancelReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApplyRulesCancelReq) Reset() {
	*x = ApplyRulesCancelReq{}
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApplyRulesCancelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyRulesCancelReq) ProtoMessage() {}

func (x *ApplyRulesCancelReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyRulesCancelReq.ProtoReflect.Descriptor instead.
func (*ApplyRulesCancelReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_rules_proto_rawDescGZIP(), []int{7}
}

type ApplyRulesStatusReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApplyRulesStatusReq) Reset() {
	*x = ApplyRulesStatusReq{}
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApplyRulesStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyRulesStatusReq) ProtoMessage() {}

func (x *ApplyRulesStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_request_rules_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyRulesStatusReq.ProtoReflect.Descriptor instead.
func (*ApplyRulesStatusReq) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_request_rules_proto_rawDescGZIP(), []int{8}
}

var File_external_iam_v2_request_rules_proto protoreflect.FileDescriptor

var file_external_iam_v2_request_rules_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76,
	0x32, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x1a,
	0x22, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdc, 0x03, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e,
	0x52, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x43,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x3a, 0x8a, 0x02, 0x92, 0x41, 0x86, 0x02, 0x0a, 0x2d, 0xd2, 0x01, 0x02, 0x69,
	0x64, 0xd2, 0x01, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0xd2, 0x01,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x04, 0x74, 0x79, 0x70, 0x65, 0xd2, 0x01, 0x0a, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xd4, 0x01, 0x12, 0xd1, 0x01, 0x7b,
	0x22, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x72,
	0x75, 0x6c, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x69,
	0x64, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22,
	0x4d, 0x79, 0x20, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x52, 0x75, 0x6c, 0x65, 0x22,
	0x2c, 0x20, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4e, 0x4f, 0x44, 0x45, 0x22,
	0x2c, 0x20, 0x22, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3a, 0x20,
	0x5b, 0x7b, 0x22, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x3a, 0x20, 0x22,
	0x43, 0x48, 0x45, 0x46, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x22, 0x2c, 0x20, 0x22, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x4d, 0x45, 0x4d, 0x42, 0x45,
	0x52, 0x5f, 0x4f, 0x46, 0x22, 0x2c, 0x20, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3a,
	0x20, 0x5b, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x22, 0x2c, 0x20,
	0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6f, 0x22, 0x5d, 0x7d, 0x5d, 0x7d,
	0x22, 0xcb, 0x03, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x52,
	0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x43, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x3a, 0xf9, 0x01, 0x92, 0x41, 0xf5, 0x01, 0x0a, 0x2d, 0xd2, 0x01, 0x02, 0x69, 0x64,
	0xd2, 0x01, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0xd2, 0x01, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x04, 0x74, 0x79, 0x70, 0x65, 0xd2, 0x01, 0x0a, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xc3, 0x01, 0x12, 0xc0, 0x01, 0x7b, 0x22,
	0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x72, 0x75,
	0x6c, 0x65, 0x22, 0x2c, 0x20, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x69, 0x64,
	0x22, 0x3a, 0x20, 0x22, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x22, 0x2c, 0x20, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4d,
	0x79, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x20, 0x52, 0x75, 0x6c, 0x65, 0x22, 0x2c,
	0x20, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4e, 0x4f, 0x44, 0x45, 0x22, 0x2c,
	0x20, 0x22, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3a, 0x20, 0x5b,
	0x7b, 0x22, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x43,
	0x48, 0x45, 0x46, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x22, 0x2c, 0x20, 0x22, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x3a, 0x20, 0x22, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53,
	0x22, 0x2c, 0x20, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x22, 0x5d, 0x7d, 0x5d, 0x7d, 0x22, 0x54,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x3a, 0x17, 0x92, 0x41, 0x14,
	0x0a, 0x12, 0xd2, 0x01, 0x02, 0x69, 0x64, 0xd2, 0x01, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x22, 0x34, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x46, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x0a,
	0x92, 0x41, 0x07, 0x0a, 0x05, 0xd2, 0x01, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x3a, 0x17, 0x92, 0x41, 0x14, 0x0a,
	0x12, 0xd2, 0x01, 0x02, 0x69, 0x64, 0xd2, 0x01, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x22, 0x15, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_iam_v2_request_rules_proto_rawDescOnce sync.Once
	file_external_iam_v2_request_rules_proto_rawDescData = file_external_iam_v2_request_rules_proto_rawDesc
)

func file_external_iam_v2_request_rules_proto_rawDescGZIP() []byte {
	file_external_iam_v2_request_rules_proto_rawDescOnce.Do(func() {
		file_external_iam_v2_request_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_iam_v2_request_rules_proto_rawDescData)
	})
	return file_external_iam_v2_request_rules_proto_rawDescData
}

var file_external_iam_v2_request_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_external_iam_v2_request_rules_proto_goTypes = []any{
	(*CreateRuleReq)(nil),          // 0: chef.automate.api.iam.v2.CreateRuleReq
	(*UpdateRuleReq)(nil),          // 1: chef.automate.api.iam.v2.UpdateRuleReq
	(*GetRuleReq)(nil),             // 2: chef.automate.api.iam.v2.GetRuleReq
	(*ListRulesReq)(nil),           // 3: chef.automate.api.iam.v2.ListRulesReq
	(*ListRulesForProjectReq)(nil), // 4: chef.automate.api.iam.v2.ListRulesForProjectReq
	(*DeleteRuleReq)(nil),          // 5: chef.automate.api.iam.v2.DeleteRuleReq
	(*ApplyRulesStartReq)(nil),     // 6: chef.automate.api.iam.v2.ApplyRulesStartReq
	(*ApplyRulesCancelReq)(nil),    // 7: chef.automate.api.iam.v2.ApplyRulesCancelReq
	(*ApplyRulesStatusReq)(nil),    // 8: chef.automate.api.iam.v2.ApplyRulesStatusReq
	(common.RuleType)(0),           // 9: chef.automate.api.iam.v2.RuleType
	(*common.Condition)(nil),       // 10: chef.automate.api.iam.v2.Condition
}
var file_external_iam_v2_request_rules_proto_depIdxs = []int32{
	9,  // 0: chef.automate.api.iam.v2.CreateRuleReq.type:type_name -> chef.automate.api.iam.v2.RuleType
	10, // 1: chef.automate.api.iam.v2.CreateRuleReq.conditions:type_name -> chef.automate.api.iam.v2.Condition
	9,  // 2: chef.automate.api.iam.v2.UpdateRuleReq.type:type_name -> chef.automate.api.iam.v2.RuleType
	10, // 3: chef.automate.api.iam.v2.UpdateRuleReq.conditions:type_name -> chef.automate.api.iam.v2.Condition
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_external_iam_v2_request_rules_proto_init() }
func file_external_iam_v2_request_rules_proto_init() {
	if File_external_iam_v2_request_rules_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_iam_v2_request_rules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_iam_v2_request_rules_proto_goTypes,
		DependencyIndexes: file_external_iam_v2_request_rules_proto_depIdxs,
		MessageInfos:      file_external_iam_v2_request_rules_proto_msgTypes,
	}.Build()
	File_external_iam_v2_request_rules_proto = out.File
	file_external_iam_v2_request_rules_proto_rawDesc = nil
	file_external_iam_v2_request_rules_proto_goTypes = nil
	file_external_iam_v2_request_rules_proto_depIdxs = nil
}
