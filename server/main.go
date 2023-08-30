package main

import (
	"context"
	"fmt"
	"net"

	"grpcmodel/config"
	"grpcmodel/constants"
	"grpcmodel/controller"
	c "grpcmodel/customer"
	"grpcmodel/services"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func intiDatabase(client *mongo.Client){
	customerCollection:=config.GetCollection(client,"BankDatabase","Customer")
	controller.CustomerService=services.InitCustomerService(customerCollection,context.Background())
}

func main(){
	mongoclient,err:=config.ConnectDatabase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	intiDatabase(mongoclient)
	lis,err:=net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s:=grpc.NewServer()
	c.RegisterCustomerServiceServer(s,&controller.RPCServer{})

	fmt.Println("sever listening on",constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}