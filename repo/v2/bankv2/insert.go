package bankv2

import (
	"context"
	"time"

	"gitlab.com/Splash07/project-meta-data/model"
)

// Insert func
func (r Repo) Insert(data model.BankV2) (bankID int, err error) {

	collection := r.Session.GetCollectionV2(r.Collection)

	// get ID for client
	data.ID = r.Session.NextID(r.Collection)

	now := time.Now()
	data.UpdatedDate = &now
	data.CreatedDate = &now

	//
	data.UpdatedIP = data.CreatedIP
	data.UpdatedClient = data.CreatedClient
	data.UpdatedEmployee = data.CreatedEmployee
	data.UpdatedSource = data.CreatedSource

	_, err = collection.InsertOne(context.Background(), data)
	if err != nil {
		return
	}

	bankID = data.ID
	return
}
