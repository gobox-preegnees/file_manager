package main

import (
	"log"
	"time"

	controllerGrpc "github.com/gobox-preegnees/file_manager/v1/pkg/controller/grpc/server"
	stubCli "github.com/gobox-preegnees/file_manager/v1/pkg/controller/grpc/stub"
)

func main() {

	server := controllerGrpc.NewServer()
	go func() {
		log.Fatal(server.Run())
	}()
	time.Sleep(2 * time.Second)
	stubCli.Get()
	time.Sleep(5 * time.Second)
}