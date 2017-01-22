# shopping-cart-api
Golang shopping cart with REST API and support for rules on discounts, promos, bundles, and comes with Swagger documentations

![enter image description here](http://lh3.googleusercontent.com/HVtYL9NyyPxsXggIV3s0JcF851OAtpODZtaobNfp425PTJQJx_qUKEEGkWQCm1dZrqGwLXoycKnhhyM9l7kGQwrP47EP7Q)

## FEATURES
### Works as a REST API for Shopping Cart (backend)
#### Current version only covers API endpoints with automated testing of APIs via go test or curl
#### Front-end does not exist yet (TODO in the future)
### Ready to test and install in Google Appengine
### Swagger API Documentation is provided
* [Sample Swagger Doc](https://ulapph-cloud-sca.appspot.com/api/)
* It was not used to generate the server side code; only used for testing and documentation
* See cart/static/api/swagger.json

### Available Rules
This shopping cart is API driver and also Rules driven though the implementation will be improved further to make the rules outside of the code. More rules will be added or if you have any suggestions, please let me know.

#### Rules for BuyThreePayTwoOnly
* If you buy 3 items, you pay only two items
* In reality this could come from a database or external config which can be updated anytime
##### Scenario: A 3 for 2 deal on Unlimited 1GB Sims. So for example, if you buy 3 Unlimited 1GB Sims, you will pay the price of 2 only for the first month.

    var Rule_BuyThreePayTwoOnly = map[string]bool{"ult_small":true,
    }

#### Rules for BulkDiscountMoreThanThree
* The price will drop to $$ each for the first month, if the customer buys more than x items.
* In reality this could come from a database or external config which can be updated anytime
##### Scenario: The Unlimited 5GB Sim will have a bulk discount applied; whereby the price will drop to $39.90 each for the first month, if the customer buys more than 3.

    var Rule_BulkDiscountMoreThanThree = map[string]float64{"ult_large":39.90,
    }

#### Rules for BundleFreeForEveryItemBought
* We will bundle in a free item X free-of-charge with every Y sold
* In reality this could come from a database or external config which can be updated anytime
##### Scenario: We will bundle in a free 1 GB Data-pack free-of-charge with every Unlimited 2GB sold.

    var Rule_BundleFreeForEveryItemBought = map[string]string{"ult_medium":"1gb",
    }

#### Rules for PromoCodeDiscount
* Adding the promo code X will apply a $$ discount across the board.
* In reality this could come from a database or external config which can be updated anytime
##### Scenario: Adding the promo code 'I<3AMAYSIM' will apply a 10% discount across the board.

    var Rule_PromoCodeDiscount = map[string]float64{"I<3AMAYSIM":10,
    }
	
### Handlers
Below are the handlers which serves as the handlers for REST API calls. These are currently hand-coded 100% but in the future we will use Swagger.JSON to automatically generate some of these server codes.

	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/products", handleFuncProducts)
	http.HandleFunc("/cart", handleFuncCart)
	http.HandleFunc("/process", handleFuncProcess)

### Functions
These are the major functions used by the shopping cart api.

#### func handleFuncProducts(w http.ResponseWriter, r *http.Request) {}
Handles management of products using POST, GET, DELETE 

#### func handleFuncCart(w http.ResponseWriter, r *http.Request) {}
Handles management of shopping cart using POST, GET, DELETE

#### func handleFuncProcess(w http.ResponseWriter, r *http.Request) {}
Handles pre-checkout and post-checkout processing and does most of the rules engine processing, application of discounts, bundles etc.

##### func (s *CartProc) Check_Rule_BuyThreePayTwoOnly(w http.ResponseWriter, r *http.Request) error {}
The rules engine for Buy Three Pay Two Only rule

##### func (s *CartProc) Check_Rule_BulkDiscountMoreThanThree(w http.ResponseWriter, r *http.Request) error {}
The rules engine for the Bulk Discount More Than Three rule

##### func (s *CartProc) Check_Rule_BundleFreeForEveryItemBought(w http.ResponseWriter, r *http.Request) error {}
The rules engine for the Bundle Free for Every Item Bought

##### func (s *CartProc) Check_Rule_PromoCodeDiscount(w http.ResponseWriter, r *http.Request, promo string) error {}
The rules engine for the Promo Code Discounts

##### func (s *CartProc) All_Others_No_Rule_Processing(w http.ResponseWriter, r *http.Request) error {}
If no rules apply to other cart items, this logic is being applied


## INSTALLATION ON GOOGLE APPENGINE

The shopping cart api can be deployed to your appengine project in less 10 minutes, provided you have prior experience with GAE. Otherwise, please consult Google Appengine docs first.

# Create your Appengine project ID
* Then, update app.yaml to indicate your new appengine project ID
* Execute the updategae.sh script

### Install script

    ./updategae.sh

### Visit your appengine project (demo is shown below)
* https://ulapph-cloud-sca.appspot.com 

### View Swagger API Documentation
* https://ulapph-cloud-sca.appspot.com/api/

## AUTOMATED TESTING VIA GO TEST AND CURL

#### Rules for BuyThreePayTwoOnly
##### Scenario: A 3 for 2 deal on Unlimited 1GB Sims. So for example, if you buy 3 Unlimited 1GB Sims, you will pay the price of 2 only for the first month.

    cd test-cases/cart-test-scenario1
    go test -v
	//after this command, all tables are refreshed with data
    [CLR0001] Test clear the database tables in appengine --> Products
    [CLR0002] Test clear the database tables in appengine --> Cart
    [ADD0001] Test add product 1 to catalog --> ult_small Unlimited 1GB $24.90
    [ADD0002] Test add product 2 to catalog --> ult_medium Unlimited 2GB $29.90
    [ADD0003] Test add product 3 to catalog --> ult_large Unlimited 5GB $44.90
    [ADD0004] Test add product 4 to catalog --> 1gb 1 GB Data-pack $9.90
    [SC10001] Test Add 3 items of ult_small Unlimited 1GB $24.90
    [SC10001] Test Add 1 items of ult_large Unlimited 5GB $44.90
    PASS
    ok      github.com/edwindvinas/shopping-cart-api/cart-test-scenario1    3.266s

	//execute curl to simulate checkout of cart
    $ curl -X GET --header 'Accept: application/json' \
		    'http://ulapph-cloud-sca.appspot.com/process'
    {
      "total": 94.69999999999999,
      "rules": {
        "ult_small": "Rule_BuyThreePayTwoOnly"
      },
      "current": [
        {
          "code": "ult_large",
          "name": "Unlimited 5GB",
          "price": 44.9,
          "items": 1,
          "status": false
        },
        {
          "code": "ult_small",
          "name": "Unlimited 1GB",
          "price": 24.9,
          "items": 3,
          "status": false
        }
      ]
    }

#### Rules for BulkDiscountMoreThanThree
##### Scenario: The Unlimited 5GB Sim will have a bulk discount applied; whereby the price will drop to $39.90 each for the first month, if the customer buys more than 3.
    cd test-cases/cart-test-scenario2
    go test -v
	//after this command, all tables are refreshed with data
    [CLR0001] Test clear the database tables in appengine --> Products
    [CLR0002] Test clear the database tables in appengine --> Cart
    [ADD0001] Test add product 1 to catalog --> ult_small Unlimited 1GB $24.90
    [ADD0002] Test add product 2 to catalog --> ult_medium Unlimited 2GB $29.90
    [ADD0003] Test add product 3 to catalog --> ult_large Unlimited 5GB $44.90
    [ADD0004] Test add product 4 to catalog --> 1gb 1 GB Data-pack $9.90
    [SC20001] Test Add 2 items of Unlimited 1 GB for $24.90
    [SC20001] Test Add 4 items of Unlimited 5 GB for $209.40
    PASS
    ok      github.com/edwindvinas/shopping-cart-api/cart-test-scenario2    3.068s

	//execute curl to simulate checkout of cart
    $ curl -X GET --header 'Accept: application/json' \
			    'http://ulapph-cloud-sca.appspot.com/process'
    {
      "total": 209.39999999999998,
      "rules": {
        "ult_large": "Rule_BulkDiscountMoreThanThree"
      },
      "current": [
        {
          "code": "ult_large",
          "name": "Unlimited 5GB",
          "price": 44.9,
          "items": 4,
          "status": false
        },
        {
          "code": "ult_small",
          "name": "Unlimited 1GB",
          "price": 24.9,
          "items": 2,
          "status": false
        }
      ]
    }

#### Rules for BundleFreeForEveryItemBought
##### Scenario: We will bundle in a free 1 GB Data-pack free-of-charge with every Unlimited 2GB sold.
    cd test-cases/cart-test-scenario2
    go test -v
	//after this command, all tables are refreshed with data
    [CLR0001] Test clear the database tables in appengine --> Products
    [CLR0002] Test clear the database tables in appengine --> Cart
    [ADD0001] Test add product 1 to catalog --> ult_small Unlimited 1GB $24.90
    [ADD0002] Test add product 2 to catalog --> ult_medium Unlimited 2GB $29.90
    [ADD0003] Test add product 3 to catalog --> ult_large Unlimited 5GB $44.90
    [ADD0004] Test add product 4 to catalog --> 1gb 1 GB Data-pack $9.90
    [SC30001] Test Add 1 item of  Unlimited 1 GB for $24.90
    [SC30001] Test Add 2 items of Unlimited 2GB for $29.90
    PASS
    ok      github.com/edwindvinas/shopping-cart-api/cart-test-scenario3    2.837s
    
    //execute curl to simulate checkout of cart
      $ curl -X GET --header 'Accept: application/json' \
			      'http://ulapph-cloud-sca.appspot.com/process'
    {
      "total": 84.69999999999999,
      "rules": {},
      "current": [
        {
          "code": "ult_large",
          "name": "Unlimited 2GB",
          "price": 29.9,
          "items": 2,
          "status": false
        },
        {
          "code": "ult_small",
          "name": "Unlimited 1GB",
          "price": 24.9,
          "items": 1,
          "status": false
        }
      ]
    }

#### Rules for PromoCodeDiscount
##### Scenario: Adding the promo code 'I<3AMAYSIM' will apply a 10% discount across the board.
    cd test-cases/cart-test-scenario2
    go test -v
	//after this command, all tables are refreshed with data
    [CLR0001] Test clear the database tables in appengine --> Products
    [CLR0002] Test clear the database tables in appengine --> Cart
    [ADD0001] Test add product 1 to catalog --> ult_small Unlimited 1GB $24.90
    [ADD0002] Test add product 2 to catalog --> ult_medium Unlimited 2GB $29.90
    [ADD0003] Test add product 3 to catalog --> ult_large Unlimited 5GB $44.90
    [ADD0004] Test add product 4 to catalog --> 1gb 1 GB Data-pack $9.90
    [SC40001] Test Add 1 item of  Unlimited 1 GB for $24.90
    [SC40001] Test Add 1 item of 1 GB Data-pack $9.90
    PASS
    ok      github.com/edwindvinas/shopping-cart-api/cart-test-scenario4    2.955s

	//execute curl to simulate checkout of cart
    $ curl -X GET --header 'Accept: application/json' 'http://ulapph-cloud-sca.appspot.com/process?promo_code=I%3C3AMAYSIM'
    {
      "total": 31.319999999999997,
      "rules": {
        "I\u003c3AMAYSIM": "Rule_PromoCodeDiscount"
      },
      "current": [
        {
          "code": "1gb",
          "name": "1 GB Data-pack",
          "price": 9.9,
          "items": 1,
          "status": false
        },
        {
          "code": "ult_small",
          "name": "Unlimited 1GB",
          "price": 24.9,
          "items": 1,
          "status": false
        }
      ]
    }

## TODOs
### Swagger to generate server-side codes automatically
* This is the swagger file which can be used to automatically generate the API server (in the future)
* See https://github.com/go-swagger/go-swagger for more info
* In reality, we shouldn't be coding the API server by hand coz go-swagger takes the swagger.json as input and it generates the API endpoints skeleton
* All we need to do is code the business that will handle each API call

### Create front-end using ngCart or a really nice option such as QOR
[ngCart Github](https://github.com/snapjay/ngCart)
[QOR](https://github.com/qor/qor/)

### Use external database or configuration for the rules engine
* This is to avoid recompiling the code everytime we need to update the rules

### Of course, more coding of the missing endpoints such as rules management and also very important to incorporate is the payment services such as Paypal etc


## ACKNOWLEDGMENT
### Used for end-to-end testing of REST API
* https://github.com/gavv/httpexpect
### Appengine gotodos used as basis for datastore database
* https://github.com/GoogleCloudPlatform/appengine-angular-gotodos
### Go swagger UI used as API Documentation tool
* https://github.com/go-swagger/go-swagger
* https://github.com/swagger-api/swagger-ui
### Google Appengine
* https://cloud.google.com/appengine/
### Golang
* https://golang.org/
### And of course...
* Github for being awesome!
* Notepad+++ as very reliable and fast editor!
* Git Bash for Windows
* https://stackedit.io/editor for creating this markdown for Github!


## Thank you!