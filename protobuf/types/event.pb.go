// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/types/event.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// Event represents a single event run in engine.
type Event struct {
	// Hash is a unique hash to identify event.
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// instanceHash is hash of instance that can proceed an execution.
	InstanceHash []byte `protobuf:"bytes,2,opt,name=instanceHash,proto3" json:"instanceHash,omitempty"`
	// key is the key of the event.
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// data is the data for the event.
	Data                 *_struct.Struct `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_118d36ccf50c3798, []int{0}
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

func (m *Event) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Event) GetInstanceHash() []byte {
	if m != nil {
		return m.InstanceHash
	}
	return nil
}

func (m *Event) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Event) GetData() *_struct.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "types.Event")
}

func init() { proto.RegisterFile("protobuf/types/event.proto", fileDescriptor_118d36ccf50c3798) }

var fileDescriptor_118d36ccf50c3798 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b,
	0xd1, 0x03, 0x0b, 0x0a, 0xb1, 0x82, 0x85, 0xa4, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5,
	0xe1, 0x2a, 0x8b, 0x4b, 0x8a, 0x4a, 0x93, 0xa1, 0x8a, 0x94, 0xea, 0xb8, 0x58, 0x5d, 0x41, 0x7a,
	0x84, 0x84, 0xb8, 0x58, 0x32, 0x12, 0x8b, 0x33, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xc0,
	0x6c, 0x21, 0x25, 0x2e, 0x9e, 0xcc, 0xbc, 0xe2, 0x92, 0xc4, 0xbc, 0xe4, 0x54, 0x0f, 0x90, 0x1c,
	0x13, 0x58, 0x0e, 0x45, 0x4c, 0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x52, 0x82, 0x59, 0x81, 0x51,
	0x83, 0x33, 0x08, 0xc4, 0x14, 0xd2, 0xe6, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x51, 0x60,
	0xd4, 0xe0, 0x36, 0x12, 0xd7, 0x83, 0xd8, 0xaf, 0x07, 0xb3, 0x5f, 0x2f, 0x18, 0x6c, 0x7f, 0x10,
	0x58, 0x91, 0x93, 0x51, 0x94, 0x41, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae,
	0x7e, 0x6e, 0x6a, 0x71, 0xba, 0x6e, 0x5a, 0x7e, 0x69, 0x5e, 0x4a, 0x62, 0x49, 0x66, 0x7e, 0x9e,
	0x7e, 0x6a, 0x5e, 0x7a, 0x66, 0x1e, 0x92, 0xd3, 0xc1, 0x3e, 0x4a, 0x62, 0x03, 0xf3, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x90, 0x8e, 0x9f, 0xcd, 0xfd, 0x00, 0x00, 0x00,
}
