package model

import "time"

//mssql xorm

// Service struct
type Service struct {
	ShippingServiceID int    `xorm:"ShippingServiceID"`
	Name              string `xorm:"Name"`
	ServiceType       int    `xorm:"ServiceType"`
	IsEnabled         bool   `xorm:"IsEnabled"`
}

type ServiceChangeRequest struct {
	Name         string    `xorm:"Name" json:"name"`
	ServiceType  int       `xorm:"ServiceType" json:"service_type"`
	DisplayOrder int       `xorm:"DisplayOrder" json:"display_order"`
	EndOrderTime time.Time `xorm:"EndOrderTime" json:"end_order_time"`
	Weight       int       `xorm:"Weight" json:"weight"`
	Type         int       `xorm:"Type" json:"type"`
	ExtID        string    `xorm:"ExtId" json:"ext_id"`
	IsEnabled    bool      `xorm:"IsEnabled" json:"is_enabled"`
}

// IsExists struct
func (m Service) IsExists() (ok bool) {
	if m.ShippingServiceID != 0 {
		ok = true
	}
	return
}
