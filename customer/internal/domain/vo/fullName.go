package vo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/exceptions"
	"regexp"
)

var NameRegex = regexp.MustCompile(`^[A-Z][a-z]{1,254}$`)

type FullName struct {
	FirstName  string
	MiddleName string
	LastName   string
}

// NewFullName creates new FullName value object
func NewFullName(firstName, middleName, lastName string) (FullName, error) {
	if !NameRegex.MatchString(firstName) {
		return FullName{}, exceptions.InvalidFullName
	}
	if !NameRegex.MatchString(middleName) {
		return FullName{}, exceptions.InvalidFullName
	}
	if !NameRegex.MatchString(lastName) {
		return FullName{}, exceptions.InvalidFullName
	}

	return FullName{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
	}, nil
}
