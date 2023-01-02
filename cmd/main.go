package main

import (
	"context"

	postgresqlAdapter "github.com/gobox-preegnees/file_manager/internal/adapters/dao/postgresql"
	kafkaAdapter "github.com/gobox-preegnees/file_manager/internal/adapters/message_broker/kafka"
	config "github.com/gobox-preegnees/file_manager/internal/config"
	grpcController "github.com/gobox-preegnees/file_manager/internal/controller/grpc"
	kafkaController "github.com/gobox-preegnees/file_manager/internal/controller/kafka"
	services "github.com/gobox-preegnees/file_manager/internal/domain/services"
	usecase "github.com/gobox-preegnees/file_manager/internal/domain/usecase"
	"golang.org/x/sync/errgroup"

	"github.com/sirupsen/logrus"
)

func main() {
	
	cnf := config.GetConfig("C:\\Users\\secrr\\Desktop\\fileManagerNew\\file_manager\\cnf.yml")
	ctx := context.TODO()

	dao, err := postgresqlAdapter.NewPosgresql(ctx, cnf.Pg)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	if cnf.Debug {
		logger.SetLevel(logrus.DebugLevel)
		logger.SetReportCaller(true)
	}

	messageBroker := kafkaAdapter.NewProducer(kafkaAdapter.ProducerConf{
		Log:   logger,
		Topic: "errors",
		Addrs: []string{cnf.Kafka},
	})
	defer messageBroker.Close()

	service := services.NewServices(services.ConfServices{
		Ctx:           ctx,
		Log:           logger,
		MessageBroker: messageBroker,
	})
	
	fileUsecase := usecase.NewFileUsecase(logger, dao)
	ownerUsecase := usecase.NewOwnerUsecase(logger, dao, service)
	statesUsecase := usecase.NewStateUsecase(logger, dao, service)

	grpcServer := grpcController.NewServer(grpcController.GrpcServerConf{
		Socket:       cnf.App,
		FileUsecase:  fileUsecase,
		OwnerUsecase: ownerUsecase,
	})
	
	g := new(errgroup.Group)

	g.Go(func() error {
        return grpcServer.Run()
	})

	consumer := kafkaController.NewConsumer(kafkaController.ConsumerCnf{
		Ctx: ctx,
		Log: logger,
		Topic: "states",
		Addrs: []string{cnf.Kafka},
		GroupId: "states_id",
		Partition: 0,
		StateUsecase: statesUsecase,
	})
	
    g.Go(func() error {
        return consumer.Run()
	})

	if err := g.Wait(); err!= nil {
        panic(err)
	}
}
