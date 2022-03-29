package application

import (
	"fmt"
	"log"
	"os"
)

func GetEnv(envName string) string {
	envVar := os.Getenv(envName)
	if envVar == "" {
		msg := fmt.Sprintf("environment variable %s not found", envName)
		log.Panicf(msg)
		os.Exit(1)
	}

	return envVar
}
