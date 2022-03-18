package postgres

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/VicOsewe/Order-service/domain"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type OrderService struct {
	DB *gorm.DB
}

func New(db *gorm.DB) OrderService {
	return OrderService{db}
}

func InitializeDatabase() (*gorm.DB, error) {
	password := getEnv("DBPASSWORD")
	user := getEnv("DBUSER")
	dbname := getEnv("DBNAME")
	host := getEnv("DBHOST")
	port, err := strconv.Atoi(getEnv("DBPORT"))
	if err != nil {
		return nil, fmt.Errorf("failed to get port with error: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connected to the database successfully")

	db.AutoMigrate(&domain.Customer{})
	return db, nil

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
