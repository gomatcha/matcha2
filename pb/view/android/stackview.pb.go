// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/pb/view/android/stackview.proto

package android

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import matcha "gomatcha.io/matcha/pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type StackChildView struct {
	ScreenId int64 `protobuf:"varint,3,opt,name=screenId" json:"screenId,omitempty"`
}

func (m *StackChildView) Reset()                    { *m = StackChildView{} }
func (m *StackChildView) String() string            { return proto.CompactTextString(m) }
func (*StackChildView) ProtoMessage()               {}
func (*StackChildView) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *StackChildView) GetScreenId() int64 {
	if m != nil {
		return m.ScreenId
	}
	return 0
}

type StackView struct {
	Children []*StackChildView `protobuf:"bytes,1,rep,name=children" json:"children,omitempty"`
}

func (m *StackView) Reset()                    { *m = StackView{} }
func (m *StackView) String() string            { return proto.CompactTextString(m) }
func (*StackView) ProtoMessage()               {}
func (*StackView) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *StackView) GetChildren() []*StackChildView {
	if m != nil {
		return m.Children
	}
	return nil
}

type StackBar struct {
	Title            string          `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Subtitle         string          `protobuf:"bytes,3,opt,name=subtitle" json:"subtitle,omitempty"`
	Color            *matcha.Color   `protobuf:"bytes,4,opt,name=color" json:"color,omitempty"`
	Items            []*StackBarItem `protobuf:"bytes,5,rep,name=items" json:"items,omitempty"`
	BackButtonHidden bool            `protobuf:"varint,2,opt,name=backButtonHidden" json:"backButtonHidden,omitempty"`
}

func (m *StackBar) Reset()                    { *m = StackBar{} }
func (m *StackBar) String() string            { return proto.CompactTextString(m) }
func (*StackBar) ProtoMessage()               {}
func (*StackBar) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *StackBar) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *StackBar) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *StackBar) GetColor() *matcha.Color {
	if m != nil {
		return m.Color
	}
	return nil
}

func (m *StackBar) GetItems() []*StackBarItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *StackBar) GetBackButtonHidden() bool {
	if m != nil {
		return m.BackButtonHidden
	}
	return false
}

type StackBarItem struct {
	Title    string                  `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Icon     *matcha.ImageOrResource `protobuf:"bytes,3,opt,name=icon" json:"icon,omitempty"`
	IconTint *matcha.Color           `protobuf:"bytes,2,opt,name=iconTint" json:"iconTint,omitempty"`
	Disabled bool                    `protobuf:"varint,4,opt,name=Disabled" json:"Disabled,omitempty"`
}

func (m *StackBarItem) Reset()                    { *m = StackBarItem{} }
func (m *StackBarItem) String() string            { return proto.CompactTextString(m) }
func (*StackBarItem) ProtoMessage()               {}
func (*StackBarItem) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *StackBarItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *StackBarItem) GetIcon() *matcha.ImageOrResource {
	if m != nil {
		return m.Icon
	}
	return nil
}

func (m *StackBarItem) GetIconTint() *matcha.Color {
	if m != nil {
		return m.IconTint
	}
	return nil
}

func (m *StackBarItem) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

type StackEvent struct {
	Id []int64 `protobuf:"varint,1,rep,packed,name=id" json:"id,omitempty"`
}

func (m *StackEvent) Reset()                    { *m = StackEvent{} }
func (m *StackEvent) String() string            { return proto.CompactTextString(m) }
func (*StackEvent) ProtoMessage()               {}
func (*StackEvent) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *StackEvent) GetId() []int64 {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*StackChildView)(nil), "matcha.view.android.StackChildView")
	proto.RegisterType((*StackView)(nil), "matcha.view.android.StackView")
	proto.RegisterType((*StackBar)(nil), "matcha.view.android.StackBar")
	proto.RegisterType((*StackBarItem)(nil), "matcha.view.android.StackBarItem")
	proto.RegisterType((*StackEvent)(nil), "matcha.view.android.StackEvent")
}

