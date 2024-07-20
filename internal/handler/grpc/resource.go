package grpchandler

import (
	"context"

	"github.com/charanck/ABAC/internal/handler/grpc/adapter"
	"github.com/charanck/ABAC/internal/service"
	abac "github.com/charanck/ABAC/protobuf/generated"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type Resource struct {
	abac.UnimplementedResourceServer
	resourceService service.Resource
}

func NewResource(resourceService service.Resource) Resource {
	return Resource{
		resourceService: resourceService,
	}
}

func (r *Resource) CreateResource(ctx context.Context, request *abac.CreateResourceRequest) (*abac.CreateResourceResponse, error) {
	id, err := r.resourceService.Create(adapter.ResourceGRPCToModel(request))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	response := abac.CreateResourceResponse{
		Id: id,
	}
	return &response, nil
}

func (r *Resource) GetResource(ctx context.Context, request *abac.GetResourceRequest) (*abac.GetResourceResponse, error) {
	return nil, nil
}

func (r *Resource) UpdateResource(ctx context.Context, request *abac.UpdateResourceRequest) (*abac.UpdateResourceResponse, error) {
	return nil, nil
}

func (r *Resource) DeleteResource(ctx context.Context, request *abac.DeleteResourceRequest) (*abac.DeleteResourceResponse, error) {
	return nil, nil
}

func (r *Resource) ListResource(ctx context.Context, request *abac.ListResourceRequest) (*abac.ListResourceResponse, error) {
	return nil, nil
}
