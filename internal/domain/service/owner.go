package service

import (
	"context"
	"time"

	grpcController "github.com/gobox-preegnees/file_manager/internal/controller/grpc"
	daoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/dao"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"
	dtoService "github.com/gobox-preegnees/file_manager/internal/domain"

	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../../mocks/domain/service/dao_owner/owner.go -package=dao_owner -source=owner.go
type IDaoOwner interface {
	DeleteOwner(ctx context.Context, deleteOwnerDTO daoDTO.DeleteOwnerDTO) error
	SaveOwner(ctx context.Context, saveOwnerDTO daoDTO.SaveOwnerDTO) (int, error)
}

//go:generate mockgen -destination=../../mocks/domain/service/service_message_owner/owner.go -package=service_message_owner -source=owner.go
type IServiceMessageOwner interface {
	sendMessage(message entity.Message) error
}

type ownerService struct {
	log     *logrus.Logger
	dao     IDaoOwner
	serviceMessage IServiceMessageOwner
}

func NewOwnerUsecase(log *logrus.Logger, dao IDaoOwner, serviceMessage IServiceMessageOwner) *ownerService {

	return &ownerService{
		log:     log,
		dao:     dao,
		serviceMessage: serviceMessage,
	}
}

func (o ownerService) CreateOwner(ctx context.Context, createOwnerReqDTO dtoService.CreateOwnerReqDTO) (int, error) {

	id, err := o.dao.SaveOwner(ctx, daoDTO.SaveOwnerDTO{
		Identifier: daoDTO.Identifier{
			Username: createOwnerReqDTO.Owner.Username,
			Folder: createOwnerReqDTO.Owner.Folder,
		},
	})
	if err != nil {
		o.log.Error(err)
		if err := o.serviceMessage.sendMessage(entity.Message{
			Message:   err.Error(),
			Timestamp: time.Now().UTC().Unix(),
			IsErr:     true,
		}); err != nil {
			o.log.Fatal(err)
		}
		return 0, err
	}
	return id, nil
}

func (o ownerService) DeleteOwner(ctx context.Context, deleteOwnerReqDTO dtoService.DeleteOwnerReqDTO) error {

	if err := o.dao.DeleteOwner(ctx, daoDTO.DeleteOwnerDTO{
		OwnerID: deleteOwnerReqDTO.Id,
	}); err != nil {
		o.log.Error(err)
		if err := o.serviceMessage.sendMessage(entity.Message{
			Message:   err.Error(),
			Timestamp: time.Now().UTC().Unix(),
			IsErr:     true,
		}); err != nil {
			o.log.Fatal(err)
		}
		return err
	}
	return nil
}

var _ grpcController.IOwnerService = (*ownerService)(nil)
