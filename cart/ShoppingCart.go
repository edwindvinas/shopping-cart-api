package shoppingCart

import (
	"fmt"
	"net/http"
	"encoding/json"
	"appengine/datastore"
	"appengine"
	"io"
)

//Data struct for Product table
type Product struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

//Data struct for Cart table
type Cart struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Items int64 `json:"items"`
	Status bool `json:"status"`
}

//Data struct for Total Amount and Overall Items
type CartProc struct {
	Total float64 `json:"total"`
	Rules map[string]string  `json:"rules"`
	Current []Cart `json:"current"`
}

//Rules for BuyThreePayTwoOnly
//If you buy 3 items, you pay only two items
//In reality this could come from a database or external config which can be updated anytime
var Rule_BuyThreePayTwoOnly = map[string]bool{"ult_small":true,
}

//Rules for BulkDiscountMoreThanThree
//The price will drop to $$ each for the first month, if the customer buys more than x items.
//In reality this could come from a database or external config which can be updated anytime
var Rule_BulkDiscountMoreThanThree = map[string]float64{"ult_large":39.90,
}

//Rules for BundleFreeForEveryItemBought
//We will bundle in a free item X free-of-charge with every Y sold
//In reality this could come from a database or external config which can be updated anytime
var Rule_BundleFreeForEveryItemBought = map[string]string{"ult_medium":"1gb",
}

//Rules for PromoCodeDiscount
//Adding the promo code X will apply a $$ discount across the board.
//In reality this could come from a database or external config which can be updated anytime
var Rule_PromoCodeDiscount = map[string]float64{"I<3AMAYSIM":10,
}

//Handlers
func init() {
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/products", handleFuncProducts)
	http.HandleFunc("/cart", handleFuncCart)
	http.HandleFunc("/process", handleFuncProcess)
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", rootHTML)
}

const rootHTML = `<h1>Shopping Cart API</h1>
<hr>
<h3>Note: GUI is not yet implemented. API backend is covered and so the testing done is using curl commands.</h3>
<hr>
<a href="/api/">API Documentation</a>
`

/////////////////////////////////////////////////////////////
// Products
////////////////////////////////////////////////////////////
//Handler for Products
func handleFuncProducts(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	
	c.Infof("handleFuncProducts")
	val, err := handleProducts(c, r)
	if err == nil {
		err = json.NewEncoder(w).Encode(val)
	}
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("api error: %#v", err)))
		return	
	}
}
 
func handleProducts(c appengine.Context, r *http.Request) (interface{}, error) {
	switch r.Method {
		
		case "POST":
			c.Infof("POST")
			product, err := decodeProduct(r.Body)
			if err != nil {
				c.Errorf("%v", err)
				return nil, err
			}
			return product.save(c)
			
		case "GET":
			c.Infof("GET")
			return getAllProducts(c)
			
		case "DELETE":
			c.Infof("DELETE")
			return nil, deleteProducts(c)
	}
	c.Infof("method not implemented")
	return nil, fmt.Errorf("method not implemented")
	
}
 
func defaultProductList(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "ProductList", "default", 0, nil)
}
 
func (t *Product) save(c appengine.Context) (*Product, error) {
	c.Infof("save")
	_, err := datastore.Put(c, getKeyProduct(c,t.Code), t)
	if err != nil {
		c.Errorf("%v", err)
		return nil, err
	}
	return t, nil
}
 
func decodeProduct(r io.ReadCloser) (*Product, error) {
	defer r.Close()
	var product Product
	err := json.NewDecoder(r).Decode(&product)
	return &product, err
}
 
func getAllProducts(c appengine.Context) ([]Product, error) {
	c.Infof("getAllProducts")
	products := []Product{}
	_, err := datastore.NewQuery("Product").Ancestor(defaultProductList(c)).Order("Code").GetAll(c, &products)
	if err != nil {
		c.Errorf("%v", err)
		return nil, err
	}
	return products, nil
}
 
func deleteProducts(c appengine.Context) error {
	c.Infof("deleteProducts")
	q := datastore.NewQuery("Product").KeysOnly()
	keys, err := q.GetAll(c, nil)
	if err != nil {
		return err
	}
	err = datastore.DeleteMulti(c, keys)
	if err != nil {
		return err
	}
	c.Infof("deleteProducts DeleteMulti")
	return err
}

func getKeyProduct(c appengine.Context, Code string) *datastore.Key {
	c.Infof("getKeyProduct")
	return datastore.NewKey(c, "Product", Code, 0, nil)
}

////////////////////////////////////////////
// Cart
////////////////////////////////////////////
//Handler for Cart
func handleFuncCart(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	
	c.Infof("handleFuncCarts")
	val, err := handleCarts(c, r)
	if err == nil {
		err = json.NewEncoder(w).Encode(val)
	}
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("api error: %#v", err)))
		return	
	}
}
 
