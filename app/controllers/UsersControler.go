package controllers

import (
	"fmt"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/guillaume-boutin/frameworkless-golang/app/models"
	"github.com/guillaume-boutin/frameworkless-golang/lib"
)

type usersController struct{}

var UsersController = usersController{}

func (hc usersController) Index(c *routing.Context) error {
	db, _ := lib.Db()
	defer db.Close()

	users := []models.User{}
	db.Find(&users)
	fmt.Println(users)

	return c.Write("This is the users page!")
}
