package model

import (
	"time"
)

// BankV2 struct
type BankV2 struct {
	ID        int    `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name,omitempty"`
	ShortName string `json:"short_name" bson:"short_name,omitempty"`
	Active    bool   `json:"active" bson:"active,omitempty"`
	HasBranch bool   `json:"has_branch" bson:"has_branch"`
	Status    int    `json:"status" bson:"status,omitempty"`

	//
	CreatedIP       string     `json:"created_ip,omitempty" bson:"created_ip,omitempty"`
	CreatedEmployee int        `json:"created_employee,omitempty" bson:"created_employee,omitempty"`
	CreatedClient   int        `json:"created_client,omitempty" bson:"created_client,omitempty"`
	CreatedSource   string     `json:"created_source,omitempty" bson:"created_source,omitempty"`
	CreatedDate     *time.Time `json:"created_date,omitempty" bson:"created_date,omitempty"`

	//
	UpdatedIP       string     `json:"updated_ip,omitempty" bson:"updated_ip,omitempty"`
	UpdatedEmployee int        `json:"updated_employee,omitempty" bson:"updated_employee,omitempty"`
	UpdatedClient   int        `json:"updated_client,omitempty" bson:"updated_client,omitempty"`
	UpdatedSource   string     `json:"updated_source,omitempty" bson:"updated_source,omitempty"`
	UpdatedDate     *time.Time `json:"updated_date,omitempty" bson:"updated_date,omitempty"`
}

// IsExists struct
func (m BankV2) IsExists() (ok bool) {
	if m.ID != 0 {
		ok = true
	}
	return
}
