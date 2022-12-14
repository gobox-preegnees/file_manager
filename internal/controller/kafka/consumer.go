package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	dtoService "github.com/gobox-preegnees/file_manager/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../../mocks/kafka/consumer/state/usecase.go -package=kafka -source=consumer.go
type IStateSerivce interface {
	SetState(context.Context, dtoService.SetStateReqDTO)
}

var (
	ErrValidate    = errors.New("Err Validate")
	ErrInvalidData = errors.New("Err Invalid Data")
	ErrReadMessage = errors.New("Err Read Message consumer")
)

// kafka.
type consumer struct {
	ctx          context.Context
	log          *logrus.Logger
	reader       *kafka.Reader
	stateService IStateSerivce
}

// KafkaCnf. Config for consumer
type ConsumerCnf struct {
	Ctx          context.Context
	Log          *logrus.Logger
	Topic        string
	Addrs        []string
	GroupId      string
	Partition    int
	StateService IStateSerivce
}

// New. Create new consumer instance
func NewConsumer(cnf ConsumerCnf) *consumer {

	if conn, err := kafka.Dial("tcp", cnf.Addrs[0]); err != nil {
		conn.Close()
		cnf.Log.Fatal(err)
	} else {
		conn.Close()
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  cnf.Addrs,
		GroupID:  cnf.GroupId,
		Topic:    cnf.Topic,
		Logger:   cnf.Log,
		MinBytes: 0,
	})

	return &consumer{
		ctx:          cnf.Ctx,
		log:          cnf.Log,
		reader:       reader,
		stateService: cnf.StateService,
	}
}

// Run. Run consuming message.
// Returning: err when reader is not reading | err when validation | err when set state
func (k *consumer) Run() error {

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

		go func() {
			k.stateService.SetState(context.Background(), state)
			k.log.Debugf("success set state: %v", state)
		}()
	}
}

// validateMessage. Conducts validation
func (k *consumer) validateMessage(msg []byte) (dtoService.SetStateReqDTO, error) {

	validate := validator.New()
	var state dtoService.SetStateReqDTO

	err := json.Unmarshal(msg, &state)
	if err != nil {
		return dtoService.SetStateReqDTO{}, fmt.Errorf("$%w {msg:{%v}} {error:%v}", ErrInvalidData, msg, err)
	}

	err = validate.Struct(&state)
	if err != nil {
		return dtoService.SetStateReqDTO{}, fmt.Errorf("$%w {msg:{%v}} {error:%v}", ErrValidate, msg, err)
	}
	return state, nil
}
