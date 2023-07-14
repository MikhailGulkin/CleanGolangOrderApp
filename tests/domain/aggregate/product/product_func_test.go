package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/product/vo"
	"github.com/google/uuid"
	"testing"
)

func TestGetAllProducts(t *testing.T) {
	ids := []uuid.UUID{uuid.New(), uuid.New()}
	for index, id := range vo.GetProductIDs(ids) {
		if id.ToString() != ids[index].String() {
			t.Error("Incorrect function work")
		}
	}
}
