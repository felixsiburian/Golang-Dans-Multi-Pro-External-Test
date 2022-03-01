package usecase

import (
	"Golang-Dans-Multi-Pro-External-Test/service"
	"Golang-Dans-Multi-Pro-External-Test/service/model"
	"errors"
	"log"
	"strings"
)

type userUsecase struct {
	userRepo service.IUserRepository
	toolRepo service.IToolsRepository
}

func NewUserUsecase(
	userRepo service.IUserRepository,
	toolRepo service.IToolsRepository,
) service.IUserUsecase {
	return userUsecase{
		userRepo: userRepo,
		toolRepo: toolRepo,
	}
}

func (u userUsecase) RegisterUser(args model.User) (err error) {
	var hashedPass string

	//validate buyer
	if len(args.Name) <= 0 {
		err = errors.New("invalid fullname")
		log.Print(err)
		return err
	}

	if len(args.PhoneNumber) <= 0 {
		err = errors.New("invalid phone number")
		log.Print(err)
		return err
	}

	if len(args.Password) <= 0 {
		err = errors.New("invalid password")
		log.Print(err)
		return err
	}

	email := strings.TrimSpace(args.Email)
	if len(email) <= 0 {
		err = errors.New("invalid email")
		log.Print(err)
		return err
	}

	emailValidation, err := u.toolRepo.CheckEmailValidation(email)
	if err != nil {
		err = errors.New("invalid email")
		return err
	}

	if !emailValidation {
		err = errors.New("email is invalid")
		log.Print(err)
		return err
	}

	if strings.TrimSpace(args.Password) != "" {
		hashedPass, err = u.toolRepo.SaltAndHash(args.Password)
		if err != nil {
			log.Print(err)
			err = errors.New(" while hashed password")
			return err
		}
	}

	args.Email = email
	args.Password = hashedPass

	err = u.userRepo.RegisterUser(args)
	if err != nil {
		return err
	}

	return err
}

func (u userUsecase) AuthUser(args model.LoginRequest) (res model.LoginResponses, err error) {
	if len(args.Username) <= 0 && len(args.Password) <= 0 {
		err = errors.New("invalid username or password")
		return res, err
	}

	res, err = u.userRepo.LoginUser(args)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, err
}
