package adapter

import (
	"github.com/charanck/ABAC/internal/model"
	abac "github.com/charanck/ABAC/protobuf/generated"
)

func ResourceGRPCToModel(request *abac.CreateResourceRequest) model.Resource {
	return model.Resource{
		Name:        request.Name,
		OwnerId:     request.OwnerId,
		PolicyId:    request.PolicyId,
		Description: request.Description,
	}
}
