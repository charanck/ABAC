package cmd

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	grpchandler "github.com/charanck/ABAC/internal/handler/grpc"
	httphandler "github.com/charanck/ABAC/internal/handler/http"
	api "github.com/charanck/ABAC/internal/handler/http/generated"
	"github.com/charanck/ABAC/internal/repository"
	"github.com/charanck/ABAC/internal/service"
	abac "github.com/charanck/ABAC/protobuf/generated"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

const grpcPort = 3001

func StartServer() {
	// Connect to DB
	db, err := sqlx.Open("sqlite3", "./db/db.sqlite")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	// Setup dependencies
	resourceRepository := repository.NewResource(db)
	resourceService := service.NewResource(&resourceRepository)
	resourceGRPCHandler := grpchandler.NewResource(resourceService)
	resourceHTTPHandler := httphandler.NewResource(&resourceService)

	// Start http server
	go func() {
		e := echo.New()
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
		}))
		var strictHandler api.ServerInterface = api.NewStrictHandler(&resourceHTTPHandler, nil)
		api.RegisterHandlers(e, strictHandler)
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

	// Register the grpc handlers
	abac.RegisterHealthServer(grpcServer, grpchandler.HealthServer{})
	abac.RegisterResourceServer(grpcServer, &resourceGRPCHandler)

	grpcServer.Serve(lis)

}
