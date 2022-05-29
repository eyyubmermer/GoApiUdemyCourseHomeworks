package main

import (
	. "eyyub/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Application started.")

	r := mux.NewRouter()

	r.HandleFunc("/api/products", PostProductHandler).Methods("POST")
	r.HandleFunc("/api/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", GetProductHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", PutProductHandler).Methods("PUT")
	r.HandleFunc("/api/products/{id}", DeleteProductHandler).Methods("Delete")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()

}
