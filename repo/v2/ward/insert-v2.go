package ward

import (
	"context"
	"strconv"
	"time"

	"gitlab.com/Splash07/project-meta-data/model"
)

// InsertV2 ..
func (r Mongo) InsertV2(ward model.WardV2) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	// Get Next ID
	ward.WardCode = strconv.Itoa(r.Session.NextID(r.Collection))
	//
	ward.UpdatedIP = ward.CreatedIP
	ward.UpdatedClient = ward.CreatedClient
	ward.UpdatedEmployee = ward.CreatedEmployee
	ward.UpdatedSource = ward.CreatedSource

	now := time.Now()
	ward.CreatedDate = &now
	ward.UpdatedDate = &now

	_, err = collection.InsertOne(context.Background(), ward)
	return
}
