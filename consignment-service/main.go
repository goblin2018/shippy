package main

import (
	"context"
	"log"
	"os"

	pb "github.com/goblin2018/shippy/consignment-service/proto/consignment"
	"github.com/micro/go-micro/v2"
)

const (
	port      = ":50051"
	defautlDB = "mongodb://datastore:27017"
)

func main() {

	service := micro.NewService(
		micro.Name("shippy.consignment.service"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")
	if len(uri) == 0 {
		uri = defautlDB
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}

	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shippy").Collection("vessels")
	repository := &MongoRepository{vesselCollection}

	h := &handler{repository}

	if err := pb.RegisterShippingServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
