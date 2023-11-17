 package main

 import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/abdullahelwalid/go-bookstore-api/pkg/routes"
 )

 func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))
 }
