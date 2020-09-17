package model

//mssql xorm

// Province struct
type Province struct {
	ProvinceID   int    `xorm:"ProvinceID" bson:"_id" json:"province_id"`
	ProvinceName string `xorm:"ProvinceName" bson:"ProvinceName" json:"province_name"`
	Code         string `xorm:"Code" bson:"Code" json:"code"`
}

// ProvinceChangeRequest struct
type ProvinceChangeRequest struct {
	ProvinceName string `xorm:"ProvinceName" bson:"ProvinceName" json:"province_name" validate:"required"`
	CountryID    int    `xorm:"CountryID" bson:"CountryID" json:"country_id"`
	Code         string `xorm:"Code" bson:"Code" json:"code" validate:"required"`
}

// IsExists struct
func (m Province) IsExists() (ok bool) {
	if m.ProvinceID != 0 {
		ok = true
	}
	return
}
