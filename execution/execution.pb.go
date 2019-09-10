// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: execution.proto

package execution

import (
	bytes "bytes"
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
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdf, 0x6a, 0xdb, 0x3e,
	0x14, 0xc7, 0xe3, 0x26, 0x71, 0xea, 0x93, 0x3f, 0xbf, 0xfc, 0x04, 0x03, 0xd3, 0x41, 0x13, 0x74,
	0x31, 0x4c, 0x4b, 0xed, 0x35, 0xa1, 0x1b, 0x14, 0xc6, 0x86, 0xdb, 0x8e, 0x95, 0x31, 0x18, 0x0e,
	0xbb, 0xd9, 0x9d, 0xe2, 0xa8, 0xb6, 0x69, 0x23, 0x19, 0x49, 0xde, 0xd6, 0x37, 0xda, 0xa3, 0xf4,
	0x11, 0xc6, 0x2e, 0x02, 0xdb, 0x23, 0xf4, 0x09, 0x86, 0xe4, 0xba, 0x8b, 0x57, 0x18, 0xa3, 0xf4,
	0xce, 0x47, 0xe7, 0x7c, 0x3f, 0xe7, 0xab, 0xa3, 0x83, 0xe1, 0x3f, 0xfa, 0x85, 0xc6, 0x85, 0xca,
	0x38, 0xf3, 0x73, 0xc1, 0x15, 0x47, 0x6d, 0x75, 0x99, 0x53, 0xb9, 0x85, 0x13, 0x9e, 0xf0, 0xc0,
	0x1c, 0xcd, 0x8b, 0xb3, 0x40, 0x47, 0x26, 0x30, 0x5f, 0x65, 0xe9, 0xd6, 0xe3, 0xdb, 0xb4, 0xd1,
	0x04, 0x52, 0x89, 0x22, 0x56, 0x65, 0x12, 0x7f, 0xb3, 0xc1, 0x39, 0xa9, 0xd8, 0x68, 0x06, 0xad,
	0x94, 0xc8, 0xd4, 0xb5, 0xc6, 0x96, 0xd7, 0x0b, 0x5f, 0x5e, 0xad, 0x46, 0x8d, 0xef, 0xab, 0xd1,
	0x6e, 0x92, 0xa9, 0xb4, 0x98, 0xfb, 0x31, 0x5f, 0x06, 0x4b, 0x2a, 0x93, 0xbd, 0x33, 0x5e, 0xb0,
	0x05, 0xd1, 0x8a, 0x80, 0xb2, 0x24, 0x63, 0x34, 0xd0, 0x2a, 0xff, 0x0d, 0x91, 0xe9, 0xf5, 0x6a,
	0xb4, 0xa9, 0x83, 0x43, 0xbc, 0x87, 0x23, 0x03, 0x43, 0x0b, 0x80, 0x9c, 0x08, 0xca, 0x94, 0xce,
	0xbb, 0x1b, 0x06, 0x7d, 0x7c, 0x3f, 0x74, 0xbf, 0x44, 0x33, 0xb2, 0xa4, 0x87, 0x13, 0x1c, 0xad,
	0x71, 0xd1, 0x1c, 0x1c, 0xfa, 0xa9, 0x6a, 0xd2, 0x7c, 0xa8, 0x26, 0x53, 0x1c, 0xfd, 0xc6, 0xa2,
	0x29, 0xd8, 0x52, 0x11, 0x55, 0x48, 0xb7, 0x35, 0xb6, 0xbc, 0xc1, 0xa4, 0xef, 0x9b, 0x89, 0xfa,
	0x33, 0x73, 0x18, 0xf6, 0x6a, 0x97, 0xbf, 0x29, 0x45, 0x29, 0xf4, 0x32, 0x26, 0x15, 0x61, 0x31,
	0x35, 0xde, 0xda, 0x0f, 0xe5, 0xed, 0x00, 0x47, 0x35, 0x32, 0xda, 0x85, 0x8e, 0x22, 0xf2, 0xfc,
	0x2d, 0xbd, 0x74, 0xed, 0xb1, 0xe5, 0x39, 0xe1, 0xff, 0x7f, 0x28, 0x9e, 0xe1, 0xa8, 0xaa, 0x40,
	0xaf, 0xc0, 0xce, 0x58, 0x5e, 0x28, 0xe9, 0x76, 0xc6, 0x96, 0xd7, 0x9d, 0x3c, 0xf2, 0x75, 0x7b,
	0xbf, 0xda, 0x15, 0x7f, 0x66, 0xb6, 0xe4, 0x0e, 0xe2, 0x39, 0x8e, 0x6e, 0x74, 0xe8, 0x05, 0x74,
	0x78, 0xa1, 0x0c, 0x62, 0xf3, 0x6f, 0x88, 0xfa, 0x58, 0x2a, 0x0d, 0xc2, 0xd0, 0xa6, 0x42, 0x70,
	0xe1, 0x3a, 0xc6, 0x6b, 0xbd, 0xaa, 0x4c, 0xa1, 0x27, 0xd0, 0x52, 0x24, 0x91, 0x2e, 0x8c, 0x9b,
	0x9e, 0x13, 0xa2, 0xeb, 0xd5, 0x68, 0xb0, 0xe6, 0x65, 0xff, 0x29, 0x8e, 0x4c, 0x1e, 0x25, 0xd0,
	0xcd, 0x05, 0x8f, 0xa9, 0x94, 0x66, 0xc4, 0x5d, 0x33, 0xe2, 0x93, 0xfb, 0x8d, 0xb8, 0xd6, 0x61,
	0x1f, 0x47, 0xeb, 0x64, 0xb4, 0xa3, 0x37, 0x80, 0xe6, 0xa7, 0xc7, 0x6e, 0xcf, 0xb8, 0xbe, 0x63,
	0x69, 0x62, 0x1e, 0x5e, 0x57, 0xec, 0xbc, 0x03, 0xbb, 0x5c, 0x0c, 0xd4, 0x85, 0xce, 0x07, 0x76,
	0xce, 0xf8, 0x67, 0x36, 0x6c, 0xe8, 0xe0, 0x48, 0x50, 0xa2, 0xe8, 0x62, 0x68, 0xa1, 0x01, 0xc0,
	0x29, 0x7b, 0x2f, 0x78, 0x22, 0xa8, 0x94, 0xc3, 0x0d, 0xd4, 0x07, 0xe7, 0x88, 0x2f, 0xf3, 0x0b,
	0xaa, 0xd3, 0x4d, 0x04, 0x60, 0xbf, 0x26, 0xd9, 0x05, 0x5d, 0x0c, 0x5b, 0xe1, 0xc1, 0xd5, 0x8f,
	0xed, 0xc6, 0xd7, 0x9f, 0xdb, 0xd6, 0xc7, 0x7f, 0xb8, 0xd0, 0xed, 0xef, 0x62, 0x6e, 0x9b, 0xe7,
	0x98, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x87, 0xdb, 0x08, 0x50, 0x42, 0x04, 0x00, 0x00,
}

func (this *Execution) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Execution)
	if !ok {
		that2, ok := that.(Execution)
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
	if !this.ParentHash.Equal(that1.ParentHash) {
		return false
	}
	if !this.EventHash.Equal(that1.EventHash) {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if !this.InstanceHash.Equal(that1.InstanceHash) {
		return false
	}
	if this.TaskKey != that1.TaskKey {
		return false
	}
	if !this.Inputs.Equal(that1.Inputs) {
		return false
	}
	if !this.Outputs.Equal(that1.Outputs) {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	if len(this.Tags) != len(that1.Tags) {
		return false
	}
	for i := range this.Tags {
		if this.Tags[i] != that1.Tags[i] {
			return false
		}
	}
	if !this.ProcessHash.Equal(that1.ProcessHash) {
		return false
	}
	if this.StepID != that1.StepID {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
