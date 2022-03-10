package auth_service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/zahidakhyar/app-test/backend/entity"
	auth_dto "github.com/zahidakhyar/app-test/backend/src/auth/dto"
	user_dto "github.com/zahidakhyar/app-test/backend/src/user/dto"
	user_service "github.com/zahidakhyar/app-test/backend/src/user/service"
)

type AuthServiceInterface interface {
	Store(user auth_dto.RegisterDto) entity.User
	Update(user user_dto.UpdateUserDto) entity.User
	Profile(userID string) entity.User
	VerifyCredential(email string, password string) interface{}
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type AuthService struct {
	userService user_service.UserServiceInterface
}

func NewAuthService(userService user_service.UserServiceInterface) AuthServiceInterface {
	return &AuthService{
		userService: userService,
	}
}

func (service *AuthService) VerifyCredential(email string, password string) interface{} {
	return service.userService.VerifyCredential(email, password)
}

func (service *AuthService) Store(user auth_dto.RegisterDto) entity.User {
	userData := entity.User{}

	err := smapping.FillStruct(&userData, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map fields: %v", err)
	}

	res := service.userService.Store(userData)
	return res
}

func (service *AuthService) Update(user user_dto.UpdateUserDto) entity.User {
	userData := entity.User{}

	err := smapping.FillStruct(&userData, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map fields: %v", err)
	}

	res := service.userService.Update(userData)
	return res
}

func (service *AuthService) Profile(userID string) entity.User {
	return service.userService.Profile(userID)
}

func (service *AuthService) FindByEmail(email string) entity.User {
	return service.userService.FindByEmail(email)
}

func (service *AuthService) IsDuplicateEmail(email string) bool {
	res := service.userService.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
