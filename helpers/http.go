package helpers

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

func createError(c beego.Controller, statusCode int, errMessage string, errCode string) {
	type ErrorResponse struct {
		Message   string `json:"message"`
		ErrorCode string `json:"errorCode"`
	}

	errResponse := ErrorResponse{
		Message:   errMessage,
		ErrorCode: errCode,
	}
	errJson, _ := json.Marshal(errResponse)

	c.Ctx.Output.Header("Content-Type", "application/json")
	c.CustomAbort(statusCode, string(errJson))
}

func BadRequestError(c beego.Controller, errMessage string, errCode string) {
	createError(c, 400, errMessage, errCode)
}

func InternalError(c beego.Controller, errMessage string, errCode string) {
	createError(c, 500, errMessage, errCode)
}
