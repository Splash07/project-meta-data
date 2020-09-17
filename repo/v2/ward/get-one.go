package ward

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-meta-data/constants"
	"gitlab.com/Splash07/project-meta-data/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetOne func
func (r Repo) GetOne(wardCode string) (result model.Ward, err error) {
	session, db := r.Session.GetCollection(r.Collection)
	err = db.Find(bson.M{"_id": wardCode}).One(&result)
	session.Close()
	return
}

// GetOne func
func (r Mongo) GetOne(wardCode string) (result model.WardV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	filter := bson.M{"IsDeleted": 0, "IsEnable": 1, "_id": wardCode}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = nil
	}
	return
}

// GetActiveOneV2 func
func (r Mongo) GetActiveOneV2(wardCode string) (result model.WardV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	filter := bson.M{
		"_id":    wardCode,
		"Status": constants.Status["ACTIVE"],
	}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = nil
	}
	return
}

// GetOneV2 ..
func (r Mongo) GetOneV2(wardCode string) (result model.WardV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	filter := bson.M{
		"_id":    wardCode,
		"Status": bson.M{"$ne": constants.Status["DELETE"]},
	}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = nil
	}
	return
}
