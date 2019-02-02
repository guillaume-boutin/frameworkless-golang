package main

import (
	"github.com/guillaume-boutin/frameworkless-golang/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, _ := gorm.Open("mysql", "root:root@/frameworkless_golang?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	db.CreateTable(&models.User{})

}
