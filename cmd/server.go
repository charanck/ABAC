package cmd

import (
	"fmt"
	"log"
	"net"
	"net/http"

	grpchandler "github.com/charanck/ABAC/internal/handler/grpc"
	abac "github.com/charanck/ABAC/protobuf/generated"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

const grpcPort = 3001

func StartServer() {
	// Start http server
	go func() {
		e := echo.New()
		e.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, abac.HealthResponse{
				Message: "OK",
			})
		})

		if err := e.Start(":3000"); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Start grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	abac.RegisterHealthServer(grpcServer, grpchandler.HealthServer{})
	grpcServer.Serve(lis)

}
