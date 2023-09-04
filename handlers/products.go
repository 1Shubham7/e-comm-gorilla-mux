package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/1shubham7/e-comm/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l*log.Logger) *Products {
	return &Products{l}
}

// anything adds to the handler interface  needs to have a serve http method
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// we will now use encoding/json package to convert the struct we ahve into a json representative
	// listOfProducts := data.GetProducts()
	// now listOfProducts contains the list of products, but how do we return that to the user
	// we do that by converting listOfProducts into json. and this the the way

	// d, err := json.Marshal(listOfProducts)
	// instead of Marshal we are using this better alternative

	// err := listOfProducts.ToJSON(rw) //ToJSON is in data/products
	// if err != nil {
		// http.Error(rw, "unable to convert data to json", http.StatusInternalServerError)
	// }
	// rw.Write(d) instead of Marshal we are using this better alternative



	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost{
		p.addProduct(rw, *r) // in tutorial it was just r instead of *r
		return
	}

	if r.Method == http.MethodPut{
		p.l.Println("PUT Method activated")
		// we will need the id from the url provided by user

		// now we are extracting the id from the path
		
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, - 1) 

		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts (rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET Products activated")
	listOfProducts := data.GetProducts()
	err := listOfProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert data to JSON", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct (rw http.ResponseWriter, r http.Request){
	p.l.Println("POST Request activated")

	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "can't decode data from JSON", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", product) //use %# for better representation than just %

	// adding it to our fake database
	data.AddProductToDatabase(product)
}

func (p Products ) updateProducts (id int, rw http.ResponseWriter, r *http.Request){
	p.l.Println("PUT Request activated")

	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "can't decode data from JSON", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound{
		http.Error(rw, "Product not found", http.StatusNotFound)
		return 
	}

	if err != nil{
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}