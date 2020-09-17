package route

import (
	"github.com/labstack/echo/v4"
	groute "gitlab.ghn.vn/online/common/gstuff/route"
	handlerV2 "gitlab.com/Splash07/project-meta-data/web/handler/v2"
)

// MasterData route
func MasterData(e *echo.Echo) {

	// API ///////////////////////////////////
	API := groute.APIRoute(e)

	// bank
	API.Any("/bank", handlerV2.Bank.GetV2)
	API.Any("/banks", handlerV2.Bank.GetAllV2)
	API.Any("/banks/check-has-branch", handlerV2.Bank.CheckHasBranch)

	// APIV2 ///////////////////////////////////
	APIV2 := API.Group("/v2")

	// bank
	APIV2.Any("/bank", handlerV2.Bank.GetBankV2)
	APIV2.Any("/banks", handlerV2.Bank.GetMany)
	APIV2.Any("/banks/active", handlerV2.Bank.GetAllActiveV2)
	APIV2.Any("/bank/add", handlerV2.Bank.AddNewBank)
	APIV2.Any("/bank/remove", handlerV2.Bank.RemoveBank)
	APIV2.Any("/bank/update", handlerV2.Bank.UpdateBank)
	APIV2.Any("/bank/switch-status", handlerV2.Bank.SwitchStatus)

	// district
	APIV2.Any("/district", handlerV2.DistrictV2.GetOne)                    // Might be removed ...
	APIV2.Any("/district/getOneByCode", handlerV2.DistrictV2.GetOneByCode) // Might be removed ...
	APIV2.Any("/districts", handlerV2.DistrictV2.GetAllByProvince)         // Might be removed ...
	APIV2.Any("/district/update", handlerV2.DistrictV2.Update)             // Might be removed ..
	APIV2.Any("/district/delete", handlerV2.DistrictV2.Delete)             // Might be removed ..
	APIV2.Any("/district/insert", handlerV2.DistrictV2.Insert)             // Might be removed ..
	APIV2.Any("/district/v2", handlerV2.DistrictV2.GetActiveOneV2)
	APIV2.Any("/district/detail-v2", handlerV2.DistrictV2.GetOneDetailV2)
	APIV2.Any("/districts/v2", handlerV2.DistrictV2.GetAllByProvinceV2)
	APIV2.Any("/district/by-name", handlerV2.DistrictV2.GetByName)
	APIV2.Any("/district/update-v2", handlerV2.DistrictV2.UpdateV2)
	APIV2.Any("/district/remove-v2", handlerV2.DistrictV2.RemoveV2)
	APIV2.Any("/district/add-v2", handlerV2.DistrictV2.AddV2)
	APIV2.Any("/district/switch-status", handlerV2.DistrictV2.SwitchStatus)

	// ward
	APIV2.Any("/ward", handlerV2.WardV2.GetOne)        // Might be removed ..
	APIV2.Any("/wards", handlerV2.WardV2.GetAll)       // Might be removed ..
	APIV2.Any("/ward/update", handlerV2.WardV2.Update) // Might be removed ..
	APIV2.Any("/ward/delete", handlerV2.WardV2.Delete) // Might be removed ..
	APIV2.Any("/ward/insert", handlerV2.WardV2.Insert) // Might be removed ..
	APIV2.Any("/ward/v2", handlerV2.WardV2.GetActiveOneV2)
	APIV2.Any("/ward/detail-v2", handlerV2.WardV2.GetOneDetailV2)
	APIV2.Any("/wards/v2", handlerV2.WardV2.GetAllByDistrictV2)
	APIV2.Any("/ward/detail-v2", handlerV2.WardV2.GetOneDetailV2)
	APIV2.Any("/ward/by-name", handlerV2.WardV2.GetByName)
	APIV2.Any("/ward/update-v2", handlerV2.WardV2.UpdateV2)
	APIV2.Any("/ward/remove-v2", handlerV2.WardV2.RemoveV2)
	APIV2.Any("/ward/add-v2", handlerV2.WardV2.AddV2)
	APIV2.Any("/ward/switch-status", handlerV2.WardV2.SwitchStatus)

	// province
	APIV2.Any("/provinces", handlerV2.ProvinceV2.GetAll)       // Might be removed ..
	APIV2.Any("/province/update", handlerV2.ProvinceV2.Update) // Might be removed ..
	APIV2.Any("/province/delete", handlerV2.ProvinceV2.Delete) // Might be removed ..
	APIV2.Any("/province/insert", handlerV2.ProvinceV2.Insert) // Might be removed ..

	APIV2.Any("/province/v2", handlerV2.ProvinceV2.GetActiveOneV2)
	APIV2.Any("/province/detail-v2", handlerV2.ProvinceV2.GetOneDetailV2)
	APIV2.Any("/provinces/v2", handlerV2.ProvinceV2.GetAllV2)
	APIV2.Any("/province/by-name", handlerV2.ProvinceV2.GetByName)
	APIV2.Any("/province/update-v2", handlerV2.ProvinceV2.UpdateV2)
	APIV2.Any("/province/remove-v2", handlerV2.ProvinceV2.RemoveV2)
	APIV2.Any("/province/add-v2", handlerV2.ProvinceV2.AddV2)
	APIV2.Any("/province/switch-status", handlerV2.ProvinceV2.SwitchStatus)
}
