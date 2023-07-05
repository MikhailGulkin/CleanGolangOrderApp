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
		resp, err := http.Post(fmt.Sprintf("%s/products", server.URL), "application/json", CreateValidProduct()) //nolint:bodyclose

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
		conftest.CleanTables(server.DB)
	})

}
