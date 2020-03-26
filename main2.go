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

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://gomongodb:kirmali15%40#MK@cluster0-qiqg2.mongodb.net/test?retryWrites=true&w=majority", ))

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	quickStartDatabase := client.Database("gomongodb")

	podcastsCollection := quickStartDatabase.Collection("podcasts")
	episodesCollection := quickStartDatabase.Collection("episodes")

	cursor, err := episodesCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	//var episodes []bson.M
	//
	//if err = cursor.All(ctx, &episodes); err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, eposode := range episodes {
	//	fmt.Println(eposode["title"])
	//}

	defer cursor.Close(ctx)

	for cursor.Next(ctx){
		var episode bson.M
		if err = cursor.Decode(&episode);err != nil {
			log.Fatal(err)
		}
	}

	var podcast bson.M

	if err = podcastsCollection.FindOne(ctx,bson.M{}).Decode(&podcast); err != nil{
		log.Fatal(err)
	}
	fmt.Println(podcast)

	filterCursor,err :=episodesCollection.Find(ctx,bson.M{"duration":22})

	if err != nil{
		log.Fatal(err)
	}

	var episodesFiltered []bson.M

	if err = filterCursor.All(ctx,&episodesFiltered);err != nil{
		log.Fatal(err)
	}

	fmt.Println(episodesFiltered)

	opts := options.Find()

	opts.SetSort(bson.D{{"duration", -1}})

	sortCursor, err :=episodesCollection.Find(ctx,bson.D{
		{"duration",bson.D{
			{"$gt",24},
		}},
	},opts)

	if err != nil{
		log.Fatal(err)
	}

	var episodesSorted []bson.M

	if err = sortCursor.All(ctx,&episodesSorted); err != nil{
		log.Fatal(err)
	}

	fmt.Println(episodesSorted)
}
