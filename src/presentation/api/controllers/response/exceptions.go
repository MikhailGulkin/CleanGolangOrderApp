package response

type ExceptionResponse struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}
