package controllers

import (
	"encoding/json"
	"fmt"
	"ssoer/helpers"
	"ssoer/models"

	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
)

type UsersController struct {
	beego.Controller
}

const minPasswordLength = 8

// type UserResponse struct {
// 	Message string `json:"message"`
// }

// @Title Get user list
// @Description Get user list
// @Param   pageNumber    query   int false  0  "Fetch result page number"
// @Param   pageSize    query   int false  20  "Max number of items returned"
// @Success 200 {object} controllers.UserListResponseSchema
// @router / [get]
func (c *UsersController) Get() {
	var response UserListResponseSchema
	var pageNumber = c.Ctx.Input.Param(":pageNumber")
	fmt.Println("pageNumber:", pageNumber)

	o := orm.NewOrm()
	cnt, _ := o.
		QueryTable(new(models.User)).
		OrderBy("DateJoined").
		Filter("Deleted", false).
		RelatedSel().
		All(&response.Items)

	response.Meta.ItemsCount = int(cnt)

	c.Data["json"] = response
	c.ServeJSON()
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

	validateParams := func() {
		if len(userRequest.Email) == 0 {
			helpers.BadRequestError(c.Controller, "Invalid email", "ERROR_EMAIL")
		}
		if len(userRequest.Password) < minPasswordLength {
			helpers.BadRequestError(
				c.Controller,
				fmt.Sprintf("Password should contain at least %d characters", minPasswordLength),
				"ERROR_PASSWORD",
			)
		}
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &userRequest); err != nil {
		helpers.InternalError(c.Controller, err.Error(), "INIVALID_DATA")
	}

	validateParams()

	o := orm.NewOrm()

	var user models.User = models.User{
		Uuid:       uuid.New().String(),
		Email:      userRequest.Email,
		Password:   helpers.HashPassword(userRequest.Password),
		FirstName:  userRequest.FirstName,
		LastName:   userRequest.LastName,
		IsActive:   true,
		Deleted:    false,
		DateJoined: helpers.GetUTCNow(),
	}

	if _, err := o.Insert(&user); err != nil {
		helpers.InternalError(c.Controller, err.Error(), "ERROR_INSERT")
	}

	c.Data["json"] = user
	c.ServeJSON()
}
