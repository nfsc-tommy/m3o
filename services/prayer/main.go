package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/m3o/m3o/services/prayer/handler"
	pb "github.com/m3o/m3o/services/prayer/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("prayer"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPrayerHandler(srv.Server(), handler.New(srv.Client()))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
