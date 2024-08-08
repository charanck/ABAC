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

func ModelToListResourceResponse(resources []model.Resource) *abac.ListResourceResponse {
	response := abac.ListResourceResponse{}
	for _, resource := range resources {
		response.Data = append(response.Data, ModelToGetResourceResponse(resource))
	}
	return &response
}

func UpdateResourceRequestToModel(request *abac.UpdateResourceRequest) model.Resource {
	return model.Resource{
		Id:          request.Data.Id,
		Name:        request.Data.Name,
		OwnerId:     request.Data.OwnerId,
		PolicyId:    request.Data.PolicyId,
		Description: request.Data.Description,
	}
}
