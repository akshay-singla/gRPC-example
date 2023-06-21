package main

import (
	"context"
	"log"

	crudpb "github.com/akshay-singla/gRPC-example/internal/pb/example"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := crudpb.NewCrudServiceClient(conn)

	ctx := context.Background()

	// Create
	createReq := &crudpb.CreateRequest{
		Id:   uuid.New().String(),
		Data: "Some data",
	}
	createRes, err := client.Create(ctx, createReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create response: %v", createRes)

	// Read
	readReq := &crudpb.ReadRequest{
		Id: createReq.GetId(),
	}
	readRes, err := client.Read(ctx, readReq)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Printf("Read response: %v", readRes)

	// Update
	updateReq := &crudpb.UpdateRequest{
		Id:   createReq.GetId(),
		Data: "Updated data",
	}
	updateRes, err := client.Update(ctx, updateReq)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update response: %v", updateRes)

	// Delete
	deleteReq := &crudpb.DeleteRequest{
		Id: createReq.GetId(),
	}
	deleteRes, err := client.Delete(ctx, deleteReq)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete response: %v", deleteRes)
}
