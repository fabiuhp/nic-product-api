package main

import (
	"log"
	"net/http"
	"nic-product-api/handlers"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)
	hh := handlers.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe("8080", sm)
}
