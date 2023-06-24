package user_service

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/zahidakhyar/app-test/backend/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	WithContext(ctx context.Context) UserServiceInterface

	Store(user entity.User) entity.User
	Update(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	Profile(userID string) entity.User
}

type userConnection struct {
	ctx context.Context

	connection *gorm.DB
}

func NewUserService(db *gorm.DB) UserServiceInterface {
	return &userConnection{
		ctx:        context.Background(),
		connection: db,
	}
}

func (db *userConnection) WithContext(ctx context.Context) UserServiceInterface {
	return &userConnection{
		ctx:        ctx,
		connection: db.connection.WithContext(ctx),
	}
}

func (db *userConnection) Store(user entity.User) entity.User {
	user.Password = db.hashPassword([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) Update(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = db.hashPassword([]byte(user.Password))
	}

	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	db.connection.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		return false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}

	return user
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).First(&user)
	return user
}

func (db *userConnection) Profile(userID string) entity.User {
	var user entity.User
	db.connection.Find(&user, userID)
	return user
}

func (db *userConnection) hashPassword(password []byte) string {
	defer newrelic.FromContext(db.ctx).StartSegment("hashPassword").End()

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		panic("Error while hashing password")
	}

	return string(hash)
}
