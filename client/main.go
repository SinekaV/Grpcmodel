package main

import (
	"context"
	"fmt"

	"log"

	"grpcmodel/constants"
	pb "grpcmodel/customer"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(constants.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)
	

	response, err := client.CreateCustomer(context.Background(), &pb.CustomerRequest{
		CustomerId: 321,
		Firstname:  "Sineka",
		Lastname:   "V",
		BankId:     12345,
		Balance:    100000,
	})
	if err != nil {
		log.Fatalf("Failed to call: %v", err)
	}

	fmt.Printf("CustomerID: %d\nCreatedTime:%v\n", response.CustomerId,response.CreatedAt)
}