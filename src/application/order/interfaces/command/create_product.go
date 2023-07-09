package command

type CreateProductCommand struct {
	Price       float64 `json:"price" binding:"required"`
	Discount    int32   `json:"discount" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Name        string  `json:"name" binding:"required"`
}
type CreateProduct interface {
	Create(command CreateProductCommand) error
}
