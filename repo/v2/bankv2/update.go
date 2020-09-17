package bankv2

import (
	"context"
	"errors"
	"time"

	"gitlab.com/Splash07/project-meta-data/model"
	"go.mongodb.org/mongo-driver/bson"
)

// Update ..
func (r Repo) Update(request model.BankV2) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	now := time.Now()
	request.UpdatedDate = &now

	filter := bson.M{"_id": request.ID}
	update := bson.M{"$set": request}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return
	}
	if res.MatchedCount != 1 {
		err = errors.New("Khong tim thay ngan hang can duoc cap nhat")
		return
	}
	return
}
