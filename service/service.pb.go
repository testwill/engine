// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package service

import (
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
	// 940 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x97, 0xdf, 0x72, 0xdb, 0x44,
	0x14, 0xc6, 0x9b, 0x48, 0xb6, 0xe3, 0x93, 0xf8, 0x82, 0xbd, 0x5a, 0x3c, 0x50, 0xa9, 0x62, 0x06,
	0x5c, 0x70, 0xec, 0xc6, 0x6e, 0x43, 0x93, 0x4e, 0x80, 0x9a, 0xd2, 0xa1, 0xcc, 0xf0, 0x4f, 0x81,
	0x61, 0xa6, 0x33, 0xbd, 0x58, 0x4b, 0x1b, 0x65, 0x49, 0xb4, 0xab, 0xae, 0x56, 0x1e, 0x7c, 0xcf,
	0x0b, 0xf0, 0x18, 0x3c, 0x02, 0xb7, 0x5c, 0xf5, 0x01, 0xb8, 0xe2, 0x42, 0x33, 0xbc, 0x82, 0x9e,
	0x80, 0xd1, 0xda, 0x4e, 0xe4, 0x58, 0x76, 0xfe, 0x94, 0x1b, 0x7a, 0x67, 0x65, 0xce, 0xf7, 0xfb,
	0xa4, 0x73, 0xbe, 0xa3, 0x8d, 0xa0, 0x11, 0x53, 0x39, 0x62, 0x1e, 0xed, 0x44, 0x52, 0x28, 0x81,
	0x20, 0xa4, 0x71, 0xd0, 0x51, 0xe3, 0x88, 0xc6, 0x4d, 0x27, 0x10, 0x81, 0xe8, 0xea, 0xbf, 0x0f,
	0x93, 0xa3, 0x6e, 0x7e, 0xa5, 0x2f, 0xf4, 0xaf, 0x49, 0xbd, 0xf3, 0x17, 0x86, 0xda, 0xe1, 0x84,
	0x80, 0x02, 0x30, 0x8f, 0x49, 0x7c, 0x8c, 0xc1, 0x5e, 0x6b, 0x6d, 0x0d, 0x0e, 0x5f, 0xa5, 0xd6,
	0xad, 0xbf, 0x53, 0xeb, 0xa3, 0x80, 0xa9, 0xe3, 0x64, 0xd8, 0xf1, 0x44, 0xd8, 0xcd, 0xe1, 0xdb,
	0x47, 0x22, 0xe1, 0x3e, 0x51, 0x4c, 0xf0, 0x2e, 0xe5, 0x01, 0xe3, 0xb4, 0x9b, 0xab, 0x3a, 0x5f,
	0x92, 0xf8, 0x38, 0x4b, 0xad, 0x77, 0xf2, 0x8b, 0x7d, 0x67, 0xdb, 0xb1, 0x47, 0xe4, 0x94, 0xf9,
	0x44, 0xd1, 0x7d, 0x47, 0xd2, 0x97, 0x09, 0x93, 0xd4, 0x77, 0x5c, 0x6d, 0x80, 0xbe, 0x07, 0x23,
	0x66, 0x3e, 0xde, 0xb2, 0xd7, 0x5a, 0xf5, 0xc1, 0xa7, 0x59, 0x6a, 0x3d, 0x9a, 0x88, 0x38, 0x09,
	0xe9, 0xfe, 0x4e, 0xaf, 0x4c, 0xda, 0x8e, 0x24, 0xe3, 0x8a, 0xc4, 0x1e, 0x63, 0xed, 0x90, 0xfc,
	0x72, 0xb0, 0xdb, 0x6f, 0xfb, 0x22, 0x24, 0x8c, 0x3b, 0x6e, 0xce, 0x42, 0x4f, 0xc0, 0xcc, 0xd5,
	0x78, 0x4d, 0x33, 0xef, 0x65, 0xa9, 0xd5, 0x2e, 0x32, 0x2f, 0x41, 0x3a, 0xae, 0x56, 0xa3, 0x67,
	0xb0, 0xe9, 0xd3, 0xd8, 0x93, 0x2c, 0xca, 0x1f, 0x0f, 0xaf, 0x6b, 0xd8, 0x07, 0x59, 0x6a, 0xbd,
	0x57, 0x80, 0xcd, 0xdd, 0x5f, 0x91, 0x51, 0xd4, 0x22, 0x09, 0x0d, 0x4f, 0xf0, 0x23, 0x16, 0x24,
	0x52, 0xf7, 0x0a, 0x6f, 0xd8, 0x6b, 0xad, 0xcd, 0xde, 0x9d, 0xce, 0xf9, 0x80, 0x3a, 0xd3, 0xc6,
	0x77, 0x3e, 0x2f, 0x16, 0x0e, 0xee, 0xe6, 0x8d, 0xcf, 0x52, 0xeb, 0x4e, 0xc1, 0xf3, 0x61, 0x79,
	0x3b, 0xe7, 0x2d, 0xd0, 0x73, 0xa8, 0x28, 0x12, 0x9f, 0xc4, 0xb8, 0x62, 0x1b, 0xad, 0xcd, 0x1e,
	0x2e, 0xf3, 0xfa, 0x81, 0xc4, 0x27, 0x83, 0x0f, 0xb3, 0xd4, 0x7a, 0xbf, 0x80, 0x7f, 0x50, 0xc4,
	0xfb, 0x6c, 0x44, 0xdb, 0xe7, 0x1e, 0x13, 0x24, 0x7a, 0x01, 0x55, 0x3a, 0xa2, 0x5c, 0xc5, 0xb8,
	0xaa, 0xe1, 0x6f, 0x97, 0xc1, 0xbf, 0xc8, 0x2b, 0x16, 0xe8, 0xbb, 0x2b, 0xe8, 0x53, 0x28, 0xe2,
	0xb0, 0xe5, 0xd3, 0x88, 0x72, 0x9f, 0x72, 0x8f, 0xd1, 0x18, 0xd7, 0xb4, 0xc9, 0xed, 0x32, 0x93,
	0x27, 0xb3, 0xba, 0xf1, 0x82, 0xd3, 0xc7, 0x2b, 0x9c, 0xe6, 0xf8, 0xe8, 0x2b, 0x00, 0x49, 0x23,
	0x11, 0x33, 0x25, 0xe4, 0x18, 0xd7, 0xf5, 0xa0, 0x2f, 0xd2, 0xf6, 0x8a, 0x34, 0x11, 0x32, 0x45,
	0xc3, 0x48, 0x8d, 0xdb, 0x89, 0x64, 0x8e, 0x5b, 0x50, 0xa3, 0x67, 0x50, 0x8d, 0x45, 0x22, 0x3d,
	0x8a, 0x1b, 0x9a, 0xb3, 0x93, 0xa5, 0xd6, 0x76, 0x31, 0x7d, 0xfd, 0x4b, 0xe3, 0x37, 0x05, 0x34,
	0x7f, 0x5f, 0x87, 0x8a, 0x6e, 0x22, 0xda, 0x03, 0xe3, 0x84, 0x8e, 0xb1, 0x59, 0x1a, 0xc1, 0xfb,
	0xcb, 0x22, 0x98, 0x6b, 0xd0, 0xa3, 0xb9, 0x5d, 0xb8, 0xa8, 0xdd, 0x59, 0xa6, 0xfd, 0xcf, 0x57,
	0xe0, 0x05, 0x98, 0x3e, 0x51, 0x04, 0x1b, 0x7a, 0x96, 0xef, 0x96, 0xcd, 0xf2, 0x3b, 0x22, 0x49,
	0x48, 0x15, 0x95, 0x0b, 0xcd, 0xef, 0xaf, 0x18, 0xa5, 0xc6, 0x36, 0x7f, 0x33, 0xc0, 0xcc, 0xd3,
	0x3c, 0x6b, 0xd5, 0x46, 0xe9, 0xad, 0x3e, 0xfc, 0x5f, 0xb4, 0x8a, 0x40, 0x95, 0xf1, 0x28, 0x39,
	0xdb, 0xae, 0x6b, 0x36, 0x6b, 0xe5, 0x86, 0x4d, 0xc0, 0xc8, 0x83, 0x9a, 0x48, 0x94, 0xf6, 0xa8,
	0xdd, 0xc4, 0x63, 0xd5, 0x6e, 0xcd, 0xc8, 0xcd, 0x5f, 0x4d, 0xa8, 0x9f, 0x21, 0xde, 0x84, 0xc1,
	0x9c, 0x80, 0x99, 0x37, 0x08, 0x1b, 0x9a, 0xf1, 0x53, 0x96, 0x5a, 0x87, 0xcb, 0x42, 0x5a, 0x76,
	0x54, 0x09, 0x4e, 0xc5, 0xd1, 0xc1, 0xa1, 0x92, 0x8c, 0x07, 0xf6, 0x37, 0x49, 0x38, 0xa4, 0xd2,
	0x1e, 0x08, 0x71, 0x4a, 0x09, 0xb7, 0xbf, 0x1d, 0xfe, 0x4c, 0x3d, 0x65, 0x3f, 0xe6, 0x63, 0xc7,
	0xd5, 0x26, 0x68, 0x1b, 0x36, 0x84, 0xb6, 0x25, 0xa7, 0x7a, 0xf1, 0x37, 0x06, 0x6f, 0x65, 0xa9,
	0xd5, 0x98, 0x5b, 0x7c, 0xf7, 0xac, 0x24, 0x2f, 0x97, 0x34, 0xa2, 0x44, 0x51, 0x5f, 0xbf, 0xc1,
	0x16, 0xcb, 0xf7, 0x1c, 0xf7, 0xac, 0x04, 0x31, 0xa8, 0x0a, 0x6d, 0x89, 0xe1, 0x2a, 0xf3, 0xef,
	0x65, 0xa9, 0xd5, 0x29, 0xf6, 0xfc, 0x5e, 0xf1, 0x61, 0x13, 0xce, 0x5e, 0x26, 0xb4, 0x7d, 0x31,
	0x6b, 0x13, 0x83, 0xe6, 0x9f, 0x06, 0x34, 0xe6, 0x0e, 0x35, 0xf4, 0x35, 0xd4, 0x46, 0xe2, 0x34,
	0x09, 0x69, 0x8c, 0xd7, 0x6c, 0xa3, 0x55, 0x1f, 0xf4, 0xb3, 0xd4, 0xea, 0x2e, 0x1b, 0x69, 0x91,
	0x5e, 0x1c, 0xcd, 0x8c, 0x81, 0x7e, 0x84, 0xcd, 0xe9, 0xcf, 0xa7, 0x52, 0x84, 0x78, 0xbd, 0x14,
	0xd9, 0xbb, 0x0a, 0xb2, 0xc8, 0x41, 0x4f, 0xa1, 0x12, 0x09, 0xa9, 0x62, 0xfd, 0xca, 0x5a, 0xfc,
	0x37, 0xa2, 0xbf, 0x14, 0x28, 0xa4, 0x0a, 0x49, 0xe4, 0xb8, 0x13, 0x39, 0xfa, 0x0c, 0x4c, 0x22,
	0x83, 0x18, 0x9b, 0x1a, 0xd3, 0xce, 0x52, 0xab, 0xb5, 0xf2, 0xb4, 0x9d, 0x8b, 0x70, 0xae, 0x44,
	0x8f, 0xa1, 0xe6, 0x89, 0x30, 0x24, 0xdc, 0xc7, 0x95, 0xeb, 0x1d, 0x01, 0x33, 0x1d, 0xfa, 0x04,
	0x0c, 0xca, 0x47, 0xfa, 0x85, 0xb2, 0x78, 0x0f, 0xbb, 0xcb, 0x1e, 0x85, 0xf2, 0x91, 0xe3, 0xe6,
	0xc2, 0xe6, 0x1f, 0x26, 0xc0, 0xf9, 0x59, 0xfb, 0x3a, 0xcb, 0x7c, 0x00, 0x15, 0x16, 0x92, 0xe0,
	0xda, 0xdb, 0x3c, 0x51, 0x15, 0xb3, 0xf3, 0x1a, 0x83, 0x5e, 0x96, 0x1d, 0xa3, 0x14, 0xd9, 0xbf,
	0x79, 0x76, 0xcc, 0xd2, 0xec, 0xdc, 0xbf, 0x6e, 0x76, 0xae, 0x30, 0xb7, 0x9b, 0x66, 0xe7, 0xc1,
	0x55, 0xb3, 0x53, 0x2f, 0xbd, 0x87, 0xbd, 0x4b, 0xb3, 0x33, 0xe8, 0xbe, 0xfa, 0xe7, 0xf6, 0xad,
	0xe7, 0x77, 0x2f, 0xff, 0x74, 0x98, 0x7e, 0xbd, 0x0c, 0xab, 0xfa, 0x73, 0xa4, 0xff, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x2b, 0x25, 0x13, 0xd6, 0xcf, 0x0c, 0x00, 0x00,
}
