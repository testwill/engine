// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/api/process.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_mesg_foundation_engine_hash "github.com/mesg-foundation/engine/hash"
	process "github.com/mesg-foundation/engine/process"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// The request's data for the `Create` API.
type CreateProcessRequest struct {
	// Process's name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// List of nodes of the process.
	Nodes []*process.Process_Node `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// List of edges of the process.
	Edges                []*process.Process_Edge `protobuf:"bytes,5,rep,name=edges,proto3" json:"edges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CreateProcessRequest) Reset()         { *m = CreateProcessRequest{} }
func (m *CreateProcessRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProcessRequest) ProtoMessage()    {}
func (*CreateProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e79ddb1ea8859998, []int{0}
}
func (m *CreateProcessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProcessRequest.Unmarshal(m, b)
}
func (m *CreateProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProcessRequest.Marshal(b, m, deterministic)
}
func (m *CreateProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProcessRequest.Merge(m, src)
}
func (m *CreateProcessRequest) XXX_Size() int {
	return xxx_messageInfo_CreateProcessRequest.Size(m)
}
func (m *CreateProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProcessRequest proto.InternalMessageInfo

func (m *CreateProcessRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateProcessRequest) GetNodes() []*process.Process_Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *CreateProcessRequest) GetEdges() []*process.Process_Edge {
	if m != nil {
		return m.Edges
	}
	return nil
}

// The response's data for the `Create` API.
type CreateProcessResponse struct {
	// The process's hash created.
	Hash                 github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *CreateProcessResponse) Reset()         { *m = CreateProcessResponse{} }
func (m *CreateProcessResponse) String() string { return proto.CompactTextString(m) }
func (*CreateProcessResponse) ProtoMessage()    {}
func (*CreateProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e79ddb1ea8859998, []int{1}
}
func (m *CreateProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProcessResponse.Unmarshal(m, b)
}
func (m *CreateProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProcessResponse.Marshal(b, m, deterministic)
}
func (m *CreateProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProcessResponse.Merge(m, src)
}
func (m *CreateProcessResponse) XXX_Size() int {
	return xxx_messageInfo_CreateProcessResponse.Size(m)
}
func (m *CreateProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProcessResponse proto.InternalMessageInfo

// The request's data for the `Delete` API.
type DeleteProcessRequest struct {
	// The process's hash to delete.
	Hash                 github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *DeleteProcessRequest) Reset()         { *m = DeleteProcessRequest{} }
func (m *DeleteProcessRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProcessRequest) ProtoMessage()    {}
func (*DeleteProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e79ddb1ea8859998, []int{2}
}
func (m *DeleteProcessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProcessRequest.Unmarshal(m, b)
}
func (m *DeleteProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProcessRequest.Marshal(b, m, deterministic)
}
func (m *DeleteProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProcessRequest.Merge(m, src)
}
func (m *DeleteProcessRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProcessRequest.Size(m)
}
func (m *DeleteProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProcessRequest proto.InternalMessageInfo

// The response's data for the `Delete` API, doesn't contain anything.
type DeleteProcessResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProcessResponse) Reset()         { *m = DeleteProcessResponse{} }
func (m *DeleteProcessResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteProcessResponse) ProtoMessage()    {}
func (*DeleteProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e79ddb1ea8859998, []int{3}
}
func (m *DeleteProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProcessResponse.Unmarshal(m, b)
}
func (m *DeleteProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProcessResponse.Marshal(b, m, deterministic)
}
func (m *DeleteProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProcessResponse.Merge(m, src)
}
func (m *DeleteProcessResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteProcessResponse.Size(m)
}
func (m *DeleteProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProcessResponse proto.InternalMessageInfo

// The request's data for the `Get` API.
type GetProcessRequest struct {
	// The process's hash to fetch.
	Hash                 github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *GetProcessRequest) Reset()         { *m = GetProcessRequest{} }
func (m *GetProcessRequest) String() string { return proto.CompactTextString(m) }
func (*GetProcessRequest) ProtoMessage()    {}
func (*GetProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e79ddb1ea8859998, []int{4}
}
func (m *GetProcessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProcessRequest.Unmarshal(m, b)
}
func (m *GetProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProcessRequest.Marshal(b, m, deterministic)
}
func (m *GetProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProcessRequest.Merge(m, src)
}
func (m *GetProcessRequest) XXX_Size() int {
	return xxx_messageInfo_GetProcessRequest.Size(m)
}
func (m *GetProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProcessRequest proto.InternalMessageInfo

// The request's data for the `List` API.
type ListProcessRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProcessRequest) Reset()         { *m = ListProcessRequest{} }
func (m *ListProcessRequest) String() string { return proto.CompactTextString(m) }
func (*ListProcessRequest) ProtoMessage()    {}
func (*ListProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e79ddb1ea8859998, []int{5}
}
func (m *ListProcessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProcessRequest.Unmarshal(m, b)
}
func (m *ListProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProcessRequest.Marshal(b, m, deterministic)
}
func (m *ListProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProcessRequest.Merge(m, src)
}
func (m *ListProcessRequest) XXX_Size() int {
	return xxx_messageInfo_ListProcessRequest.Size(m)
}
func (m *ListProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListProcessRequest proto.InternalMessageInfo

// The response's data for the `List` API.
type ListProcessResponse struct {
	// List of processes that match the request's filters.
	Processes            []*process.Process `protobuf:"bytes,1,rep,name=processes,proto3" json:"processes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListProcessResponse) Reset()         { *m = ListProcessResponse{} }
func (m *ListProcessResponse) String() string { return proto.CompactTextString(m) }
func (*ListProcessResponse) ProtoMessage()    {}
func (*ListProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e79ddb1ea8859998, []int{6}
}
func (m *ListProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProcessResponse.Unmarshal(m, b)
}
func (m *ListProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProcessResponse.Marshal(b, m, deterministic)
}
func (m *ListProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProcessResponse.Merge(m, src)
}
func (m *ListProcessResponse) XXX_Size() int {
	return xxx_messageInfo_ListProcessResponse.Size(m)
}
func (m *ListProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListProcessResponse proto.InternalMessageInfo

func (m *ListProcessResponse) GetProcesses() []*process.Process {
	if m != nil {
		return m.Processes
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateProcessRequest)(nil), "mesg.api.CreateProcessRequest")
	proto.RegisterType((*CreateProcessResponse)(nil), "mesg.api.CreateProcessResponse")
	proto.RegisterType((*DeleteProcessRequest)(nil), "mesg.api.DeleteProcessRequest")
	proto.RegisterType((*DeleteProcessResponse)(nil), "mesg.api.DeleteProcessResponse")
	proto.RegisterType((*GetProcessRequest)(nil), "mesg.api.GetProcessRequest")
	proto.RegisterType((*ListProcessRequest)(nil), "mesg.api.ListProcessRequest")
	proto.RegisterType((*ListProcessResponse)(nil), "mesg.api.ListProcessResponse")
}

func init() { proto.RegisterFile("protobuf/api/process.proto", fileDescriptor_e79ddb1ea8859998) }

var fileDescriptor_e79ddb1ea8859998 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0x6e, 0xe2, 0x40,
	0x0c, 0x4e, 0x20, 0xb0, 0x8b, 0x77, 0x2f, 0x3b, 0x80, 0x36, 0xca, 0xb2, 0x0b, 0xca, 0x09, 0x69,
	0xd5, 0x89, 0x0a, 0xa7, 0x5e, 0x69, 0x11, 0x48, 0xad, 0xaa, 0x2a, 0xc7, 0xaa, 0x52, 0x3b, 0x10,
	0x93, 0x44, 0x2a, 0x99, 0x94, 0x99, 0x1c, 0xfa, 0x0a, 0x7d, 0x8f, 0xbe, 0x4b, 0x9f, 0xa1, 0x07,
	0x9e, 0xa5, 0xca, 0x24, 0xfc, 0x87, 0x5b, 0x7b, 0x1b, 0xdb, 0x9f, 0x3f, 0xdb, 0x9f, 0x3d, 0x60,
	0xc5, 0x0b, 0x2e, 0xf9, 0x24, 0x99, 0x39, 0x2c, 0x0e, 0x9d, 0x78, 0xc1, 0xa7, 0x28, 0x04, 0x55,
	0x4e, 0xf2, 0x7d, 0x8e, 0xc2, 0xa7, 0x2c, 0x0e, 0x2d, 0xdb, 0xe7, 0x3e, 0x77, 0xd6, 0xd0, 0xd4,
	0x52, 0x86, 0x7a, 0x65, 0x68, 0xab, 0xb5, 0x0e, 0xcb, 0xe7, 0x18, 0xc5, 0x2e, 0x97, 0xfd, 0xa2,
	0x43, 0xe3, 0x7c, 0x81, 0x4c, 0xe2, 0x4d, 0xe6, 0x77, 0xf1, 0x29, 0x41, 0x21, 0x09, 0x01, 0x23,
	0x62, 0x73, 0x34, 0x4b, 0x1d, 0xbd, 0x5b, 0x73, 0xd5, 0x9b, 0x50, 0xa8, 0x44, 0xdc, 0x43, 0x61,
	0x1a, 0x9d, 0x72, 0xf7, 0x47, 0xcf, 0xa4, 0xaa, 0x11, 0x45, 0x4b, 0xf3, 0x74, 0x7a, 0xcd, 0x3d,
	0x74, 0x33, 0x58, 0x8a, 0x47, 0xcf, 0x47, 0x61, 0x56, 0x8e, 0xe3, 0x87, 0x9e, 0x8f, 0x6e, 0x06,
	0xb3, 0x1f, 0xa0, 0xb9, 0xd7, 0x8b, 0x88, 0x79, 0x24, 0x90, 0x8c, 0xc0, 0x08, 0x98, 0x08, 0x4c,
	0xbd, 0xa3, 0x77, 0x7f, 0x0e, 0xfa, 0x6f, 0xcb, 0xb6, 0xf6, 0xbe, 0x6c, 0xff, 0xf7, 0x43, 0x19,
	0x24, 0x13, 0x3a, 0xe5, 0x73, 0x27, 0x65, 0x3e, 0x99, 0xf1, 0x24, 0xf2, 0x98, 0x0c, 0x79, 0xe4,
	0x60, 0xe4, 0x87, 0x11, 0x3a, 0x69, 0x16, 0x1d, 0x33, 0x11, 0xb8, 0x8a, 0xc0, 0xbe, 0x87, 0xc6,
	0x05, 0x3e, 0xe2, 0xc1, 0xb4, 0x9f, 0x56, 0xe0, 0x37, 0x34, 0xf7, 0x0a, 0x64, 0x23, 0xd8, 0x77,
	0xf0, 0x6b, 0x84, 0xf2, 0xab, 0xca, 0x36, 0x80, 0x5c, 0x85, 0x62, 0x8f, 0xde, 0x1e, 0x43, 0x7d,
	0xc7, 0x9b, 0xab, 0x79, 0x0a, 0xb5, 0xfc, 0x08, 0x50, 0x98, 0xba, 0x5a, 0x4d, 0xbd, 0x60, 0x35,
	0xee, 0x06, 0xd5, 0x7b, 0x2d, 0xc1, 0xb7, 0xdc, 0x4d, 0x2e, 0xa1, 0x9a, 0x6d, 0x89, 0xfc, 0xa3,
	0xab, 0x4b, 0xa4, 0x45, 0x37, 0x64, 0xb5, 0x8f, 0xc6, 0x73, 0x51, 0xb4, 0x94, 0x2c, 0xd3, 0x6b,
	0x9b, 0xac, 0x68, 0x45, 0xdb, 0x64, 0xc5, 0x0a, 0x6b, 0xe4, 0x0c, 0xca, 0x23, 0x94, 0xe4, 0xcf,
	0x06, 0x79, 0x20, 0xb9, 0x55, 0x34, 0xa9, 0xad, 0x91, 0x21, 0x18, 0xa9, 0x54, 0xa4, 0xb5, 0xc9,
	0x3d, 0x14, 0xd4, 0xfa, 0x7b, 0x24, 0xba, 0xea, 0x60, 0x50, 0xb9, 0x2d, 0xb3, 0x38, 0x9c, 0x54,
	0xd5, 0xe7, 0xea, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x41, 0xb2, 0xa9, 0xc6, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProcessClient is the client API for Process service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessClient interface {
	// Create a Process from a Process Definition.
	// It will return an unique identifier which is used to interact with the Process.
	Create(ctx context.Context, in *CreateProcessRequest, opts ...grpc.CallOption) (*CreateProcessResponse, error)
	// Delete a process.
	// An error is returned if one or more Instances of the process are running.
	Delete(ctx context.Context, in *DeleteProcessRequest, opts ...grpc.CallOption) (*DeleteProcessResponse, error)
	// Get returns a process matching the criteria of the request.
	Get(ctx context.Context, in *GetProcessRequest, opts ...grpc.CallOption) (*process.Process, error)
	// List returns processes specified in a request.
	List(ctx context.Context, in *ListProcessRequest, opts ...grpc.CallOption) (*ListProcessResponse, error)
}

type processClient struct {
	cc *grpc.ClientConn
}

func NewProcessClient(cc *grpc.ClientConn) ProcessClient {
	return &processClient{cc}
}

func (c *processClient) Create(ctx context.Context, in *CreateProcessRequest, opts ...grpc.CallOption) (*CreateProcessResponse, error) {
	out := new(CreateProcessResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.Process/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processClient) Delete(ctx context.Context, in *DeleteProcessRequest, opts ...grpc.CallOption) (*DeleteProcessResponse, error) {
	out := new(DeleteProcessResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.Process/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processClient) Get(ctx context.Context, in *GetProcessRequest, opts ...grpc.CallOption) (*process.Process, error) {
	out := new(process.Process)
	err := c.cc.Invoke(ctx, "/mesg.api.Process/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processClient) List(ctx context.Context, in *ListProcessRequest, opts ...grpc.CallOption) (*ListProcessResponse, error) {
	out := new(ListProcessResponse)
	err := c.cc.Invoke(ctx, "/mesg.api.Process/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessServer is the server API for Process service.
type ProcessServer interface {
	// Create a Process from a Process Definition.
	// It will return an unique identifier which is used to interact with the Process.
	Create(context.Context, *CreateProcessRequest) (*CreateProcessResponse, error)
	// Delete a process.
	// An error is returned if one or more Instances of the process are running.
	Delete(context.Context, *DeleteProcessRequest) (*DeleteProcessResponse, error)
	// Get returns a process matching the criteria of the request.
	Get(context.Context, *GetProcessRequest) (*process.Process, error)
	// List returns processes specified in a request.
	List(context.Context, *ListProcessRequest) (*ListProcessResponse, error)
}

// UnimplementedProcessServer can be embedded to have forward compatible implementations.
type UnimplementedProcessServer struct {
}

func (*UnimplementedProcessServer) Create(ctx context.Context, req *CreateProcessRequest) (*CreateProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedProcessServer) Delete(ctx context.Context, req *DeleteProcessRequest) (*DeleteProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedProcessServer) Get(ctx context.Context, req *GetProcessRequest) (*process.Process, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedProcessServer) List(ctx context.Context, req *ListProcessRequest) (*ListProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterProcessServer(s *grpc.Server, srv ProcessServer) {
	s.RegisterService(&_Process_serviceDesc, srv)
}

func _Process_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.Process/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServer).Create(ctx, req.(*CreateProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Process_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.Process/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServer).Delete(ctx, req.(*DeleteProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Process_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.Process/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServer).Get(ctx, req.(*GetProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Process_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mesg.api.Process/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServer).List(ctx, req.(*ListProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Process_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mesg.api.Process",
	HandlerType: (*ProcessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Process_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Process_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Process_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Process_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/api/process.proto",
}
