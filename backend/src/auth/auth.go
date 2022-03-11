package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/thedevsaddam/govalidator"
	"github.com/zahidakhyar/app-test/backend/entity"
	"github.com/zahidakhyar/app-test/backend/helper"
	auth_dto "github.com/zahidakhyar/app-test/backend/src/auth/dto"
	auth_service "github.com/zahidakhyar/app-test/backend/src/auth/service"
	user_dto "github.com/zahidakhyar/app-test/backend/src/user/dto"
)

type AuthControllerInterface interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type AuthController struct {
	authService auth_service.AuthServiceInterface
	jwtService  auth_service.JwtServiceInterface
}

func NewAuthController(authService auth_service.AuthServiceInterface, jwtService auth_service.JwtServiceInterface) AuthControllerInterface {
	return &AuthController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginDto auth_dto.LoginDto

	rules := govalidator.MapData{
		"email":    []string{"required", "min:4", "max:20", "email"},
		"password": []string{"required", "min:4"},
	}

	opts := govalidator.Options{
		Request: ctx.Request,
		Data:    &loginDto,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	if e != nil {
		response := helper.BuildErrorResponse("Invalid request body", e, helper.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	auth := c.authService.VerifyCredential(loginDto.Email, loginDto.Password)
	if result, ok := auth.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(result.ID, 10))
		result.Token = generatedToken

		response := helper.BuildResponse(true, "Ok!", result)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.BuildErrorResponse("Invalid credential", "Invalid credential", helper.EmptyResponse{})
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (c *AuthController) Register(ctx *gin.Context) {
	var registerDto auth_dto.RegisterDto

	rules := govalidator.MapData{
		"name":     []string{"required"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"password": []string{"required", "min:4"},
	}

	opts := govalidator.Options{
		Request: ctx.Request,
		Data:    &registerDto,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	if e != nil {
		response := helper.BuildErrorResponse("Invalid request body", e, helper.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDto.Email) {
		response := helper.BuildErrorResponse("Duplicate email", "Duplicate email", helper.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.Store(registerDto)
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = generatedToken
		response := helper.BuildResponse(true, "Ok!", createdUser)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *AuthController) Update(ctx *gin.Context) {
	var updateDto user_dto.UpdateUserDto

	rules := govalidator.MapData{
		"name":     []string{"required"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"password": []string{"min:4"},
	}

	opts := govalidator.Options{
		Request: ctx.Request,
		Data:    &updateDto,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	if e != nil {
		response := helper.BuildErrorResponse("Invalid request body", e, helper.EmptyResponse{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	updateDto.ID = id
	updatedUser := c.authService.Update(updateDto)
	response := helper.BuildResponse(true, "Ok!", updatedUser)
	ctx.JSON(http.StatusOK, response)
}

func (c *AuthController) Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	profile := c.authService.Profile(id)
	response := helper.BuildResponse(true, "Ok!", profile)
	ctx.JSON(http.StatusOK, response)
}
