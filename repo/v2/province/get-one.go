package province

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-meta-data/constants"
	"gitlab.com/Splash07/project-meta-data/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetOne func
// Might be removed
func (r Mongo) GetOne(provinceID int) (result model.ProvinceV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	filter := bson.M{"IsDeleted": 0, "IsEnable": 1, "_id": provinceID}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = nil
	}
	return
}

// GetActiveOneV2 func
func (r Mongo) GetActiveOneV2(provinceID int) (result model.ProvinceV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	filter := bson.M{
		"_id":    provinceID,
		"Status": constants.Status["ACTIVE"],
	}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = nil
	}
	return
}

// GetOneV2 ..
func (r Mongo) GetOneV2(provinceID int) (result model.ProvinceV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	filter := bson.M{
		"_id":    provinceID,
		"Status": bson.M{"$ne": constants.Status["DELETE"]},
	}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = nil
	}
	return
}
