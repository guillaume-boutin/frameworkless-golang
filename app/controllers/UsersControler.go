package controllers

import (
	routing "github.com/go-ozzo/ozzo-routing"
	actions "github.com/guillaume-boutin/frameworkless-golang/app/actions/userActions"
	"github.com/guillaume-boutin/frameworkless-golang/lib"
)

type usersController struct {
	lib.Controller
}

var UsersController = usersController{}

func (ctrl usersController) Index(c *routing.Context) error {
	body := actions.UserIndex.Run()

	return ctrl.Respond(c).StatusCode(200).Json(&body)
}
