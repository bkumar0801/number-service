package main

import (
	"flag"
	"log"
	"net/http"

	h "taudience.com/number-service/handlers"
)

func main() {
	listenAddr := flag.String("http.addr", ":8070", "http listen address")
	flag.Parse()

	http.HandleFunc("/numbers", h.HandleNumbers)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
