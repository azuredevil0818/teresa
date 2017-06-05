// Code generated by protoc-gen-go.
// source: pkg/protobuf/app/app.proto
// DO NOT EDIT!

/*
Package app is a generated protocol buffer package.

It is generated from these files:
	pkg/protobuf/app/app.proto

It has these top-level messages:
	CreateRequest
	LogsRequest
	LogsResponse
	Empty
*/
package app

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateRequest struct {
	Name        string                   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Team        string                   `protobuf:"bytes,2,opt,name=team" json:"team,omitempty"`
	ProcessType string                   `protobuf:"bytes,3,opt,name=process_type,json=processType" json:"process_type,omitempty"`
	Limits      *CreateRequest_Limits    `protobuf:"bytes,4,opt,name=limits" json:"limits,omitempty"`
	AutoScale   *CreateRequest_AutoScale `protobuf:"bytes,5,opt,name=auto_scale,json=autoScale" json:"auto_scale,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *CreateRequest) GetProcessType() string {
	if m != nil {
		return m.ProcessType
	}
	return ""
}

func (m *CreateRequest) GetLimits() *CreateRequest_Limits {
	if m != nil {
		return m.Limits
	}
	return nil
}

func (m *CreateRequest) GetAutoScale() *CreateRequest_AutoScale {
	if m != nil {
		return m.AutoScale
	}
	return nil
}

type CreateRequest_Limits struct {
	Default        []*CreateRequest_Limits_LimitRangeQuantity `protobuf:"bytes,1,rep,name=default" json:"default,omitempty"`
	DefaultRequest []*CreateRequest_Limits_LimitRangeQuantity `protobuf:"bytes,2,rep,name=default_request,json=defaultRequest" json:"default_request,omitempty"`
}

func (m *CreateRequest_Limits) Reset()                    { *m = CreateRequest_Limits{} }
func (m *CreateRequest_Limits) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest_Limits) ProtoMessage()               {}
func (*CreateRequest_Limits) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CreateRequest_Limits) GetDefault() []*CreateRequest_Limits_LimitRangeQuantity {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *CreateRequest_Limits) GetDefaultRequest() []*CreateRequest_Limits_LimitRangeQuantity {
	if m != nil {
		return m.DefaultRequest
	}
	return nil
}

type CreateRequest_Limits_LimitRangeQuantity struct {
	Quantity string `protobuf:"bytes,1,opt,name=quantity" json:"quantity,omitempty"`
	Resource string `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
}

func (m *CreateRequest_Limits_LimitRangeQuantity) Reset() {
	*m = CreateRequest_Limits_LimitRangeQuantity{}
}
func (m *CreateRequest_Limits_LimitRangeQuantity) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_Limits_LimitRangeQuantity) ProtoMessage()    {}
func (*CreateRequest_Limits_LimitRangeQuantity) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

func (m *CreateRequest_Limits_LimitRangeQuantity) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *CreateRequest_Limits_LimitRangeQuantity) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type CreateRequest_AutoScale struct {
	CpuTargetUtilization int64 `protobuf:"varint,1,opt,name=cpu_target_utilization,json=cpuTargetUtilization" json:"cpu_target_utilization,omitempty"`
	Max                  int64 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
	Min                  int64 `protobuf:"varint,3,opt,name=min" json:"min,omitempty"`
}

func (m *CreateRequest_AutoScale) Reset()                    { *m = CreateRequest_AutoScale{} }
func (m *CreateRequest_AutoScale) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest_AutoScale) ProtoMessage()               {}
func (*CreateRequest_AutoScale) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *CreateRequest_AutoScale) GetCpuTargetUtilization() int64 {
	if m != nil {
		return m.CpuTargetUtilization
	}
	return 0
}

func (m *CreateRequest_AutoScale) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *CreateRequest_AutoScale) GetMin() int64 {
	if m != nil {
		return m.Min
	}
	return 0
}

type LogsRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Lines  int64  `protobuf:"varint,2,opt,name=lines" json:"lines,omitempty"`
	Follow bool   `protobuf:"varint,3,opt,name=follow" json:"follow,omitempty"`
}

