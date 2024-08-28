package adapter

import (
	api "github.com/charanck/ABAC/internal/handler/http/generated"
	"github.com/charanck/ABAC/internal/model"
)

func ModelToList200JSONResponse(resources []model.Resource) api.List200JSONResponse {
	var response []api.Resource
	for _, resource := range resources {
		response = append(response, api.Resource{
			Id:          &resource.Id,
			Name:        &resource.Name,
			OwnerId:     &resource.OwnerId,
			PolicyId:    &resource.PolicyId,
			Updated:     &resource.Updated,
			Created:     &resource.Created,
			Deleted:     &resource.Deleted,
			Description: &resource.Description,
		})
	}
	return response
}

func CreateRequestObjectToModel(request *api.Resource) model.Resource {
	return model.Resource{
		Name:        *request.Name,
		OwnerId:     *request.OwnerId,
		PolicyId:    *request.PolicyId,
		Description: *request.Description,
	}
}

func ModelToGetByIdResponseObject(resource model.Resource) api.GetById200JSONResponse {
	return api.GetById200JSONResponse{
		Id:          &resource.Id,
		Created:     &resource.Created,
		Updated:     &resource.Updated,
		Deleted:     &resource.Deleted,
		Description: &resource.Description,
		Name:        &resource.Name,
		OwnerId:     &resource.OwnerId,
		PolicyId:    &resource.PolicyId,
	}
}

func UpdateByIdRequestObjectToModel(request api.UpdateByIdRequestObject) model.Resource {
	return model.Resource{
		Id:          request.ResourceId,
		Name:        *request.Body.Data.Name,
		OwnerId:     *request.Body.Data.OwnerId,
		PolicyId:    *request.Body.Data.PolicyId,
		Description: *request.Body.Data.Description,
	}
}
