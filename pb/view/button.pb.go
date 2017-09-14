// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/pb/view/button.proto

package view

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import matcha "gomatcha.io/matcha/pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Button struct {
	Str     string        `protobuf:"bytes,1,opt,name=str" json:"str,omitempty"`
	Enabled bool          `protobuf:"varint,2,opt,name=enabled" json:"enabled,omitempty"`
	Color   *matcha.Color `protobuf:"bytes,3,opt,name=color" json:"color,omitempty"`
}

func (m *Button) Reset()                    { *m = Button{} }
func (m *Button) String() string            { return proto.CompactTextString(m) }
func (*Button) ProtoMessage()               {}
func (*Button) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

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

func (m *Button) GetColor() *matcha.Color {
	if m != nil {
		return m.Color
	}
	return nil
}

func init() {
	proto.RegisterType((*Button)(nil), "matcha.view.Button")
}

func init() { proto.RegisterFile("gomatcha.io/matcha/pb/view/button.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xcf, 0xcf, 0x4d,
	0x2c, 0x49, 0xce, 0x48, 0xd4, 0xcb, 0xcc, 0xd7, 0x87, 0xb0, 0xf4, 0x0b, 0x92, 0xf4, 0xcb, 0x32,
	0x53, 0xcb, 0xf5, 0x93, 0x4a, 0x4b, 0x4a, 0xf2, 0xf3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xb8, 0xa1, 0xca, 0x40, 0x32, 0x52, 0x8a, 0xd8, 0x75, 0x65, 0xe6, 0x26, 0xa6, 0xa7, 0x42, 0xd4,
	0x2b, 0x45, 0x72, 0xb1, 0x39, 0x81, 0xf5, 0x0b, 0x09, 0x70, 0x31, 0x17, 0x97, 0x14, 0x49, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x12, 0x5c, 0xec, 0xa9, 0x79, 0x89, 0x49, 0x39,
	0xa9, 0x29, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x30, 0xae, 0x90, 0x32, 0x17, 0x6b, 0x72,
	0x7e, 0x4e, 0x7e, 0x91, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0xb7, 0x11, 0xaf, 0x1e, 0xd4, 0x1a, 0x67,
	0x90, 0x60, 0x10, 0x44, 0xce, 0xc9, 0x9a, 0x4b, 0x2a, 0x33, 0x5f, 0x0f, 0xee, 0x04, 0x28, 0x55,
	0x90, 0x04, 0x76, 0x9b, 0x13, 0x47, 0x40, 0x12, 0xc4, 0xe2, 0x28, 0x16, 0x10, 0x7f, 0x11, 0x13,
	0x8f, 0x2f, 0x58, 0x41, 0x58, 0x66, 0x6a, 0x79, 0x40, 0x52, 0x12, 0x1b, 0xd8, 0x79, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x70, 0x4c, 0xf3, 0xf9, 0x00, 0x00, 0x00,
}