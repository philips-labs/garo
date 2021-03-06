// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/garo/service.proto

package garo

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// GetRepoConfigurationRequest request the configuration for a given repository and organization
type GetRepoConfigurationRequest struct {
	Organisation         string   `protobuf:"bytes,1,opt,name=organisation,proto3" json:"organisation,omitempty"`
	Repository           string   `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRepoConfigurationRequest) Reset()         { *m = GetRepoConfigurationRequest{} }
func (m *GetRepoConfigurationRequest) String() string { return proto.CompactTextString(m) }
func (*GetRepoConfigurationRequest) ProtoMessage()    {}
func (*GetRepoConfigurationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb44b64dbde6c2, []int{0}
}

func (m *GetRepoConfigurationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRepoConfigurationRequest.Unmarshal(m, b)
}
func (m *GetRepoConfigurationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRepoConfigurationRequest.Marshal(b, m, deterministic)
}
func (m *GetRepoConfigurationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRepoConfigurationRequest.Merge(m, src)
}
func (m *GetRepoConfigurationRequest) XXX_Size() int {
	return xxx_messageInfo_GetRepoConfigurationRequest.Size(m)
}
func (m *GetRepoConfigurationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRepoConfigurationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRepoConfigurationRequest proto.InternalMessageInfo

func (m *GetRepoConfigurationRequest) GetOrganisation() string {
	if m != nil {
		return m.Organisation
	}
	return ""
}

func (m *GetRepoConfigurationRequest) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

// RepoConfigurationResponse returns agent configuration options
type RepoConfigurationResponse struct {
	Repository           string   `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	MaxConcurrentRunners uint32   `protobuf:"varint,2,opt,name=maxConcurrentRunners,proto3" json:"maxConcurrentRunners,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepoConfigurationResponse) Reset()         { *m = RepoConfigurationResponse{} }
func (m *RepoConfigurationResponse) String() string { return proto.CompactTextString(m) }
func (*RepoConfigurationResponse) ProtoMessage()    {}
func (*RepoConfigurationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07eb44b64dbde6c2, []int{1}
}

func (m *RepoConfigurationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepoConfigurationResponse.Unmarshal(m, b)
}
func (m *RepoConfigurationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepoConfigurationResponse.Marshal(b, m, deterministic)
}
func (m *RepoConfigurationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepoConfigurationResponse.Merge(m, src)
}
func (m *RepoConfigurationResponse) XXX_Size() int {
	return xxx_messageInfo_RepoConfigurationResponse.Size(m)
}
func (m *RepoConfigurationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RepoConfigurationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RepoConfigurationResponse proto.InternalMessageInfo

func (m *RepoConfigurationResponse) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *RepoConfigurationResponse) GetMaxConcurrentRunners() uint32 {
	if m != nil {
		return m.MaxConcurrentRunners
	}
	return 0
}

func init() {
	proto.RegisterType((*GetRepoConfigurationRequest)(nil), "philips.garo.garo.GetRepoConfigurationRequest")
	proto.RegisterType((*RepoConfigurationResponse)(nil), "philips.garo.garo.RepoConfigurationResponse")
}

func init() { proto.RegisterFile("rpc/garo/service.proto", fileDescriptor_07eb44b64dbde6c2) }

var fileDescriptor_07eb44b64dbde6c2 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0x89, 0xc8, 0x81, 0x83, 0x16, 0x06, 0x91, 0xbb, 0x13, 0x44, 0xb6, 0xb2, 0x90, 0x1c,
	0x9c, 0xbf, 0x40, 0xaf, 0xb0, 0x8f, 0x9d, 0x5d, 0x5c, 0xc6, 0x18, 0xd0, 0x99, 0x38, 0x93, 0x88,
	0xfa, 0x47, 0xfc, 0xbb, 0x72, 0x39, 0x0b, 0xf7, 0x5c, 0x6d, 0x52, 0xbc, 0xf9, 0x26, 0x8f, 0xf7,
	0x06, 0x8e, 0x25, 0xf7, 0x8b, 0x18, 0x84, 0x17, 0x8a, 0xf2, 0x9a, 0x7a, 0x74, 0x59, 0xb8, 0xb0,
	0x3d, 0xcc, 0x8f, 0xe9, 0x29, 0x65, 0x75, 0xeb, 0x59, 0x7b, 0xba, 0x00, 0x27, 0x37, 0x58, 0x3c,
	0x66, 0x5e, 0x31, 0x3d, 0xa4, 0x58, 0x25, 0x94, 0xc4, 0xe4, 0xf1, 0xa5, 0xa2, 0x16, 0xdb, 0xc1,
	0x3e, 0x4b, 0x0c, 0x94, 0xb4, 0xc9, 0x53, 0x73, 0x66, 0xce, 0xf7, 0xfc, 0x40, 0xb3, 0xa7, 0x00,
	0x82, 0x99, 0x35, 0x15, 0x96, 0xf7, 0xe9, 0x4e, 0x23, 0x7e, 0x28, 0x1d, 0xc3, 0x6c, 0xe4, 0x7f,
	0xcd, 0x4c, 0x8a, 0x5b, 0xcb, 0x66, 0x7b, 0xd9, 0x2e, 0xe1, 0xe8, 0x39, 0xbc, 0xad, 0x98, 0xfa,
	0x2a, 0x82, 0x54, 0x7c, 0x25, 0x42, 0xd1, 0x66, 0x73, 0xe0, 0x47, 0x67, 0xcb, 0x4f, 0x03, 0xb3,
	0xab, 0x88, 0x54, 0x06, 0x96, 0xb7, 0x9b, 0x2a, 0xec, 0x07, 0xcc, 0xbf, 0x13, 0x6f, 0x2c, 0x06,
	0x90, 0x75, 0xee, 0x57, 0x47, 0xee, 0x9f, 0x82, 0xe6, 0x17, 0x23, 0xfc, 0x9f, 0x69, 0xaf, 0x27,
	0x77, 0xbb, 0x6b, 0xe2, 0x7e, 0xd2, 0xee, 0x71, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x52,
	0x9f, 0xc3, 0xa9, 0x01, 0x00, 0x00,
}
