package ward

import (
	"context"
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
)

// RemoveFromNameExt func
func (r Mongo) RemoveFromNameExt(wardCode string, updatedBy int, nameExtToRemove ...string) (err error) {
	collection := r.Session.GetCollectionV2(r.Collection)

	loc, _ := time.LoadLocation("Asia/Bangkok")
	now := time.Now().In(loc)
	nowStr := now.Format(time.RFC822)

	filter := bson.M{"_id": wardCode}
	update := make(bson.M)
	update["$pull"] = bson.M{"NameExtension": bson.M{"$in": nameExtToRemove}}
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
