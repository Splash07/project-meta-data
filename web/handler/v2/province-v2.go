package v2

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gitlab.com/Splash07/project-meta-data/constants"
	"gitlab.com/Splash07/project-meta-data/model"
	provinceRepo "gitlab.com/Splash07/project-meta-data/repo/v2/province"
)

//ProvinceV2 handler
var ProvinceV2 provincev2Handler

func init() {
	ProvinceV2 = provincev2Handler{}
}

type provincev2Handler struct {
}

// Get func
func (o provincev2Handler) GetOne(c echo.Context) (err error) {
	type myRequest struct {
		ID int `json:"id" query:"id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	result, err := provinceRepo.NewMongo().GetOne(request.ID)
	if err != nil {
		return
	}
	if !result.IsExists() {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	return c.JSON(success(result))
}

// Get func
func (o provincev2Handler) GetByName(c echo.Context) (err error) {
	type myRequest struct {
		Name string `json:"name" query:"name" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	results, err := provinceRepo.NewMongo().GetByName(strings.TrimSpace(request.Name))
	if err != nil {
		return
	}
	if len(results) < 1 {
		return errTemplate(constants.ProvinceNameNotValid, err)
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
	return errTemplate(constants.ProvinceNameNotValid, err)
}

// Index func
func (o provincev2Handler) GetAll(c echo.Context) (err error) {
	result, err := provinceRepo.NewMongo().GetAll()
	if err != nil {
		return
	}
	return c.JSON(success(result))
}

func (o provincev2Handler) Update(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID     int      `json:"province_id" query:"province_id" validate:"required"`
		Priority       int      `json:"priority" query:"priority"`
		IsEnable       int      `json:"is_enable" query:"is_enable"`
		RegionID       int      `json:"region_id" query:"region_id" validate:"required"`
		UpdatedBy      int      `json:"updated_by" query:"updated_by" validate:"required"`
		Name           string   `json:"name" query:"name"`
		Code           string   `json:"code" query:"code"`
		ProvinceEncode string   `json:"province_encode" query:"province_encode"`
		NameExtension  []string `json:"name_extension" query:"name_extension"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}

	province, err := provinceRepo.NewMongo().GetOne(request.ProvinceID)
	if err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}

	if !province.IsExists() {
		err = errTemplate(constants.WardNotFound, err)
		return
	}

	if request.Name != "" {
		province.ProvinceName = request.Name
	}

	if len(request.NameExtension) > 0 {
		province.NameExtension = request.NameExtension
	}

	if request.ProvinceEncode != "" {
		province.ProvinceEncode = request.ProvinceEncode
	}

	if request.Code != "" {
		province.Code = request.Code
	}

	if request.Priority > 0 {
		province.Priority = request.Priority
	}

	province.IsEnable = request.IsEnable
	province.UpdatedBy = request.UpdatedBy

	err = provinceRepo.NewMongo().Update(&province)
	if err != nil {
		return
	}
	return c.JSON(success(province))
}

// Delete
func (o provincev2Handler) Delete(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID int `json:"province_id" query:"province_id" validate:"required"`
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
	province, err := provinceRepo.NewMongo().GetOne(request.ProvinceID)
	if err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}

	if !province.IsExists() {
		err = errTemplate(constants.WardNotFound, err)
		return
	}

	province.IsDeleted = 1
	province.UpdatedBy = request.DeletedBy

	err = provinceRepo.NewMongo().Update(&province)
	if err != nil {
		return
	}

	return c.JSON(success(province))
}

func (o provincev2Handler) Insert(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID     int      `json:"province_id" query:"province_id" validate:"required"`
		Priority       int      `json:"priority" query:"priority" validate:"required"`
		CreatedBy      int      `json:"created_by" query:"created_by" validate:"required"`
		Name           string   `json:"name" query:"name"  validate:"required"`
		ProvinceEncode string   `json:"province_encode" query:"province_encode"`
		NameExtension  []string `json:"name_extension" query:"name_extension" validate:"required"`
		CountryID      int      `json:"country_id" query:"country_id" validate:"required"`
		Code           string   `json:"code" query:"code" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		err = errTemplate(constants.UserErrCommon, err)
		return
	}

	province := model.ProvinceV2{}

	province.ProvinceID = request.ProvinceID
	province.ProvinceName = request.Name
	province.NameExtension = request.NameExtension
	province.ProvinceEncode = request.ProvinceEncode
	province.Code = request.Code
	province.CountryID = request.CountryID
	province.Priority = request.Priority
	province.IsEnable = 1
	province.IsDeleted = 0
	province.UpdatedBy = request.CreatedBy

	err = provinceRepo.NewMongo().Insert(&province)
	if err != nil {
		return
	}
	return c.JSON(success(province))
}

//=======================V2=======================//

// GetOneV2 ..
func (o provincev2Handler) GetActiveOneV2(c echo.Context) (err error) {
	type myRequest struct {
		ID int `json:"id" query:"id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	result, err := provinceRepo.NewMongo().GetActiveOneV2(request.ID)
	if err != nil {
		return
	}
	return c.JSON(success(result))
}

// GetOneDetailV2 ..
func (o provincev2Handler) GetOneDetailV2(c echo.Context) (err error) {
	type myRequest struct {
		ID int `json:"id" query:"id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	result, err := provinceRepo.NewMongo().GetOneV2(request.ID)
	if err != nil {
		return
	}
	return c.JSON(success(result))
}

// GetAllV2 ..
func (o provincev2Handler) GetAllV2(c echo.Context) (err error) {
	results, err := provinceRepo.NewMongo().GetAllV2()
	if err != nil {
		return
	}
	return c.JSON(success(results))
}

// AddV2 ..
func (o provincev2Handler) AddV2(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceName   string   `json:"province_name" validate:"required"`
		NameExtension  []string `json:"name_extension"`
		ProvinceEncode string   `json:"province_encode"`
		Priority       int      `json:"priority"`
		RegionID       int      `json:"region_id" validate:"required"`
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

	err = provinceRepo.NewMongo().InsertV2(model.ProvinceV2{
		ProvinceName:   request.ProvinceName,
		NameExtension:  request.NameExtension,
		ProvinceEncode: request.ProvinceEncode,
		CountryID:      1,
		Priority:       request.Priority,
		CanUpdateCOD:   *request.CanUpdateCOD,
		Status:         constants.Status["ACTIVE"],
		RegionID:       request.RegionID,

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
func (o provincev2Handler) UpdateV2(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID            int      `json:"id" validate:"required"`
		ProvinceName          string   `json:"province_name" validate:"required"`
		Priority              int      `json:"priority"`
		Code                  string   `json:"code"`
		RegionID              int      `json:"region_id"`
		CanUpdateCOD          *bool    `json:"can_update_cod" validate:"required"`
		ProvinceEncode        string   `json:"province_encode"`
		NameExtensionToAdd    []string `json:"name_extension_add"`
		NameExtensionToRemove []string `json:"name_extension_remove"`

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

	err = provinceRepo.NewMongo().UpdateV2(provinceRepo.UpdateProvinceRequest{
		ProvinceID:            request.ProvinceID,
		ProvinceName:          request.ProvinceName,
		Priority:              request.Priority,
		Code:                  request.Code,
		RegionID:              request.RegionID,
		ProvinceEncode:        request.ProvinceEncode,
		NameExtensionToAdd:    request.NameExtensionToAdd,
		NameExtensionToRemove: request.NameExtensionToRemove,
		CanUpdateCOD:          *request.CanUpdateCOD,

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
func (o provincev2Handler) RemoveV2(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID int `json:"id" validate:"required"`
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

	province, err := provinceRepo.NewMongo().GetOneV2(request.ProvinceID)
	if err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}

	if !province.IsExists() {
		return errTemplate(constants.ProvinceNotFound, "Tinh can xoa khong ton tai")
	}

	err = provinceRepo.NewMongo().UpdateDocumentV2(model.ProvinceV2{
		ProvinceID:      request.ProvinceID,
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
func (o provincev2Handler) SwitchStatus(c echo.Context) (err error) {
	type myRequest struct {
		ProvinceID int `json:"id" validate:"required"`
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

	province, err := provinceRepo.NewMongo().GetOneV2(request.ProvinceID)
	if err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}

	if !province.IsExists() {
		return errTemplate(constants.ProvinceNotFound, "Tinh can doi status khong ton tai")
	}

	var newStatus int
	switch province.Status {
	case constants.Status["ACTIVE"]:
		newStatus = constants.Status["NOT_ACTIVE"]
	case constants.Status["NOT_ACTIVE"]:
		newStatus = constants.Status["ACTIVE"]
	case constants.Status["DELETE"]:
		return errTemplate(constants.ServerErrorCommon, "Khong the update tinh da bi xoa")
	default:
		return errTemplate(constants.ServerErrorCommon, "Status cua tinh khong hop le")
	}

	err = provinceRepo.NewMongo().UpdateDocumentV2(model.ProvinceV2{
		ProvinceID:      request.ProvinceID,
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
