package postgres

import (
	"fmt"
	"log"

	"github.com/VicOsewe/Order-service/application"
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

func NewOrderService() *OrderService {
	m := OrderService{
		DB: InitializeDatabase(),
	}
	return &m

}
func InitializeDatabase() *gorm.DB {
	password := application.GetEnv("DBPASSWORD")
	user := application.GetEnv("DBUSER")
	dbname := application.GetEnv("DBNAME")
	host := application.GetEnv("DBHOST")
	port := application.GetEnv("DBPORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connected to the database successfully")

	db.AutoMigrate(&domain.Customer{})
	return db

}

//CreateCustomer creates a record of a customer in the database
func (db *OrderService) CreateCustomer(customer *domain.Customer) (*domain.Customer, error) {
	customer.ID = application.NewUUID()
	if err := db.DB.Create(customer).Error; err != nil {
		return nil, fmt.Errorf(
			"can't create a new marketing record: err: %v",
			err,
		)
	}
	return customer, nil
}
