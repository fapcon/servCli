package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"servCliGrantex/client/pb"
	"time"
)

func main() {
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewMoneyServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req, err := client.GetMoneyValues(ctx, &pb.MoneyRequest{Input: "btcusd"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(req.GetOutput())
}
