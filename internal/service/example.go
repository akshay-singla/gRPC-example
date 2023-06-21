package service

import (
	"context"

	crudpb "github.com/akshay-singla/gRPC-example/internal/pb/example"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	crudpb.UnimplementedCrudServiceServer
	DB map[string]string
}

func (s *Server) Create(ctx context.Context, req *crudpb.CreateRequest) (*crudpb.CreateResponse, error) {
	id := req.GetId()
	data := req.GetData()

	s.DB[id] = data

	return &crudpb.CreateResponse{Success: true}, nil
}

func (s *Server) Read(ctx context.Context, req *crudpb.ReadRequest) (*crudpb.ReadResponse, error) {
	id := req.GetId()

	data, ok := s.DB[id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Record not found")
	}

	return &crudpb.ReadResponse{Data: data}, nil
}

func (s *Server) Update(ctx context.Context, req *crudpb.UpdateRequest) (*crudpb.UpdateResponse, error) {
	id := req.GetId()
	data := req.GetData()

	_, ok := s.DB[id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Record not found")
	}

	s.DB[id] = data

	return &crudpb.UpdateResponse{Success: true}, nil
}

func (s *Server) Delete(ctx context.Context, req *crudpb.DeleteRequest) (*crudpb.DeleteResponse, error) {
	id := req.GetId()

	_, ok := s.DB[id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Record not found")
	}

	delete(s.DB, id)

	return &crudpb.DeleteResponse{Success: true}, nil
}
