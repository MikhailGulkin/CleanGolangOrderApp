package command

type CreateProductCommand struct {
	Price       float64
	Discount    int32
	Description string
	Name        string
}
type CreateProduct interface {
	Create(command CreateProductCommand) error
}
