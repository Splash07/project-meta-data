package script

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// WardStatusRunner ..
var WardStatusRunner wardStatusRunner

type wardStatusRunner struct{}

func (r wardStatusRunner) failOnError(err error) {
	if err != nil {
		log.Println(err)
		time.Sleep(1 * time.Second)
		r.Run()
	}
}

// Run ..
func (r *wardStatusRunner) Run() (err error) {
	// init Mongo
	cfg.MongoV2.Get("master-data").Init()

	collection := cfg.MongoV2.Get("master-data").GetCollectionV2("Ward")

	filter := bson.M{}
	update := bson.M{"$set": bson.M{"Status": 1}}

	_, err = collection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		r.failOnError(err)
	}
	return
}
