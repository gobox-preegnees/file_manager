package kafka

import (
	"context"
	"time"
	"errors"

	mbDTO "github.com/gobox-preegnees/file_manager/internal/adapters/message_broker"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type producer struct {
	log            *logrus.Logger
	writer         *kafka.Writer
}

type ProducerConf struct {
	Log     *logrus.Logger
	Topic   string
	Addrs   []string
}

func NewProducer(cnf ProducerConf) *producer {

	w := &kafka.Writer{
		Addr:                   kafka.TCP(cnf.Addrs...),
		Topic:                  cnf.Topic,
		AllowAutoTopicCreation: true,
		Logger:                 cnf.Log,
	}
	return &producer{
		log:     cnf.Log,
		writer:  w,
	}
}

func (p *producer) PublishErr(publishErrReqDTO mbDTO.PublishErrReqDTO) error {

	for i := 0; i < 3; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := p.writer.WriteMessages(ctx, kafka.Message{
			Value: publishErrReqDTO.Error,
		})
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}
		if err != nil {
			return err
		}
	}
    return nil
}

func (p *producer) Close() error {

	return p.writer.Close()
}
