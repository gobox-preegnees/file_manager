package grpc

import (
	"context"
	"net"

	pb "github.com/gobox-preegnees/file_manager/api/grpc"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"
	daoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/dao"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// IFileUsecase.
//go:generate mockgen -destination=../../mocks/grpc/server/file/usecase.go -package=grpc -source=server.go
type IFileUsecase interface {
	GetFiles(ctx context.Context, identifier entity.Identifier, ownerId, fileId int) (files []daoDTO.FullFile, err error)
	SaveFile(ctx context.Context, identifier entity.Identifier, file entity.File, client string) (id int, err error)
	RenameFile(ctx context.Context, identifier entity.Identifier, oldFilName, newFileName, client string) (err error)
	DeleteFile(ctx context.Context, identifier entity.Identifier, fileName, client string) (err error)
}

// IOwnerUsecase.
//go:generate mockgen -destination=../../mocks/grpc/server/owner/usecase.go -package=grpc -source=server.go
type IOwnerUsecase interface {
	CreateOwner(ctx context.Context, owner entity.Owner) (int, error)
	DeleteOwner(ctx context.Context, id int) error
}

// server.
type server struct {
	pb.UnimplementedFileManagerServer
	pb.UnimplementedOwnerManagerServer

	fileUsecase  IFileUsecase
	ownerUsecase IOwnerUsecase
	socket       string
	log          *logrus.Logger
}

// GrpcServerConf.
type GrpcServerConf struct {
	Socket       string
	FileUsecase  IFileUsecase
	OwnerUsecase IOwnerUsecase
}

// NewServer.
func NewServer(conf GrpcServerConf) server {

	return server{
		fileUsecase:  conf.FileUsecase,
		ownerUsecase: conf.OwnerUsecase,
		socket:       conf.Socket,
	}
}

// Run. 
func (s *server) Run() error {

	listener, err := net.Listen("tcp", s.socket)
	if err != nil {
		return err
	}

	baseServer := grpc.NewServer(grpc.EmptyServerOption{})
	pb.RegisterFileManagerServer(baseServer, &server{})
	pb.RegisterOwnerManagerServer(baseServer, &server{})
	return baseServer.Serve(listener)
}
