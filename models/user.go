package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int       `json:"-"`
	Uuid       string    `orm:"column(uuid);type(uuid);null;unique" json:"uuid"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	IsActive   bool      `json:"-"`
	Deleted    bool      `json:"-"`
	DateJoined time.Time `json:"dateJoined"`
	LastLogin  time.Time `json:"lastLogin"`
}

type UserCreateRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (u *User) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(
		new(User),
	)
}
