package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func SetUpDB() error {

	por := getEnv("DBPORT")
	password := getEnv("DBPASSWORD")
	user := getEnv("DBUSER")
	dbname := getEnv("DBNAME")
	host := getEnv("DBHOST")
	port, err := strconv.Atoi(por)
	if err != nil {
		return fmt.Errorf("failed to get port with error: %v", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return fmt.Errorf("failed to open database:%v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to establish a connection to the database:%v", err)
	}
	return nil

}

func getEnv(envName string) string {
	envVar := os.Getenv(envName)
	if envVar == "" {
		msg := fmt.Sprintf("environment variable %s not found", envName)
		log.Panicf(msg)
		os.Exit(1)
	}

	return envVar
}
