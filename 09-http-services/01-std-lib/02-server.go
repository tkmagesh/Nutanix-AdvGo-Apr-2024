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
	routes map[string]func(http.ResponseWriter, *http.Request)
}

func (appServer *AppServer) Register(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	appServer.routes[pattern] = handler
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	resourcePath := r.URL.Path
	if handler, exists := appServer.routes[resourcePath]; exists {
		handler(w, r)
		return
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}

// application specific
func IndexHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func ProductsHanlder(w http.ResponseWriter, r *http.Request) {
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
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers will be served")
}

func main() {
	server := NewAppServer()
	server.Register("/", IndexHanlder)
	server.Register("/products", ProductsHanlder)
	server.Register("/customers", CustomersHandler)
	http.ListenAndServe(":8080", server)
}
