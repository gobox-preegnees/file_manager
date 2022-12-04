package server

import (
	"context"
	"net"

	protobuf "github.com/gobox-preegnees/file_manager/pkg/proto"

	"google.golang.org/grpc"
)

type IFileWorker interface {
	// OpenFile 
}

const TCP = "tcp"
const SOCKET = "localhost:50051"

type server struct {
	protobuf.FileManagerServer
}

func New() server {
	
	return server{}
}

func (s server) Run() error {

	listener, err := net.Listen(TCP, SOCKET)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	protobuf.RegisterFileManagerServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (s *server) Recover(ctx context.Context, infoForRecover *protobuf.InfoForRecover) (*protobuf.Resp, error) {
	
	return nil, nil
}

func (s *server) Delete(ctx context.Context, infoForDeletion *protobuf.InfoForDeletion) (*protobuf.Resp, error) {
	
	return nil, nil
}

func (s *server) Rename(ctx context.Context, infoForRenaming *protobuf.InfoForRenaming) (*protobuf.Resp, error) {
	
	return nil, nil
}

func (s *server) Check(ctx context.Context, infoForCheck *protobuf.InfoForCheck) (*protobuf.Resp, error) {
	
	return nil, nil
}

func (s *server) SaveBatch(ctx context.Context, batch *protobuf.Batch) (*protobuf.Resp, error) {
	
	return nil, nil
}

func (s *server) GetBatch(ctx context.Context, getBatchReq *protobuf.GetBatchReq) (*protobuf.Batch, error) {
	
	return nil, nil
}

func (s *server) CreateFolder(ctx context.Context, createFolderReq *protobuf.CreateFolderReq) (*protobuf.Resp, error) {
	
	return nil, nil
}