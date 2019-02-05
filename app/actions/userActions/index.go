package actions

import (
	"github.com/guillaume-boutin/frameworkless-golang/app/models"
	"github.com/guillaume-boutin/frameworkless-golang/app/resources"
	"github.com/guillaume-boutin/frameworkless-golang/lib"
)

type userIndex struct {
	lib.Action
}

var UserIndex = userIndex{}

func (a userIndex) Run() []resources.UserResource {
	db, _ := lib.Db()
	defer db.Close()

	users := models.Users{}
	db.Find(&users)

	return users.Resource()
}
