package controllers

import (
	"fmt"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/guillaume-boutin/frameworkless-golang/lib"
)

type homeController struct{}

var HomeController = homeController{}

func (hc homeController) Index(c *routing.Context) error {
	// env := (c.Get("Env")).(map[string]string)
	// fmt.Println(env["APP_NAME"])

	// fmt.Println(c.Query("name"))

	appName := lib.Env("APP_NAME")
	fmt.Println(appName)

	return c.Write("Hello world!")
}
