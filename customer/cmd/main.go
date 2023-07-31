package main

// import (
//
//	"context"
//	"fmt"
//	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/infrastructure/db/models"
//	"github.com/MikhailGulkin/simpleGoOrderApp/pkg/customer/servicespb"
//	"github.com/google/uuid"
//	"google.golang.org/grpc"!Ё
//	"google.golang.org/grpc/reflection"
//	"gorm.io/gorm"
//	"net"
//
// )
func main() {
	//lis, _ := net.Listen("tcp", fmt.Sprintf(":%d", 50052))
	//
	//serv := grpc.NewServer()
	//servicespb.RegisterCustomerServiceServer(serv, &CustomerService{})
	//reflection.Register(serv)
	//if err := serv.Serve(lis); err != nil {
	//	log.Fatalf("failed to server: %v", err)
	//}
	//	//var conf config.Config
	//	//load.LoadConfig(&conf, "", "")
	//	//conn := db.BuildConnection(conf.DBConfig)
	//	//customerRepo := CustomerRepository{db: conn}
	//	//fmt.Println(customerRepo)
}

//type CustomerService struct {
//	servicespb.CustomerServiceServer
//}
//
//func (s *CustomerService) CreateCustomer(
//	_ context.Context, request *servicespb.CreateCustomerRequest,
//) (*servicespb.CreateCustomerResponse, error) {
//	var response servicespb.CreateCustomerResponse
//
//	newCustomer := models.Customer{
//		Base:        models.Base{ID: uuid.New()},
//		FirstName:   request.FirstName,
//		LastName:    request.LastName,
//		PhoneNumber: request.PhoneNumber,
//		Email:       request.Email,
//	}
//	fmt.Println(newCustomer.ID, newCustomer.Email, request.Email)
//	response.Id = newCustomer.ID.String()
//
//	return &response, nil
//}

//
//type CustomerRepository struct {
//	db *gorm.DB
//}
//
//func (r *CustomerRepository) GetCustomerByID(id uuid.UUID) (models.Customer, error) {
//	var customer models.Customer
//	err := r.db.First(&customer, id).Error
//	return customer, err
//}
//
//func (r *CustomerRepository) GetAllCustomers() ([]models.Customer, error) {
//	var customers []models.Customer
//	err := r.db.Find(&customers).Error
//	return customers, err
//}
//func (r *CustomerRepository) CreateCustomer(customer *models.Customer) error {
//	return r.db.Create(customer).Error
//}
