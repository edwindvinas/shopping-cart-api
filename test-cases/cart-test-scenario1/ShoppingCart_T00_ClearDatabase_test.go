//This test will clear database prior to executing all tests
package shoppingCart

import (
	"net/http"
	"testing"
	"fmt"
	"github.com/gavv/httpexpect"
)

const (
	API_URL = `https://ulapph-cloud-sca.appspot.com`
)

func TestClearDatabase(t *testing.T) {
	
	e := httpexpect.New(t, API_URL)
	
	printComment("CLR0001", "Test clear the database tables in appengine --> Products")
	e.DELETE("/products").
		Expect().
		Status(http.StatusOK)

}

func TestClearDatabase2(t *testing.T) {
	
	e := httpexpect.New(t, API_URL)
	
	printComment("CLR0002", "Test clear the database tables in appengine --> Cart")
	e.DELETE("/cart").
		Expect().
		Status(http.StatusOK)

}


func printComment(id, msg string) {
	fmt.Printf("[%v] %v\n", id, msg)
}