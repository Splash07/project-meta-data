package province

import (
	"context"
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
)

// AppendNameExtRequest ..
type AppendNameExtRequest struct {
	ProvinceID int
	NewNameExt []string

	UpdatedIP       string
	UpdatedEmployee int
	UpdatedClient   int
	UpdatedSource   string
	UpdatedDate     time.Time
}

// AppendToNameExt func
func (r Mongo) AppendToNameExt(request AppendNameExtRequest) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	request.UpdatedDate = time.Now()

	filter := bson.M{"_id": request.ProvinceID}

	update := make(bson.M)
	update["$addToSet"] = bson.M{"NameExtension": bson.M{"$each": request.NewNameExt}}
	update["$set"] = bson.M{
		"UpdatedIP":     request.UpdatedIP,
		"UpdatedSource": request.UpdatedSource,
		"UpdatedDate":   request.UpdatedDate,
	}

	if request.UpdatedEmployee != 0 {
		update["$set"] = bson.M{"UpdatedEmployee": request.UpdatedEmployee}
	}
	if request.UpdatedClient != 0 {
		update["$set"] = bson.M{"UpdatedCLient": request.UpdatedClient}
	}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return
	}
	if res.MatchedCount != 1 {
		err = errors.New("Không tìm thấy Tỉnh cần sửa đổi. Hãy đảm bảo rằng province_id được truyền vào là chính xác")
		return
	}
	return
}
