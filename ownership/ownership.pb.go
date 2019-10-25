// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ownership.proto

package ownership

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

// Ownership is a ownership relation between one owner and a resource.
type Ownership struct {
	// Service's hash.
	Hash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash" hash:"-" validate:"required"`
	// The owner of the resource.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty" hash:"name:2" validate:"required"`
	// Resource's hash.
	//
	// Types that are valid to be assigned to Resource:
	//	*Ownership_ServiceHash
	Resource             isOwnership_Resource `protobuf_oneof:"resource"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Ownership) Reset()         { *m = Ownership{} }
func (m *Ownership) String() string { return proto.CompactTextString(m) }
func (*Ownership) ProtoMessage()    {}
func (*Ownership) Descriptor() ([]byte, []int) {
	return fileDescriptor_21ae26e0dccf9d04, []int{0}
}
func (m *Ownership) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ownership.Unmarshal(m, b)
}
func (m *Ownership) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ownership.Marshal(b, m, deterministic)
}
func (m *Ownership) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ownership.Merge(m, src)
}
func (m *Ownership) XXX_Size() int {
	return xxx_messageInfo_Ownership.Size(m)
}
func (m *Ownership) XXX_DiscardUnknown() {
	xxx_messageInfo_Ownership.DiscardUnknown(m)
}

var xxx_messageInfo_Ownership proto.InternalMessageInfo

type isOwnership_Resource interface {
	isOwnership_Resource()
	Equal(interface{}) bool
}

type Ownership_ServiceHash struct {
	ServiceHash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,3,opt,name=serviceHash,proto3,oneof,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"serviceHash,omitempty" hash:"name:3"`
}

func (*Ownership_ServiceHash) isOwnership_Resource() {}

func (m *Ownership) GetResource() isOwnership_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Ownership) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Ownership_ServiceHash)(nil),
	}
}

func init() {
	proto.RegisterType((*Ownership)(nil), "mesg.types.Ownership")
}

func init() { proto.RegisterFile("ownership.proto", fileDescriptor_21ae26e0dccf9d04) }

var fileDescriptor_21ae26e0dccf9d04 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x2f, 0xcf, 0x4b,
	0x2d, 0x2a, 0xce, 0xc8, 0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0x4d, 0x2d,
	0x4e, 0xd7, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0x52, 0x4a, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x8b,
	0x27, 0x95, 0xa6, 0xe9, 0x83, 0x78, 0x60, 0x0e, 0x98, 0x05, 0x51, 0xaf, 0xb4, 0x98, 0x89, 0x8b,
	0xd3, 0x1f, 0x66, 0x86, 0x50, 0x3a, 0x17, 0x4b, 0x46, 0x62, 0x71, 0x86, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x8f, 0x53, 0xf0, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0x6b, 0xa7, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x83, 0x8c, 0xd7, 0x4d, 0xcb, 0x2f, 0xcd, 0x4b, 0x49,
	0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x4f, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0xd5, 0x07, 0xe9, 0xd2, 0xf3,
	0x48, 0x2c, 0xce, 0xf8, 0x74, 0x4f, 0x5e, 0x06, 0xc4, 0xb1, 0x52, 0xd2, 0x55, 0x52, 0x28, 0x4b,
	0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xb5, 0x52, 0x2a, 0x4a, 0x2d, 0x2c, 0xcd, 0x2c, 0x4a, 0x4d,
	0x51, 0x0a, 0x02, 0x5b, 0x20, 0x64, 0xcd, 0xc5, 0x0a, 0x76, 0xb9, 0x04, 0x93, 0x02, 0xa3, 0x06,
	0xa7, 0x93, 0xea, 0xa7, 0x7b, 0xf2, 0x8a, 0x10, 0x6d, 0x79, 0x89, 0xb9, 0xa9, 0x56, 0x46, 0xd8,
	0xf5, 0x42, 0xf4, 0x08, 0xa5, 0x70, 0x71, 0x17, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x82, 0x6c,
	0x94, 0x60, 0x06, 0x3b, 0xd6, 0x81, 0x74, 0x87, 0xf2, 0x22, 0xd9, 0x68, 0xac, 0xe4, 0xc1, 0x10,
	0x84, 0x6c, 0xac, 0x13, 0x17, 0x17, 0x47, 0x51, 0x6a, 0x71, 0x7e, 0x69, 0x51, 0x72, 0xaa, 0x93,
	0xe9, 0x89, 0x87, 0x72, 0x0c, 0x2b, 0x1e, 0xc9, 0x31, 0x46, 0x11, 0x61, 0x05, 0x3c, 0x4a, 0x92,
	0xd8, 0xc0, 0x61, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x18, 0x33, 0x3a, 0xc5, 0xa6, 0x01,
	0x00, 0x00,
}

func (this *Ownership) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Ownership)
	if !ok {
		that2, ok := that.(Ownership)
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
	if this.Owner != that1.Owner {
		return false
	}
	if that1.Resource == nil {
		if this.Resource != nil {
			return false
		}
	} else if this.Resource == nil {
		return false
	} else if !this.Resource.Equal(that1.Resource) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Ownership_ServiceHash) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Ownership_ServiceHash)
	if !ok {
		that2, ok := that.(Ownership_ServiceHash)
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
	if !this.ServiceHash.Equal(that1.ServiceHash) {
		return false
	}
	return true
}
