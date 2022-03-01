package service

import "Golang-Dans-Multi-Pro-External-Test/service/model"

type IUserUsecase interface {
	RegisterUser(args model.User) (err error)
	AuthUser(args model.LoginRequest) (res model.LoginResponses, err error)
}

type IJobUsecase interface {
	GetJobList(args model.JobRequest) (res model.JobResponse, err error)
	GetJobDetail(id string) (res model.JobElement, err error)
}
