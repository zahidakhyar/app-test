package controller

import "github.com/gin-gonic/gin"

type AuthInterface interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}

type Auth struct{}

func NewAuth() AuthInterface {
	return &Auth{}
}

func (auth *Auth) Login(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Login",
	})
}

func (auth *Auth) Register(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Register",
	})
}
