package main

import (
	"grpc-go"
	"log"
	"net"

	"github.com/tuthahthien02/grpc-course/caculator/caculatorpb"
)

type server struct{}

const PORT = "0.0.0.0:50059"

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	caculatorpb.RegisterCalulatorServer(s, &server{})

	err = s.Server(lis)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
