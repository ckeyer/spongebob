// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manager_server.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	manager_server.proto
	task_server.proto

It has these top-level messages:
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import ckeyer_api "github.com/ckeyer/api/types"

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
	Join(ctx context.Context, in *ckeyer_api.Node, opts ...grpc.CallOption) (*ckeyer_api.Task, error)
	ReportStatus(ctx context.Context, opts ...grpc.CallOption) (Controller_ReportStatusClient, error)
}

type controllerClient struct {
	cc *grpc.ClientConn
}

func NewControllerClient(cc *grpc.ClientConn) ControllerClient {
	return &controllerClient{cc}
}

func (c *controllerClient) Join(ctx context.Context, in *ckeyer_api.Node, opts ...grpc.CallOption) (*ckeyer_api.Task, error) {
	out := new(ckeyer_api.Task)
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
	Send(*ckeyer_api.NodeStatus) error
	Recv() (*ckeyer_api.Task, error)
	grpc.ClientStream
}

type controllerReportStatusClient struct {
	grpc.ClientStream
}

func (x *controllerReportStatusClient) Send(m *ckeyer_api.NodeStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controllerReportStatusClient) Recv() (*ckeyer_api.Task, error) {
	m := new(ckeyer_api.Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Controller service

type ControllerServer interface {
	Join(context.Context, *ckeyer_api.Node) (*ckeyer_api.Task, error)
	ReportStatus(Controller_ReportStatusServer) error
}

func RegisterControllerServer(s *grpc.Server, srv ControllerServer) {
	s.RegisterService(&_Controller_serviceDesc, srv)
}

func _Controller_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ckeyer_api.Node)
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
		return srv.(ControllerServer).Join(ctx, req.(*ckeyer_api.Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller_ReportStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControllerServer).ReportStatus(&controllerReportStatusServer{stream})
}

type Controller_ReportStatusServer interface {
	Send(*ckeyer_api.Task) error
	Recv() (*ckeyer_api.NodeStatus, error)
	grpc.ServerStream
}

type controllerReportStatusServer struct {
	grpc.ServerStream
}

func (x *controllerReportStatusServer) Send(m *ckeyer_api.Task) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controllerReportStatusServer) Recv() (*ckeyer_api.NodeStatus, error) {
	m := new(ckeyer_api.NodeStatus)
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
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xbd, 0x4a, 0x44, 0x31,
	0x10, 0x85, 0x89, 0xf8, 0x47, 0xd8, 0x62, 0x09, 0x62, 0x71, 0xb5, 0xba, 0x85, 0xc8, 0x16, 0x89,
	0x68, 0xb7, 0x58, 0x69, 0x67, 0x21, 0xa2, 0x56, 0x36, 0x92, 0xbb, 0x0e, 0xd9, 0xb8, 0x77, 0x73,
	0x42, 0x66, 0x56, 0xd8, 0xd6, 0x57, 0xb0, 0xf5, 0xad, 0x7c, 0x05, 0x1f, 0x44, 0xbc, 0xb1, 0x51,
	0xac, 0x06, 0xce, 0xf9, 0xe6, 0x1b, 0x46, 0xef, 0x2d, 0x7d, 0xf2, 0x81, 0xca, 0x23, 0x53, 0x79,
	0xa1, 0x62, 0x73, 0x81, 0xc0, 0x1c, 0xcc, 0x16, 0xb4, 0xa6, 0x62, 0x67, 0x58, 0x5a, 0xce, 0x48,
	0x81, 0x3a, 0x74, 0xb5, 0xe3, 0xe6, 0x30, 0x00, 0xa1, 0x27, 0xe7, 0x73, 0x74, 0x3e, 0x25, 0x88,
	0x97, 0x88, 0xc4, 0xb5, 0x6e, 0x8e, 0x42, 0x94, 0xf9, 0xaa, 0xfb, 0x5e, 0x75, 0xd5, 0x32, 0x80,
	0xb2, 0xce, 0xc4, 0x6e, 0x0e, 0x96, 0xca, 0x9d, 0xbe, 0x2b, 0xad, 0x2f, 0x91, 0xa4, 0xa0, 0xef,
	0xa9, 0x98, 0x73, 0xbd, 0x79, 0x85, 0x98, 0xcc, 0xd8, 0xfe, 0x9c, 0xf6, 0x39, 0xda, 0x6b, 0x3c,
	0x51, 0xf3, 0x2b, 0xb9, 0xf7, 0xbc, 0x68, 0xc7, 0xaf, 0x1f, 0x9f, 0x6f, 0x1b, 0xba, 0xdd, 0x72,
	0xcf, 0x88, 0x69, 0xaa, 0x26, 0xe6, 0x46, 0x8f, 0x6e, 0x29, 0xa3, 0xc8, 0x9d, 0x78, 0x59, 0xb1,
	0xd9, 0xff, 0x6b, 0xa9, 0xf9, 0x3f, 0x2e, 0x33, 0xb8, 0x46, 0xed, 0x8e, 0xe3, 0x01, 0x99, 0xaa,
	0xc9, 0xb1, 0x3a, 0x51, 0x17, 0xbb, 0x0f, 0xdb, 0xf5, 0xdd, 0xae, 0xce, 0xb3, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x0f, 0x79, 0xec, 0xb7, 0x2a, 0x01, 0x00, 0x00,
}
