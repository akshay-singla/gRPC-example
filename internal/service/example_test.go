package service_test

import (
	"context"
	"testing"

	crudpb "github.com/akshay-singla/gRPC-example/internal/pb/example"
	"github.com/akshay-singla/gRPC-example/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestCrudServer_Create(t *testing.T) {
	srv := &service.Server{DB: make(map[string]string)}
	req := &crudpb.CreateRequest{
		Id:   "1",
		Data: "Test data",
	}

	res, err := srv.Create(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.True(t, res.Success)
}

func TestCrudServer_Read(t *testing.T) {
	srv := &service.Server{DB: map[string]string{"1": "Test data"}}
	req := &crudpb.ReadRequest{
		Id: "1",
	}

	res, err := srv.Read(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "Test data", res.Data)
}

func TestCrudServer_Update(t *testing.T) {
	srv := &service.Server{DB: map[string]string{"1": "Test data"}}
	req := &crudpb.UpdateRequest{
		Id:   "1",
		Data: "Updated data",
	}

	res, err := srv.Update(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.True(t, res.Success)
	assert.Equal(t, "Updated data", srv.DB["1"])
}

func TestCrudServer_Delete(t *testing.T) {
	srv := &service.Server{DB: map[string]string{"1": "Test data"}}
	req := &crudpb.DeleteRequest{
		Id: "1",
	}

	res, err := srv.Delete(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.True(t, res.Success)
	_, ok := srv.DB["1"]
	assert.False(t, ok)
}
