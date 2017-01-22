//This test will add the products to the catalog by calling API
package shoppingCart

import (
	"net/http"
	"testing"
	"github.com/gavv/httpexpect"
)

//ult_small Unlimited 1GB $24.90

func TestAddProductToCatalog(t *testing.T) {
	
	e := httpexpect.New(t, API_URL)
	
	printComment("ADD0001", "Test add product 1 to catalog --> ult_small Unlimited 1GB $24.90")
	product := map[string]interface{}{
		"code": "ult_small",
		"name": "Unlimited 1GB",
		"price": 24.90,
	}

	e.POST("/products").
		WithJSON(product).
		Expect().
		Status(http.StatusOK)

}

//ult_medium Unlimited 2GB $29.90
func TestAddProductToCatalog2(t *testing.T) {
	
	e := httpexpect.New(t, API_URL)
	
	printComment("ADD0002", "Test add product 2 to catalog --> ult_medium Unlimited 2GB $29.90")
	product := map[string]interface{}{
		"code": "ult_medium",
		"name": "Unlimited 2GB",
		"price": 29.90,
	}

	e.POST("/products").
		WithJSON(product).
		Expect().
		Status(http.StatusOK)

}

//ult_large Unlimited 5GB $44.90
func TestAddProductToCatalog3(t *testing.T) {
	
	e := httpexpect.New(t, API_URL)
	
	printComment("ADD0003", "Test add product 3 to catalog --> ult_large Unlimited 5GB $44.90")
	product := map[string]interface{}{
		"code": "ult_large",
		"name": "Unlimited 5GB",
		"price": 44.90,
	}

	e.POST("/products").
		WithJSON(product).
		Expect().
		Status(http.StatusOK)

}

//1gb 1 GB Data-pack $9.90
func TestAddProductToCatalog4(t *testing.T) {
	
	e := httpexpect.New(t, API_URL)
	
	printComment("ADD0004", "Test add product 4 to catalog --> 1gb 1 GB Data-pack $9.90")
	product := map[string]interface{}{
		"code": "1gb",
		"name": "1 GB Data-pack",
		"price": 9.90,
	}

	e.POST("/products").
		WithJSON(product).
		Expect().
		Status(http.StatusOK)

}