func init() { proto.RegisterFile("gomatcha.io/matcha/pb/view/android/stackview.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x95, 0x93, 0x0d, 0x64, 0x27, 0xb0, 0x20, 0x83, 0x44, 0xb4, 0xe2, 0x90, 0xcd, 0x5e, 0xc2,
	0x87, 0x12, 0x29, 0x1c, 0x38, 0x22, 0xb2, 0x20, 0x51, 0xa9, 0x88, 0xca, 0x20, 0x0e, 0xdc, 0x9c,
	0xd8, 0x6a, 0x2d, 0x12, 0xbb, 0x72, 0xdc, 0xf6, 0xaf, 0x70, 0xe6, 0xc7, 0xf0, 0xbb, 0x90, 0x9d,
	0x34, 0xda, 0xaa, 0xed, 0xc9, 0x9e, 0xf7, 0xde, 0x8c, 0xdf, 0xcc, 0x18, 0xca, 0xa5, 0xea, 0xa8,
	0x69, 0x56, 0x34, 0x17, 0xaa, 0x18, 0x6e, 0xc5, 0xba, 0x2e, 0xb6, 0x82, 0xef, 0x0a, 0x2a, 0x99,
	0x56, 0x82, 0x15, 0xbd, 0xa1, 0xcd, 0x6f, 0x8b, 0xe4, 0x6b, 0xad, 0x8c, 0xc2, 0xcf, 0xc6, 0x0c,
	0x07, 0x8d, 0xa2, 0xeb, 0x9b, 0xd3, 0x85, 0x44, 0x47, 0x97, 0x7c, 0xc8, 0x4b, 0xdf, 0xc2, 0xd5,
	0x77, 0x5b, 0xea, 0x6e, 0x25, 0x5a, 0xf6, 0x53, 0xf0, 0x1d, 0xbe, 0x86, 0xb0, 0x6f, 0x34, 0xe7,
	0x72, 0xc6, 0x62, 0x3f, 0x41, 0x99, 0x4f, 0xa6, 0x38, 0x9d, 0xc3, 0xa5, 0x53, 0x3b, 0xe1, 0x07,
	0x08, 0x1b, 0x9b, 0xa5, 0xb9, 0x8c, 0x51, 0xe2, 0x67, 0x51, 0x79, 0x9b, 0x9f, 0x70, 0x91, 0x1f,
	0xd6, 0x27, 0x53, 0x52, 0xfa, 0x0f, 0x41, 0xe8, 0xc8, 0x8a, 0x6a, 0xfc, 0x1c, 0x02, 0x23, 0x4c,
	0xcb, 0x63, 0x94, 0xa0, 0xec, 0x92, 0x0c, 0x81, 0x33, 0xb3, 0xa9, 0x07, 0xc2, 0x77, 0xc4, 0x14,
	0xe3, 0x5b, 0x08, 0x1a, 0xd5, 0x2a, 0x1d, 0x5f, 0x24, 0x28, 0x8b, 0xca, 0xc7, 0xfb, 0xc7, 0xef,
	0x2c, 0x48, 0x06, 0x0e, 0xbf, 0x87, 0x40, 0x18, 0xde, 0xf5, 0x71, 0xe0, 0x1c, 0xde, 0x9c, 0x77,
	0x58, 0x51, 0x3d, 0x33, 0xbc, 0x23, 0x83, 0x1e, 0xbf, 0x86, 0xa7, 0xb5, 0x45, 0x37, 0xc6, 0x28,
	0xf9, 0x45, 0x30, 0xc6, 0x65, 0xec, 0x25, 0x28, 0x0b, 0xc9, 0x11, 0x9e, 0xfe, 0x41, 0xf0, 0xe8,
	0x7e, 0x8d, 0x33, 0xcd, 0xbc, 0x81, 0x0b, 0xd1, 0x28, 0xe9, 0x1a, 0x89, 0xca, 0x17, 0x7b, 0x2b,
	0x33, 0xbb, 0x8e, 0x6f, 0x9a, 0xf0, 0x5e, 0x6d, 0x74, 0xc3, 0x89, 0x13, 0xe1, 0x57, 0x10, 0xda,
	0xf3, 0x87, 0x90, 0xc6, 0xbd, 0x7b, 0xd4, 0xe0, 0x44, 0xdb, 0x21, 0x7d, 0x12, 0x3d, 0xad, 0x5b,
	0xce, 0xdc, 0x2c, 0x42, 0x32, 0xc5, 0xe9, 0x4b, 0x00, 0xe7, 0xec, 0xf3, 0x96, 0x4b, 0x83, 0xaf,
	0xc0, 0x13, 0xcc, 0x2d, 0xcb, 0x27, 0x9e, 0x60, 0xd5, 0x1c, 0x52, 0xa1, 0xf2, 0xe9, 0x97, 0x8c,
	0xc7, 0xba, 0x3e, 0x18, 0x4f, 0x15, 0x2d, 0xea, 0x69, 0xeb, 0xbf, 0x1e, 0x8e, 0xe8, 0x5f, 0xef,
	0xc9, 0x57, 0x27, 0xff, 0x38, 0xc4, 0x8b, 0xaa, 0x7e, 0xe0, 0xbe, 0xd4, 0xbb, 0xff, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x34, 0xf3, 0xc4, 0xf2, 0xc0, 0x02, 0x00, 0x00,
}