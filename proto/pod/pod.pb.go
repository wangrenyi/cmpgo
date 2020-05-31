// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/wangrenyi/kmpgo/proto/pod/pod.proto

/*
Package pod is a generated protocol buffer package.

It is generated from these files:
	github.com/wangrenyi/kmpgo/proto/pod/pod.proto

It has these top-level messages:
*/
package pod

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/wangrenyi/kmpgo/proto/common"

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

// Client API for PodInterface service

type PodInterfaceClient interface {
	PodCreate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PodUpdate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PodUpdateStatus(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PodDelete(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PodDeleteCollection(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PodGet(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PodList(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PodWatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PodPatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PodGetEphemeralContainers(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PodUpdateEphemeralContainers(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	// pod expansion
	PodBind(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PodEvict(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
	PodGetLogs(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error)
}

type podInterfaceClient struct {
	cc *grpc.ClientConn
}

func NewPodInterfaceClient(cc *grpc.ClientConn) PodInterfaceClient {
	return &podInterfaceClient{cc}
}

func (c *podInterfaceClient) PodCreate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podInterfaceClient) PodUpdate(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podInterfaceClient) PodUpdateStatus(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodUpdateStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podInterfaceClient) PodDelete(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podInterfaceClient) PodDeleteCollection(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodDeleteCollection", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podInterfaceClient) PodGet(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podInterfaceClient) PodList(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podInterfaceClient) PodWatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodWatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podInterfaceClient) PodPatch(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodPatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podInterfaceClient) PodGetEphemeralContainers(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodGetEphemeralContainers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podInterfaceClient) PodUpdateEphemeralContainers(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodUpdateEphemeralContainers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podInterfaceClient) PodBind(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodBind", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podInterfaceClient) PodEvict(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodEvict", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podInterfaceClient) PodGetLogs(ctx context.Context, in *common.InvokeServiceRequest, opts ...grpc.CallOption) (*common.InvokeServiceResponse, error) {
	out := new(common.InvokeServiceResponse)
	err := grpc.Invoke(ctx, "/pod.PodInterface/PodGetLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PodInterface service

type PodInterfaceServer interface {
	PodCreate(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PodUpdate(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PodUpdateStatus(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PodDelete(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PodDeleteCollection(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PodGet(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PodList(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PodWatch(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PodPatch(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PodGetEphemeralContainers(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PodUpdateEphemeralContainers(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	// pod expansion
	PodBind(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PodEvict(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
	PodGetLogs(context.Context, *common.InvokeServiceRequest) (*common.InvokeServiceResponse, error)
}

func RegisterPodInterfaceServer(s *grpc.Server, srv PodInterfaceServer) {
	s.RegisterService(&_PodInterface_serviceDesc, srv)
}

func _PodInterface_PodCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodCreate(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodInterface_PodUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodUpdate(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodInterface_PodUpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodUpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodUpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodUpdateStatus(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodInterface_PodDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodDelete(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodInterface_PodDeleteCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodDeleteCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodDeleteCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodDeleteCollection(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodInterface_PodGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodGet(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodInterface_PodList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodList(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodInterface_PodWatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodWatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodWatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodWatch(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodInterface_PodPatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodPatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodPatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodPatch(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodInterface_PodGetEphemeralContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodGetEphemeralContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodGetEphemeralContainers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodGetEphemeralContainers(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodInterface_PodUpdateEphemeralContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodUpdateEphemeralContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodUpdateEphemeralContainers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodUpdateEphemeralContainers(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodInterface_PodBind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodBind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodBind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodBind(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodInterface_PodEvict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodEvict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodEvict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodEvict(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodInterface_PodGetLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodInterfaceServer).PodGetLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pod.PodInterface/PodGetLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodInterfaceServer).PodGetLogs(ctx, req.(*common.InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PodInterface_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pod.PodInterface",
	HandlerType: (*PodInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PodCreate",
			Handler:    _PodInterface_PodCreate_Handler,
		},
		{
			MethodName: "PodUpdate",
			Handler:    _PodInterface_PodUpdate_Handler,
		},
		{
			MethodName: "PodUpdateStatus",
			Handler:    _PodInterface_PodUpdateStatus_Handler,
		},
		{
			MethodName: "PodDelete",
			Handler:    _PodInterface_PodDelete_Handler,
		},
		{
			MethodName: "PodDeleteCollection",
			Handler:    _PodInterface_PodDeleteCollection_Handler,
		},
		{
			MethodName: "PodGet",
			Handler:    _PodInterface_PodGet_Handler,
		},
		{
			MethodName: "PodList",
			Handler:    _PodInterface_PodList_Handler,
		},
		{
			MethodName: "PodWatch",
			Handler:    _PodInterface_PodWatch_Handler,
		},
		{
			MethodName: "PodPatch",
			Handler:    _PodInterface_PodPatch_Handler,
		},
		{
			MethodName: "PodGetEphemeralContainers",
			Handler:    _PodInterface_PodGetEphemeralContainers_Handler,
		},
		{
			MethodName: "PodUpdateEphemeralContainers",
			Handler:    _PodInterface_PodUpdateEphemeralContainers_Handler,
		},
		{
			MethodName: "PodBind",
			Handler:    _PodInterface_PodBind_Handler,
		},
		{
			MethodName: "PodEvict",
			Handler:    _PodInterface_PodEvict_Handler,
		},
		{
			MethodName: "PodGetLogs",
			Handler:    _PodInterface_PodGetLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/wangrenyi/kmpgo/proto/pod/pod.proto",
}

func init() { proto.RegisterFile("github.com/wangrenyi/kmpgo/proto/pod/pod.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0xd4, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x00, 0x50, 0x41, 0x58, 0xd7, 0x41, 0x10, 0xea, 0xc9, 0x65, 0xbd, 0xf8, 0x01, 0x2d, 0xa8,
	0x5f, 0x60, 0x5d, 0xd6, 0xea, 0x1e, 0x82, 0xab, 0x08, 0x82, 0x87, 0x6c, 0x32, 0xb6, 0x61, 0xdb,
	0x4c, 0x4c, 0xa7, 0x15, 0xff, 0xd6, 0x4f, 0x11, 0xb7, 0xa5, 0x37, 0xf1, 0x90, 0x7a, 0x28, 0x25,
	0x24, 0xf3, 0x98, 0xcc, 0x0c, 0x81, 0x38, 0x37, 0x5c, 0x34, 0x9b, 0x58, 0x51, 0x95, 0x7c, 0x48,
	0x9b, 0x7b, 0xb4, 0x9f, 0x26, 0xd9, 0x56, 0x2e, 0xa7, 0xc4, 0x79, 0x62, 0x4a, 0x1c, 0xe9, 0x9f,
	0x2f, 0xde, 0xad, 0xa2, 0x7d, 0x47, 0x7a, 0x76, 0xf5, 0x67, 0x90, 0xa2, 0xaa, 0x22, 0xdb, 0xff,
	0xba, 0xd0, 0x8b, 0xaf, 0x29, 0x1c, 0x09, 0xd2, 0x99, 0x65, 0xf4, 0x6f, 0x52, 0x61, 0x74, 0x07,
	0x87, 0x82, 0x74, 0xea, 0x51, 0x32, 0x46, 0xf3, 0xb8, 0x3f, 0x9c, 0xd9, 0x96, 0xb6, 0xb8, 0x46,
	0xdf, 0x1a, 0x85, 0x0f, 0xf8, 0xde, 0x60, 0xcd, 0xb3, 0xb3, 0x5f, 0x76, 0x6b, 0x47, 0xb6, 0xc6,
	0xf3, 0xbd, 0xde, 0x7a, 0x72, 0x7a, 0x04, 0x4b, 0xc0, 0xf1, 0x60, 0xad, 0x59, 0x72, 0x53, 0x8f,
	0x93, 0xdd, 0x0d, 0x96, 0x18, 0x9e, 0xdd, 0x23, 0x9c, 0x0c, 0x56, 0x4a, 0x65, 0x89, 0x8a, 0x0d,
	0xd9, 0x50, 0x75, 0x09, 0x13, 0x41, 0x7a, 0x89, 0x1c, 0x0a, 0xdd, 0xc2, 0x81, 0x20, 0xbd, 0x32,
	0x75, 0xb0, 0x94, 0xc1, 0x54, 0x90, 0x7e, 0x96, 0xac, 0x8a, 0x71, 0x28, 0x31, 0x06, 0xf5, 0x02,
	0xa7, 0x5d, 0xa1, 0x16, 0xae, 0xc0, 0x0a, 0xbd, 0x2c, 0x53, 0xb2, 0x2c, 0x8d, 0x45, 0x1f, 0x3c,
	0x26, 0xaf, 0x30, 0x1f, 0x06, 0xef, 0x1f, 0xf8, 0xae, 0x35, 0xd7, 0xc6, 0xea, 0x71, 0xea, 0xb9,
	0x68, 0x8d, 0x0a, 0xee, 0xf2, 0x3d, 0x40, 0x57, 0xcf, 0x15, 0xe5, 0xa1, 0x37, 0xdc, 0x4c, 0x76,
	0x2f, 0xcd, 0xe5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0x61, 0x6b, 0x94, 0xd6, 0x04, 0x00,
	0x00,
}
