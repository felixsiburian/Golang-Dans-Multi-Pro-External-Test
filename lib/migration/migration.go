package migration

import (
	db2 "Golang-Dans-Multi-Pro-External-Test/lib/db"
	"Golang-Dans-Multi-Pro-External-Test/service/model"
	"errors"
)

func InitTable() (err error) {
	db := db2.ConnectionGorm()

	res := db.AutoMigrate(&model.User{})
	if res.Error != nil {
		err = errors.New("[migration][initTable] while migrate table")
		return err
	}

	defer db.Close()
	return err
}