func handleCarts(c appengine.Context, r *http.Request) (interface{}, error) {
	switch r.Method {
		
		case "POST":
			c.Infof("POST")
			cart, err := decodeCart(r.Body)
			if err != nil {
				c.Errorf("%v", err)
				return nil, err
			}
			return cart.save(c)
			
		case "GET":
			c.Infof("GET")
			return getAllCarts(c)
			
		case "DELETE":
			c.Infof("DELETE")
			return nil, deleteCarts(c)
	}
	c.Infof("method not implemented")
	return nil, fmt.Errorf("method not implemented")
	
}
 
func defaultCartList(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "CartList", "default", 0, nil)
}
 
func (t *Cart) save(c appengine.Context) (*Cart, error) {
	c.Infof("save")
	_, err := datastore.Put(c, getKeyCart(c,t.Code), t)
	if err != nil {
		c.Errorf("%v", err)
		return nil, err
	}
	return t, nil
}
 
func decodeCart(r io.ReadCloser) (*Cart, error) {
	defer r.Close()
	var product Cart
	err := json.NewDecoder(r).Decode(&product)
	return &product, err
}
 
func getAllCarts(c appengine.Context) ([]Cart, error) {
	c.Infof("getAllCarts")
	products := []Cart{}
	_, err := datastore.NewQuery("Cart").Ancestor(defaultCartList(c)).Order("Code").GetAll(c, &products)
	if err != nil {
		c.Errorf("%v", err)
		return nil, err
	}
	return products, nil
}
 
func deleteCarts(c appengine.Context) error {
	c.Infof("deleteCarts")
	q := datastore.NewQuery("Cart").KeysOnly()
	keys, err := q.GetAll(c, nil)
	if err != nil {
		return err
	}
	err = datastore.DeleteMulti(c, keys)
	if err != nil {
		return err
	}
	c.Infof("deleteCarts DeleteMulti")
	return err
}

func getKeyCart(c appengine.Context, Code string) *datastore.Key {
	c.Infof("getKeyCart")
	return datastore.NewKey(c, "Cart", Code, 0, nil)
}

//////////////////////////////////////
// Process
//////////////////////////////////////
//Handler for Cart
func handleFuncProcess(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	
	promo_code := r.FormValue("promo_code")
	
	//Call initiated by browser when a JS or ajax attempts ot recompute the cart items
	//Get all items in the cart and process
	//For each item which has applicable pricing rules, include them in processing
	//Return total and list of items
	
	shop := newCartProc()
	//Rules for BuyThreePayTwoOnly
	//If you buy 3 items, you pay only two items
	c.Infof("Check_Rule_BuyThreePayTwoOnly")
	shop.Check_Rule_BuyThreePayTwoOnly(w,r)

	//Rules for BulkDiscountMoreThanThree
	//The price will drop to $$ each for the first month, if the customer buys more than x items.
	c.Infof("Check_Rule_BulkDiscountMoreThanThree")
	shop.Check_Rule_BulkDiscountMoreThanThree(w,r)

	//Rules for BundleFreeForEveryItemBought
	//We will bundle in a free item X free-of-charge with every Y sold
	c.Infof("Check_Rule_BundleFreeForEveryItemBought")
	shop.Check_Rule_BundleFreeForEveryItemBought(w,r)
	
	//Process those items w/ no rules applied
	c.Infof("All_Others_No_Rule_Processing")
	shop.All_Others_No_Rule_Processing(w,r)
	
	//Finally, see if any promos to be processed
	//Check rules if there are promos to be applied
	c.Infof("promo_code: %v", promo_code)
	if promo_code != "" {
		//Rules for PromoCodeDiscount
		//Adding the promo code X will apply a $$ discount across the board.
		c.Infof("Check_Rule_PromoCodeDiscount")
		shop.Check_Rule_PromoCodeDiscount(w,r,promo_code)
	}
	
	//return data as json
	cp := CartProc {
		Total: shop.Total,
		Rules: shop.Rules,
		Current: shop.Current,
	}
	
	data,_ := json.MarshalIndent(cp, "", "  ")
	w.Write(data)
	
}

//Initialize cart proc
func newCartProc() *CartProc {
	return &CartProc{
		Total: 0,
		Rules: map[string]string{},
		Current: nil,
	}
}

func (s *CartProc) Check_Rule_BuyThreePayTwoOnly(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	
	c.Infof("Check_Rule_BuyThreePayTwoOnly")
	//For each entry in rule, scan the cart items
	
	for k, v := range Rule_BuyThreePayTwoOnly {
		c.Infof("key: %v -> value: %v", k, v)
		q := datastore.NewQuery("Cart").Filter("Code =", k)

		recCount, _  := q.Count(c)
		items := make([]Cart, 0, recCount)
		if _, err := q.GetAll(c, &items); err != nil {
			return err
		 }
		
		var total float64
		applied := false
		for _, p := range items {
			c.Infof("cart item: %v", p)
			if p.Items >= 3 {
				//temp; need to fix this logic; fixed to 3 only for now; what if cx has 6 or 9 items etc
				total = float64((p.Items - 1)) * p.Price
				applied = true
			}
		}
		s.Total = s.Total + total
		if applied == true {
			s.Rules[k] = "Rule_BuyThreePayTwoOnly"
		}
		
	}
	return nil
}

