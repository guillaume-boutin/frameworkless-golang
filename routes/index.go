package routes

import (
	routing "github.com/go-ozzo/ozzo-routing"
	ctrl "github.com/guillaume-boutin/frameworkless-golang/app/controllers"
	"github.com/guillaume-boutin/frameworkless-golang/middleware"
)

func Index(r *routing.Router) {

	r.Use(middleware.Env)
	r.Get("/", ctrl.HomeController.Index)
}

func Register(r *routing.Router) {
	Index(r)
	Products(r)
}
