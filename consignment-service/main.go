package main

import (
	"context"
	"log"
	"sync"

	pb "github.com/goblin2018/shippy/consignment-service/proto/consignment"
	"github.com/micro/go-micro/v2"
)

const (
	port = ":50051"
)



type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	log.Println("create consignment: ", consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type consignmentService struct {
	repo repository
}

func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	log.Println("CreateConsignment: ", req)
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = consignment
	log.Println("res is ", res)

	return nil
}

func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	res = &pb.Response{}
	res.Consignments = s.repo.GetAll()
	return nil
}

func main() {
	repo := &Repository{}

	service := micro.NewService(
		micro.Name("shippy.consignment.service"),
	)

	service.Init()
	if err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
