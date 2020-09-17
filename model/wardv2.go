package model

import "time"

// WardV2 ..
type WardV2 struct {
	WardCode      string   `json:"WardCode,omitempty" bson:"_id,omitempty,omitempty"`
	DistrictID    int      `json:"DistrictID,omitempty" bson:"DistrictID,omitempty,omitempty"`
	WardName      string   `json:"WardName,omitempty" bson:"WardName,omitempty,omitempty"`
	NameExtension []string `json:"NameExtension,omitempty" bson:"NameExtension,omitempty"`
	WardEncode    string   `json:"WardEncode,omitempty" bson:"WardEncode,omitempty"`
	IsEnable      int      `json:"IsEnable,omitempty" bson:"IsEnable,omitempty"`
	IsDeleted     int      `json:"IsDeleted,omitempty" bson:"IsDeleted,omitempty"`
	Priority      int      `json:"Priority,omitempty" bson:"Priority,omitempty"`
	CanUpdateCOD  bool     `json:"CanUpdateCOD,omitempty" bson:"CanUpdateCOD"`
	UpdatedBy     int      `json:"UpdatedBy,omitempty" bson:"UpdatedBy,omitempty"`
	CreatedAt     string   `json:"CreatedAt,omitempty" bson:"CreatedAt,omitempty"`
	UpdatedAt     string   `json:"UpdatedAt,omitempty" bson:"UpdatedAt,omitempty"`
	SupportType   int      `json:"SupportType,omitempty" bson:"SupportType,omitempty"`

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
func (m WardV2) IsExists() (ok bool) {
	if m.WardCode != "" {
		ok = true
	}
	return
}
