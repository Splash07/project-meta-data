package district

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-meta-data/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetOneByCode func
func (r Mongo) GetOneByCode(districtCode string) (result model.DistrictV2, err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	filter := bson.M{"IsDeleted": 0, "IsEnable": 1, "Code": districtCode}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		err = nil
	}
	return
}
