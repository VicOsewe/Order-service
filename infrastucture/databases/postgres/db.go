package postgres

import (
	"fmt"
	"log"

	"github.com/VicOsewe/Order-service/application"
	"github.com/VicOsewe/Order-service/domain/dao"
	"github.com/google/uuid"
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

	db.AutoMigrate(&dao.Customer{}, &dao.Order{}, &dao.Product{}, &dao.OrderProduct{}, &dao.OTP{})
	return db

}

//SaveOTP persists the otp record
func (db *OrderService) SaveOTP(otp dao.OTP) error {
	if err := db.DB.Create(otp).Error; err != nil {
		return fmt.Errorf("error saving otp record: %v", err)
	}
	return nil
}

//CreateCustomer creates a record of a customer in the database
func (db *OrderService) CreateCustomer(customer *dao.Customer) (*dao.Customer, error) {
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
func (db *OrderService) CreateProduct(product *dao.Product) (*dao.Product, error) {
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
func (db *OrderService) CreateOrder(order *dao.Order, orderProducts *[]dao.OrderProduct) (*dao.Order, error) {
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
func (db *OrderService) GetCustomerByID(customerID string) (*dao.Customer, error) {
	customer := dao.Customer{}
	if err := db.DB.Where(
		&dao.Customer{ID: customerID}).
		Find(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (db *OrderService) GetProductByID(productID string) (*dao.Product, error) {
	product := dao.Product{}
	if err := db.DB.Where(
		&dao.Product{ID: productID}).
		Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (db *OrderService) GetCustomerByPhoneNumber(phoneNumber string) (*dao.Customer, error) {
	customer := dao.Customer{}
	if err := db.DB.Where(
		&dao.Customer{PhoneNumber: phoneNumber}).
		Find(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (db *OrderService) GetProductByName(name string) (*dao.Product, error) {
	product := dao.Product{}
	if err := db.DB.Where(
		&dao.Product{Name: name}).
		Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (db *OrderService) GetAllCustomerOrdersByCustomerID(customerID string) (*[]dao.Order, error) {
	orders := []dao.Order{}
	if err := db.DB.Where(
		&dao.Order{
			CustomerID: customerID,
		}).
		Find(&orders).
		Error; err != nil {
		return nil, err
	}

	return &orders, nil
}

func (db *OrderService) GetAllProducts() (*[]dao.Product, error) {

	products := []dao.Product{}
	if err := db.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

//UpdateCustomer updates the records of a customer
func (db *OrderService) UpdateCustomer(customer *dao.Customer) (*dao.Customer, error) {
	if err := db.DB.Where(&dao.Customer{PhoneNumber: customer.PhoneNumber}).Updates(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}
