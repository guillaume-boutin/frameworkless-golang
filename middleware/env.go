package middleware

import (
	"log"
	"os"
	"strings"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/joho/godotenv"
)

func Env(c *routing.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

		return err
	}

	environ := os.Environ()
	envMap := map[string]string{}

	for _, str := range environ {
		fields := strings.Split(str, "=")
		envMap[fields[0]] = fields[1]
	}

	c.Set("Env", envMap)

	return c.Next()
}
