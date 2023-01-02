package usecase

import (
	"context"
	"time"

	kafkaController "github.com/gobox-preegnees/file_manager/internal/controller/kafka"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"

	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../../mocks/domain/usecase/dao/state/state.go -package=usecase_dao_state -source=file.go
type IDaoState interface {
	SetState(entity.State) error
}
//go:generate mockgen -destination=../../mocks/domain/usecase/service/state/state.go -package=usecase_service_state -source=file.go
type IServiceState interface {
	SendMessage(message entity.Message) error
}

type stateUsecase struct {
	log      *logrus.Logger
	daoState IDaoState
	service  IServiceState
}

func NewStateUsecase(log *logrus.Logger, daoState IDaoState, service IServiceState) *stateUsecase {

	return &stateUsecase{
		log:      log,
		daoState: daoState,
		service: service,
	}
}

func (s *stateUsecase) SetState(ctx context.Context, state entity.State) {

	if err := s.daoState.SetState(state); err != nil {
		s.log.Error(err)
		if err := s.service.SendMessage(entity.Message{
			Message:   err.Error(),
			Timestamp: time.Now().UTC().Unix(),
			IsErr:     true,
		}); err != nil {
			s.log.Fatal(err)
		}
	}
}

var _ kafkaController.IStateUsecase = (*stateUsecase)(nil)
