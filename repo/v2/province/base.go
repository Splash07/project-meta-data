package province

import (
	"sync"

	"gitlab.ghn.vn/online/common/config"
	"gitlab.com/Splash07/project-meta-data/repo"
)

// Repo type
type Repo repo.MSSQL
type Mongo repo.MongoV2

var (
	instance      *Repo
	instanceMongo *Mongo
	once          sync.Once
	onceMongo     sync.Once
)

// New ..
func New() *Repo {
	once.Do(func() {
		instance = &Repo{
			Session:    config.GetConfig().MSSQL.Get("master"),
			Table:      "Province",
			PrimaryKey: "ProvinceID",
		}
	})
	return instance
}

// New ..
func NewMongo() *Mongo {
	onceMongo.Do(func() {
		instanceMongo = &Mongo{
			Session:    config.GetConfig().MongoV2.Get("master-data"),
			Collection: "Province",
		}
	})
	return instanceMongo
}
