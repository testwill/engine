// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/execution/internal/types/msg.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// The message to create an Execution.
type MsgCreate struct {
	// The msg's signer.
	Signer github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty" validate:"required,accaddress"`
	// taskKey to filter executions.
	TaskKey string        `protobuf:"bytes,2,opt,name=taskKey,proto3" json:"taskKey,omitempty" validate:"required,printascii"`
	Inputs  *types.Struct `protobuf:"bytes,3,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// tags the execution.
	Tags                 []string                                    `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" validate:"dive,printascii"`
	ParentHash           github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,5,opt,name=parentHash,proto3,casttype=github.com/mesg-foundation/engine/hash.Hash" json:"parentHash,omitempty" validate:"required_without=EventHash,omitempty,hash"`
	EventHash            github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,6,opt,name=eventHash,proto3,casttype=github.com/mesg-foundation/engine/hash.Hash" json:"eventHash,omitempty" validate:"required_without=ParentHash,omitempty,hash"`
	ProcessHash          github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,7,opt,name=processHash,proto3,casttype=github.com/mesg-foundation/engine/hash.Hash" json:"processHash,omitempty" validate:"omitempty,hash"`
	NodeKey              string                                      `protobuf:"bytes,8,opt,name=nodeKey,proto3" json:"nodeKey,omitempty"`
	ExecutorHash         github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,9,opt,name=executorHash,proto3,casttype=github.com/mesg-foundation/engine/hash.Hash" json:"executorHash,omitempty" validate:"required,hash"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *MsgCreate) Reset()         { *m = MsgCreate{} }
func (m *MsgCreate) String() string { return proto.CompactTextString(m) }
func (*MsgCreate) ProtoMessage()    {}
func (*MsgCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c59efd9270dae25, []int{0}
}
func (m *MsgCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgCreate.Unmarshal(m, b)
}
func (m *MsgCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgCreate.Marshal(b, m, deterministic)
}
func (m *MsgCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreate.Merge(m, src)
}
func (m *MsgCreate) XXX_Size() int {
	return xxx_messageInfo_MsgCreate.Size(m)
}
func (m *MsgCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreate proto.InternalMessageInfo

func (m *MsgCreate) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func (m *MsgCreate) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *MsgCreate) GetInputs() *types.Struct {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *MsgCreate) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *MsgCreate) GetParentHash() github_com_mesg_foundation_engine_hash.Hash {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *MsgCreate) GetEventHash() github_com_mesg_foundation_engine_hash.Hash {
	if m != nil {
		return m.EventHash
	}
	return nil
}

func (m *MsgCreate) GetProcessHash() github_com_mesg_foundation_engine_hash.Hash {
	if m != nil {
		return m.ProcessHash
	}
	return nil
}

func (m *MsgCreate) GetNodeKey() string {
	if m != nil {
		return m.NodeKey
	}
	return ""
}

func (m *MsgCreate) GetExecutorHash() github_com_mesg_foundation_engine_hash.Hash {
	if m != nil {
		return m.ExecutorHash
	}
	return nil
}

// The message to update an Execution.
type MsgUpdate struct {
	// The execution's executor.
	Executor github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=executor,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"executor,omitempty" validate:"required,accaddress"`
	// Hash represents execution.
	Hash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,2,opt,name=hash,proto3,casttype=github.com/mesg-foundation/engine/hash.Hash" json:"hash,omitempty" validate:"required,hash"`
	// result pass to execution
	//
	// Types that are valid to be assigned to Result:
	//	*MsgUpdate_Outputs
	//	*MsgUpdate_Error
	Result isMsgUpdate_Result `protobuf_oneof:"result"`
	// Start time
	Start uint64 `protobuf:"varint,5,opt,name=start,proto3" json:"start,omitempty" validate:"required"`
	// Stop time
	Stop                 uint64   `protobuf:"varint,6,opt,name=stop,proto3" json:"stop,omitempty" validate:"required,gtfield=Start"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgUpdate) Reset()         { *m = MsgUpdate{} }
func (m *MsgUpdate) String() string { return proto.CompactTextString(m) }
func (*MsgUpdate) ProtoMessage()    {}
func (*MsgUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c59efd9270dae25, []int{1}
}
func (m *MsgUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgUpdate.Unmarshal(m, b)
}
func (m *MsgUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgUpdate.Marshal(b, m, deterministic)
}
func (m *MsgUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdate.Merge(m, src)
}
func (m *MsgUpdate) XXX_Size() int {
	return xxx_messageInfo_MsgUpdate.Size(m)
}
func (m *MsgUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdate proto.InternalMessageInfo

type isMsgUpdate_Result interface {
	isMsgUpdate_Result()
}

type MsgUpdate_Outputs struct {
	Outputs *types.Struct `protobuf:"bytes,3,opt,name=outputs,proto3,oneof" json:"outputs,omitempty"`
}
type MsgUpdate_Error struct {
	Error string `protobuf:"bytes,4,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (*MsgUpdate_Outputs) isMsgUpdate_Result() {}
func (*MsgUpdate_Error) isMsgUpdate_Result()   {}

func (m *MsgUpdate) GetResult() isMsgUpdate_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *MsgUpdate) GetExecutor() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *MsgUpdate) GetHash() github_com_mesg_foundation_engine_hash.Hash {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *MsgUpdate) GetOutputs() *types.Struct {
	if x, ok := m.GetResult().(*MsgUpdate_Outputs); ok {
		return x.Outputs
	}
	return nil
}

