package router

import (
	"Golang-Dans-Multi-Pro-External-Test/service"
	"Golang-Dans-Multi-Pro-External-Test/service/delivery/handler"
	"github.com/labstack/echo"
)

func NewRouter(
	e *echo.Echo,
	userCase service.IUserUsecase,
	jobCase service.IJobUsecase,
) {
	u := handler.NewUserHandler(e, userCase)
	j := handler.NewJobHandler(e, jobCase)

	r := e.Group("v1")

	r.POST("/user/register", u.RegisterUser)
	r.POST("/user/auth", u.AuthUser)

	r.GET("/job/get", j.GetJobList)
	r.GET("/job/get/:id", j.GetJobDetaiil)
}
