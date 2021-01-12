package main

import (
	"context"

	pb "github.com/goblin2018/shippy/consignment-service/proto/consignment"

	vesselPb "github.com/goblin2018/shippy/vessel-service/proto/vessel"
)

type handler struct {
	repository
	vesselClient vesselPb.VesselService
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	vesselRes, err := s.vesselClient.FindAvailable(ctx, &vesselPb.S)
}
