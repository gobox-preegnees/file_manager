package grpc

import (
	"context"
	"net"

	pb "github.com/gobox-preegnees/file_manager/api/grpc"
	daoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/dao"
	dtoService "github.com/gobox-preegnees/file_manager/internal/domain"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

//go:generate mockgen -destination=../../mocks/grpc/file_service/service.go -package=file_service -source=server.go
type IFileService interface {
	GetFiles(ctx context.Context, getFilesReqDTO dtoService.GetFilesReqDTO) (files []daoDTO.FullFile, err error)
	SaveFile(ctx context.Context, saveFileReqDTO dtoService.SaveFileReqDTO) (id int, err error)
	RenameFile(ctx context.Context, renameFileReqDTO dtoService.RenameFileReqDTO) (err error)
	DeleteFile(ctx context.Context, deleteFileReqDTO dtoService.DeleteFileReqDTO) (err error)
}

//go:generate mockgen -destination=../../mocks/grpc/server/owner_service/owner.go -package=file_service -source=server.go
type IOwnerService interface {
	CreateOwner(ctx context.Context, createOwnerReqDTO dtoService.CreateOwnerReqDTO) (id int, err error)
	DeleteOwner(ctx context.Context, deleteOwnerReqDTO dtoService.DeleteOwnerReqDTO) (err error)
}

// server.
type server struct {
	pb.UnimplementedFileManagerServer
	pb.UnimplementedOwnerManagerServer

	log          *logrus.Logger
	address      string
	fileService  IFileService
	ownerService IOwnerService
}

// GrpcServerConf.
type CnfGrpcServer struct {
	Log *logrus.Logger
	Address      string
	FileService  IFileService
	OwnerService IOwnerService
}

// NewServer.
func NewServer(cnf CnfGrpcServer) *server {

	return &server{
		log: cnf.Log,
		fileService:  cnf.FileService,
		ownerService: cnf.OwnerService,
		address:      cnf.Address,
	}
}

// Run.
func (s server) Run() error {

	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	s.log.Infof("File Manager server listening on %s", s.address)

	baseServer := grpc.NewServer(grpc.EmptyServerOption{})
	pb.RegisterFileManagerServer(baseServer, &server{})
	pb.RegisterOwnerManagerServer(baseServer, &server{})
	return baseServer.Serve(listener)
}
