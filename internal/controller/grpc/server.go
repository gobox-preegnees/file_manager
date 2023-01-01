package grpc

import (
	"context"
	"net"

	pb "github.com/gobox-preegnees/file_manager/api/grpc"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type IFileUsecase interface {
	GetFiles(ctx context.Context, identifier entity.Identifier) (files []entity.File, err error)
	SaveFile(ctx context.Context, identifier entity.Identifier, file entity.File) (err error)
	RenameFile(ctx context.Context, identifier entity.Identifier, oldPath, newPath, hash string) (err error)
	DeleteFile(ctx context.Context, identifier entity.Identifier, path, hash string) (err error)
}

type server struct {
	pb.UnimplementedFileManagerServer

	fileUsecase IFileUsecase
	socket      string
	log         *logrus.Logger
}

type GrpcServerConf struct {
	Socket      string
	FileUsecase IFileUsecase
}

func New(conf GrpcServerConf) server {

	return server{
		fileUsecase: conf.FileUsecase,
		socket:      conf.Socket,
	}
}

func (s *server) Serve() error {

	listener, err := net.Listen("tcp", s.socket)
	if err != nil {
		return err
	}

	baseServer := grpc.NewServer(grpc.EmptyServerOption{})
	pb.RegisterFileManagerServer(baseServer, &server{})
	return baseServer.Serve(listener)
}

const CODE_OK = 100

func (s *server) DeleteFile(ctx context.Context, req *pb.DeleteFileReq) (*pb.StandardResponse, error) {

	err := s.fileUsecase.DeleteFile(
		ctx,
		entity.Identifier{
			Username: req.Identifier.Username,
			FolderID: req.Identifier.FolderID,
			ClientID: req.Identifier.ClientID,
		},
		req.Path,
		req.Hash,
	)
	if err != nil {
		return nil, err
	}
	return &pb.StandardResponse{
		Status: CODE_OK,
	}, nil
}

func (s *server) GetFiles(ctx context.Context, req *pb.GetFilesReq) (*pb.GetFilesResp, error) {

	files, err := s.fileUsecase.GetFiles(
		ctx,
		entity.Identifier{
			Username: req.Identifier.Username,
			FolderID: req.Identifier.FolderID,
			ClientID: req.Identifier.ClientID,
		},
	)
	if err != nil {
		return nil, err
	}

	pbFiles := make([]*pb.File, len(files))
	for _, file := range files {
		pbFiles = append(pbFiles, &pb.File{
			Path:        file.FileName,
			Hash:        file.HashSum,
			VirtualName: file.VirtualName,
			ModTime:     file.ModTime,
			Size:        file.SizeFile,
		})
	}

	return &pb.GetFilesResp{
		StandardResponse: &pb.StandardResponse{
			Status: CODE_OK,
		},
		File: pbFiles,
	}, nil
}

func (s *server) RenameFile(ctx context.Context, req *pb.RenameFileReq) (*pb.StandardResponse, error) {

	err := s.fileUsecase.RenameFile(ctx,
		entity.Identifier{
			Username: req.Identifier.Username,
			FolderID: req.Identifier.FolderID,
			ClientID: req.Identifier.ClientID,
		},
		req.RenameInfo.OldPath,
		req.RenameInfo.NewPath,
		req.RenameInfo.Hash,
	)
	if err != nil {
		return nil, err
	}

	return &pb.StandardResponse{
		Status: CODE_OK,
	}, nil
}

func (s *server) SaveFile(ctx context.Context, req *pb.SaveFileReq) (*pb.StandardResponse, error) {

	err := s.fileUsecase.SaveFile(
		ctx,
		entity.Identifier{
			Username: req.Identifier.Username,
			FolderID: req.Identifier.FolderID,
			ClientID: req.Identifier.ClientID,
		},
		entity.File{
			FileName:    req.File.Path,
			HashSum:     req.File.Hash,
			SizeFile:    req.File.Size,
			ModTime:     req.File.ModTime,
			VirtualName: req.File.VirtualName,
		},
	)
	if err != nil {
		return nil, err
	}

	return &pb.StandardResponse{
		Status: CODE_OK,
	}, nil
}
