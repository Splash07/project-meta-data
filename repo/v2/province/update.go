package province

import (
	"context"
	"time"

	"github.com/globalsign/mgo/bson"
	"gitlab.com/Splash07/project-meta-data/model"
)

// Update func
func (r Mongo) Update(data *model.ProvinceV2) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	// override data
	data.UpdatedAt = time.Now().String()

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": data.ProvinceID}, bson.M{"$set": data})
	return
}
