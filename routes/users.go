package routes

import (
	routing "github.com/go-ozzo/ozzo-routing"
	ctrl "github.com/guillaume-boutin/frameworkless-golang/app/controllers"
)

func Users(r *routing.Router) {

	pr := r.Group("/users")

	pr.Get("", ctrl.UsersController.Index)
}
