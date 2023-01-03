package service

import (
	"context"
	"time"

	daoDTO "github.com/gobox-preegnees/file_manager/internal/adapters/dao"
	kafkaController "github.com/gobox-preegnees/file_manager/internal/controller/kafka"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"

	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../../mocks/domain/service/dao_state/state.go -package=dao_state -source=state.go
type IDaoState interface {
	SetState(ctx context.Context, setStateReqDTO daoDTO.SetStateReqDTO) error
}

//go:generate mockgen -destination=../../mocks/domain/service/service_message_state/state.go -package=service_message_state -source=state.go
type IServiceMessageState interface {
	SendMessage(message entity.Message) error
}

type stateService struct {
	log            *logrus.Logger
	daoState       IDaoState
	serviceMessage IServiceMessageState
}

type CnfStateService struct {
	Log            *logrus.Logger
    DaoState       IDaoState
    ServiceMessage IServiceMessageState
}

func NewStateService(cnf CnfStateService) *stateService {

	return &stateService{
		log:            cnf.Log,
		daoState:       cnf.DaoState,
		serviceMessage: cnf.ServiceMessage,
	}
}

func (s stateService) SetState(ctx context.Context, state entity.State) {

	if err := s.daoState.SetState(ctx, daoDTO.SetStateReqDTO{
		Identifier: daoDTO.Identifier{
			Username: state.Username,
			Folder:   state.Folder,
		},
		FileName:    state.FileName,
		VirtualName: state.VirtualName,
		State:       state.State,
	}); err != nil {
		s.log.Error(err)
		if err := s.serviceMessage.SendMessage(entity.Message{
			Message:   err.Error(),
			Timestamp: time.Now().UTC().Unix(),
			IsErr:     true,
		}); err != nil {
			s.log.Fatal(err)
		}
	}
}

var _ kafkaController.IStateSerivce = (*stateService)(nil)