func (m *LogsRequest) Reset()                    { *m = LogsRequest{} }
func (m *LogsRequest) String() string            { return proto.CompactTextString(m) }
func (*LogsRequest) ProtoMessage()               {}
func (*LogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LogsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogsRequest) GetLines() int64 {
	if m != nil {
		return m.Lines
	}
	return 0
}

func (m *LogsRequest) GetFollow() bool {
	if m != nil {
		return m.Follow
	}
	return false
}

type LogsResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *LogsResponse) Reset()                    { *m = LogsResponse{} }
func (m *LogsResponse) String() string            { return proto.CompactTextString(m) }
func (*LogsResponse) ProtoMessage()               {}
func (*LogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogsResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*CreateRequest)(nil), "app.CreateRequest")
	proto.RegisterType((*CreateRequest_Limits)(nil), "app.CreateRequest.Limits")
	proto.RegisterType((*CreateRequest_Limits_LimitRangeQuantity)(nil), "app.CreateRequest.Limits.LimitRangeQuantity")
	proto.RegisterType((*CreateRequest_AutoScale)(nil), "app.CreateRequest.AutoScale")
	proto.RegisterType((*LogsRequest)(nil), "app.LogsRequest")
	proto.RegisterType((*LogsResponse)(nil), "app.LogsResponse")
	proto.RegisterType((*Empty)(nil), "app.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for App service

type AppClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error)
	Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (App_LogsClient, error)
}

type appClient struct {
	cc *grpc.ClientConn
}

func NewAppClient(cc *grpc.ClientConn) AppClient {
	return &appClient{cc}
}

