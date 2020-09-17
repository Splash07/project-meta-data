package v2

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gitlab.ghn.vn/online/common/config"
	ghandler "gitlab.ghn.vn/online/common/gstuff/handler"
)

var cfg = config.GetConfig()

func getRequestID(c echo.Context) (requestID string) {
	return ghandler.GetRequestID(c)
}

// success func
func success(data interface{}) (int, interface{}) {
	return ghandler.Success(data)
}

func errTemplate(codeMessageKey string, errInput interface{}) (err error) {

	err = fmt.Errorf("%v | %s ", errInput, codeMessageKey)
	return
}
