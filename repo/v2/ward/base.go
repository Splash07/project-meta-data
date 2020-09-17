package ward

import (
	"sync"

	"gitlab.ghn.vn/online/common/config"
	"gitlab.com/Splash07/project-meta-data/repo"
)

// Repo type
type Repo repo.Mongo
type Mongo repo.MongoV2

var (
	instance *Repo
	once     sync.Once
	instanceMongo 	*Mongo
	onceMongo     	sync.Once
)

// New ..
func New() *Repo {
	once.Do(func() {
		instance = &Repo{
			Session:    config.GetConfig().Mongo.Get("f1pds"),
			Collection: "Ward",
		}
	})
	return instance
}

func NewMongo() *Mongo {
	onceMongo.Do(func() {
		instanceMongo = &Mongo{
			Session:config.GetConfig().MongoV2.Get("master-data"),
			Collection: "Ward",
		}
	})
	return instanceMongo
}
