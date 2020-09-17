package ward

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-meta-data/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetByName func
func (r Mongo) GetByDistrict(DistrictID int) (results []model.WardV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	var filter interface{}
	if DistrictID > 0 {
		filter = bson.M{"DistrictID" : DistrictID, "IsDeleted" : 0}
	} else {
		filter = bson.M{"IsDeleted" : 0}
	}
	var cursor *mongo.Cursor
	cursor, err = collection.Find(context.Background(), filter)
	if err == nil && cursor != nil {
		for cursor.Next(context.TODO()) {

			// create a value into which the single document can be decoded
			result := model.WardV2{}
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