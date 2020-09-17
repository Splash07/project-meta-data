package bankv2

import (
	"context"
	"errors"
	"time"

	"gitlab.com/Splash07/project-meta-data/constants"
	"go.mongodb.org/mongo-driver/bson"
)

// RemoveBankRequest ..
type RemoveBankRequest struct {
	BankID int `bson:"_id"`

	UpdatedIP       string     `bson:"updated_ip"`
	UpdatedEmployee int        `bson:"updated_employee"`
	UpdatedSource   string     `bson:"updated_source"`
	UpdatedDate     *time.Time `bson:"updated_date"`
}

// RemoveBank ..
func (r Repo) RemoveBank(removalRequest RemoveBankRequest) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	filter := bson.M{"_id": removalRequest.BankID}

	now := time.Now()
	removalRequest.UpdatedDate = &now
	update := bson.M{"$set": removalRequest}
	update["$set"] = bson.M{"status": constants.Status["DELETE"]}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return
	}
	if res.MatchedCount != 1 {
		err = errors.New("Khong tim thay ngan hang can xoa")
		return
	}
	return
}
