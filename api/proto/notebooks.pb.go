// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notebooks.proto

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

type NotebookExportRequest struct {
	NotebookId           string   `protobuf:"bytes,1,opt,name=notebook_id,json=notebookId,proto3" json:"notebook_id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotebookExportRequest) Reset()         { *m = NotebookExportRequest{} }
func (m *NotebookExportRequest) String() string { return proto.CompactTextString(m) }
func (*NotebookExportRequest) ProtoMessage()    {}
func (*NotebookExportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a2fb83cea3dee99, []int{0}
}

func (m *NotebookExportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotebookExportRequest.Unmarshal(m, b)
}
func (m *NotebookExportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotebookExportRequest.Marshal(b, m, deterministic)
}
func (m *NotebookExportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotebookExportRequest.Merge(m, src)
}
func (m *NotebookExportRequest) XXX_Size() int {
	return xxx_messageInfo_NotebookExportRequest.Size(m)
}
func (m *NotebookExportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotebookExportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotebookExportRequest proto.InternalMessageInfo

func (m *NotebookExportRequest) GetNotebookId() string {
	if m != nil {
		return m.NotebookId
	}
	return ""
}

func (m *NotebookExportRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type NotebookCellRequest struct {
	NotebookId           string   `protobuf:"bytes,1,opt,name=notebook_id,json=notebookId,proto3" json:"notebook_id,omitempty"`
	CellId               string   `protobuf:"bytes,2,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
	Input                string   `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	Offset               uint64   `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                uint64   `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	Type                 string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	CurrentlyEditing     bool     `protobuf:"varint,8,opt,name=currently_editing,json=currentlyEditing,proto3" json:"currently_editing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotebookCellRequest) Reset()         { *m = NotebookCellRequest{} }
func (m *NotebookCellRequest) String() string { return proto.CompactTextString(m) }
func (*NotebookCellRequest) ProtoMessage()    {}
func (*NotebookCellRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a2fb83cea3dee99, []int{1}
}

func (m *NotebookCellRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotebookCellRequest.Unmarshal(m, b)
}
func (m *NotebookCellRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotebookCellRequest.Marshal(b, m, deterministic)
}
func (m *NotebookCellRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotebookCellRequest.Merge(m, src)
}
func (m *NotebookCellRequest) XXX_Size() int {
	return xxx_messageInfo_NotebookCellRequest.Size(m)
}
func (m *NotebookCellRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotebookCellRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotebookCellRequest proto.InternalMessageInfo

func (m *NotebookCellRequest) GetNotebookId() string {
	if m != nil {
		return m.NotebookId
	}
	return ""
}

func (m *NotebookCellRequest) GetCellId() string {
	if m != nil {
		return m.CellId
	}
	return ""
}

func (m *NotebookCellRequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *NotebookCellRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *NotebookCellRequest) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *NotebookCellRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NotebookCellRequest) GetCurrentlyEditing() bool {
	if m != nil {
		return m.CurrentlyEditing
	}
	return false
}

type NotebookMetadata struct {
	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Creator      string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	CreatedTime  int64  `protobuf:"varint,4,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	ModifiedTime int64  `protobuf:"varint,5,opt,name=modified_time,json=modifiedTime,proto3" json:"modified_time,omitempty"`
	NotebookId   string `protobuf:"bytes,7,opt,name=notebook_id,json=notebookId,proto3" json:"notebook_id,omitempty"`
	// Deprecated
	Cells                []string            `protobuf:"bytes,6,rep,name=cells,proto3" json:"cells,omitempty"`
	CellMetadata         []*NotebookCell     `protobuf:"bytes,11,rep,name=cell_metadata,json=cellMetadata,proto3" json:"cell_metadata,omitempty"`
	LatestCellId         string              `protobuf:"bytes,8,opt,name=latest_cell_id,json=latestCellId,proto3" json:"latest_cell_id,omitempty"`
	Hidden               bool                `protobuf:"varint,9,opt,name=hidden,proto3" json:"hidden,omitempty"`
	AvailableDownloads   *AvailableDownloads `protobuf:"bytes,10,opt,name=available_downloads,json=availableDownloads,proto3" json:"available_downloads,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NotebookMetadata) Reset()         { *m = NotebookMetadata{} }
func (m *NotebookMetadata) String() string { return proto.CompactTextString(m) }
func (*NotebookMetadata) ProtoMessage()    {}
func (*NotebookMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a2fb83cea3dee99, []int{2}
}

func (m *NotebookMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotebookMetadata.Unmarshal(m, b)
}
func (m *NotebookMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotebookMetadata.Marshal(b, m, deterministic)
}
func (m *NotebookMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotebookMetadata.Merge(m, src)
}
func (m *NotebookMetadata) XXX_Size() int {
	return xxx_messageInfo_NotebookMetadata.Size(m)
}
func (m *NotebookMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NotebookMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NotebookMetadata proto.InternalMessageInfo

func (m *NotebookMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NotebookMetadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NotebookMetadata) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *NotebookMetadata) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *NotebookMetadata) GetModifiedTime() int64 {
	if m != nil {
		return m.ModifiedTime
	}
	return 0
}

func (m *NotebookMetadata) GetNotebookId() string {
	if m != nil {
		return m.NotebookId
	}
	return ""
}

func (m *NotebookMetadata) GetCells() []string {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *NotebookMetadata) GetCellMetadata() []*NotebookCell {
	if m != nil {
		return m.CellMetadata
	}
	return nil
}

func (m *NotebookMetadata) GetLatestCellId() string {
	if m != nil {
		return m.LatestCellId
	}
	return ""
}

func (m *NotebookMetadata) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *NotebookMetadata) GetAvailableDownloads() *AvailableDownloads {
	if m != nil {
		return m.AvailableDownloads
	}
	return nil
}

