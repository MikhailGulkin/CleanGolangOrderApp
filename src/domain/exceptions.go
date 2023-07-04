package domain

type CustomException struct {
	Message string
	Ctx     string
}

func (e *CustomException) Error() string {
	return e.Message + " " + e.Ctx
}

type InvalidUUIDCreation struct {
	CustomException
	uuidError string
}

func (e InvalidUUIDCreation) Exception(context string, uuidError string) InvalidUUIDCreation {
	return InvalidUUIDCreation{
		CustomException{
			Message: "When create uuid error occurred",
			Ctx:     context,
		},
		uuidError,
	}
}
func (e *InvalidUUIDCreation) Error() string {
	return e.CustomException.Error() + e.uuidError
}
