package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/m3o/m3o/services/helloworld/handler"
	pb "github.com/m3o/m3o/services/helloworld/proto"
	"github.com/m3o/m3o/services/pkg/tracing"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("helloworld"),
	)

	// Register Handler
	pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld))

	traceCloser := tracing.SetupOpentracing("helloworld")
	defer traceCloser.Close()

	// Run the service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
