package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/guillaume-boutin/frameworkless-golang/bootstrap"
	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	APP_NAME := os.Getenv("APP_NAME")
	fmt.Println(APP_NAME)

	ctn := bootstrap.Container()
	test := ctn.Get("test")
	fmt.Println(test)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
