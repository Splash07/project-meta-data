package model

import "time"

// DistrictV2 ..
type DistrictV2 struct {
	DistrictID       int      `json:"DistrictID,omitempty" bson:"_id,omitempty"`
	ProvinceID       int      `json:"ProvinceID,omitempty" bson:"ProvinceID,omitempty"`
	DistrictName     string   `json:"DistrictName,omitempty" bson:"DistrictName,omitempty"`
	Code             string   `json:"Code,omitempty" bson:"Code,omitempty"`
	Type             int      `json:"Type,omitempty" bson:"Type,omitempty"`
	SupportType      int      `json:"SupportType,omitempty" bson:"SupportType,omitempty"`
	NameExtension    []string `json:"NameExtension,omitempty" bson:"NameExtension,omitempty"`
	DistrictEncode   string   `json:"DistrictEncode,omitempty" bson:"DistrictEncode,omitempty"`
	Priority         int      `json:"Priority,omitempty" bson:"Priority,omitempty"`
	IsEnable         int      `json:"IsEnable,omitempty" bson:"IsEnable,omitempty"`
	IsDeleted        int      `json:"IsDeleted,omitempty" bson:"IsDeleted,omitempty"`
	UpdatedBy        int      `json:"UpdatedBy,omitempty" bson:"UpdatedBy,omitempty"`
	CreatedAt        string   `json:"CreatedAt,omitempty" bson:"CreatedAt,omitempty"`
	UpdatedAt        string   `json:"UpdatedAt,omitempty" bson:"UpdatedAt,omitempty"`
	IsRepresentative int      `json:"IsRepresentative,omitempty" bson:"IsRepresentative,omitempty"`
	CanUpdateCOD     bool     `json:"CanUpdateCOD,omitempty" bson:"CanUpdateCOD"`
	Status           int      `json:"Status,omitempty" bson:"Status,omitempty"`
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
func (m DistrictV2) IsExists() (ok bool) {
	if m.DistrictID > 0 {
		ok = true
	}
	return
}
