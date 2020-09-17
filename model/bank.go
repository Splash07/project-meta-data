package model

import (

	//mssql xorm

	"gitlab.ghn.vn/online/common/gtype"
)

// Bank struct
type Bank struct {
	BankID      int               `xorm:"BankID"`
	BankName    string            `xorm:"BankName"`
	ShortName   string            `xorm:"ShortName"`
	Inactive    bool              `xorm:"Inactive"`
	CreatedDate *gtype.GHNSQLTime `xorm:"CreatedDate"`
	UpdatedDate *gtype.GHNSQLTime `xorm:"UpdatedDate"`
	HasBranch   bool              `json:"HasBranch"`
}

// IsExists struct
func (m Bank) IsExists() (ok bool) {
	if m.BankID != 0 {
		ok = true
	}
	return
}
