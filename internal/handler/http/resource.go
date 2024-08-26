package httphandler

import (
	"context"
	"errors"

	"github.com/charanck/ABAC/internal/handler/http/adapter"
	api "github.com/charanck/ABAC/internal/handler/http/generated"
	"github.com/charanck/ABAC/internal/service"
	"github.com/charanck/ABAC/internal/util"
)

type Resource struct {
	resourceService service.Resource
}

func NewResource(resourceService *service.Resource) Resource {
	return Resource{
		resourceService: *resourceService,
	}
}

func (r *Resource) List(ctx context.Context, request api.ListRequestObject) (api.ListResponseObject, error) {
	resources, err := r.resourceService.List()
	if err != nil {
		var apiError util.ApiError
		if errors.As(err, &apiError) {
			return api.ListdefaultJSONResponse{
				Body: api.Error{
					Code:    int32(apiError.HTTPErrorCode),
					Message: apiError.Error(),
				},
				StatusCode: apiError.HTTPErrorCode,
			}, nil
		} else {
			return api.ListdefaultJSONResponse{
				Body: api.Error{
					Code:    500,
					Message: "internal server error",
				},
				StatusCode: 500,
			}, nil
		}
	}
	return adapter.ModelToList200JSONResponse(resources), nil
}

func (r *Resource) Create(ctx context.Context, request api.CreateRequestObject) (api.CreateResponseObject, error) {
	return api.Create201Response{}, nil
}

func (r *Resource) DeleteById(ctx context.Context, request api.DeleteByIdRequestObject) (api.DeleteByIdResponseObject, error) {
	return api.DeleteById200Response{}, nil
}

func (r *Resource) GetById(ctx context.Context, request api.GetByIdRequestObject) (api.GetByIdResponseObject, error) {
	return api.GetById200JSONResponse{}, nil
}

func (r *Resource) UpdateById(ctx context.Context, request api.UpdateByIdRequestObject) (api.UpdateByIdResponseObject, error) {
	return api.UpdateById200Response{}, nil
}
