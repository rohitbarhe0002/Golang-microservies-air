package main

import (
	"log"
	"net/http"
	common "rohitbarche2000/common"
pb "rohitbarche2000/common/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddr = "localhost:3000"
)

func main() {
	conn,err := grpc.Dial(orderServiceAddr,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil {
		log.Fatalf("faild to dial server:%v",err)
	}
	defer conn.Close()

	log.Println("dialing order service at",orderServiceAddr)
	c := pb.NewOrderServiceClient(conn)
	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)
	log.Printf("Starting HTTP server at %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("failed to start http server")
	}

}
