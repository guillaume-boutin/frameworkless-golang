package controllers

import (
	routing "github.com/go-ozzo/ozzo-routing"
)

type productsController struct{}

var ProductsController = productsController{}

func (hc productsController) Index(c *routing.Context) error {
	return c.Write("This is the products page!")
}
