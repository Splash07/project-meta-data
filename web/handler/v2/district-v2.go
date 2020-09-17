package v2

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gitlab.com/Splash07/project-meta-data/constants"
	"gitlab.com/Splash07/project-meta-data/model"
	districtRepo "gitlab.com/Splash07/project-meta-data/repo/v2/district"
	provinceRepo "gitlab.com/Splash07/project-meta-data/repo/v2/province"
)

//DistrictV2 handler
var DistrictV2 districtv2Handler

func init() {
	DistrictV2 = districtv2Handler{}
}

type districtv2Handler struct{}

// GetByID func
func (o districtv2Handler) GetOne(c echo.Context) (err error) {
	type myRequest struct {
		ID int `json:"id" query:"id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, errDistrict := districtRepo.NewMongo().GetOne(request.ID)
	if errDistrict != nil {
		err = errDistrict
		return
	}
	if !result.IsExists() {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(success(result))
}

// GetOneByCode func
func (o districtv2Handler) GetOneByCode(c echo.Context) (err error) {
	type myRequest struct {
		Code string `json:"code" query:"code" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, errDistrict := districtRepo.NewMongo().GetOneByCode(request.Code)
	if errDistrict != nil {
		err = errDistrict
		return
	}
	if !result.IsExists() {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(success(result))
}

// Get func
func (o districtv2Handler) GetByName(c echo.Context) (err error) {
	type myRequest struct {
		Name       string `json:"name" query:"name" validate:"required"`
		ProvinceID int    `json:"province_id" query:"province_id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}
	if err = c.Validate(request); err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}
	results, err := districtRepo.NewMongo().GetByName(strings.TrimSpace(request.Name), request.ProvinceID)
	if err != nil {
		return
	}

	if len(results) < 1 {
		return errTemplate(constants.DistrictNameNotValid, err)
	}

	if len(results) == 1 {
		return c.JSON(success(results[0]))
	}

	for _, result := range results {
		for _, name := range result.NameExtension {
			if strings.ToLower(name) == strings.ToLower(strings.TrimSpace(request.Name)) {
				return c.JSON(success(result))
			}
		}
	}
	return errTemplate(constants.DistrictNameNotValid, err)
}

// Get func
func (o districtv2Handler) GetAllByProvince(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID int `json:"province_id" query:"province_id"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}
	if err = c.Validate(request); err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}
	result, err := districtRepo.NewMongo().GetAll(request.ProvinceID)
	if err != nil {
		return
	}

	return c.JSON(success(result))
}

// Ward Update
func (o districtv2Handler) Update(c echo.Context) (err error) {
	type myRequest struct {
		DistrictID     int      `json:"district_id" query:"district_id" validate:"required"`
		ProvinceID     int      `json:"province_id" query:"province_id"`
		Name           string   `json:"name" query:"name"`
		DistrictEncode string   `json:"district_encode" query:"district_encode"`
		NameExtension  []string `json:"name_extension" query:"name_extension"`
		Priority       int      `json:"priority" query:"priority"`
		Type           int      `json:"type" query:"type"`
		Code           string   `json:"code" query:"code"`
		IsEnable       int      `json:"is_enable" query:"is_enable"`
		UpdatedBy      int      `json:"updated_by" query:"updated_by" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}

	district, err := districtRepo.NewMongo().GetOne(request.DistrictID)
	if err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}

	if !district.IsExists() {
		err = errTemplate(constants.WardNotFound, err)
		return
	}

	if request.Name != "" {
		district.DistrictName = request.Name
	}

	if request.DistrictEncode != "" {
		district.DistrictEncode = request.DistrictEncode
	}

	if request.Code != "" {
		district.Code = request.Code
	}

	if len(request.NameExtension) > 0 {
		district.NameExtension = request.NameExtension
	}

	if request.Priority > 0 {
		district.Priority = request.Priority
	}

	if request.Type > 0 {
		district.Type = request.Type
	}

	district.IsEnable = request.IsEnable
	district.UpdatedBy = request.UpdatedBy

	err = districtRepo.NewMongo().Update(&district)
	if err != nil {
		return
	}
	return c.JSON(success(district))
}

// Delete
func (o districtv2Handler) Delete(c echo.Context) (err error) {
	type myRequest struct {
		DistrictID int `json:"district_id" query:"district_id" validate:"required"`
		DeletedBy  int `json:"deleted_by" query:"deleted_by" validate:"required"`
	}

	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}
	district, err := districtRepo.NewMongo().GetOne(request.DistrictID)
	if err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}

	if !district.IsExists() {
		err = errTemplate(constants.WardNotFound, err)
		return
	}

	district.IsDeleted = 1
	district.UpdatedBy = request.DeletedBy

	err = districtRepo.NewMongo().Update(&district)
	if err != nil {
		return
	}

	return c.JSON(success(district))
}

// Ward Insert
func (o districtv2Handler) Insert(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID     int      `json:"province_id" query:"province_id" validate:"required"`
		Name           string   `json:"name" query:"name" validate:"required"`
		DistrictEncode string   `json:"district_encode" query:"district_encode" validate:"required"`
		NameExtension  []string `json:"name_extension" query:"name_extension" validate:"required"`
		Priority       int      `json:"priority" query:"priority" validate:"required"`
		Type           int      `json:"type" query:"type" validate:"required"`
		Code           string   `json:"code" query:"code" validate:"required"`
		CreatedBy      int      `json:"updated_by" query:"updated_by" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}

	district := model.DistrictV2{}

	district.DistrictName = request.Name
	district.ProvinceID = request.ProvinceID
	district.DistrictEncode = request.DistrictEncode
	district.Code = request.Code
	district.NameExtension = request.NameExtension
	district.Priority = request.Priority
	district.Type = request.Type
	district.IsEnable = 1
	district.IsDeleted = 0
	district.UpdatedBy = request.CreatedBy
	err = districtRepo.NewMongo().Insert(&district)

	if err != nil {
		return
	}
	return c.JSON(success(district))
}

//============================================================
//=========================== V2 =============================

// GetAllByProvinceV2 ..
func (o districtv2Handler) GetAllByProvinceV2(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID *int `json:"province_id" query:"province_id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	province, err := provinceRepo.NewMongo().GetOneV2(*request.ProvinceID)
	if err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if !province.IsExists() {
		return errTemplate(constants.ProvinceNotFound, "Province ID khong ton tai")
	}

	districts, err := districtRepo.NewMongo().GetAllV2(*request.ProvinceID)
	if err != nil {
		return
	}
	return c.JSON(success(districts))
}

// GetActiveOneV2 ..
func (o districtv2Handler) GetActiveOneV2(c echo.Context) (err error) {
	type myRequest struct {
		DistrictID int `json:"id" query:"id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	district, err := districtRepo.NewMongo().GetActiveOneV2(request.DistrictID)
	if err != nil {
		return
	}
	return c.JSON(success(district))

}

// GetOneDetailV2 ..
func (o districtv2Handler) GetOneDetailV2(c echo.Context) (err error) {
	type myRequest struct {
		DistrictID int `json:"id" query:"id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	district, err := districtRepo.NewMongo().GetOneV2(request.DistrictID)
	if err != nil {
		return
	}
	return c.JSON(success(district))

}

// AddV2 ..
func (o districtv2Handler) AddV2(c echo.Context) (err error) {
	type myRequest struct {
		DistrictName   string   `json:"district_name" validate:"required"`
		ProvinceID     int      `json:"province_id" validate:"required"`
		NameExtension  []string `json:"name_extension"`
		DistrictEncode string   `json:"district_encode"`
		Priority       int      `json:"priority"`
		Type           int      `json:"type" validate:"required"`
		SupportType    int      `json:"support_type" validate:"required"`
		CanUpdateCOD   *bool    `json:"can_update_cod" validate:"required"`

		//
		CreatedIP       string `json:"created_ip" validate:"required"`
		CreatedEmployee int    `json:"created_employee" validate:"required"`
		CreatedClient   int    `json:"created_client"`
		CreatedSource   string `json:"created_source" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}

	province, err := provinceRepo.NewMongo().GetOneV2(request.ProvinceID)
	if err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if !province.IsExists() {
		return errTemplate(constants.ProvinceNotFound, "Province ID khong ton tai")
	}

	if request.Type < 1 || request.Type > 3 {
		return errTemplate(constants.UserErrCommon, "Type khong hop le")
	}

	if request.SupportType < 1 || request.SupportType > 3 {
		return errTemplate(constants.UserErrCommon, "Support Type khong hop le")
	}

	err = districtRepo.NewMongo().InsertV2(model.DistrictV2{
		DistrictName:   request.DistrictName,
		ProvinceID:     request.ProvinceID,
		NameExtension:  request.NameExtension,
		DistrictEncode: request.DistrictEncode,
		Priority:       request.Priority,
		Type:           request.Type,
		SupportType:    request.SupportType,
		CanUpdateCOD:   *request.CanUpdateCOD,
		Status:         constants.Status["ACTIVE"],

		CreatedIP:       request.CreatedIP,
		CreatedEmployee: request.CreatedEmployee,
		CreatedClient:   request.CreatedClient,
		CreatedSource:   request.CreatedSource,
	})
	if err != nil {
		return
	}
	return c.JSON(success(nil))
}

// UpdateV2 ..
func (o districtv2Handler) UpdateV2(c echo.Context) (err error) {
	type myRequest struct {
		DistrictID            int      `json:"id" validate:"required"`
		DistrictName          string   `json:"district_name" validate:"required"`
		Priority              int      `json:"priority"`
		DistrictEncode        string   `json:"district_encode"`
		Type                  int      `json:"type" validate:"required"`
		SupportType           int      `json:"support_type" validate:"required"`
		NameExtensionToAdd    []string `json:"name_extension_add"`
		NameExtensionToRemove []string `json:"name_extension_remove"`
		CanUpdateCOD          *bool    `json:"can_update_cod" validate:"required"`

		//
		UpdatedIP       string `json:"updated_ip" validate:"required"`
		UpdatedEmployee int    `json:"updated_employee" validate:"required"`
		UpdatedClient   int    `json:"updated_client"`
		UpdatedSource   string `json:"updated_source" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}

	if request.Type < 1 || request.Type > 3 {
		return errTemplate(constants.UserErrCommon, "Type khong hop le")
	}

	if request.SupportType < 1 || request.SupportType > 3 {
		return errTemplate(constants.UserErrCommon, "Support Type khong hop le")
	}

	err = districtRepo.NewMongo().UpdateV2(districtRepo.UpdateDistrictRequest{
		DistrictID:            request.DistrictID,
		DistrictName:          request.DistrictName,
		Priority:              request.Priority,
		DistrictEncode:        request.DistrictEncode,
		Type:                  request.Type,
		SupportType:           request.SupportType,
		CanUpdateCOD:          *request.CanUpdateCOD,
		NameExtensionToAdd:    request.NameExtensionToAdd,
		NameExtensionToRemove: request.NameExtensionToRemove,

		//
		UpdatedIP:       request.UpdatedIP,
		UpdatedEmployee: request.UpdatedEmployee,
		UpdatedClient:   request.UpdatedClient,
		UpdatedSource:   request.UpdatedSource,
	})
	if err != nil {
		return errTemplate(constants.ServerErrorCommon, err)
	}

	return c.JSON(success(nil))
}

// RemoveV2 ..
func (o districtv2Handler) RemoveV2(c echo.Context) (err error) {
	type myRequest struct {
		DistrictID int `json:"id" validate:"required"`
		//
		UpdatedIP       string `json:"updated_ip" validate:"required"`
		UpdatedEmployee int    `json:"updated_employee" validate:"required"`
		UpdatedClient   int    `json:"updated_client"`
		UpdatedSource   string `json:"updated_source" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}

	err = districtRepo.NewMongo().UpdateDocumentV2(model.DistrictV2{
		DistrictID:      request.DistrictID,
		Status:          constants.Status["DELETE"],
		UpdatedIP:       request.UpdatedIP,
		UpdatedEmployee: request.UpdatedEmployee,
		UpdatedClient:   request.UpdatedClient,
		UpdatedSource:   request.UpdatedSource,
	})
	if err != nil {
		return
	}
	return c.JSON(success(nil))
}

// SwitchStatus
func (o districtv2Handler) SwitchStatus(c echo.Context) (err error) {
	type myRequest struct {
		DistrictID int `json:"id" validate:"required"`
		//
		UpdatedIP       string `json:"updated_ip" validate:"required"`
		UpdatedEmployee int    `json:"updated_employee" validate:"required"`
		UpdatedClient   int    `json:"updated_client"`
		UpdatedSource   string `json:"updated_source" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}

	district, err := districtRepo.NewMongo().GetOneV2(request.DistrictID)
	if err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}

	if !district.IsExists() {
		return errTemplate(constants.DistrictNotFound, "Quan can doi status khong ton tai")
	}

	var newStatus int
	switch district.Status {
	case constants.Status["ACTIVE"]:
		newStatus = constants.Status["NOT_ACTIVE"]
	case constants.Status["NOT_ACTIVE"]:
		newStatus = constants.Status["ACTIVE"]
	case constants.Status["DELETE"]:
		return errTemplate(constants.ServerErrorCommon, "Khong the update quan da bi xoa")
	default:
		return errTemplate(constants.ServerErrorCommon, "Status cua quan khong hop le")
	}

	err = districtRepo.NewMongo().UpdateDocumentV2(model.DistrictV2{
		DistrictID:      request.DistrictID,
		Status:          newStatus,
		UpdatedIP:       request.UpdatedIP,
		UpdatedEmployee: request.UpdatedEmployee,
		UpdatedClient:   request.UpdatedClient,
		UpdatedSource:   request.UpdatedSource,
	})
	if err != nil {
		return
	}
	return c.JSON(success(nil))
}
