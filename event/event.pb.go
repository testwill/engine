// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: event.proto

package event

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_mesg_foundation_engine_hash "github.com/mesg-foundation/engine/hash"
	types "github.com/mesg-foundation/engine/protobuf/types"
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

// Event represents a single event run in engine.
type Event struct {
	// Hash is a unique hash to identify event.
	Hash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,casttype=github.com/mesg-foundation/engine/hash.Hash" json:"hash,omitempty" hash:"-" validate:"required,hash"`
	// instanceHash is hash of instance that can proceed an execution.
	InstanceHash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,2,opt,name=instanceHash,proto3,casttype=github.com/mesg-foundation/engine/hash.Hash" json:"instanceHash,omitempty" hash:"name:2" validate:"required,hash"`
	// key is the key of the event.
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty" hash:"name:3" validate:"required,printascii"`
	// data is the data for the event.
	Data                 *types.Struct `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty" hash:"name:4"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Event)(nil), "mesg.types.Event")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xff, 0x69, 0xfb, 0x17, 0xdc, 0xd6, 0x83, 0x01, 0x21, 0x54, 0x30, 0x71, 0x0f, 0x5a,
	0xb0, 0xdd, 0x95, 0xd6, 0x53, 0xf0, 0x14, 0x10, 0x3c, 0x47, 0x2f, 0x7a, 0xdb, 0x24, 0xd3, 0x64,
	0xd1, 0xec, 0xc6, 0xec, 0xa6, 0xd0, 0x83, 0xef, 0xe7, 0x53, 0x04, 0x7c, 0x85, 0x1c, 0x3d, 0xc9,
	0x6e, 0x55, 0x2a, 0x08, 0x5e, 0xbc, 0xe5, 0x9b, 0x99, 0xef, 0xfb, 0x65, 0x67, 0xd0, 0x10, 0x56,
	0x20, 0x34, 0xa9, 0x6a, 0xa9, 0xa5, 0x8b, 0x4a, 0x50, 0x39, 0xd1, 0xeb, 0x0a, 0xd4, 0x18, 0xe7,
	0x32, 0x97, 0xd4, 0xd6, 0x93, 0x66, 0x49, 0x8d, 0xb2, 0xc2, 0x7e, 0x6d, 0xe6, 0xc7, 0x87, 0x5f,
	0x6d, 0xeb, 0xa1, 0x4a, 0xd7, 0x4d, 0xfa, 0x11, 0x86, 0xbb, 0x1e, 0xfa, 0x7f, 0x65, 0xc2, 0xdd,
	0x02, 0x0d, 0x0a, 0xa6, 0x0a, 0xcf, 0x09, 0x9c, 0xc9, 0x28, 0xba, 0xed, 0x5a, 0xff, 0xd8, 0xe8,
	0x10, 0xcf, 0x70, 0xb0, 0x62, 0x8f, 0x3c, 0x63, 0x1a, 0x42, 0x5c, 0xc3, 0x53, 0xc3, 0x6b, 0xc8,
	0xa6, 0xa6, 0x87, 0xdf, 0x5a, 0xff, 0x2c, 0xe7, 0xba, 0x68, 0x12, 0x92, 0xca, 0x92, 0x9a, 0x1f,
	0x9b, 0x2d, 0x65, 0x23, 0x32, 0xa6, 0xb9, 0x14, 0x14, 0x44, 0xce, 0x05, 0x50, 0x33, 0x4a, 0xae,
	0x99, 0x2a, 0x62, 0x4b, 0x70, 0x9f, 0xd1, 0x88, 0x0b, 0xa5, 0x99, 0x48, 0xc1, 0x54, 0xbd, 0x9e,
	0x25, 0xde, 0x75, 0xad, 0x7f, 0xb2, 0x21, 0x0a, 0x56, 0x42, 0x38, 0xff, 0x3b, 0xec, 0x37, 0x9c,
	0x1b, 0xa1, 0xfe, 0x03, 0xac, 0xbd, 0x7e, 0xe0, 0x4c, 0x76, 0xa3, 0xf3, 0xae, 0xf5, 0xa7, 0x5b,
	0xd4, 0xc5, 0x8f, 0xd4, 0xaa, 0xe6, 0x42, 0x33, 0x95, 0x72, 0x8e, 0x63, 0x63, 0x76, 0x2f, 0xd1,
	0x20, 0x63, 0x9a, 0x79, 0x83, 0xc0, 0x99, 0x0c, 0xe7, 0x07, 0xc4, 0x9e, 0xe4, 0x73, 0xcf, 0xe4,
	0xc6, 0x6e, 0x38, 0xda, 0xef, 0x5a, 0x7f, 0x6f, 0x2b, 0xfb, 0x02, 0xc7, 0xd6, 0x15, 0xcd, 0x5e,
	0x5e, 0x8f, 0xfe, 0xdd, 0x9f, 0xfe, 0xfe, 0x04, 0x7b, 0xf6, 0x64, 0xc7, 0x06, 0x2f, 0xde, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x48, 0xc8, 0x52, 0x5d, 0x06, 0x02, 0x00, 0x00,
}
