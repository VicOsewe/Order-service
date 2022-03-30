package postgres

import (
	"fmt"
	"log"

	"github.com/VicOsewe/Order-service/application"
	"github.com/VicOsewe/Order-service/domain"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

//OrderService ...
type OrderService struct {
	DB *gorm.DB
}

//New ..
func New(db *gorm.DB) OrderService {
	return OrderService{db}
}

//NewOrderService ...
func NewOrderService() *OrderService {
	m := OrderService{
		DB: InitializeDatabase(),
	}
	return &m

}

//InitializeDatabase ...
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

	db.AutoMigrate(&domain.Customer{}, &domain.Order{}, &domain.Product{}, &domain.OrderProduct{})
	return db

}

//CreateCustomer creates a record of a customer in the database
func (db *OrderService) CreateCustomer(customer *domain.Customer) (*domain.Customer, error) {
	customer.ID = uuid.New().String()
	if err := db.DB.Create(customer).Error; err != nil {
		return nil, fmt.Errorf(
			"can't create customer record: err: %v",
			err,
		)
	}
	return customer, nil
}

//CreateProduct creates a record of a product in the database
func (db *OrderService) CreateProduct(product *domain.Product) (*domain.Product, error) {
	product.ID = uuid.New().String()
	if err := db.DB.Create(product).Error; err != nil {
		return nil, fmt.Errorf(
			"can't create a product record: err: %v",
			err,
		)
	}

	return product, nil
}

//CreateOrder creates a record of an order int he database
func (db *OrderService) CreateOrder(order *domain.Order, orderProducts *[]domain.OrderProduct) (*domain.Order, error) {
	order.ID = uuid.New().String()
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		for _, orderProduct := range *orderProducts {
			orderProduct.ID = uuid.New().String()
			orderProduct.OrderID = order.ID
			if err := tx.Create(&orderProduct).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return order, nil
}

//GetCustomerByID retrieves a record of a customer in the database using the customer id
func (db *OrderService) GetCustomerByID(customerID string) (*domain.Customer, error) {
	customer := domain.Customer{}
	if err := db.DB.Where(
		&domain.Customer{ID: customerID}).
		Find(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

//GetProductByID ...
func (db *OrderService) GetProductByID(productID string) (*domain.Product, error) {
	product := domain.Product{}
	if err := db.DB.Where(
		&domain.Product{ID: productID}).
		Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

//GetCustomerByPhoneNumber ...
func (db *OrderService) GetCustomerByPhoneNumber(phoneNumber string) (*domain.Customer, error) {
	customer := domain.Customer{}
	if err := db.DB.Where(
		&domain.Customer{PhoneNumber: phoneNumber}).
		Find(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

//GetProductByName ...
func (db *OrderService) GetProductByName(name string) (*domain.Product, error) {
	product := domain.Product{}
	if err := db.DB.Where(
		&domain.Product{Name: name}).
		Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

//GetAllCustomerOrdersByCustomerID ...
func (db *OrderService) GetAllCustomerOrdersByCustomerID(customerID string) (*[]domain.Order, error) {
	orders := []domain.Order{}
	if err := db.DB.Where(
		&domain.Order{
			CustomerID: customerID,
		}).
		Find(&orders).
		Error; err != nil {
		return nil, err
	}

	return &orders, nil
}

//GetAllProducts ...
func (db *OrderService) GetAllProducts() (*[]domain.Product, error) {

	products := []domain.Product{}
	if err := db.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}
