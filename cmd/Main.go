package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"log"
	"net"
	"testTask/internal/app"
	spec "testTask/internal/pkg/models"
	proto "testTask/proto/gen/proto"
)

const restServerAddress = "8080"
const grpcServerAddress = "8090"

func main() {

	srv, err := net.Listen("tcp", fmt.Sprintf(`:%s`, grpcServerAddress))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTransferBoxApiServer(grpcServer, new(app.ProtoHandler))

	e := echo.New()
	spec.RegisterHandlers(e, app.RestHandler{})

	fmt.Println(fmt.Sprintf(`Server running in address : %s`, grpcServerAddress))

	// стартуем
	go func() {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", restServerAddress)))
	}()
	if err = grpcServer.Serve(srv); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