func (m *MsgUpdate) GetError() string {
	if x, ok := m.GetResult().(*MsgUpdate_Error); ok {
		return x.Error
	}
	return ""
}

func (m *MsgUpdate) GetStart() uint64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *MsgUpdate) GetStop() uint64 {
	if m != nil {
		return m.Stop
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MsgUpdate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MsgUpdate_Outputs)(nil),
		(*MsgUpdate_Error)(nil),
	}
}

func init() {
	proto.RegisterType((*MsgCreate)(nil), "mesg.execution.types.MsgCreate")
	proto.RegisterType((*MsgUpdate)(nil), "mesg.execution.types.MsgUpdate")
}

func init() {
	proto.RegisterFile("x/execution/internal/types/msg.proto", fileDescriptor_6c59efd9270dae25)
}

var fileDescriptor_6c59efd9270dae25 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5f, 0x4f, 0xd4, 0x40,
	0x10, 0xe7, 0xa4, 0xf4, 0xb8, 0x85, 0xa7, 0xf5, 0x5f, 0xc5, 0x84, 0x9e, 0x8d, 0x26, 0x97, 0xc8,
	0xb5, 0x11, 0xf5, 0x41, 0x03, 0x26, 0xd4, 0x98, 0x90, 0xa0, 0x89, 0x29, 0xf1, 0xc5, 0x17, 0xb2,
	0xb4, 0x43, 0xbb, 0xe1, 0xba, 0x5b, 0x77, 0xa7, 0x08, 0x5f, 0x40, 0xbf, 0x84, 0xef, 0x7e, 0xab,
	0x7e, 0x88, 0x3e, 0xfa, 0x64, 0xba, 0xbd, 0x1e, 0x10, 0x40, 0x12, 0xd1, 0xa7, 0xbb, 0xc9, 0xfc,
	0x7e, 0xf3, 0xfb, 0xcd, 0xec, 0x4c, 0xc9, 0xe3, 0xe3, 0x00, 0x8e, 0x21, 0x2e, 0x91, 0x4b, 0x11,
	0x70, 0x81, 0xa0, 0x04, 0x9b, 0x04, 0x78, 0x52, 0x80, 0x0e, 0x72, 0x9d, 0xfa, 0x85, 0x92, 0x28,
	0xe9, 0x9d, 0x1c, 0x74, 0xea, 0xcf, 0x80, 0xbe, 0xc9, 0xaf, 0x78, 0xa9, 0x4c, 0x65, 0x60, 0x10,
	0xfb, 0xe5, 0x41, 0xd0, 0x44, 0x26, 0x30, 0xff, 0x5a, 0xe6, 0xca, 0xc3, 0x59, 0xba, 0xad, 0xa9,
	0x51, 0x95, 0x31, 0xb6, 0x49, 0xef, 0xa7, 0x4d, 0x06, 0x1f, 0x74, 0xfa, 0x56, 0x01, 0x43, 0xa0,
	0x87, 0xc4, 0xd6, 0x3c, 0x15, 0xa0, 0x9c, 0xde, 0xb0, 0x37, 0x5a, 0x0e, 0x77, 0xeb, 0xca, 0x5d,
	0x3d, 0x62, 0x13, 0x9e, 0x30, 0x84, 0xd7, 0x9e, 0x82, 0x2f, 0x25, 0x57, 0x90, 0xac, 0xb1, 0x38,
	0x66, 0x49, 0xa2, 0x40, 0x6b, 0xef, 0x57, 0xe5, 0x8e, 0x53, 0x8e, 0x59, 0xb9, 0xef, 0xc7, 0x32,
	0x0f, 0x62, 0xa9, 0x73, 0xa9, 0xa7, 0x3f, 0x63, 0x9d, 0x1c, 0xb6, 0xa2, 0xfe, 0x56, 0x1c, 0x6f,
	0xb5, 0x8c, 0x68, 0x2a, 0x41, 0x37, 0x48, 0x1f, 0x99, 0x3e, 0xdc, 0x81, 0x13, 0xe7, 0xd6, 0xb0,
	0x37, 0x1a, 0x84, 0xde, 0x15, 0x6a, 0x85, 0xe2, 0x02, 0x99, 0x8e, 0x39, 0xf7, 0xa2, 0x8e, 0x42,
	0xc7, 0xc4, 0xe6, 0xa2, 0x28, 0x51, 0x3b, 0xf3, 0xc3, 0xde, 0x68, 0x69, 0xfd, 0xae, 0x6f, 0x06,
	0xd4, 0xf5, 0xea, 0xef, 0x9a, 0x2e, 0xa3, 0x29, 0x88, 0xae, 0x13, 0x0b, 0x59, 0xaa, 0x1d, 0x6b,
	0x38, 0x3f, 0x1a, 0x84, 0xab, 0x75, 0xe5, 0xae, 0x9c, 0x2a, 0x25, 0xfc, 0x08, 0xce, 0xa9, 0x18,
	0x2c, 0xfd, 0xde, 0x23, 0xa4, 0x60, 0x0a, 0x04, 0x6e, 0x33, 0x9d, 0x39, 0x0b, 0x66, 0x24, 0x69,
	0x5d, 0xb9, 0x2f, 0x2e, 0x9a, 0xdc, 0xfb, 0xca, 0x31, 0x93, 0x25, 0x6e, 0xbe, 0x3b, 0x9a, 0xe2,
	0xd7, 0x64, 0xce, 0x11, 0xf2, 0x02, 0x4f, 0xd6, 0x32, 0xa6, 0xb3, 0x66, 0x50, 0x4f, 0xcf, 0x0c,
	0xaa, 0x71, 0x3b, 0x3e, 0x90, 0xa5, 0x48, 0x98, 0x79, 0x78, 0x10, 0x29, 0x17, 0x10, 0x34, 0x50,
	0xbf, 0xa1, 0x47, 0x67, 0xa4, 0xe9, 0xb7, 0x1e, 0x19, 0x40, 0x57, 0xd8, 0xb1, 0x8d, 0x91, 0xac,
	0xae, 0xdc, 0x97, 0x7f, 0x30, 0xf2, 0x71, 0x46, 0xbf, 0xa9, 0x93, 0x53, 0x69, 0x2a, 0xc8, 0x52,
	0xa1, 0x64, 0x0c, 0x5a, 0x1b, 0x27, 0x7d, 0xe3, 0xe4, 0x7d, 0x5d, 0xb9, 0x0f, 0x4e, 0x9d, 0xdc,
	0x50, 0xed, 0xac, 0x00, 0x75, 0x48, 0x5f, 0xc8, 0x04, 0x9a, 0x1d, 0x59, 0x6c, 0x76, 0x24, 0xea,
	0x42, 0x2a, 0xc9, 0x72, 0x7b, 0x0c, 0x52, 0x19, 0x2b, 0x03, 0x63, 0x65, 0xa7, 0xae, 0x5c, 0xe7,
	0x92, 0x15, 0xfa, 0x2b, 0x27, 0xe7, 0x04, 0xbc, 0x1f, 0xf3, 0xe6, 0x52, 0x3e, 0x15, 0x4d, 0x69,
	0x2a, 0xc9, 0x62, 0x97, 0xfd, 0x9f, 0xb7, 0x32, 0x13, 0xa1, 0x7b, 0xc4, 0x6a, 0x9c, 0x99, 0x53,
	0xf9, 0xc7, 0x7d, 0x9a, 0xc2, 0xf4, 0x19, 0xe9, 0xcb, 0x12, 0xaf, 0xbd, 0xa8, 0xed, 0xb9, 0xa8,
	0xc3, 0xd1, 0x7b, 0x64, 0x01, 0x94, 0x92, 0xca, 0xb1, 0x9a, 0xb7, 0xd9, 0x9e, 0x8b, 0xda, 0x90,
	0x8e, 0xc9, 0x82, 0x46, 0xa6, 0xd0, 0x9c, 0x8c, 0x15, 0xde, 0xaf, 0x2b, 0xf7, 0xf6, 0x45, 0xb3,
	0x5e, 0xd4, 0xa2, 0xe8, 0x2b, 0x62, 0x69, 0x94, 0x85, 0xd9, 0x6b, 0x2b, 0x7c, 0x52, 0x57, 0xee,
	0xa3, 0x4b, 0x5a, 0x4b, 0xf1, 0x80, 0xc3, 0x24, 0xd9, 0xdc, 0x6d, 0x38, 0x5e, 0x64, 0x28, 0xe1,
	0x22, 0xb1, 0x15, 0xe8, 0x72, 0x82, 0xe1, 0x9b, 0xcf, 0x1b, 0xd7, 0xf7, 0x7c, 0xf5, 0x97, 0x76,
	0xdf, 0x36, 0x7d, 0x3e, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x33, 0x84, 0xa3, 0x8e, 0x05,
	0x00, 0x00,
}
