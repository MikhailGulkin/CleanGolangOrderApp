package services

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/customer/aggregate"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/customer/vo"
	"github.com/google/uuid"
)

type Service struct {
}

func (Service) CreateUser(
	firstName string,
	middleName string,
	lastName string,
	addressID uuid.UUID,
) (aggregate.Customer, error) {
	fullName, err := vo.FullName{}.Create(firstName, middleName, lastName)
	if err != nil {
		return aggregate.Customer{}, err
	}
	customer := aggregate.Create(fullName, addressID)
	return customer, nil
}
