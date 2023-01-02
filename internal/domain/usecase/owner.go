package usecase

import (
	"context"
	"time"

	daoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/dao"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"

	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../../mocks/domain/usecase/dao/owner/owner.go -package=usecase_dao_owner -source=file.go
type IDaoOwner interface {
	DeleteOwner(ctx context.Context, deleteOwnerDTO daoDTO.DeleteOwnerDTO) error
	SaveOwner(ctx context.Context, saveOwnerDTO daoDTO.SaveOwnerDTO) (int, error)
}

//go:generate mockgen -destination=../../mocks/domain/usecase/service/owner/owner.go -package=usecase_service_owner -source=file.go
type IServiceOwner interface {
	SendMessage(message entity.Message) error
}

type ownerUsecase struct {
	log     *logrus.Logger
	dao     IDaoOwner
	service IServiceState
}

func NewOwnerUsecase(log *logrus.Logger, dao IDaoOwner, service IServiceOwner) *ownerUsecase {

	return &ownerUsecase{
		log:     log,
		dao:     dao,
		service: service,
	}
}

func (o *ownerUsecase) CreateOwner(ctx context.Context, owner entity.Owner) (int, error) {

	id, err := o.dao.SaveOwner(ctx, daoDTO.SaveOwnerDTO{
		Identifier: daoDTO.Identifier{},
	})
	if err != nil {
		o.log.Error(err)
		if err := o.service.SendMessage(entity.Message{
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

func (o *ownerUsecase) DeleteOwner(ctx context.Context, id int) error {

	if err := o.dao.DeleteOwner(ctx, daoDTO.DeleteOwnerDTO{
		OwnerID: id,
	}); err != nil {
		o.log.Error(err)
		if err := o.service.SendMessage(entity.Message{
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
