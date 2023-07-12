package product

import "github.com/google/uuid"

type ProductID struct {
	Value uuid.UUID
}

func GetProductIDs(productIDs []uuid.UUID) []ProductID {
	ids := make([]ProductID, len(productIDs))
	for i, productID := range productIDs {
		ids[i] = ProductID{Value: productID}
	}
	return ids
}
func (id ProductID) ToString() string {
	return id.Value.String()
}
