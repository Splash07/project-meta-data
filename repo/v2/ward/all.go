package ward

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-meta-data/constants"
	"gitlab.com/Splash07/project-meta-data/model"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// All func
func (r Repo) All(districtID int) (result []model.Ward, err error) {
	session, db := r.Session.GetCollection(r.Collection)
	condition := bson.M{}
	if districtID > 0 {
		condition["DistrictID"] = districtID
	}
	err = db.Find(condition).All(&result)
	session.Close()
	return
}

// GetAllV2 func ..
func (r Mongo) GetAllV2(districtID int) (wards []model.WardV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	filter := bson.M{}
	if districtID > 0 {
		filter["DistrictID"] = districtID
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
		ward := model.WardV2{}
		err = cur.Decode(&ward)
		if err != nil {
			return
		}
		wards = append(wards, ward)
	}
	return
}
