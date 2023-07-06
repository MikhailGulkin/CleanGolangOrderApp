package product

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/tests/integration/conftest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestProducts(t *testing.T) {
	server := conftest.StartServer()
	defer server.Done()

	t.Run("Success create product test", func(t *testing.T) {
		conftest.CleanTables(server.DB)

		resp, err := http.Post(fmt.Sprintf("%s/products", server.URL), "application/json", CreateValidByteProduct()) //nolint:bodyclose

		if err != nil {
			t.Fatalf("Expected no errorHandler, got %v", err)
		}
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})
	t.Run("Invalid discount test", func(t *testing.T) {
		conftest.CleanTables(server.DB)

		resp, err := http.Post(fmt.Sprintf("%s/products", server.URL), "application/json", CreateInvalidDiscountByteProduct()) //nolint:bodyclose

		if err != nil {
			t.Fatalf("Expected no errorHandler, got %v", err)
		}
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
	t.Run("Invalid price test", func(t *testing.T) {
		conftest.CleanTables(server.DB)

		resp, err := http.Post(fmt.Sprintf("%s/products", server.URL), "application/json", CreateInvalidPriceByteProduct()) //nolint:bodyclose

		if err != nil {
			t.Fatalf("Expected no errorHandler, got %v", err)
		}
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
