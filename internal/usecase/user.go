package usecase

import (
	"github.com/fierzahaikkal/neocourse-be-boilerplate-golang/internal/entity"
	userModel "github.com/fierzahaikkal/neocourse-be-boilerplate-golang/internal/model/user"
	"github.com/fierzahaikkal/neocourse-be-boilerplate-golang/internal/repository"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	userRepo *repository.UserRepository
	log      *log.Logger
}

func NewAuthUsaCase(userRepo *repository.UserRepository, log *log.Logger) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
		log:      log,
	}
}

func (u *AuthUseCase) Register(req *userModel.SignUpRequest) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashed),
		Name:     req.Name,
	}
	err = u.userRepo.Register(&user)
	if err != nil {
		return err
	}
	return nil
}

func (u *AuthUseCase) SignIn(req *userModel.SignInRequest) error {
	var user entity.User
	err := u.userRepo.DB.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return err
	}
	return nil
}
