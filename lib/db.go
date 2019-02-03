package lib

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Db() (*gorm.DB, error) {
	dialect := Env("DB_CONNECTION")

	params := fmt.Sprintf(
		"%v:%v@/%v?charset=utf8&parseTime=True&loc=Local",
		Env("DB_USERNAME"),
		Env("DB_PASSWORD"),
		Env("DB_DATABASE"),
	)

	fmt.Println(params)

	db, err := gorm.Open(dialect, params)
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	return db, err
}
