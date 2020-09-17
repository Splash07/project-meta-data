package district

import (
	"context"
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
)

// AppendToNameExt func
func (r Mongo) AppendToNameExt(districtID int, updatedBy int, newNameExt ...string) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	loc, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().In(loc)
	nowStr := now.Format(time.RFC822)

	filter := bson.M{"_id": districtID}
	update := make(bson.M)
	update["$addToSet"] = bson.M{"NameExtension": bson.M{"$each": newNameExt}}
	update["$set"] = bson.M{"UpdatedBy": updatedBy, "UpdatedAt": nowStr}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return
	}
	if res.MatchedCount != 1 {
		err = errors.New("Không tìm thấy Quận cần sửa đổi. Hãy đảm bảo rằng district_id được truyền vào là chính xác")
		return
	}
	return
}
