package province

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-meta-data/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetOne func
func (r Mongo) GetByName(name string) (results []model.ProvinceV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	filter := bson.M{"NameExtension": primitive.Regex{Pattern: name, Options: "im"}}
	var cursor *mongo.Cursor
	cursor, err = collection.Find(context.Background(), filter)
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
