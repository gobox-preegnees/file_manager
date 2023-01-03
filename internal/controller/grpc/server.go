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

// IFileService.
//go:generate mockgen -destination=../../mocks/grpc/file_service/service.go -package=file_service -source=server.go
type IFileService interface {
	GetFiles(ctx context.Context, identifier entity.Identifier, ownerId, fileId int) (files []daoDTO.FullFile, err error)
	SaveFile(ctx context.Context, identifier entity.Identifier, file entity.File, client string) (id int, err error)
	RenameFile(ctx context.Context, identifier entity.Identifier, oldFilName, newFileName, client string) (err error)
	DeleteFile(ctx context.Context, identifier entity.Identifier, fileName, client string) (err error)
}

// IOwnerService.
//go:generate mockgen -destination=../../mocks/grpc/server/owner_service/owner.go -package=file_service -source=server.go
type IOwnerService interface {
	CreateOwner(ctx context.Context, owner entity.Owner) (int, error)
	DeleteOwner(ctx context.Context, id int) error
}

// server.
type server struct {
	pb.UnimplementedFileManagerServer
	pb.UnimplementedOwnerManagerServer

	fileService  IFileService
	ownerService IOwnerService
	socket       string
	log          *logrus.Logger
}

// GrpcServerConf.
type GrpcServerConf struct {
	Socket       string
	FileService  IFileService
	OwnerService IOwnerService
}

// NewServer.
func NewServer(conf GrpcServerConf) server {

	return server{
		fileService:  conf.FileService,
		ownerService: conf.OwnerService,
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
