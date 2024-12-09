package main

import (
	"context"
	"log"
	pb "rohitbarche2000/common/api"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrderSrvice
}

func NewGRPCHandler(grpcServer *grpc.Server,service OrderSrvice) {
	handler := &grpcHandler{
		service: service,
	}
	pb.RegisterOrderServiceServer(grpcServer, handler)

}
func (h *grpcHandler) CreateOrder(ctx context.Context, p  *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println(" New order Recived!",p)
	o := &pb.Order{
		ID: "42",
	}
	return o, nil

}
