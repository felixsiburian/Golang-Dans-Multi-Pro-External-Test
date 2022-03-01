package usecase

import (
	"Golang-Dans-Multi-Pro-External-Test/service"
	"Golang-Dans-Multi-Pro-External-Test/service/model"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type jobUsecase struct {
}

func NewJobUsecase() service.IJobUsecase {
	return jobUsecase{}
}

func (j jobUsecase) GetJobList(args model.JobRequest) (res model.JobResponse, err error) {
	var url string

	if len(args.Description) > 0 && len(args.Location) > 0 && len(args.FullTime) > 0 {
		url = fmt.Sprintf("%s?description=%s&location=%s&full_time=%s", os.Getenv("jobsUrl"), args.Description, args.Location, args.FullTime)
	}

	if len(args.Description) > 0 && len(args.Location) > 0 && len(args.FullTime) == 0 {
		url = fmt.Sprintf("%s?description=%s&location=%s", os.Getenv("jobsUrl"), args.Description, args.Location)
	}

	if len(args.Description) > 0 && len(args.Location) == 0 && len(args.FullTime) > 0 {
		url = fmt.Sprintf("%s?description=%s&full_time=%s", os.Getenv("jobsUrl"), args.Description, args.FullTime)
	}

	if len(args.Description) == 0 && len(args.Location) > 0 && len(args.FullTime) > 0 {
		url = fmt.Sprintf("%s?location=%s&full_time=%s", os.Getenv("jobsUrl"), args.Location, args.FullTime)
	}

	if len(args.Description) > 0 && len(args.Location) == 0 && len(args.FullTime) == 0 {
		url = fmt.Sprintf("%s?description=%s", os.Getenv("jobsUrl"), args.Description)
	}

	if len(args.Description) == 0 && len(args.Location) > 0 && len(args.FullTime) == 0 {
		url = fmt.Sprintf("%s?location=%s", os.Getenv("jobsUrl"), args.Location)
	}

	if len(args.Description) == 0 && len(args.Location) == 0 && len(args.FullTime) > 0 {
		url = fmt.Sprintf("%s?full_time=%s", os.Getenv("jobsUrl"), args.FullTime)
	}

	if len(args.Description) == 0 && len(args.Location) == 0 && len(args.FullTime) == 0 {
		url = fmt.Sprintf("%s", os.Getenv("jobsUrl"))
	}

	fmt.Println("URL : ", url)
	resp, err := http.Get(url)
	if err != nil {
		err = errors.New("while get data from url")
		return res, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return res, err
	}

	err = json.Unmarshal(responseData, &res)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, err
}

func (j jobUsecase) GetJobDetail(id string) (res model.JobElement, err error) {
	url := fmt.Sprintf(
		"%s/%s", os.Getenv("jobsDetailUrl"), id,
	)

	resp, err := http.Get(url)
	if err != nil {
		err = errors.New("while get data from url")
		return res, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return res, err
	}

	err = json.Unmarshal(responseData, &res)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, err
}
