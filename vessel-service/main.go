package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/enriikke/shipper/vessel-service/proto/vessel"
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
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