type Notebooks struct {
	Items                []*NotebookMetadata `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Notebooks) Reset()         { *m = Notebooks{} }
func (m *Notebooks) String() string { return proto.CompactTextString(m) }
func (*Notebooks) ProtoMessage()    {}
func (*Notebooks) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a2fb83cea3dee99, []int{3}
}

func (m *Notebooks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notebooks.Unmarshal(m, b)
}
func (m *Notebooks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notebooks.Marshal(b, m, deterministic)
}
func (m *Notebooks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notebooks.Merge(m, src)
}
func (m *Notebooks) XXX_Size() int {
	return xxx_messageInfo_Notebooks.Size(m)
}
func (m *Notebooks) XXX_DiscardUnknown() {
	xxx_messageInfo_Notebooks.DiscardUnknown(m)
}

var xxx_messageInfo_Notebooks proto.InternalMessageInfo

func (m *Notebooks) GetItems() []*NotebookMetadata {
	if m != nil {
		return m.Items
	}
	return nil
}

type NotebookCell struct {
	Input     string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output    string   `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Data      string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	CellId    string   `protobuf:"bytes,4,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
	Messages  []string `protobuf:"bytes,5,rep,name=messages,proto3" json:"messages,omitempty"`
	Timestamp int64    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The type of this cell.
	Type                 string   `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	CurrentlyEditing     bool     `protobuf:"varint,8,opt,name=currently_editing,json=currentlyEditing,proto3" json:"currently_editing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotebookCell) Reset()         { *m = NotebookCell{} }
func (m *NotebookCell) String() string { return proto.CompactTextString(m) }
func (*NotebookCell) ProtoMessage()    {}
func (*NotebookCell) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a2fb83cea3dee99, []int{4}
}

func (m *NotebookCell) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotebookCell.Unmarshal(m, b)
}
func (m *NotebookCell) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotebookCell.Marshal(b, m, deterministic)
}
func (m *NotebookCell) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotebookCell.Merge(m, src)
}
func (m *NotebookCell) XXX_Size() int {
	return xxx_messageInfo_NotebookCell.Size(m)
}
func (m *NotebookCell) XXX_DiscardUnknown() {
	xxx_messageInfo_NotebookCell.DiscardUnknown(m)
}

var xxx_messageInfo_NotebookCell proto.InternalMessageInfo

func (m *NotebookCell) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *NotebookCell) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *NotebookCell) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *NotebookCell) GetCellId() string {
	if m != nil {
		return m.CellId
	}
	return ""
}

func (m *NotebookCell) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *NotebookCell) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *NotebookCell) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NotebookCell) GetCurrentlyEditing() bool {
	if m != nil {
		return m.CurrentlyEditing
	}
	return false
}

type NotebookFileUploadRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Filename             string   `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	NotebookId           string   `protobuf:"bytes,3,opt,name=notebook_id,json=notebookId,proto3" json:"notebook_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotebookFileUploadRequest) Reset()         { *m = NotebookFileUploadRequest{} }
func (m *NotebookFileUploadRequest) String() string { return proto.CompactTextString(m) }
func (*NotebookFileUploadRequest) ProtoMessage()    {}
func (*NotebookFileUploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a2fb83cea3dee99, []int{5}
}

func (m *NotebookFileUploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotebookFileUploadRequest.Unmarshal(m, b)
}
func (m *NotebookFileUploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotebookFileUploadRequest.Marshal(b, m, deterministic)
}
func (m *NotebookFileUploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotebookFileUploadRequest.Merge(m, src)
}
func (m *NotebookFileUploadRequest) XXX_Size() int {
	return xxx_messageInfo_NotebookFileUploadRequest.Size(m)
}
func (m *NotebookFileUploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotebookFileUploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotebookFileUploadRequest proto.InternalMessageInfo

func (m *NotebookFileUploadRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *NotebookFileUploadRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *NotebookFileUploadRequest) GetNotebookId() string {
	if m != nil {
		return m.NotebookId
	}
	return ""
}

type NotebookFileUploadResponse struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotebookFileUploadResponse) Reset()         { *m = NotebookFileUploadResponse{} }
func (m *NotebookFileUploadResponse) String() string { return proto.CompactTextString(m) }
func (*NotebookFileUploadResponse) ProtoMessage()    {}
func (*NotebookFileUploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a2fb83cea3dee99, []int{6}
}

func (m *NotebookFileUploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotebookFileUploadResponse.Unmarshal(m, b)
}
func (m *NotebookFileUploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotebookFileUploadResponse.Marshal(b, m, deterministic)
}
func (m *NotebookFileUploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotebookFileUploadResponse.Merge(m, src)
}
func (m *NotebookFileUploadResponse) XXX_Size() int {
	return xxx_messageInfo_NotebookFileUploadResponse.Size(m)
}
func (m *NotebookFileUploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotebookFileUploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotebookFileUploadResponse proto.InternalMessageInfo

func (m *NotebookFileUploadResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*NotebookExportRequest)(nil), "proto.NotebookExportRequest")
	proto.RegisterType((*NotebookCellRequest)(nil), "proto.NotebookCellRequest")
	proto.RegisterType((*NotebookMetadata)(nil), "proto.NotebookMetadata")
	proto.RegisterType((*Notebooks)(nil), "proto.Notebooks")
	proto.RegisterType((*NotebookCell)(nil), "proto.NotebookCell")
	proto.RegisterType((*NotebookFileUploadRequest)(nil), "proto.NotebookFileUploadRequest")
	proto.RegisterType((*NotebookFileUploadResponse)(nil), "proto.NotebookFileUploadResponse")
}

func init() {
	proto.RegisterFile("notebooks.proto", fileDescriptor_6a2fb83cea3dee99)
}

var fileDescriptor_6a2fb83cea3dee99 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0xb2, 0xf9, 0x9c, 0x4d, 0xa1, 0xb8, 0xd0, 0xba, 0x11, 0x12, 0x4b, 0xe0, 0x10, 0x09,
	0x91, 0x88, 0x72, 0x41, 0xdc, 0xa0, 0x14, 0xa9, 0x08, 0x38, 0xac, 0xe0, 0xc2, 0x25, 0x72, 0xd7,
	0x93, 0x60, 0xe1, 0x5d, 0x2f, 0x6b, 0xa7, 0x69, 0x7f, 0x10, 0xff, 0x88, 0x9f, 0xc0, 0x0f, 0x41,
	0xeb, 0x8f, 0x34, 0x4d, 0x39, 0x20, 0x4e, 0x3b, 0xef, 0xf9, 0xc5, 0x9e, 0x37, 0xf3, 0x02, 0x77,
	0x4a, 0x65, 0xf0, 0x4c, 0xa9, 0xef, 0x7a, 0x52, 0xd5, 0xca, 0x28, 0xd2, 0xb6, 0x9f, 0x61, 0x32,
	0x97, 0x6a, 0xe5, 0xb9, 0xd1, 0x07, 0xb8, 0xff, 0xc9, 0xcb, 0x4e, 0x2e, 0x2a, 0x55, 0x9b, 0x0c,
	0x7f, 0x2c, 0x51, 0x1b, 0xf2, 0x10, 0x92, 0xf0, 0xfb, 0x99, 0xe0, 0x34, 0x4a, 0xa3, 0x71, 0x3f,
	0x83, 0x40, 0x9d, 0x72, 0x42, 0xa0, 0x65, 0x2e, 0x2b, 0xa4, 0xb7, 0xec, 0x89, 0xad, 0x47, 0xbf,
	0x22, 0xd8, 0x0b, 0xd7, 0x1d, 0xa3, 0x94, 0xff, 0x7c, 0xd9, 0x01, 0x74, 0x73, 0x94, 0xb2, 0x39,
	0x74, 0xf7, 0x75, 0x1a, 0x78, 0xca, 0xc9, 0x3d, 0x68, 0x8b, 0xb2, 0x5a, 0x1a, 0x1a, 0x5b, 0xda,
	0x01, 0xb2, 0x0f, 0x1d, 0x35, 0x9f, 0x6b, 0x34, 0xb4, 0x95, 0x46, 0xe3, 0x56, 0xe6, 0x51, 0xa3,
	0xce, 0xd5, 0xb2, 0x34, 0xb4, 0x6d, 0x69, 0x07, 0xd6, 0x9d, 0x76, 0xae, 0x3a, 0x25, 0x4f, 0xe1,
	0x6e, 0xbe, 0xac, 0x6b, 0x2c, 0x8d, 0xbc, 0x9c, 0x21, 0x17, 0x46, 0x94, 0x0b, 0xda, 0x4b, 0xa3,
	0x71, 0x2f, 0xdb, 0x5d, 0x1f, 0x9c, 0x38, 0x7e, 0xf4, 0x33, 0x86, 0xdd, 0x60, 0xeb, 0x23, 0x1a,
	0xc6, 0x99, 0x61, 0xcd, 0xad, 0x25, 0x2b, 0xd0, 0x9b, 0xb1, 0x35, 0x49, 0x21, 0xe1, 0xa8, 0xf3,
	0x5a, 0x54, 0x46, 0xa8, 0xd2, 0x5b, 0xd9, 0xa4, 0x08, 0x85, 0x6e, 0x5e, 0x23, 0x33, 0xaa, 0xf6,
	0x8e, 0x02, 0x24, 0x8f, 0x60, 0x60, 0x4b, 0xe4, 0x33, 0x23, 0x0a, 0xb4, 0xce, 0xe2, 0x2c, 0xf1,
	0xdc, 0x67, 0x51, 0x20, 0x79, 0x0c, 0x3b, 0x85, 0xe2, 0x62, 0x2e, 0x82, 0xa6, 0x6d, 0x35, 0x83,
	0x40, 0x5a, 0xd1, 0xd6, 0xac, 0xbb, 0x37, 0x66, 0xdd, 0x0c, 0x09, 0xa5, 0xd4, 0xb4, 0x93, 0xc6,
	0xcd, 0x48, 0x2d, 0x20, 0x2f, 0x61, 0xc7, 0x6e, 0xa0, 0xf0, 0xfe, 0x68, 0x92, 0xc6, 0xe3, 0xe4,
	0x68, 0xcf, 0xe5, 0x64, 0x72, 0x6d, 0xab, 0x83, 0x46, 0xb9, 0x1e, 0xc4, 0x13, 0xb8, 0x2d, 0x99,
	0x41, 0x6d, 0x66, 0x61, 0x85, 0x3d, 0xfb, 0xe6, 0xc0, 0xb1, 0xc7, 0x6e, 0x91, 0xfb, 0xd0, 0xf9,
	0x26, 0x38, 0xc7, 0x92, 0xf6, 0xed, 0x94, 0x3d, 0x22, 0xef, 0x61, 0x8f, 0x9d, 0x33, 0x21, 0xd9,
	0x99, 0xc4, 0x19, 0x57, 0xab, 0x52, 0x2a, 0xc6, 0x35, 0x85, 0x34, 0x1a, 0x27, 0x47, 0x87, 0xfe,
	0xf5, 0xd7, 0x41, 0xf1, 0x36, 0x08, 0x32, 0xc2, 0x6e, 0x70, 0xa3, 0x57, 0xd0, 0x0f, 0x7d, 0x6a,
	0xf2, 0x0c, 0xda, 0xc2, 0x60, 0xa1, 0x69, 0x64, 0x8d, 0x1c, 0x6c, 0x19, 0x09, 0xed, 0x67, 0x4e,
	0x35, 0xfa, 0x1d, 0xc1, 0x60, 0xd3, 0xe4, 0x55, 0xf2, 0xa2, 0xed, 0xe4, 0x2d, 0x4d, 0x43, 0xfb,
	0x9c, 0x3a, 0xd4, 0xa4, 0xc1, 0x4e, 0xcd, 0x2d, 0xd5, 0xd6, 0x9b, 0xa1, 0x6e, 0x5d, 0x0b, 0xf5,
	0x10, 0x7a, 0x05, 0x6a, 0xcd, 0x16, 0xa8, 0x69, 0xdb, 0x2e, 0x61, 0x8d, 0xc9, 0x03, 0xe8, 0x37,
	0xab, 0xd5, 0x86, 0x15, 0x95, 0x4d, 0x6c, 0x9c, 0x5d, 0x11, 0xeb, 0x28, 0x77, 0xff, 0x37, 0xca,
	0x12, 0x0e, 0x83, 0xcb, 0x77, 0x42, 0xe2, 0x97, 0xaa, 0x99, 0x5c, 0xf8, 0x9b, 0x06, 0x13, 0xd1,
	0x86, 0x89, 0x21, 0xf4, 0xe6, 0x42, 0xa2, 0x8d, 0xba, 0xb3, 0xbc, 0xc6, 0xdb, 0x51, 0x8b, 0xb7,
	0xa3, 0x36, 0x9a, 0xc0, 0xf0, 0x6f, 0xaf, 0xe9, 0x4a, 0x95, 0x1a, 0xc9, 0x2e, 0xc4, 0xcb, 0x5a,
	0xfa, 0xd7, 0x9a, 0xf2, 0xcd, 0xf3, 0xaf, 0xd3, 0xd5, 0x6a, 0x35, 0x39, 0x47, 0xa9, 0x72, 0xc1,
	0xf1, 0x62, 0x92, 0xab, 0x62, 0xba, 0x50, 0x92, 0x95, 0x8b, 0xa9, 0x23, 0x6b, 0x56, 0x19, 0x55,
	0x4f, 0x59, 0x25, 0xa6, 0x76, 0x9f, 0x67, 0x1d, 0xfb, 0x79, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff,
	0xde, 0xe7, 0xac, 0x44, 0xee, 0x04, 0x00, 0x00,
}
