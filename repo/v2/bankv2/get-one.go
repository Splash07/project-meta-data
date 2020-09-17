package bankv2

import (
	"context"

	"gitlab.com/Splash07/project-meta-data/constants"
	"gitlab.com/Splash07/project-meta-data/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetByID ..
// Might be removed
func (r Repo) GetByID(id int) (result model.BankV2, err error) {

	collection := r.Session.GetCollectionV2(r.Collection)

	filter := bson.M{}
	filter["_id"] = id

	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = nil
	}

	return
}

// GetByIDV2 ..
func (r Repo) GetByIDV2(id int) (result model.BankV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	filter := bson.M{
		"_id":    id,
		"status": constants.Status["ACTIVE"],
	}

	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = nil
	}
	return
}
