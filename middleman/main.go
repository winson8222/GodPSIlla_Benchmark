package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"middleman/constants"
	"middleman/model"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", constants.GRPC_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	model.RegisterMiddlemanServer(s, &server{})
	log.Printf("middleman listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
