package bankv2

import (
	"context"

	"gitlab.com/Splash07/project-meta-data/constants"
	"gitlab.com/Splash07/project-meta-data/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// All ..
func (r Repo) All() (results []model.BankV2, err error) {

	collection := r.Session.GetCollectionV2(r.Collection)

	// check filter
	filter := bson.M{"active": true}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		bank := model.BankV2{}
		err := cur.Decode(&bank)
		if err != nil {
			return nil, err
		}

		results = append(results, bank)
	}

	return
}

// GetAllV2 ..
func (r Repo) GetAllV2(status []int, offsetID, limit int) (results []model.BankV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	opts := options.Find()
	if limit < 0 || limit > 100 {
		limit = 100
	}
	opts.SetLimit(int64(limit))
	opts.SetSort(bson.M{"_id": -1})

	filter := bson.M{}
	if offsetID != 0 {
		filter["_id"] = bson.M{"$lt": offsetID}
	}

	if len(status) == 0 {
		status = []int{constants.Status["ACTIVE"], constants.Status["NOT_ACTIVE"]}
	}
	filter["status"] = bson.M{"$in": status}

	cur, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		bank := model.BankV2{}
		err = cur.Decode(&bank)
		if err != nil {
			return
		}
		results = append(results, bank)
	}
	return
}
