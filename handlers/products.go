// Package classificationof Product API
// 
// Documentation for Product API
// 
// Schemes: http
// BasePath: /
// Version: 1.0.0
// 
// Consumes:
//  -application/json
// 
// Produces:
//  -application/json
// swagger:meta

package handlers

import (
	"log"
	"net/http"
	"context"
	"strconv"
	"github.com/gorilla/mux"
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

func (p *Products) AddProduct (rw http.ResponseWriter, r *http.Request){
	p.l.Println("POST Request activated")

	product := r.Context().Value(KeyProduct{}).(data.Product)
	// adding it to our fake database
	data.AddProductToDatabase(&product)
}

func (p Products ) UpdateProducts (rw http.ResponseWriter, r *http.Request){
	
	vars := mux.Vars(r) //r is the http.request
	id, err := strconv.Atoi(vars["id"]) // this is how we get ID
	//strconv.Atoi is used to convert a string representation of an integer into an actual integer value

	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("PUT Request activated", id)
	product := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &product)
	if err == data.ErrProductNotFound{
		http.Error(rw, "Product not found", http.StatusNotFound)
		return 
	}

	if err != nil{
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

// func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
// 	id := getProductID(r)

// 	p.l.Println("[DEBUG] deleting record id", id)

// 	err := data.DeleteProduct(id)
// 	if err == data.ErrProductNotFound {
// 		p.l.Println("[ERROR] deleting record id does not exist")

// 		rw.WriteHeader(http.StatusNotFound)
// 		data.ToJSON(&GenericError{Message: err.Error()}, rw)
// 		return
// 	}

// 	if err != nil {
// 		p.l.Println("[ERROR] deleting record", err)

// 		rw.WriteHeader(http.StatusInternalServerError)
// 		data.ToJSON(&GenericError{Message: err.Error()}, rw)
// 		return
// 	}

// 	rw.WriteHeader(http.StatusNoContent)
// }

type KeyProduct struct{}

func (p Products) MiddlewareProductValication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r*http.Request){
		product := data.Product{}
		err := product.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "can't decode data from JSON", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, product) //defining a context
		request := r.WithContext(ctx)
		next.ServeHTTP(rw, request) //next is just a http Handler
		//ServeHTTP() simply calls the code inside of the original function.
		
	})
}