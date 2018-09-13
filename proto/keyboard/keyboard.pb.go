// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/gomatcha/matcha/proto/keyboard/keyboard.proto

package keyboard

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

type Type int32

const (
	Type_TEXT_TYPE      Type = 0
	Type_NUMBER_TYPE    Type = 1
	Type_EMAIL_TYPE     Type = 2
	Type_URL_TYPE       Type = 3
	Type_PHONE_TYPE     Type = 4
	Type_DATE_TIME_TYPE Type = 5
)

var Type_name = map[int32]string{
	0: "TEXT_TYPE",
	1: "NUMBER_TYPE",
	2: "EMAIL_TYPE",
	3: "URL_TYPE",
	4: "PHONE_TYPE",
	5: "DATE_TIME_TYPE",
}

var Type_value = map[string]int32{
	"TEXT_TYPE":      0,
	"NUMBER_TYPE":    1,
	"EMAIL_TYPE":     2,
	"URL_TYPE":       3,
	"PHONE_TYPE":     4,
	"DATE_TIME_TYPE": 5,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b97cd04f987885c, []int{0}
}

type Appearance int32

const (
	Appearance_DEFAULT_APPEARANCE Appearance = 0
	Appearance_LIGHT_APPEARANCE   Appearance = 1
	Appearance_DARK_APPEARANCE    Appearance = 2
)

var Appearance_name = map[int32]string{
	0: "DEFAULT_APPEARANCE",
	1: "LIGHT_APPEARANCE",
	2: "DARK_APPEARANCE",
}

var Appearance_value = map[string]int32{
	"DEFAULT_APPEARANCE": 0,
	"LIGHT_APPEARANCE":   1,
	"DARK_APPEARANCE":    2,
}

func (x Appearance) String() string {
	return proto.EnumName(Appearance_name, int32(x))
}

func (Appearance) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b97cd04f987885c, []int{1}
}

type ReturnType int32

const (
	ReturnType_DEFAULT_RETURN_TYPE        ReturnType = 0
	ReturnType_GO_RETURN_TYPE             ReturnType = 1
	ReturnType_GOOGLE_RETURN_TYPE         ReturnType = 2
	ReturnType_JOIN_RETURN_TYPE           ReturnType = 3
	ReturnType_NEXT_RETURN_TYPE           ReturnType = 4
	ReturnType_ROUTE_RETURN_TYPE          ReturnType = 5
	ReturnType_SEARCH_RETURN_TYPE         ReturnType = 6
	ReturnType_SEND_RETURN_TYPE           ReturnType = 7
	ReturnType_YAHOO_RETURN_TYPE          ReturnType = 8
	ReturnType_DONE_RETURN_TYPE           ReturnType = 9
	ReturnType_EMERGENCY_CALL_RETURN_TYPE ReturnType = 10
	ReturnType_CONTINUE_RETURN_TYPE       ReturnType = 11
)

var ReturnType_name = map[int32]string{
	0:  "DEFAULT_RETURN_TYPE",
	1:  "GO_RETURN_TYPE",
	2:  "GOOGLE_RETURN_TYPE",
	3:  "JOIN_RETURN_TYPE",
	4:  "NEXT_RETURN_TYPE",
	5:  "ROUTE_RETURN_TYPE",
	6:  "SEARCH_RETURN_TYPE",
	7:  "SEND_RETURN_TYPE",
	8:  "YAHOO_RETURN_TYPE",
	9:  "DONE_RETURN_TYPE",
	10: "EMERGENCY_CALL_RETURN_TYPE",
	11: "CONTINUE_RETURN_TYPE",
}

var ReturnType_value = map[string]int32{
	"DEFAULT_RETURN_TYPE":        0,
	"GO_RETURN_TYPE":             1,
	"GOOGLE_RETURN_TYPE":         2,
	"JOIN_RETURN_TYPE":           3,
	"NEXT_RETURN_TYPE":           4,
	"ROUTE_RETURN_TYPE":          5,
	"SEARCH_RETURN_TYPE":         6,
	"SEND_RETURN_TYPE":           7,
	"YAHOO_RETURN_TYPE":          8,
	"DONE_RETURN_TYPE":           9,
	"EMERGENCY_CALL_RETURN_TYPE": 10,
	"CONTINUE_RETURN_TYPE":       11,
}

func (x ReturnType) String() string {
	return proto.EnumName(ReturnType_name, int32(x))
}

func (ReturnType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b97cd04f987885c, []int{2}
}

func init() {
	proto.RegisterEnum("matcha.keyboard.Type", Type_name, Type_value)
	proto.RegisterEnum("matcha.keyboard.Appearance", Appearance_name, Appearance_value)
	proto.RegisterEnum("matcha.keyboard.ReturnType", ReturnType_name, ReturnType_value)
}

func init() {
	proto.RegisterFile("github.com/gomatcha/matcha/proto/keyboard/keyboard.proto", fileDescriptor_7b97cd04f987885c)
}

