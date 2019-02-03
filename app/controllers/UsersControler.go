package controllers

import (
	"encoding/json"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/guillaume-boutin/frameworkless-golang/app/models"
	"github.com/guillaume-boutin/frameworkless-golang/lib"
)

type usersController struct{}

var UsersController = usersController{}

func (hc usersController) Index(c *routing.Context) error {
	db, _ := lib.Db()
	defer db.Close()

	users := models.Users{}
	db.Find(&users)

	collection := users.Resource()

	res, _ := json.Marshal(&collection)

	return c.Write(res)
}
