// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/mesg-foundation/core/api/core/api.proto

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	github.com/mesg-foundation/core/api/core/api.proto

It has these top-level messages:
	ListenEventRequest
	ExecuteTaskRequest
	ListenResultRequest
	StartServiceRequest
	StopServiceRequest
	EventData
	ExecuteTaskReply
	ResultData
	StartServiceReply
	StopServiceReply
	DeployServiceRequest
	DeployServiceReply
	DeleteServiceRequest
	DeleteServiceReply
*/
package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import service "github.com/mesg-foundation/core/service"

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

type ListenEventRequest struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
	EventKey  string `protobuf:"bytes,2,opt,name=eventKey" json:"eventKey,omitempty"`
}

func (m *ListenEventRequest) Reset()                    { *m = ListenEventRequest{} }
func (m *ListenEventRequest) String() string            { return proto.CompactTextString(m) }
func (*ListenEventRequest) ProtoMessage()               {}
func (*ListenEventRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListenEventRequest) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

func (m *ListenEventRequest) GetEventKey() string {
	if m != nil {
		return m.EventKey
	}
	return ""
}

type ExecuteTaskRequest struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
	TaskKey   string `protobuf:"bytes,2,opt,name=taskKey" json:"taskKey,omitempty"`
	TaskData  string `protobuf:"bytes,3,opt,name=taskData" json:"taskData,omitempty"`
}

func (m *ExecuteTaskRequest) Reset()                    { *m = ExecuteTaskRequest{} }
func (m *ExecuteTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecuteTaskRequest) ProtoMessage()               {}
func (*ExecuteTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExecuteTaskRequest) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

func (m *ExecuteTaskRequest) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *ExecuteTaskRequest) GetTaskData() string {
	if m != nil {
		return m.TaskData
	}
	return ""
}

type ListenResultRequest struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
}

func (m *ListenResultRequest) Reset()                    { *m = ListenResultRequest{} }
func (m *ListenResultRequest) String() string            { return proto.CompactTextString(m) }
func (*ListenResultRequest) ProtoMessage()               {}
func (*ListenResultRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListenResultRequest) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

type StartServiceRequest struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
}

func (m *StartServiceRequest) Reset()                    { *m = StartServiceRequest{} }
func (m *StartServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*StartServiceRequest) ProtoMessage()               {}
func (*StartServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StartServiceRequest) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

type StopServiceRequest struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
}

func (m *StopServiceRequest) Reset()                    { *m = StopServiceRequest{} }
func (m *StopServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*StopServiceRequest) ProtoMessage()               {}
func (*StopServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StopServiceRequest) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

type EventData struct {
	Error     string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	EventKey  string `protobuf:"bytes,2,opt,name=eventKey" json:"eventKey,omitempty"`
	EventData string `protobuf:"bytes,3,opt,name=eventData" json:"eventData,omitempty"`
}

func (m *EventData) Reset()                    { *m = EventData{} }
func (m *EventData) String() string            { return proto.CompactTextString(m) }
func (*EventData) ProtoMessage()               {}
func (*EventData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EventData) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *EventData) GetEventKey() string {
	if m != nil {
		return m.EventKey
	}
	return ""
}

func (m *EventData) GetEventData() string {
	if m != nil {
		return m.EventData
	}
	return ""
}

type ExecuteTaskReply struct {
	ExecutionID string `protobuf:"bytes,1,opt,name=executionID" json:"executionID,omitempty"`
	Error       string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *ExecuteTaskReply) Reset()                    { *m = ExecuteTaskReply{} }
func (m *ExecuteTaskReply) String() string            { return proto.CompactTextString(m) }
func (*ExecuteTaskReply) ProtoMessage()               {}
func (*ExecuteTaskReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ExecuteTaskReply) GetExecutionID() string {
	if m != nil {
		return m.ExecutionID
	}
	return ""
}

