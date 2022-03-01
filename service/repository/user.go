package repository

import (
	db2 "Golang-Dans-Multi-Pro-External-Test/lib/db"
	"Golang-Dans-Multi-Pro-External-Test/service"
	"Golang-Dans-Multi-Pro-External-Test/service/model"
	"Golang-Dans-Multi-Pro-External-Test/service/repository/queries"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type userRepository struct {
	toolRepo  service.IToolsRepository
	tokenRepo service.ITokenRepository
}

func NewUserRepository(
	toolRepo service.IToolsRepository,
	tokenRepo service.ITokenRepository,
) service.IUserRepository {
	return userRepository{
		toolRepo:  toolRepo,
		tokenRepo: tokenRepo,
	}
}

func (u userRepository) RegisterUser(args model.User) (err error) {
	db := db2.ConnectionGorm()

	args.IsActive = true
	args.CreatedDate = time.Now()
	args.CreatedBy = args.Name

	getUserbyEmail := db.Exec(queries.QueryGetBuyerByEmail, args.Email)
	if getUserbyEmail.RowsAffected > 0 {
		err = errors.New("user already exist")
		return err
	}

	if getUserbyEmail.Error != nil {
		err = errors.New(fmt.Sprintf("err while getUserbyEmail: %v", getUserbyEmail.Error.Error()))
		return err
	}

	res := db.Debug().Create(&args)
	if res.Error != nil {
		err = errors.New(fmt.Sprintf("failed while create user : %v", res.Error.Error()))
		return err
	}

	defer db.Close()
	return err
}

func (u userRepository) LoginUser(args model.LoginRequest) (res model.LoginResponses, err error) {
	var (
		user      model.User
		tokenArgs model.TokenArgs
		db        = db2.ConnectionGorm()
	)

	rows := db.Debug().Model(&model.User{}).Where("username = ?", args.Username).Take(&user)
	if rows.Error != nil {
		log.Println(rows.Error.Error())
		err = errors.New(fmt.Sprintf("wrong username or password"))
		return res, err
	}

	err = u.toolRepo.VerifyPassword(user.Password, args.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Print(err)
		err = errors.New("wrong username or password")
		return res, err
	}

	tokenArgs.ID = user.ID
	tokenArgs.Name = user.Name
	tokenArgs.Email = user.Email

	res, err = u.tokenRepo.CreateUserToken(tokenArgs)
	if err != nil {
		log.Print(err)
		err = errors.New(fmt.Sprintf("failed while create token : %v", err.Error()))
		return res, err
	}

	defer db.Close()
	return res, err
}
