// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.0
// source: external/iam/v2/common/policy.proto

package common

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

type Type int32

const (
	Type_CHEF_MANAGED Type = 0
	Type_CUSTOM       Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "CHEF_MANAGED",
		1: "CUSTOM",
	}
	Type_value = map[string]int32{
		"CHEF_MANAGED": 0,
		"CUSTOM":       1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_external_iam_v2_common_policy_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_external_iam_v2_common_policy_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_external_iam_v2_common_policy_proto_rawDescGZIP(), []int{0}
}

type Statement_Effect int32

const (
	Statement_ALLOW Statement_Effect = 0
	Statement_DENY  Statement_Effect = 1
)

// Enum value maps for Statement_Effect.
var (
	Statement_Effect_name = map[int32]string{
		0: "ALLOW",
		1: "DENY",
	}
	Statement_Effect_value = map[string]int32{
		"ALLOW": 0,
		"DENY":  1,
	}
)

func (x Statement_Effect) Enum() *Statement_Effect {
	p := new(Statement_Effect)
	*p = x
	return p
}

func (x Statement_Effect) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Statement_Effect) Descriptor() protoreflect.EnumDescriptor {
	return file_external_iam_v2_common_policy_proto_enumTypes[1].Descriptor()
}

func (Statement_Effect) Type() protoreflect.EnumType {
	return &file_external_iam_v2_common_policy_proto_enumTypes[1]
}

func (x Statement_Effect) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Statement_Effect.Descriptor instead.
func (Statement_Effect) EnumDescriptor() ([]byte, []int) {
	return file_external_iam_v2_common_policy_proto_rawDescGZIP(), []int{1, 0}
}

type Version_VersionNumber int32

const (
	Version_V0 Version_VersionNumber = 0
	Version_V1 Version_VersionNumber = 1
	Version_V2 Version_VersionNumber = 2
)

// Enum value maps for Version_VersionNumber.
var (
	Version_VersionNumber_name = map[int32]string{
		0: "V0",
		1: "V1",
		2: "V2",
	}
	Version_VersionNumber_value = map[string]int32{
		"V0": 0,
		"V1": 1,
		"V2": 2,
	}
)

func (x Version_VersionNumber) Enum() *Version_VersionNumber {
	p := new(Version_VersionNumber)
	*p = x
	return p
}

func (x Version_VersionNumber) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Version_VersionNumber) Descriptor() protoreflect.EnumDescriptor {
	return file_external_iam_v2_common_policy_proto_enumTypes[2].Descriptor()
}

func (Version_VersionNumber) Type() protoreflect.EnumType {
	return &file_external_iam_v2_common_policy_proto_enumTypes[2]
}

func (x Version_VersionNumber) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Version_VersionNumber.Descriptor instead.
func (Version_VersionNumber) EnumDescriptor() ([]byte, []int) {
	return file_external_iam_v2_common_policy_proto_rawDescGZIP(), []int{4, 0}
}

type Policy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name for the policy.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Unique ID. Cannot be changed.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// This doc-comment is ignored for an enum.
	Type Type `protobuf:"varint,3,opt,name=type,proto3,enum=chef.automate.api.iam.v2.Type" json:"type,omitempty"`
	// Members affected by this policy. May be empty.
	Members []string `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
	// Statements for the policy. Will contain one or more.
	Statements []*Statement `protobuf:"bytes,5,rep,name=statements,proto3" json:"statements,omitempty"`
	// List of projects this policy belongs to. May be empty.
	Projects      []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Policy) Reset() {
	*x = Policy{}
	mi := &file_external_iam_v2_common_policy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_common_policy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_common_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Policy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Policy) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_CHEF_MANAGED
}

func (x *Policy) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Policy) GetStatements() []*Statement {
	if x != nil {
		return x.Statements
	}
	return nil
}

func (x *Policy) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

type Statement struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This doc-comment is ignored for an enum.
	Effect Statement_Effect `protobuf:"varint,1,opt,name=effect,proto3,enum=chef.automate.api.iam.v2.Statement_Effect" json:"effect,omitempty"`
	// Actions defined inline. May be empty.
	// Best practices recommend that you use custom roles rather than inline actions where practical.
	Actions []string `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	// The role defines a set of actions that the statement is scoped to.
	Role string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	// DEPRECATED: Resources defined inline. Use projects instead.
	Resources []string `protobuf:"bytes,5,rep,name=resources,proto3" json:"resources,omitempty"`
	// The project list defines the set of resources that the statement is scoped to. May be empty.
	Projects      []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Statement) Reset() {
	*x = Statement{}
	mi := &file_external_iam_v2_common_policy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statement) ProtoMessage() {}

