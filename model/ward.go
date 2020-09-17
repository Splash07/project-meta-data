package model

//mssql xorm

// Ward struct
type Ward struct {
	WardCode   string `json:"WardCode,omitempty" bson:"_id,omitempty"`
	DistrictID int    `json:"DistrictID,omitempty" bson:"DistrictID,omitempty"`
	WardName   string `json:"WardName,omitempty" bson:"WardName,omitempty"`
}

// IsExists struct
func (m Ward) IsExists() (ok bool) {
	if m.WardCode != "" {
		ok = true
	}
	return
}
