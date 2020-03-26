package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)
func main()  {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://gomongodb:<password>@cluster0-qiqg2.mongodb.net/test?retryWrites=true&w=majority"))

	if err !=nil{
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err !=nil{
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	err = client.Ping(ctx,readpref.Primary())

	if err !=nil{
		log.Fatal(err)
	}

	database,err :=client.ListDatabaseNames(ctx,bson.M{})

	if err !=nil{
		log.Fatal(err)
	}

fmt.Println("connected to database.........",database)
}
