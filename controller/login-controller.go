package controller

import (
	"example.com/entity"
	"example.com/service"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	dbService  service.DBService
	jwtService service.JWTService
}

func NewLoginController(dbService service.DBService, jwtService service.JWTService) LoginController {
	return loginController{dbService: dbService, jwtService: jwtService}
}

func (c loginController) Login(ctx *gin.Context) string {
	var data entity.SignUpData
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		return ""
	}
	err = c.dbService.Find(data)
	if err == nil {
		return c.jwtService.GenerateToken(data.Phone)
	}
	return ""
}
