// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.0
// source: interservice/authn/authenticate.proto

package authn

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AuthenticateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	mi := &file_interservice_authn_authenticate_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateRequest) ProtoMessage() {}

func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_authn_authenticate_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_interservice_authn_authenticate_proto_rawDescGZIP(), []int{0}
}

type AuthenticateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This could be either "user:{local,ldap,saml}:<some-id>",
	// "tls:service:<some-id> or "token:<some-id>",
	// depending on the authentication method that was successful.
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty" toml:"subject,omitempty" mapstructure:"subject,omitempty"`
	// Only human users have teams. The teams are provided either by the external
	// IdP (in which case they're extracted from the id_token; TODO), or, for local
	// users, by teams-service (TODO).
	Teams         []string `protobuf:"bytes,2,rep,name=teams,proto3" json:"teams,omitempty" toml:"teams,omitempty" mapstructure:"teams,omitempty"`
	Requestor     string   `protobuf:"bytes,3,opt,name=requestor,proto3" json:"requestor,omitempty" toml:"requestor,omitempty" mapstructure:"requestor,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticateResponse) Reset() {
	*x = AuthenticateResponse{}
	mi := &file_interservice_authn_authenticate_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateResponse) ProtoMessage() {}

func (x *AuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_authn_authenticate_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_interservice_authn_authenticate_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateResponse) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *AuthenticateResponse) GetTeams() []string {
	if x != nil {
		return x.Teams
	}
	return nil
}

func (x *AuthenticateResponse) GetRequestor() string {
	if x != nil {
		return x.Requestor
	}
	return ""
}

var File_interservice_authn_authenticate_proto protoreflect.FileDescriptor

var file_interservice_authn_authenticate_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x64, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65,
	0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x32, 0xad,
	0x01, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x0c, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65,
	0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_authn_authenticate_proto_rawDescOnce sync.Once
	file_interservice_authn_authenticate_proto_rawDescData = file_interservice_authn_authenticate_proto_rawDesc
)

func file_interservice_authn_authenticate_proto_rawDescGZIP() []byte {
	file_interservice_authn_authenticate_proto_rawDescOnce.Do(func() {
		file_interservice_authn_authenticate_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_authn_authenticate_proto_rawDescData)
	})
	return file_interservice_authn_authenticate_proto_rawDescData
}

var file_interservice_authn_authenticate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_interservice_authn_authenticate_proto_goTypes = []any{
	(*AuthenticateRequest)(nil),  // 0: chef.automate.domain.authn.AuthenticateRequest
	(*AuthenticateResponse)(nil), // 1: chef.automate.domain.authn.AuthenticateResponse
}
var file_interservice_authn_authenticate_proto_depIdxs = []int32{
	0, // 0: chef.automate.domain.authn.AuthenticationService.Authenticate:input_type -> chef.automate.domain.authn.AuthenticateRequest
	1, // 1: chef.automate.domain.authn.AuthenticationService.Authenticate:output_type -> chef.automate.domain.authn.AuthenticateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_interservice_authn_authenticate_proto_init() }
func file_interservice_authn_authenticate_proto_init() {
	if File_interservice_authn_authenticate_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interservice_authn_authenticate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interservice_authn_authenticate_proto_goTypes,
		DependencyIndexes: file_interservice_authn_authenticate_proto_depIdxs,
		MessageInfos:      file_interservice_authn_authenticate_proto_msgTypes,
	}.Build()
	File_interservice_authn_authenticate_proto = out.File
	file_interservice_authn_authenticate_proto_rawDesc = nil
	file_interservice_authn_authenticate_proto_goTypes = nil
	file_interservice_authn_authenticate_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	// Authenticate inspects the request's metadata -- for this, an empty argument
	// is just enough. Getting a response means it was authenticated successfully.
	// If the metadata does not contain what is needed to authenticate the
	// request, or the tokens are wrong, the AuthenticationService will return the
	// corresponding error code, with details in the message.
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.AuthenticationService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
type AuthenticationServiceServer interface {
	// Authenticate inspects the request's metadata -- for this, an empty argument
	// is just enough. Getting a response means it was authenticated successfully.
	// If the metadata does not contain what is needed to authenticate the
	// request, or the tokens are wrong, the AuthenticationService will return the
	// corresponding error code, with details in the message.
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
}

// UnimplementedAuthenticationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (*UnimplementedAuthenticationServiceServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.AuthenticationService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.authn.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthenticationService_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interservice/authn/authenticate.proto",
}
