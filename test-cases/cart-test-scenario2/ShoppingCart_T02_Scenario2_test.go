//This is for Scenario 2
/*Scenario 	Items Added 			Expected Cart Total 		Expected Cart Items
2 			2 x Unlimited 1 GB	$209.40					2 x Unlimited 1 GB
			4 x Unlimited 5 GB							4 x Unlimited 5 GB
*/

package shoppingCart

import (
	"net/http"
	"testing"
	"github.com/gavv/httpexpect"
)

//Add 2 items of Unlimited 1 GB for $24.90
func TestCartScenario(t *testing.T) {
	
	e := httpexpect.New(t, API_URL)
	
	printComment("SC20001", "Test Add 2 items of Unlimited 1 GB for $24.90")
	cart := map[string]interface{}{
		"code": "ult_small",
		"name": "Unlimited 1GB",
		"price": 24.90,
		"items": 2,
	}

	e.POST("/cart").
		WithJSON(cart).
		Expect().
		Status(http.StatusOK)

}

//Add 4 items of Unlimited 5 GB for $209.40
func TestCartScenario2(t *testing.T) {
	
	e := httpexpect.New(t, API_URL)
	
	printComment("SC20001", "Test Add 4 items of Unlimited 5 GB for $209.40")
	cart := map[string]interface{}{
		"code": "ult_large",
		"name": "Unlimited 5GB",
		"price": 44.90,
		"items": 4,
	}

	e.POST("/cart").
		WithJSON(cart).
		Expect().
		Status(http.StatusOK)

}
