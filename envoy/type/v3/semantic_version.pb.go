// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/v3/semantic_version.proto

package envoy_type_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type SemanticVersion struct {
	MajorNumber          uint32   `protobuf:"varint,1,opt,name=major_number,json=majorNumber,proto3" json:"major_number,omitempty"`
	MinorNumber          uint32   `protobuf:"varint,2,opt,name=minor_number,json=minorNumber,proto3" json:"minor_number,omitempty"`
	Patch                uint32   `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SemanticVersion) Reset()         { *m = SemanticVersion{} }
func (m *SemanticVersion) String() string { return proto.CompactTextString(m) }
func (*SemanticVersion) ProtoMessage()    {}
func (*SemanticVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d8ec35e89ef71eb, []int{0}
}

func (m *SemanticVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SemanticVersion.Unmarshal(m, b)
}
func (m *SemanticVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SemanticVersion.Marshal(b, m, deterministic)
}
func (m *SemanticVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SemanticVersion.Merge(m, src)
}
func (m *SemanticVersion) XXX_Size() int {
	return xxx_messageInfo_SemanticVersion.Size(m)
}
func (m *SemanticVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_SemanticVersion.DiscardUnknown(m)
}

var xxx_messageInfo_SemanticVersion proto.InternalMessageInfo

func (m *SemanticVersion) GetMajorNumber() uint32 {
	if m != nil {
		return m.MajorNumber
	}
	return 0
}

func (m *SemanticVersion) GetMinorNumber() uint32 {
	if m != nil {
		return m.MinorNumber
	}
	return 0
}

func (m *SemanticVersion) GetPatch() uint32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func init() {
	proto.RegisterType((*SemanticVersion)(nil), "envoy.type.v3.SemanticVersion")
}

func init() {
	proto.RegisterFile("envoy/type/v3/semantic_version.proto", fileDescriptor_5d8ec35e89ef71eb)
}

var fileDescriptor_5d8ec35e89ef71eb = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x33, 0xd6, 0x2f, 0x4e, 0xcd, 0x4d, 0xcc, 0x2b,
	0xc9, 0x4c, 0x8e, 0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x05, 0xab, 0xd2, 0x03, 0xa9, 0xd2, 0x2b, 0x33, 0x96, 0x92, 0x2d, 0x4d, 0x29, 0x48,
	0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0x2e, 0x49,
	0x2c, 0x29, 0x2d, 0x86, 0xa8, 0x96, 0x52, 0xc4, 0x90, 0x86, 0x9a, 0x96, 0x99, 0x97, 0x0e, 0x51,
	0xa2, 0x34, 0x81, 0x91, 0x8b, 0x3f, 0x18, 0x6a, 0x57, 0x18, 0x44, 0x52, 0x48, 0x91, 0x8b, 0x27,
	0x37, 0x31, 0x2b, 0xbf, 0x28, 0x3e, 0xaf, 0x34, 0x37, 0x29, 0xb5, 0x48, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x37, 0x88, 0x1b, 0x2c, 0xe6, 0x07, 0x16, 0x02, 0x2b, 0xc9, 0xcc, 0x43, 0x28, 0x61, 0x82,
	0x2a, 0x01, 0x89, 0x41, 0x95, 0x88, 0x70, 0xb1, 0x16, 0x24, 0x96, 0x24, 0x67, 0x48, 0x30, 0x83,
	0xe5, 0x20, 0x1c, 0x2b, 0xc5, 0x59, 0x47, 0x3b, 0xe4, 0x64, 0xb8, 0xa4, 0x90, 0xfc, 0x81, 0x66,
	0xbd, 0x93, 0xed, 0xae, 0x86, 0x13, 0x17, 0xd9, 0x98, 0x04, 0x98, 0xb8, 0xa4, 0x33, 0xf3, 0xf5,
	0xc0, 0x0a, 0x0b, 0x8a, 0xf2, 0x2b, 0x2a, 0xf5, 0x50, 0xfc, 0xee, 0x24, 0x82, 0xa6, 0x2f, 0x00,
	0xe4, 0x9f, 0x00, 0xc6, 0x24, 0x36, 0xb0, 0xc7, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe8,
	0x8d, 0xa3, 0xd0, 0x51, 0x01, 0x00, 0x00,
}
