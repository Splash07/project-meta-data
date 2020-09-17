package model

import (
	"time"
)

// BankBranch struct
type BankBranch struct {
	ID      int    `json:"_id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address,omitempty"`
	Phone   string `json:"phone" bson:"phone,omitempty"`
	BankID  int    `json:"bank_id" bson:"bank_id"`

	//
	CreatedIP       string    `json:"created_ip" bson:"created_ip,omitempty"`
	CreatedEmployee int       `json:"created_employee" bson:"created_employee,omitempty"`
	CreatedClient   int       `json:"created_client" bson:"created_client,omitempty"`
	CreatedSource   string    `json:"created_source" bson:"created_source,omitempty"`
	CreatedDate     time.Time `json:"created_date" bson:"created_date,omitempty"`

	//
	UpdatedIP       string    `json:"updated_ip" bson:"updated_ip,omitempty"`
	UpdatedEmployee int       `json:"updated_employee" bson:"updated_employee,omitempty"`
	UpdatedClient   int       `json:"updated_client" bson:"updated_client,omitempty"`
	UpdatedSource   string    `json:"updated_source" bson:"updated_source,omitempty"`
	UpdatedDate     time.Time `json:"updated_date" bson:"updated_date,omitempty"`
}

// IsExists struct
func (m BankBranch) IsExists() (ok bool) {
	if m.ID != 0 {
		ok = true
	}
	return
}
