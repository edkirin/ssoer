package controllers

import "ssoer/models"

type MetaSchema struct {
	ItemsCount uint `json:"count"`
	PageCount  uint `json:"pageCount"`
	PageNumber uint `json:"pageNumber"`
}

type UserListResponseSchema struct {
	Meta  MetaSchema     `json:"meta"`
	Items []*models.User `json:"items"`
}
