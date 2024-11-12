package adapter

import (
	"github.com/charanck/ABAC/internal/model"
	abac "github.com/charanck/ABAC/protobuf/generated"
)

func CreateAttributeRequestToModel(request *abac.CreateAttributeRequest) model.Attribute {
	attribute := model.Attribute{
		Name: request.Name,
		Type: request.Type,
	}
	switch request.Type {
	case "string":
		attribute.StringValue = request.StringValue
	case "integer":
		attribute.IntegerValue = int(request.IntegerValue)
	case "float":
		attribute.FloatValue = request.FloatValue
	case "bool":
		attribute.BoolValue = request.BoolValue
	case "date":
		attribute.DateValue = request.DateValue.AsTime()
	}
	return attribute
}
