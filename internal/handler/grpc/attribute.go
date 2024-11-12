package grpchandler

import (
	"github.com/charanck/ABAC/internal/service"
	abac "github.com/charanck/ABAC/protobuf/generated"
)

type Attribute struct {
	abac.UnimplementedAttributeServer
	Service service.Attribute
}

func NewAttribute(service service.Attribute) Attribute {
	return Attribute{
		Service: service,
	}
}
