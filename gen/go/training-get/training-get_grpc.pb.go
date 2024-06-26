// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: training-get/training-get.proto

package trainingGet

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

// TrainingGetServiceClient is the client API for TrainingGetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainingGetServiceClient interface {
	GetTraining(ctx context.Context, in *GetTrainingRequest, opts ...grpc.CallOption) (*Training, error)
	GetTrainingByWeekDay(ctx context.Context, in *GetTrainingsByWeekDay, opts ...grpc.CallOption) (*GetResponse, error)
	GetTrainingByTrainerAndWeekDay(ctx context.Context, in *GetByTrainer, opts ...grpc.CallOption) (*GetResponse, error)
	GetTrainingByUserAndWeekDay(ctx context.Context, in *GetByUser, opts ...grpc.CallOption) (*GetResponse, error)
}

type trainingGetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainingGetServiceClient(cc grpc.ClientConnInterface) TrainingGetServiceClient {
	return &trainingGetServiceClient{cc}
}

func (c *trainingGetServiceClient) GetTraining(ctx context.Context, in *GetTrainingRequest, opts ...grpc.CallOption) (*Training, error) {
	out := new(Training)
	err := c.cc.Invoke(ctx, "/trainingGet.TrainingGetService/GetTraining", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingGetServiceClient) GetTrainingByWeekDay(ctx context.Context, in *GetTrainingsByWeekDay, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/trainingGet.TrainingGetService/GetTrainingByWeekDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingGetServiceClient) GetTrainingByTrainerAndWeekDay(ctx context.Context, in *GetByTrainer, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/trainingGet.TrainingGetService/GetTrainingByTrainerAndWeekDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainingGetServiceClient) GetTrainingByUserAndWeekDay(ctx context.Context, in *GetByUser, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/trainingGet.TrainingGetService/GetTrainingByUserAndWeekDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainingGetServiceServer is the server API for TrainingGetService service.
// All implementations must embed UnimplementedTrainingGetServiceServer
// for forward compatibility
type TrainingGetServiceServer interface {
	GetTraining(context.Context, *GetTrainingRequest) (*Training, error)
	GetTrainingByWeekDay(context.Context, *GetTrainingsByWeekDay) (*GetResponse, error)
	GetTrainingByTrainerAndWeekDay(context.Context, *GetByTrainer) (*GetResponse, error)
	GetTrainingByUserAndWeekDay(context.Context, *GetByUser) (*GetResponse, error)
	mustEmbedUnimplementedTrainingGetServiceServer()
}

// UnimplementedTrainingGetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrainingGetServiceServer struct {
}

func (UnimplementedTrainingGetServiceServer) GetTraining(context.Context, *GetTrainingRequest) (*Training, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTraining not implemented")
}
func (UnimplementedTrainingGetServiceServer) GetTrainingByWeekDay(context.Context, *GetTrainingsByWeekDay) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrainingByWeekDay not implemented")
}
func (UnimplementedTrainingGetServiceServer) GetTrainingByTrainerAndWeekDay(context.Context, *GetByTrainer) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrainingByTrainerAndWeekDay not implemented")
}
func (UnimplementedTrainingGetServiceServer) GetTrainingByUserAndWeekDay(context.Context, *GetByUser) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrainingByUserAndWeekDay not implemented")
}
func (UnimplementedTrainingGetServiceServer) mustEmbedUnimplementedTrainingGetServiceServer() {}

// UnsafeTrainingGetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainingGetServiceServer will
// result in compilation errors.
type UnsafeTrainingGetServiceServer interface {
	mustEmbedUnimplementedTrainingGetServiceServer()
}

func RegisterTrainingGetServiceServer(s grpc.ServiceRegistrar, srv TrainingGetServiceServer) {
	s.RegisterService(&TrainingGetService_ServiceDesc, srv)
}

func _TrainingGetService_GetTraining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrainingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingGetServiceServer).GetTraining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trainingGet.TrainingGetService/GetTraining",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingGetServiceServer).GetTraining(ctx, req.(*GetTrainingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingGetService_GetTrainingByWeekDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrainingsByWeekDay)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingGetServiceServer).GetTrainingByWeekDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trainingGet.TrainingGetService/GetTrainingByWeekDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingGetServiceServer).GetTrainingByWeekDay(ctx, req.(*GetTrainingsByWeekDay))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingGetService_GetTrainingByTrainerAndWeekDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTrainer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingGetServiceServer).GetTrainingByTrainerAndWeekDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trainingGet.TrainingGetService/GetTrainingByTrainerAndWeekDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingGetServiceServer).GetTrainingByTrainerAndWeekDay(ctx, req.(*GetByTrainer))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainingGetService_GetTrainingByUserAndWeekDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainingGetServiceServer).GetTrainingByUserAndWeekDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trainingGet.TrainingGetService/GetTrainingByUserAndWeekDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainingGetServiceServer).GetTrainingByUserAndWeekDay(ctx, req.(*GetByUser))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainingGetService_ServiceDesc is the grpc.ServiceDesc for TrainingGetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainingGetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trainingGet.TrainingGetService",
	HandlerType: (*TrainingGetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTraining",
			Handler:    _TrainingGetService_GetTraining_Handler,
		},
		{
			MethodName: "GetTrainingByWeekDay",
			Handler:    _TrainingGetService_GetTrainingByWeekDay_Handler,
		},
		{
			MethodName: "GetTrainingByTrainerAndWeekDay",
			Handler:    _TrainingGetService_GetTrainingByTrainerAndWeekDay_Handler,
		},
		{
			MethodName: "GetTrainingByUserAndWeekDay",
			Handler:    _TrainingGetService_GetTrainingByUserAndWeekDay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "training-get/training-get.proto",
}
