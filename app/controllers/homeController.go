package controllers

import (
	"fmt"

	routing "github.com/go-ozzo/ozzo-routing"
)

type homeController struct{}

var HomeController = homeController{}

func (hc homeController) Index(c *routing.Context) error {
	env := (c.Get("Env")).(map[string]string)
	fmt.Println(env["APP_NAME"])

	fmt.Println(c.Query("name"))

	return c.Write("Hello world!")
}
