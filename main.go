package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/kzmake/micro-helloworld/handler"
	"github.com/kzmake/micro-helloworld/subscriber"

	helloworld "github.com/kzmake/micro-helloworld/proto/helloworld"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("kzmake.example.service.helloworld"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	helloworld.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("kzmake.example.service.helloworld", service.Server(), new(subscriber.Helloworld))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
