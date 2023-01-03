package service

import (
	"context"
	"encoding/json"

	mbDTO "github.com/gobox-preegnees/file_manager/internal/adapters/message_broker"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"

	"github.com/sirupsen/logrus"
)

type message struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	IsErr     bool   `json:"is_err"`
}

//go:generate mockgen -destination=../../mocks/domain/service/message_broker_message/message.go -package=message_broker_message -source=message.go
type IMessageBroker interface {
	PublishErr(mbDTO.PublishErrReqDTO) error
}

type messageService struct {
	ctx           context.Context
	log           *logrus.Logger
	messageBroker IMessageBroker
}

type CnfMessageServices struct {
	Ctx           context.Context
	Log           *logrus.Logger
	MessageBroker IMessageBroker
}

func NewMessageServices(cnf CnfMessageServices) *messageService {

	return &messageService{
		ctx:           cnf.Ctx,
		log:           cnf.Log,
		messageBroker: cnf.MessageBroker,
	}
}

func (s messageService) SendMessage(message entity.Message) error {

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


