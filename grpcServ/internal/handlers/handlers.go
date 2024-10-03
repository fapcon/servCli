package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	fmt.Println("successfully based")
	return &pb2.MoneyResponse{Output: res.Timestamp}, err
}
