package v2

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"gitlab.com/Splash07/project-meta-data/constants"
	"gitlab.com/Splash07/project-meta-data/model"
	bankV2Repo "gitlab.com/Splash07/project-meta-data/repo/v2/bankv2"
)

//Bank handler
var Bank bankHandler

func init() {
	Bank = bankHandler{}
}

type bankHandler struct {
}

//============================================
//=====================V2=====================
// GetV2 func
func (o bankHandler) GetV2(c echo.Context) (err error) {
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

	result, err := bankV2Repo.New().GetByID(request.ID)
	if err != nil {
		return
	}

	if !result.IsExists() {
		return fmt.Errorf("Bank not existed")
	}

	response := model.Bank{
		BankID:    result.ID,
		BankName:  result.Name,
		ShortName: result.ShortName,
		Inactive:  !result.Active,
		HasBranch: result.HasBranch,
	}

	return c.JSON(success(response))
}

// GetBankV2 ..
func (o bankHandler) GetBankV2(c echo.Context) (err error) {
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

	result, err := bankV2Repo.New().GetByIDV2(request.ID)
	if err != nil {
		return
	}

	type myResponse struct {
		ID          int        `json:"id"`
		Name        string     `json:"name"`
		ShortName   string     `json:"short_name"`
		HasBranch   bool       `json:"has_branch"`
		Status      int        `json:"status"`
		CreatedDate *time.Time `json:"created_date,omitempty"`
		UpdatedDate *time.Time `json:"updated_date,omitempty"`
	}
	response := myResponse{}

	jsonByte, err := json.Marshal(result)
	if err != nil {
		return errTemplate(constants.ServerErrorCommon, err)
	}
	if err = json.Unmarshal(jsonByte, &response); err != nil {
		return errTemplate(constants.ServerErrorCommon, err)
	}

	return c.JSON(success(response))
}

// GetAllV2 func
func (o bankHandler) GetAllV2(c echo.Context) (err error) {
	banks, err := bankV2Repo.New().All()
	if err != nil {
		return
	}

	response := []model.Bank{}
	for _, bank := range banks {
		temp := model.Bank{
			BankID:    bank.ID,
			BankName:  bank.Name,
			ShortName: bank.ShortName,
			Inactive:  !bank.Active,
			HasBranch: bank.HasBranch,
		}

		response = append(response, temp)
	}

	return c.JSON(success(response))
}

// GetMany ..
func (o bankHandler) GetMany(c echo.Context) (err error) {
	result := []model.BankV2{}
	offsetID := 0
	limit := 100
	for {
		banks, err := bankV2Repo.New().GetAllV2([]int{}, offsetID, limit)
		if err != nil {
			return errTemplate(constants.ServerErrorCommon, err)
		}
		result = append(result, banks...)
		if len(banks) < limit {
			break
		}
		offsetID = banks[len(banks)-1].ID
	}
	return c.JSON(success(result))
}

// GetMany ..
func (o bankHandler) GetAllActiveV2(c echo.Context) (err error) {
	result := []model.BankV2{}
	offsetID := 0
	limit := 100
	for {
		banks, err := bankV2Repo.New().GetAllV2([]int{constants.Status["ACTIVE"]}, offsetID, limit)
		if err != nil {
			return errTemplate(constants.ServerErrorCommon, err)
		}
		result = append(result, banks...)
		if len(banks) < limit {
			break
		}
		offsetID = banks[len(banks)-1].ID
	}
	return c.JSON(success(result))
}

// CheckHasBranch ..
func (o bankHandler) CheckHasBranch(c echo.Context) (err error) {

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

	bank, err := bankV2Repo.New().GetByID(request.ID)
	if err != nil {
		return
	}

	if !bank.IsExists() {
		return fmt.Errorf("Bank not existed")
	}

	return c.JSON(success(map[string]interface{}{
		"has_branch": bank.HasBranch,
	}))
}

