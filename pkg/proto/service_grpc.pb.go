// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: pkg/proto/service.proto

package __

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

// FileManagerClient is the client API for FileManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileManagerClient interface {
	Recover(ctx context.Context, in *InfoForRecover, opts ...grpc.CallOption) (*Resp, error)
	Delete(ctx context.Context, in *InfoForDeletion, opts ...grpc.CallOption) (*Resp, error)
	Rename(ctx context.Context, in *InfoForRenaming, opts ...grpc.CallOption) (*Resp, error)
	Check(ctx context.Context, in *InfoForCheck, opts ...grpc.CallOption) (*Resp, error)
	SaveBatch(ctx context.Context, in *Batch, opts ...grpc.CallOption) (*Resp, error)
	GetBatch(ctx context.Context, in *GetBatchReq, opts ...grpc.CallOption) (*Batch, error)
	CreateFolder(ctx context.Context, in *CreateFolderReq, opts ...grpc.CallOption) (*Resp, error)
}

type fileManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewFileManagerClient(cc grpc.ClientConnInterface) FileManagerClient {
	return &fileManagerClient{cc}
}

func (c *fileManagerClient) Recover(ctx context.Context, in *InfoForRecover, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/service.FileManager/Recover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) Delete(ctx context.Context, in *InfoForDeletion, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/service.FileManager/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) Rename(ctx context.Context, in *InfoForRenaming, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/service.FileManager/Rename", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) Check(ctx context.Context, in *InfoForCheck, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/service.FileManager/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) SaveBatch(ctx context.Context, in *Batch, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/service.FileManager/SaveBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) GetBatch(ctx context.Context, in *GetBatchReq, opts ...grpc.CallOption) (*Batch, error) {
	out := new(Batch)
	err := c.cc.Invoke(ctx, "/service.FileManager/GetBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) CreateFolder(ctx context.Context, in *CreateFolderReq, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/service.FileManager/CreateFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileManagerServer is the server API for FileManager service.
// All implementations must embed UnimplementedFileManagerServer
// for forward compatibility
type FileManagerServer interface {
	Recover(context.Context, *InfoForRecover) (*Resp, error)
	Delete(context.Context, *InfoForDeletion) (*Resp, error)
	Rename(context.Context, *InfoForRenaming) (*Resp, error)
	Check(context.Context, *InfoForCheck) (*Resp, error)
	SaveBatch(context.Context, *Batch) (*Resp, error)
	GetBatch(context.Context, *GetBatchReq) (*Batch, error)
	CreateFolder(context.Context, *CreateFolderReq) (*Resp, error)
	mustEmbedUnimplementedFileManagerServer()
}

// UnimplementedFileManagerServer must be embedded to have forward compatible implementations.
type UnimplementedFileManagerServer struct {
}

func (UnimplementedFileManagerServer) Recover(context.Context, *InfoForRecover) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recover not implemented")
}
func (UnimplementedFileManagerServer) Delete(context.Context, *InfoForDeletion) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFileManagerServer) Rename(context.Context, *InfoForRenaming) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rename not implemented")
}
func (UnimplementedFileManagerServer) Check(context.Context, *InfoForCheck) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedFileManagerServer) SaveBatch(context.Context, *Batch) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBatch not implemented")
}
func (UnimplementedFileManagerServer) GetBatch(context.Context, *GetBatchReq) (*Batch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatch not implemented")
}
func (UnimplementedFileManagerServer) CreateFolder(context.Context, *CreateFolderReq) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFolder not implemented")
}
func (UnimplementedFileManagerServer) mustEmbedUnimplementedFileManagerServer() {}

// UnsafeFileManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileManagerServer will
// result in compilation errors.
type UnsafeFileManagerServer interface {
	mustEmbedUnimplementedFileManagerServer()
}

func RegisterFileManagerServer(s grpc.ServiceRegistrar, srv FileManagerServer) {
	s.RegisterService(&FileManager_ServiceDesc, srv)
}

func _FileManager_Recover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoForRecover)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).Recover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.FileManager/Recover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).Recover(ctx, req.(*InfoForRecover))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoForDeletion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.FileManager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).Delete(ctx, req.(*InfoForDeletion))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoForRenaming)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.FileManager/Rename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).Rename(ctx, req.(*InfoForRenaming))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoForCheck)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.FileManager/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).Check(ctx, req.(*InfoForCheck))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_SaveBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Batch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).SaveBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.FileManager/SaveBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).SaveBatch(ctx, req.(*Batch))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_GetBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).GetBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.FileManager/GetBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).GetBatch(ctx, req.(*GetBatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_CreateFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFolderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).CreateFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.FileManager/CreateFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).CreateFolder(ctx, req.(*CreateFolderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FileManager_ServiceDesc is the grpc.ServiceDesc for FileManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.FileManager",
	HandlerType: (*FileManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recover",
			Handler:    _FileManager_Recover_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FileManager_Delete_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _FileManager_Rename_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _FileManager_Check_Handler,
		},
		{
			MethodName: "SaveBatch",
			Handler:    _FileManager_SaveBatch_Handler,
		},
		{
			MethodName: "GetBatch",
			Handler:    _FileManager_GetBatch_Handler,
		},
		{
			MethodName: "CreateFolder",
			Handler:    _FileManager_CreateFolder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/service.proto",
}
