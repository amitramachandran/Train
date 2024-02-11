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
	if err != nil {
		log.Fatalf("Cannot create listener: %+v", err)
	}

	serviceRegister := grpc.NewServer()

	bt.RegisterBoardingServer(serviceRegister, handlers.NewServer())
	err = serviceRegister.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve %s", err)
	}
}
