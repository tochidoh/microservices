package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/tochidoh/microservices/product_api/handlers"
)

func main() {
	l := log.New(os.Stdout, "products_api ", log.LstdFlags)

	// create handlers
	ph := handlers.NewProducts(l) // from separate handler package that needs to be imported

	// serve mux
	sm := http.NewServeMux()
	sm.Handle("/", ph) // default paths will be handled by the product handler

	// create a server
	s := http.Server{
		Addr:         "localhost:8080",
		Handler:      sm, // serve mux is default handler
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// go routine to start server separate thread
	go func() {
		l.Println("starting server on port 8080")

		err := s.ListenAndServe() // should block unless something bad happens
		if err != nil {
			l.Printf("error starting server: %s\n", err)
			os.Exit(1) // 1 means error exit
		}
	}()

	// graceful shutdown
	c := make(chan os.Signal, 1)   // a channel of size 1 that contains os signal types
	signal.Notify(c, os.Interrupt) // pushes signal to chan
	signal.Notify(c, os.Kill)

	// block until signal
	sig := <-c
	log.Println("got signal", sig)

	// context
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second) // cancel case?
	s.Shutdown(ctx)                                                     // shutdown server
}
