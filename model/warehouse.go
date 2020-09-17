package model

import "time"

// Warehouse struct
type Warehouse struct {
	WarehouseID        int       `xorm:"WarehouseID" json:"warehouse_id"`
	WarehouseCode      string    `xorm:"WarehouseCode" json:"warehouse_code"`
	WarehouseName      string    `xorm:"WarehouseName" json:"warehouse_name"`
	WarehouseAddress   string    `xorm:"WarehouseAddress" json:"warehouse_address"`
	WarehouseManager   string    `xorm:"WarehouseManager" json:"warehouse_manager"`
	WarehousePhoneNum  string    `xorm:"WarehousePhoneNum" json:"warehouse_phonenum"`
	WarehouseShortName string    `xorm:"WarehouseShortName" json:"warehouse_shortname"`
	WarehouseFax       string    `xorm:"WarehouseFax" json:"warehouse_fax"`
	Latitude           float32   `xorm:"Latitude" json:"latitude"`
	Longtitude         float32   `xorm:"Longtitude" json:"longtitude"`
	DistrictID         int       `xorm:"DistrictID" json:"district_id"`
	LeftNo             int       `xorm:"LeftNo" json:"left_no"`
	RightNo            int       `xorm:"RightNo" json:"right_no"`
	CreatedDate        time.Time `xorm:"CreatedDate"`
	MaxSOLoadPerTimes  int       `xorm:"MaxSOLoadPerTimes" json:"max_so_load_per_time"`
	DeactivateDate     time.Time `xorm:"DeactivateDate" json:"deactive_time"`
	NumPickPerDay      int       `xorm:"NumPickPerDay" json:"num_pick_per_day"`
	NumDeliverPerDay   int       `xorm:"NumDeliverPerDay" json:"num_delivery_per_day"`
	NumReturnPerDay    int       `xorm:"NumReturnPerDay" json:"num_return_per_day"`
	NumPSOPerDay       int       `xorm:"NumPSOPerDay" json:"num_pso_per_day"`
	NumDSOPerDay       int       `xorm:"NumDSOPerDay" json:"num_dso_per_day"`
	NumRSOPerDay       int       `xorm:"NumRSOPerDay" json:"num_rso_per_day"`
	UpdatedDate        time.Time `xorm:"UpdatedDate"`
	StationType        int       `xorm:"StationType" json:"station_type"`
}

type WarehouseChangeRequest struct {
	WarehouseCode      string    `xorm:"WarehouseCode" json:"warehouse_code"`
	WarehouseName      string    `xorm:"WarehouseName" json:"warehouse_name"`
	WarehouseAddress   string    `xorm:"WarehouseAddress" json:"warehouse_address"`
	WarehouseManager   string    `xorm:"WarehouseManager" json:"warehouse_manager"`
	WarehousePhoneNum  string    `xorm:"WarehousePhoneNum" json:"warehouse_phonenum"`
	WarehouseShortName string    `xorm:"WarehouseShortName" json:"warehouse_shortname"`
	WarehouseFax       string    `xorm:"WarehouseFax" json:"warehouse_fax"`
	Latitude           float32   `xorm:"Latitude" json:"latitude"`
	Longtitude         float32   `xorm:"Longtitude" json:"longtitude"`
	DistrictID         int       `xorm:"DistrictID" json:"district_id"`
	LeftNo             int       `xorm:"LeftNo" json:"left_no"`
	RightNo            int       `xorm:"RightNo" json:"right_no"`
	CreatedDate        time.Time `xorm:"CreatedDate"`
	MaxSOLoadPerTimes  int       `xorm:"MaxSOLoadPerTimes" json:"max_so_load_per_time"`
	DeactivateDate     time.Time `xorm:"DeactivateDate" json:"deactive_time"`
	NumPickPerDay      int       `xorm:"NumPickPerDay" json:"num_pick_per_day"`
	NumDeliverPerDay   int       `xorm:"NumDeliverPerDay" json:"num_delivery_per_day"`
	NumReturnPerDay    int       `xorm:"NumReturnPerDay" json:"num_return_per_day"`
	NumPSOPerDay       int       `xorm:"NumPSOPerDay" json:"num_pso_per_day"`
	NumDSOPerDay       int       `xorm:"NumDSOPerDay" json:"num_dso_per_day"`
	NumRSOPerDay       int       `xorm:"NumRSOPerDay" json:"num_rso_per_day"`
	UpdatedDate        time.Time `xorm:"UpdatedDate"`
	StationType        int       `xorm:"StationType" json:"station_type"`
}

// IsExists func
func (m Warehouse) IsExists() (ok bool) {
	if m.WarehouseID != 0 {
		ok = true
	}
	return
}
