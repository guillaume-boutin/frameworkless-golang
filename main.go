package main

import (
	"fmt"
	"log"
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/guillaume-boutin/frameworkless-golang/routes"
)

func handler(c *routing.Context) error {
	env := (c.Get("Env")).(map[string]string)
	fmt.Println(env["APP_NAME"])

	fmt.Println(c.Query("name"))

	return c.Write("Hello world!")
}

func main() {

	router := routing.New()
	routes.Register(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
