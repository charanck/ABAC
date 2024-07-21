package adapter

import (
	"github.com/charanck/ABAC/internal/model"
	abac "github.com/charanck/ABAC/protobuf/generated"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateResourceRequestToModel(request *abac.CreateResourceRequest) model.Resource {
	return model.Resource{
		Name:        request.Name,
		OwnerId:     request.OwnerId,
		PolicyId:    request.PolicyId,
		Description: request.Description,
	}
}

func ModelToGetResourceResponse(resource model.Resource) *abac.GetResourceResponse {
	return &abac.GetResourceResponse{
		Id:          resource.Id,
		OwnerId:     resource.OwnerId,
		PolicyId:    resource.PolicyId,
		Name:        resource.Name,
		Description: resource.Description,
		Updated:     timestamppb.New(resource.Updated),
		Created:     timestamppb.New(resource.Created),
		Deleted:     timestamppb.New(resource.Deleted),
	}
}