func (c *appClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/app.App/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (App_LogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_App_serviceDesc.Streams[0], c.cc, "/app.App/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &appLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type App_LogsClient interface {
	Recv() (*LogsResponse, error)
	grpc.ClientStream
}

type appLogsClient struct {
	grpc.ClientStream
}

func (x *appLogsClient) Recv() (*LogsResponse, error) {
	m := new(LogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for App service

type AppServer interface {
	Create(context.Context, *CreateRequest) (*Empty, error)
	Logs(*LogsRequest, App_LogsServer) error
}

func RegisterAppServer(s *grpc.Server, srv AppServer) {
	s.RegisterService(&_App_serviceDesc, srv)
}

func _App_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AppServer).Logs(m, &appLogsServer{stream})
}

type App_LogsServer interface {
	Send(*LogsResponse) error
	grpc.ServerStream
}

type appLogsServer struct {
	grpc.ServerStream
}

func (x *appLogsServer) Send(m *LogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _App_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _App_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _App_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/protobuf/app/app.proto",
}

func init() { proto.RegisterFile("pkg/protobuf/app/app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0x3a, 0x75, 0x9b, 0x49, 0x81, 0x32, 0xaa, 0x2a, 0x63, 0x71, 0x08, 0x3e, 0xe5,
	0x00, 0x29, 0x04, 0x6e, 0x9c, 0x2a, 0x04, 0xa7, 0x48, 0x88, 0xa5, 0xbd, 0x62, 0x6d, 0xcd, 0x24,
	0x5a, 0x61, 0x7b, 0xb7, 0xde, 0x59, 0xd1, 0xf0, 0x46, 0x3c, 0x20, 0x77, 0xb4, 0xeb, 0x4d, 0x28,
	0x2a, 0x70, 0xe0, 0x60, 0xf9, 0x9f, 0x7f, 0x66, 0xbe, 0xb1, 0x76, 0xd6, 0x50, 0x98, 0x2f, 0xeb,
	0x33, 0xd3, 0x6b, 0xd6, 0x57, 0x6e, 0x75, 0x26, 0x8d, 0xf1, 0xcf, 0x3c, 0x18, 0x98, 0x4a, 0x63,
	0xca, 0xef, 0x23, 0xb8, 0xf7, 0xa6, 0x27, 0xc9, 0x24, 0xe8, 0xda, 0x91, 0x65, 0x44, 0x18, 0x75,
	0xb2, 0xa5, 0x3c, 0x99, 0x26, 0xb3, 0xb1, 0x08, 0xda, 0x7b, 0x4c, 0xb2, 0xcd, 0xf7, 0x06, 0xcf,
	0x6b, 0x7c, 0x02, 0x47, 0xa6, 0xd7, 0x35, 0x59, 0x5b, 0xf1, 0xc6, 0x50, 0x9e, 0x86, 0xdc, 0x24,
	0x7a, 0x17, 0x1b, 0x43, 0xf8, 0x02, 0xb2, 0x46, 0xb5, 0x8a, 0x6d, 0x3e, 0x9a, 0x26, 0xb3, 0xc9,
	0xe2, 0xd1, 0xdc, 0x4f, 0xff, 0x6d, 0xdc, 0x7c, 0x19, 0x0a, 0x44, 0x2c, 0xc4, 0xd7, 0x00, 0xd2,
	0xb1, 0xae, 0x6c, 0x2d, 0x1b, 0xca, 0xf7, 0x43, 0xdb, 0xe3, 0x3f, 0xb4, 0x9d, 0x3b, 0xd6, 0x1f,
	0x7d, 0x8d, 0x18, 0xcb, 0xad, 0x2c, 0x7e, 0x24, 0x90, 0x0d, 0x3c, 0x7c, 0x07, 0x07, 0x9f, 0x69,
	0x25, 0x5d, 0xc3, 0x79, 0x32, 0x4d, 0x67, 0x93, 0xc5, 0xd3, 0xbf, 0xce, 0x1e, 0x5e, 0x42, 0x76,
	0x6b, 0xfa, 0xe0, 0x64, 0xc7, 0x8a, 0x37, 0x62, 0xdb, 0x8c, 0x97, 0xf0, 0x20, 0xca, 0xaa, 0x1f,
	0xba, 0xf2, 0xbd, 0xff, 0xe0, 0xdd, 0x8f, 0x90, 0x58, 0x59, 0x2c, 0x01, 0xef, 0x56, 0x61, 0x01,
	0x87, 0xd7, 0x51, 0xc7, 0xe3, 0xdf, 0xc5, 0x3e, 0xd7, 0x93, 0xd5, 0xae, 0xaf, 0x29, 0xae, 0x61,
	0x17, 0x17, 0x04, 0xe3, 0xdd, 0x79, 0xe0, 0x2b, 0x38, 0xad, 0x8d, 0xab, 0x58, 0xf6, 0x6b, 0xe2,
	0xca, 0xb1, 0x6a, 0xd4, 0x37, 0xc9, 0x4a, 0x77, 0x01, 0x99, 0x8a, 0x93, 0xda, 0xb8, 0x8b, 0x90,
	0xbc, 0xfc, 0x95, 0xc3, 0x63, 0x48, 0x5b, 0x79, 0x13, 0xc8, 0xa9, 0xf0, 0x32, 0x38, 0xaa, 0x0b,
	0x6b, 0xf5, 0x8e, 0xea, 0xca, 0xf7, 0x30, 0x59, 0xea, 0xb5, 0xfd, 0xd7, 0x45, 0x39, 0x81, 0xfd,
	0x46, 0x75, 0x64, 0x23, 0x68, 0x08, 0xf0, 0x14, 0xb2, 0x95, 0x6e, 0x1a, 0xfd, 0x35, 0xd0, 0x0e,
	0x45, 0x8c, 0xca, 0x12, 0x8e, 0x06, 0xa0, 0x35, 0xba, 0xb3, 0xf1, 0x9a, 0xdd, 0xf0, 0x96, 0xe8,
	0x75, 0x79, 0x00, 0xfb, 0x6f, 0x5b, 0xc3, 0x9b, 0xc5, 0x27, 0x48, 0xcf, 0x8d, 0xc1, 0x19, 0x64,
	0xc3, 0xa1, 0x23, 0xde, 0xdd, 0x40, 0x01, 0xc1, 0x0b, 0x0d, 0xf8, 0x0c, 0x46, 0x9e, 0x8e, 0xc7,
	0xc1, 0xbb, 0xf5, 0xe5, 0xc5, 0xc3, 0x5b, 0xce, 0x30, 0xfa, 0x79, 0x72, 0x95, 0x85, 0xbf, 0xe2,
	0xe5, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0xbe, 0xe1, 0x1c, 0x33, 0x03, 0x00, 0x00,
}
