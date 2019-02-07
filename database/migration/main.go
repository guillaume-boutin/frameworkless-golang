package main

import (
	"log"

	gormigrate "gopkg.in/gormigrate.v1"

	"github.com/guillaume-boutin/frameworkless-golang/lib"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := lib.Db()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.LogMode(true)

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "201802062142",
			Migrate: func(tx *gorm.DB) error {
				// it's a good pratice to copy the struct inside the function,
				// so side effects are prevented if the original struct changes during the time
				type User struct {
					gorm.Model
					Email        string `gorm:"size:255;UNIQUE;NOT NULL"`
					HashPassword string `gorm:"size:255;NOT NULL"`
				}
				return tx.AutoMigrate(&User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users").Error
			},
		},
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")

}
