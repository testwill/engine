// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/api/instance.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	types "github.com/mesg-foundation/engine/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request's data for the `Get` API.
type GetInstanceRequest struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInstanceRequest) Reset()         { *m = GetInstanceRequest{} }
func (m *GetInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*GetInstanceRequest) ProtoMessage()    {}
func (*GetInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d44b8f4a870f63, []int{0}
}

func (m *GetInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInstanceRequest.Unmarshal(m, b)
}
func (m *GetInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInstanceRequest.Marshal(b, m, deterministic)
}
func (m *GetInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInstanceRequest.Merge(m, src)
}
func (m *GetInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_GetInstanceRequest.Size(m)
}
func (m *GetInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInstanceRequest proto.InternalMessageInfo

func (m *GetInstanceRequest) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// The request's data for the `List` API.
type ListInstancesRequest struct {
	// Filter by Services' hash.
	ServiceHash          []byte   `protobuf:"bytes,1,opt,name=serviceHash,proto3" json:"serviceHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListInstancesRequest) Reset()         { *m = ListInstancesRequest{} }
func (m *ListInstancesRequest) String() string { return proto.CompactTextString(m) }
func (*ListInstancesRequest) ProtoMessage()    {}
func (*ListInstancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d44b8f4a870f63, []int{1}
}

func (m *ListInstancesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInstancesRequest.Unmarshal(m, b)
}
func (m *ListInstancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInstancesRequest.Marshal(b, m, deterministic)
}
func (m *ListInstancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInstancesRequest.Merge(m, src)
}
func (m *ListInstancesRequest) XXX_Size() int {
	return xxx_messageInfo_ListInstancesRequest.Size(m)
}
func (m *ListInstancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInstancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListInstancesRequest proto.InternalMessageInfo

func (m *ListInstancesRequest) GetServiceHash() []byte {
	if m != nil {
		return m.ServiceHash
	}
	return nil
}

// The response's data for the `List` API.
type ListInstancesResponse struct {
	// List of instances that match the request's filters.
	Instances            []*types.Instance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListInstancesResponse) Reset()         { *m = ListInstancesResponse{} }
func (m *ListInstancesResponse) String() string { return proto.CompactTextString(m) }
func (*ListInstancesResponse) ProtoMessage()    {}
func (*ListInstancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d44b8f4a870f63, []int{2}
}

func (m *ListInstancesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListInstancesResponse.Unmarshal(m, b)
}
func (m *ListInstancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListInstancesResponse.Marshal(b, m, deterministic)
}
func (m *ListInstancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListInstancesResponse.Merge(m, src)
}
func (m *ListInstancesResponse) XXX_Size() int {
	return xxx_messageInfo_ListInstancesResponse.Size(m)
}
func (m *ListInstancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListInstancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListInstancesResponse proto.InternalMessageInfo

func (m *ListInstancesResponse) GetInstances() []*types.Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

// The request's data for the `Create` API.
type CreateInstanceRequest struct {
	// Service's hash.
	ServiceHash []byte `protobuf:"bytes,1,opt,name=serviceHash,proto3" json:"serviceHash,omitempty"`
	// Environmental variables to apply to the Instance.
	Env                  []string `protobuf:"bytes,2,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateInstanceRequest) Reset()         { *m = CreateInstanceRequest{} }
func (m *CreateInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateInstanceRequest) ProtoMessage()    {}
func (*CreateInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d44b8f4a870f63, []int{3}
}

func (m *CreateInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInstanceRequest.Unmarshal(m, b)
}
func (m *CreateInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInstanceRequest.Marshal(b, m, deterministic)
}
func (m *CreateInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInstanceRequest.Merge(m, src)
}
func (m *CreateInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateInstanceRequest.Size(m)
}
func (m *CreateInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInstanceRequest proto.InternalMessageInfo

func (m *CreateInstanceRequest) GetServiceHash() []byte {
	if m != nil {
		return m.ServiceHash
	}
	return nil
}

func (m *CreateInstanceRequest) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

// The response's data for the `Create` API.
type CreateInstanceResponse struct {
	// The instance's hash created.
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateInstanceResponse) Reset()         { *m = CreateInstanceResponse{} }
func (m *CreateInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateInstanceResponse) ProtoMessage()    {}
func (*CreateInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d44b8f4a870f63, []int{4}
}

func (m *CreateInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInstanceResponse.Unmarshal(m, b)
}
func (m *CreateInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInstanceResponse.Marshal(b, m, deterministic)
}
func (m *CreateInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInstanceResponse.Merge(m, src)
}
func (m *CreateInstanceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateInstanceResponse.Size(m)
}
func (m *CreateInstanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInstanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInstanceResponse proto.InternalMessageInfo

func (m *CreateInstanceResponse) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// The request's data for the `Delete` API.
type DeleteInstanceRequest struct {
	// Instance's hash
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// If true, any persistent data (volumes) that belongs to the instance and its dependencies will also be deleted.
	DeleteData           bool     `protobuf:"varint,2,opt,name=deleteData,proto3" json:"deleteData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteInstanceRequest) Reset()         { *m = DeleteInstanceRequest{} }
func (m *DeleteInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteInstanceRequest) ProtoMessage()    {}
func (*DeleteInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d44b8f4a870f63, []int{5}
}

func (m *DeleteInstanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteInstanceRequest.Unmarshal(m, b)
}
func (m *DeleteInstanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteInstanceRequest.Marshal(b, m, deterministic)
}
func (m *DeleteInstanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteInstanceRequest.Merge(m, src)
}
func (m *DeleteInstanceRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteInstanceRequest.Size(m)
}
func (m *DeleteInstanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteInstanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteInstanceRequest proto.InternalMessageInfo

func (m *DeleteInstanceRequest) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *DeleteInstanceRequest) GetDeleteData() bool {
	if m != nil {
		return m.DeleteData
	}
	return false
}

// The response's data for the `Delete` API.
type DeleteInstanceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteInstanceResponse) Reset()         { *m = DeleteInstanceResponse{} }
func (m *DeleteInstanceResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteInstanceResponse) ProtoMessage()    {}
func (*DeleteInstanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_71d44b8f4a870f63, []int{6}
}

func (m *DeleteInstanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteInstanceResponse.Unmarshal(m, b)
}
func (m *DeleteInstanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteInstanceResponse.Marshal(b, m, deterministic)
}
func (m *DeleteInstanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteInstanceResponse.Merge(m, src)
}
func (m *DeleteInstanceResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteInstanceResponse.Size(m)
}
func (m *DeleteInstanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteInstanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteInstanceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetInstanceRequest)(nil), "api.GetInstanceRequest")
	proto.RegisterType((*ListInstancesRequest)(nil), "api.ListInstancesRequest")
	proto.RegisterType((*ListInstancesResponse)(nil), "api.ListInstancesResponse")
	proto.RegisterType((*CreateInstanceRequest)(nil), "api.CreateInstanceRequest")
	proto.RegisterType((*CreateInstanceResponse)(nil), "api.CreateInstanceResponse")
	proto.RegisterType((*DeleteInstanceRequest)(nil), "api.DeleteInstanceRequest")
	proto.RegisterType((*DeleteInstanceResponse)(nil), "api.DeleteInstanceResponse")
}

func init() { proto.RegisterFile("protobuf/api/instance.proto", fileDescriptor_71d44b8f4a870f63) }

var fileDescriptor_71d44b8f4a870f63 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xa5, 0x94, 0x10, 0x18, 0x4c, 0x34, 0x13, 0xc1, 0x5a, 0xa2, 0x69, 0xf6, 0xd4, 0x83, 0x96,
	0x88, 0x17, 0x6f, 0x1e, 0x20, 0xa2, 0xc1, 0x53, 0xff, 0xc1, 0x82, 0x63, 0xd8, 0xc4, 0x94, 0xb5,
	0xbb, 0x90, 0xf8, 0x0f, 0xfc, 0xd9, 0x86, 0xed, 0x07, 0xcd, 0xb2, 0x89, 0xb7, 0xcd, 0x9b, 0x37,
	0x6f, 0xde, 0xbc, 0x59, 0x18, 0xcb, 0x7c, 0xab, 0xb7, 0xab, 0xdd, 0xe7, 0x84, 0x4b, 0x31, 0x11,
	0x99, 0xd2, 0x3c, 0x5b, 0x53, 0x62, 0x50, 0xf4, 0xb9, 0x14, 0xe1, 0x4d, 0xcd, 0xd0, 0x3f, 0x92,
	0x94, 0xc5, 0x61, 0x31, 0xe0, 0x82, 0xf4, 0x5b, 0x09, 0xa6, 0xf4, 0xbd, 0x23, 0xa5, 0x11, 0xa1,
	0xb3, 0xe1, 0x6a, 0x13, 0x78, 0x91, 0x17, 0x9f, 0xa5, 0xe6, 0xcd, 0x9e, 0xe0, 0xf2, 0x5d, 0xa8,
	0x9a, 0xaa, 0x2a, 0x6e, 0x04, 0x03, 0x45, 0xf9, 0x5e, 0xac, 0xe9, 0xf5, 0xd8, 0xd2, 0x84, 0xd8,
	0x0b, 0x0c, 0xad, 0x4e, 0x25, 0xb7, 0x99, 0x22, 0xbc, 0x87, 0x7e, 0x65, 0x47, 0x05, 0x5e, 0xe4,
	0xc7, 0x83, 0xe9, 0x79, 0x62, 0x6c, 0x26, 0xb5, 0xa3, 0x23, 0x83, 0x2d, 0x61, 0x38, 0xcb, 0x89,
	0x6b, 0xb2, 0xed, 0xfe, 0x6b, 0x01, 0x2f, 0xc0, 0xa7, 0x6c, 0x1f, 0xb4, 0x23, 0x3f, 0xee, 0xa7,
	0x87, 0x27, 0xbb, 0x83, 0x91, 0x2d, 0x56, 0xba, 0x72, 0x2d, 0xbf, 0x84, 0xe1, 0x9c, 0xbe, 0xe8,
	0x74, 0xb4, 0x83, 0x8c, 0xb7, 0x00, 0x1f, 0x86, 0x3c, 0xe7, 0x9a, 0x07, 0xed, 0xc8, 0x8b, 0x7b,
	0x69, 0x03, 0x61, 0x01, 0x8c, 0x6c, 0xb1, 0x62, 0xf4, 0xf4, 0xb7, 0x0d, 0xbd, 0x0a, 0xc4, 0x07,
	0xf0, 0x17, 0xa4, 0xf1, 0x2a, 0xe1, 0x52, 0x24, 0xa7, 0x47, 0x0a, 0xed, 0xa8, 0x58, 0x0b, 0x9f,
	0xa1, 0x73, 0x48, 0x1a, 0xaf, 0x4d, 0x8f, 0xeb, 0x5c, 0x61, 0xe8, 0x2a, 0x15, 0xe3, 0x59, 0x0b,
	0x67, 0xd0, 0x2d, 0x52, 0xc1, 0x82, 0xe7, 0xcc, 0x3b, 0x1c, 0x3b, 0x6b, 0x4d, 0x91, 0x62, 0xbf,
	0x52, 0xc4, 0x99, 0x5c, 0x29, 0xe2, 0x0e, 0x82, 0xb5, 0x56, 0x5d, 0xf3, 0x3f, 0x1f, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xf7, 0xd0, 0xc5, 0xdf, 0xe2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InstanceClient is the client API for Instance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InstanceClient interface {
	// Get returns an Instance matching the criteria of the request.
	Get(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*types.Instance, error)
	// List returns all Instances matching the criteria of the request.
	List(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
	// Create an Instance from a Service's hash and custom environmental variables.
	// It will return an unique identifier which is used to interact with the Instance.
	Create(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error)
	// Delete an Instance.
	Delete(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*DeleteInstanceResponse, error)
}

type instanceClient struct {
	cc *grpc.ClientConn
}

func NewInstanceClient(cc *grpc.ClientConn) InstanceClient {
	return &instanceClient{cc}
}

func (c *instanceClient) Get(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*types.Instance, error) {
	out := new(types.Instance)
	err := c.cc.Invoke(ctx, "/api.Instance/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) List(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := c.cc.Invoke(ctx, "/api.Instance/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) Create(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error) {
	out := new(CreateInstanceResponse)
	err := c.cc.Invoke(ctx, "/api.Instance/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceClient) Delete(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*DeleteInstanceResponse, error) {
	out := new(DeleteInstanceResponse)
	err := c.cc.Invoke(ctx, "/api.Instance/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstanceServer is the server API for Instance service.
type InstanceServer interface {
	// Get returns an Instance matching the criteria of the request.
	Get(context.Context, *GetInstanceRequest) (*types.Instance, error)
	// List returns all Instances matching the criteria of the request.
	List(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
	// Create an Instance from a Service's hash and custom environmental variables.
	// It will return an unique identifier which is used to interact with the Instance.
	Create(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error)
	// Delete an Instance.
	Delete(context.Context, *DeleteInstanceRequest) (*DeleteInstanceResponse, error)
}

// UnimplementedInstanceServer can be embedded to have forward compatible implementations.
type UnimplementedInstanceServer struct {
}

func (*UnimplementedInstanceServer) Get(ctx context.Context, req *GetInstanceRequest) (*types.Instance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedInstanceServer) List(ctx context.Context, req *ListInstancesRequest) (*ListInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedInstanceServer) Create(ctx context.Context, req *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedInstanceServer) Delete(ctx context.Context, req *DeleteInstanceRequest) (*DeleteInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterInstanceServer(s *grpc.Server, srv InstanceServer) {
	s.RegisterService(&_Instance_serviceDesc, srv)
}

func _Instance_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Instance/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Get(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Instance/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).List(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Instance/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Create(ctx, req.(*CreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Instance_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Instance/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServer).Delete(ctx, req.(*DeleteInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Instance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Instance",
	HandlerType: (*InstanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Instance_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Instance_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Instance_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Instance_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/instance.proto",
}
