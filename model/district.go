package model

import "time"

//mssql xorm

// District struct
type District struct {
	DistrictID       int       `xorm:"DistrictID"`
	ProvinceID       int       `xorm:"ProvinceID"`
	DistrictName     string    `xorm:"DistrictName"`
	Code             string    `xorm:"Code"`
	MiningText       string    `xorm:"MiningText"`
	Priority         int       `xorm:"Priority"`
	IsRepresentative bool      `xorm:"IsRepresentative"`
	Type             int       `xorm:"Type"`
	SupportType      int       `xorm:"SupportType"`
	Description      string    `xorm:"Description"`
	CreatedDate      time.Time `xorm:"CreatedDate"`
	UpdatedDate      time.Time `xorm:"UpdatedDate"`
}

//DistrictChangeRequest struct
type DistrictChangeRequest struct {
	ProvinceID       int       `xorm:"ProvinceID" json:"province_id"`
	DistrictName     string    `xorm:"DistrictName" json:"district_name"`
	Code             string    `xorm:"Code" json:"code"`
	MiningText       string    `xorm:"MiningText" json:"mining_text"`
	Priority         int       `xorm:"Priority" json:"priority"`
	IsRepresentative bool      `xorm:"IsRepresentative" json:"is_representative"`
	Type             int       `xorm:"Type" json:"type"`
	SupportType      int       `xorm:"SupportType" json:"support_type"`
	Description      string    `xorm:"Description" json:"description"`
	CreatedDate      time.Time `xorm:"CreatedDate" json:"created_date"`
	UpdatedDate      time.Time `xorm:"UpdatedDate" json:"updated_date"`
}

// IsExists struct
func (m District) IsExists() (ok bool) {
	if m.DistrictID != 0 {
		ok = true
	}
	return
}