func (s *CartProc) Check_Rule_BulkDiscountMoreThanThree(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	
	c.Infof("Check_Rule_BulkDiscountMoreThanThree")
	//For each entry in rule, scan the cart items

	for k, v := range Rule_BulkDiscountMoreThanThree {
		c.Infof("key: %v -> value: %v", k, v)
		q := datastore.NewQuery("Cart").Filter("Code =", k)

		recCount, _  := q.Count(c)
		items := make([]Cart, 0, recCount)
		if _, err := q.GetAll(c, &items); err != nil {
			return err
		 }
		
		var total float64
		applied := false
		for _, p := range items {
			c.Infof("cart item: %v", p)
			if p.Items >= 3 {
				discPrice := v
				total = float64(p.Items) * discPrice
				s.Total = s.Total + total
				applied = true
			}
		}
		if applied == true {
			s.Rules[k] = "Rule_BulkDiscountMoreThanThree"
		}
		

	}
	
	return nil
}

func (s *CartProc) Check_Rule_BundleFreeForEveryItemBought(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	
	c.Infof("Check_Rule_BundleFreeForEveryItemBought")
	//For each entry in rule, scan the cart items
	
	for k, v := range Rule_BundleFreeForEveryItemBought {
		c.Infof("key: %v -> value: %v", k, v)
		q := datastore.NewQuery("Cart").Filter("Code =", k)

		recCount, _  := q.Count(c)
		items := make([]Cart, 0, recCount)
		if _, err := q.GetAll(c, &items); err != nil {
			return err
		 }
		
		promCtr := 0
		applied := false
		for _, p := range items {
			c.Infof("cart item: %v", p)
			promCtr++
			applied = true
		}
		c.Infof("cart item promCtr: %v", promCtr)
			
		if applied == true {
			//get item details for this promo item
			err := AddPromoItemToCart(w,r,v)
			if err != nil {
				c.Infof("AddPromoItemToCart error: %v", v)
				return nil
			}
			s.Rules[k] = "Rule_BulkDiscountMoreThanThree"
		}
		

	}
	
	return nil
}

func (s *CartProc) Check_Rule_PromoCodeDiscount(w http.ResponseWriter, r *http.Request, promo string) error {
	c := appengine.NewContext(r)
	
	c.Infof("Check_Rule_PromoCodeDiscount")
	//For each entry in rule, scan the cart items
	
	for k, v := range Rule_PromoCodeDiscount {
		c.Infof("key: %v -> value: %v", k, v)
		
		_, ok := Rule_PromoCodeDiscount[promo]
		if ok == true {
			//promo should be applied
			discount := v/100
			c.Infof("s.Total: %v", s.Total)
			s.Total = s.Total - (s.Total * discount)
			c.Infof("discount: %v", discount)
			c.Infof("s.Total(w/ promo): %v", s.Total)
			s.Rules[k] = "Rule_PromoCodeDiscount"
		}
	}
	
	return nil
}

func (s *CartProc) All_Others_No_Rule_Processing(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	
	c.Infof("All_Others_No_Rule_Processing")
	//For all other entries w/ no rules applied, process normally

		q := datastore.NewQuery("Cart").Order("Code")

		recCount, _  := q.Count(c)
		items := make([]Cart, 0, recCount)
		if _, err := q.GetAll(c, &items); err != nil {
			return err
		 }
		
		for _, p := range items {
			c.Infof("cart item: %v", p)
			_, ok := s.Rules[p.Code]
			if ok == false {
				//normal process this item
				s.Total = s.Total + (p.Price* float64(p.Items))
			}
		}
		//append to cart current items
		s.Current = items
		
	
	return nil
}

func AddPromoItemToCart(w http.ResponseWriter, r *http.Request, code string) (err error) {
	c := appengine.NewContext(r)
	
	c.Infof("AddPromoItemToCart")
	c.Infof("AddPromoItemToCart code: %v", code)
	q := datastore.NewQuery("Product").Filter("Code =", code).Limit(1)
	
	recCount, _  := q.Count(c)
	items := make([]Product, 0, recCount)
	if _, err := q.GetAll(c, &items); err != nil {
		return err
	}
	
	for _, p := range items {
		c.Infof("cart item: %v", p)
		//insert new promo item to cart
		
		cart := Cart {
			Code: p.Code,
			Name: p.Name,
			Price: 0,
			Items: 1,			
		}

		cart.save(c)	
	}
	
	return err
}


