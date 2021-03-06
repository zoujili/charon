// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/piotrkowalczuk/charon/pb/rpc/charond/v1/common.proto

package charond // import "github.com/piotrkowalczuk/charon/pb/rpc/charond/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Order represents single field within OrderBy clause.
type Order struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Descending           bool     `protobuf:"varint,2,opt,name=descending,proto3" json:"descending,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_6fa0804b816886fd, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (dst *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(dst, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Order) GetDescending() bool {
	if m != nil {
		return m.Descending
	}
	return false
}

func init() {
	proto.RegisterType((*Order)(nil), "charon.rpc.charond.v1.Order")
}

func init() {
	proto.RegisterFile("github.com/piotrkowalczuk/charon/pb/rpc/charond/v1/common.proto", fileDescriptor_common_6fa0804b816886fd)
}

var fileDescriptor_common_6fa0804b816886fd = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xc8, 0xcc, 0x2f, 0x29, 0xca, 0xce, 0x2f, 0x4f,
	0xcc, 0x49, 0xae, 0x2a, 0xcd, 0xd6, 0x4f, 0xce, 0x48, 0x2c, 0xca, 0xcf, 0xd3, 0x2f, 0x48, 0xd2,
	0x2f, 0x2a, 0x48, 0x86, 0xf2, 0x52, 0xf4, 0xcb, 0x0c, 0xf5, 0x93, 0xf3, 0x73, 0x73, 0xf3, 0xf3,
	0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0x21, 0x12, 0x7a, 0x45, 0x05, 0xc9, 0x7a, 0x50,
	0x35, 0x7a, 0x65, 0x86, 0x4a, 0xd6, 0x5c, 0xac, 0xfe, 0x45, 0x29, 0xa9, 0x45, 0x42, 0x42, 0x5c,
	0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x1c,
	0x17, 0x57, 0x4a, 0x6a, 0x71, 0x72, 0x6a, 0x5e, 0x4a, 0x66, 0x5e, 0xba, 0x04, 0x93, 0x02, 0xa3,
	0x06, 0x47, 0x10, 0x92, 0x88, 0x53, 0x02, 0x97, 0x42, 0x72, 0x7e, 0xae, 0x1e, 0xcc, 0x69, 0xd8,
	0x2c, 0x08, 0x60, 0x8c, 0xb2, 0x22, 0xdd, 0xe9, 0xd6, 0x50, 0x66, 0x12, 0x1b, 0xd8, 0xf1, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x79, 0xac, 0x27, 0xff, 0x00, 0x00, 0x00,
}
