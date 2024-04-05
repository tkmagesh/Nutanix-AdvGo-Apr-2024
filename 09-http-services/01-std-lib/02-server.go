package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

// library

type Middleware func(http.HandlerFunc) http.HandlerFunc

type AppServer struct {
	routes      map[string]http.HandlerFunc
	middlewares []Middleware
}

func (appServer *AppServer) Register(pattern string, handler http.HandlerFunc) {
	for _, middleware := range appServer.middlewares {
		handler = middleware(handler)
	}
	appServer.routes[pattern] = handler
}

func (appServer *AppServer) Use(middleware Middleware) {
	appServer.middlewares = append(appServer.middlewares, middleware)
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes: make(map[string]http.HandlerFunc),
	}
}

func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resourcePath := r.URL.Path
	if handler, exists := appServer.routes[resourcePath]; exists {
		handler(w, r)
		return
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}

// application specific
func IndexHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IndexHandler invoked")
	time.Sleep(2 * time.Second)
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

// cross cutting concerns
func loggerMiddleware(original http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v %s %s\n", time.Now(), r.Method, r.URL.Path)
		original(w, r)
	}
}

func profileMiddleware(original http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		original(w, r)
		elapsed := time.Since(start)
		fmt.Println("Time Taken :", elapsed)
	}
}

func main() {
	server := NewAppServer()
	server.Use(profileMiddleware)
	server.Use(loggerMiddleware)
	server.Register("/", IndexHanlder)
	server.Register("/products", ProductsHanlder)
	server.Register("/customers", CustomersHandler)
	http.ListenAndServe(":8080", server)
}
