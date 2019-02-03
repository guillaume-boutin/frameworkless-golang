package main

import (
	"log"
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/guillaume-boutin/frameworkless-golang/routes"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	router := routing.New()

	routes.Register(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
