// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/function.proto

package envoy_api_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_ "github.com/envoyproxy/go-control-plane/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type Function struct {
	Uri                  string             `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Subset               string             `protobuf:"bytes,2,opt,name=subset,proto3" json:"subset,omitempty"`
	Timeout              int32              `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Glimit               *TrafficLimit      `protobuf:"bytes,4,opt,name=glimit,proto3" json:"glimit,omitempty"`
	Breakers             *core.RollBreakers `protobuf:"bytes,5,opt,name=breakers,proto3" json:"breakers,omitempty"`
	Degrade              *Downgrade         `protobuf:"bytes,6,opt,name=degrade,proto3" json:"degrade,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Function) Reset()         { *m = Function{} }
func (m *Function) String() string { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()    {}
func (*Function) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d2470d82096fbb, []int{0}
}

func (m *Function) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Function.Unmarshal(m, b)
}
func (m *Function) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Function.Marshal(b, m, deterministic)
}
func (m *Function) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Function.Merge(m, src)
}
func (m *Function) XXX_Size() int {
	return xxx_messageInfo_Function.Size(m)
}
func (m *Function) XXX_DiscardUnknown() {
	xxx_messageInfo_Function.DiscardUnknown(m)
}

var xxx_messageInfo_Function proto.InternalMessageInfo

func (m *Function) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Function) GetSubset() string {
	if m != nil {
		return m.Subset
	}
	return ""
}

func (m *Function) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Function) GetGlimit() *TrafficLimit {
	if m != nil {
		return m.Glimit
	}
	return nil
}

func (m *Function) GetBreakers() *core.RollBreakers {
	if m != nil {
		return m.Breakers
	}
	return nil
}

func (m *Function) GetDegrade() *Downgrade {
	if m != nil {
		return m.Degrade
	}
	return nil
}

type TrafficLimit struct {
	WindowMils           int32    `protobuf:"varint,1,opt,name=windowMils,proto3" json:"windowMils,omitempty"`
	BucketMils           int32    `protobuf:"varint,2,opt,name=bucketMils,proto3" json:"bucketMils,omitempty"`
	Times                int32    `protobuf:"varint,3,opt,name=times,proto3" json:"times,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficLimit) Reset()         { *m = TrafficLimit{} }
func (m *TrafficLimit) String() string { return proto.CompactTextString(m) }
func (*TrafficLimit) ProtoMessage()    {}
func (*TrafficLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d2470d82096fbb, []int{1}
}

