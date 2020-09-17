package model

//mssql xorm

// ProvinceRegion struct
type ProvinceRegion struct {
	ProvinceRegionID int `xorm:"ProvinceRegionID"`
	ProvinceID       int `xorm:"ProvinceID"`
	RegionID         int `xorm:"RegionID"`
}

// IsExists struct
func (m ProvinceRegion) IsExists() (ok bool) {
	if m.ProvinceRegionID != 0 {
		ok = true
	}
	return
}
