package routes

import (
	routing "github.com/go-ozzo/ozzo-routing"
	ctrl "github.com/guillaume-boutin/frameworkless-golang/app/controllers"
)

func Products(r *routing.Router) {

	pr := r.Group("/products")

	pr.Get("", ctrl.ProductsController.Index)
}
