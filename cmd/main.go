package main

import (
	"context"
	// "os"

	config "github.com/gobox-preegnees/file_manager/internal/config"

	grpcController "github.com/gobox-preegnees/file_manager/internal/controller/grpc"
	kafkaController "github.com/gobox-preegnees/file_manager/internal/controller/kafka"

	postgresAdapter "github.com/gobox-preegnees/file_manager/internal/adapters/dao/postgresql"
	kafkaAdapter "github.com/gobox-preegnees/file_manager/internal/adapters/message_broker/kafka"
	service "github.com/gobox-preegnees/file_manager/internal/domain/service"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func main() {
	var path string
	// if  == "" {
	path = "C:\\Users\\secrr\\Desktop\\fileManagerNew\\file_manager\\cnf.yml"
	// } else {
	// path = os.Args[1]
	// }
	cnf := config.GetConfig(path)

	logger := logrus.New()
	logger.SetReportCaller(true)
	if cnf.Debug {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.ErrorLevel)
	}

	ctx, cancel := context.WithCancel(context.TODO())

	dao, err := postgresAdapter.NewPosgresql(postgresAdapter.CnfPostgres{
		Ctx: ctx,
		Url: cnf.Pg.Url,
	})
	if err != nil {
		logger.Fatal(err)
	}

	messageBroker := kafkaAdapter.NewProducer(kafkaAdapter.ProducerConf{
		Log:       logger,
		ErrTopic:  cnf.Kafka.Producer.ErrTopic,
		Addrs:     cnf.Kafka.Addr,
		Attempts:  cnf.Kafka.Producer.Attempts,
		Timeout:   cnf.Kafka.Producer.Timeout,
		Sleeptime: cnf.Kafka.Producer.Sleeptime,
	})
	messageService := service.NewMessageServices(service.CnfMessageServices{
		Ctx:           ctx,
		Log:           logger,
		MessageBroker: messageBroker,
	})

	stateService := service.NewStateService(service.CnfStateService{
		Log:            logger,
		DaoState:       dao,
		ServiceMessage: messageService,
	})
	fileService := service.NewFileService(service.CnfFileService{
		Log:     logger,
		DaoFile: dao,
	})
	ownerService := service.NewOwnerService(service.CnfOwnerService{
		Log:            logger,
		DaoOwner:       dao,
		ServiceMessage: messageService,
	})

	gController := grpcController.NewServer(grpcController.CnfGrpcServer{
		Log:          logger,
		Address:      cnf.App.Addr,
		FileService:  fileService,
		OwnerService: ownerService,
	})
	kController := kafkaController.NewConsumer(kafkaController.ConsumerCnf{
		Ctx:          ctx,
		Log:          logger,
		Topic:        cnf.Kafka.Consumer.StateTopic,
		Addrs:        cnf.Kafka.Addr,
		GroupId:      cnf.Kafka.Consumer.GroupId,
		Partition:    cnf.Kafka.Consumer.Partition,
		StateService: stateService,
	})

	g := new(errgroup.Group)
	g.Go(func() error {
		return gController.Run()
	})
	g.Go(func() error {
		return kController.Run()
	})
	if err := g.Wait(); err != nil {
		cancel()
		logger.Fatal(err)
	}
}
