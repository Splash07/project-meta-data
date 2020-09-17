package district

import (
	"context"
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
	"gitlab.ghn.vn/online/common/config"
	"gitlab.com/Splash07/project-meta-data/constants"
	"gitlab.com/Splash07/project-meta-data/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// UpdateDistrictRequest ..
type UpdateDistrictRequest struct {
	DistrictID            int
	DistrictName          string
	Priority              int
	DistrictEncode        string
	Type                  int
	SupportType           int
	NameExtensionToAdd    []string
	NameExtensionToRemove []string
	CanUpdateCOD          bool
	//
	UpdatedIP       string
	UpdatedEmployee int
	UpdatedClient   int
	UpdatedSource   string
}

// UpdateV2 ..
func (r Mongo) UpdateV2(request UpdateDistrictRequest) (err error) {
	ctx := context.Background()

	db := r.Session.ConClient.Database(config.GetConfig().MongoV2.Get("master-data").DBName)

	collection := db.Collection(r.Collection)

	filter := bson.M{"_id": request.DistrictID, "Status": bson.M{"$ne": constants.Status["DELETE"]}}

	updateData := make(bson.M)
	updateData["DistrictName"] = request.DistrictName
	updateData["DistrictEncode"] = request.DistrictEncode
	updateData["Priority"] = request.Priority
	updateData["Type"] = request.Type
	updateData["SupportType"] = request.SupportType
	updateData["CanUpdateCOD"] = request.CanUpdateCOD
	updateData["UpdatedIP"] = request.UpdatedIP
	updateData["UpdatedEmployee"] = request.UpdatedEmployee
	updateData["UpdatedSource"] = request.UpdatedSource
	updateData["UpdatedDate"] = time.Now()

	if request.UpdatedClient != 0 {
		updateData["UpdatedClient"] = request.UpdatedClient
	}

	// transaction
	err = db.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			return err
		}

		res, err := collection.UpdateOne(sessionContext, filter, bson.M{"$set": updateData})
		if err != nil || res.MatchedCount != 1 {
			sessionContext.AbortTransaction(sessionContext)
			if res.MatchedCount != 1 {
				return errors.New("Khong tim thay quan can duoc update")
			}
			return err
		}

		if len(request.NameExtensionToAdd) > 0 {
			addNameExtUpdate := bson.M{}
			addNameExtUpdate["$addToSet"] = bson.M{"NameExtension": bson.M{"$each": request.NameExtensionToAdd}}

			_, err = collection.UpdateOne(sessionContext, filter, addNameExtUpdate)
			if err != nil || res.MatchedCount != 1 {
				sessionContext.AbortTransaction(sessionContext)
				if res.MatchedCount != 1 {
					return errors.New("Khong tim thay quan can duoc update")
				}
				return err
			}
		}
		if len(request.NameExtensionToRemove) > 0 {
			removeNameExtUpdate := bson.M{}
			removeNameExtUpdate["$pull"] = bson.M{"NameExtension": bson.M{"$in": request.NameExtensionToRemove}}

			res, err = collection.UpdateOne(sessionContext, filter, removeNameExtUpdate)
			if err != nil || res.MatchedCount != 1 {
				sessionContext.AbortTransaction(sessionContext)
				if res.MatchedCount != 1 {
					return errors.New("Khong tim thay quan can duoc update")
				}
				return err
			}
		}

		sessionContext.CommitTransaction(sessionContext)
		return nil
	})
	return
}

// UpdateDocumentV2 ..
func (r Mongo) UpdateDocumentV2(district model.DistrictV2) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	now := time.Now()
	district.UpdatedDate = &now

	filter := bson.M{"_id": district.DistrictID}
	update := bson.M{"$set": district}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return
	}
	if res.MatchedCount != 1 {
		err = errors.New("Khong tim thay quan can duoc cap nhat")
		return
	}
	return
}
