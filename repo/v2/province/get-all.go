package province

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-meta-data/constants"
	"gitlab.com/Splash07/project-meta-data/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAll func
func (r Mongo) GetAll() (results []model.ProvinceV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	var cursor *mongo.Cursor
	cursor, err = collection.Find(context.Background(), bson.M{"IsDeleted": 0})
	if err == nil && cursor != nil {
		for cursor.Next(context.TODO()) {

			// create a value into which the single document can be decoded
			result := model.ProvinceV2{}
			err = cursor.Decode(&result)
			if err != nil {
				return nil, err
			}

			results = append(results, result)
		}

		// Close the cursor once finished
		_ = cursor.Close(context.TODO())
	}
	if err == mongo.ErrNoDocuments {
		err = nil
	}
	return
}

// GetAllV2 func
func (r Mongo) GetAllV2() (results []model.ProvinceV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	filter := bson.M{"Status": bson.M{"$ne": constants.Status["DELETE"]}}

	opts := options.Find()
	opts.SetSort(bson.M{"_id": -1})

	cur, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		result := model.ProvinceV2{}
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}
	return
}
