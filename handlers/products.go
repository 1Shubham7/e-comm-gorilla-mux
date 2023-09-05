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

func (p *Products) GetProducts (rw http.ResponseWriter, r *http.Request) {
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