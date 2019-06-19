// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/api/service.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import definition "github.com/mesg-foundation/core/protobuf/definition"

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

type CreateServiceRequest struct {
	Definition           *definition.Service `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreateServiceRequest) Reset()         { *m = CreateServiceRequest{} }
func (m *CreateServiceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServiceRequest) ProtoMessage()    {}
func (*CreateServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_c439b291bb00d662, []int{0}
}
func (m *CreateServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceRequest.Unmarshal(m, b)
}
func (m *CreateServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceRequest.Marshal(b, m, deterministic)
}
func (dst *CreateServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceRequest.Merge(dst, src)
}
func (m *CreateServiceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateServiceRequest.Size(m)
}
func (m *CreateServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceRequest proto.InternalMessageInfo

func (m *CreateServiceRequest) GetDefinition() *definition.Service {
	if m != nil {
		return m.Definition
	}
	return nil
}

type CreateServiceResponse struct {
	Sid                  string   `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceResponse) Reset()         { *m = CreateServiceResponse{} }
func (m *CreateServiceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateServiceResponse) ProtoMessage()    {}
func (*CreateServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_c439b291bb00d662, []int{1}
}
func (m *CreateServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceResponse.Unmarshal(m, b)
}
func (m *CreateServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceResponse.Marshal(b, m, deterministic)
}
func (dst *CreateServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceResponse.Merge(dst, src)
}
func (m *CreateServiceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateServiceResponse.Size(m)
}
func (m *CreateServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceResponse proto.InternalMessageInfo

func (m *CreateServiceResponse) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *CreateServiceResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// GetServiceRequest defines request to retrieve a single service.
type GetServiceRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceRequest) Reset()         { *m = GetServiceRequest{} }
func (m *GetServiceRequest) String() string { return proto.CompactTextString(m) }
func (*GetServiceRequest) ProtoMessage()    {}
func (*GetServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_c439b291bb00d662, []int{2}
}
func (m *GetServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceRequest.Unmarshal(m, b)
}
func (m *GetServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceRequest.Marshal(b, m, deterministic)
}
func (dst *GetServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceRequest.Merge(dst, src)
}
func (m *GetServiceRequest) XXX_Size() int {
	return xxx_messageInfo_GetServiceRequest.Size(m)
}
func (m *GetServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceRequest proto.InternalMessageInfo

func (m *GetServiceRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateServiceRequest)(nil), "api.CreateServiceRequest")
	proto.RegisterType((*CreateServiceResponse)(nil), "api.CreateServiceResponse")
	proto.RegisterType((*GetServiceRequest)(nil), "api.GetServiceRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceXClient is the client API for ServiceX service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceXClient interface {
	Create(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error)
	// Get returns a single Service specified in a request.
	Get(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*definition.Service, error)
}

type serviceXClient struct {
	cc *grpc.ClientConn
}

func NewServiceXClient(cc *grpc.ClientConn) ServiceXClient {
	return &serviceXClient{cc}
}

func (c *serviceXClient) Create(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*CreateServiceResponse, error) {
	out := new(CreateServiceResponse)
	err := c.cc.Invoke(ctx, "/api.ServiceX/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceXClient) Get(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*definition.Service, error) {
	out := new(definition.Service)
	err := c.cc.Invoke(ctx, "/api.ServiceX/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceXServer is the server API for ServiceX service.
type ServiceXServer interface {
	Create(context.Context, *CreateServiceRequest) (*CreateServiceResponse, error)
	// Get returns a single Service specified in a request.
	Get(context.Context, *GetServiceRequest) (*definition.Service, error)
}

func RegisterServiceXServer(s *grpc.Server, srv ServiceXServer) {
	s.RegisterService(&_ServiceX_serviceDesc, srv)
}

func _ServiceX_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceXServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceX/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceXServer).Create(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceX_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceXServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceX/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceXServer).Get(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceX_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ServiceX",
	HandlerType: (*ServiceXServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ServiceX_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ServiceX_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/service.proto",
}

func init() { proto.RegisterFile("protobuf/api/service.proto", fileDescriptor_service_c439b291bb00d662) }

var fileDescriptor_service_c439b291bb00d662 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x52, 0x85, 0x30,
	0x10, 0x85, 0x6f, 0xc4, 0xb9, 0xa3, 0x6b, 0xa3, 0xeb, 0xcf, 0x5c, 0x53, 0x69, 0x1a, 0xad, 0xc2,
	0x0c, 0xd8, 0x5a, 0x38, 0x16, 0x14, 0x76, 0xd8, 0xd8, 0x06, 0x59, 0x86, 0x6d, 0x48, 0x24, 0xc1,
	0x17, 0xf0, 0xc5, 0x1d, 0x83, 0x88, 0x22, 0xdd, 0xce, 0x97, 0x2f, 0x9b, 0x73, 0x02, 0xd2, 0xf5,
	0x36, 0xd8, 0x6a, 0x68, 0x52, 0xe3, 0x38, 0xf5, 0xd4, 0xbf, 0xf3, 0x2b, 0xe9, 0x08, 0x31, 0x31,
	0x8e, 0xe5, 0xf5, 0x8f, 0x50, 0x53, 0xc3, 0x1d, 0x07, 0xb6, 0xdd, 0x5f, 0x4f, 0x3d, 0xc1, 0xd9,
	0x63, 0x4f, 0x26, 0xd0, 0xf3, 0x88, 0x4b, 0x7a, 0x1b, 0xc8, 0x07, 0xcc, 0x01, 0xe6, 0x3b, 0x3b,
	0x71, 0x25, 0x6e, 0x8f, 0xb2, 0x53, 0x3d, 0x23, 0x3d, 0xf9, 0xbf, 0x34, 0x75, 0x0f, 0xe7, 0x8b,
	0x65, 0xde, 0xd9, 0xce, 0x13, 0x1e, 0x43, 0xe2, 0xb9, 0x8e, 0x6b, 0x0e, 0xcb, 0xaf, 0x11, 0x11,
	0xf6, 0x5b, 0xe3, 0xdb, 0xdd, 0x5e, 0x44, 0x71, 0x56, 0x37, 0x70, 0x52, 0x50, 0x58, 0x04, 0x99,
	0x44, 0x31, 0x8b, 0xd9, 0x87, 0x80, 0x83, 0x6f, 0xed, 0x05, 0x1f, 0x60, 0x3b, 0x3e, 0x8a, 0x97,
	0xda, 0x38, 0xd6, 0x6b, 0x75, 0xa4, 0x5c, 0x3b, 0x1a, 0xc3, 0xa9, 0x0d, 0xde, 0x41, 0x52, 0x50,
	0xc0, 0x8b, 0x28, 0xfd, 0x8b, 0x20, 0xd7, 0x7a, 0xab, 0x4d, 0xb5, 0x8d, 0x3f, 0x98, 0x7f, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0x95, 0x61, 0x97, 0x87, 0x01, 0x00, 0x00,
}
