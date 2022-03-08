package controllers

import "ssoer/models"

type MetaSchema struct {
	ItemsCount int `json:"count"`
	PageCount  int `json:"pageCount"`
	PageNumber int `json:"pageNumber"`
}

type UserListResponseSchema struct {
	Meta  MetaSchema     `json:"meta"`
	Items []*models.User `json:"items"`
}
