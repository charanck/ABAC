package grpchandler

import (
	"context"
	"errors"

	"github.com/charanck/ABAC/internal/handler/grpc/adapter"
	"github.com/charanck/ABAC/internal/service"
	"github.com/charanck/ABAC/internal/util"
	abac "github.com/charanck/ABAC/protobuf/generated"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type Attribute struct {
	abac.UnimplementedAttributeServer
	Service service.Attribute
}

func NewAttribute(service service.Attribute) Attribute {
	return Attribute{
		Service: service,
	}
}

func (a *Attribute) CreateAttribute(ctx context.Context, request *abac.CreateAttributeRequest) (*abac.CreateAttributeResponse, error) {
	attributeTypes := map[string]bool{"string": true, "integer": true, "float": true, "bool": true, "date": true}
	if _, ok := attributeTypes[request.Type]; !ok {
		return nil, status.Error(codes.InvalidArgument, "invalid type")
	}
	attribute := adapter.CreateAttributeRequestToModel(request)
	attributeId, err := a.Service.Create(attribute)
	if err != nil {
		var apiError util.ApiError
		if errors.As(err, &apiError) {
			return nil, status.Error(apiError.GRPCErrorCode, apiError.ErrorMessage)
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &abac.CreateAttributeResponse{
		Id: attributeId,
	}, nil
}
