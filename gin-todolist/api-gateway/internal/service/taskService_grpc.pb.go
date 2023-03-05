// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: internal/service/pb/taskService.proto

package service

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

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	TaskCreate(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	TaskUpdate(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	TaskShow(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TasksDetailResponse, error)
	TaskDelete(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) TaskCreate(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pb.TaskService/TaskCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) TaskUpdate(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pb.TaskService/TaskUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) TaskShow(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TasksDetailResponse, error) {
	out := new(TasksDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.TaskService/TaskShow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) TaskDelete(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pb.TaskService/TaskDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	TaskCreate(context.Context, *TaskRequest) (*CommonResponse, error)
	TaskUpdate(context.Context, *TaskRequest) (*CommonResponse, error)
	TaskShow(context.Context, *TaskRequest) (*TasksDetailResponse, error)
	TaskDelete(context.Context, *TaskRequest) (*CommonResponse, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) TaskCreate(context.Context, *TaskRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskCreate not implemented")
}
func (UnimplementedTaskServiceServer) TaskUpdate(context.Context, *TaskRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskUpdate not implemented")
}
func (UnimplementedTaskServiceServer) TaskShow(context.Context, *TaskRequest) (*TasksDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskShow not implemented")
}
func (UnimplementedTaskServiceServer) TaskDelete(context.Context, *TaskRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskDelete not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_TaskCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TaskService/TaskCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskCreate(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_TaskUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TaskService/TaskUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskUpdate(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_TaskShow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskShow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TaskService/TaskShow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskShow(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_TaskDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TaskService/TaskDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskDelete(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskCreate",
			Handler:    _TaskService_TaskCreate_Handler,
		},
		{
			MethodName: "TaskUpdate",
			Handler:    _TaskService_TaskUpdate_Handler,
		},
		{
			MethodName: "TaskShow",
			Handler:    _TaskService_TaskShow_Handler,
		},
		{
			MethodName: "TaskDelete",
			Handler:    _TaskService_TaskDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/service/pb/taskService.proto",
}
