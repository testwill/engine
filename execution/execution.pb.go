// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: execution.proto

package execution

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// Status represents the status of a single execution.
// Note that a valid execution must have only one status
// flag at time.
type Status int32

const (
	// Unknown status represents any status unknown to execution.
	Status_Unknown Status = 0
	// Created is an initial status after execution creation.
	Status_Created Status = 1
	// InProgress informs that processing of execution has been started.
	Status_InProgress Status = 2
	// Completed is a success status after execution was processed.
	Status_Completed Status = 3
	// Failed is an error status after execution was processed.
	Status_Failed Status = 4
)

var Status_name = map[int32]string{
	0: "Unknown",
	1: "Created",
	2: "InProgress",
	3: "Completed",
	4: "Failed",
}

var Status_value = map[string]int32{
	"Unknown":    0,
	"Created":    1,
	"InProgress": 2,
	"Completed":  3,
	"Failed":     4,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_776e2c5022e94aef, []int{0}
}

// Execution represents a single execution run in engine.
type Execution struct {
	// Hash is a unique hash to identify execution.
	Hash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,1,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash" hash:"-"`
	// parentHash is the unique hash of parent execution.
	// if execution is triggered by another one,
	// dependency execution considered as the parent.
	ParentHash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,2,opt,name=parentHash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"parentHash" hash:"name:2"`
	// eventHash is unique event hash.
	EventHash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,3,opt,name=eventHash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"eventHash" hash:"name:3"`
	// Status is the current status of execution.
	Status Status `protobuf:"varint,4,opt,name=status,proto3,enum=types.Status" json:"status,omitempty" hash:"-"`
	// instanceHash is hash of the instance that can proceed an execution
	InstanceHash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,5,opt,name=instanceHash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"instanceHash" hash:"name:5"`
	// taskKey is the key of the task of this execution.
	TaskKey string `protobuf:"bytes,6,opt,name=taskKey,proto3" json:"taskKey,omitempty" hash:"name:6"`
	// inputs data of the execution.
	Inputs *types.Struct `protobuf:"bytes,7,opt,name=inputs,proto3" json:"inputs,omitempty" hash:"name:7"`
	// outputs are the returned data of successful execution.
	Outputs *types.Struct `protobuf:"bytes,8,opt,name=outputs,proto3" json:"outputs,omitempty" hash:"-"`
	// error message of a failed execution.
	Error string `protobuf:"bytes,9,opt,name=error,proto3" json:"error,omitempty" hash:"-"`
	// tags are optionally associated with execution by the user.
	Tags []string `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty" hash:"name:10"`
	// processHash is the unique hash of the process associated to this execution.
	ProcessHash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,11,opt,name=processHash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"processHash" hash:"name:11"`
	// step of the workflow.
	StepID               string   `protobuf:"bytes,12,opt,name=stepID,proto3" json:"stepID,omitempty" hash:"name:12"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Execution) Reset()         { *m = Execution{} }
func (m *Execution) String() string { return proto.CompactTextString(m) }
func (*Execution) ProtoMessage()    {}
func (*Execution) Descriptor() ([]byte, []int) {
	return fileDescriptor_776e2c5022e94aef, []int{0}
}
func (m *Execution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Execution.Unmarshal(m, b)
}
func (m *Execution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Execution.Marshal(b, m, deterministic)
}
func (m *Execution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Execution.Merge(m, src)
}
func (m *Execution) XXX_Size() int {
	return xxx_messageInfo_Execution.Size(m)
}
func (m *Execution) XXX_DiscardUnknown() {
	xxx_messageInfo_Execution.DiscardUnknown(m)
}

var xxx_messageInfo_Execution proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("types.Status", Status_name, Status_value)
	proto.RegisterType((*Execution)(nil), "types.Execution")
}

func init() { proto.RegisterFile("execution.proto", fileDescriptor_776e2c5022e94aef) }

