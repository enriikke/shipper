package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/enriikke/shipper/consignment-service/proto/consignment"
	vesselProto "github.com/enriikke/shipper/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

const defaultHost = "localhost:27017"

func main() {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Panicf("could not connect to datastore with host %s - %v", host, err)
	}

	defer session.Close()

	srv := micro.NewService(
		// This name must match the package name given in your protobuf def
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	// Init will parse the command line flags
	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
