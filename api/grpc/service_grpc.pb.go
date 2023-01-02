// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: api/grpc/service.proto

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
	GetFiles(ctx context.Context, in *GetFilesReq, opts ...grpc.CallOption) (*GetFilesResp, error)
	SaveFile(ctx context.Context, in *SaveFileReq, opts ...grpc.CallOption) (*SaveFileResp, error)
	DeleteFile(ctx context.Context, in *DeleteFileReq, opts ...grpc.CallOption) (*DeleteFileResp, error)
	RenameFile(ctx context.Context, in *RenameFileReq, opts ...grpc.CallOption) (*RenameFileResp, error)
}

type fileManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewFileManagerClient(cc grpc.ClientConnInterface) FileManagerClient {
	return &fileManagerClient{cc}
}

func (c *fileManagerClient) GetFiles(ctx context.Context, in *GetFilesReq, opts ...grpc.CallOption) (*GetFilesResp, error) {
	out := new(GetFilesResp)
	err := c.cc.Invoke(ctx, "/contract.FileManager/GetFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) SaveFile(ctx context.Context, in *SaveFileReq, opts ...grpc.CallOption) (*SaveFileResp, error) {
	out := new(SaveFileResp)
	err := c.cc.Invoke(ctx, "/contract.FileManager/SaveFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) DeleteFile(ctx context.Context, in *DeleteFileReq, opts ...grpc.CallOption) (*DeleteFileResp, error) {
	out := new(DeleteFileResp)
	err := c.cc.Invoke(ctx, "/contract.FileManager/DeleteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) RenameFile(ctx context.Context, in *RenameFileReq, opts ...grpc.CallOption) (*RenameFileResp, error) {
	out := new(RenameFileResp)
	err := c.cc.Invoke(ctx, "/contract.FileManager/RenameFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileManagerServer is the server API for FileManager service.
// All implementations must embed UnimplementedFileManagerServer
// for forward compatibility
type FileManagerServer interface {
	GetFiles(context.Context, *GetFilesReq) (*GetFilesResp, error)
	SaveFile(context.Context, *SaveFileReq) (*SaveFileResp, error)
	DeleteFile(context.Context, *DeleteFileReq) (*DeleteFileResp, error)
	RenameFile(context.Context, *RenameFileReq) (*RenameFileResp, error)
	mustEmbedUnimplementedFileManagerServer()
}

// UnimplementedFileManagerServer must be embedded to have forward compatible implementations.
type UnimplementedFileManagerServer struct {
}

func (UnimplementedFileManagerServer) GetFiles(context.Context, *GetFilesReq) (*GetFilesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiles not implemented")
}
func (UnimplementedFileManagerServer) SaveFile(context.Context, *SaveFileReq) (*SaveFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveFile not implemented")
}
func (UnimplementedFileManagerServer) DeleteFile(context.Context, *DeleteFileReq) (*DeleteFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedFileManagerServer) RenameFile(context.Context, *RenameFileReq) (*RenameFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameFile not implemented")
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

func _FileManager_GetFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).GetFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contract.FileManager/GetFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).GetFiles(ctx, req.(*GetFilesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_SaveFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).SaveFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contract.FileManager/SaveFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).SaveFile(ctx, req.(*SaveFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contract.FileManager/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).DeleteFile(ctx, req.(*DeleteFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_RenameFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).RenameFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contract.FileManager/RenameFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).RenameFile(ctx, req.(*RenameFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FileManager_ServiceDesc is the grpc.ServiceDesc for FileManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contract.FileManager",
	HandlerType: (*FileManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFiles",
			Handler:    _FileManager_GetFiles_Handler,
		},
		{
			MethodName: "SaveFile",
			Handler:    _FileManager_SaveFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _FileManager_DeleteFile_Handler,
		},
		{
			MethodName: "RenameFile",
			Handler:    _FileManager_RenameFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/service.proto",
}

// OwnerManagerClient is the client API for OwnerManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OwnerManagerClient interface {
	CreateOwner(ctx context.Context, in *CreateOwnerReq, opts ...grpc.CallOption) (*CreateOwnerResp, error)
	DeleteOwner(ctx context.Context, in *DeleteOwnerReq, opts ...grpc.CallOption) (*DeleteOwnerResp, error)
}

type ownerManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewOwnerManagerClient(cc grpc.ClientConnInterface) OwnerManagerClient {
	return &ownerManagerClient{cc}
}

func (c *ownerManagerClient) CreateOwner(ctx context.Context, in *CreateOwnerReq, opts ...grpc.CallOption) (*CreateOwnerResp, error) {
	out := new(CreateOwnerResp)
	err := c.cc.Invoke(ctx, "/contract.OwnerManager/CreateOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownerManagerClient) DeleteOwner(ctx context.Context, in *DeleteOwnerReq, opts ...grpc.CallOption) (*DeleteOwnerResp, error) {
	out := new(DeleteOwnerResp)
	err := c.cc.Invoke(ctx, "/contract.OwnerManager/DeleteOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerManagerServer is the server API for OwnerManager service.
// All implementations must embed UnimplementedOwnerManagerServer
// for forward compatibility
type OwnerManagerServer interface {
	CreateOwner(context.Context, *CreateOwnerReq) (*CreateOwnerResp, error)
	DeleteOwner(context.Context, *DeleteOwnerReq) (*DeleteOwnerResp, error)
	mustEmbedUnimplementedOwnerManagerServer()
}

// UnimplementedOwnerManagerServer must be embedded to have forward compatible implementations.
type UnimplementedOwnerManagerServer struct {
}

func (UnimplementedOwnerManagerServer) CreateOwner(context.Context, *CreateOwnerReq) (*CreateOwnerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOwner not implemented")
}
func (UnimplementedOwnerManagerServer) DeleteOwner(context.Context, *DeleteOwnerReq) (*DeleteOwnerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOwner not implemented")
}
func (UnimplementedOwnerManagerServer) mustEmbedUnimplementedOwnerManagerServer() {}

// UnsafeOwnerManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OwnerManagerServer will
// result in compilation errors.
type UnsafeOwnerManagerServer interface {
	mustEmbedUnimplementedOwnerManagerServer()
}

func RegisterOwnerManagerServer(s grpc.ServiceRegistrar, srv OwnerManagerServer) {
	s.RegisterService(&OwnerManager_ServiceDesc, srv)
}

func _OwnerManager_CreateOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOwnerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerManagerServer).CreateOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contract.OwnerManager/CreateOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerManagerServer).CreateOwner(ctx, req.(*CreateOwnerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnerManager_DeleteOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOwnerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerManagerServer).DeleteOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contract.OwnerManager/DeleteOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerManagerServer).DeleteOwner(ctx, req.(*DeleteOwnerReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OwnerManager_ServiceDesc is the grpc.ServiceDesc for OwnerManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OwnerManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contract.OwnerManager",
	HandlerType: (*OwnerManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOwner",
			Handler:    _OwnerManager_CreateOwner_Handler,
		},
		{
			MethodName: "DeleteOwner",
			Handler:    _OwnerManager_DeleteOwner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/service.proto",
}