// AddNewBank ..
func (o bankHandler) AddNewBank(c echo.Context) (err error) {
	type myRequest struct {
		Name      string `json:"name" validate:"required"`
		ShortName string `json:"short_name" validate:"required"`
		HasBranch bool   `json:"has_branch"`

		//
		CreatedIP       string `json:"created_ip" validate:"required"`
		CreatedEmployee int    `json:"created_employee" validate:"required"`
		CreatedSource   string `json:"created_source" validate:"required"`
		CreatedClient   int    `json:"created_client"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}

	_, err = bankV2Repo.New().Insert(model.BankV2{
		Name:      request.Name,
		ShortName: request.ShortName,
		Status:    constants.Status["ACTIVE"],
		HasBranch: request.HasBranch,

		//
		CreatedIP:       request.CreatedIP,
		CreatedEmployee: request.CreatedEmployee,
		CreatedSource:   request.CreatedSource,
		CreatedClient:   request.CreatedClient,
	})
	if err != nil {
		return errTemplate(constants.ServerErrorCommon, err)
	}
	return c.JSON(success(nil))
}

// RemoveBank ..
func (o bankHandler) RemoveBank(c echo.Context) (err error) {
	type myRequest struct {
		BankID int `json:"id" query:"id" validate:"required"`

		//
		UpdatedIP       string `json:"updated_ip" validate:"required"`
		UpdatedEmployee int    `json:"updated_employee" validate:"required"`
		UpdatedSource   string `json:"updated_source" validate:"required"`
		UpdatedClient   int    `json:"updated_client"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}

	bank, err := bankV2Repo.New().GetByID(request.BankID)
	if err != nil {
		return errTemplate(constants.BankNotFound, err)
	}

	if bank.Status != constants.Status["ACTIVE"] && bank.Status != constants.Status["NOT_ACTIVE"] {
		return errTemplate(constants.ServerErrorCommon, errors.New("Bank da duoc xoa"))
	}

	err = bankV2Repo.New().Update(model.BankV2{
		ID:        request.BankID,
		HasBranch: bank.HasBranch,
		Status:    constants.Status["DELETE"],
		//
		UpdatedIP:       request.UpdatedIP,
		UpdatedEmployee: request.UpdatedEmployee,
		UpdatedSource:   request.UpdatedSource,
		UpdatedClient:   request.UpdatedClient,
	})
	if err != nil {
		return errTemplate("Khong the xoa bank", err)
	}
	return c.JSON(success(nil))
}

// UpdateBank ..
func (o bankHandler) UpdateBank(c echo.Context) (err error) {
	type myRequest struct {
		BankID    int    `json:"id" validate:"required"`
		Name      string `json:"name" validate:"required"`
		ShortName string `json:"short_name" validate:"required"`
		HasBranch bool   `json:"has_branch"`

		//
		UpdatedIP       string `json:"updated_ip" validate:"required"`
		UpdatedEmployee int    `json:"updated_employee" validate:"required"`
		UpdatedSource   string `json:"updated_source" validate:"required"`
		UpdatedClient   int    `json:"updated_client"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}

	bank, err := bankV2Repo.New().GetByID(request.BankID)
	if err != nil {
		return errTemplate(constants.BankNotFound, err)
	}
	if bank.Status == constants.Status["DELETE"] {
		return errTemplate(constants.ServerErrorCommon, errors.New("Khong the cap nhat status cua bank da bi xoa"))
	}

	updateRequest := model.BankV2{
		ID:        request.BankID,
		Name:      request.Name,
		ShortName: request.ShortName,
		HasBranch: request.HasBranch,
		//
		UpdatedIP:       request.UpdatedIP,
		UpdatedEmployee: request.UpdatedEmployee,
		UpdatedSource:   request.UpdatedSource,
		UpdatedClient:   request.UpdatedClient,
	}

	err = bankV2Repo.New().Update(updateRequest)
	if err != nil {
		return errTemplate(constants.ServerErrorCommon, err)
	}
	return c.JSON(success(nil))
}

// SwitchStatus ..
func (o bankHandler) SwitchStatus(c echo.Context) (err error) {
	type myRequest struct {
		BankID int `json:"id" validate:"required"`

		UpdatedIP       string `json:"updated_ip" validate:"required"`
		UpdatedEmployee int    `json:"updated_employee" validate:"required"`
		UpdatedSource   string `json:"updated_source" validate:"required"`
		UpdatedClient   int    `json:"updated_client"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}
	if err = c.Validate(request); err != nil {
		return errTemplate(constants.UserErrCommon, err)
	}

	bank, err := bankV2Repo.New().GetByID(request.BankID)
	if err != nil {
		return errTemplate(constants.BankNotFound, err)
	}
	if !bank.IsExists() {
		return errTemplate(constants.BankNotFound, "Bank khong ton tai")
	}

	if bank.Status == constants.Status["DELETE"] {
		return errTemplate(constants.ServerErrorCommon, errors.New("Khong the chuyen status cua bank da bi xoa"))
	}

	var newStatus int
	switch bank.Status {
	case constants.Status["ACTIVE"]:
		newStatus = constants.Status["NOT_ACTIVE"]
	case constants.Status["NOT_ACTIVE"]:
		newStatus = constants.Status["ACTIVE"]
	case constants.Status["DELETE"]:
		return errTemplate(constants.ServerErrorCommon, errors.New("Khong the chuyen status cua bank da bi xoa"))
	default:
		return errTemplate(constants.ServerErrorCommon, errors.New("Status cua bank khong hop ly"))
	}

	err = bankV2Repo.New().Update(model.BankV2{
		ID:        bank.ID,
		Status:    newStatus,
		HasBranch: bank.HasBranch,

		UpdatedIP:       request.UpdatedIP,
		UpdatedEmployee: request.UpdatedEmployee,
		UpdatedSource:   request.UpdatedSource,
		UpdatedClient:   request.UpdatedClient,
	})
	if err != nil {
		return errTemplate(constants.ServerErrorCommon, err)
	}
	return c.JSON(success(nil))
}
