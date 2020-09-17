package ward

import (
	"context"
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
)

// AppendToNameExt func
func (r Mongo) AppendToNameExt(wardCode string, updatedBy int, newNameExt ...string) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return
	}
	now := time.Now().In(loc)
	nowStr := now.Format(time.RFC822)

	filter := bson.M{"_id": wardCode}
	update := make(bson.M)
	update["$addToSet"] = bson.M{"NameExtension": bson.M{"$each": newNameExt}}
	update["$set"] = bson.M{"UpdatedBy": updatedBy, "UpdatedAt": nowStr}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return
	}
	if res.MatchedCount != 1 {
		err = errors.New("Không tìm thấy Phường/Xã cần sửa đổi. Hãy đảm bảo rằng ward_id được truyền vào là chính xác")
		return
	}
	return
}
