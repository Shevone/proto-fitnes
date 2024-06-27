// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: sub-notification/sub-notification.proto

package subNotification

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SubscriptionNotificationClient is the client API for SubscriptionNotification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionNotificationClient interface {
	TrainUpdated(ctx context.Context, in *RequestTrain, opts ...grpc.CallOption) (*Response, error)
	TrainRemoved(ctx context.Context, in *RequestTrain, opts ...grpc.CallOption) (*Response, error)
	TrainCreatedByTrainer(ctx context.Context, in *RequestTrainer, opts ...grpc.CallOption) (*Response, error)
}

type subscriptionNotificationClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionNotificationClient(cc grpc.ClientConnInterface) SubscriptionNotificationClient {
	return &subscriptionNotificationClient{cc}
}

func (c *subscriptionNotificationClient) TrainUpdated(ctx context.Context, in *RequestTrain, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/subNotification.SubscriptionNotification/TrainUpdated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionNotificationClient) TrainRemoved(ctx context.Context, in *RequestTrain, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/subNotification.SubscriptionNotification/TrainRemoved", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionNotificationClient) TrainCreatedByTrainer(ctx context.Context, in *RequestTrainer, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/subNotification.SubscriptionNotification/TrainCreatedByTrainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionNotificationServer is the server API for SubscriptionNotification service.
// All implementations must embed UnimplementedSubscriptionNotificationServer
// for forward compatibility
type SubscriptionNotificationServer interface {
	TrainUpdated(context.Context, *RequestTrain) (*Response, error)
	TrainRemoved(context.Context, *RequestTrain) (*Response, error)
	TrainCreatedByTrainer(context.Context, *RequestTrainer) (*Response, error)
	mustEmbedUnimplementedSubscriptionNotificationServer()
}

// UnimplementedSubscriptionNotificationServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriptionNotificationServer struct {
}

func (UnimplementedSubscriptionNotificationServer) TrainUpdated(context.Context, *RequestTrain) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainUpdated not implemented")
}
func (UnimplementedSubscriptionNotificationServer) TrainRemoved(context.Context, *RequestTrain) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainRemoved not implemented")
}
func (UnimplementedSubscriptionNotificationServer) TrainCreatedByTrainer(context.Context, *RequestTrainer) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrainCreatedByTrainer not implemented")
}
func (UnimplementedSubscriptionNotificationServer) mustEmbedUnimplementedSubscriptionNotificationServer() {
}

// UnsafeSubscriptionNotificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionNotificationServer will
// result in compilation errors.
type UnsafeSubscriptionNotificationServer interface {
	mustEmbedUnimplementedSubscriptionNotificationServer()
}

func RegisterSubscriptionNotificationServer(s grpc.ServiceRegistrar, srv SubscriptionNotificationServer) {
	s.RegisterService(&SubscriptionNotification_ServiceDesc, srv)
}

func _SubscriptionNotification_TrainUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTrain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionNotificationServer).TrainUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subNotification.SubscriptionNotification/TrainUpdated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionNotificationServer).TrainUpdated(ctx, req.(*RequestTrain))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionNotification_TrainRemoved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTrain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionNotificationServer).TrainRemoved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subNotification.SubscriptionNotification/TrainRemoved",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionNotificationServer).TrainRemoved(ctx, req.(*RequestTrain))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionNotification_TrainCreatedByTrainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTrainer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionNotificationServer).TrainCreatedByTrainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subNotification.SubscriptionNotification/TrainCreatedByTrainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionNotificationServer).TrainCreatedByTrainer(ctx, req.(*RequestTrainer))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionNotification_ServiceDesc is the grpc.ServiceDesc for SubscriptionNotification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionNotification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subNotification.SubscriptionNotification",
	HandlerType: (*SubscriptionNotificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TrainUpdated",
			Handler:    _SubscriptionNotification_TrainUpdated_Handler,
		},
		{
			MethodName: "TrainRemoved",
			Handler:    _SubscriptionNotification_TrainRemoved_Handler,
		},
		{
			MethodName: "TrainCreatedByTrainer",
			Handler:    _SubscriptionNotification_TrainCreatedByTrainer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sub-notification/sub-notification.proto",
}