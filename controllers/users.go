package controllers

import (
	"encoding/json"
	"ssoer/helpers"
	"ssoer/models"

	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
)

type UsersController struct {
	beego.Controller
}

// type UserResponse struct {
// 	Message string `json:"message"`
// }

// @Title Create User
// @Description Create new user
// @Param	body		body 	models.UserCreateRequest	true		"body for user content"
// @Success 200 {object} models.User
// @router / [post]
func (c *UsersController) Post() {
	var userRequest models.UserCreateRequest

	var err = json.Unmarshal(c.Ctx.Input.RequestBody, &userRequest)
	if err != nil {
		helpers.InternalError(c.Controller, err.Error())
	}

	o := orm.NewOrm()

	var user models.User = models.User{
		Uuid:       uuid.New().String(),
		Email:      userRequest.Email,
		FirstName:  userRequest.FirstName,
		LastName:   userRequest.LastName,
		IsActive:   true,
		Deleted:    false,
		DateJoined: helpers.GetUTCNow(),
	}

	_, err = o.Insert(&user)
	if err != nil {
		helpers.InternalError(c.Controller, err.Error())
	}

	c.Data["json"] = user
	c.ServeJSON()
}
