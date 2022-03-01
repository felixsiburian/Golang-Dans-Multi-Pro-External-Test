package main

import (
	"Golang-Dans-Multi-Pro-External-Test/lib/migration"
	"Golang-Dans-Multi-Pro-External-Test/service/config"
	"Golang-Dans-Multi-Pro-External-Test/service/delivery/router"
	"Golang-Dans-Multi-Pro-External-Test/service/repository"
	"Golang-Dans-Multi-Pro-External-Test/service/tools"
	"Golang-Dans-Multi-Pro-External-Test/service/usecase"
	"fmt"
	"github.com/labstack/echo"
	"os"
)

func main() {
	app := config.Config{}

	app.CatchError(app.InitEnv())
	app.CatchError(migration.InitTable())

	Start()
}

func Start() {
	e := echo.New()

	//	reg repos
	toolRepo := tools.NewToolRepository()
	tokenRepo := tools.NewTokenRepository(toolRepo)
	userRepo := repository.NewUserRepository(toolRepo, tokenRepo)

	//	register usecase
	userUsecase := usecase.NewUserUsecase(userRepo, toolRepo)
	jobUsecase := usecase.NewJobUsecase()

	router.NewRouter(e, userUsecase, jobUsecase)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s%v", os.Getenv("APP_HOST"), ":", os.Getenv("APP_PORT"))))
}
