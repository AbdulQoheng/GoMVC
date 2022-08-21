package service

import (
	"log"

	"github.com/AbdulQoheng/go-mvc/app/models"
	"github.com/AbdulQoheng/go-mvc/database/dto"
	"github.com/AbdulQoheng/go-mvc/database/entity"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

// AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userModels models.UserModels
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRep models.UserModels) AuthService {
	return &authService{
		userModels: userRep,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userModels.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userModels.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(email string) entity.User {
	return service.userModels.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userModels.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
