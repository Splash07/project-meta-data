package bankv2

import (
	"sync"

	"gitlab.ghn.vn/online/common/config"
	"gitlab.com/Splash07/project-meta-data/repo"
)

// Repo type
type Repo repo.MongoV2

var (
	instance *Repo
	once     sync.Once
)

// New ..
func New() *Repo {
	once.Do(func() {
		instance = &Repo{
			Session:    config.GetConfig().MongoV2.Get("master-data"),
			Collection: "Banks",
		}
	})
	return instance
}
