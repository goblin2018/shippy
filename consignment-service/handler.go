package main

import (
	"context"
	"errors"

	pb "github.com/goblin2018/shippy/consignment-service/proto/consignment"

	vesselPb "github.com/goblin2018/shippy/vessel-service/proto/vessel"
)

type handler struct {
	repository
	vesselClient vesselPb.VesselService
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	s.vesselClient.Create(ctx, &vesselPb.Vessel{
		Id:        "ok",
		Capacity:  9999999,
		MaxWeight: 9999999,
		Name:      "good",
		Avaiable:  true,
		OwnerId:   "ok2",
	})
	vesselRes, err := s.vesselClient.FindAvailable(ctx, &vesselPb.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	if vesselRes == nil {
		return errors.New("error fetching vessel, returned nil")
	}

	if err != nil {
		return err
	}

	req.VesselId = vesselRes.Vessel.Id
	if err = s.repository.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}
	res.Created = true
	res.Consignment = req
	return nil
}

func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repository.GetAll(ctx)
	if err != nil {
		return err
	}

	res.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}
