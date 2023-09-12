package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/1shubham7/e-comm/handlers"
	"github.com/gorilla/mux"
	// "github.com/1shubham7/e-comm/env"
	// "io"
)

// var bindAddress = env.String("BIND_ADDRESS", false, ":9090","Bind adddress for the server")

func main() {

	// env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// handlers
	producthandler := handlers.NewProducts(l)

	servemux := mux.NewRouter()
	getRouter := servemux.Methods("GET").Subrouter() // this creates a route for a perticular http "GET" request and then .subrouter makes it a router again so that you can use handler in it.
	getRouter.HandleFunc("/", producthandler.GetProducts)

	putRouter := servemux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", producthandler.UpdateProducts) // this thing inside url is called regex, mux willl automatically extract this data, it means anything in 0-9 and + means one or more
	//here we just gave the id of the product we want to update

	// a subrouter has use() func that takes middleware therefore
	putRouter.Use(producthandler.MiddlewareProductValication)

	postRouter := servemux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", producthandler.AddProduct)
	postRouter.Use(producthandler.MiddlewareProductValication)

	server := &http.Server{
		Addr: ":6000",
		Handler: servemux,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 5 *time.Second,
		WriteTimeout: 10 *time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	} ()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	log.Println("Performing Graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
	// http.ListenAndServe(":6000", servemux)
	// second parameter is for http handler. if we say "nil" in second parameter, 
	// the server will take the default http handler, here we specified the http handler
}