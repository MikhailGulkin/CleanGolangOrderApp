package reader

import "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/models"

func CalculateTotalProductPrice(product models.Product) float64 {
	return product.Price - ((float64(product.Discount) / 100) * 100)
}
func CalculateTotalOrderPrice(products []models.Product) float64 {
	var totalPrice float64
	for _, product := range products {
		totalPrice += CalculateTotalProductPrice(product)
	}
	return totalPrice
}
