package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	authServiceURL, _ := url.Parse("http://auth-service:8081")
	orderServiceURL, _ := url.Parse("http://order-service:8082")

	authProxy := httputil.NewSingleHostReverseProxy(authServiceURL)
	orderProxy := httputil.NewSingleHostReverseProxy(orderServiceURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/users") {
			authProxy.ServeHTTP(w, r)
			return
		}
		if strings.HasPrefix(r.URL.Path, "/orders") {
			orderProxy.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Service not found")
	})

	fmt.Println("API Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
