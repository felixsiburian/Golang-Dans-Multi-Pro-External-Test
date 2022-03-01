package handler

import (
	"Golang-Dans-Multi-Pro-External-Test/service"
	"Golang-Dans-Multi-Pro-External-Test/service/model"
	"encoding/json"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type UserHandler struct {
	e        *echo.Echo
	userCase service.IUserUsecase
}

func NewUserHandler(
	e *echo.Echo,
	userCase service.IUserUsecase,
) *UserHandler {
	return &UserHandler{
		e:        e,
		userCase: userCase,
	}
}

func (ox *UserHandler) RegisterUser(ec echo.Context) error {
	var form model.User
	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		log.Println(err)
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	err = ox.userCase.RegisterUser(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": err.Error(),
			"en": err.Error(),
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id": "Berhasil",
		"en": "Success",
	})
}

func (ox *UserHandler) AuthUser(ec echo.Context) error {
	var form model.LoginRequest

	err := json.NewDecoder(ec.Request().Body).Decode(&form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	res, err := ox.userCase.AuthUser(form)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	tokens := map[string]string{
		"access_token":  res.AccessToken,
		"refresh_token": res.RefreshToken,
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":    "Berhasil",
		"en":    "Success",
		"token": tokens,
	})
}
