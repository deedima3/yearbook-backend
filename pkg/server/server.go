package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

// Server acts as an object that will listens to http request
type Server struct {
	Address string
	Handler http.Handler
}

// ListenAndServe listens to http network with given address and handler. This method blocks the execution
func (s Server) ListenAndServe() {
	go func() {
		if err := http.ListenAndServe(s.Address, s.Handler); err != nil {
			log.Printf("ERROR (starting server): %v\n", err)
		}
	}()
	log.Printf("INFO (server): server started, listening to %s\n", s.Address)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("INFO (server): shutting down server")
}

// ProvideServer acts as a Factory Design Pattern to provide Server struct object in a consistent manner
func ProvideServer(address string, handler http.Handler) Server {
	return Server{Address: address, Handler: handler}
}
