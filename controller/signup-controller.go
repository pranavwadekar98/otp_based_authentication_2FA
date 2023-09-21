package controller

import (
	"example.com/entity"
	"example.com/service"
	"github.com/gin-gonic/gin"
)

type SignUpController interface {
	Save(ctx *gin.Context) error
	// 	ShowAll(ctx *gin.Context)
}

type signUpController struct {
	service service.DBService
}

func NewSignUpController(service service.DBService) SignUpController {
	return signUpController{service: service}
}

func (c signUpController) Save(ctx *gin.Context) error {
	var data entity.SignUpData
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		return err
	}
	err = c.service.Save(data)
	if err != nil {
		return err
	}
	return nil
}
