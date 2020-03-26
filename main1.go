package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)
func main()  {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://gomongodb:kirmali15%40#MK@cluster0-qiqg2.mongodb.net/test?retryWrites=true&w=majority",))

	if err !=nil{
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err !=nil{
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)
	quickStartDatabase := client.Database("gomongodb")
	podcastsCollection := quickStartDatabase.Collection("podcasts")
	episodesCollection := quickStartDatabase.Collection("episodes")
	podcastResult,err :=podcastsCollection.InsertOne(ctx,bson.D{
		{Key:"title",Value:"MongoDB"},
		{Key:"author",Value:"karan rathore"},
		{"tag",bson.A{"developement","programming","coding"}},
	})
	if err !=nil{
		log.Fatal(err)
	}

	fmt.Println(podcastResult.InsertedID)
	episodeResult,err :=episodesCollection.InsertMany(ctx,[]interface{}{
		bson.D{
			{"podcast",podcastResult.InsertedID},
			{"title","1"},
			{"description","2"},
			{"duration","3 "},

		},
		bson.D{
			{"podcast",podcastResult.InsertedID},
			{"title","11"},
			{"description","22"},
			{"duration","33"},

		},
	})
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(episodeResult.InsertedIDs)
}
