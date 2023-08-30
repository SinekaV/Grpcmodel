package services

import (
	"context"
	"grpcmodel/interfaces"
	"grpcmodel/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}


func InitCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

func (c * CustomerService)CreateCustomer(Info *models.CustomerRequest)(*models.CustomerResponse,error){
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"customer_id": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := c.CustomerCollection.Indexes().CreateOne(c.ctx, indexModel)
	if err != nil {
		log.Fatal(err)
	}
	Info.IsActive = true
	Info.CreatedAt = time.Now()
	Info.UpdatedAt = Info.CreatedAt

	res,err:=c.CustomerCollection.InsertOne(c.ctx,&Info)
	if err!=nil{
		return nil,err
	}
	var newUser *models.CustomerResponse
	query:=bson.M{"_id":res.InsertedID}

	err = c.CustomerCollection.FindOne(c.ctx,query).Decode(&newUser)
	if err != nil{
		return nil,err
	}
	return newUser,nil
}