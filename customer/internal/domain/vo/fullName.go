package vo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/exceptions"
)

type FullName struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func (FullName) Create(firstName, middleName, lastName string) (FullName, error) {
	if firstName == "" {
		err := exceptions.EmptyName{}.Exception("First name cannot be empty", firstName)
		return FullName{}, &err
	}
	if middleName == "" {
		err := exceptions.EmptyName{}.Exception("First name cannot be empty", middleName)
		return FullName{}, &err
	}
	if lastName == "" {
		err := exceptions.EmptyName{}.Exception("First name cannot be empty", lastName)
		return FullName{}, &err
	}

	return FullName{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
	}, nil
}
