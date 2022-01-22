package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/tochidoh/microservices/product_api/data"
	"github.com/tochidoh/microservices/product_api/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product_api", log.LstdFlags)
	validation := data.NewValidation()

	// create handlers
	productHandler := handlers.NewProducts(logger, validation)

	// create new serve mux
	serveMux := mux.NewRouter()

	getRequest := serveMux.Methods(http.MethodGet).Subrouter()
	getRequest.HandleFunc("/products", productHandler.ListAll)
	getRequest.HandleFunc("/products/{id:[0-9]+}", productHandler.ListSingle)

	putRequest := serveMux.Methods(http.MethodPut).Subrouter()
	putRequest.HandleFunc("/products", productHandler.Update)
	putRequest.Use(productHandler.MiddlewareValidateProduct)

	postRequest := serveMux.Methods(http.MethodPost).Subrouter()
	postRequest.HandleFunc("/products", productHandler.Create)
	postRequest.Use(productHandler.MiddlewareValidateProduct)

	deleteRequest := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteRequest.HandleFunc("/products/{id:[0-9]+}", productHandler.Delete)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRequest.Handle("/docs", sh)
	getRequest.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// create new server
	server := http.Server{
		Addr:         "localhost:9090",
		Handler:      serveMux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// start server
	go func() {
		logger.Println("starting server on port 9090")

		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown server
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	signal.Notify(channel, os.Kill)

	// block until signal received
	sig := <-channel
	logger.Println("got signal:", sig)

	// gracefully shutdown server
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
