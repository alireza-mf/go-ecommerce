package services

import (
	"github.com/alireza-mf/go-ecommerce/models"
	"github.com/alireza-mf/go-ecommerce/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func ProvideUserService(t repositories.UserRepository) UserService {
	return UserService{UserRepository: t}
}

// FindById
func (u *UserService) FindById(UserId uint) (*models.User, error) {
	return u.UserRepository.FindByUserId(UserId)
}