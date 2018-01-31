// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/ckeyer/spongebob/protos/types.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Node struct {
	Info *HostInfo `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Node) GetInfo() *HostInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type Task struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NodeStatus struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *NodeStatus) Reset()                    { *m = NodeStatus{} }
func (m *NodeStatus) String() string            { return proto.CompactTextString(m) }
func (*NodeStatus) ProtoMessage()               {}
func (*NodeStatus) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *NodeStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HostInfo struct {
	Hostname   string   `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Ips        []string `protobuf:"bytes,2,rep,name=ips" json:"ips,omitempty"`
	Memory     int64    `protobuf:"varint,3,opt,name=memory" json:"memory,omitempty"`
	NumCpus    int64    `protobuf:"varint,4,opt,name=numCpus" json:"numCpus,omitempty"`
	NumPhyCpus int64    `protobuf:"varint,5,opt,name=numPhyCpus" json:"numPhyCpus,omitempty"`
	UnixTime   int64    `protobuf:"varint,6,opt,name=unixTime" json:"unixTime,omitempty"`
}

func (m *HostInfo) Reset()                    { *m = HostInfo{} }
func (m *HostInfo) String() string            { return proto.CompactTextString(m) }
func (*HostInfo) ProtoMessage()               {}
func (*HostInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *HostInfo) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *HostInfo) GetIps() []string {
	if m != nil {
		return m.Ips
	}
	return nil
}

func (m *HostInfo) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *HostInfo) GetNumCpus() int64 {
	if m != nil {
		return m.NumCpus
	}
	return 0
}

func (m *HostInfo) GetNumPhyCpus() int64 {
	if m != nil {
		return m.NumPhyCpus
	}
	return 0
}

func (m *HostInfo) GetUnixTime() int64 {
	if m != nil {
		return m.UnixTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Node)(nil), "ckeyer.com.spongebob.protos.Node")
	proto.RegisterType((*Task)(nil), "ckeyer.com.spongebob.protos.Task")
	proto.RegisterType((*NodeStatus)(nil), "ckeyer.com.spongebob.protos.NodeStatus")
	proto.RegisterType((*HostInfo)(nil), "ckeyer.com.spongebob.protos.HostInfo")
}

func init() { proto.RegisterFile("github.com/ckeyer/spongebob/protos/types.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x41, 0x4b, 0xf4, 0x30,
	0x10, 0x86, 0xe9, 0xd7, 0x7e, 0xeb, 0xee, 0x78, 0x91, 0x1c, 0x24, 0x54, 0x91, 0x52, 0x10, 0x7a,
	0x4a, 0x40, 0x4f, 0x1e, 0xd5, 0x8b, 0x5e, 0x44, 0xea, 0x9e, 0xbc, 0xa5, 0x6b, 0xb6, 0x0d, 0x6b,
	0x32, 0xa1, 0x49, 0xc0, 0xfe, 0x21, 0x7f, 0xa7, 0x34, 0xdd, 0x2d, 0x7b, 0x10, 0x4f, 0x99, 0x37,
	0xcf, 0x33, 0x6f, 0x08, 0xb0, 0x56, 0xf9, 0x2e, 0x34, 0x6c, 0x83, 0x9a, 0x6f, 0x76, 0x72, 0x90,
	0x3d, 0x77, 0x16, 0x4d, 0x2b, 0x1b, 0x6c, 0xb8, 0xed, 0xd1, 0xa3, 0xe3, 0x7e, 0xb0, 0xd2, 0xb1,
	0x18, 0xc8, 0xc5, 0x24, 0x8d, 0x3e, 0x9b, 0xc5, 0x89, 0xb9, 0xfc, 0xb2, 0x45, 0x6c, 0x3f, 0x25,
	0x17, 0x56, 0x71, 0x61, 0x0c, 0x7a, 0xe1, 0x15, 0x9a, 0xfd, 0x6a, 0x79, 0x0f, 0xd9, 0x0b, 0x7e,
	0x48, 0x72, 0x07, 0x99, 0x32, 0x5b, 0xa4, 0x49, 0x91, 0x54, 0xa7, 0x37, 0xd7, 0xec, 0x8f, 0x46,
	0xf6, 0x84, 0xce, 0x3f, 0x9b, 0x2d, 0xd6, 0x71, 0xa5, 0xcc, 0x21, 0x5b, 0x0b, 0xb7, 0x23, 0x04,
	0x32, 0x23, 0xb4, 0x8c, 0x15, 0xab, 0x3a, 0xce, 0x65, 0x01, 0x30, 0xd6, 0xbf, 0x79, 0xe1, 0x83,
	0xfb, 0xd5, 0xf8, 0x4e, 0x60, 0x79, 0x28, 0x24, 0x39, 0x2c, 0x3b, 0x74, 0xfe, 0x48, 0x9a, 0x33,
	0x39, 0x83, 0x54, 0x59, 0x47, 0xff, 0x15, 0x69, 0xb5, 0xaa, 0xc7, 0x91, 0x9c, 0xc3, 0x42, 0x4b,
	0x8d, 0xfd, 0x40, 0xd3, 0x22, 0xa9, 0xd2, 0x7a, 0x9f, 0x08, 0x85, 0x13, 0x13, 0xf4, 0xa3, 0x0d,
	0x8e, 0x66, 0x11, 0x1c, 0x22, 0xb9, 0x02, 0x30, 0x41, 0xbf, 0x76, 0x43, 0x84, 0xff, 0x23, 0x3c,
	0xba, 0x19, 0xdf, 0x0f, 0x46, 0x7d, 0xad, 0x95, 0x96, 0x74, 0x11, 0xe9, 0x9c, 0x1f, 0x96, 0xef,
	0x8b, 0xe9, 0xff, 0xcd, 0x74, 0xde, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xf5, 0x38, 0x82,
	0xa7, 0x01, 0x00, 0x00,
}
