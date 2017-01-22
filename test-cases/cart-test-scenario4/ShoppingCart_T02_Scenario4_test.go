//This is for Scenario 4
/*Scenario 	Items Added 			Expected Cart Total 		Expected Cart Items
4 			1 x Unlimited 1 GB	$31.32					1 x Unlimited 1 GB
			1 x 1 GB Data-pack							1 x 1 GB Data-pack
			'I<3AMAYSIM' Promo Applied
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
	
	printComment("SC40001", "Test Add 1 item of  Unlimited 1 GB for $24.90")
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

//Add 1 item of 1 GB Data-pack $9.90
func TestCartScenario2(t *testing.T) {
	
	e := httpexpect.New(t, API_URL)
	
	printComment("SC40001", "Test Add 1 item of 1 GB Data-pack $9.90")
	cart := map[string]interface{}{
		"code": "1gb",
		"name": "1 GB Data-pack",
		"price": 9.90,
		"items": 1,
	}

	e.POST("/cart").
		WithJSON(cart).
		Expect().
		Status(http.StatusOK)

}
