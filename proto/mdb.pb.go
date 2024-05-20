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
	MachineType_MACHINE_TYPE_INTEL        MachineType = 3
	MachineType_MACHINE_TYPE_APPLE        MachineType = 4
	MachineType_MACHINE_TYPE_PHONE        MachineType = 5
	MachineType_MACHINE_TYPE_TABLET       MachineType = 6
	MachineType_MACHINE_TYPE_AMD          MachineType = 7
)

// Enum value maps for MachineType.
var (
	MachineType_name = map[int32]string{
		0: "MACHINE_TYPE_UNKNOWN",
		1: "MACHINE_TYPE_RASPBERRY_PI",
		2: "MACHINE_TYPE_IOT_DEVICE",
		3: "MACHINE_TYPE_INTEL",
		4: "MACHINE_TYPE_APPLE",
		5: "MACHINE_TYPE_PHONE",
		6: "MACHINE_TYPE_TABLET",
		7: "MACHINE_TYPE_AMD",
	}
	MachineType_value = map[string]int32{
		"MACHINE_TYPE_UNKNOWN":      0,
		"MACHINE_TYPE_RASPBERRY_PI": 1,
		"MACHINE_TYPE_IOT_DEVICE":   2,
		"MACHINE_TYPE_INTEL":        3,
		"MACHINE_TYPE_APPLE":        4,
		"MACHINE_TYPE_PHONE":        5,
		"MACHINE_TYPE_TABLET":       6,
		"MACHINE_TYPE_AMD":          7,
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

type MachineUse int32

const (
	MachineUse_MACHINE_USE_UNKNOWN            MachineUse = 0
	MachineUse_MACHINE_USE_DEV_DESKTOP        MachineUse = 1
	MachineUse_MACHINE_USE_DEV_SERVER         MachineUse = 2
	MachineUse_MACHINE_USE_KUBERNETES_CLUSTER MachineUse = 3
	MachineUse_MACHINE_USE_LOCAL_CLUSTER      MachineUse = 4
	MachineUse_MACHINE_USE_NOT_IN_USE         MachineUse = 5
	MachineUse_MACHINE_USE_PI_SERVER          MachineUse = 6
)

// Enum value maps for MachineUse.
var (
	MachineUse_name = map[int32]string{
		0: "MACHINE_USE_UNKNOWN",
		1: "MACHINE_USE_DEV_DESKTOP",
		2: "MACHINE_USE_DEV_SERVER",
		3: "MACHINE_USE_KUBERNETES_CLUSTER",
		4: "MACHINE_USE_LOCAL_CLUSTER",
		5: "MACHINE_USE_NOT_IN_USE",
		6: "MACHINE_USE_PI_SERVER",
	}
	MachineUse_value = map[string]int32{
		"MACHINE_USE_UNKNOWN":            0,
		"MACHINE_USE_DEV_DESKTOP":        1,
		"MACHINE_USE_DEV_SERVER":         2,
		"MACHINE_USE_KUBERNETES_CLUSTER": 3,
		"MACHINE_USE_LOCAL_CLUSTER":      4,
		"MACHINE_USE_NOT_IN_USE":         5,
		"MACHINE_USE_PI_SERVER":          6,
	}
)

func (x MachineUse) Enum() *MachineUse {
	p := new(MachineUse)
	*p = x
	return p
}

func (x MachineUse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MachineUse) Descriptor() protoreflect.EnumDescriptor {
	return file_mdb_proto_enumTypes[1].Descriptor()
}

func (MachineUse) Type() protoreflect.EnumType {
	return &file_mdb_proto_enumTypes[1]
}

func (x MachineUse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MachineUse.Descriptor instead.
func (MachineUse) EnumDescriptor() ([]byte, []int) {
	return file_mdb_proto_rawDescGZIP(), []int{1}
}

type MachineErrors int32

const (
	MachineErrors_MACHINE_ERROR_NONE         MachineErrors = 0
	MachineErrors_MACHINE_ERROR_MISSING_TYPE MachineErrors = 1
	MachineErrors_MACHINE_ERROR_UNSTABLE_IP  MachineErrors = 2
	MachineErrors_MACHINE_ERROR_MISSING_USE  MachineErrors = 3
)

// Enum value maps for MachineErrors.
var (
	MachineErrors_name = map[int32]string{
		0: "MACHINE_ERROR_NONE",
		1: "MACHINE_ERROR_MISSING_TYPE",
		2: "MACHINE_ERROR_UNSTABLE_IP",
		3: "MACHINE_ERROR_MISSING_USE",
	}
	MachineErrors_value = map[string]int32{
		"MACHINE_ERROR_NONE":         0,
		"MACHINE_ERROR_MISSING_TYPE": 1,
		"MACHINE_ERROR_UNSTABLE_IP":  2,
		"MACHINE_ERROR_MISSING_USE":  3,
	}
)

func (x MachineErrors) Enum() *MachineErrors {
	p := new(MachineErrors)
	*p = x
	return p
}

func (x MachineErrors) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MachineErrors) Descriptor() protoreflect.EnumDescriptor {
	return file_mdb_proto_enumTypes[2].Descriptor()
}

func (MachineErrors) Type() protoreflect.EnumType {
	return &file_mdb_proto_enumTypes[2]
}

func (x MachineErrors) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MachineErrors.Descriptor instead.
func (MachineErrors) EnumDescriptor() ([]byte, []int) {
	return file_mdb_proto_rawDescGZIP(), []int{2}
}

type MachineConnectivity int32

const (
	MachineConnectivity_MACHINE_CONNECTIVITY_UNKNOWN   MachineConnectivity = 0
	MachineConnectivity_MACHINE_CONNECTIVITY_WIRED     MachineConnectivity = 1
	MachineConnectivity_MACHINE_CONNECTIVITY_EPHEMERAL MachineConnectivity = 2
)

// Enum value maps for MachineConnectivity.
var (
	MachineConnectivity_name = map[int32]string{
		0: "MACHINE_CONNECTIVITY_UNKNOWN",
		1: "MACHINE_CONNECTIVITY_WIRED",
		2: "MACHINE_CONNECTIVITY_EPHEMERAL",
	}
	MachineConnectivity_value = map[string]int32{
		"MACHINE_CONNECTIVITY_UNKNOWN":   0,
		"MACHINE_CONNECTIVITY_WIRED":     1,
		"MACHINE_CONNECTIVITY_EPHEMERAL": 2,
	}
)

func (x MachineConnectivity) Enum() *MachineConnectivity {
	p := new(MachineConnectivity)
	*p = x
	return p
}

func (x MachineConnectivity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MachineConnectivity) Descriptor() protoreflect.EnumDescriptor {
	return file_mdb_proto_enumTypes[3].Descriptor()
}

func (MachineConnectivity) Type() protoreflect.EnumType {
	return &file_mdb_proto_enumTypes[3]
}

func (x MachineConnectivity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MachineConnectivity.Descriptor instead.
func (MachineConnectivity) EnumDescriptor() ([]byte, []int) {
	return file_mdb_proto_rawDescGZIP(), []int{3}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentMachine *Machine      `protobuf:"bytes,1,opt,name=current_machine,json=currentMachine,proto3" json:"current_machine,omitempty"`
	IssueId        int32         `protobuf:"varint,2,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	IssueType      MachineErrors `protobuf:"varint,3,opt,name=issue_type,json=issueType,proto3,enum=mdb.MachineErrors" json:"issue_type,omitempty"`
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

func (x *Config) GetIssueType() MachineErrors {
	if x != nil {
		return x.IssueType
	}
	return MachineErrors_MACHINE_ERROR_NONE
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
	Type            MachineType         `protobuf:"varint,5,opt,name=type,proto3,enum=mdb.MachineType" json:"type,omitempty"`
	Use             MachineUse          `protobuf:"varint,6,opt,name=use,proto3,enum=mdb.MachineUse" json:"use,omitempty"`
	Connectivity    MachineConnectivity `protobuf:"varint,8,opt,name=connectivity,proto3,enum=mdb.MachineConnectivity" json:"connectivity,omitempty"`
	MarkedForDelete bool                `protobuf:"varint,7,opt,name=marked_for_delete,json=markedForDelete,proto3" json:"marked_for_delete,omitempty"`
	Version         string              `protobuf:"bytes,9,opt,name=version,proto3" json:"version,omitempty"`
	LastUpdated     int64               `protobuf:"varint,10,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
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

func (x *Machine) GetUse() MachineUse {
	if x != nil {
		return x.Use
	}
	return MachineUse_MACHINE_USE_UNKNOWN
}

func (x *Machine) GetConnectivity() MachineConnectivity {
	if x != nil {
		return x.Connectivity
	}
	return MachineConnectivity_MACHINE_CONNECTIVITY_UNKNOWN
}

func (x *Machine) GetMarkedForDelete() bool {
	if x != nil {
		return x.MarkedForDelete
	}
	return false
}

func (x *Machine) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Machine) GetLastUpdated() int64 {
	if x != nil {
		return x.LastUpdated
	}
	return 0
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

type UpdateMachineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string      `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	NewType  MachineType `protobuf:"varint,2,opt,name=new_type,json=newType,proto3,enum=mdb.MachineType" json:"new_type,omitempty"`
	Ipv4     uint32      `protobuf:"fixed32,3,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Remove   bool        `protobuf:"varint,4,opt,name=remove,proto3" json:"remove,omitempty"`
}

func (x *UpdateMachineRequest) Reset() {
	*x = UpdateMachineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mdb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMachineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMachineRequest) ProtoMessage() {}

func (x *UpdateMachineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mdb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMachineRequest.ProtoReflect.Descriptor instead.
func (*UpdateMachineRequest) Descriptor() ([]byte, []int) {
	return file_mdb_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateMachineRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *UpdateMachineRequest) GetNewType() MachineType {
	if x != nil {
		return x.NewType
	}
	return MachineType_MACHINE_TYPE_UNKNOWN
}

func (x *UpdateMachineRequest) GetIpv4() uint32 {
	if x != nil {
		return x.Ipv4
	}
	return 0
}

func (x *UpdateMachineRequest) GetRemove() bool {
	if x != nil {
		return x.Remove
	}
	return false
}

type UpdateMachineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateMachineResponse) Reset() {
	*x = UpdateMachineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mdb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMachineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMachineResponse) ProtoMessage() {}

func (x *UpdateMachineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mdb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMachineResponse.ProtoReflect.Descriptor instead.
func (*UpdateMachineResponse) Descriptor() ([]byte, []int) {
	return file_mdb_proto_rawDescGZIP(), []int{6}
}

var File_mdb_proto protoreflect.FileDescriptor

var file_mdb_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x64, 0x62,
	0x22, 0x8d, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x35, 0x0a, 0x0f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x12, 0x31, 0x0a,
	0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x54, 0x0a, 0x03, 0x4d, 0x64, 0x62, 0x12, 0x28, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x64, 0x62, 0x2e,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xdb, 0x02, 0x0a, 0x07, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07,
	0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x61, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x75, 0x73,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x52, 0x03, 0x75, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x46, 0x6f,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x22, 0x29, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x70, 0x76, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x22,
	0x40, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x64, 0x62, 0x2e,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x73, 0x22, 0x8b, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x07, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x22,
	0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0xda, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x41, 0x43, 0x48,
	0x49, 0x4e, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x52, 0x41, 0x53, 0x50, 0x42, 0x45, 0x52, 0x52, 0x59, 0x5f, 0x50, 0x49, 0x10,
	0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x4f, 0x54, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0x02, 0x12, 0x16,
	0x0a, 0x12, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49,
	0x4e, 0x54, 0x45, 0x4c, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x16,
	0x0a, 0x12, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x48, 0x4f, 0x4e, 0x45, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x54, 0x10, 0x06, 0x12,
	0x14, 0x0a, 0x10, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x41, 0x4d, 0x44, 0x10, 0x07, 0x2a, 0xd8, 0x01, 0x0a, 0x0a, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x55, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f,
	0x55, 0x53, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1b, 0x0a,
	0x17, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x44, 0x45, 0x56,
	0x5f, 0x44, 0x45, 0x53, 0x4b, 0x54, 0x4f, 0x50, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x41,
	0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x52, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e,
	0x45, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53,
	0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x41,
	0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f,
	0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x41, 0x43,
	0x48, 0x49, 0x4e, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f,
	0x55, 0x53, 0x45, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45,
	0x5f, 0x55, 0x53, 0x45, 0x5f, 0x50, 0x49, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x06,
	0x2a, 0x85, 0x01, 0x0a, 0x0d, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x41,
	0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x41,
	0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x54,
	0x41, 0x42, 0x4c, 0x45, 0x5f, 0x49, 0x50, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x41, 0x43,
	0x48, 0x49, 0x4e, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4e, 0x47, 0x5f, 0x55, 0x53, 0x45, 0x10, 0x03, 0x2a, 0x7b, 0x0a, 0x13, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12,
	0x20, 0x0a, 0x1c, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x43, 0x4f, 0x4e,
	0x4e, 0x45, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x57, 0x49, 0x52, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x22, 0x0a, 0x1e, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x43, 0x4f, 0x4e,
	0x4e, 0x45, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x45, 0x50, 0x48, 0x45, 0x4d, 0x45,
	0x52, 0x41, 0x4c, 0x10, 0x02, 0x32, 0x9d, 0x01, 0x0a, 0x0a, 0x4d, 0x44, 0x42, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x19, 0x2e, 0x6d,
	0x64, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x63,
	0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_mdb_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_mdb_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_mdb_proto_goTypes = []interface{}{
	(MachineType)(0),              // 0: mdb.MachineType
	(MachineUse)(0),               // 1: mdb.MachineUse
	(MachineErrors)(0),            // 2: mdb.MachineErrors
	(MachineConnectivity)(0),      // 3: mdb.MachineConnectivity
	(*Config)(nil),                // 4: mdb.Config
	(*Mdb)(nil),                   // 5: mdb.Mdb
	(*Machine)(nil),               // 6: mdb.Machine
	(*ListMachinesRequest)(nil),   // 7: mdb.ListMachinesRequest
	(*ListMachinesResponse)(nil),  // 8: mdb.ListMachinesResponse
	(*UpdateMachineRequest)(nil),  // 9: mdb.UpdateMachineRequest
	(*UpdateMachineResponse)(nil), // 10: mdb.UpdateMachineResponse
}
var file_mdb_proto_depIdxs = []int32{
	6,  // 0: mdb.Config.current_machine:type_name -> mdb.Machine
	2,  // 1: mdb.Config.issue_type:type_name -> mdb.MachineErrors
	6,  // 2: mdb.Mdb.machines:type_name -> mdb.Machine
	4,  // 3: mdb.Mdb.config:type_name -> mdb.Config
	0,  // 4: mdb.Machine.type:type_name -> mdb.MachineType
	1,  // 5: mdb.Machine.use:type_name -> mdb.MachineUse
	3,  // 6: mdb.Machine.connectivity:type_name -> mdb.MachineConnectivity
	6,  // 7: mdb.ListMachinesResponse.machines:type_name -> mdb.Machine
	0,  // 8: mdb.UpdateMachineRequest.new_type:type_name -> mdb.MachineType
	7,  // 9: mdb.MDBService.ListMachines:input_type -> mdb.ListMachinesRequest
	9,  // 10: mdb.MDBService.UpdateMachine:input_type -> mdb.UpdateMachineRequest
	8,  // 11: mdb.MDBService.ListMachines:output_type -> mdb.ListMachinesResponse
	10, // 12: mdb.MDBService.UpdateMachine:output_type -> mdb.UpdateMachineResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
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
		file_mdb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMachineRequest); i {
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
		file_mdb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMachineResponse); i {
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
			NumEnums:      4,
			NumMessages:   7,
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
