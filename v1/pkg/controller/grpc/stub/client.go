package studclient

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/gobox-preegnees/file_manager/v1/pkg/contract"
)

const SOCKET = "localhost:50051"

func Get() {
	conn, err := grpc.Dial(SOCKET, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFileManagerClient(conn)

	r, err := c.SaveFile(context.TODO(), &pb.SaveFileReq{
		Identification: &pb.Identification{
			Username: "username",
		},
		File: &pb.File{
			Path: "path_____",
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Description)
}