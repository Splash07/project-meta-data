package province

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

// UpdateProvinceRequest ..
type UpdateProvinceRequest struct {
	ProvinceID            int
	ProvinceName          string
	Priority              int
	Code                  string
	RegionID              int
	ProvinceEncode        string
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
func (r Mongo) UpdateV2(request UpdateProvinceRequest) (err error) {
	ctx := context.Background()

	db := r.Session.ConClient.Database(config.GetConfig().MongoV2.Get("master-data").DBName)
	// defer db.Client().Disconnect(ctx)

	collection := db.Collection(r.Collection)

	filter := bson.M{"_id": request.ProvinceID, "Status": bson.M{"$ne": constants.Status["DELETE"]}}

	updateData := make(bson.M)
	updateData["ProvinceName"] = request.ProvinceName
	updateData["ProvinceEncode"] = request.ProvinceEncode
	updateData["Priority"] = request.Priority
	updateData["RegionID"] = request.RegionID
	updateData["CanUpdateCOD"] = request.CanUpdateCOD
	updateData["UpdatedIP"] = request.UpdatedIP
	updateData["UpdatedEmployee"] = request.UpdatedEmployee
	updateData["UpdatedSource"] = request.UpdatedSource
	updateData["UpdatedDate"] = time.Now()

	if request.Code != "" {
		updateData["Code"] = request.Code
	}
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
				return errors.New("Khong tim thay Province can duoc update")
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
					return errors.New("Khong tim thay Province can duoc update")
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
					return errors.New("Khong tim thay Province can duoc update")
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
func (r Mongo) UpdateDocumentV2(province model.ProvinceV2) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	now := time.Now()
	province.UpdatedDate = &now

	filter := bson.M{"_id": province.ProvinceID}
	update := bson.M{"$set": province}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return
	}
	if res.MatchedCount != 1 {
		err = errors.New("Khong tim thay tinh can duoc cap nhat")
		return
	}
	return
}
