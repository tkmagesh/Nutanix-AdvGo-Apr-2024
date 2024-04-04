package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Cost     float32 `json:"cost"`
	Category string  `json:"-"`
}

var products []Product = []Product{
	{100, "Pen", 10, "Stationary"},
	{101, "Pencil", 5, "Stationary"},
	{102, "Marker", 50, "Stationary"},
}

type AppServer struct {
}

func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello there!\n"))
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello World!")
	case "/products":
		// fmt.Fprintln(w, "All the products will be served")
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "error serializing the data", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, "error in request body", http.StatusBadRequest)
			}
			products = append(products, newProduct)
			w.WriteHeader(http.StatusCreated)
		default:
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		}
	case "/customers":
		fmt.Fprintln(w, "All the customers will be served")
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
	}
}

func main() {
	server := &AppServer{}
	http.ListenAndServe(":8080", server)
}