func (m *ExecuteTaskReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ResultData struct {
	Error       string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	ExecutionID string `protobuf:"bytes,2,opt,name=executionID" json:"executionID,omitempty"`
	TaskKey     string `protobuf:"bytes,3,opt,name=taskKey" json:"taskKey,omitempty"`
	OutputKey   string `protobuf:"bytes,4,opt,name=outputKey" json:"outputKey,omitempty"`
	OutputData  string `protobuf:"bytes,5,opt,name=outputData" json:"outputData,omitempty"`
}

func (m *ResultData) Reset()                    { *m = ResultData{} }
func (m *ResultData) String() string            { return proto.CompactTextString(m) }
func (*ResultData) ProtoMessage()               {}
func (*ResultData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ResultData) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ResultData) GetExecutionID() string {
	if m != nil {
		return m.ExecutionID
	}
	return ""
}

func (m *ResultData) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *ResultData) GetOutputKey() string {
	if m != nil {
		return m.OutputKey
	}
	return ""
}

func (m *ResultData) GetOutputData() string {
	if m != nil {
		return m.OutputData
	}
	return ""
}

type StartServiceReply struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *StartServiceReply) Reset()                    { *m = StartServiceReply{} }
func (m *StartServiceReply) String() string            { return proto.CompactTextString(m) }
func (*StartServiceReply) ProtoMessage()               {}
func (*StartServiceReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StartServiceReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type StopServiceReply struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *StopServiceReply) Reset()                    { *m = StopServiceReply{} }
func (m *StopServiceReply) String() string            { return proto.CompactTextString(m) }
func (*StopServiceReply) ProtoMessage()               {}
func (*StopServiceReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StopServiceReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DeployServiceRequest struct {
	Service *service.Service `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *DeployServiceRequest) Reset()                    { *m = DeployServiceRequest{} }
func (m *DeployServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*DeployServiceRequest) ProtoMessage()               {}
func (*DeployServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeployServiceRequest) GetService() *service.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type DeployServiceReply struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
}

func (m *DeployServiceReply) Reset()                    { *m = DeployServiceReply{} }
func (m *DeployServiceReply) String() string            { return proto.CompactTextString(m) }
func (*DeployServiceReply) ProtoMessage()               {}
func (*DeployServiceReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeployServiceReply) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

type DeleteServiceRequest struct {
	ServiceID string `protobuf:"bytes,1,opt,name=serviceID" json:"serviceID,omitempty"`
}

func (m *DeleteServiceRequest) Reset()                    { *m = DeleteServiceRequest{} }
func (m *DeleteServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteServiceRequest) ProtoMessage()               {}
func (*DeleteServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *DeleteServiceRequest) GetServiceID() string {
	if m != nil {
		return m.ServiceID
	}
	return ""
}

type DeleteServiceReply struct {
}

func (m *DeleteServiceReply) Reset()                    { *m = DeleteServiceReply{} }
func (m *DeleteServiceReply) String() string            { return proto.CompactTextString(m) }
func (*DeleteServiceReply) ProtoMessage()               {}
func (*DeleteServiceReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*ListenEventRequest)(nil), "api.ListenEventRequest")
	proto.RegisterType((*ExecuteTaskRequest)(nil), "api.ExecuteTaskRequest")
	proto.RegisterType((*ListenResultRequest)(nil), "api.ListenResultRequest")
	proto.RegisterType((*StartServiceRequest)(nil), "api.StartServiceRequest")
	proto.RegisterType((*StopServiceRequest)(nil), "api.StopServiceRequest")
	proto.RegisterType((*EventData)(nil), "api.EventData")
	proto.RegisterType((*ExecuteTaskReply)(nil), "api.ExecuteTaskReply")
	proto.RegisterType((*ResultData)(nil), "api.ResultData")
	proto.RegisterType((*StartServiceReply)(nil), "api.StartServiceReply")
	proto.RegisterType((*StopServiceReply)(nil), "api.StopServiceReply")
	proto.RegisterType((*DeployServiceRequest)(nil), "api.DeployServiceRequest")
	proto.RegisterType((*DeployServiceReply)(nil), "api.DeployServiceReply")
	proto.RegisterType((*DeleteServiceRequest)(nil), "api.DeleteServiceRequest")
	proto.RegisterType((*DeleteServiceReply)(nil), "api.DeleteServiceReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Core service

type CoreClient interface {
	ListenEvent(ctx context.Context, in *ListenEventRequest, opts ...grpc.CallOption) (Core_ListenEventClient, error)
	ExecuteTask(ctx context.Context, in *ExecuteTaskRequest, opts ...grpc.CallOption) (*ExecuteTaskReply, error)
	ListenResult(ctx context.Context, in *ListenResultRequest, opts ...grpc.CallOption) (Core_ListenResultClient, error)
	StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*StartServiceReply, error)
	StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*StopServiceReply, error)
	DeployService(ctx context.Context, in *DeployServiceRequest, opts ...grpc.CallOption) (*DeployServiceReply, error)
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceReply, error)
}

type coreClient struct {
	cc *grpc.ClientConn
}

func NewCoreClient(cc *grpc.ClientConn) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) ListenEvent(ctx context.Context, in *ListenEventRequest, opts ...grpc.CallOption) (Core_ListenEventClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Core_serviceDesc.Streams[0], c.cc, "/api.Core/ListenEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreListenEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Core_ListenEventClient interface {
	Recv() (*EventData, error)
	grpc.ClientStream
}

type coreListenEventClient struct {
	grpc.ClientStream
}

func (x *coreListenEventClient) Recv() (*EventData, error) {
	m := new(EventData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coreClient) ExecuteTask(ctx context.Context, in *ExecuteTaskRequest, opts ...grpc.CallOption) (*ExecuteTaskReply, error) {
	out := new(ExecuteTaskReply)
	err := grpc.Invoke(ctx, "/api.Core/ExecuteTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) ListenResult(ctx context.Context, in *ListenResultRequest, opts ...grpc.CallOption) (Core_ListenResultClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Core_serviceDesc.Streams[1], c.cc, "/api.Core/ListenResult", opts...)
	if err != nil {
		return nil, err
	}
	x := &coreListenResultClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Core_ListenResultClient interface {
	Recv() (*ResultData, error)
	grpc.ClientStream
}

type coreListenResultClient struct {
	grpc.ClientStream
}

func (x *coreListenResultClient) Recv() (*ResultData, error) {
	m := new(ResultData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coreClient) StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*StartServiceReply, error) {
	out := new(StartServiceReply)
	err := grpc.Invoke(ctx, "/api.Core/StartService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*StopServiceReply, error) {
	out := new(StopServiceReply)
	err := grpc.Invoke(ctx, "/api.Core/StopService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) DeployService(ctx context.Context, in *DeployServiceRequest, opts ...grpc.CallOption) (*DeployServiceReply, error) {
	out := new(DeployServiceReply)
	err := grpc.Invoke(ctx, "/api.Core/DeployService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceReply, error) {
	out := new(DeleteServiceReply)
	err := grpc.Invoke(ctx, "/api.Core/DeleteService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Core service

type CoreServer interface {
	ListenEvent(*ListenEventRequest, Core_ListenEventServer) error
	ExecuteTask(context.Context, *ExecuteTaskRequest) (*ExecuteTaskReply, error)
	ListenResult(*ListenResultRequest, Core_ListenResultServer) error
	StartService(context.Context, *StartServiceRequest) (*StartServiceReply, error)
	StopService(context.Context, *StopServiceRequest) (*StopServiceReply, error)
	DeployService(context.Context, *DeployServiceRequest) (*DeployServiceReply, error)
	DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceReply, error)
}

func RegisterCoreServer(s *grpc.Server, srv CoreServer) {
	s.RegisterService(&_Core_serviceDesc, srv)
}

func _Core_ListenEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenEventRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoreServer).ListenEvent(m, &coreListenEventServer{stream})
}

type Core_ListenEventServer interface {
	Send(*EventData) error
	grpc.ServerStream
}

type coreListenEventServer struct {
	grpc.ServerStream
}

func (x *coreListenEventServer) Send(m *EventData) error {
	return x.ServerStream.SendMsg(m)
}

func _Core_ExecuteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).ExecuteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Core/ExecuteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).ExecuteTask(ctx, req.(*ExecuteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_ListenResult_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenResultRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoreServer).ListenResult(m, &coreListenResultServer{stream})
}

type Core_ListenResultServer interface {
	Send(*ResultData) error
	grpc.ServerStream
}

type coreListenResultServer struct {
	grpc.ServerStream
}

func (x *coreListenResultServer) Send(m *ResultData) error {
	return x.ServerStream.SendMsg(m)
}

func _Core_StartService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).StartService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Core/StartService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).StartService(ctx, req.(*StartServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_StopService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).StopService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Core/StopService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).StopService(ctx, req.(*StopServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_DeployService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).DeployService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Core/DeployService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).DeployService(ctx, req.(*DeployServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Core_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Core/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Core_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Core",
	HandlerType: (*CoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteTask",
			Handler:    _Core_ExecuteTask_Handler,
		},
		{
			MethodName: "StartService",
			Handler:    _Core_StartService_Handler,
		},
		{
			MethodName: "StopService",
			Handler:    _Core_StopService_Handler,
		},
		{
			MethodName: "DeployService",
			Handler:    _Core_DeployService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _Core_DeleteService_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenEvent",
			Handler:       _Core_ListenEvent_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenResult",
			Handler:       _Core_ListenResult_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/mesg-foundation/core/api/core/api.proto",
}

func init() { proto.RegisterFile("github.com/mesg-foundation/core/api/core/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0xce, 0xab, 0x2d, 0x19, 0x17, 0x08, 0x4b, 0xa0, 0xc6, 0xaa, 0x50, 0xb5, 0xa7, 0x82, 0x44,
	0x82, 0x52, 0xb8, 0x20, 0xa1, 0x4a, 0x21, 0x39, 0xf0, 0x10, 0x87, 0x84, 0x13, 0x9c, 0x5c, 0x33,
	0xb4, 0x56, 0x5d, 0xef, 0x62, 0xaf, 0x2b, 0xfc, 0x5f, 0xf8, 0x5b, 0xfc, 0x1f, 0xb4, 0xbb, 0x7e,
	0xac, 0x1f, 0x84, 0xe4, 0x64, 0xef, 0x3c, 0xbe, 0xf9, 0x66, 0xbf, 0x99, 0x85, 0xd9, 0xa5, 0x2f,
	0xae, 0x92, 0x8b, 0x89, 0xc7, 0x6e, 0xa6, 0x37, 0x18, 0x5f, 0xbe, 0xf8, 0xc1, 0x92, 0xf0, 0xbb,
	0x2b, 0x7c, 0x16, 0x4e, 0x3d, 0x16, 0xe1, 0xd4, 0xe5, 0x7e, 0xf1, 0x33, 0xe1, 0x11, 0x13, 0x8c,
	0xf4, 0x5d, 0xee, 0x3b, 0xaf, 0xff, 0x97, 0x18, 0x63, 0x74, 0xeb, 0x7b, 0xc5, 0x57, 0xe7, 0xd2,
	0xcf, 0x40, 0x3e, 0xf9, 0xb1, 0xc0, 0x70, 0x79, 0x8b, 0xa1, 0x58, 0xe1, 0xcf, 0x04, 0x63, 0x41,
	0x8e, 0x61, 0x98, 0x85, 0xbd, 0x5f, 0xd8, 0xdd, 0x93, 0xee, 0xe9, 0x70, 0x55, 0x1a, 0x88, 0x03,
	0x77, 0x50, 0x46, 0x7f, 0xc4, 0xd4, 0xee, 0x29, 0x67, 0x71, 0xa6, 0x57, 0x40, 0x96, 0xbf, 0xd0,
	0x4b, 0x04, 0x7e, 0x71, 0xe3, 0xeb, 0xed, 0xf0, 0x6c, 0x38, 0x10, 0x6e, 0x7c, 0x5d, 0xc2, 0xe5,
	0x47, 0x59, 0x49, 0xfe, 0x2e, 0x5c, 0xe1, 0xda, 0x7d, 0x5d, 0x29, 0x3f, 0xd3, 0x33, 0x78, 0xa8,
	0x99, 0xaf, 0x30, 0x4e, 0x82, 0xed, 0xa8, 0xcb, 0xa4, 0xb5, 0x70, 0x23, 0xb1, 0xd6, 0x96, 0xed,
	0x92, 0x66, 0x40, 0xd6, 0x82, 0xf1, 0x9d, 0x72, 0xbe, 0xc1, 0x50, 0xdd, 0xa8, 0xa4, 0x4a, 0xc6,
	0xb0, 0x87, 0x51, 0xc4, 0xa2, 0x2c, 0x4c, 0x1f, 0x36, 0x5d, 0xa3, 0x04, 0xc7, 0x3c, 0x3d, 0xeb,
	0xbc, 0x34, 0xd0, 0x0f, 0x30, 0xaa, 0x5c, 0x32, 0x0f, 0x52, 0x72, 0x02, 0x16, 0x2a, 0x9b, 0xcf,
	0xc2, 0x82, 0x90, 0x69, 0x2a, 0x59, 0xf4, 0x0c, 0x16, 0xf4, 0x77, 0x17, 0x40, 0xdf, 0xe0, 0x06,
	0xaa, 0x35, 0xf0, 0x5e, 0x13, 0xdc, 0xd0, 0xb0, 0x5f, 0xd5, 0xf0, 0x18, 0x86, 0x2c, 0x11, 0x3c,
	0x51, 0x7d, 0x0e, 0x74, 0x2b, 0x85, 0x81, 0x3c, 0x05, 0xd0, 0x07, 0xd5, 0xe9, 0x9e, 0x72, 0x1b,
	0x16, 0xfa, 0x0c, 0x1e, 0x54, 0x05, 0x93, 0xbd, 0xb6, 0x92, 0xa4, 0xa7, 0x30, 0xaa, 0xc8, 0xf4,
	0xef, 0xc8, 0x39, 0x8c, 0x17, 0xc8, 0x03, 0x96, 0xd6, 0x24, 0x7d, 0x0e, 0x07, 0x99, 0x82, 0x2a,
	0xde, 0x9a, 0x8d, 0x26, 0xf9, 0xb6, 0xe4, 0x91, 0x79, 0x80, 0x1c, 0x8a, 0x1a, 0x86, 0xac, 0xb7,
	0x79, 0x28, 0x5e, 0xc9, 0xba, 0x01, 0x0a, 0xdc, 0x69, 0x94, 0xc6, 0xb2, 0x52, 0x25, 0x8b, 0x07,
	0xe9, 0xec, 0x4f, 0x1f, 0x06, 0xef, 0x58, 0x84, 0xe4, 0x0d, 0x58, 0xc6, 0x06, 0x93, 0xa3, 0x89,
	0x7c, 0x18, 0x9a, 0x3b, 0xed, 0xdc, 0x53, 0x8e, 0x62, 0x28, 0x69, 0xe7, 0x65, 0x97, 0x9c, 0x83,
	0x65, 0x0c, 0x52, 0x96, 0xdb, 0xdc, 0x5f, 0xe7, 0x51, 0xd3, 0xc1, 0x83, 0x94, 0x76, 0xc8, 0x5b,
	0x38, 0x34, 0x97, 0x90, 0xd8, 0x46, 0xf5, 0xca, 0x5e, 0x3a, 0xf7, 0x95, 0xa7, 0x9c, 0x34, 0x55,
	0x7f, 0x0e, 0x87, 0xa6, 0xba, 0x59, 0x7a, 0xcb, 0x86, 0x3a, 0x8f, 0x5b, 0x3c, 0x9a, 0xc2, 0x39,
	0x58, 0x86, 0xec, 0x59, 0x0f, 0xcd, 0x7d, 0xcd, 0x7a, 0xa8, 0x4f, 0x08, 0xed, 0x90, 0x25, 0xdc,
	0xad, 0x28, 0x49, 0x9e, 0xa8, 0xc8, 0xb6, 0x09, 0x71, 0x8e, 0xda, 0x5c, 0x06, 0x8c, 0x21, 0x53,
	0x01, 0xd3, 0x14, 0xbc, 0x80, 0xa9, 0xab, 0x4a, 0x3b, 0xf3, 0xfd, 0xaf, 0x03, 0xf9, 0x5c, 0x5f,
	0xec, 0xab, 0xf7, 0xf9, 0xec, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x80, 0x9f, 0xc6, 0x11,
	0x06, 0x00, 0x00,
}
