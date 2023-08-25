package vo

type FullName struct {
	FirstName  string
	MiddleName string
	LastName   string
}

// NewFullName creates new FullName value object
func NewFullName(firstName, middleName, lastName string) (FullName, error) {
	if firstName == "" {
		// TODO: add exception
		return FullName{}, nil
	}
	if middleName == "" {
		// TODO: add exception
		return FullName{}, nil
	}
	if lastName == "" {
		// TODO: add exception
		return FullName{}, nil
	}

	return FullName{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
	}, nil
}
