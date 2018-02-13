// Code generated by protoc-gen-go. DO NOT EDIT.
// source: daemon.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	daemon.proto
	task_server.proto

It has these top-level messages:
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ckeyer_api "github.com/ckeyer/api"
import ckeyer_api1 "github.com/ckeyer/api"

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

// Client API for DaemonController service

type DaemonControllerClient interface {
	GetStatus(ctx context.Context, in *ckeyer_api1.Empty, opts ...grpc.CallOption) (*ckeyer_api.ProcessStatus, error)
	Reload(ctx context.Context, in *ckeyer_api1.Empty, opts ...grpc.CallOption) (*ckeyer_api.ProcessStatus, error)
}

type daemonControllerClient struct {
	cc *grpc.ClientConn
}

func NewDaemonControllerClient(cc *grpc.ClientConn) DaemonControllerClient {
	return &daemonControllerClient{cc}
}

func (c *daemonControllerClient) GetStatus(ctx context.Context, in *ckeyer_api1.Empty, opts ...grpc.CallOption) (*ckeyer_api.ProcessStatus, error) {
	out := new(ckeyer_api.ProcessStatus)
	err := grpc.Invoke(ctx, "/ckeyer.com.spongebob.protos.DaemonController/GetStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonControllerClient) Reload(ctx context.Context, in *ckeyer_api1.Empty, opts ...grpc.CallOption) (*ckeyer_api.ProcessStatus, error) {
	out := new(ckeyer_api.ProcessStatus)
	err := grpc.Invoke(ctx, "/ckeyer.com.spongebob.protos.DaemonController/Reload", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DaemonController service

type DaemonControllerServer interface {
	GetStatus(context.Context, *ckeyer_api1.Empty) (*ckeyer_api.ProcessStatus, error)
	Reload(context.Context, *ckeyer_api1.Empty) (*ckeyer_api.ProcessStatus, error)
}

func RegisterDaemonControllerServer(s *grpc.Server, srv DaemonControllerServer) {
	s.RegisterService(&_DaemonController_serviceDesc, srv)
}

func _DaemonController_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ckeyer_api1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonControllerServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ckeyer.com.spongebob.protos.DaemonController/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonControllerServer).GetStatus(ctx, req.(*ckeyer_api1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonController_Reload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ckeyer_api1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonControllerServer).Reload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ckeyer.com.spongebob.protos.DaemonController/Reload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonControllerServer).Reload(ctx, req.(*ckeyer_api1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DaemonController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ckeyer.com.spongebob.protos.DaemonController",
	HandlerType: (*DaemonControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _DaemonController_GetStatus_Handler,
		},
		{
			MethodName: "Reload",
			Handler:    _DaemonController_Reload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "daemon.proto",
}

func init() { proto.RegisterFile("daemon.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x49, 0x4c, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4e, 0xce, 0x4e, 0xad, 0x4c, 0x2d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2b, 0x2e, 0xc8, 0xcf, 0x4b, 0x4f, 0x4d, 0xca, 0x4f, 0x82, 0xc8,
	0x15, 0x4b, 0x29, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0x81, 0x24, 0xf5, 0x21, 0xea, 0xf4, 0x13,
	0x0b, 0x32, 0xf5, 0x8b, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x21, 0x8a, 0xa4, 0x14, 0xb1, 0xab, 0x29,
	0x2d, 0xc9, 0xcc, 0x81, 0x2a, 0x31, 0xea, 0x64, 0xe4, 0x12, 0x70, 0x01, 0x5b, 0xea, 0x9c, 0x9f,
	0x57, 0x52, 0x94, 0x9f, 0x93, 0x93, 0x5a, 0x24, 0x64, 0xcd, 0xc5, 0xe9, 0x9e, 0x5a, 0x12, 0x0c,
	0x36, 0x4a, 0x48, 0x50, 0x0f, 0xea, 0x8c, 0xc4, 0x82, 0x4c, 0x3d, 0xd7, 0xdc, 0x82, 0x92, 0x4a,
	0x29, 0x49, 0x64, 0xa1, 0x80, 0xa2, 0xfc, 0xe4, 0xd4, 0xe2, 0x62, 0x88, 0x6a, 0x25, 0x06, 0x21,
	0x0b, 0x2e, 0xb6, 0xa0, 0xd4, 0x9c, 0xfc, 0xc4, 0x14, 0x52, 0x75, 0x3a, 0x71, 0x44, 0xb1, 0x41,
	0x3c, 0x97, 0x04, 0xa1, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xc0, 0x45, 0x98, 0x10,
	0x01, 0x00, 0x00,
}