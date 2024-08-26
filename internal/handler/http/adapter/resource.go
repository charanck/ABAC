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
