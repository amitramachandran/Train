package main

import (
	"log"
	"net"

	bt "github.com/amitramachandran/train-grpc/boardTrain"
	"github.com/amitramachandran/train-grpc/server/handlers"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8099")
	log.Printf("listening on %s\n", lis.Addr())
	if err != nil {
		log.Fatalf("Cannot create listener: %+v", err)
	}

	serviceRegister := grpc.NewServer()

	bt.RegisterBoardingServer(serviceRegister, handlers.NewServer())
	err = serviceRegister.Serve(lis)
	if err != nil {
		log.Printf("impossible to serve %s", err)
		serviceRegister.GracefulStop()
	}
}
