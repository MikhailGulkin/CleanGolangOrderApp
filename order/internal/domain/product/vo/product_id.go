package vo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/common/id"
	"github.com/google/uuid"
)

type ProductID struct {
	id.ID
	Value uuid.UUID
}

func GetProductIDs(productIDs []uuid.UUID) []id.ID {
	ids := make([]id.ID, len(productIDs))
	for i, productID := range productIDs {
		ids[i] = ProductID{Value: productID}
	}
	return ids
}
func (id ProductID) ToString() string {
	return id.Value.String()
}
