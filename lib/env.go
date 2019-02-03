package lib

import (
	"os"
)

func Env(key string) string {
	return os.Getenv(key)
}

func LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}