var fileDescriptor_7b97cd04f987885c = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x4b, 0x8f, 0xda, 0x30,
	0x10, 0xc7, 0x4b, 0x78, 0x14, 0x86, 0x16, 0xa8, 0xa1, 0x0f, 0xf5, 0xd0, 0x7b, 0x39, 0x84, 0x43,
	0x2f, 0xbd, 0x3a, 0x89, 0x9b, 0xa4, 0x24, 0x76, 0x64, 0x1c, 0x69, 0xd9, 0x4b, 0x94, 0x64, 0x23,
	0x58, 0xad, 0x20, 0x51, 0x14, 0x0e, 0x7c, 0x9d, 0xbd, 0xed, 0xb7, 0x5c, 0x39, 0x0f, 0x84, 0x4f,
	0xf6, 0xfc, 0x66, 0xfc, 0x1f, 0xcf, 0xe8, 0x0f, 0x7f, 0x0f, 0xcf, 0xd5, 0xf1, 0x92, 0xe8, 0x69,
	0x7e, 0xda, 0x1c, 0xf2, 0x53, 0x5c, 0xa5, 0xc7, 0x78, 0xd3, 0x1e, 0x45, 0x99, 0x57, 0xf9, 0xe6,
	0x25, 0xbb, 0x26, 0x79, 0x5c, 0x3e, 0xdd, 0x2e, 0x7a, 0xcd, 0xd1, 0xbc, 0xa9, 0xd2, 0x3b, 0xbc,
	0x3e, 0xc2, 0x40, 0x5c, 0x8b, 0x0c, 0x7d, 0x86, 0x89, 0x20, 0x0f, 0x22, 0x12, 0xfb, 0x80, 0x2c,
	0x3e, 0xa0, 0x39, 0x4c, 0x69, 0xe8, 0x1b, 0x84, 0x37, 0xa0, 0x87, 0x66, 0x00, 0xc4, 0xc7, 0xae,
	0xd7, 0xc4, 0x1a, 0xfa, 0x04, 0xe3, 0x90, 0xb7, 0x51, 0x5f, 0x66, 0x03, 0x87, 0x51, 0xd2, 0xc4,
	0x03, 0x84, 0x60, 0x66, 0x61, 0x41, 0x22, 0xe1, 0xfa, 0x2d, 0x1b, 0xae, 0x19, 0x00, 0x2e, 0x8a,
	0x2c, 0x2e, 0xe3, 0x73, 0x9a, 0xa1, 0x6f, 0x80, 0x2c, 0xf2, 0x0f, 0x87, 0x9e, 0x88, 0x70, 0x10,
	0x10, 0xcc, 0x31, 0x35, 0x65, 0xe3, 0x15, 0x2c, 0x3c, 0xd7, 0x76, 0x14, 0xda, 0x43, 0x4b, 0x98,
	0x5b, 0x98, 0x6f, 0xef, 0xa1, 0xb6, 0x7e, 0xd3, 0x00, 0x78, 0x56, 0x5d, 0xca, 0x73, 0x3d, 0xc1,
	0x77, 0x58, 0x76, 0x8a, 0x9c, 0x88, 0x90, 0xd3, 0x6e, 0x16, 0x04, 0x33, 0x9b, 0x29, 0xac, 0x27,
	0xdb, 0xdb, 0x8c, 0xd9, 0x1e, 0x51, 0xb8, 0x26, 0xdb, 0xff, 0x67, 0x2e, 0x55, 0x68, 0x5f, 0x52,
	0x2a, 0x97, 0x73, 0x4f, 0x07, 0xe8, 0x2b, 0x7c, 0xe1, 0x2c, 0x14, 0xaa, 0xc4, 0x50, 0x4a, 0xef,
	0x08, 0xe6, 0xa6, 0xa3, 0xf0, 0x91, 0x14, 0xd9, 0x11, 0x6a, 0x29, 0xf4, 0xa3, 0x14, 0xd9, 0x63,
	0x87, 0xa9, 0xff, 0x1b, 0xcb, 0x62, 0x4b, 0xee, 0xf3, 0x9e, 0x4e, 0xd0, 0x2f, 0xf8, 0x49, 0x7c,
	0xc2, 0x6d, 0x42, 0xcd, 0x7d, 0x64, 0x62, 0xcf, 0x53, 0xf2, 0x80, 0x7e, 0xc0, 0xca, 0x64, 0x54,
	0xb8, 0x34, 0x54, 0x5f, 0x4e, 0x0d, 0x01, 0xbf, 0xd3, 0xfc, 0xa4, 0xb7, 0xbe, 0xe9, 0x3c, 0xa3,
	0xb7, 0x47, 0xed, 0x8d, 0x9b, 0x27, 0x0c, 0x08, 0x92, 0x6d, 0x7b, 0x7f, 0x1c, 0x77, 0xf4, 0x55,
	0x5b, 0xf8, 0x75, 0x75, 0x97, 0x0a, 0x8c, 0x64, 0x54, 0x3f, 0xfc, 0xf3, 0x1e, 0x00, 0x00, 0xff,
	0xff, 0x26, 0x7a, 0x8c, 0xd6, 0x90, 0x02, 0x00, 0x00,
}
