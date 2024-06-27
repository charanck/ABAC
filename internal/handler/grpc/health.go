package grpchandler

import (
	"context"

	abac "github.com/charanck/ABAC/protobuf/generated"
)

type HealthServer struct {
	abac.UnimplementedHealthServer
}

func (t HealthServer) Health(context.Context, *abac.Void) (*abac.HealthResponse, error) {
	return &abac.HealthResponse{
		Message: "OK",
	}, nil
}
