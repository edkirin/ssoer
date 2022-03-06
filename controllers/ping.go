package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type PingController struct {
	beego.Controller
}

type PingResponse struct {
	Message string `json:"message"`
}

// @Title Ping
// @Description Returns Pong message
// @Success 200 {object} controllers.PingResponse
// @router / [get]
func (c *PingController) Get() {
	c.Data["json"] = PingResponse{
		Message: "Pong!",
	}
	c.ServeJSON()
}
