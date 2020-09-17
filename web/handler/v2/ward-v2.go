package v2

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gitlab.com/Splash07/project-meta-data/constants"
	"gitlab.com/Splash07/project-meta-data/model"
	districtRepo "gitlab.com/Splash07/project-meta-data/repo/v2/district"
	wardRepo "gitlab.com/Splash07/project-meta-data/repo/v2/ward"
)

//Ward handler
var WardV2 wardv2Handler

func init() {
	WardV2 = wardv2Handler{}
}

type wardv2Handler struct {
}

// GetOne func
// Might be removed ..
func (o wardv2Handler) GetOne(c echo.Context) (err error) {
	type myRequest struct {
		Code string `json:"code" query:"code" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	result, err := wardRepo.NewMongo().GetOne(request.Code)
	if err != nil {
		return
	}
	return c.JSON(success(result))
}

// GetAll func
// Might be removed ..
func (o wardv2Handler) GetAll(c echo.Context) (err error) {

	type myRequest struct {
		DistrictID int `json:"district_id" query:"district_id"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if request.DistrictID < 0 {
		return fmt.Errorf("District id is invalid: %d", request.DistrictID)
	}

	result, errWard := wardRepo.NewMongo().GetByDistrict(request.DistrictID)
	if errWard != nil {
		err = errWard
		return
	}
	return c.JSON(success(result))
}

// GetByName func
func (o wardv2Handler) GetByName(c echo.Context) (err error) {

	type myRequest struct {
		Name       string `json:"name" query:"name" validate:"required"`
		DistrictID int    `json:"district_id" query:"district_id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}

	results, errWard := wardRepo.NewMongo().GetByName(strings.TrimSpace(request.Name), request.DistrictID)
	if errWard != nil {
		return errTemplate(constants.WardNameNotValid, err)
	}

	if len(results) < 1 {
		return errTemplate(constants.WardNameNotValid, err)
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
	return errTemplate(constants.WardNameNotValid, err)
}

// Ward Update
// Might be removed ..
func (o wardv2Handler) Update(c echo.Context) (err error) {
	type myRequest struct {
		WardCode      string   `json:"ward_code" query:"ward_code" validate:"required"`
		Name          string   `json:"name" query:"name"`
		DistrictID    int      `json:"district_id" query:"district_id"`
		NameExtension []string `json:"name_extension" query:"name_extension"`
		WardEncode    string   `json:"ward_encode" query:"ward_encode"`
		Priority      int      `json:"priority" query:"priority"`
		IsEnable      int      `json:"is_enable" query:"is_enable"`
		UpdatedBy     int      `json:"updated_by" query:"updated_by" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}

	ward, err := wardRepo.NewMongo().GetOne(request.WardCode)
	if err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}

	if !ward.IsExists() {
		err = errTemplate(constants.WardNotFound, err)
		return
	}

	if request.Name != "" {
		ward.WardName = request.Name
	}

	if request.WardEncode != "" {
		ward.WardEncode = request.WardEncode
	}

	if len(request.NameExtension) > 0 {
		ward.NameExtension = request.NameExtension
	}

	if request.DistrictID > 0 {
		ward.DistrictID = request.DistrictID
	}

	if request.Priority > 0 {
		ward.Priority = request.Priority
	}

	ward.IsEnable = request.IsEnable
	ward.UpdatedBy = request.UpdatedBy

	err = wardRepo.NewMongo().Update(&ward)
	if err != nil {
		return
	}
	return c.JSON(success(ward))
}