func (m *TrafficLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficLimit.Unmarshal(m, b)
}
func (m *TrafficLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficLimit.Marshal(b, m, deterministic)
}
func (m *TrafficLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficLimit.Merge(m, src)
}
func (m *TrafficLimit) XXX_Size() int {
	return xxx_messageInfo_TrafficLimit.Size(m)
}
func (m *TrafficLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficLimit.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficLimit proto.InternalMessageInfo

func (m *TrafficLimit) GetWindowMils() int32 {
	if m != nil {
		return m.WindowMils
	}
	return 0
}

func (m *TrafficLimit) GetBucketMils() int32 {
	if m != nil {
		return m.BucketMils
	}
	return 0
}

func (m *TrafficLimit) GetTimes() int32 {
	if m != nil {
		return m.Times
	}
	return 0
}

type Downgrade struct {
	Enable               bool     `protobuf:"varint,3,opt,name=enable,proto3" json:"enable,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Downgrade) Reset()         { *m = Downgrade{} }
func (m *Downgrade) String() string { return proto.CompactTextString(m) }
func (*Downgrade) ProtoMessage()    {}
func (*Downgrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d2470d82096fbb, []int{2}
}

func (m *Downgrade) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Downgrade.Unmarshal(m, b)
}
func (m *Downgrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Downgrade.Marshal(b, m, deterministic)
}
func (m *Downgrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Downgrade.Merge(m, src)
}
func (m *Downgrade) XXX_Size() int {
	return xxx_messageInfo_Downgrade.Size(m)
}
func (m *Downgrade) XXX_DiscardUnknown() {
	xxx_messageInfo_Downgrade.DiscardUnknown(m)
}

var xxx_messageInfo_Downgrade proto.InternalMessageInfo

func (m *Downgrade) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *Downgrade) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Function)(nil), "envoy.api.v2.Function")
	proto.RegisterType((*TrafficLimit)(nil), "envoy.api.v2.TrafficLimit")
	proto.RegisterType((*Downgrade)(nil), "envoy.api.v2.Downgrade")
}

func init() { proto.RegisterFile("envoy/api/v2/function.proto", fileDescriptor_c2d2470d82096fbb) }

var fileDescriptor_c2d2470d82096fbb = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x8e, 0xd3, 0x30,
	0x14, 0xc6, 0x95, 0xce, 0x34, 0xd3, 0x9a, 0x41, 0x1a, 0x59, 0x88, 0x09, 0x01, 0x4a, 0x55, 0x58,
	0x74, 0x95, 0x88, 0xce, 0x82, 0x05, 0xbb, 0x0a, 0xb1, 0x02, 0x69, 0x14, 0xc1, 0x1a, 0x39, 0xc9,
	0x4b, 0x64, 0x8d, 0x6b, 0x5b, 0xfe, 0x93, 0x92, 0x1d, 0x37, 0x60, 0xcb, 0x59, 0x38, 0x01, 0x5b,
	0xae, 0x32, 0x07, 0x40, 0x28, 0x8e, 0xd3, 0x69, 0x06, 0xb1, 0xf3, 0xf3, 0xef, 0xfb, 0xe2, 0xf7,
	0xde, 0x17, 0xf4, 0x14, 0x78, 0x23, 0xda, 0x94, 0x48, 0x9a, 0x36, 0x9b, 0xb4, 0xb2, 0xbc, 0x30,
	0x54, 0xf0, 0x44, 0x2a, 0x61, 0x04, 0x3e, 0x77, 0x30, 0x21, 0x92, 0x26, 0xcd, 0x26, 0x7e, 0x35,
	0x92, 0x16, 0x42, 0x41, 0xaa, 0x04, 0x63, 0x5f, 0x72, 0x05, 0xe4, 0x06, 0x54, 0xef, 0x89, 0xa3,
	0x5e, 0x65, 0x5a, 0x09, 0xa9, 0x04, 0x55, 0x00, 0x37, 0x9e, 0x3c, 0xa9, 0x85, 0xa8, 0x19, 0xa4,
	0xae, 0xca, 0x6d, 0x95, 0x12, 0xde, 0x7a, 0xb4, 0xb8, 0x8f, 0x4a, 0xab, 0xc8, 0x5d, 0x23, 0xf1,
	0xb3, 0xfb, 0x5c, 0x1b, 0x65, 0x0b, 0xf3, 0x3f, 0xf7, 0x5e, 0x11, 0x29, 0x41, 0x69, 0xcf, 0x5f,
	0xfa, 0xc6, 0x39, 0x17, 0xc6, 0x7d, 0x55, 0xa7, 0x25, 0x48, 0x05, 0xc5, 0xf1, 0x13, 0x0b, 0x5b,
	0x4a, 0x32, 0xd2, 0xec, 0x68, 0xad, 0x88, 0x01, 0xcf, 0x9f, 0xff, 0xc3, 0xb5, 0x21, 0xc6, 0x0e,
	0x6f, 0x5c, 0x36, 0x84, 0xd1, 0x92, 0x18, 0x48, 0x87, 0x43, 0x0f, 0x56, 0xb7, 0x01, 0x9a, 0xbd,
	0xf7, 0x6b, 0xc5, 0x17, 0xe8, 0xc4, 0x2a, 0x1a, 0x05, 0xcb, 0x60, 0x3d, 0xcf, 0xba, 0x23, 0x7e,
	0x8c, 0x42, 0x6d, 0x73, 0x0d, 0x26, 0x9a, 0xb8, 0x4b, 0x5f, 0xe1, 0x08, 0x9d, 0x19, 0xba, 0x03,
	0x61, 0x4d, 0x74, 0xb2, 0x0c, 0xd6, 0xd3, 0x6c, 0x28, 0xf1, 0x06, 0x85, 0x35, 0xa3, 0x3b, 0x6a,
	0xa2, 0xd3, 0x65, 0xb0, 0x7e, 0xb0, 0x89, 0x93, 0xe3, 0x94, 0x92, 0x4f, 0x8a, 0x54, 0x15, 0x2d,
	0x3e, 0x74, 0x8a, 0xcc, 0x2b, 0xf1, 0x5b, 0x34, 0xf3, 0x29, 0xe9, 0x68, 0xea, 0x5c, 0x2f, 0xc6,
	0xae, 0x2e, 0xcd, 0x24, 0x13, 0x8c, 0x6d, 0xbd, 0x2c, 0x3b, 0x18, 0xf0, 0x6b, 0x74, 0x56, 0x42,
	0xad, 0x48, 0x09, 0x51, 0xe8, 0xbc, 0x97, 0x63, 0xef, 0x3b, 0xb1, 0xe7, 0x0e, 0x67, 0x83, 0x6e,
	0x55, 0xa2, 0xf3, 0xe3, 0x3e, 0xf0, 0x02, 0xa1, 0x3d, 0xe5, 0xa5, 0xd8, 0x7f, 0xa4, 0x4c, 0xbb,
	0xf1, 0xa7, 0xd9, 0xd1, 0x4d, 0xc7, 0x73, 0x5b, 0xdc, 0x80, 0x71, 0x7c, 0xd2, 0xf3, 0xbb, 0x1b,
	0xfc, 0x08, 0x4d, 0xbb, 0xf1, 0xb5, 0xdf, 0x45, 0x5f, 0xac, 0xde, 0xa0, 0xf9, 0xe1, 0xed, 0x6e,
	0x91, 0xc0, 0x49, 0xce, 0xc0, 0x69, 0x66, 0x99, 0xaf, 0x30, 0x46, 0xa7, 0xb9, 0x28, 0x5b, 0xb7,
	0xac, 0x79, 0xe6, 0xce, 0xdb, 0xcf, 0xb7, 0x3f, 0xfe, 0x7c, 0x9f, 0xc6, 0xb8, 0xff, 0x57, 0x93,
	0x42, 0xf0, 0x8a, 0xd6, 0xc9, 0xe1, 0xe7, 0x6f, 0xae, 0x7e, 0x7e, 0xfb, 0xf5, 0x3b, 0x9c, 0x5c,
	0x04, 0x28, 0xa6, 0xa2, 0x1f, 0x56, 0x2a, 0xf1, 0xb5, 0x1d, 0xcd, 0xbd, 0x7d, 0x38, 0xc4, 0x7a,
	0xdd, 0x05, 0x7d, 0x1d, 0xe4, 0xa1, 0x4b, 0xfc, 0xea, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2,
	0x04, 0xe6, 0xa8, 0x54, 0x03, 0x00, 0x00,
}