var fileDescriptor_776e2c5022e94aef = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0xad, 0xd8, 0x96, 0xa3, 0xf1, 0x9f, 0xba, 0x7b, 0xa9, 0x08, 0xa5, 0x36, 0x7b, 0x28,
	0x26, 0x21, 0x72, 0x6d, 0xd3, 0x16, 0x72, 0x09, 0x28, 0x49, 0x69, 0x28, 0x85, 0x22, 0xd3, 0x4b,
	0x6f, 0x6b, 0x7b, 0xb2, 0x16, 0xb1, 0x77, 0xc5, 0xee, 0xaa, 0x6d, 0xde, 0x30, 0xcf, 0x50, 0x8a,
	0xa1, 0xaf, 0x90, 0x27, 0x28, 0x5a, 0x45, 0x8d, 0xd5, 0x40, 0x29, 0x21, 0x37, 0xcd, 0xce, 0x37,
	0xbf, 0xfd, 0x76, 0x66, 0x10, 0x3c, 0xc1, 0xef, 0x38, 0x4f, 0x4d, 0x2c, 0x45, 0x90, 0x28, 0x69,
	0x24, 0xa9, 0x9b, 0xab, 0x04, 0xf5, 0xde, 0x73, 0x2e, 0x25, 0x5f, 0xe1, 0xd0, 0x1e, 0xce, 0xd2,
	0x8b, 0xa1, 0x36, 0x2a, 0x9d, 0x9b, 0x5c, 0xb4, 0x47, 0xb9, 0xe4, 0xf2, 0x2e, 0x97, 0x45, 0x36,
	0xb0, 0x5f, 0xb9, 0x86, 0xfe, 0x74, 0xc1, 0x3b, 0x2b, 0xe0, 0x64, 0x0a, 0xb5, 0x25, 0xd3, 0x4b,
	0xdf, 0xe9, 0x3b, 0x83, 0x56, 0x78, 0x7c, 0xbd, 0xe9, 0x55, 0x7e, 0x6c, 0x7a, 0x07, 0x3c, 0x36,
	0xcb, 0x74, 0x16, 0xcc, 0xe5, 0x7a, 0xb8, 0x46, 0xcd, 0x0f, 0x2f, 0x64, 0x2a, 0x16, 0x2c, 0xab,
	0x18, 0xa2, 0xe0, 0xb1, 0xc0, 0x61, 0x56, 0x15, 0xbc, 0x67, 0x7a, 0x79, 0xb3, 0xe9, 0xed, 0x66,
	0xc1, 0x11, 0x3d, 0xa4, 0x91, 0x85, 0x91, 0x05, 0x40, 0xc2, 0x14, 0x0a, 0x93, 0xe5, 0xfd, 0x1d,
	0x8b, 0x3e, 0x7d, 0x18, 0xba, 0x9d, 0xa3, 0x05, 0x5b, 0xe3, 0xd1, 0x98, 0x46, 0x5b, 0x5c, 0x32,
	0x03, 0x0f, 0xbf, 0x16, 0x97, 0x54, 0x1f, 0xeb, 0x92, 0x09, 0x8d, 0xee, 0xb0, 0x64, 0x02, 0xae,
	0x36, 0xcc, 0xa4, 0xda, 0xaf, 0xf5, 0x9d, 0x41, 0x67, 0xdc, 0x0e, 0xec, 0x18, 0x82, 0xa9, 0x3d,
	0x0c, 0x5b, 0xa5, 0xc7, 0xdf, 0x4a, 0xc9, 0x12, 0x5a, 0xb1, 0xd0, 0x86, 0x89, 0x39, 0x5a, 0x6f,
	0xf5, 0xc7, 0xf2, 0xf6, 0x9a, 0x46, 0x25, 0x32, 0x39, 0x80, 0x86, 0x61, 0xfa, 0xf2, 0x03, 0x5e,
	0xf9, 0x6e, 0xdf, 0x19, 0x78, 0xe1, 0xd3, 0xbf, 0x2a, 0xde, 0xd0, 0xa8, 0x50, 0x90, 0x10, 0xdc,
	0x58, 0x24, 0xa9, 0xd1, 0x7e, 0xa3, 0xef, 0x0c, 0x9a, 0xe3, 0x67, 0x41, 0xbe, 0x4b, 0x41, 0xb1,
	0x2f, 0xc1, 0xd4, 0xee, 0xd2, 0x3d, 0xc8, 0x5b, 0x1a, 0xdd, 0x56, 0x92, 0x63, 0x68, 0xc8, 0xd4,
	0x58, 0xc8, 0xee, 0xbf, 0x21, 0xe5, 0xd6, 0x14, 0x55, 0x84, 0x42, 0x1d, 0x95, 0x92, 0xca, 0xf7,
	0xac, 0xdf, 0xb2, 0x2a, 0x4f, 0x91, 0x97, 0x50, 0x33, 0x8c, 0x6b, 0x1f, 0xfa, 0xd5, 0x81, 0x17,
	0x92, 0x9b, 0x4d, 0xaf, 0xb3, 0xe5, 0x66, 0xf4, 0x8a, 0x46, 0x36, 0x4f, 0x38, 0x34, 0x13, 0x25,
	0xe7, 0xa8, 0xb5, 0x6d, 0x73, 0xd3, 0xb6, 0xf9, 0xec, 0x61, 0x6d, 0x2e, 0xdd, 0x30, 0xa2, 0xd1,
	0x36, 0x99, 0xec, 0x67, 0x5b, 0x80, 0xc9, 0xf9, 0xa9, 0xdf, 0xb2, 0xae, 0xef, 0x59, 0x1a, 0xdb,
	0xe1, 0x67, 0x8a, 0xfd, 0x8f, 0xe0, 0xe6, 0xcb, 0x41, 0x9a, 0xd0, 0xf8, 0x2c, 0x2e, 0x85, 0xfc,
	0x26, 0xba, 0x95, 0x2c, 0x38, 0x51, 0xc8, 0x0c, 0x2e, 0xba, 0x0e, 0xe9, 0x00, 0x9c, 0x8b, 0x4f,
	0x4a, 0x72, 0x85, 0x5a, 0x77, 0x77, 0x48, 0x1b, 0xbc, 0x13, 0xb9, 0x4e, 0x56, 0x98, 0xa5, 0xab,
	0x04, 0xc0, 0x7d, 0xc7, 0xe2, 0x15, 0x2e, 0xba, 0xb5, 0x70, 0x74, 0xfd, 0xeb, 0x45, 0xe5, 0xcb,
	0x7f, 0x3c, 0xe6, 0xcf, 0xff, 0x62, 0xe6, 0xda, 0x51, 0x4c, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xf8, 0x72, 0x51, 0x0c, 0x43, 0x04, 0x00, 0x00,
}