// Ward Insert
// Might be removed ..
func (o wardv2Handler) Insert(c echo.Context) (err error) {
	type myRequest struct {
		WardCode      string   `json:"ward_code" query:"ward_code" validate:"required"`
		Name          string   `json:"name" query:"name"`
		DistrictID    int      `json:"district_id" query:"district_id"`
		NameExtension []string `json:"name_extension" query:"name_extension"`
		WardEncode    string   `json:"ward_encode" query:"ward_encode"`
		Priority      int      `json:"priority" query:"priority"`
		CreatedBy     int      `json:"created_by" query:"created_by" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}

	ward := model.WardV2{}

	ward.WardName = request.Name
	ward.WardEncode = request.WardEncode
	ward.NameExtension = request.NameExtension
	ward.DistrictID = request.DistrictID
	ward.Priority = request.Priority
	ward.IsEnable = 1
	ward.UpdatedBy = request.CreatedBy

	err = wardRepo.NewMongo().Insert(&ward)
	if err != nil {
		return
	}
	return c.JSON(success(ward))
}

// Delete
// Might be removed ..
func (o wardv2Handler) Delete(c echo.Context) (err error) {
	type myRequest struct {
		WardCode  string `json:"ward_code" query:"ward_code" validate:"required"`
		DeletedBy int    `json:"deleted_by" query:"deleted_by" validate:"required"`
	}

	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}
	ward, err := wardRepo.NewMongo().GetOne(request.WardCode)
	if err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}

	if !ward.IsExists() {
		err = errTemplate(constants.WardNotFound, err)
		return
	}

	ward.IsDeleted = 1
	ward.UpdatedBy = request.DeletedBy

	err = wardRepo.NewMongo().Update(&ward)
	if err != nil {
		return
	}

	return c.JSON(success(ward))
}

//============================================================
//============================ V2 ============================

// GetAllByDistrictV2 ..
func (o wardv2Handler) GetAllByDistrictV2(c echo.Context) (err error) {
	type myRequest struct {
		DistrictID *int `json:"district_id" query:"district_id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	district, err := districtRepo.NewMongo().GetOneV2(*request.DistrictID)
	if err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if !district.IsExists() {
		return errTemplate(constants.ProvinceNotFound, "Province ID khong ton tai")
	}

	wards, err := wardRepo.NewMongo().GetAllV2(*request.DistrictID)
	if err != nil {
		return
	}
	return c.JSON(success(wards))
}

// GetActiveOneV2 ..
func (o wardv2Handler) GetActiveOneV2(c echo.Context) (err error) {
	type myRequest struct {
		WardCode string `json:"code" query:"code" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	ward, err := wardRepo.NewMongo().GetActiveOneV2(request.WardCode)
	if err != nil {
		return
	}
	return c.JSON(success(ward))

}

// GetOneDetailV2 ..
func (o wardv2Handler) GetOneDetailV2(c echo.Context) (err error) {
	type myRequest struct {
		WardCode string `json:"code" query:"code" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	ward, err := wardRepo.NewMongo().GetOneV2(request.WardCode)
	if err != nil {
		return
	}
	return c.JSON(success(ward))

}

// AddV2 ..
func (o wardv2Handler) AddV2(c echo.Context) (err error) {
	type myRequest struct {
		WardName      string   `json:"ward_name" validate:"required"`
		DistrictID    int      `json:"district_id" validate:"required"`
		NameExtension []string `json:"name_extension"`
		WardEncode    string   `json:"ward_encode"`
		Priority      int      `json:"priority"`
		CanUpdateCOD  *bool    `json:"can_update_cod" validate:"required"`
		SupportType   int      `json:"support_type" validate:"required"`

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

	if request.SupportType < 1 || request.SupportType > 3 {
		return errTemplate(constants.UserErrCommon, "Support Type khong hop le")
	}

	district, err := districtRepo.NewMongo().GetOneV2(request.DistrictID)
	if err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if !district.IsExists() {
		return errTemplate(constants.DistrictNotFound, "District ID khong ton tai")
	}

	err = wardRepo.NewMongo().InsertV2(model.WardV2{
		WardName:      request.WardName,
		DistrictID:    request.DistrictID,
		NameExtension: request.NameExtension,
		WardEncode:    request.WardEncode,
		Priority:      request.Priority,
		Status:        constants.Status["ACTIVE"],
		CanUpdateCOD:  *request.CanUpdateCOD,
		SupportType:   request.SupportType,

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
func (o wardv2Handler) UpdateV2(c echo.Context) (err error) {
	type myRequest struct {
		WardCode              string   `json:"code" validate:"required"`
		WardName              string   `json:"ward_name" validate:"required"`
		Priority              int      `json:"priority"`
		WardEncode            string   `json:"ward_encode"`
		NameExtensionToAdd    []string `json:"name_extension_add"`
		NameExtensionToRemove []string `json:"name_extension_remove"`
		CanUpdateCOD          *bool    `json:"can_update_cod" validate:"required"`
		SupportType           int      `json:"support_type" validate:"required"`
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

	if request.SupportType < 1 || request.SupportType > 3 {
		return errTemplate(constants.UserErrCommon, "Support Type khong hop le")
	}

	err = wardRepo.NewMongo().UpdateV2(wardRepo.UpdateWardRequest{
		WardCode:              request.WardCode,
		WardName:              request.WardName,
		Priority:              request.Priority,
		WardEncode:            request.WardEncode,
		NameExtensionToAdd:    request.NameExtensionToAdd,
		NameExtensionToRemove: request.NameExtensionToRemove,
		CanUpdateCOD:          *request.CanUpdateCOD,
		SupportType:           request.SupportType,
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
func (o wardv2Handler) RemoveV2(c echo.Context) (err error) {
	type myRequest struct {
		WardCode string `json:"code" validate:"required"`
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

	err = wardRepo.NewMongo().UpdateDocumentV2(model.WardV2{
		WardCode:        request.WardCode,
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
func (o wardv2Handler) SwitchStatus(c echo.Context) (err error) {
	type myRequest struct {
		WardCode string `json:"code" validate:"required"`
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

	ward, err := wardRepo.NewMongo().GetOneV2(request.WardCode)
	if err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}

	if !ward.IsExists() {
		return errTemplate(constants.WardNotFound, "Phuong can doi status khong ton tai")
	}

	var newStatus int
	switch ward.Status {
	case constants.Status["ACTIVE"]:
		newStatus = constants.Status["NOT_ACTIVE"]
	case constants.Status["NOT_ACTIVE"]:
		newStatus = constants.Status["ACTIVE"]
	case constants.Status["DELETE"]:
		return errTemplate(constants.ServerErrorCommon, "Khong the update phuong/xoa da bi xoa")
	default:
		return errTemplate(constants.ServerErrorCommon, "Status cua phuong/xoa khong hop le")
	}

	err = wardRepo.NewMongo().UpdateDocumentV2(model.WardV2{
		WardCode:        request.WardCode,
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
