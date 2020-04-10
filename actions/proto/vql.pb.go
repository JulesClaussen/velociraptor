// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vql.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	proto1 "www.velocidex.com/golang/velociraptor/artifacts/proto"
	_ "www.velocidex.com/golang/velociraptor/proto"
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

type VQLRequest struct {
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	VQL                  string   `protobuf:"bytes,1,opt,name=VQL,proto3" json:"VQL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VQLRequest) Reset()         { *m = VQLRequest{} }
func (m *VQLRequest) String() string { return proto.CompactTextString(m) }
func (*VQLRequest) ProtoMessage()    {}
func (*VQLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e14263f319df14a, []int{0}
}

func (m *VQLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLRequest.Unmarshal(m, b)
}
func (m *VQLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLRequest.Marshal(b, m, deterministic)
}
func (m *VQLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLRequest.Merge(m, src)
}
func (m *VQLRequest) XXX_Size() int {
	return xxx_messageInfo_VQLRequest.Size(m)
}
func (m *VQLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VQLRequest proto.InternalMessageInfo

func (m *VQLRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VQLRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *VQLRequest) GetVQL() string {
	if m != nil {
		return m.VQL
	}
	return ""
}

type VQLEnv struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VQLEnv) Reset()         { *m = VQLEnv{} }
func (m *VQLEnv) String() string { return proto.CompactTextString(m) }
func (*VQLEnv) ProtoMessage()    {}
func (*VQLEnv) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e14263f319df14a, []int{1}
}

func (m *VQLEnv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLEnv.Unmarshal(m, b)
}
func (m *VQLEnv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLEnv.Marshal(b, m, deterministic)
}
func (m *VQLEnv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLEnv.Merge(m, src)
}
func (m *VQLEnv) XXX_Size() int {
	return xxx_messageInfo_VQLEnv.Size(m)
}
func (m *VQLEnv) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLEnv.DiscardUnknown(m)
}

var xxx_messageInfo_VQLEnv proto.InternalMessageInfo

func (m *VQLEnv) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *VQLEnv) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type VQLCollectorArgs struct {
	Env                  []*VQLEnv          `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	Query                []*VQLRequest      `protobuf:"bytes,2,rep,name=Query,proto3" json:"Query,omitempty"`
	MaxRow               uint64             `protobuf:"varint,4,opt,name=max_row,json=maxRow,proto3" json:"max_row,omitempty"`
	MaxWait              uint64             `protobuf:"varint,6,opt,name=max_wait,json=maxWait,proto3" json:"max_wait,omitempty"`
	OpsPerSecond         float32            `protobuf:"fixed32,24,opt,name=ops_per_second,json=opsPerSecond,proto3" json:"ops_per_second,omitempty"`
	Artifacts            []*proto1.Artifact `protobuf:"bytes,5,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	Timeout              uint64             `protobuf:"varint,25,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *VQLCollectorArgs) Reset()         { *m = VQLCollectorArgs{} }
func (m *VQLCollectorArgs) String() string { return proto.CompactTextString(m) }
func (*VQLCollectorArgs) ProtoMessage()    {}
func (*VQLCollectorArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e14263f319df14a, []int{2}
}

func (m *VQLCollectorArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLCollectorArgs.Unmarshal(m, b)
}
func (m *VQLCollectorArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLCollectorArgs.Marshal(b, m, deterministic)
}
func (m *VQLCollectorArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLCollectorArgs.Merge(m, src)
}
func (m *VQLCollectorArgs) XXX_Size() int {
	return xxx_messageInfo_VQLCollectorArgs.Size(m)
}
func (m *VQLCollectorArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLCollectorArgs.DiscardUnknown(m)
}

var xxx_messageInfo_VQLCollectorArgs proto.InternalMessageInfo

func (m *VQLCollectorArgs) GetEnv() []*VQLEnv {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *VQLCollectorArgs) GetQuery() []*VQLRequest {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *VQLCollectorArgs) GetMaxRow() uint64 {
	if m != nil {
		return m.MaxRow
	}
	return 0
}

func (m *VQLCollectorArgs) GetMaxWait() uint64 {
	if m != nil {
		return m.MaxWait
	}
	return 0
}

func (m *VQLCollectorArgs) GetOpsPerSecond() float32 {
	if m != nil {
		return m.OpsPerSecond
	}
	return 0
}

func (m *VQLCollectorArgs) GetArtifacts() []*proto1.Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func (m *VQLCollectorArgs) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type VQLTypeMap struct {
	Column               string   `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VQLTypeMap) Reset()         { *m = VQLTypeMap{} }
func (m *VQLTypeMap) String() string { return proto.CompactTextString(m) }
func (*VQLTypeMap) ProtoMessage()    {}
func (*VQLTypeMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e14263f319df14a, []int{3}
}

func (m *VQLTypeMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLTypeMap.Unmarshal(m, b)
}
func (m *VQLTypeMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLTypeMap.Marshal(b, m, deterministic)
}
func (m *VQLTypeMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLTypeMap.Merge(m, src)
}
func (m *VQLTypeMap) XXX_Size() int {
	return xxx_messageInfo_VQLTypeMap.Size(m)
}
func (m *VQLTypeMap) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLTypeMap.DiscardUnknown(m)
}

var xxx_messageInfo_VQLTypeMap proto.InternalMessageInfo

func (m *VQLTypeMap) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *VQLTypeMap) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type VQLResponse struct {
	Response             string        `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
	Columns              []string      `protobuf:"bytes,2,rep,name=Columns,proto3" json:"Columns,omitempty"`
	Types                []*VQLTypeMap `protobuf:"bytes,8,rep,name=types,proto3" json:"types,omitempty"`
	QueryId              uint64        `protobuf:"varint,5,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	Part                 uint64        `protobuf:"varint,6,opt,name=part,proto3" json:"part,omitempty"`
	Query                *VQLRequest   `protobuf:"bytes,3,opt,name=Query,proto3" json:"Query,omitempty"`
	Timestamp            uint64        `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TotalRows            uint64        `protobuf:"varint,7,opt,name=total_rows,json=totalRows,proto3" json:"total_rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VQLResponse) Reset()         { *m = VQLResponse{} }
func (m *VQLResponse) String() string { return proto.CompactTextString(m) }
func (*VQLResponse) ProtoMessage()    {}
func (*VQLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e14263f319df14a, []int{4}
}

func (m *VQLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLResponse.Unmarshal(m, b)
}
func (m *VQLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLResponse.Marshal(b, m, deterministic)
}
func (m *VQLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLResponse.Merge(m, src)
}
func (m *VQLResponse) XXX_Size() int {
	return xxx_messageInfo_VQLResponse.Size(m)
}
func (m *VQLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VQLResponse proto.InternalMessageInfo

func (m *VQLResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func (m *VQLResponse) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *VQLResponse) GetTypes() []*VQLTypeMap {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *VQLResponse) GetQueryId() uint64 {
	if m != nil {
		return m.QueryId
	}
	return 0
}

func (m *VQLResponse) GetPart() uint64 {
	if m != nil {
		return m.Part
	}
	return 0
}

func (m *VQLResponse) GetQuery() *VQLRequest {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *VQLResponse) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *VQLResponse) GetTotalRows() uint64 {
	if m != nil {
		return m.TotalRows
	}
	return 0
}

// FIXME: We replicate a small subset of GRR's elaborate knowledgebase
// protos here because the GUI API plugins use this to construct the
// GRR APIs. When we re-implement the API plugins, refactor this into
// a more sane structure.
type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e14263f319df14a, []int{5}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type VQLEventTable struct {
	Event                []*VQLCollectorArgs `protobuf:"bytes,1,rep,name=event,proto3" json:"event,omitempty"`
	Version              uint64              `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VQLEventTable) Reset()         { *m = VQLEventTable{} }
func (m *VQLEventTable) String() string { return proto.CompactTextString(m) }
func (*VQLEventTable) ProtoMessage()    {}
func (*VQLEventTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e14263f319df14a, []int{6}
}

func (m *VQLEventTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLEventTable.Unmarshal(m, b)
}
func (m *VQLEventTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLEventTable.Marshal(b, m, deterministic)
}
func (m *VQLEventTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLEventTable.Merge(m, src)
}
func (m *VQLEventTable) XXX_Size() int {
	return xxx_messageInfo_VQLEventTable.Size(m)
}
func (m *VQLEventTable) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLEventTable.DiscardUnknown(m)
}

