// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/gomatcha/matcha/proto/view/button.proto

package view

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "github.com/gomatcha/matcha/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Button struct {
	Str                  string        `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	Enabled              bool          `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Color                *proto1.Color `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Button) Reset()         { *m = Button{} }
func (m *Button) String() string { return proto.CompactTextString(m) }
func (*Button) ProtoMessage()    {}
func (*Button) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5f2c18f180cc1e3, []int{0}
}
func (m *Button) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Button.Unmarshal(m, b)
}
func (m *Button) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Button.Marshal(b, m, deterministic)
}
func (m *Button) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Button.Merge(m, src)
}
func (m *Button) XXX_Size() int {
	return xxx_messageInfo_Button.Size(m)
}
func (m *Button) XXX_DiscardUnknown() {
	xxx_messageInfo_Button.DiscardUnknown(m)
}

var xxx_messageInfo_Button proto.InternalMessageInfo

func (m *Button) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *Button) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Button) GetColor() *proto1.Color {
	if m != nil {
		return m.Color
	}
	return nil
}

func init() {
	proto.RegisterType((*Button)(nil), "matcha.view.Button")
}

func init() {
	proto.RegisterFile("github.com/gomatcha/matcha/proto/view/button.proto", fileDescriptor_b5f2c18f180cc1e3)
}

var fileDescriptor_b5f2c18f180cc1e3 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0xd4,
	0x87, 0x52, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0xfa, 0x65, 0x99, 0xa9, 0xe5, 0xfa, 0x49, 0xa5, 0x25,
	0x25, 0xf9, 0x79, 0x7a, 0x60, 0x11, 0x21, 0x6e, 0x88, 0xbc, 0x1e, 0x48, 0x46, 0x4a, 0x87, 0xa0,
	0x01, 0x99, 0xb9, 0x89, 0xe9, 0xa9, 0x10, 0xad, 0x4a, 0x91, 0x5c, 0x6c, 0x4e, 0x60, 0xa3, 0x84,
	0x04, 0xb8, 0x98, 0x8b, 0x4b, 0x8a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21,
	0x09, 0x2e, 0xf6, 0xd4, 0xbc, 0xc4, 0xa4, 0x9c, 0xd4, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x8e,
	0x20, 0x18, 0x57, 0x48, 0x99, 0x8b, 0x35, 0x39, 0x3f, 0x27, 0xbf, 0x48, 0x82, 0x59, 0x81, 0x51,
	0x83, 0xdb, 0x88, 0x57, 0x0f, 0xea, 0x00, 0x67, 0x90, 0x60, 0x10, 0x44, 0xce, 0xc9, 0x8d, 0x4b,
	0x35, 0x39, 0x3f, 0x57, 0x0f, 0xea, 0x1c, 0x98, 0x53, 0x60, 0x4a, 0xc1, 0xd6, 0x83, 0x5d, 0xec,
	0xc4, 0x11, 0x90, 0x04, 0x71, 0x43, 0x14, 0x0b, 0x88, 0xbf, 0x88, 0x89, 0xc7, 0x17, 0xac, 0x26,
	0x2c, 0x33, 0xb5, 0x3c, 0x20, 0x29, 0x89, 0x0d, 0xac, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x5a, 0x4b, 0xe2, 0xb7, 0x1a, 0x01, 0x00, 0x00,
}
