// Code generated by protoc-gen-go. DO NOT EDIT.
// source: download.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateDownloadRequest struct {
	FlowId               string   `protobuf:"bytes,1,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	HuntId               string   `protobuf:"bytes,3,opt,name=hunt_id,json=huntId,proto3" json:"hunt_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDownloadRequest) Reset()         { *m = CreateDownloadRequest{} }
func (m *CreateDownloadRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDownloadRequest) ProtoMessage()    {}
func (*CreateDownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ce52b48c9eea83, []int{0}
}

func (m *CreateDownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDownloadRequest.Unmarshal(m, b)
}
func (m *CreateDownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDownloadRequest.Marshal(b, m, deterministic)
}
func (m *CreateDownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDownloadRequest.Merge(m, src)
}
func (m *CreateDownloadRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDownloadRequest.Size(m)
}
func (m *CreateDownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDownloadRequest proto.InternalMessageInfo

func (m *CreateDownloadRequest) GetFlowId() string {
	if m != nil {
		return m.FlowId
	}
	return ""
}

func (m *CreateDownloadRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *CreateDownloadRequest) GetHuntId() string {
	if m != nil {
		return m.HuntId
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateDownloadRequest)(nil), "proto.CreateDownloadRequest")
}

func init() {
	proto.RegisterFile("download.proto", fileDescriptor_f7ce52b48c9eea83)
}

var fileDescriptor_f7ce52b48c9eea83 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xc9, 0x2f, 0xcf,
	0xcb, 0xc9, 0x4f, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x69,
	0x5c, 0xa2, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x2e, 0x50, 0xe9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4,
	0xe2, 0x12, 0x21, 0x71, 0x2e, 0xf6, 0xb4, 0x9c, 0xfc, 0xf2, 0xf8, 0xcc, 0x14, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x36, 0x10, 0xd7, 0x33, 0x45, 0x48, 0x9a, 0x8b, 0x33, 0x39, 0x27, 0x33,
	0x35, 0xaf, 0x04, 0x24, 0xc5, 0x04, 0x96, 0xe2, 0x80, 0x08, 0x78, 0xa6, 0x80, 0x74, 0x65, 0x94,
	0x42, 0xa4, 0x98, 0x21, 0xba, 0x40, 0x5c, 0xcf, 0x14, 0x27, 0xc3, 0x28, 0xfd, 0xf2, 0xf2, 0x72,
	0xbd, 0xb2, 0xd4, 0x9c, 0xfc, 0xe4, 0xcc, 0x94, 0xd4, 0x0a, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xf4,
	0xfc, 0x9c, 0xc4, 0xbc, 0x74, 0x7d, 0x88, 0x60, 0x51, 0x62, 0x41, 0x49, 0x7e, 0x91, 0x7e, 0x62,
	0x41, 0xa6, 0x3e, 0xd8, 0x69, 0x49, 0x6c, 0x60, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x0b,
	0xf9, 0x37, 0x2e, 0xba, 0x00, 0x00, 0x00,
}
