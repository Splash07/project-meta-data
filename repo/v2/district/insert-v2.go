package district

import (
	"context"
	"time"

	"gitlab.com/Splash07/project-meta-data/model"
)

// InsertV2 ..
func (r Mongo) InsertV2(district model.DistrictV2) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	district.DistrictID = r.Session.NextID(r.Collection)
	//
	district.UpdatedIP = district.CreatedIP
	district.UpdatedClient = district.CreatedClient
	district.UpdatedEmployee = district.CreatedEmployee
	district.UpdatedSource = district.CreatedSource

	now := time.Now()
	district.CreatedDate = &now
	district.UpdatedDate = &now

	_, err = collection.InsertOne(context.Background(), district)
	return
}
