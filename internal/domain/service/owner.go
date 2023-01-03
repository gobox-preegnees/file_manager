package service

import (
	"context"
	"time"

	daoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/dao"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"

	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../../mocks/domain/service/dao_owner/owner.go -package=dao_owner -source=owner.go
type IDaoOwner interface {
	DeleteOwner(ctx context.Context, deleteOwnerDTO daoDTO.DeleteOwnerDTO) error
	SaveOwner(ctx context.Context, saveOwnerDTO daoDTO.SaveOwnerDTO) (int, error)
}

//go:generate mockgen -destination=../../mocks/domain/service/service_message_owner/owner.go -package=service_message_owner -source=owner.go
type IServiceMessageOwner interface {
	SendMessage(message entity.Message) error
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

func (o ownerService) CreateOwner(ctx context.Context, owner entity.Owner) (int, error) {

	id, err := o.dao.SaveOwner(ctx, daoDTO.SaveOwnerDTO{
		Identifier: daoDTO.Identifier{},
	})
	if err != nil {
		o.log.Error(err)
		if err := o.serviceMessage.SendMessage(entity.Message{
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

func (o ownerService) DeleteOwner(ctx context.Context, id int) error {

	if err := o.dao.DeleteOwner(ctx, daoDTO.DeleteOwnerDTO{
		OwnerID: id,
	}); err != nil {
		o.log.Error(err)
		if err := o.serviceMessage.SendMessage(entity.Message{
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
