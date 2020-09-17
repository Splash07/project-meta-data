package district

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-meta-data/constants"
	"gitlab.com/Splash07/project-meta-data/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAll func
func (r Mongo) GetAll(ProvinceID int) (results []model.DistrictV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	var filter = bson.M{"IsDeleted": 0, "IsRepresentative": 0}
	if ProvinceID > 0 {
		filter["ProvinceID"] = ProvinceID
	}
	var cursor *mongo.Cursor
	cursor, err = collection.Find(context.Background(), filter)
	if err == nil && cursor != nil {
		for cursor.Next(context.TODO()) {

			// create a value into which the single document can be decoded
			result := model.DistrictV2{}
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

// GetAllV2 func ..
func (r Mongo) GetAllV2(provinceID int) (districts []model.DistrictV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	filter := bson.M{}
	if provinceID > 0 {
		filter["ProvinceID"] = provinceID
	}
	filter["Status"] = bson.M{"$ne": constants.Status["DELETE"]}

	opts := options.Find()
	opts.SetSort(bson.M{"_id": -1})

	cur, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		district := model.DistrictV2{}
		err = cur.Decode(&district)
		if err != nil {
			return
		}
		districts = append(districts, district)
	}
	return
}
