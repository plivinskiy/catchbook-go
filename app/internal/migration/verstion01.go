package migration

import (
	"catchbook/internal/model"
	"fmt"
	"gorm.io/gorm"
)

type V01 struct {
	Db *gorm.DB
}

func (v V01) Migrate() {
	var u model.User
	var f model.Fish
	var a model.Address
	err := v.Db.Migrator().AutoMigrate(&u, &f, &a)
	fmt.Println("migrate V01")
	if err != nil {
		fmt.Println(err.Error())
	}
}
