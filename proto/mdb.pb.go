// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: mdb.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MachineType int32

const (
	MachineType_MACHINE_TYPE_UNKNOWN      MachineType = 0
	MachineType_MACHINE_TYPE_RASPBERRY_PI MachineType = 1
	MachineType_MACHINE_TYPE_IOT_DEVICE   MachineType = 2
)

// Enum value maps for MachineType.
var (
	MachineType_name = map[int32]string{
		0: "MACHINE_TYPE_UNKNOWN",
		1: "MACHINE_TYPE_RASPBERRY_PI",
		2: "MACHINE_TYPE_IOT_DEVICE",
	}
	MachineType_value = map[string]int32{
		"MACHINE_TYPE_UNKNOWN":      0,
		"MACHINE_TYPE_RASPBERRY_PI": 1,
		"MACHINE_TYPE_IOT_DEVICE":   2,
	}
)

func (x MachineType) Enum() *MachineType {
	p := new(MachineType)
	*p = x
	return p
}

func (x MachineType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MachineType) Descriptor() protoreflect.EnumDescriptor {
	return file_mdb_proto_enumTypes[0].Descriptor()
}

func (MachineType) Type() protoreflect.EnumType {
	return &file_mdb_proto_enumTypes[0]
}

func (x MachineType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MachineType.Descriptor instead.
func (MachineType) EnumDescriptor() ([]byte, []int) {
	return file_mdb_proto_rawDescGZIP(), []int{0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentMachine *Machine `protobuf:"bytes,1,opt,name=current_machine,json=currentMachine,proto3" json:"current_machine,omitempty"`
	IssueId        int32    `protobuf:"varint,2,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_mdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_mdb_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetCurrentMachine() *Machine {
	if x != nil {
		return x.CurrentMachine
	}
	return nil
}

func (x *Config) GetIssueId() int32 {
	if x != nil {
		return x.IssueId
	}
	return 0
}

type Mdb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machines []*Machine `protobuf:"bytes,1,rep,name=machines,proto3" json:"machines,omitempty"`
	Config   *Config    `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Mdb) Reset() {
	*x = Mdb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mdb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mdb) ProtoMessage() {}

func (x *Mdb) ProtoReflect() protoreflect.Message {
	mi := &file_mdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mdb.ProtoReflect.Descriptor instead.
func (*Mdb) Descriptor() ([]byte, []int) {
	return file_mdb_proto_rawDescGZIP(), []int{1}
}

func (x *Mdb) GetMachines() []*Machine {
	if x != nil {
		return x.Machines
	}
	return nil
}

func (x *Mdb) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type Machine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ipv4       uint32 `protobuf:"fixed32,1,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Hostname   string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Mac        string `protobuf:"bytes,3,opt,name=mac,proto3" json:"mac,omitempty"`
	Controller string `protobuf:"bytes,4,opt,name=controller,proto3" json:"controller,omitempty"`
	// These are the user specified eleemnts
	Type MachineType `protobuf:"varint,5,opt,name=type,proto3,enum=mdb.MachineType" json:"type,omitempty"`
}

func (x *Machine) Reset() {
	*x = Machine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Machine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Machine) ProtoMessage() {}

func (x *Machine) ProtoReflect() protoreflect.Message {
	mi := &file_mdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Machine.ProtoReflect.Descriptor instead.
func (*Machine) Descriptor() ([]byte, []int) {
	return file_mdb_proto_rawDescGZIP(), []int{2}
}

func (x *Machine) GetIpv4() uint32 {
	if x != nil {
		return x.Ipv4
	}
	return 0
}

func (x *Machine) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Machine) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *Machine) GetController() string {
	if x != nil {
		return x.Controller
	}
	return ""
}

func (x *Machine) GetType() MachineType {
	if x != nil {
		return x.Type
	}
	return MachineType_MACHINE_TYPE_UNKNOWN
}

type ListMachinesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ipv4 uint32 `protobuf:"fixed32,1,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
}

func (x *ListMachinesRequest) Reset() {
	*x = ListMachinesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMachinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMachinesRequest) ProtoMessage() {}

func (x *ListMachinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMachinesRequest.ProtoReflect.Descriptor instead.
func (*ListMachinesRequest) Descriptor() ([]byte, []int) {
	return file_mdb_proto_rawDescGZIP(), []int{3}
}

func (x *ListMachinesRequest) GetIpv4() uint32 {
	if x != nil {
		return x.Ipv4
	}
	return 0
}

type ListMachinesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machines []*Machine `protobuf:"bytes,1,rep,name=machines,proto3" json:"machines,omitempty"`
}

func (x *ListMachinesResponse) Reset() {
	*x = ListMachinesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMachinesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMachinesResponse) ProtoMessage() {}

func (x *ListMachinesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMachinesResponse.ProtoReflect.Descriptor instead.
func (*ListMachinesResponse) Descriptor() ([]byte, []int) {
	return file_mdb_proto_rawDescGZIP(), []int{4}
}

func (x *ListMachinesResponse) GetMachines() []*Machine {
	if x != nil {
		return x.Machines
	}
	return nil
}

var File_mdb_proto protoreflect.FileDescriptor

var file_mdb_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x64, 0x62,
	0x22, 0x5a, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x35, 0x0a, 0x0f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x03,
	0x4d, 0x64, 0x62, 0x12, 0x28, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x52, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x23, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x6d, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x91, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x04, 0x69, 0x70,
	0x76, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x04, 0x69, 0x70, 0x76,
	0x34, 0x22, 0x40, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x64,
	0x62, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x73, 0x2a, 0x63, 0x0a, 0x0b, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19,
	0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x41, 0x53,
	0x50, 0x42, 0x45, 0x52, 0x52, 0x59, 0x5f, 0x50, 0x49, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4d,
	0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4f, 0x54, 0x5f,
	0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0x02, 0x32, 0x53, 0x0a, 0x0a, 0x4d, 0x44, 0x42, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mdb_proto_rawDescOnce sync.Once
	file_mdb_proto_rawDescData = file_mdb_proto_rawDesc
)

func file_mdb_proto_rawDescGZIP() []byte {
	file_mdb_proto_rawDescOnce.Do(func() {
		file_mdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_mdb_proto_rawDescData)
	})
	return file_mdb_proto_rawDescData
}

var file_mdb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mdb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mdb_proto_goTypes = []interface{}{
	(MachineType)(0),             // 0: mdb.MachineType
	(*Config)(nil),               // 1: mdb.Config
	(*Mdb)(nil),                  // 2: mdb.Mdb
	(*Machine)(nil),              // 3: mdb.Machine
	(*ListMachinesRequest)(nil),  // 4: mdb.ListMachinesRequest
	(*ListMachinesResponse)(nil), // 5: mdb.ListMachinesResponse
}
var file_mdb_proto_depIdxs = []int32{
	3, // 0: mdb.Config.current_machine:type_name -> mdb.Machine
	3, // 1: mdb.Mdb.machines:type_name -> mdb.Machine
	1, // 2: mdb.Mdb.config:type_name -> mdb.Config
	0, // 3: mdb.Machine.type:type_name -> mdb.MachineType
	3, // 4: mdb.ListMachinesResponse.machines:type_name -> mdb.Machine
	4, // 5: mdb.MDBService.ListMachines:input_type -> mdb.ListMachinesRequest
	5, // 6: mdb.MDBService.ListMachines:output_type -> mdb.ListMachinesResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_mdb_proto_init() }
func file_mdb_proto_init() {
	if File_mdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mdb); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Machine); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMachinesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMachinesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mdb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mdb_proto_goTypes,
		DependencyIndexes: file_mdb_proto_depIdxs,
		EnumInfos:         file_mdb_proto_enumTypes,
		MessageInfos:      file_mdb_proto_msgTypes,
	}.Build()
	File_mdb_proto = out.File
	file_mdb_proto_rawDesc = nil
	file_mdb_proto_goTypes = nil
	file_mdb_proto_depIdxs = nil
}
