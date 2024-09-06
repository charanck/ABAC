package grpchandler

import (
	"context"
	"errors"
	"strings"

	"github.com/charanck/ABAC/internal/handler/grpc/adapter"
	"github.com/charanck/ABAC/internal/service"
	"github.com/charanck/ABAC/internal/util"
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
	id, err := r.resourceService.Create(adapter.CreateResourceRequestToModel(request))
	if err != nil {
		var apiError util.ApiError
		if errors.As(err, &apiError) {
			return nil, status.Error(apiError.GRPCErrorCode, apiError.ErrorMessage)
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	response := abac.CreateResourceResponse{
		Id: id,
	}
	return &response, nil
}

func (r *Resource) GetResource(ctx context.Context, request *abac.GetResourceRequest) (*abac.GetResourceResponse, error) {
	resource, err := r.resourceService.GetById(request.GetId())
	if err != nil {
		var apiError util.ApiError
		if errors.As(err, &apiError) {
			return nil, status.Error(apiError.GRPCErrorCode, apiError.ErrorMessage)
		}
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return adapter.ModelToGetResourceResponse(resource), nil
}

func (r *Resource) UpdateResource(ctx context.Context, request *abac.UpdateResourceRequest) (*abac.UpdateResourceResponse, error) {
	if strings.TrimSpace(request.Data.Id) == "" {
		return nil, status.Errorf(codes.InvalidArgument, "resource id is required")
	}
	if !request.FieldMask.IsValid(&abac.UpdateResourceRequest_Data{}) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid fieldmask")
	}
	request.FieldMask.Normalize()
	_, err := r.resourceService.UpdateById(adapter.UpdateResourceRequestToModel(request), request.FieldMask.Paths)
	if err != nil {
		var apiError util.ApiError
		if errors.As(err, &apiError) {
			return nil, status.Error(apiError.GRPCErrorCode, apiError.ErrorMessage)
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return &abac.UpdateResourceResponse{
		Id: request.Data.Id,
	}, nil
}

func (r *Resource) DeleteResource(ctx context.Context, request *abac.DeleteResourceRequest) (*abac.DeleteResourceResponse, error) {
	resourceId, err := r.resourceService.DeleteById(request.Id)
	if err != nil {
		var apiError util.ApiError
		if errors.As(err, &apiError) {
			return nil, status.Error(apiError.GRPCErrorCode, apiError.ErrorMessage)
		}
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return &abac.DeleteResourceResponse{Id: resourceId}, nil
}

func (r *Resource) ListResource(ctx context.Context, request *abac.ListResourceRequest) (*abac.ListResourceResponse, error) {
	pageNumber, pageSize := 1, 10
	if request != nil && request.PagingMetadata != nil {
		pageNumber = int(request.PagingMetadata.GetPageNumber())
		pageSize = int(request.PagingMetadata.GetPageSize())
	}
	resources, err := r.resourceService.List(pageNumber, pageSize)
	if err != nil {
		var apiError util.ApiError
		if errors.As(err, &apiError) {
			return nil, status.Error(apiError.GRPCErrorCode, apiError.ErrorMessage)
		}
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return adapter.ModelToListResourceResponse(resources), nil
}
