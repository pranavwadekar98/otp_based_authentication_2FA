package main

import (
	"io"
	"net/http"
	"os"

	"example.com/controller"
	"example.com/middlewares"
	"example.com/service"
	"github.com/gin-gonic/gin"
)

var (
	dbService  service.DBService  = service.New()
	jwtService service.JWTService = service.NewJWTService()

	signUpController controller.SignUpController = controller.NewSignUpController(dbService)
	loginController  controller.LoginController  = controller.NewLoginController(dbService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	apiRoutes := server.Group("/api")
	{
		// apiRoutes.GET("/videos", func(ctx *gin.Context) {
		// 	ctx.JSON(200, videoController.FindAll())
		// })
		apiRoutes.POST("/signup", func(ctx *gin.Context) {
			err := signUpController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusCreated, gin.H{"message": "User is created"})
			}
		})
		apiRoutes.POST("/login", func(ctx *gin.Context) {
			token := loginController.Login(ctx)
			if token != "" {
				ctx.JSON(http.StatusOK, gin.H{"token": token})
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User does not exists or does not have access"})
			}
		})

		// server.POST("/login", func(ctx *gin.Context) {
		// 	token := loginController.Login(ctx)
		// 	if token != "" {
		// 		ctx.JSON(http.StatusOK, gin.H{
		// 			"token": token,
		// 		})
		// 	} else {
		// 		ctx.JSON(http.StatusUnauthorized, gin.H{
		// 			"error": "User does not exists or does not have access",
		// 		})
		// 	}
		// })

	}
	server.Run(":8081")
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdXRoLWJhY2tlbmQ6YXBwIiwic3ViIjoiODRhZjAyZjUtZmE3Yi00M2M0LWJhNjktOTUzZWE2MWY1ZTE3In0.EGYI9q007MXtJnYbAr9_ZwXEWEcP1s-nliJi5YlMkcQ
