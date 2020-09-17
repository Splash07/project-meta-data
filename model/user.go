package model

//mssql xorm

// User struct
type User struct {
	UserID   int    `xorm:"UserID"`
	FullName string `xorm:"FullName"`
	Email    string `xorm:"Email"`
	Phone    string `xorm:"Phone"`
}

// IsExists struct
func (m User) IsExists() (ok bool) {
	if m.UserID != 0 {
		ok = true
	}
	return
}
