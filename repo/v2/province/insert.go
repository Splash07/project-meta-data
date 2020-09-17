package province

import (
	"context"
	"time"

	"gitlab.com/Splash07/project-meta-data/model"
)

// Insert ..
// Might be removed
func (r Mongo) Insert(data *model.ProvinceV2) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	// override data
	now := time.Now()
	data.CreatedAt = now.String()
	data.UpdatedAt = now.String()
	//
	_, err = collection.InsertOne(context.Background(), data)
	return
}
