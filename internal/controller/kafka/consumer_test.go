package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/gobox-preegnees/file_manager/internal/domain/entity"
	mockCons "github.com/gobox-preegnees/file_manager/internal/mocks/kafka/consumer/state"

	"github.com/golang/mock/gomock"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func TestValidateModel(t *testing.T) {

	data := []struct {
		state entity.State
		err   bool
	}{
		{
			state: entity.State{
				Username:    "username",
				Folder:      "folder",
				FileName:    "fileName",
				VirtualName: "virtualName",
				FileSize:    16,
				State:       200,
			},
			err: false,
		},
		{
			state: entity.State{
				Username:    "username",
				Folder:      "folder",
				FileName:    "fileName",
				VirtualName: "virtualName",
				FileSize:    16,
				State:       100,
			},
			err: true,
		},
		{
			state: entity.State{
				Username:    "username",
				Folder:      "folder",
				FileName:    "fileName",
				VirtualName: "",
				FileSize:    16,
				State:       200,
			},
			err: true,
		},
	}

	kafka := NewConsumer(ConsumerCnf{
		Ctx:          context.TODO(),
		Log:          getLogger(),
		Topic:        "topic",
		Addrs:        []string{"localhost:29092"},
		GroupId:      "id",
		Partition:    0,
		StateService: nil,
	})

	for _, d := range data {
		t.Run("test", func(t *testing.T) {
			jsonData, err := json.Marshal(d.state)
			if err != nil {
				t.Fatal(err)
			}

			_, err = kafka.validateMessage(jsonData)
			if err != nil {
				if !d.err {
					t.Fatalf("Expected no error, got %v", err)
				}
			}
		})
	}
}

func TestConsumerWork(t *testing.T) {

	addr := []string{"localhost:29092"}
	topic := "topic"
	groupId := "id"
	partition := 0
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	clear(addr, topic)
	defer clear(addr, topic)

	states := []entity.State{
		{
			Username:    "1",
			Folder:      "1",
			FileName:    "1",
			VirtualName: "1",
			FileSize:    1,
			State:       200,
		},
		{
			Username:    "2",
			Folder:      "2",
			FileName:    "2",
			VirtualName: "2",
			FileSize:    2,
			State:       200,
		},
		{
			Username:    "3",
			Folder:      "3",
			FileName:    "3",
			VirtualName: "3",
			FileSize:    3,
			State:       200,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockConsumer := mockCons.NewMockIStateUsecase(mockCtrl)
	mockConsumer.EXPECT().SetState(gomock.Any(), gomock.Any()).AnyTimes()

	go func() {
		defer cancel()
		for _, s := range states {
			bs, err := json.Marshal(s)
			if err != nil {
				t.Fatal(err)
			}
			if err := produce(t, addr, topic, bs); err != nil {
				t.Fatal(err)
			}
		}
	}()

	cnf := ConsumerCnf{
		Ctx:          ctx,
		Log:          getLogger(),
		Topic:        topic,
		Addrs:        []string{addr[0]},
		GroupId:      groupId,
		Partition:    partition,
		StateService: mockConsumer,
	}
	kafka := NewConsumer(cnf)
	err := kafka.Run()
	if err != nil {
		t.Fatal(err)
	}
}

func clear(addrs []string, topic string) {

	conn, err := kafka.Dial("tcp", addrs[0])
	controller, err := conn.Controller()
	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	controllerConn.DeleteTopics(topic)
	if err != nil {
		panic(err)
	} else {
		controllerConn.Close()
		conn.Close()
	}
}

func produce(t *testing.T, addrs []string, topic string, message []byte) error {

	w := &kafka.Writer{
		Addr:                   kafka.TCP(addrs...),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}

	messages := []kafka.Message{
		{
			Value: message,
		},
	}

	const retries = 3
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := w.WriteMessages(ctx, messages...)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			return err
		}
	}

	if err := w.Close(); err != nil {
		return err
	}
	return nil
}

func getLogger() *logrus.Logger {

	logger := logrus.StandardLogger()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)
	return logger
}
