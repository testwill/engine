// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: instance.proto

package instance

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_mesg_foundation_engine_hash "github.com/mesg-foundation/engine/hash"
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

// Instance represents service's instance.
type Instance struct {
	Hash                 github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash" hash:"-" validate:"required,hash"`
	ServiceHash          github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,2,opt,name=serviceHash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"serviceHash" hash:"name:2" validate:"required,hash"`
	EnvHash              github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,3,opt,name=envHash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"envHash" hash:"name:3" validate:"required,hash"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd22322185b2070b, []int{0}
}
func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (m *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(m, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Instance)(nil), "mesg.types.Instance")
}

func init() { proto.RegisterFile("instance.proto", fileDescriptor_fd22322185b2070b) }

var fileDescriptor_fd22322185b2070b = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0x2b, 0x2e,
	0x49, 0xcc, 0x4b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x4d, 0x2d, 0x4e,
	0xd7, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0x52, 0x4a, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x8b, 0x27,
	0x95, 0xa6, 0xe9, 0x83, 0x78, 0x60, 0x0e, 0x98, 0x05, 0x51, 0xaf, 0xf4, 0x89, 0x89, 0x8b, 0xc3,
	0x13, 0x6a, 0x84, 0x50, 0x0e, 0x17, 0x4b, 0x46, 0x62, 0x71, 0x86, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x8f, 0x53, 0xc4, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0x6b, 0xa7, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x83, 0x4c, 0xd7, 0x4d, 0xcb, 0x2f, 0xcd, 0x4b, 0x49, 0x2c,
	0xc9, 0xcc, 0xcf, 0xd3, 0x4f, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0xd5, 0x07, 0xe9, 0xd2, 0xf3, 0x48,
	0x2c, 0xce, 0xf8, 0x74, 0x4f, 0x5e, 0x11, 0xc4, 0xb1, 0x52, 0xd2, 0x55, 0x52, 0x28, 0x4b, 0xcc,
	0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xb5, 0x52, 0x2a, 0x4a, 0x2d, 0x2c, 0xcd, 0x2c, 0x4a, 0x4d, 0xd1,
	0x01, 0xc9, 0x29, 0x05, 0x81, 0x6d, 0x11, 0xaa, 0xe7, 0xe2, 0x2e, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x05, 0xe9, 0x94, 0x60, 0x02, 0x5b, 0x1a, 0x4b, 0x9e, 0xa5, 0x6a, 0x10, 0x4b, 0xf3, 0x12,
	0x73, 0x53, 0xad, 0x8c, 0xf0, 0xd8, 0x8c, 0x6c, 0xa3, 0x50, 0x39, 0x17, 0x7b, 0x6a, 0x5e, 0x19,
	0xd8, 0x72, 0x66, 0x6a, 0x59, 0x6e, 0x8c, 0xc7, 0x72, 0x98, 0x6d, 0x4e, 0x26, 0x27, 0x1e, 0xca,
	0x31, 0xac, 0x78, 0x24, 0xc7, 0x18, 0xa5, 0x45, 0xd8, 0x26, 0x58, 0x04, 0x27, 0xb1, 0x81, 0x63,
	0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x1b, 0xdc, 0x00, 0xf3, 0x01, 0x00, 0x00,
}

func (this *Instance) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Instance)
	if !ok {
		that2, ok := that.(Instance)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Hash.Equal(that1.Hash) {
		return false
	}
	if !this.ServiceHash.Equal(that1.ServiceHash) {
		return false
	}
	if !this.EnvHash.Equal(that1.EnvHash) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
