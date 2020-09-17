package province

import (
	"context"
	"time"

	"gitlab.com/Splash07/project-meta-data/model"
)

// InsertV2 ..
func (r Mongo) InsertV2(province model.ProvinceV2) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	province.ProvinceID = r.Session.NextID(r.Collection)
	//
	province.UpdatedIP = province.CreatedIP
	province.UpdatedClient = province.CreatedClient
	province.UpdatedEmployee = province.CreatedEmployee
	province.UpdatedSource = province.CreatedSource

	now := time.Now()
	province.CreatedDate = &now
	province.UpdatedDate = &now

	_, err = collection.InsertOne(context.Background(), province)
	return
}
