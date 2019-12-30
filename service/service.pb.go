// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package service

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

// Service represents the service's type.
type Service struct {
	// Service's hash.
	Hash github_com_mesg_foundation_engine_hash.Hash `protobuf:"bytes,10,opt,name=hash,proto3,customtype=github.com/mesg-foundation/engine/hash.Hash" json:"hash" hash:"-" validate:"required"`
	// Service's sid.
	Sid string `protobuf:"bytes,12,opt,name=sid,proto3" json:"sid,omitempty" hash:"name:12" validate:"required,printascii,max=63,domain"`
	// Service's name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" hash:"name:1" validate:"required,printascii"`
	// Service's description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" hash:"name:2" validate:"printascii"`
	// Configurations related to the service
	Configuration Service_Configuration `protobuf:"bytes,8,opt,name=configuration,proto3" json:"configuration" hash:"name:8" validate:"required"`
	// The list of tasks this service can execute.
	Tasks []*Service_Task `protobuf:"bytes,5,rep,name=tasks,proto3" json:"tasks,omitempty" hash:"name:5" validate:"dive,required"`
	// The list of events this service can emit.
	Events []*Service_Event `protobuf:"bytes,6,rep,name=events,proto3" json:"events,omitempty" hash:"name:6" validate:"dive,required"`
	// The container dependencies this service requires.
	Dependencies []*Service_Dependency `protobuf:"bytes,7,rep,name=dependencies,proto3" json:"dependencies,omitempty" hash:"name:7" validate:"dive,required"`
	// Service's repository url.
	Repository string `protobuf:"bytes,9,opt,name=repository,proto3" json:"repository,omitempty" hash:"name:9" validate:"omitempty,uri"`
	// The hash id of service's source code on IPFS.
	Source               string   `protobuf:"bytes,13,opt,name=source,proto3" json:"source,omitempty" hash:"name:13" validate:"required,printascii"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

// Events are emitted by the service whenever the service wants.
type Service_Event struct {
	// Event's key.
	Key string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty" hash:"name:4" validate:"printascii"`
	// Event's name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" hash:"name:1" validate:"printascii"`
	// Event's description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" hash:"name:2" validate:"printascii"`
	// List of data of this event.
	Data                 []*Service_Parameter `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" hash:"name:3" validate:"dive,required"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Service_Event) Reset()         { *m = Service_Event{} }
func (m *Service_Event) String() string { return proto.CompactTextString(m) }
func (*Service_Event) ProtoMessage()    {}
func (*Service_Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 0}
}
func (m *Service_Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Event.Unmarshal(m, b)
}
func (m *Service_Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Event.Marshal(b, m, deterministic)
}
func (m *Service_Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Event.Merge(m, src)
}
func (m *Service_Event) XXX_Size() int {
	return xxx_messageInfo_Service_Event.Size(m)
}
func (m *Service_Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Event proto.InternalMessageInfo

// Task is a function that requires inputs and returns output.
type Service_Task struct {
	// Task's key.
	Key string `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty" hash:"name:8" validate:"printascii"`
	// Task's name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" hash:"name:1" validate:"printascii"`
	// Task's description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" hash:"name:2" validate:"printascii"`
	// List inputs of this task.
	Inputs []*Service_Parameter `protobuf:"bytes,6,rep,name=inputs,proto3" json:"inputs,omitempty" hash:"name:6" validate:"dive,required"`
	// List of tasks outputs.
	Outputs              []*Service_Parameter `protobuf:"bytes,7,rep,name=outputs,proto3" json:"outputs,omitempty" hash:"name:7" validate:"dive,required"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Service_Task) Reset()         { *m = Service_Task{} }
func (m *Service_Task) String() string { return proto.CompactTextString(m) }
func (*Service_Task) ProtoMessage()    {}
func (*Service_Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 1}
}
func (m *Service_Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Task.Unmarshal(m, b)
}
func (m *Service_Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Task.Marshal(b, m, deterministic)
}
func (m *Service_Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Task.Merge(m, src)
}
func (m *Service_Task) XXX_Size() int {
	return xxx_messageInfo_Service_Task.Size(m)
}
func (m *Service_Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Task proto.InternalMessageInfo

// Parameter describes the task's inputs, the task's outputs, and the event's data.
type Service_Parameter struct {
	// Parameter's key.
	Key string `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty" hash:"name:8" validate:"printascii"`
	// Parameter's name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" hash:"name:1" validate:"printascii"`
	// Parameter's description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" hash:"name:2" validate:"printascii"`
	// Parameter's type: `String`, `Number`, `Boolean`, `Object` or `Any`.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" hash:"name:3" validate:"required,printascii,oneof=String Number Boolean Object Any"`
	// Set the parameter as optional.
	Optional bool `protobuf:"varint,4,opt,name=optional,proto3" json:"optional,omitempty" hash:"name:4"`
	// Mark a parameter as an array of the defined type.
	Repeated bool `protobuf:"varint,9,opt,name=repeated,proto3" json:"repeated,omitempty" hash:"name:9"`
	// Optional object structure type when type is set to `Object`.
	Object               []*Service_Parameter `protobuf:"bytes,10,rep,name=object,proto3" json:"object,omitempty" hash:"name:10" validate:"unique,dive,required"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Service_Parameter) Reset()         { *m = Service_Parameter{} }
func (m *Service_Parameter) String() string { return proto.CompactTextString(m) }
func (*Service_Parameter) ProtoMessage()    {}
func (*Service_Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 2}
}
func (m *Service_Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Parameter.Unmarshal(m, b)
}
func (m *Service_Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Parameter.Marshal(b, m, deterministic)
}
func (m *Service_Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Parameter.Merge(m, src)
}
func (m *Service_Parameter) XXX_Size() int {
	return xxx_messageInfo_Service_Parameter.Size(m)
}
func (m *Service_Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Parameter proto.InternalMessageInfo

// A configuration is the configuration of the main container of the service's instance.
type Service_Configuration struct {
	// List of volumes.
	Volumes []string `protobuf:"bytes,1,rep,name=volumes,proto3" json:"volumes,omitempty" hash:"name:1" validate:"unique,dive,printascii"`
	// List of volumes mounted from other dependencies.
	VolumesFrom []string `protobuf:"bytes,2,rep,name=volumesFrom,proto3" json:"volumesFrom,omitempty" hash:"name:2" validate:"unique,dive,printascii"`
	// List of ports the container exposes.
	Ports []string `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty" hash:"name:3" validate:"unique,dive,portmap"`
	// Args to pass to the container.
	Args []string `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty" hash:"name:5" validate:"dive,printascii"`
	// Command to run the container.
	Command string `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty" hash:"name:4" validate:"printascii"`
	// Default env vars to apply to service's instance on runtime.
	Env                  []string `protobuf:"bytes,6,rep,name=env,proto3" json:"env,omitempty" hash:"name:6" validate:"unique,dive,env"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service_Configuration) Reset()         { *m = Service_Configuration{} }
func (m *Service_Configuration) String() string { return proto.CompactTextString(m) }
func (*Service_Configuration) ProtoMessage()    {}
func (*Service_Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 3}
}
func (m *Service_Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Configuration.Unmarshal(m, b)
}
func (m *Service_Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Configuration.Marshal(b, m, deterministic)
}
func (m *Service_Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Configuration.Merge(m, src)
}
func (m *Service_Configuration) XXX_Size() int {
	return xxx_messageInfo_Service_Configuration.Size(m)
}
func (m *Service_Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Configuration proto.InternalMessageInfo

// A dependency is a configuration of an other container that runs separately from the service.
type Service_Dependency struct {
	// Dependency's key.
	Key string `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty" hash:"name:8" validate:"printascii"`
	// Image's name of the container.
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty" hash:"name:1" validate:"printascii"`
	// List of volumes.
	Volumes []string `protobuf:"bytes,2,rep,name=volumes,proto3" json:"volumes,omitempty" hash:"name:2" validate:"unique,dive,printascii"`
	// List of volumes mounted from other dependencies.
	VolumesFrom []string `protobuf:"bytes,3,rep,name=volumesFrom,proto3" json:"volumesFrom,omitempty" hash:"name:3" validate:"unique,dive,printascii"`
	// List of ports the container exposes.
	Ports []string `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty" hash:"name:4" validate:"unique,dive,portmap"`
	// Args to pass to the container.
	Args []string `protobuf:"bytes,6,rep,name=args,proto3" json:"args,omitempty" hash:"name:6" validate:"dive,printascii"`
	// Command to run the container.
	Command string `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty" hash:"name:5" validate:"printascii"`
	// Default env vars to apply to service's instance on runtime.
	Env                  []string `protobuf:"bytes,9,rep,name=env,proto3" json:"env,omitempty" hash:"name:9" validate:"unique,dive,env"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service_Dependency) Reset()         { *m = Service_Dependency{} }
func (m *Service_Dependency) String() string { return proto.CompactTextString(m) }
func (*Service_Dependency) ProtoMessage()    {}
func (*Service_Dependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 4}
}
func (m *Service_Dependency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service_Dependency.Unmarshal(m, b)
}
func (m *Service_Dependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service_Dependency.Marshal(b, m, deterministic)
}
func (m *Service_Dependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service_Dependency.Merge(m, src)
}
func (m *Service_Dependency) XXX_Size() int {
	return xxx_messageInfo_Service_Dependency.Size(m)
}
func (m *Service_Dependency) XXX_DiscardUnknown() {
	xxx_messageInfo_Service_Dependency.DiscardUnknown(m)
}

var xxx_messageInfo_Service_Dependency proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Service)(nil), "mesg.types.Service")
	proto.RegisterType((*Service_Event)(nil), "mesg.types.Service.Event")
	proto.RegisterType((*Service_Task)(nil), "mesg.types.Service.Task")
	proto.RegisterType((*Service_Parameter)(nil), "mesg.types.Service.Parameter")
	proto.RegisterType((*Service_Configuration)(nil), "mesg.types.Service.Configuration")
	proto.RegisterType((*Service_Dependency)(nil), "mesg.types.Service.Dependency")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 945 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x97, 0xdf, 0x72, 0xdb, 0x44,
	0x14, 0xc6, 0xeb, 0x48, 0xb6, 0xe3, 0x93, 0xf8, 0x82, 0xbd, 0x5a, 0x3c, 0x10, 0xa9, 0x62, 0x06,
	0x5c, 0x70, 0xec, 0xc6, 0x6e, 0x43, 0x93, 0x4e, 0x80, 0x9a, 0xd2, 0xa1, 0xcc, 0xf0, 0xcf, 0x81,
	0x61, 0xa6, 0x33, 0xbd, 0x58, 0x4b, 0x1b, 0x65, 0x49, 0xb4, 0xab, 0xae, 0x56, 0x1e, 0x7c, 0xcf,
	0x0b, 0xf0, 0x16, 0xf0, 0x08, 0xdc, 0x72, 0xd5, 0x07, 0xe0, 0x8a, 0x0b, 0xcd, 0xc0, 0x23, 0xe8,
	0x09, 0x18, 0xad, 0xed, 0x44, 0x8e, 0x65, 0xe7, 0x4f, 0xb9, 0xa1, 0x77, 0x56, 0xe6, 0x7c, 0xbf,
	0x4f, 0x3a, 0xe7, 0x3b, 0xda, 0x08, 0xea, 0x11, 0x95, 0x23, 0xe6, 0xd2, 0x76, 0x28, 0x85, 0x12,
	0x08, 0x02, 0x1a, 0xf9, 0x6d, 0x35, 0x0e, 0x69, 0xd4, 0x70, 0x7c, 0xe1, 0x8b, 0x8e, 0xfe, 0xfb,
	0x30, 0x3e, 0xea, 0x64, 0x57, 0xfa, 0x42, 0xff, 0x9a, 0xd4, 0x3b, 0x7f, 0x62, 0xa8, 0x1e, 0x4e,
	0x08, 0xc8, 0x07, 0xf3, 0x98, 0x44, 0xc7, 0x18, 0xec, 0x52, 0x73, 0xb3, 0x7f, 0xf8, 0x32, 0xb1,
	0x6e, 0xfd, 0x95, 0x58, 0x1f, 0xf8, 0x4c, 0x1d, 0xc7, 0xc3, 0xb6, 0x2b, 0x82, 0x4e, 0x06, 0xdf,
	0x3e, 0x12, 0x31, 0xf7, 0x88, 0x62, 0x82, 0x77, 0x28, 0xf7, 0x19, 0xa7, 0x9d, 0x4c, 0xd5, 0xfe,
	0x9c, 0x44, 0xc7, 0x69, 0x62, 0xbd, 0x95, 0x5d, 0xec, 0x3b, 0xdb, 0x8e, 0x3d, 0x22, 0xa7, 0xcc,
	0x23, 0x8a, 0xee, 0x3b, 0x92, 0xbe, 0x88, 0x99, 0xa4, 0x9e, 0x33, 0xd0, 0x06, 0xe8, 0x5b, 0x30,
	0x22, 0xe6, 0xe1, 0x4d, 0xbb, 0xd4, 0xac, 0xf5, 0x3f, 0x4e, 0x13, 0xeb, 0xe1, 0x44, 0xc4, 0x49,
	0x40, 0xf7, 0x77, 0xba, 0x45, 0xd2, 0x56, 0x28, 0x19, 0x57, 0x24, 0x72, 0x19, 0x6b, 0x05, 0xe4,
	0xa7, 0x83, 0xdd, 0x5e, 0xcb, 0x13, 0x01, 0x61, 0xdc, 0x19, 0x64, 0x2c, 0xf4, 0x18, 0xcc, 0x4c,
	0x8d, 0x4b, 0x9a, 0x79, 0x37, 0x4d, 0xac, 0x56, 0x9e, 0x79, 0x09, 0xd2, 0x19, 0x68, 0x35, 0x7a,
	0x0a, 0x1b, 0x1e, 0x8d, 0x5c, 0xc9, 0xc2, 0xec, 0xf1, 0xf0, 0x9a, 0x86, 0xbd, 0x97, 0x26, 0xd6,
	0x3b, 0x39, 0xd8, 0xdc, 0xfd, 0xe5, 0x19, 0x79, 0x2d, 0x92, 0x50, 0x77, 0x05, 0x3f, 0x62, 0x7e,
	0x2c, 0x75, 0xaf, 0xf0, 0xba, 0x5d, 0x6a, 0x6e, 0x74, 0x6f, 0xb7, 0xcf, 0x07, 0xd4, 0x9e, 0x36,
	0xbe, 0xfd, 0x69, 0xbe, 0xb0, 0x7f, 0x27, 0x6b, 0x7c, 0x9a, 0x58, 0xb7, 0x73, 0x9e, 0x0f, 0x8a,
	0xdb, 0x39, 0x6f, 0x81, 0x9e, 0x41, 0x59, 0x91, 0xe8, 0x24, 0xc2, 0x65, 0xdb, 0x68, 0x6e, 0x74,
	0x71, 0x91, 0xd7, 0x77, 0x24, 0x3a, 0xe9, 0xbf, 0x9f, 0x26, 0xd6, 0xbb, 0x39, 0xfc, 0xfd, 0x3c,
	0xde, 0x63, 0x23, 0xda, 0x3a, 0xf7, 0x98, 0x20, 0xd1, 0x73, 0xa8, 0xd0, 0x11, 0xe5, 0x2a, 0xc2,
	0x15, 0x0d, 0x7f, 0xb3, 0x08, 0xfe, 0x59, 0x56, 0xb1, 0x40, 0xdf, 0x5d, 0x41, 0x9f, 0x42, 0x11,
	0x87, 0x4d, 0x8f, 0x86, 0x94, 0x7b, 0x94, 0xbb, 0x8c, 0x46, 0xb8, 0xaa, 0x4d, 0xb6, 0x8a, 0x4c,
	0x1e, 0xcf, 0xea, 0xc6, 0x0b, 0x4e, 0x1f, 0xae, 0x70, 0x9a, 0xe3, 0xa3, 0x2f, 0x00, 0x24, 0x0d,
	0x45, 0xc4, 0x94, 0x90, 0x63, 0x5c, 0xd3, 0x83, 0xbe, 0x48, 0xdb, 0xcb, 0xd3, 0x44, 0xc0, 0x14,
	0x0d, 0x42, 0x35, 0x6e, 0xc5, 0x92, 0x39, 0x83, 0x9c, 0x1a, 0x3d, 0x85, 0x4a, 0x24, 0x62, 0xe9,
	0x52, 0x5c, 0xd7, 0x9c, 0x9d, 0x34, 0xb1, 0xb6, 0xf3, 0xe9, 0xeb, 0x5d, 0x1a, 0xbf, 0x29, 0xa0,
	0xf1, 0xdb, 0x1a, 0x94, 0x75, 0x13, 0xd1, 0x1e, 0x18, 0x27, 0x74, 0x8c, 0xcd, 0xc2, 0x08, 0xde,
	0x5b, 0x16, 0xc1, 0x4c, 0x83, 0x1e, 0xce, 0xed, 0xc2, 0x45, 0xed, 0xce, 0x32, 0xed, 0x7f, 0xbe,
	0x02, 0xcf, 0xc1, 0xf4, 0x88, 0x22, 0xd8, 0xd0, 0xb3, 0x7c, 0xbb, 0x68, 0x96, 0xdf, 0x10, 0x49,
	0x02, 0xaa, 0xa8, 0x5c, 0x68, 0x7e, 0x6f, 0xc5, 0x28, 0x35, 0xb6, 0xf1, 0x8b, 0x01, 0x66, 0x96,
	0xe6, 0x59, 0xab, 0xd6, 0x0b, 0x6f, 0xf5, 0xc1, 0xff, 0xa2, 0x55, 0x04, 0x2a, 0x8c, 0x87, 0xf1,
	0xd9, 0x76, 0x5d, 0xb3, 0x59, 0x2b, 0x37, 0x6c, 0x02, 0x46, 0x2e, 0x54, 0x45, 0xac, 0xb4, 0x47,
	0xf5, 0x26, 0x1e, 0xab, 0x76, 0x6b, 0x46, 0x6e, 0xfc, 0x6c, 0x42, 0xed, 0x0c, 0xf1, 0x3a, 0x0c,
	0xe6, 0x04, 0xcc, 0xac, 0x41, 0xd8, 0xd0, 0x8c, 0x1f, 0xd2, 0xc4, 0x3a, 0x5c, 0x16, 0xd2, 0xa2,
	0xa3, 0x4a, 0x70, 0x2a, 0x8e, 0x0e, 0x0e, 0x95, 0x64, 0xdc, 0xb7, 0xbf, 0x8a, 0x83, 0x21, 0x95,
	0x76, 0x5f, 0x88, 0x53, 0x4a, 0xb8, 0xfd, 0xf5, 0xf0, 0x47, 0xea, 0x2a, 0xfb, 0x11, 0x1f, 0x3b,
	0x03, 0x6d, 0x82, 0xb6, 0x61, 0x5d, 0x68, 0x5b, 0x72, 0xaa, 0x17, 0x7f, 0xbd, 0xff, 0x46, 0x9a,
	0x58, 0xf5, 0xb9, 0xc5, 0x1f, 0x9c, 0x95, 0x64, 0xe5, 0x92, 0x86, 0x94, 0x28, 0xea, 0xe9, 0x37,
	0xd8, 0x62, 0xf9, 0x9e, 0x33, 0x38, 0x2b, 0x41, 0x0c, 0x2a, 0x42, 0x5b, 0x62, 0xb8, 0xca, 0xfc,
	0xbb, 0x69, 0x62, 0xb5, 0xf3, 0x3d, 0xbf, 0x9b, 0x7f, 0xd8, 0x98, 0xb3, 0x17, 0x31, 0x6d, 0x5d,
	0xcc, 0xda, 0xc4, 0xa0, 0xf1, 0x87, 0x01, 0xf5, 0xb9, 0x43, 0x0d, 0x7d, 0x09, 0xd5, 0x91, 0x38,
	0x8d, 0x03, 0x1a, 0xe1, 0x92, 0x6d, 0x34, 0x6b, 0xfd, 0x5e, 0x9a, 0x58, 0x9d, 0x65, 0x23, 0xcd,
	0xd3, 0xf3, 0xa3, 0x99, 0x31, 0xd0, 0xf7, 0xb0, 0x31, 0xfd, 0xf9, 0x44, 0x8a, 0x00, 0xaf, 0x15,
	0x22, 0xbb, 0x57, 0x41, 0xe6, 0x39, 0xe8, 0x09, 0x94, 0x43, 0x21, 0x55, 0xa4, 0x5f, 0x59, 0x8b,
	0xff, 0x46, 0xf4, 0x96, 0x02, 0x85, 0x54, 0x01, 0x09, 0x9d, 0xc1, 0x44, 0x8e, 0x3e, 0x01, 0x93,
	0x48, 0x3f, 0xc2, 0xa6, 0xc6, 0xb4, 0xd2, 0xc4, 0x6a, 0xae, 0x3c, 0x6d, 0xe7, 0x22, 0x9c, 0x29,
	0xd1, 0x23, 0xa8, 0xba, 0x22, 0x08, 0x08, 0xf7, 0x70, 0xf9, 0x7a, 0x47, 0xc0, 0x4c, 0x87, 0x3e,
	0x02, 0x83, 0xf2, 0x91, 0x7e, 0xa1, 0x2c, 0xde, 0xc3, 0xee, 0xb2, 0x47, 0xa1, 0x7c, 0xe4, 0x0c,
	0x32, 0x61, 0xe3, 0x77, 0x13, 0xe0, 0xfc, 0xac, 0x7d, 0x95, 0x65, 0x3e, 0x80, 0x32, 0x0b, 0x88,
	0x7f, 0xed, 0x6d, 0x9e, 0xa8, 0xf2, 0xd9, 0x79, 0x85, 0x41, 0x2f, 0xcb, 0x8e, 0x51, 0x88, 0xec,
	0xdd, 0x3c, 0x3b, 0x66, 0x61, 0x76, 0xee, 0x5d, 0x37, 0x3b, 0x57, 0x98, 0xdb, 0x4d, 0xb3, 0x73,
	0xff, 0xaa, 0xd9, 0xa9, 0x15, 0xde, 0xc3, 0xde, 0xa5, 0xd9, 0xe9, 0xf7, 0x5e, 0xfe, 0xbd, 0x75,
	0xeb, 0xd7, 0x7f, 0xb6, 0x4a, 0xcf, 0xee, 0x5c, 0xfe, 0xf9, 0x30, 0xfd, 0x82, 0x19, 0x56, 0xf4,
	0x27, 0x49, 0xef, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xba, 0xa3, 0x55, 0xd3, 0x0c, 0x00,
	0x00,
}

func (this *Service) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Service)
	if !ok {
		that2, ok := that.(Service)
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
	if this.Sid != that1.Sid {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if !this.Configuration.Equal(&that1.Configuration) {
		return false
	}
	if len(this.Tasks) != len(that1.Tasks) {
		return false
	}
	for i := range this.Tasks {
		if !this.Tasks[i].Equal(that1.Tasks[i]) {
			return false
		}
	}
	if len(this.Events) != len(that1.Events) {
		return false
	}
	for i := range this.Events {
		if !this.Events[i].Equal(that1.Events[i]) {
			return false
		}
	}
	if len(this.Dependencies) != len(that1.Dependencies) {
		return false
	}
	for i := range this.Dependencies {
		if !this.Dependencies[i].Equal(that1.Dependencies[i]) {
			return false
		}
	}
	if this.Repository != that1.Repository {
		return false
	}
	if this.Source != that1.Source {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Service_Event) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Service_Event)
	if !ok {
		that2, ok := that.(Service_Event)
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
	if this.Key != that1.Key {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Data) != len(that1.Data) {
		return false
	}
	for i := range this.Data {
		if !this.Data[i].Equal(that1.Data[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Service_Task) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Service_Task)
	if !ok {
		that2, ok := that.(Service_Task)
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
	if this.Key != that1.Key {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if len(this.Inputs) != len(that1.Inputs) {
		return false
	}
	for i := range this.Inputs {
		if !this.Inputs[i].Equal(that1.Inputs[i]) {
			return false
		}
	}
	if len(this.Outputs) != len(that1.Outputs) {
		return false
	}
	for i := range this.Outputs {
		if !this.Outputs[i].Equal(that1.Outputs[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Service_Parameter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Service_Parameter)
	if !ok {
		that2, ok := that.(Service_Parameter)
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
	if this.Key != that1.Key {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Optional != that1.Optional {
		return false
	}
	if this.Repeated != that1.Repeated {
		return false
	}
	if len(this.Object) != len(that1.Object) {
		return false
	}
	for i := range this.Object {
		if !this.Object[i].Equal(that1.Object[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Service_Configuration) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Service_Configuration)
	if !ok {
		that2, ok := that.(Service_Configuration)
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
	if len(this.Volumes) != len(that1.Volumes) {
		return false
	}
	for i := range this.Volumes {
		if this.Volumes[i] != that1.Volumes[i] {
			return false
		}
	}
	if len(this.VolumesFrom) != len(that1.VolumesFrom) {
		return false
	}
	for i := range this.VolumesFrom {
		if this.VolumesFrom[i] != that1.VolumesFrom[i] {
			return false
		}
	}
	if len(this.Ports) != len(that1.Ports) {
		return false
	}
	for i := range this.Ports {
		if this.Ports[i] != that1.Ports[i] {
			return false
		}
	}
	if len(this.Args) != len(that1.Args) {
		return false
	}
	for i := range this.Args {
		if this.Args[i] != that1.Args[i] {
			return false
		}
	}
	if this.Command != that1.Command {
		return false
	}
	if len(this.Env) != len(that1.Env) {
		return false
	}
	for i := range this.Env {
		if this.Env[i] != that1.Env[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Service_Dependency) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Service_Dependency)
	if !ok {
		that2, ok := that.(Service_Dependency)
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
	if this.Key != that1.Key {
		return false
	}
	if this.Image != that1.Image {
		return false
	}
	if len(this.Volumes) != len(that1.Volumes) {
		return false
	}
	for i := range this.Volumes {
		if this.Volumes[i] != that1.Volumes[i] {
			return false
		}
	}
	if len(this.VolumesFrom) != len(that1.VolumesFrom) {
		return false
	}
	for i := range this.VolumesFrom {
		if this.VolumesFrom[i] != that1.VolumesFrom[i] {
			return false
		}
	}
	if len(this.Ports) != len(that1.Ports) {
		return false
	}
	for i := range this.Ports {
		if this.Ports[i] != that1.Ports[i] {
			return false
		}
	}
	if len(this.Args) != len(that1.Args) {
		return false
	}
	for i := range this.Args {
		if this.Args[i] != that1.Args[i] {
			return false
		}
	}
	if this.Command != that1.Command {
		return false
	}
	if len(this.Env) != len(that1.Env) {
		return false
	}
	for i := range this.Env {
		if this.Env[i] != that1.Env[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}