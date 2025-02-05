// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.0
// source: external/report_manager/report_manager.proto

package report_manager

import (
	context "context"
	_ "github.com/chef/automate/api/external/annotations/iam"
	common "github.com/chef/automate/api/external/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type ListDownloadReportRequestsResponse struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Data          []*DownloadRequestResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDownloadReportRequestsResponse) Reset() {
	*x = ListDownloadReportRequestsResponse{}
	mi := &file_external_report_manager_report_manager_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDownloadReportRequestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDownloadReportRequestsResponse) ProtoMessage() {}

func (x *ListDownloadReportRequestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_report_manager_report_manager_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDownloadReportRequestsResponse.ProtoReflect.Descriptor instead.
func (*ListDownloadReportRequestsResponse) Descriptor() ([]byte, []int) {
	return file_external_report_manager_report_manager_proto_rawDescGZIP(), []int{0}
}

func (x *ListDownloadReportRequestsResponse) GetData() []*DownloadRequestResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type DownloadRequestResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	AcknowledgementId string                 `protobuf:"bytes,1,opt,name=acknowledgement_id,json=acknowledgementId,proto3" json:"acknowledgement_id,omitempty"`
	Status            string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	ReportSize        int64                  `protobuf:"varint,3,opt,name=report_size,json=reportSize,proto3" json:"report_size,omitempty"`
	ErrMessage        string                 `protobuf:"bytes,4,opt,name=err_message,json=errMessage,proto3" json:"err_message,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	EndedAt           *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
	Duration          string                 `protobuf:"bytes,7,opt,name=duration,proto3" json:"duration,omitempty"`
	ReportType        string                 `protobuf:"bytes,8,opt,name=report_type,json=reportType,proto3" json:"report_type,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *DownloadRequestResponse) Reset() {
	*x = DownloadRequestResponse{}
	mi := &file_external_report_manager_report_manager_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequestResponse) ProtoMessage() {}

func (x *DownloadRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_report_manager_report_manager_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequestResponse.ProtoReflect.Descriptor instead.
func (*DownloadRequestResponse) Descriptor() ([]byte, []int) {
	return file_external_report_manager_report_manager_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadRequestResponse) GetAcknowledgementId() string {
	if x != nil {
		return x.AcknowledgementId
	}
	return ""
}

func (x *DownloadRequestResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DownloadRequestResponse) GetReportSize() int64 {
	if x != nil {
		return x.ReportSize
	}
	return 0
}

func (x *DownloadRequestResponse) GetErrMessage() string {
	if x != nil {
		return x.ErrMessage
	}
	return ""
}

func (x *DownloadRequestResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DownloadRequestResponse) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

func (x *DownloadRequestResponse) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *DownloadRequestResponse) GetReportType() string {
	if x != nil {
		return x.ReportType
	}
	return ""
}

type ExportFromReportManagerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportFromReportManagerRequest) Reset() {
	*x = ExportFromReportManagerRequest{}
	mi := &file_external_report_manager_report_manager_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportFromReportManagerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportFromReportManagerRequest) ProtoMessage() {}

func (x *ExportFromReportManagerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_report_manager_report_manager_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportFromReportManagerRequest.ProtoReflect.Descriptor instead.
func (*ExportFromReportManagerRequest) Descriptor() ([]byte, []int) {
	return file_external_report_manager_report_manager_proto_rawDescGZIP(), []int{2}
}

func (x *ExportFromReportManagerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_external_report_manager_report_manager_proto protoreflect.FileDescriptor

var file_external_report_manager_report_manager_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x22, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd1, 0x02,
	0x0a, 0x17, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x63, 0x6b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x30, 0x0a, 0x1e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xfc, 0x02, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xdb, 0x01, 0x0a,
	0x1a, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x44, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f, 0x8a, 0xb5, 0x18, 0x35, 0x0a,
	0x16, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x3a,
	0x6c, 0x69, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x30, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x85, 0x01, 0x0a, 0x17, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x40, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_external_report_manager_report_manager_proto_rawDescOnce sync.Once
	file_external_report_manager_report_manager_proto_rawDescData = file_external_report_manager_report_manager_proto_rawDesc
)

func file_external_report_manager_report_manager_proto_rawDescGZIP() []byte {
	file_external_report_manager_report_manager_proto_rawDescOnce.Do(func() {
		file_external_report_manager_report_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_report_manager_report_manager_proto_rawDescData)
	})
	return file_external_report_manager_report_manager_proto_rawDescData
}

