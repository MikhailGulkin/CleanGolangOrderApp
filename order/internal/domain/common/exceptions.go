package common

type CustomException struct {
	Message string
	Ctx     string
}

func (e *CustomException) Error() string {
	return e.Message + " " + e.Ctx
}
