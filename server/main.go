package main

import (
	"log"
	"net"

	"github.com/emreturgutce/grpc-tutorialedge/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
