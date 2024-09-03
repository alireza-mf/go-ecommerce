package services

import (
	"errors"
	"time"

	"github.com/alireza-mf/go-ecommerce/models"
	"github.com/alireza-mf/go-ecommerce/repositories"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func ProvideUserService(t repositories.UserRepository) UserService {
	return UserService{UserRepository: t}
}

// FindById
func (u *UserService) FindById(UserId string) (*models.User, error) {
	return u.UserRepository.FindByUserId(UserId)
}

// CreateUser
func (u *UserService) CreateUser(model *models.RegisterUser) (*models.User, error) {
	isEmailExisted, err := u.UserRepository.FindByEmail(model.Email)
	if err != nil {
		return nil, err
	}
	if isEmailExisted != nil {
		return nil, errors.New("CreateUser::email_is_existed")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(model.Password), 8)

	userModel := &models.User{
		UserId:          uuid.Must(uuid.NewRandom()).String(),
		Email:           model.Email,
		Password:        string(hashedPassword),
		DeliveryAddress: model.DeliveryAddress,
		CreatedAt:       time.Now(),
	}

	return u.UserRepository.Create(userModel)
}
