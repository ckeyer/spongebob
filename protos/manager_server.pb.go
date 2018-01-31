// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manager_server.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	manager_server.proto
	task_server.proto
	types.proto

It has these top-level messages:
	HTTPOption
	Node
	Task
	NodeStatus
	NodeInfo
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Controller service

type ControllerClient interface {
	Join(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Task, error)
	ReportStatus(ctx context.Context, opts ...grpc.CallOption) (Controller_ReportStatusClient, error)
}

type controllerClient struct {
	cc *grpc.ClientConn
}

func NewControllerClient(cc *grpc.ClientConn) ControllerClient {
	return &controllerClient{cc}
}

func (c *controllerClient) Join(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/ckeyer.com.spongebob.protos.Controller/Join", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerClient) ReportStatus(ctx context.Context, opts ...grpc.CallOption) (Controller_ReportStatusClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Controller_serviceDesc.Streams[0], c.cc, "/ckeyer.com.spongebob.protos.Controller/ReportStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &controllerReportStatusClient{stream}
	return x, nil
}

type Controller_ReportStatusClient interface {
	Send(*NodeStatus) error
	Recv() (*Task, error)
	grpc.ClientStream
}

type controllerReportStatusClient struct {
	grpc.ClientStream
}

func (x *controllerReportStatusClient) Send(m *NodeStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controllerReportStatusClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Controller service

type ControllerServer interface {
	Join(context.Context, *Node) (*Task, error)
	ReportStatus(Controller_ReportStatusServer) error
}

func RegisterControllerServer(s *grpc.Server, srv ControllerServer) {
	s.RegisterService(&_Controller_serviceDesc, srv)
}

func _Controller_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ckeyer.com.spongebob.protos.Controller/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).Join(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller_ReportStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControllerServer).ReportStatus(&controllerReportStatusServer{stream})
}

type Controller_ReportStatusServer interface {
	Send(*Task) error
	Recv() (*NodeStatus, error)
	grpc.ServerStream
}

type controllerReportStatusServer struct {
	grpc.ServerStream
}

func (x *controllerReportStatusServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controllerReportStatusServer) Recv() (*NodeStatus, error) {
	m := new(NodeStatus)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Controller_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ckeyer.com.spongebob.protos.Controller",
	HandlerType: (*ControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _Controller_Join_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReportStatus",
			Handler:       _Controller_ReportStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "manager_server.proto",
}

func init() { proto.RegisterFile("manager_server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x4d, 0xcc, 0x4b,
	0x4c, 0x4f, 0x2d, 0x8a, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x4e, 0xce, 0x4e, 0xad, 0x4c, 0x2d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2b, 0x2e, 0xc8,
	0xcf, 0x4b, 0x4f, 0x4d, 0xca, 0x4f, 0x82, 0xc8, 0x15, 0x4b, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7,
	0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7,
	0x15, 0x43, 0xa4, 0xa5, 0xb8, 0x4b, 0x2a, 0x0b, 0x52, 0xa1, 0x1c, 0xa3, 0x07, 0x8c, 0x5c, 0x5c,
	0xce, 0xf9, 0x79, 0x25, 0x45, 0xf9, 0x39, 0x39, 0xa9, 0x45, 0x42, 0x71, 0x5c, 0x2c, 0x5e, 0xf9,
	0x99, 0x79, 0x42, 0x8a, 0x7a, 0x78, 0xcc, 0xd7, 0xf3, 0xcb, 0x4f, 0x49, 0x95, 0xc2, 0xaf, 0x24,
	0x24, 0xb1, 0x38, 0x5b, 0x49, 0xa0, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0x5c, 0x4a, 0xac, 0xfa, 0x59,
	0xf9, 0x99, 0x79, 0x56, 0x8c, 0x5a, 0x42, 0x45, 0x5c, 0x3c, 0x41, 0xa9, 0x05, 0xf9, 0x45, 0x25,
	0xc1, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x42, 0xea, 0x04, 0xed, 0x81, 0x28, 0x24, 0xc6, 0x36, 0x21,
	0xb0, 0x6d, 0x3c, 0x4a, 0xec, 0xfa, 0xc5, 0x60, 0x3d, 0x56, 0x8c, 0x5a, 0x1a, 0x8c, 0x06, 0x8c,
	0x4e, 0x1c, 0x51, 0x6c, 0x10, 0x65, 0x49, 0x10, 0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa2,
	0xf0, 0xa8, 0x5f, 0x53, 0x01, 0x00, 0x00,
}