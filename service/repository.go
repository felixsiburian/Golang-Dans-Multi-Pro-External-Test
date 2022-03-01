package service

import "Golang-Dans-Multi-Pro-External-Test/service/model"

type IUserRepository interface {
	RegisterUser(args model.User) (err error)
	LoginUser(args model.LoginRequest) (res model.LoginResponses, err error)
}

type IToolsRepository interface {
	SaltAndHash(pwd string) (s string, err error)
	CheckEmailValidation(email string) (res bool, err error)
	VerifyPassword(hashedPwd, pwd string) (err error)
	GUID() string
}

type ITokenRepository interface {
	CreateUserToken(args model.TokenArgs) (res model.LoginResponses, err error)
}
