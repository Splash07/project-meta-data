package district

import (
	"context"
	"time"

	"gitlab.com/Splash07/project-meta-data/model"
)

// Insert func
func (r Mongo) Insert(data *model.DistrictV2) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)
	// override data
	now := time.Now()
	data.CreatedAt = now.String()
	data.UpdatedAt = now.String()
	//
	_, err = collection.InsertOne(context.Background(), data)
	return
}
