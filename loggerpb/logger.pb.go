// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.3
// source: loggerpb/logger.proto

package loggerpb

import (
	context "context"
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

type Status int32

const (
	Status_OK    Status = 0
	Status_ERROR Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "OK",
		2: "ERROR",
	}
	Status_value = map[string]int32{
		"OK":    0,
		"ERROR": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_loggerpb_logger_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_loggerpb_logger_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_loggerpb_logger_proto_rawDescGZIP(), []int{0}
}

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cases     int32 `protobuf:"varint,1,opt,name=cases,proto3" json:"cases,omitempty"`
	Death     int32 `protobuf:"varint,2,opt,name=death,proto3" json:"death,omitempty"`
	Recovered int32 `protobuf:"varint,3,opt,name=recovered,proto3" json:"recovered,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loggerpb_logger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_loggerpb_logger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_loggerpb_logger_proto_rawDescGZIP(), []int{0}
}

func (x *Report) GetCases() int32 {
	if x != nil {
		return x.Cases
	}
	return 0
}

func (x *Report) GetDeath() int32 {
	if x != nil {
		return x.Death
	}
	return 0
}

func (x *Report) GetRecovered() int32 {
	if x != nil {
		return x.Recovered
	}
	return 0
}

type DataLoggerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string  `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Source    string  `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Report    *Report `protobuf:"bytes,3,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *DataLoggerRequest) Reset() {
	*x = DataLoggerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loggerpb_logger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataLoggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataLoggerRequest) ProtoMessage() {}

func (x *DataLoggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loggerpb_logger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataLoggerRequest.ProtoReflect.Descriptor instead.
func (*DataLoggerRequest) Descriptor() ([]byte, []int) {
	return file_loggerpb_logger_proto_rawDescGZIP(), []int{1}
}

func (x *DataLoggerRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *DataLoggerRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *DataLoggerRequest) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type DataLoggerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summary *Report `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	Status  Status  `protobuf:"varint,2,opt,name=status,proto3,enum=logger.Status" json:"status,omitempty"`
}

func (x *DataLoggerResponse) Reset() {
	*x = DataLoggerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loggerpb_logger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataLoggerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataLoggerResponse) ProtoMessage() {}

func (x *DataLoggerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loggerpb_logger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataLoggerResponse.ProtoReflect.Descriptor instead.
func (*DataLoggerResponse) Descriptor() ([]byte, []int) {
	return file_loggerpb_logger_proto_rawDescGZIP(), []int{2}
}

func (x *DataLoggerResponse) GetSummary() *Report {
	if x != nil {
		return x.Summary
	}
	return nil
}

func (x *DataLoggerResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

var File_loggerpb_logger_proto protoreflect.FileDescriptor

var file_loggerpb_logger_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x22,
	0x52, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x73,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x65, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x64, 0x65, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x65, 0x64, 0x22, 0x71, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x26,
	0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x66, 0x0a, 0x12, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x1b,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x32, 0x53, 0x0a, 0x11, 0x44,
	0x61, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3e, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x19, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loggerpb_logger_proto_rawDescOnce sync.Once
	file_loggerpb_logger_proto_rawDescData = file_loggerpb_logger_proto_rawDesc
)

func file_loggerpb_logger_proto_rawDescGZIP() []byte {
	file_loggerpb_logger_proto_rawDescOnce.Do(func() {
		file_loggerpb_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_loggerpb_logger_proto_rawDescData)
	})
	return file_loggerpb_logger_proto_rawDescData
}

var file_loggerpb_logger_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_loggerpb_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_loggerpb_logger_proto_goTypes = []interface{}{
	(Status)(0),                // 0: logger.Status
	(*Report)(nil),             // 1: logger.Report
	(*DataLoggerRequest)(nil),  // 2: logger.DataLoggerRequest
	(*DataLoggerResponse)(nil), // 3: logger.DataLoggerResponse
}
var file_loggerpb_logger_proto_depIdxs = []int32{
	1, // 0: logger.DataLoggerRequest.report:type_name -> logger.Report
	1, // 1: logger.DataLoggerResponse.summary:type_name -> logger.Report
	0, // 2: logger.DataLoggerResponse.status:type_name -> logger.Status
	2, // 3: logger.DataLoggerService.Log:input_type -> logger.DataLoggerRequest
	3, // 4: logger.DataLoggerService.Log:output_type -> logger.DataLoggerResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_loggerpb_logger_proto_init() }
func file_loggerpb_logger_proto_init() {
	if File_loggerpb_logger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loggerpb_logger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_loggerpb_logger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataLoggerRequest); i {
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
		file_loggerpb_logger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataLoggerResponse); i {
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
			RawDescriptor: file_loggerpb_logger_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loggerpb_logger_proto_goTypes,
		DependencyIndexes: file_loggerpb_logger_proto_depIdxs,
		EnumInfos:         file_loggerpb_logger_proto_enumTypes,
		MessageInfos:      file_loggerpb_logger_proto_msgTypes,
	}.Build()
	File_loggerpb_logger_proto = out.File
	file_loggerpb_logger_proto_rawDesc = nil
	file_loggerpb_logger_proto_goTypes = nil
	file_loggerpb_logger_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataLoggerServiceClient is the client API for DataLoggerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataLoggerServiceClient interface {
	Log(ctx context.Context, in *DataLoggerRequest, opts ...grpc.CallOption) (*DataLoggerResponse, error)
}

type dataLoggerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataLoggerServiceClient(cc grpc.ClientConnInterface) DataLoggerServiceClient {
	return &dataLoggerServiceClient{cc}
}

func (c *dataLoggerServiceClient) Log(ctx context.Context, in *DataLoggerRequest, opts ...grpc.CallOption) (*DataLoggerResponse, error) {
	out := new(DataLoggerResponse)
	err := c.cc.Invoke(ctx, "/logger.DataLoggerService/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataLoggerServiceServer is the server API for DataLoggerService service.
type DataLoggerServiceServer interface {
	Log(context.Context, *DataLoggerRequest) (*DataLoggerResponse, error)
}

// UnimplementedDataLoggerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDataLoggerServiceServer struct {
}

func (*UnimplementedDataLoggerServiceServer) Log(context.Context, *DataLoggerRequest) (*DataLoggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}

func RegisterDataLoggerServiceServer(s *grpc.Server, srv DataLoggerServiceServer) {
	s.RegisterService(&_DataLoggerService_serviceDesc, srv)
}

func _DataLoggerService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataLoggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataLoggerServiceServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logger.DataLoggerService/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataLoggerServiceServer).Log(ctx, req.(*DataLoggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataLoggerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logger.DataLoggerService",
	HandlerType: (*DataLoggerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _DataLoggerService_Log_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loggerpb/logger.proto",
}
