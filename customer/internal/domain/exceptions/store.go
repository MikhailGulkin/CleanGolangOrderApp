package exceptions

import "errors"

var (
	ErrCustomerAlreadyExists = errors.New("customer already exists")
)
