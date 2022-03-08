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

// @Title Get user list
// @Description Get user list
// @Param   pageNumber		query	int	false	"Fetch result page number"
// @Param   pageSize    	query	int	false	"Max number of items returned"
// @Success 200 {object} controllers.UserListResponseSchema
// @router / [get]
func (c *UsersController) Get() {
	var response UserListResponseSchema
	pageNumber, err := c.GetUint32("pageNumber")
	if err != nil {
		pageNumber = 1
	}
	pageSize, err := c.GetUint32("pageSize")
	if err != nil {
		pageSize = 10
	}

	o := orm.NewOrm()
	itemsCount, _ := o.
		QueryTable(new(models.User)).
		OrderBy("FirstName").
		Filter("Deleted", false).
		Offset((pageNumber - 1) * pageSize).
		Limit(pageSize).
		All(&response.Items)

	totalCount, _ := o.
		QueryTable(new(models.User)).
		Filter("Deleted", false).
		Count()

	response.Meta.ItemsCount = uint(itemsCount)
	response.Meta.PageNumber = uint(pageNumber)
	response.Meta.PageCount = uint(uint(totalCount)/uint(pageSize+1) + 1)

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
