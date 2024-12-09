package main

import "context"
import pb "rohitbarche2000/common/api"

type OrderSrvice interface {
	CreateOrder(context.Context) error
	ValidateOrder(context.Context,*pb.CreateOrderRequest)error
}


type OrderStore interface {
	Create(context.Context) error
}