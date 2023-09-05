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
	// "io"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// handlers
	producthandler := handlers.NewProducts(l)

	servemux := mux.NewRouter()
	getRouter := servemux.Methods("GET").Subrouter() // this creates a route for a perticular http "GET" request and then .subrouter makes it a router again so that you can use handler in it.
	getRouter.HandleFunc("/", producthandler.GetProducts)
	// servemux.Handle("/", producthandler)

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
			l.Fatal(err)
		}
	} ()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Performing Graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
	// http.ListenAndServe(":6000", servemux)
	// second parameter is for http handler. if we say "nil" in second parameter, 
	// the server will take the default http handler, here we specified the http handler
}
