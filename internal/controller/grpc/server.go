package grpc

import (
	"context"
	"net"

	pb "github.com/gobox-preegnees/file_manager/api/contract"
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
