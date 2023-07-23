package consts

type SagaStatus string

const (
	Pending  SagaStatus = "Pending"
	Rejected SagaStatus = "Rejected"
	Approved SagaStatus = "Approved"
)
