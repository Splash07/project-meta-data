package model

//mssql xorm

// PermissionConfiguration struct
type PermissionConfiguration struct {
	DataPermissionConfigurationID int `xorm:"DataPermissionConfigurationID"`
	WarehouseID                   int `xorm:"WarehouseID"`
	DistrictID                    int `xorm:"DistrictID"`
}

// IsExists struct
func (m PermissionConfiguration) IsExists() (ok bool) {
	if m.DataPermissionConfigurationID != 0 {
		ok = true
	}
	return
}