var file_external_report_manager_report_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_external_report_manager_report_manager_proto_goTypes = []any{
	(*ListDownloadReportRequestsResponse)(nil), // 0: chef.automate.api.report_manager.ListDownloadReportRequestsResponse
	(*DownloadRequestResponse)(nil),            // 1: chef.automate.api.report_manager.DownloadRequestResponse
	(*ExportFromReportManagerRequest)(nil),     // 2: chef.automate.api.report_manager.ExportFromReportManagerRequest
	(*timestamppb.Timestamp)(nil),              // 3: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                      // 4: google.protobuf.Empty
	(*common.ExportData)(nil),                  // 5: chef.automate.api.common.ExportData
}
var file_external_report_manager_report_manager_proto_depIdxs = []int32{
	1, // 0: chef.automate.api.report_manager.ListDownloadReportRequestsResponse.data:type_name -> chef.automate.api.report_manager.DownloadRequestResponse
	3, // 1: chef.automate.api.report_manager.DownloadRequestResponse.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: chef.automate.api.report_manager.DownloadRequestResponse.ended_at:type_name -> google.protobuf.Timestamp
	4, // 3: chef.automate.api.report_manager.ReportManagerService.ListDownloadReportRequests:input_type -> google.protobuf.Empty
	2, // 4: chef.automate.api.report_manager.ReportManagerService.ExportFromReportManager:input_type -> chef.automate.api.report_manager.ExportFromReportManagerRequest
	0, // 5: chef.automate.api.report_manager.ReportManagerService.ListDownloadReportRequests:output_type -> chef.automate.api.report_manager.ListDownloadReportRequestsResponse
	5, // 6: chef.automate.api.report_manager.ReportManagerService.ExportFromReportManager:output_type -> chef.automate.api.common.ExportData
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_external_report_manager_report_manager_proto_init() }
func file_external_report_manager_report_manager_proto_init() {
	if File_external_report_manager_report_manager_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_report_manager_report_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_report_manager_report_manager_proto_goTypes,
		DependencyIndexes: file_external_report_manager_report_manager_proto_depIdxs,
		MessageInfos:      file_external_report_manager_report_manager_proto_msgTypes,
	}.Build()
	File_external_report_manager_report_manager_proto = out.File
	file_external_report_manager_report_manager_proto_rawDesc = nil
	file_external_report_manager_report_manager_proto_goTypes = nil
	file_external_report_manager_report_manager_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReportManagerServiceClient is the client API for ReportManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReportManagerServiceClient interface {
	// List Download Report Requests
	//
	// Returns the details of the download report requests placed by the user.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//reportmanager:requests:list
	ListDownloadReportRequests(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListDownloadReportRequestsResponse, error)
	ExportFromReportManager(ctx context.Context, in *ExportFromReportManagerRequest, opts ...grpc.CallOption) (ReportManagerService_ExportFromReportManagerClient, error)
}

type reportManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportManagerServiceClient(cc grpc.ClientConnInterface) ReportManagerServiceClient {
	return &reportManagerServiceClient{cc}
}

func (c *reportManagerServiceClient) ListDownloadReportRequests(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListDownloadReportRequestsResponse, error) {
	out := new(ListDownloadReportRequestsResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.report_manager.ReportManagerService/ListDownloadReportRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportManagerServiceClient) ExportFromReportManager(ctx context.Context, in *ExportFromReportManagerRequest, opts ...grpc.CallOption) (ReportManagerService_ExportFromReportManagerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ReportManagerService_serviceDesc.Streams[0], "/chef.automate.api.report_manager.ReportManagerService/ExportFromReportManager", opts...)
	if err != nil {
		return nil, err
	}
	x := &reportManagerServiceExportFromReportManagerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReportManagerService_ExportFromReportManagerClient interface {
	Recv() (*common.ExportData, error)
	grpc.ClientStream
}

type reportManagerServiceExportFromReportManagerClient struct {
	grpc.ClientStream
}

func (x *reportManagerServiceExportFromReportManagerClient) Recv() (*common.ExportData, error) {
	m := new(common.ExportData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReportManagerServiceServer is the server API for ReportManagerService service.
type ReportManagerServiceServer interface {
	// List Download Report Requests
	//
	// Returns the details of the download report requests placed by the user.
	//
	// Authorization Action:
	// ```
	// ```
	//
	//reportmanager:requests:list
	ListDownloadReportRequests(context.Context, *emptypb.Empty) (*ListDownloadReportRequestsResponse, error)
	ExportFromReportManager(*ExportFromReportManagerRequest, ReportManagerService_ExportFromReportManagerServer) error
}

// UnimplementedReportManagerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReportManagerServiceServer struct {
}

func (*UnimplementedReportManagerServiceServer) ListDownloadReportRequests(context.Context, *emptypb.Empty) (*ListDownloadReportRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDownloadReportRequests not implemented")
}
func (*UnimplementedReportManagerServiceServer) ExportFromReportManager(*ExportFromReportManagerRequest, ReportManagerService_ExportFromReportManagerServer) error {
	return status.Errorf(codes.Unimplemented, "method ExportFromReportManager not implemented")
}

func RegisterReportManagerServiceServer(s *grpc.Server, srv ReportManagerServiceServer) {
	s.RegisterService(&_ReportManagerService_serviceDesc, srv)
}

func _ReportManagerService_ListDownloadReportRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportManagerServiceServer).ListDownloadReportRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.report_manager.ReportManagerService/ListDownloadReportRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportManagerServiceServer).ListDownloadReportRequests(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportManagerService_ExportFromReportManager_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExportFromReportManagerRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReportManagerServiceServer).ExportFromReportManager(m, &reportManagerServiceExportFromReportManagerServer{stream})
}

type ReportManagerService_ExportFromReportManagerServer interface {
	Send(*common.ExportData) error
	grpc.ServerStream
}

type reportManagerServiceExportFromReportManagerServer struct {
	grpc.ServerStream
}

func (x *reportManagerServiceExportFromReportManagerServer) Send(m *common.ExportData) error {
	return x.ServerStream.SendMsg(m)
}

var _ReportManagerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.report_manager.ReportManagerService",
	HandlerType: (*ReportManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDownloadReportRequests",
			Handler:    _ReportManagerService_ListDownloadReportRequests_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExportFromReportManager",
			Handler:       _ReportManagerService_ExportFromReportManager_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "external/report_manager/report_manager.proto",
}
