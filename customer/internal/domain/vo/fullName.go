package vo

import "github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/exceptions"

type FullName struct {
	FirstName  string
	MiddleName string
	LastName   string
}

// NewFullName creates new FullName value object
func NewFullName(firstName, middleName, lastName string) (FullName, error) {
	if firstName == "" {
		return FullName{}, exceptions.InvalidFullNameLength
	}
	if middleName == "" {
		return FullName{}, exceptions.InvalidFullNameLength
	}
	if lastName == "" {
		return FullName{}, exceptions.InvalidFullNameLength
	}

	return FullName{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
	}, nil
}
