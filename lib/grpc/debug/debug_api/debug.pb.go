// -*- mode: protobuf; indent-tabs-mode: t; c-basic-offset: 8; tab-width: 8 -*-

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.0
// source: grpc/debug/debug_api/debug.proto

package debug_api

import (
	context "context"
	_ "github.com/chef/automate/api/external/annotations/iam"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SetLogLevelRequest_Level int32

const (
	SetLogLevelRequest_UNKNOWN SetLogLevelRequest_Level = 0
	SetLogLevelRequest_DEBUG   SetLogLevelRequest_Level = 1
	SetLogLevelRequest_INFO    SetLogLevelRequest_Level = 2
	SetLogLevelRequest_WARN    SetLogLevelRequest_Level = 3
	SetLogLevelRequest_FATAL   SetLogLevelRequest_Level = 4
)

// Enum value maps for SetLogLevelRequest_Level.
var (
	SetLogLevelRequest_Level_name = map[int32]string{
		0: "UNKNOWN",
		1: "DEBUG",
		2: "INFO",
		3: "WARN",
		4: "FATAL",
	}
	SetLogLevelRequest_Level_value = map[string]int32{
		"UNKNOWN": 0,
		"DEBUG":   1,
		"INFO":    2,
		"WARN":    3,
		"FATAL":   4,
	}
)

func (x SetLogLevelRequest_Level) Enum() *SetLogLevelRequest_Level {
	p := new(SetLogLevelRequest_Level)
	*p = x
	return p
}

func (x SetLogLevelRequest_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetLogLevelRequest_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_debug_debug_api_debug_proto_enumTypes[0].Descriptor()
}

func (SetLogLevelRequest_Level) Type() protoreflect.EnumType {
	return &file_grpc_debug_debug_api_debug_proto_enumTypes[0]
}

func (x SetLogLevelRequest_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetLogLevelRequest_Level.Descriptor instead.
func (SetLogLevelRequest_Level) EnumDescriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{3, 0}
}

type Chunk struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Chunk         []byte                 `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type TraceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Duration      *durationpb.Duration   `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TraceRequest) Reset() {
	*x = TraceRequest{}
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceRequest) ProtoMessage() {}

func (x *TraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceRequest.ProtoReflect.Descriptor instead.
func (*TraceRequest) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{1}
}

func (x *TraceRequest) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type ProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProfileName   string                 `protobuf:"bytes,1,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	SampleRate    int64                  `protobuf:"varint,2,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
	Duration      *durationpb.Duration   `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProfileRequest) Reset() {
	*x = ProfileRequest{}
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileRequest) ProtoMessage() {}

func (x *ProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileRequest.ProtoReflect.Descriptor instead.
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{2}
}

func (x *ProfileRequest) GetProfileName() string {
	if x != nil {
		return x.ProfileName
	}
	return ""
}

func (x *ProfileRequest) GetSampleRate() int64 {
	if x != nil {
		return x.SampleRate
	}
	return 0
}

func (x *ProfileRequest) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type SetLogLevelRequest struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Level         SetLogLevelRequest_Level `protobuf:"varint,1,opt,name=level,proto3,enum=chef.automate.api.debug.SetLogLevelRequest_Level" json:"level,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetLogLevelRequest) Reset() {
	*x = SetLogLevelRequest{}
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetLogLevelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLogLevelRequest) ProtoMessage() {}

func (x *SetLogLevelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLogLevelRequest.ProtoReflect.Descriptor instead.
func (*SetLogLevelRequest) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{3}
}

func (x *SetLogLevelRequest) GetLevel() SetLogLevelRequest_Level {
	if x != nil {
		return x.Level
	}
	return SetLogLevelRequest_UNKNOWN
}

type SetLogLevelResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetLogLevelResponse) Reset() {
	*x = SetLogLevelResponse{}
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetLogLevelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLogLevelResponse) ProtoMessage() {}

func (x *SetLogLevelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLogLevelResponse.ProtoReflect.Descriptor instead.
func (*SetLogLevelResponse) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{4}
}

type VersionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VersionRequest) Reset() {
	*x = VersionRequest{}
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionRequest) ProtoMessage() {}

func (x *VersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionRequest.ProtoReflect.Descriptor instead.
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{5}
}

type VersionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	GitRef        string                 `protobuf:"bytes,2,opt,name=git_ref,json=gitRef,proto3" json:"git_ref,omitempty"`
	GoVersion     string                 `protobuf:"bytes,3,opt,name=go_version,json=goVersion,proto3" json:"go_version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{6}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *VersionResponse) GetGitRef() string {
	if x != nil {
		return x.GitRef
	}
	return ""
}

func (x *VersionResponse) GetGoVersion() string {
	if x != nil {
		return x.GoVersion
	}
	return ""
}

var File_grpc_debug_debug_api_debug_proto protoreflect.FileDescriptor

const file_grpc_debug_debug_api_debug_proto_rawDesc = "" +
	"\n" +
	" grpc/debug/debug_api/debug.proto\x12\x17chef.automate.api.debug\x1a\x1egoogle/protobuf/duration.proto\x1a*external/annotations/iam/annotations.proto\"\x1d\n" +
	"\x05Chunk\x12\x14\n" +
	"\x05chunk\x18\x01 \x01(\fR\x05chunk\"E\n" +
	"\fTraceRequest\x125\n" +
	"\bduration\x18\x01 \x01(\v2\x19.google.protobuf.DurationR\bduration\"\x8b\x01\n" +
	"\x0eProfileRequest\x12!\n" +
	"\fprofile_name\x18\x01 \x01(\tR\vprofileName\x12\x1f\n" +
	"\vsample_rate\x18\x02 \x01(\x03R\n" +
	"sampleRate\x125\n" +
	"\bduration\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\bduration\"\x9d\x01\n" +
	"\x12SetLogLevelRequest\x12G\n" +
	"\x05level\x18\x01 \x01(\x0e21.chef.automate.api.debug.SetLogLevelRequest.LevelR\x05level\">\n" +
	"\x05Level\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\t\n" +
	"\x05DEBUG\x10\x01\x12\b\n" +
	"\x04INFO\x10\x02\x12\b\n" +
	"\x04WARN\x10\x03\x12\t\n" +
	"\x05FATAL\x10\x04\"\x15\n" +
	"\x13SetLogLevelResponse\"\x10\n" +
	"\x0eVersionRequest\"c\n" +
	"\x0fVersionResponse\x12\x18\n" +
	"\aversion\x18\x01 \x01(\tR\aversion\x12\x17\n" +
	"\agit_ref\x18\x02 \x01(\tR\x06gitRef\x12\x1d\n" +
	"\n" +
	"go_version\x18\x03 \x01(\tR\tgoVersion2\xf4\x03\n" +
	"\x05Debug\x12R\n" +
	"\x05Trace\x12%.chef.automate.api.debug.TraceRequest\x1a\x1e.chef.automate.api.debug.Chunk\"\x000\x01\x12V\n" +
	"\aProfile\x12'.chef.automate.api.debug.ProfileRequest\x1a\x1e.chef.automate.api.debug.Chunk\"\x000\x01\x12\xa3\x01\n" +
	"\vSetLogLevel\x12+.chef.automate.api.debug.SetLogLevelRequest\x1a,.chef.automate.api.debug.SetLogLevelResponse\"9\x8a\xb5\x185\n" +
	"\x17system:service:logLevel\x12\x1asystem:serviceLogLevel:set\x12\x98\x01\n" +
	"\n" +
	"GetVersion\x12'.chef.automate.api.debug.VersionRequest\x1a(.chef.automate.api.debug.VersionResponse\"7\x8a\xb5\x183\n" +
	"\x16system:service:version\x12\x19system:serviceVersion:getB3Z1github.com/chef/automate/lib/grpc/debug/debug_apib\x06proto3"

var (
	file_grpc_debug_debug_api_debug_proto_rawDescOnce sync.Once
	file_grpc_debug_debug_api_debug_proto_rawDescData []byte
)

func file_grpc_debug_debug_api_debug_proto_rawDescGZIP() []byte {
	file_grpc_debug_debug_api_debug_proto_rawDescOnce.Do(func() {
		file_grpc_debug_debug_api_debug_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_grpc_debug_debug_api_debug_proto_rawDesc), len(file_grpc_debug_debug_api_debug_proto_rawDesc)))
	})
	return file_grpc_debug_debug_api_debug_proto_rawDescData
}

var file_grpc_debug_debug_api_debug_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_debug_debug_api_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_grpc_debug_debug_api_debug_proto_goTypes = []any{
	(SetLogLevelRequest_Level)(0), // 0: chef.automate.api.debug.SetLogLevelRequest.Level
	(*Chunk)(nil),                 // 1: chef.automate.api.debug.Chunk
	(*TraceRequest)(nil),          // 2: chef.automate.api.debug.TraceRequest
	(*ProfileRequest)(nil),        // 3: chef.automate.api.debug.ProfileRequest
	(*SetLogLevelRequest)(nil),    // 4: chef.automate.api.debug.SetLogLevelRequest
	(*SetLogLevelResponse)(nil),   // 5: chef.automate.api.debug.SetLogLevelResponse
	(*VersionRequest)(nil),        // 6: chef.automate.api.debug.VersionRequest
	(*VersionResponse)(nil),       // 7: chef.automate.api.debug.VersionResponse
	(*durationpb.Duration)(nil),   // 8: google.protobuf.Duration
}
var file_grpc_debug_debug_api_debug_proto_depIdxs = []int32{
	8, // 0: chef.automate.api.debug.TraceRequest.duration:type_name -> google.protobuf.Duration
	8, // 1: chef.automate.api.debug.ProfileRequest.duration:type_name -> google.protobuf.Duration
	0, // 2: chef.automate.api.debug.SetLogLevelRequest.level:type_name -> chef.automate.api.debug.SetLogLevelRequest.Level
	2, // 3: chef.automate.api.debug.Debug.Trace:input_type -> chef.automate.api.debug.TraceRequest
	3, // 4: chef.automate.api.debug.Debug.Profile:input_type -> chef.automate.api.debug.ProfileRequest
	4, // 5: chef.automate.api.debug.Debug.SetLogLevel:input_type -> chef.automate.api.debug.SetLogLevelRequest
	6, // 6: chef.automate.api.debug.Debug.GetVersion:input_type -> chef.automate.api.debug.VersionRequest
	1, // 7: chef.automate.api.debug.Debug.Trace:output_type -> chef.automate.api.debug.Chunk
	1, // 8: chef.automate.api.debug.Debug.Profile:output_type -> chef.automate.api.debug.Chunk
	5, // 9: chef.automate.api.debug.Debug.SetLogLevel:output_type -> chef.automate.api.debug.SetLogLevelResponse
	7, // 10: chef.automate.api.debug.Debug.GetVersion:output_type -> chef.automate.api.debug.VersionResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_debug_debug_api_debug_proto_init() }
func file_grpc_debug_debug_api_debug_proto_init() {
	if File_grpc_debug_debug_api_debug_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_grpc_debug_debug_api_debug_proto_rawDesc), len(file_grpc_debug_debug_api_debug_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_debug_debug_api_debug_proto_goTypes,
		DependencyIndexes: file_grpc_debug_debug_api_debug_proto_depIdxs,
		EnumInfos:         file_grpc_debug_debug_api_debug_proto_enumTypes,
		MessageInfos:      file_grpc_debug_debug_api_debug_proto_msgTypes,
	}.Build()
	File_grpc_debug_debug_api_debug_proto = out.File
	file_grpc_debug_debug_api_debug_proto_goTypes = nil
	file_grpc_debug_debug_api_debug_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DebugClient is the client API for Debug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DebugClient interface {
	Trace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (Debug_TraceClient, error)
	Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (Debug_ProfileClient, error)
	SetLogLevel(ctx context.Context, in *SetLogLevelRequest, opts ...grpc.CallOption) (*SetLogLevelResponse, error)
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type debugClient struct {
	cc grpc.ClientConnInterface
}

func NewDebugClient(cc grpc.ClientConnInterface) DebugClient {
	return &debugClient{cc}
}

func (c *debugClient) Trace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (Debug_TraceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Debug_serviceDesc.Streams[0], "/chef.automate.api.debug.Debug/Trace", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugTraceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_TraceClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type debugTraceClient struct {
	grpc.ClientStream
}

func (x *debugTraceClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugClient) Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (Debug_ProfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Debug_serviceDesc.Streams[1], "/chef.automate.api.debug.Debug/Profile", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugProfileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_ProfileClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type debugProfileClient struct {
	grpc.ClientStream
}

func (x *debugProfileClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugClient) SetLogLevel(ctx context.Context, in *SetLogLevelRequest, opts ...grpc.CallOption) (*SetLogLevelResponse, error) {
	out := new(SetLogLevelResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.debug.Debug/SetLogLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.debug.Debug/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DebugServer is the server API for Debug service.
type DebugServer interface {
	Trace(*TraceRequest, Debug_TraceServer) error
	Profile(*ProfileRequest, Debug_ProfileServer) error
	SetLogLevel(context.Context, *SetLogLevelRequest) (*SetLogLevelResponse, error)
	GetVersion(context.Context, *VersionRequest) (*VersionResponse, error)
}

// UnimplementedDebugServer can be embedded to have forward compatible implementations.
type UnimplementedDebugServer struct {
}

func (*UnimplementedDebugServer) Trace(*TraceRequest, Debug_TraceServer) error {
	return status.Errorf(codes.Unimplemented, "method Trace not implemented")
}
func (*UnimplementedDebugServer) Profile(*ProfileRequest, Debug_ProfileServer) error {
	return status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (*UnimplementedDebugServer) SetLogLevel(context.Context, *SetLogLevelRequest) (*SetLogLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLogLevel not implemented")
}
func (*UnimplementedDebugServer) GetVersion(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}

func RegisterDebugServer(s *grpc.Server, srv DebugServer) {
	s.RegisterService(&_Debug_serviceDesc, srv)
}

func _Debug_Trace_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TraceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Trace(m, &debugTraceServer{stream})
}

type Debug_TraceServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type debugTraceServer struct {
	grpc.ServerStream
}

func (x *debugTraceServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Debug_Profile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProfileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Profile(m, &debugProfileServer{stream})
}

type Debug_ProfileServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type debugProfileServer struct {
	grpc.ServerStream
}

func (x *debugProfileServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Debug_SetLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLogLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).SetLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.debug.Debug/SetLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).SetLogLevel(ctx, req.(*SetLogLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debug_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.debug.Debug/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Debug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.debug.Debug",
	HandlerType: (*DebugServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLogLevel",
			Handler:    _Debug_SetLogLevel_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Debug_GetVersion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Trace",
			Handler:       _Debug_Trace_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Profile",
			Handler:       _Debug_Profile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc/debug/debug_api/debug.proto",
}
