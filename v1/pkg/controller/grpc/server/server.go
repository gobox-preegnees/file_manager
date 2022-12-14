package server

import (
	"context"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/gobox-preegnees/file_manager/v1/internal/adapters/storage/disk"
	services "github.com/gobox-preegnees/file_manager/v1/internal/domain/services"
	usecase "github.com/gobox-preegnees/file_manager/v1/internal/domain/usecase/file"
	protobuf "github.com/gobox-preegnees/file_manager/v1/pkg/contract"
)

const TCP = "tcp"
const SOCKET = "localhost:50051"

type IUsecase interface{
	SaveFile(ctx context.Context, file usecase.SaveFileDTO) error
}

type server struct {
	protobuf.UnimplementedFileManagerServer
	// usecase IUsecase
	usecase usecase.FileUsecase
}

func NewServer() server {
	os.Create("test.txt")
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	var stor services.IFileStorage = storage.New(f)
	var fileService usecase.IFileService = services.NewFileService(stor)
	var folderService usecase.IFolderService = services.NewFolderService(stor)
	uc := usecase.New(fileService, folderService)
	return server{
		usecase: uc,
	}
}

func (s server) Run() error {

	listener, err := net.Listen(TCP, SOCKET)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	protobuf.RegisterFileManagerServer(grpcServer, &server{})

	return grpcServer.Serve(listener)
}

func (s server) CreateFolder(ctx context.Context, in *protobuf.CreateFolderReq) (*protobuf.StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFolder not implemented")
}

func (s server) RenameFolder(ctx context.Context, in *protobuf.RenameFolderReq) (*protobuf.StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameFolder not implemented")
}

func (s server) DeleteFolder(ctx context.Context, in *protobuf.DeleteFolderReq) (*protobuf.StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFolder not implemented")
}

func (s server) RecoverFolder(ctx context.Context, in *protobuf.RecoverFolderReq) (*protobuf.StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverFolder not implemented")
}

func (s server) GetDirSchema(ctx context.Context, in *protobuf.GetDirSchemaReq) (*protobuf.GetDirSchemaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDirSchema not implemented")
}

func (s server) GetFolder(ctx context.Context, in *protobuf.GetFolderReq) (*protobuf.GetFolderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFolder not implemented")
}

func (s server) SaveFile(ctx context.Context, in *protobuf.SaveFileReq) (*protobuf.StandardResponse, error) {
	err := s.usecase.SaveFile(ctx, usecase.SaveFileDTO{
		FullFile: usecase.FullFile{
			Identifier: usecase.Identifier{
				Username: usecase.Username(in.Identification.Username),
				FolderID: usecase.FolderID(in.Identification.FolderID),
				ClientID: usecase.ClientID(in.Identification.ClientID),
			},
			File: usecase.File{
				Path: usecase.Path(in.File.Path),
			},
		},
	})
	return &protobuf.StandardResponse{
		Status: 100,
		Description: "all ok",
	}, err
}

func (s server) RenameFile(ctx context.Context, in *protobuf.RenameFileReq) (*protobuf.StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameFile not implemented")
}

func (s server) DeleteFile(ctx context.Context, in *protobuf.DeleteFileReq) (*protobuf.StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}

func (s server) RecoverFile(ctx context.Context, in *protobuf.RecoverFileReq) (*protobuf.StandardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverFile not implemented")
}

func (s server) GetFile(ctx context.Context, in *protobuf.GetFileReq) (*protobuf.GetFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
