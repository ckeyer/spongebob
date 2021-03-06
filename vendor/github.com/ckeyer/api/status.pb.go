// Code generated by protoc-gen-go. DO NOT EDIT.
// source: status.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProcessStatus struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Pid     int32  `protobuf:"varint,3,opt,name=pid" json:"pid,omitempty"`
	StartAt int64  `protobuf:"varint,4,opt,name=startAt" json:"startAt,omitempty"`
	Status  string `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
	Desc    string `protobuf:"bytes,6,opt,name=desc" json:"desc,omitempty"`
}

func (m *ProcessStatus) Reset()                    { *m = ProcessStatus{} }
func (m *ProcessStatus) String() string            { return proto.CompactTextString(m) }
func (*ProcessStatus) ProtoMessage()               {}
func (*ProcessStatus) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ProcessStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProcessStatus) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ProcessStatus) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *ProcessStatus) GetStartAt() int64 {
	if m != nil {
		return m.StartAt
	}
	return 0
}

func (m *ProcessStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ProcessStatus) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func init() {
	proto.RegisterType((*ProcessStatus)(nil), "ckeyer.api.ProcessStatus")
}

func init() { proto.RegisterFile("status.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x49, 0x2c,
	0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0xce, 0x4e, 0xad, 0x4c, 0x2d,
	0xd2, 0x4b, 0x2c, 0xc8, 0x54, 0x9a, 0xca, 0xc8, 0xc5, 0x1b, 0x50, 0x94, 0x9f, 0x9c, 0x5a, 0x5c,
	0x1c, 0x0c, 0x56, 0x23, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0x04, 0x66, 0x0b, 0x49, 0x70, 0xb1, 0x97, 0xa5, 0x16, 0x15, 0x67, 0xe6, 0xe7, 0x49,
	0x30, 0x81, 0x85, 0x61, 0x5c, 0x21, 0x01, 0x2e, 0xe6, 0x82, 0xcc, 0x14, 0x09, 0x66, 0x05, 0x46,
	0x0d, 0xd6, 0x20, 0x10, 0x13, 0xa4, 0xb6, 0xb8, 0x24, 0xb1, 0xa8, 0xc4, 0xb1, 0x44, 0x82, 0x45,
	0x81, 0x51, 0x83, 0x39, 0x08, 0xc6, 0x15, 0x12, 0xe3, 0x62, 0x83, 0xb8, 0x43, 0x82, 0x15, 0x6c,
	0x08, 0x94, 0x07, 0xb2, 0x31, 0x25, 0xb5, 0x38, 0x59, 0x82, 0x0d, 0x62, 0x23, 0x88, 0xed, 0xc4,
	0x1a, 0xc5, 0x9c, 0x58, 0x90, 0x99, 0xc4, 0x06, 0x76, 0xb1, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x2e, 0x34, 0xb9, 0x77, 0xc1, 0x00, 0x00, 0x00,
}
