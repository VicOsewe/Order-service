package application

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
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

func NewUUID() uuid.UUID {
	uuid, _ := uuid.NewUUID()
	return uuid
}
