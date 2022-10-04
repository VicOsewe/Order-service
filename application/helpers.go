package application

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
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

func GenerateOTP() (string, error) {
	totpOpt := totp.GenerateOpts{
		Issuer:      GetEnv("OTP_ISSUER"),
		AccountName: GetEnv("OTP_ACCOUNTNAME"),
	}
	key, err := totp.Generate(totpOpt)
	if err != nil {
		return "", fmt.Errorf("failed to generate key: %v", err)
	}

	code, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		return "", fmt.Errorf("failed to generate code: %v", err)
	}
	return code, nil
}

// func
