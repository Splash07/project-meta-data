package district

import (
	"sync"

	"gitlab.ghn.vn/online/common/config"
	"gitlab.com/Splash07/project-meta-data/repo"
)

// Mongo type
type Mongo repo.MongoV2

var (
	instanceMongo *Mongo
	onceMongo     sync.Once
)

// NewMongo func
func NewMongo() *Mongo {
	onceMongo.Do(func() {
		instanceMongo = &Mongo{
			Session:    config.GetConfig().MongoV2.Get("master-data"),
			Collection: "District",
		}
	})
	return instanceMongo
}