var xxx_messageInfo_VQLEventTable proto.InternalMessageInfo

func (m *VQLEventTable) GetEvent() []*VQLCollectorArgs {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *VQLEventTable) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type ClientInfo struct {
	Hostname              string   `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Fqdn                  string   `protobuf:"bytes,4,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	System                string   `protobuf:"bytes,5,opt,name=system,proto3" json:"system,omitempty"`
	Release               string   `protobuf:"bytes,6,opt,name=release,proto3" json:"release,omitempty"`
	Architecture          string   `protobuf:"bytes,7,opt,name=architecture,proto3" json:"architecture,omitempty"`
	IpAddress             string   `protobuf:"bytes,10,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Ping                  uint64   `protobuf:"varint,11,opt,name=ping,proto3" json:"ping,omitempty"`
	ClientVersion         string   `protobuf:"bytes,12,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	ClientName            string   `protobuf:"bytes,13,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	Labels                []string `protobuf:"bytes,15,rep,name=labels,proto3" json:"labels,omitempty"`
	LastInterrogateFlowId string   `protobuf:"bytes,16,opt,name=last_interrogate_flow_id,json=lastInterrogateFlowId,proto3" json:"last_interrogate_flow_id,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ClientInfo) Reset()         { *m = ClientInfo{} }
func (m *ClientInfo) String() string { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()    {}
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e14263f319df14a, []int{7}
}

func (m *ClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInfo.Unmarshal(m, b)
}
func (m *ClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInfo.Marshal(b, m, deterministic)
}
func (m *ClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInfo.Merge(m, src)
}
func (m *ClientInfo) XXX_Size() int {
	return xxx_messageInfo_ClientInfo.Size(m)
}
func (m *ClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInfo proto.InternalMessageInfo

func (m *ClientInfo) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *ClientInfo) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *ClientInfo) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *ClientInfo) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *ClientInfo) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

func (m *ClientInfo) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *ClientInfo) GetPing() uint64 {
	if m != nil {
		return m.Ping
	}
	return 0
}

func (m *ClientInfo) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *ClientInfo) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *ClientInfo) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ClientInfo) GetLastInterrogateFlowId() string {
	if m != nil {
		return m.LastInterrogateFlowId
	}
	return ""
}

func init() {
	proto.RegisterType((*VQLRequest)(nil), "proto.VQLRequest")
	proto.RegisterType((*VQLEnv)(nil), "proto.VQLEnv")
	proto.RegisterType((*VQLCollectorArgs)(nil), "proto.VQLCollectorArgs")
	proto.RegisterType((*VQLTypeMap)(nil), "proto.VQLTypeMap")
	proto.RegisterType((*VQLResponse)(nil), "proto.VQLResponse")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*VQLEventTable)(nil), "proto.VQLEventTable")
	proto.RegisterType((*ClientInfo)(nil), "proto.ClientInfo")
}

func init() {
	proto.RegisterFile("vql.proto", fileDescriptor_3e14263f319df14a)
}

var fileDescriptor_3e14263f319df14a = []byte{
	// 1574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x56, 0x4f, 0x6f, 0x24, 0x47,
	0x15, 0xd7, 0xd8, 0x1e, 0x8f, 0xa7, 0xbc, 0x4e, 0x36, 0xa5, 0xb0, 0xf4, 0x2e, 0xb0, 0x14, 0x43,
	0x00, 0x87, 0x28, 0x3d, 0xde, 0x5d, 0xad, 0x76, 0x09, 0x02, 0xe4, 0x59, 0xef, 0x22, 0x23, 0x3b,
	0xc1, 0x1d, 0xc7, 0x41, 0xab, 0x88, 0xa1, 0xa6, 0xfb, 0xcd, 0x4c, 0x69, 0xab, 0xab, 0xda, 0x55,
	0xd5, 0xd3, 0x9e, 0x43, 0x4e, 0x1c, 0xe0, 0x1b, 0x20, 0x21, 0x71, 0xe0, 0x16, 0x71, 0xe1, 0x4b,
	0x70, 0x88, 0xe0, 0x3b, 0x70, 0x80, 0x1b, 0x9f, 0x81, 0x03, 0x7a, 0xaf, 0xbb, 0x67, 0xec, 0x15,
	0x39, 0x75, 0xfd, 0xfd, 0xbd, 0xdf, 0xfb, 0xf7, 0xab, 0x66, 0xfd, 0xc5, 0xa5, 0x8e, 0x0b, 0x67,
	0x83, 0xe5, 0x5d, 0xfa, 0xdc, 0x7b, 0x9b, 0x3e, 0x43, 0x0f, 0xb9, 0x34, 0x41, 0xa5, 0xf5, 0xe6,
	0xbd, 0xbb, 0xf5, 0xea, 0x54, 0xdb, 0x6a, 0x9c, 0x43, 0x90, 0x99, 0x0c, 0xb2, 0xd9, 0xba, 0x2f,
	0x5d, 0x50, 0x53, 0x99, 0x06, 0x3f, 0xac, 0x0f, 0xb5, 0xf3, 0x7a, 0x7f, 0xf0, 0xc5, 0x06, 0x63,
	0x17, 0x67, 0x27, 0x09, 0x5c, 0x96, 0xe0, 0x03, 0xff, 0x7d, 0x87, 0x6d, 0x7d, 0x28, 0x73, 0x88,
	0x36, 0x44, 0x67, 0xbf, 0x3f, 0x0a, 0xff, 0xfa, 0xef, 0xbf, 0xff, 0xd6, 0x31, 0x5c, 0x9f, 0xcf,
	0x41, 0x18, 0x99, 0x83, 0xb0, 0x53, 0x11, 0xe6, 0xca, 0x8b, 0xcb, 0x12, 0xdc, 0x32, 0x16, 0xe7,
	0x38, 0xf6, 0x73, 0x5b, 0xea, 0x4c, 0x4c, 0x40, 0x64, 0xe0, 0x53, 0xa7, 0x8a, 0xa0, 0x16, 0x20,
	0x82, 0x15, 0xca, 0x64, 0x2a, 0x95, 0x01, 0x44, 0x35, 0x97, 0x41, 0x84, 0x65, 0x41, 0xf7, 0x95,
	0x99, 0x5a, 0x97, 0x4b, 0x15, 0xac, 0x11, 0x61, 0x0e, 0x35, 0x94, 0x70, 0x10, 0x9c, 0x82, 0x05,
	0xf8, 0x38, 0x21, 0x06, 0xfc, 0x84, 0xed, 0x1e, 0xb5, 0x70, 0xd6, 0x44, 0x9b, 0x44, 0xe8, 0x87,
	0x44, 0xe8, 0x1d, 0x3e, 0xf8, 0x94, 0x10, 0x57, 0x4c, 0x04, 0xd2, 0x28, 0x8b, 0xc2, 0x7a, 0xc8,
	0xd0, 0x72, 0x66, 0xe3, 0xe4, 0xfa, 0x75, 0x7e, 0xc8, 0x36, 0x2f, 0xce, 0x4e, 0xa2, 0x0e, 0xa1,
	0x0c, 0x09, 0xe5, 0x5d, 0xfe, 0x03, 0x74, 0xeb, 0xe2, 0xec, 0xa4, 0xc1, 0x08, 0x56, 0xc0, 0x15,
	0xa4, 0x65, 0x00, 0xd1, 0x50, 0x4b, 0xb5, 0x02, 0x13, 0xe2, 0x04, 0xef, 0x0e, 0x0e, 0xd8, 0xf6,
	0xc5, 0xd9, 0xc9, 0x73, 0xb3, 0xe0, 0xb7, 0xd9, 0xe6, 0x2b, 0x58, 0xd6, 0x60, 0x09, 0x0e, 0xf9,
	0xdb, 0xac, 0xbb, 0x90, 0xba, 0x6c, 0xe2, 0x96, 0xd4, 0x93, 0xc1, 0xdf, 0x7b, 0xec, 0xf6, 0xc5,
	0xd9, 0xc9, 0x33, 0xab, 0x35, 0xa4, 0xc1, 0xba, 0x43, 0x37, 0xf3, 0xfc, 0x33, 0xb6, 0x09, 0x66,
	0x11, 0x6d, 0x8a, 0xcd, 0xfd, 0xdd, 0x87, 0x7b, 0x75, 0x1a, 0xe2, 0x1a, 0x78, 0xf4, 0x63, 0x22,
	0xf6, 0x98, 0x3f, 0x7a, 0x6e, 0x16, 0xca, 0x59, 0x93, 0x83, 0x09, 0x62, 0x21, 0x9d, 0x92, 0x13,
	0x0d, 0x1e, 0x09, 0x4e, 0x40, 0x14, 0xce, 0x2e, 0x54, 0x06, 0x99, 0x98, 0x5a, 0xb7, 0x0e, 0x60,
	0x9c, 0x20, 0x2c, 0x7f, 0xc9, 0xba, 0x67, 0x38, 0x8d, 0x36, 0x08, 0xff, 0xad, 0x35, 0x7e, 0x93,
	0xe2, 0xd1, 0x03, 0xb2, 0xf1, 0x1e, 0x7f, 0xf7, 0xba, 0xf3, 0xaa, 0x46, 0xff, 0x0a, 0xf7, 0x6b,
	0x48, 0xfe, 0xa7, 0x0e, 0xeb, 0xe5, 0xf2, 0x6a, 0xec, 0x6c, 0x15, 0x6d, 0x89, 0xce, 0xfe, 0xd6,
	0xe8, 0xb7, 0x1d, 0x02, 0xfb, 0x9c, 0x4f, 0x10, 0x2c, 0x97, 0x57, 0x2a, 0x2f, 0x73, 0xe1, 0x6c,
	0xe5, 0x45, 0x01, 0x4e, 0x38, 0xf0, 0x85, 0x35, 0x1e, 0x62, 0x91, 0x34, 0x23, 0x2f, 0xb4, 0x74,
	0x33, 0x40, 0xde, 0xd2, 0xd4, 0xf9, 0xab, 0x94, 0xd6, 0xe8, 0x96, 0x2f, 0xb4, 0x0a, 0x42, 0xa6,
	0xce, 0x7a, 0x2f, 0xf2, 0x52, 0x07, 0x55, 0x68, 0x58, 0x41, 0xf8, 0x78, 0xf0, 0xd6, 0xa9, 0xbc,
	0x5a, 0x63, 0x17, 0xd2, 0x85, 0x87, 0x5b, 0x0f, 0x0e, 0x0e, 0x0e, 0x92, 0xed, 0x5c, 0x5e, 0x25,
	0xb6, 0xe2, 0xff, 0xe8, 0xb0, 0x1d, 0xe4, 0x57, 0x49, 0x15, 0xa2, 0x6d, 0x22, 0xf8, 0xd7, 0x9a,
	0xe0, 0x17, 0x1d, 0xfe, 0xf9, 0x0b, 0xeb, 0x84, 0xb6, 0x66, 0xb6, 0xf2, 0xb7, 0x42, 0xf4, 0x50,
	0x3a, 0x43, 0x58, 0x4a, 0x6a, 0x34, 0x56, 0xea, 0xe0, 0x85, 0x9c, 0x06, 0xa2, 0xa8, 0x3c, 0xdd,
	0x68, 0xca, 0x5c, 0x79, 0xe1, 0xe0, 0xb2, 0x54, 0xae, 0x09, 0x3d, 0x2c, 0x30, 0x3b, 0x5a, 0xf9,
	0x00, 0x06, 0x9c, 0x17, 0xd5, 0x5c, 0xa5, 0x73, 0x61, 0x60, 0x41, 0x9e, 0x4b, 0xad, 0x97, 0x22,
	0xb5, 0x79, 0xa1, 0x21, 0x40, 0x3c, 0xf8, 0xee, 0x48, 0x86, 0x74, 0x4e, 0x00, 0xe0, 0x83, 0x47,
	0xff, 0x83, 0x08, 0xf2, 0x15, 0x5c, 0x33, 0xf3, 0x70, 0xe3, 0xc1, 0x41, 0x82, 0x11, 0xfe, 0x54,
	0xaa, 0xc0, 0xbf, 0xec, 0xb0, 0x37, 0x6c, 0xe1, 0xc7, 0x05, 0xb8, 0xb1, 0x87, 0xd4, 0x9a, 0x2c,
	0x8a, 0x44, 0x67, 0x7f, 0x63, 0xf4, 0xe7, 0xda, 0xa7, 0x3f, 0x76, 0xf8, 0x1f, 0x3a, 0x87, 0x46,
	0x7c, 0x54, 0x20, 0xb3, 0x0c, 0xa6, 0xca, 0x40, 0x26, 0xa4, 0x17, 0xde, 0xe6, 0x20, 0xa4, 0x9b,
	0xa8, 0xe0, 0xa4, 0x5b, 0x8a, 0xd2, 0xa8, 0x80, 0x6d, 0x57, 0x59, 0xf7, 0xaa, 0xf1, 0x44, 0x6a,
	0x8d, 0x61, 0xc4, 0x95, 0xa6, 0xa6, 0xb4, 0xca, 0x55, 0x80, 0x2c, 0x16, 0xe7, 0xcb, 0x42, 0xa5,
	0xc4, 0x1d, 0x4b, 0xa3, 0xd0, 0xe5, 0x4c, 0x99, 0x26, 0x47, 0xa9, 0x2d, 0x4d, 0x10, 0xb6, 0xc0,
	0x4a, 0xa9, 0xa4, 0xcb, 0xfc, 0xb5, 0xee, 0x95, 0x5e, 0xc8, 0xa2, 0x70, 0xb6, 0x70, 0x4a, 0x06,
	0x88, 0x93, 0x5b, 0xb6, 0xf0, 0xbf, 0x04, 0xf7, 0x31, 0xf1, 0xe6, 0x86, 0xf5, 0x57, 0x32, 0x14,
	0x75, 0xa9, 0x30, 0xdf, 0x6c, 0x0a, 0xf3, 0xb0, 0x59, 0x1f, 0xfd, 0x8c, 0x9c, 0xfa, 0x11, 0x7f,
	0xd2, 0xae, 0x78, 0xe1, 0x31, 0xbe, 0x53, 0x67, 0x73, 0x32, 0xe4, 0xc1, 0x61, 0x64, 0x83, 0x15,
	0x73, 0xd0, 0x85, 0xa8, 0x54, 0x98, 0x5f, 0x97, 0xa2, 0x64, 0x6d, 0x82, 0x1f, 0xb1, 0x5e, 0x50,
	0x39, 0xd8, 0x32, 0x44, 0x77, 0xa9, 0x0c, 0x56, 0xb2, 0x71, 0xda, 0x94, 0x28, 0x6e, 0xdf, 0xec,
	0x1f, 0x04, 0x76, 0xa5, 0x89, 0x93, 0xf6, 0xea, 0x07, 0x7b, 0xff, 0xf9, 0x9d, 0xe8, 0xb3, 0xde,
	0xcf, 0x31, 0xbb, 0x2a, 0x1d, 0x3c, 0x25, 0xa1, 0x3c, 0x5f, 0x16, 0x70, 0x2a, 0x0b, 0x7e, 0x87,
	0x6d, 0xa7, 0x56, 0x97, 0xb9, 0x69, 0x54, 0xa0, 0x99, 0x71, 0xce, 0xb6, 0x50, 0xe7, 0x1a, 0x1d,
	0xa0, 0xf1, 0xe0, 0xcb, 0x2e, 0xdb, 0xa5, 0x06, 0xac, 0x4b, 0x99, 0x7f, 0xc0, 0x76, 0xda, 0x71,
	0x23, 0x48, 0xf7, 0x89, 0x5f, 0xc4, 0xef, 0xfc, 0xe2, 0xe3, 0x8f, 0x3e, 0x14, 0x60, 0x52, 0x8b,
	0xfd, 0xbd, 0x6a, 0x9f, 0x64, 0x75, 0x9e, 0x27, 0xac, 0xf7, 0x8c, 0x2c, 0x79, 0xea, 0xf0, 0xfe,
	0xe8, 0x29, 0x5d, 0x7d, 0xc8, 0x0f, 0x0e, 0xa9, 0x14, 0x31, 0xd3, 0x35, 0x15, 0x31, 0x07, 0x99,
	0x29, 0x33, 0xf3, 0x28, 0x18, 0x59, 0x99, 0x42, 0x26, 0x26, 0xcb, 0xeb, 0x7a, 0xd1, 0x02, 0xf1,
	0x5f, 0xb3, 0x2e, 0xf2, 0xf4, 0xd1, 0xce, 0xeb, 0x9a, 0xd1, 0x78, 0x3b, 0x7a, 0x4c, 0x46, 0x86,
	0xfc, 0xfd, 0x53, 0x59, 0x14, 0xca, 0xcc, 0xc4, 0x04, 0x42, 0x05, 0x60, 0x5a, 0x53, 0xf8, 0x34,
	0x78, 0x21, 0x4d, 0x86, 0xf8, 0xca, 0x91, 0xd4, 0xfb, 0x38, 0xa9, 0x61, 0x79, 0xc2, 0x76, 0xc8,
	0xe4, 0x58, 0x65, 0x51, 0x97, 0xf2, 0xf1, 0x84, 0xf0, 0x1e, 0xf0, 0xe1, 0xb3, 0xb9, 0xb3, 0xc6,
	0x6a, 0x3b, 0xc3, 0xa2, 0x13, 0xd6, 0x65, 0xe0, 0xea, 0x27, 0xa6, 0xcd, 0x4a, 0xd5, 0x4a, 0x00,
	0xca, 0x7a, 0x9c, 0xf4, 0x68, 0xf5, 0x38, 0xe3, 0x81, 0x6d, 0x61, 0xd7, 0x36, 0x6d, 0xfe, 0x1b,
	0xc2, 0x7b, 0xc9, 0x7f, 0x75, 0x82, 0xea, 0x42, 0xa5, 0xbb, 0x12, 0x0e, 0x21, 0xdd, 0xeb, 0xe2,
	0x22, 0xcd, 0x92, 0x7a, 0xde, 0x37, 0xfd, 0x90, 0x4a, 0x57, 0x2b, 0xe0, 0x1c, 0x68, 0x7d, 0xf5,
	0xca, 0xad, 0x33, 0x40, 0xd6, 0xf8, 0x69, 0xab, 0xae, 0xf8, 0x1a, 0xfd, 0x5f, 0x75, 0x7d, 0x87,
	0x98, 0xdc, 0xe7, 0xdf, 0x3c, 0x5f, 0x17, 0x16, 0xb6, 0x7a, 0x25, 0x7d, 0xab, 0xae, 0xd9, 0x4a,
	0x50, 0x3f, 0x63, 0x7d, 0x2c, 0x36, 0x1f, 0x64, 0x5e, 0x34, 0x8a, 0xfa, 0x53, 0xba, 0xff, 0x94,
	0xed, 0x26, 0x47, 0x2f, 0x8e, 0x64, 0x00, 0xdc, 0xaf, 0xa5, 0x7a, 0x75, 0x92, 0x68, 0xb6, 0xcc,
	0x08, 0x78, 0x86, 0x25, 0x2a, 0x09, 0x79, 0x0d, 0xc8, 0x13, 0xc6, 0x82, 0x0d, 0x52, 0xa3, 0x5e,
	0xfb, 0xa8, 0x47, 0xf0, 0x8f, 0x08, 0xfe, 0x7d, 0xfe, 0xde, 0x39, 0xee, 0x08, 0x53, 0xe6, 0x93,
	0x3a, 0xe2, 0xa4, 0xab, 0xca, 0xdc, 0x74, 0x9b, 0x82, 0x81, 0x98, 0x78, 0x38, 0xb1, 0x95, 0x1f,
	0x3c, 0x67, 0x5b, 0x9f, 0x78, 0x70, 0xfc, 0x27, 0x6c, 0xa7, 0xf4, 0xe0, 0x30, 0xef, 0x4d, 0x09,
	0x7f, 0x87, 0x90, 0xbf, 0xc1, 0xef, 0x22, 0xd7, 0x76, 0xaf, 0xcd, 0x25, 0xce, 0xe3, 0x64, 0x75,
	0x65, 0xf0, 0x97, 0x0e, 0xdb, 0xc3, 0x27, 0x0f, 0x15, 0xf4, 0x1c, 0xdf, 0x36, 0xfe, 0x09, 0xeb,
	0x92, 0x9e, 0x46, 0x1d, 0xaa, 0xc1, 0xaf, 0xaf, 0x23, 0x7b, 0xe3, 0xf5, 0x1c, 0x7d, 0x9f, 0xcc,
	0x08, 0x7e, 0xff, 0x50, 0x78, 0xa0, 0x44, 0xd5, 0x32, 0x7c, 0xed, 0x11, 0xa3, 0x2e, 0xae, 0xd1,
	0xf8, 0x88, 0xf5, 0x16, 0xe0, 0x3c, 0xfe, 0x40, 0x6c, 0x50, 0x00, 0xf6, 0xe9, 0xfe, 0x80, 0x0b,
	0xa4, 0xd9, 0x6c, 0xad, 0xd2, 0x5d, 0x43, 0x05, 0x24, 0x14, 0x27, 0xed, 0xc5, 0xc1, 0x3f, 0x37,
	0x18, 0x7b, 0x46, 0x2f, 0xe1, 0xb1, 0x99, 0x5a, 0x7e, 0x8f, 0xed, 0xcc, 0xad, 0x0f, 0xe4, 0x3a,
	0xfd, 0x94, 0x24, 0xab, 0x39, 0x76, 0xff, 0xf4, 0x32, 0x33, 0x94, 0xcb, 0x7e, 0x42, 0x63, 0x54,
	0x0a, 0xbf, 0xf4, 0x01, 0x72, 0xaa, 0xfd, 0x7e, 0xd2, 0xcc, 0x78, 0xc4, 0x7a, 0x0e, 0x34, 0x48,
	0x0f, 0x54, 0xc4, 0xfd, 0xa4, 0x9d, 0xf2, 0x01, 0xbb, 0x25, 0x5d, 0x3a, 0x57, 0x01, 0xd2, 0x50,
	0x3a, 0xa0, 0xd4, 0xf5, 0x93, 0x1b, 0x6b, 0xfc, 0x5b, 0x8c, 0xa9, 0x62, 0x2c, 0xb3, 0xcc, 0x81,
	0xf7, 0x11, 0xa3, 0x13, 0x7d, 0x55, 0x1c, 0xd6, 0x0b, 0x48, 0x04, 0xfb, 0x34, 0xda, 0x45, 0xa7,
	0x13, 0x1a, 0xf3, 0xef, 0xb1, 0x37, 0xea, 0x07, 0x7d, 0xdc, 0x86, 0xe4, 0x16, 0x5d, 0xdb, 0xab,
	0x57, 0x2f, 0xea, 0x45, 0xfe, 0x6d, 0xb6, 0xdb, 0x1c, 0x23, 0x17, 0xf7, 0xe8, 0x0c, 0xab, 0x97,
	0xe8, 0xc7, 0xec, 0x0e, 0xdb, 0xd6, 0x72, 0x02, 0xda, 0x47, 0x6f, 0xa2, 0x02, 0x25, 0xcd, 0x8c,
	0x3f, 0x61, 0x91, 0x96, 0x3e, 0x8c, 0x95, 0x09, 0xe0, 0x9c, 0x9d, 0xc9, 0x00, 0x63, 0xfa, 0x25,
	0x55, 0x59, 0x74, 0x9b, 0x50, 0xbe, 0x86, 0xfb, 0xc7, 0xeb, 0xed, 0x17, 0xda, 0x56, 0xc7, 0xd9,
	0xe8, 0xf1, 0xcb, 0x47, 0x55, 0x55, 0xc5, 0x0b, 0xd0, 0x36, 0x55, 0x19, 0x5c, 0xc5, 0xa9, 0xcd,
	0x87, 0x33, 0xab, 0xa5, 0x99, 0x0d, 0xeb, 0x45, 0x27, 0x8b, 0x60, 0xdd, 0x50, 0xa6, 0xf8, 0x2f,
	0xd7, 0xfc, 0xc9, 0x4e, 0xb6, 0xe9, 0xf3, 0xe8, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x80,
	0x94, 0xfc, 0x26, 0x0b, 0x00, 0x00,
}
