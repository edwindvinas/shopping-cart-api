//This is for Scenario 1
/*Scenario 	Items Added 			Expected Cart Total 		Expected Cart Items
1 			3 x Unlimited 1 GB	$94.70					3 x Unlimited 1 GB
			1 x Unlimited 5 GB							1 x Unlimited 5 GB
*/

package shoppingCart

import (
	"net/http"
	"testing"
	"github.com/gavv/httpexpect"
)

//Add 3 items of ult_small Unlimited 1GB $24.90
func TestCartScenario(t *testing.T) {
	
	e := httpexpect.New(t, API_URL)
	
	printComment("SC10001", "Test Add 3 items of ult_small Unlimited 1GB $24.90")
	cart := map[string]interface{}{
		"code": "ult_small",
		"name": "Unlimited 1GB",
		"price": 24.90,
		"items": 3,
	}

	e.POST("/cart").
		WithJSON(cart).
		Expect().
		Status(http.StatusOK)

}

//Add 1 items of ult_large Unlimited 5GB $44.90
func TestCartScenario2(t *testing.T) {
	
	e := httpexpect.New(t, API_URL)
	
	printComment("SC10001", "Test Add 1 items of ult_large Unlimited 5GB $44.90")
	cart := map[string]interface{}{
		"code": "ult_large",
		"name": "Unlimited 5GB",
		"price": 44.90,
		"items": 1,
	}

	e.POST("/cart").
		WithJSON(cart).
		Expect().
		Status(http.StatusOK)

}
