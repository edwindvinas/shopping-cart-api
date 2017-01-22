//This is for Scenario 3
/*Scenario 	Items Added 			Expected Cart Total 		Expected Cart Items
3 			1 x Unlimited 1 GB	$84.70					1 x Unlimited 1 GB
			2 X Unlimited 2 GB							2 X Unlimited 2 GB
														2 X 1 GB Data-pack
*/

package shoppingCart

import (
	"net/http"
	"testing"
	"github.com/gavv/httpexpect"
)

//Add 1 item of  Unlimited 1 GB for $24.90
func TestCartScenario(t *testing.T) {
	
	e := httpexpect.New(t, API_URL)
	
	printComment("SC30001", "Test Add 1 item of  Unlimited 1 GB for $24.90")
	cart := map[string]interface{}{
		"code": "ult_small",
		"name": "Unlimited 1GB",
		"price": 24.90,
		"items": 1,
	}

	e.POST("/cart").
		WithJSON(cart).
		Expect().
		Status(http.StatusOK)

}

//Add 2 items of Unlimited 2GB for $29.90
func TestCartScenario2(t *testing.T) {
	
	e := httpexpect.New(t, API_URL)
	
	printComment("SC30001", "Test Add 2 items of Unlimited 2GB for $29.90")
	cart := map[string]interface{}{
		"code": "ult_large",
		"name": "Unlimited 2GB",
		"price": 29.90,
		"items": 2,
	}

	e.POST("/cart").
		WithJSON(cart).
		Expect().
		Status(http.StatusOK)

}
