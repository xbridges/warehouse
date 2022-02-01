package httputils

import(
)

type LoginParam struct {
	ID			string	`json:"id" binding:"required"`
	Password	string	`json:"password" binding:"required"`
}