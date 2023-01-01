package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"

	"github.com/go-playground/validator/v10"
	kafkaGo "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../../mocks/kafka/consumer/usecase.go -package=mock -source=consumer.go
type IUsecase interface {
	SetState(entity.State)
}

var (
	ErrValidate    = errors.New("Err Validate")
	ErrInvalidData = errors.New("Err Invalid Data")
	ErrReadMessage = errors.New("Err Read Message consumer")
)

// kafka.
type kafka struct {
	ctx     context.Context
	log     *logrus.Logger
	reader  *kafkaGo.Reader
	usecase IUsecase
}

// KafkaCnf. Config for consumer
type KafkaConsumerCnf struct {
	Ctx       context.Context
	Log       *logrus.Logger
	Topic     string
	Addresses []string
	GroupId   string
	Partition int
	Usecase   IUsecase
}

// New. Create new consumer instance
func New(cnf KafkaConsumerCnf) *kafka {

	reader := kafkaGo.NewReader(kafkaGo.ReaderConfig{
		Brokers:  cnf.Addresses,
		GroupID:  cnf.GroupId,
		Topic:    cnf.Topic,
		Logger:   cnf.Log,
		MinBytes: 0,
	})

	if reader == nil {
		panic("Reader is nil")
	}

	return &kafka{
		ctx:     cnf.Ctx,
		log:     cnf.Log,
		reader:  reader,
		usecase: cnf.Usecase,
	}
}

// Run. Run consuming message.
// Returning: err when reader is not reading | err when validation | err when set state
func (k *kafka) Run() error {

	defer k.reader.Close()

	for {
		msg, err := k.reader.ReadMessage(k.ctx)
		if errors.Is(err, context.Canceled) {
			return nil
		} else if err != nil {
			return fmt.Errorf("$%w {error:%v}", ErrReadMessage, err)
		}
		k.log.Debugf("msg: %v", msg)

		state, err := k.validateMessage(msg.Value)
		if err != nil {
			return err
		}
		k.log.Debugf("state: %v", state)

		k.usecase.SetState(state)
		k.log.Debugf("success set state: %v", state)
	}
}

// validateMessage. Conducts validation
func (k *kafka) validateMessage(msg []byte) (entity.State, error) {

	validate := validator.New()
	var state entity.State

	err := json.Unmarshal([]byte(msg), &state)
	if err != nil {
		return entity.State{}, fmt.Errorf("$%w {msg:{%v}} {error:%v}", ErrInvalidData, msg, err)
	}

	err = validate.Struct(&state)
	if err != nil {
		return entity.State{}, fmt.Errorf("$%w {msg:{%v}} {error:%v}", ErrValidate, msg, err)
	}
	return state, nil
}
