package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"servCliGrantex/grpcServ/config"
	"servCliGrantex/grpcServ/models"
	pb2 "servCliGrantex/grpcServ/pb"
)

type Server struct {
	pb2.UnimplementedMoneyServiceServer
}

func (s *Server) GetMoneyValues(ctx context.Context, request *pb2.MoneyRequest) (*pb2.MoneyResponse, error) {
	url := "https://garantex.org/api/v2/depth?market=" + request.GetInput()
	get, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer get.Body.Close()
	res := models.CurrencyInfo{}
	err = json.NewDecoder(get.Body).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return &pb2.MoneyResponse{Output: res.Timestamp}, err
}

func main() {
	c := config.New()
	port := fmt.Sprintf(":%d", c.Server.Port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	service := Server{}
	pb2.RegisterMoneyServiceServer(server, &service)
	server.Serve(lis)
}
