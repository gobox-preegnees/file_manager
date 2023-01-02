package services

import (
	"context"
	"encoding/json"

	mbDTO "github.com/gobox-preegnees/file_manager/internal/adapters/message_broker"
	"github.com/gobox-preegnees/file_manager/internal/domain/entity"

	"github.com/sirupsen/logrus"
)

type IMessageBroker interface {
	PublishErr(mbDTO.PublishErrReqDTO) error
}

type service struct {
	ctx           context.Context
	log           *logrus.Logger
	messageBroker IMessageBroker
}

type ConfServices struct {
	Ctx           context.Context
	Log           *logrus.Logger
	MessageBroker IMessageBroker
}

func NewServices(cnf ConfServices) *service {

	return &service{
		ctx:           cnf.Ctx,
		log:           cnf.Log,
		messageBroker: cnf.MessageBroker,
	}
}

func (s *service) SendMessage(message entity.Message) error {

	if message.IsErr {
		jData, err := json.Marshal(message)
		if err != nil {
			return err
		}
		return s.messageBroker.PublishErr(mbDTO.PublishErrReqDTO{
			Error: jData,
		})
	}
	return nil
}


