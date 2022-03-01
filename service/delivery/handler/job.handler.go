package handler

import (
	"Golang-Dans-Multi-Pro-External-Test/service"
	"Golang-Dans-Multi-Pro-External-Test/service/model"
	"github.com/labstack/echo"
	"net/http"
)

type JobUsecaseHandler struct {
	e          *echo.Echo
	jobUsecase service.IJobUsecase
}

func NewJobHandler(
	e *echo.Echo,
	jobUsecase service.IJobUsecase,
) *JobUsecaseHandler {
	return &JobUsecaseHandler{
		e:          e,
		jobUsecase: jobUsecase,
	}
}

func (ox *JobUsecaseHandler) GetJobList(ec echo.Context) error {
	description := ec.QueryParam("description")
	location := ec.QueryParam("location")
	fullTime := ec.QueryParam("full_time")

	jobArgs := model.JobRequest{
		Description: description,
		Location:    location,
		FullTime:    fullTime,
	}

	res, err := ox.jobUsecase.GetJobList(jobArgs)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":   "Berhasil",
		"en":   "Success",
		"data": res,
	})
}

func (ox *JobUsecaseHandler) GetJobDetaiil(ec echo.Context) error {
	id := ec.Param("id")

	res, err := ox.jobUsecase.GetJobDetail(id)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, map[string]interface{}{
			"id": "Kesalahan sedang terjadi. Silahkan Ulangi Beberapa saat lagi",
			"en": "Something went wrong. Please try again later",
		})
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id":   "Berhasil",
		"en":   "Success",
		"data": res,
	})

}
