package bankv2

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SwitchStatusRequest ..
type SwitchStatusRequest struct {
	BankID    int `bson:"_id"`
	NewStatus int `bson:"status"`

	UpdatedIP       string     `bson:"updated_ip,omitempty"`
	UpdatedEmployee int        `bson:"updated_employee"`
	UpdatedClient   int        `bson:"updated_client"`
	UpdatedSource   string     `bson:"updated_source"`
	UpdatedDate     *time.Time `bson:"updated_date"`
}

// SwitchStatus ..
func (r Repo) SwitchStatus(request SwitchStatusRequest) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	now := time.Now()
	request.UpdatedDate = &now

	filter := bson.M{"_id": request.BankID}
	update := bson.M{"$set": request}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	return
}
