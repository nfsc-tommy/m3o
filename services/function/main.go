package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/m3o/m3o/services/function/handler"
	pb "github.com/m3o/m3o/services/function/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("function"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterFunctionHandler(srv.Server(), handler.NewFunction())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
