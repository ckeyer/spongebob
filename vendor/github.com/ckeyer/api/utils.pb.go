// Code generated by protoc-gen-go. DO NOT EDIT.
// source: utils.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type String struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *String) Reset()                    { *m = String{} }
func (m *String) String() string            { return proto.CompactTextString(m) }
func (*String) ProtoMessage()               {}
func (*String) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "ckeyer.api.Empty")
	proto.RegisterType((*String)(nil), "ckeyer.api.String")
}

func init() { proto.RegisterFile("utils.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2d, 0xc9, 0xcc,
	0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0xce, 0x4e, 0xad, 0x4c, 0x2d, 0xd2,
	0x4b, 0x2c, 0xc8, 0x54, 0x62, 0xe7, 0x62, 0x75, 0xcd, 0x2d, 0x28, 0xa9, 0x54, 0x92, 0xe3, 0x62,
	0x0b, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x17, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x9c, 0x58, 0xa3, 0x98, 0x13, 0x0b, 0x32, 0x93,
	0xd8, 0xc0, 0x46, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x52, 0xe5, 0x31, 0xcc, 0x51, 0x00,
	0x00, 0x00,
}