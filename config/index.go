package config

import (
	"context"

	"log"

	"time"
     "grpcmodel/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDatabase()(*mongo.Client,error){
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	//it connects using the uri given
	mongoConnection:=options.Client().ApplyURI(constants.ConnectionString)
	//connect and checking whether connected 
	mongoclient,err:=mongo.Connect(ctx,mongoConnection)
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err
	}
	if err:=mongoclient.Ping(ctx,readpref.Primary());err!=nil{
		return nil,err
	}
	return mongoclient,nil
}

func GetCollection(client *mongo.Client,dbname string,collectionName string)*mongo.Collection{
	// colletions contains the single collection which we use
	
	collection:=client.Database(dbname).Collection(collectionName)
	return collection
}