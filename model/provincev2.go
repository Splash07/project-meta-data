package model

import "time"

//ProvinceV2 ..
type ProvinceV2 struct {
	ProvinceID     int      `json:"ProvinceID" bson:"_id,omitempty"`
	ProvinceName   string   `json:"ProvinceName" bson:"ProvinceName,omitempty"`
	CountryID      int      `json:"CountryID" bson:"CountryID,omitempty"`
	Code           string   `json:"Code,omitempty" bson:"Code,omitempty"`
	NameExtension  []string `json:"NameExtension,omitempty" bson:"NameExtension,omitempty"`
	ProvinceEncode string   `json:"ProvinceEncode,omitempty" bson:"ProvinceEncode,omitempty"`
	RegionID       int      `json:"RegionID,omitempty" bson:"RegionID,omitempty"`
	Priority       int      `json:"Priority,omitempty" bson:"Priority,omitempty"`
	CanUpdateCOD   bool     `json:"CanUpdateCOD,omitempty" bson:"CanUpdateCOD"`
	IsEnable       int      `json:"IsEnable,omitempty" bson:"IsEnable,omitempty"`
	IsDeleted      int      `json:"IsDeleted,omitempty" bson:"IsDeleted,omitempty"`
	UpdatedBy      int      `json:"UpdatedBy,omitempty" bson:"UpdatedBy,omitempty"`
	CreatedAt      string   `json:"CreatedAt,omitempty" bson:"CreatedAt,omitempty"`
	UpdatedAt      string   `json:"UpdatedAt,omitempty" bson:"UpdatedAt,omitempty"`

	Status int `json:"Status,omitempty" bson:"Status,omitempty"`
	//
	CreatedIP       string     `json:"CreatedIP,omitempty" bson:"CreatedIP,omitempty"`
	CreatedEmployee int        `json:"CreatedEmployee,omitempty" bson:"CreatedEmployee,omitempty"`
	CreatedClient   int        `json:"CreatedClient,omitempty" bson:"CreatedClient,omitempty"`
	CreatedSource   string     `json:"CreatedSource,omitempty" bson:"CreatedSource,omitempty"`
	CreatedDate     *time.Time `json:"CreatedDate,omitempty" bson:"CreatedDate,omitempty"`

	//
	UpdatedIP       string     `json:"UpdatedIP,omitempty" bson:"UpdatedIP,omitempty"`
	UpdatedEmployee int        `json:"UpdatedEmployee,omitempty" bson:"UpdatedEmployee,omitempty"`
	UpdatedClient   int        `json:"UpdatedClient,omitempty" bson:"UpdatedClient,omitempty"`
	UpdatedSource   string     `json:"UpdatedSource,omitempty" bson:"UpdatedSource,omitempty"`
	UpdatedDate     *time.Time `json:"UpdatedDate,omitempty" bson:"UpdatedDate,omitempty"`
}

// IsExists struct
func (m ProvinceV2) IsExists() (ok bool) {
	if m.ProvinceID > 0 {
		ok = true
	}
	return
}
