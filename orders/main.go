package main

import (
	"context"

	"log"
	"net"
	"rohitbarche2000/common"
pb "rohitbarche2000/common/api"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:3000")
)

func main() {
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to start TCP listener on %s: %v", grpcAddr, err)
	}
	defer func() {
		if cerr := l.Close(); cerr != nil {
			log.Printf("Error closing listener: %v", cerr)
		}
	}()
	log.Println("gRPC server listening on", grpcAddr)

	
	store := NewStore()
	if store == nil {
		log.Fatal("Failed to initialize store")
	}
	svc := NewService(store)
	if svc == nil {
		log.Fatal("Failed to initialize service")
	}

	// Register gRPC handlers
 NewGRPCHandler(grpcServer,svc)

	// Call CreateOrder (if necessary)
	err = svc.CreateOrder(context.Background())
	if err != nil {
		log.Fatalf("Failed to create order: %v", err)
	}

	// Start serving gRPC requests
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}


func (s *service) validateOrder(stx context.Context,p *pb.CreateOrderRequest)error{
	if len(p.Items) == 0 {
return common.ErrNoItems
	}
	mergedItems := mergeItemsQuantities(p.Items)
}


func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity{
merged := make([]*pb.ItemsWithQuantity,0)
for _,item :=range items{
merged = append(merged, item)
}

return merged
}