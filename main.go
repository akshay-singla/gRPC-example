package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	crudpb "github.com/akshay-singla/gRPC-example/internal/pb/example"
	"github.com/akshay-singla/gRPC-example/internal/service"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	crudpb.RegisterCrudServiceServer(s, &service.Server{
		DB: make(map[string]string),
	})

	log.Println("Server started on localhost:50051")

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
