package main

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/infrastructure/db/models"
	"github.com/MikhailGulkin/simpleGoOrderApp/pkg/customer/servicespb"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func main() {
	//conn, err := grpc.Dial("L")
	//var conf config.Config
	//load.LoadConfig(&conf, "", "")
	//conn := db.BuildConnection(conf.DBConfig)
	//customerRepo := CustomerRepository{db: conn}
	//fmt.Println(customerRepo)
}

type CustomerService struct {
	servicespb.CustomerServiceServer
}

type CustomerRepository struct {
	db *gorm.DB
}

func (r *CustomerRepository) GetCustomerByID(id uuid.UUID) (models.Customer, error) {
	var customer models.Customer
	err := r.db.First(&customer, id).Error
	return customer, err
}

func (r *CustomerRepository) GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}
func (r *CustomerRepository) CreateCustomer(customer *models.Customer) error {
	return r.db.Create(customer).Error
}