func (x *Statement) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_common_policy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statement.ProtoReflect.Descriptor instead.
func (*Statement) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_common_policy_proto_rawDescGZIP(), []int{1}
}

func (x *Statement) GetEffect() Statement_Effect {
	if x != nil {
		return x.Effect
	}
	return Statement_ALLOW
}

func (x *Statement) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Statement) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Statement) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Statement) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

type Role struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name for the role.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Unique ID. Cannot be changed.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Whether this policy is user created (`CUSTOM`) or chef managed (`CHEF_MANAGED`).
	Type Type `protobuf:"varint,3,opt,name=type,proto3,enum=chef.automate.api.iam.v2.Type" json:"type,omitempty"`
	// List of actions this role scopes to. Will contain one or more.
	Actions []string `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
	// List of projects this role belongs to. May be empty.
	Projects      []string `protobuf:"bytes,5,rep,name=projects,proto3" json:"projects,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Role) Reset() {
	*x = Role{}
	mi := &file_external_iam_v2_common_policy_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_common_policy_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_common_policy_proto_rawDescGZIP(), []int{2}
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_CHEF_MANAGED
}

func (x *Role) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Role) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

type Project struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name for the project.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Unique ID. Cannot be changed.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Whether this policy is user created (`CUSTOM`) or chef managed (`CHEF_MANAGED`).
	Type Type `protobuf:"varint,3,opt,name=type,proto3,enum=chef.automate.api.iam.v2.Type" json:"type,omitempty"`
	// The current status of the rules for this project.
	Status        ProjectRulesStatus `protobuf:"varint,4,opt,name=status,proto3,enum=chef.automate.api.iam.v2.ProjectRulesStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Project) Reset() {
	*x = Project{}
	mi := &file_external_iam_v2_common_policy_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_common_policy_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_common_policy_proto_rawDescGZIP(), []int{3}
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_CHEF_MANAGED
}

func (x *Project) GetStatus() ProjectRulesStatus {
	if x != nil {
		return x.Status
	}
	return ProjectRulesStatus_PROJECT_RULES_STATUS_UNSET
}

// the only values that may be returned by GetPolicyVersion
type Version struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Major         Version_VersionNumber  `protobuf:"varint,1,opt,name=major,proto3,enum=chef.automate.api.iam.v2.Version_VersionNumber" json:"major,omitempty"`
	Minor         Version_VersionNumber  `protobuf:"varint,2,opt,name=minor,proto3,enum=chef.automate.api.iam.v2.Version_VersionNumber" json:"minor,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Version) Reset() {
	*x = Version{}
	mi := &file_external_iam_v2_common_policy_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_external_iam_v2_common_policy_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_external_iam_v2_common_policy_proto_rawDescGZIP(), []int{4}
}

func (x *Version) GetMajor() Version_VersionNumber {
	if x != nil {
		return x.Major
	}
	return Version_V0
}

func (x *Version) GetMinor() Version_VersionNumber {
	if x != nil {
		return x.Minor
	}
	return Version_V0
}

var File_external_iam_v2_common_policy_proto protoreflect.FileDescriptor

