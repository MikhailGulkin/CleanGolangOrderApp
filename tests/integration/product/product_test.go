package product

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/tests/integration/conftest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestProducts(t *testing.T) {
	server := conftest.StartServer()
	defer server.Done()
	entity := CreateValidProductEntity()
	model := CreateValidProductModel(entity)

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
	t.Run("Get Product By name test", func(t *testing.T) {
		conftest.CleanTables(server.DB)
		CreateProductInDB(&model, server.DB)

		resp, err := http.Get(fmt.Sprintf("%s/products/%s", server.URL, model.Name)) //nolint:bodyclose
		if err != nil {
			t.Fatalf("Expected no errorHandler, got %v", err)
		}
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		decoder := json.NewDecoder(resp.Body)
		var data product.Product
		_ = decoder.Decode(&data)

		assert.Equal(t, data.Name, model.Name)
	})
	t.Run("Not exist product name", func(t *testing.T) {
		conftest.CleanTables(server.DB)

		resp, err := http.Get(fmt.Sprintf("%s/products/%s", server.URL, model.Name)) //nolint:bodyclose
		if err != nil {
			t.Fatalf("Expected no errorHandler, got %v", err)
		}
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
	t.Run("Not exist product when update", func(t *testing.T) {
		conftest.CleanTables(server.DB)

		req, err := http.NewRequest(
			http.MethodPut,
			fmt.Sprintf("%s/products/%s/productName", server.URL, model.ID),
			strings.NewReader(fmt.Sprintf("{\"productName\": \"%s\"}", "testNameUpdateNotExist")),
		)
		if err != nil {
			t.Fatalf("Expected no errorHandler, got %v", err)
		}

		client := &http.Client{}
		resp, errClient := client.Do(req) //nolint:bodyclose
		if errClient != nil {
			t.Fatalf("Expected no errorHandler, got %v", errClient)
		}
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
	t.Run("Success update product name", func(t *testing.T) {
		conftest.CleanTables(server.DB)
		CreateProductInDB(&model, server.DB)

		req, err := http.NewRequest(
			http.MethodPut,
			fmt.Sprintf("%s/products/%s/productName", server.URL, model.ID),
			strings.NewReader(fmt.Sprintf("{\"productName\": \"%s\"}", "testNameUpdateSuccess")),
		)
		if err != nil {
			t.Fatalf("Expected no errorHandler, got %v", err)
		}

		client := &http.Client{}
		resp, errClient := client.Do(req) //nolint:bodyclose
		if errClient != nil {
			t.Fatalf("Expected no errorHandler, got %v", errClient)
		}
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})
}
