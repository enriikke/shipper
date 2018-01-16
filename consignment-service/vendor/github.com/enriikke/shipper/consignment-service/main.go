package main

import (
	"context"
	"fmt"

	pb "github.com/enriikke/shipper/consignment-service/proto/consignment"
	micro "github.com/micro/go-micro"
)

const port = ":50051"

type IRespository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository represents a dummy repository that simulates the use of a
// datastore of some kind. It will be replaced with a real implementation in the
// future.
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// service should implement all of the methods to satisfy the service defined in
// our protobuf definition.
type service struct {
	repo IRespository
}

// CreateConsignment implements the shipping service method
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response message we created in our protobuf def.
	res.Created = true
	res.Consignment = consignment

	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	res.Consignments = s.repo.GetAll()

	return nil
}

func main() {
	repo := &Repository{}

	srv := micro.NewService(
		// This name must match the package name given in your protobuf def
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags
	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