var file_external_iam_v2_common_policy_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76,
	0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x1a,
	0x22, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x12, 0x43, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x22, 0xd6, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x42, 0x0a, 0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x06, 0x65, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x1d, 0x0a, 0x06, 0x45,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x01, 0x22, 0x94, 0x01, 0x0a, 0x04, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x22, 0xa7, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32,
	0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x07,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x32, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x45,
	0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x05,
	0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x22, 0x27, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x30, 0x10, 0x00, 0x12, 0x06,
	0x0a, 0x02, 0x56, 0x31, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x32, 0x10, 0x02, 0x2a, 0x24,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x48, 0x45, 0x46, 0x5f, 0x4d,
	0x41, 0x4e, 0x41, 0x47, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54,
	0x4f, 0x4d, 0x10, 0x01, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_external_iam_v2_common_policy_proto_rawDescOnce sync.Once
	file_external_iam_v2_common_policy_proto_rawDescData = file_external_iam_v2_common_policy_proto_rawDesc
)

func file_external_iam_v2_common_policy_proto_rawDescGZIP() []byte {
	file_external_iam_v2_common_policy_proto_rawDescOnce.Do(func() {
		file_external_iam_v2_common_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_iam_v2_common_policy_proto_rawDescData)
	})
	return file_external_iam_v2_common_policy_proto_rawDescData
}

var file_external_iam_v2_common_policy_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_external_iam_v2_common_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_external_iam_v2_common_policy_proto_goTypes = []any{
	(Type)(0),                  // 0: chef.automate.api.iam.v2.Type
	(Statement_Effect)(0),      // 1: chef.automate.api.iam.v2.Statement.Effect
	(Version_VersionNumber)(0), // 2: chef.automate.api.iam.v2.Version.VersionNumber
	(*Policy)(nil),             // 3: chef.automate.api.iam.v2.Policy
	(*Statement)(nil),          // 4: chef.automate.api.iam.v2.Statement
	(*Role)(nil),               // 5: chef.automate.api.iam.v2.Role
	(*Project)(nil),            // 6: chef.automate.api.iam.v2.Project
	(*Version)(nil),            // 7: chef.automate.api.iam.v2.Version
	(ProjectRulesStatus)(0),    // 8: chef.automate.api.iam.v2.ProjectRulesStatus
}
var file_external_iam_v2_common_policy_proto_depIdxs = []int32{
	0, // 0: chef.automate.api.iam.v2.Policy.type:type_name -> chef.automate.api.iam.v2.Type
	4, // 1: chef.automate.api.iam.v2.Policy.statements:type_name -> chef.automate.api.iam.v2.Statement
	1, // 2: chef.automate.api.iam.v2.Statement.effect:type_name -> chef.automate.api.iam.v2.Statement.Effect
	0, // 3: chef.automate.api.iam.v2.Role.type:type_name -> chef.automate.api.iam.v2.Type
	0, // 4: chef.automate.api.iam.v2.Project.type:type_name -> chef.automate.api.iam.v2.Type
	8, // 5: chef.automate.api.iam.v2.Project.status:type_name -> chef.automate.api.iam.v2.ProjectRulesStatus
	2, // 6: chef.automate.api.iam.v2.Version.major:type_name -> chef.automate.api.iam.v2.Version.VersionNumber
	2, // 7: chef.automate.api.iam.v2.Version.minor:type_name -> chef.automate.api.iam.v2.Version.VersionNumber
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_external_iam_v2_common_policy_proto_init() }
func file_external_iam_v2_common_policy_proto_init() {
	if File_external_iam_v2_common_policy_proto != nil {
		return
	}
	file_external_iam_v2_common_rules_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_iam_v2_common_policy_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_iam_v2_common_policy_proto_goTypes,
		DependencyIndexes: file_external_iam_v2_common_policy_proto_depIdxs,
		EnumInfos:         file_external_iam_v2_common_policy_proto_enumTypes,
		MessageInfos:      file_external_iam_v2_common_policy_proto_msgTypes,
	}.Build()
	File_external_iam_v2_common_policy_proto = out.File
	file_external_iam_v2_common_policy_proto_rawDesc = nil
	file_external_iam_v2_common_policy_proto_goTypes = nil
	file_external_iam_v2_common_policy_proto_depIdxs = nil
}
