package services

import (
	"errors"
	"fmt"
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
func (u *UserService) CreateUser(inputModel *models.UserInputForm) (*models.User, error) {
	fmt.Println(inputModel)
	isEmailExisted, err := u.UserRepository.FindByEmail(inputModel.Email)
	if err != nil {
		return nil, err
	}
	if isEmailExisted != nil {
		return nil, errors.New("CreateUser::email_is_existed")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(inputModel.Password), 8)

	userModel := &models.User{
		UserId:          uuid.Must(uuid.NewRandom()).String(),
		Email:           inputModel.Email,
		Password:        string(hashedPassword),
		DeliveryAddress: inputModel.DeliveryAddress,
		CreatedAt:       time.Now(),
	}

	return u.UserRepository.Create(userModel)
}
