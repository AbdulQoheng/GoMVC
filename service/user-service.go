package service

import (
	"log"

	"github.com/AbdulQoheng/go-mvc/app/models"
	"github.com/AbdulQoheng/go-mvc/database/dto"
	"github.com/AbdulQoheng/go-mvc/database/entity"
	"github.com/mashingan/smapping"
)

// UserService is a contract.....
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userModels models.UserModels
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo models.UserModels) UserService {
	return &userService{
		userModels: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userModels.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userModels.ProfileUser(userID)
}
