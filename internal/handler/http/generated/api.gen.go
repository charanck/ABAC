// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Pet defines model for Pet.
type Pet struct {
	Id   int64   `json:"id"`
	Name string  `json:"name"`
	Tag  *string `json:"tag,omitempty"`
}

// Pets defines model for Pets.
type Pets = []Pet

// ListPetsParams defines parameters for ListPets.
type ListPetsParams struct {
	// Limit How many items to return at one time (max 100)
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

// CreatePetsJSONRequestBody defines body for CreatePets for application/json ContentType.
type CreatePetsJSONRequestBody = Pet

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List all pets
	// (GET /pets)
	ListPets(ctx echo.Context, params ListPetsParams) error
	// Create a pet
	// (POST /pets)
	CreatePets(ctx echo.Context) error
	// Info for a specific pet
	// (GET /pets/{petId})
	ShowPetById(ctx echo.Context, petId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListPets converts echo context to params.
func (w *ServerInterfaceWrapper) ListPets(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListPetsParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ListPets(ctx, params)
	return err
}

// CreatePets converts echo context to params.
func (w *ServerInterfaceWrapper) CreatePets(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreatePets(ctx)
	return err
}

// ShowPetById converts echo context to params.
func (w *ServerInterfaceWrapper) ShowPetById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId string

	err = runtime.BindStyledParameterWithOptions("simple", "petId", ctx.Param("petId"), &petId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ShowPetById(ctx, petId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/pets", wrapper.ListPets)
	router.POST(baseURL+"/pets", wrapper.CreatePets)
	router.GET(baseURL+"/pets/:petId", wrapper.ShowPetById)

}

type ListPetsRequestObject struct {
	Params ListPetsParams
}

type ListPetsResponseObject interface {
	VisitListPetsResponse(w http.ResponseWriter) error
}

type ListPets200ResponseHeaders struct {
	XNext string
}

type ListPets200JSONResponse struct {
	Body    Pets
	Headers ListPets200ResponseHeaders
}

func (response ListPets200JSONResponse) VisitListPetsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-next", fmt.Sprint(response.Headers.XNext))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type ListPetsdefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response ListPetsdefaultJSONResponse) VisitListPetsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type CreatePetsRequestObject struct {
	Body *CreatePetsJSONRequestBody
}

type CreatePetsResponseObject interface {
	VisitCreatePetsResponse(w http.ResponseWriter) error
}

type CreatePets201Response struct {
}

func (response CreatePets201Response) VisitCreatePetsResponse(w http.ResponseWriter) error {
	w.WriteHeader(201)
	return nil
}

type CreatePetsdefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response CreatePetsdefaultJSONResponse) VisitCreatePetsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type ShowPetByIdRequestObject struct {
	PetId string `json:"petId"`
}

type ShowPetByIdResponseObject interface {
	VisitShowPetByIdResponse(w http.ResponseWriter) error
}

type ShowPetById200JSONResponse Pet

func (response ShowPetById200JSONResponse) VisitShowPetByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ShowPetByIddefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response ShowPetByIddefaultJSONResponse) VisitShowPetByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// List all pets
	// (GET /pets)
	ListPets(ctx context.Context, request ListPetsRequestObject) (ListPetsResponseObject, error)
	// Create a pet
	// (POST /pets)
	CreatePets(ctx context.Context, request CreatePetsRequestObject) (CreatePetsResponseObject, error)
	// Info for a specific pet
	// (GET /pets/{petId})
	ShowPetById(ctx context.Context, request ShowPetByIdRequestObject) (ShowPetByIdResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// ListPets operation middleware
func (sh *strictHandler) ListPets(ctx echo.Context, params ListPetsParams) error {
	var request ListPetsRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ListPets(ctx.Request().Context(), request.(ListPetsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListPets")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(ListPetsResponseObject); ok {
		return validResponse.VisitListPetsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// CreatePets operation middleware
func (sh *strictHandler) CreatePets(ctx echo.Context) error {
	var request CreatePetsRequestObject

	var body CreatePetsJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreatePets(ctx.Request().Context(), request.(CreatePetsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreatePets")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(CreatePetsResponseObject); ok {
		return validResponse.VisitCreatePetsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// ShowPetById operation middleware
func (sh *strictHandler) ShowPetById(ctx echo.Context, petId string) error {
	var request ShowPetByIdRequestObject

	request.PetId = petId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.ShowPetById(ctx.Request().Context(), request.(ShowPetByIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ShowPetById")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(ShowPetByIdResponseObject); ok {
		return validResponse.VisitShowPetByIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
