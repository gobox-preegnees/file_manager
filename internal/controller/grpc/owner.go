package grpc

import (
	"context"

	pb "github.com/gobox-preegnees/file_manager/api/grpc"
	dtoService "github.com/gobox-preegnees/file_manager/internal/domain"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"
)

func (s *server) CreateOwner(ctx context.Context, in *pb.CreateOwnerReq) (*pb.CreateOwnerResp, error) {

	id, err := s.ownerService.CreateOwner(ctx, dtoService.CreateOwnerReqDTO{
		Owner: entity.Owner{
			Username: in.Username,
			Folder:   in.Folder,
		},
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateOwnerResp{Id: int64(id)}, nil
}

func (s *server) DeleteOwner(ctx context.Context, in *pb.DeleteOwnerReq) (*pb.DeleteOwnerResp, error) {

	return nil, s.ownerService.DeleteOwner(ctx, dtoService.DeleteOwnerReqDTO{Id: int(in.Id)})
}
